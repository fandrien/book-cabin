package middleware

import (
	"net/http"

	"github.com/fandrien/book-cabin/constant"
	"github.com/fandrien/book-cabin/response"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(10, 20)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		if !limiter.Allow() {
			response.WriteError(
				w,
				http.StatusTooManyRequests,
				constant.ErrRateLimitExceeded,
				"too many requests",
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}
