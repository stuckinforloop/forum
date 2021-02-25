// Code generated by sqlc. DO NOT EDIT.
// source: fan.sql

package db

import (
	"context"
)

const createFan = `-- name: CreateFan :one
INSERT INTO fan (
  first_name,
  last_name,
  user_name,
  password,
  email,
  preferred_currency_id
  ) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, first_name, last_name, user_name, password, email, time_registered, time_confirmed, preferred_currency_id
`

type CreateFanParams struct {
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	UserName            string `json:"user_name"`
	Password            string `json:"password"`
	Email               string `json:"email"`
	PreferredCurrencyID int64  `json:"preferred_currency_id"`
}

func (q *Queries) CreateFan(ctx context.Context, arg CreateFanParams) (Fan, error) {
	row := q.db.QueryRowContext(ctx, createFan,
		arg.FirstName,
		arg.LastName,
		arg.UserName,
		arg.Password,
		arg.Email,
		arg.PreferredCurrencyID,
	)
	var i Fan
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.UserName,
		&i.Password,
		&i.Email,
		&i.TimeRegistered,
		&i.TimeConfirmed,
		&i.PreferredCurrencyID,
	)
	return i, err
}

const getFan = `-- name: GetFan :one
SELECT id, first_name, last_name, user_name, password, email, time_registered, time_confirmed, preferred_currency_id FROM fan
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetFan(ctx context.Context, id int64) (Fan, error) {
	row := q.db.QueryRowContext(ctx, getFan, id)
	var i Fan
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.UserName,
		&i.Password,
		&i.Email,
		&i.TimeRegistered,
		&i.TimeConfirmed,
		&i.PreferredCurrencyID,
	)
	return i, err
}

const updateEmail = `-- name: UpdateEmail :exec
UPDATE fan
SET email = $2
WHERE id = $1
`

type UpdateEmailParams struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func (q *Queries) UpdateEmail(ctx context.Context, arg UpdateEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateEmail, arg.ID, arg.Email)
	return err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE fan
SET password = $2
WHERE id = $1
`

type UpdatePasswordParams struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.ExecContext(ctx, updatePassword, arg.ID, arg.Password)
	return err
}

const updatePreferredCurrency = `-- name: UpdatePreferredCurrency :exec
UPDATE fan
SET preferred_currency_id = $2
WHERE id = $1
`

type UpdatePreferredCurrencyParams struct {
	ID                  int64 `json:"id"`
	PreferredCurrencyID int64 `json:"preferred_currency_id"`
}

func (q *Queries) UpdatePreferredCurrency(ctx context.Context, arg UpdatePreferredCurrencyParams) error {
	_, err := q.db.ExecContext(ctx, updatePreferredCurrency, arg.ID, arg.PreferredCurrencyID)
	return err
}
