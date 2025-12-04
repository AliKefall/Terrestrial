package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/AliKefall/DonemOdevi/internal/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type contextKey string

const userIDContextKey contextKey = "userID"

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	val := ctx.Value(userIDContextKey)
	if val == "" {
		return uuid.Nil, errors.New("userID could not be found in context!")

	}
	id, ok := val.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("userID type error!")
	}
	return id, nil
}

func AuthMiddleware(tokenSecret string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := auth.GetBearerToken(r.Header)
			if err != nil {
				RespondWithError(w, http.StatusUnauthorized, "Token could not be found", err)
				return
			}
			claims := &jwt.RegisteredClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
				return []byte(tokenSecret), nil
			})

			if err != nil || !token.Valid {
				RespondWithError(w, http.StatusUnauthorized, "Token is invalid", err)
				return
			}

			userID, err := uuid.Parse(claims.Subject)
			if err != nil {
				RespondWithError(w, http.StatusUnauthorized, "User verification is failed", err)
				return
			}
			ctx := context.WithValue(r.Context(), userIDContextKey, userID)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
