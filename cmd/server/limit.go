package main

import (
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

// The rate itself is set in server.go.
const (
	// Bounds the limiter's memory. Past it the table is emptied: a fresh
	// allowance for everyone costs less than an unbounded map.
	maxBuckets = 100000
	// Clients idle for idleAfter are forgotten once the table holds more
	// than sweepAbove, at most every sweepEvery so that a flood of new
	// addresses cannot make every request pay for a full scan.
	sweepAbove = 10000
	sweepEvery = 10 * time.Second
	idleAfter  = time.Minute
	// One subscriber is handed a /64, so that is one client.
	ipv6ClientBits = 64
)

// visitor is the address a request comes from, or nil when it cannot be
// read. With -trust-proxy it is the last X-Forwarded-For entry, the one the
// proxy itself appended: any earlier entry is the client's own claim.
func (s *server) visitor(r *http.Request) net.IP {
	if s.trustProxy {
		// The proxy's own entry is the last one, whether it appended to the
		// client's header line or added a line of its own.
		if lines := r.Header.Values("X-Forwarded-For"); len(lines) > 0 {
			parts := strings.Split(lines[len(lines)-1], ",")
			if ip := net.ParseIP(strings.TrimSpace(parts[len(parts)-1])); ip != nil {
				return ip
			}
		}
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}
	return net.ParseIP(host)
}

// clientIP is the key requests are rate-limited by.
func (s *server) clientIP(r *http.Request) string {
	if ip := s.visitor(r); ip != nil {
		return limiterKey(ip)
	}
	return r.RemoteAddr
}

// limiterKey makes one IPv4 address, or one IPv6 /64, one client.
func limiterKey(ip net.IP) string {
	if v4 := ip.To4(); v4 != nil {
		return v4.String()
	}
	return ip.Mask(net.CIDRMask(ipv6ClientBits, 128)).String()
}

// limiter is an in-memory token bucket per client.
type limiter struct {
	mu      sync.Mutex
	burst   float64
	perSec  float64
	buckets map[string]*bucket
	swept   time.Time
}

type bucket struct {
	tokens float64
	seen   time.Time
}

func newLimiter(burst, perSec float64) *limiter {
	return &limiter{burst: burst, perSec: perSec, buckets: map[string]*bucket{}}
}

// allow reports whether the client may make a request now, and counts it.
func (l *limiter) allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()
	l.forgetIdle(now)
	b := l.buckets[key]
	if b == nil {
		b = &bucket{tokens: l.burst, seen: now}
		l.buckets[key] = b
	}
	b.tokens = min(l.burst, b.tokens+now.Sub(b.seen).Seconds()*l.perSec)
	b.seen = now
	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}

// forgetIdle keeps the table within its bounds. The caller holds l.mu.
func (l *limiter) forgetIdle(now time.Time) {
	if len(l.buckets) > sweepAbove && now.Sub(l.swept) > sweepEvery {
		l.swept = now
		for k, b := range l.buckets {
			if now.Sub(b.seen) > idleAfter {
				delete(l.buckets, k)
			}
		}
	}
	if len(l.buckets) >= maxBuckets {
		clear(l.buckets)
	}
}
