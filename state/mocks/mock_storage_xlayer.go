// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	pgx "github.com/jackc/pgx/v4"
)

func (_m *StorageMock) GetBatchL2DataByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) ([]byte, error){
	return nil, nil
}

func (_m *StorageMock) GetBatchL2DataByNumbers(ctx context.Context, batchNumbers []uint64, dbTx pgx.Tx) (map[uint64][]byte, error){
	return nil, nil
}