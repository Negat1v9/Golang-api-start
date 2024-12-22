package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type MiddleWareManager struct{}

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := 0; i < len(xs); i++ {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}
