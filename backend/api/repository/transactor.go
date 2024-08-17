package repository

import (
	"context"
	"database/sql"
	"fmt"
	"post-api/appcontext"
	"post-api/apperror"
)

var transactorFile = "transactor.go"

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Transactor interface {
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type transactor struct {
	db *sql.DB
}

func NewTransactor(db *sql.DB) *transactor {
	return &transactor{db}
}

func (r *transactor) Begin(ctx context.Context) (context.Context, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, apperror.NewError(
			err,
			transactorFile,
			fmt.Sprintf("repository.Begin(%v)", ctx),
			"",
		)
	}

	return context.WithValue(ctx, appcontext.KeyTx, tx), nil
}

func (r *transactor) Commit(ctx context.Context) error {
	tx, err := GetTxFromContext(ctx)
	if err != nil {
		return err
	}
	if tx != nil {
		err = tx.Commit()
		if err != nil {
			return apperror.NewError(
				err,
				transactorFile,
				fmt.Sprintf("repository.Commit(%v)", ctx),
				"",
			)
		}
	}

	return nil
}

func (r *transactor) Rollback(ctx context.Context) error {
	tx, err := GetTxFromContext(ctx)
	if err != nil {
		return err
	}
	if tx != nil {
		err = tx.Rollback()
		if err != nil {
			return apperror.NewError(
				err,
				transactorFile,
				fmt.Sprintf("repository.Rollback(%v)", ctx),
				"",
			)
		}
	}

	return nil
}

func GetTxFromContext(ctx context.Context) (*sql.Tx, error) {
	trx := ctx.Value(appcontext.KeyTx)
	if trx != nil {
		tx, ok := trx.(*sql.Tx)
		if !ok {
			return nil, apperror.NewError(
				apperror.ErrTxConversion,
				transactorFile,
				fmt.Sprintf("GetTxFromContext(%v)", ctx),
				"",
			)
		}

		return tx, nil
	}

	return nil, nil
}
