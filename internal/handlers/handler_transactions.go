package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AliKefall/DonemOdevi/internal/db"
	"github.com/google/uuid"
)

type createTransactionRequest struct {
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Category  string  `json:"category"`
	Note      string  `json:"note"`
	OccuredAt string  `json:"occured_at"`
}

func CreateTransaction(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := GetUserIDFromContext(r.Context())
		if err != nil {
			RespondWithError(w, 401, "Unauthorized", err)
			return
		}

		var req createTransactionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			RespondWithError(w, 400, "Invalid JSON", err)
			return
		}

		occuredAt, err := time.Parse(time.RFC3339, req.OccuredAt)
		if err != nil {
			RespondWithError(w, 400, "Invalid occured_at format. Must be RFC3339", err)
			return
		}

		now := time.Now()

		params := db.CreateTransactionParams{
			ID:         uuid.New(),
			UserID:     userID,
			Amount:     req.Amount,
			Currency:   req.Currency,
			Category:   req.Category,
			Note:       req.Note,
			OccurredAt: occuredAt.Format(time.RFC3339),
			CreatedAt:  now.Format(time.RFC3339),
			UpdatedAt:  now.Format(time.RFC3339),
		}

		tx, err := q.CreateTransaction(r.Context(), params)
		if err != nil {
			RespondWithError(w, 500, "Failed to create transaction", err)
			return
		}
		RespondWithJson(w, 201, tx)

	}
}
