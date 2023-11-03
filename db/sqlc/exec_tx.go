package db

import (
	"context"
	"fmt"
)

// execTx executes a function within a database transaction Tx stands for transaction
// Starts with a lowercase e because we don't want this exported, we we make another function to help with that
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error { // when a function is an arguement like this, it is referred to as a callback
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
