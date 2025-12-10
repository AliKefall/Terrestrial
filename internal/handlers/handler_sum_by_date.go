package handlers

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/AliKefall/DonemOdevi/internal/db"
	"github.com/google/uuid"
)

type SumResult struct {
	Key   string  `json:"key"`
	Total float64 `json:"total"`
}

func SumQuery(ctx context.Context, dbase *sql.DB, query string, userID uuid.UUID, start, end time.Time) ([]SumResult, error) {
	rows, err := dbase.QueryContext(ctx, query, userID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []SumResult{}

	for rows.Next() {
		var key string
		var total sql.NullFloat64

		if err := rows.Scan(&key, &total); err != nil {
			return nil, err
		}
		results = append(results, SumResult{
			Key:   key,
			Total: total.Float64,
		})

	}
	return results, nil

}

func ParseDateRange(r *http.Request) (time.Time, time.Time, error) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	if startStr == "" || endStr == "" {
		return time.Time{}, time.Time{}, errors.New("start and end parameters are required")

	}

	start, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	end, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return start, end, nil
}

const querySumByDay = `
SELECT DATE(occurred_at) AS day, SUM(amount)
FROM transactions
WHERE user_id = ?
  AND occurred_at BETWEEN ? AND ?
GROUP BY day
ORDER BY day;
`

const querySumByMonth = `
SELECT strftime('%Y-%m', occurred_at) AS month, SUM(amount)
FROM transactions
WHERE user_id = ?
  AND occurred_at BETWEEN ? AND ?
GROUP BY month
ORDER BY month;
`

const querySumByYear = `
SELECT strftime('%Y', occurred_at) AS year, SUM(amount)
FROM transactions
WHERE user_id = ?
  AND occurred_at BETWEEN ? AND ?
GROUP BY year
ORDER BY year;
`

func SumByDayHandler(q *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := GetUserIDFromContext(r.Context())
		if err != nil {
			RespondWithError(w, 401, "Unauthorized", err)
			return
		}

		start, end, err := ParseDateRange(r)
		if err != nil {
			RespondWithError(w, 400, "Bad date format", err)
			return
		}

		results, err := SumQuery(r.Context(), q, querySumByDay, userID, start, end)
		if err != nil {
			RespondWithError(w, 500, "Query failed", err)
			return
		}

		RespondWithJson(w, 200, results)
	}
}

func SumByMonthHandler(q *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := GetUserIDFromContext(r.Context())
		if err != nil {
			RespondWithError(w, 401, "Unauthorized", err)
			return
		}

		start, end, err := ParseDateRange(r)
		if err != nil {
			RespondWithError(w, 400, "Invalid date range", err)
			return
		}

		results, err := SumQuery(r.Context(), q, querySumByMonth, userID, start, end)
		if err != nil {
			RespondWithError(w, 500, "Query failed", err)
			return
		}

		RespondWithJson(w, 200, results)
	}
}

func SumByYearHandler(q *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := GetUserIDFromContext(r.Context())
		if err != nil {
			RespondWithError(w, 401, "Unauthorized", err)
			return
		}

		start, end, err := ParseDateRange(r)
		if err != nil {
			RespondWithError(w, 400, "Invalid dates", err)
			return
		}

		results, err := SumQuery(r.Context(), q, querySumByYear, userID, start, end)
		if err != nil {
			RespondWithError(w, 500, "Failed query", err)
			return
		}

		RespondWithJson(w, 200, results)
	}
}
