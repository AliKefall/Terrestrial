package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AliKefall/DonemOdevi/internal/auth"
	"github.com/AliKefall/DonemOdevi/internal/db"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	UserID       string    `json:"user_id"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func Login(q *db.Queries, tokenSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req loginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "Login attempt is failed, failed to fetch login data.", err)
			return
		}
		user, err := q.GetUserByEmail(r.Context(), req.Email)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "User could not be found.", err)
			return
		}

		if !auth.ComparePassword(user.Password, req.Password) {
			RespondWithError(w, http.StatusUnauthorized, "Email or password is missing or wrong.", nil)
			return

		}

		accessToken, err := auth.MakeJWT(user.ID, tokenSecret, time.Hour*1)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "JWT could not be created", err)
			return
		}

		refreshToken, err := auth.MakeRefreshToken()
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Refresh token could not be created", err)
			return
		}
		now := time.Now().UTC()

		_, err = q.CreateRefreshToken(r.Context(), db.CreateRefreshTokenParams{
			Token:     refreshToken,
			UserID:    user.ID,
			CreatedAt: now,
			UpdatedAt: now,
			ExpiresAt: now.Add(30 * 24 * time.Hour),
		})
		if err != nil {
			RespondWithError(w, 500, "Refresh Token could not be saved.", err)
			return
		}
		RespondWithJson(w, 200, loginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			UserID:       user.ID.String(),
			ExpiresAt:    now.Add(1 * time.Hour),
		})
	}
}
