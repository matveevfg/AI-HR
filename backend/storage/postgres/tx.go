package postgres

import (
	"context"

	"github.com/uptrace/bun"
)

type txValue string

const txKey txValue = "tx"

// CtxWithTx returns a context with a transaction.
// Creates a new transaction if there is none and set in the context.
func (p *Postgres) CtxWithTx(ctx context.Context) (context.Context, error) {
	var (
		tx *bun.Tx
	)

	tx, ok := ctx.Value(txKey).(*bun.Tx)
	if !ok {
		btx, err := p.d.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}

		tx = &btx
	}

	return context.WithValue(ctx, txKey, tx), nil
}

// TxCommit commits the active transaction active from the context.
// Returns nil if there is no transaction in the context.
func (p *Postgres) TxCommit(ctx context.Context) error {
	tx, ok := ctx.Value(txKey).(*bun.Tx)
	if !ok {
		return nil
	}

	return tx.Commit()
}

// TxRollback rollbacks the active transaction active from the context.
// Returns nil if there is no transaction in the context.
func (p *Postgres) TxRollback(ctx context.Context) error {
	tx, ok := ctx.Value(txKey).(*bun.Tx)
	if !ok {
		return nil
	}

	return tx.Rollback()
}

func txFromCtx(ctx context.Context) (*bun.Tx, bool) {
	tx, ok := ctx.Value(txKey).(*bun.Tx)
	if !ok {
		return nil, false
	}

	return tx, true
}
