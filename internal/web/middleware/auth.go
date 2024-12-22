package middleware

import "net/http"

func (m *MiddleWareManager) AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// place data verification for user authorization here.
		// Checking the session or jwt token

		// You can place the user ID in the context of the request to
		// work with it in a future router

		// example:
		// next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "ctxKey", "userID")))
		next.ServeHTTP(w, r)
	})
}
