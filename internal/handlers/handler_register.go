package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AliKefall/DonemOdevi/internal/auth"
	"github.com/AliKefall/DonemOdevi/internal/db"
	"github.com/google/uuid"
)

type registerRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerResponse struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

func Register(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req registerRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "Invalid request", err)
			return
		}
		if req.Email == "" || req.Username == "" || req.Password == "" {
			RespondWithError(w, http.StatusBadRequest, "Every field of the request must be filled", err)
			return
		}

		hashed := auth.HashPassword(req.Password)
		userID := uuid.New()

		now := time.Now().UTC()

		_, err = q.CreateUser(r.Context(), db.CreateUserParams{
			ID:        userID,
			Email:     req.Email,
			Username:  req.Username,
			Password:  hashed,
			CreatedAt: now.String(),
			UpdatedAt: now.String(),
		})
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "User cannot be created", err)
			return
		}
		RespondWithJson(w, 201, registerResponse{
			ID:       userID,
			Email:    req.Email,
			Username: req.Username,
		})
	}
}
