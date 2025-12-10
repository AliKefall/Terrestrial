package handlers

import (
	"database/sql"
	"encoding/json"

	"net/http"
	"strconv"
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
			ID:       uuid.New(),
			UserID:   userID,
			Amount:   req.Amount,
			Currency: req.Currency,
			Category: sql.NullString{
				String: req.Category,
				Valid:  req.Category != "",
			},
			Note: sql.NullString{
				String: req.Note,
				Valid:  req.Note != "",
			},
			OccurredAt: occuredAt,
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		tx, err := q.CreateTransaction(r.Context(), params)
		if err != nil {
			RespondWithError(w, 500, "Failed to create transaction", err)
			return
		}
		RespondWithJson(w, 201, tx)

	}
}

func ListTransactions(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := GetUserIDFromContext(r.Context())
		if err != nil {
			RespondWithError(w, 401, "Unauthorized", err)
			return
		}

		limitStr := r.URL.Query().Get("limit")
		offsetStr := r.URL.Query().Get("offset")

		limit := 20
		offset := 0

		if limitStr != "" {
			if v, err := strconv.Atoi(limitStr); err == nil {
				limit = v
			}
		}
		if offsetStr != "" {
			if v, err := strconv.Atoi(offsetStr); err == nil {
				offset = v
			}
		}

		params := db.ListTransactionsByUserParams{
			UserID: userID,
			Limit:  int64(limit),
			Offset: int64(offset),
		}

		txs, err := q.ListTransactionsByUser(r.Context(), params)
		if err != nil {
			RespondWithError(w, 500, "Failed to list transactions", err)
			return
		}

		RespondWithJson(w, 200, txs)
	}
}
