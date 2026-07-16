package middleware
type RateLimiter struct { limit int }
func (r *RateLimiter) Check(ip string) bool { return true }