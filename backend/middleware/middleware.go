package middleware

import (
	"backend/internal/entities"
	"backend/token"
	"context"
	"net/http"
	"strings"
)

type ctxKey string

const claimsKey ctxKey = "claims"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, "Missing or Invalid token", http.StatusUnauthorized)
			return
		}
		claims, err := token.VerifyToken(strings.TrimPrefix(header, "Bearer "))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), claimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetClaims(ctx context.Context) (*entities.ClaimsData, bool) {
	c, ok := ctx.Value(claimsKey).(*entities.ClaimsData)
	return c, ok
}
