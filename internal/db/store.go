//
// Copyright 2023 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

// ExtendQuerier extends the Querier interface with custom queries
type ExtendQuerier interface {
	Querier
	GetRuleEvaluationByProfileIdAndRuleType(ctx context.Context, profileID uuid.UUID, entityType NullEntities,
		entityID uuid.NullUUID, ruleName sql.NullString) (ListRuleEvaluationsByProfileIdRow, error)
}

// Store provides all functions to execute db queries and transactions
type Store interface {
	ExtendQuerier
	CheckHealth() error
	BeginTransaction() (*sql.Tx, error)
	GetQuerierWithTransaction(tx *sql.Tx) ExtendQuerier
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	db *sql.DB
	*Queries
}

// CheckHealth checks the health of the database
func (s *SQLStore) CheckHealth() error {
	return s.db.Ping()
}

// BeginTransaction begins a new transaction
func (s *SQLStore) BeginTransaction() (*sql.Tx, error) {
	return s.db.Begin()
}

// GetQuerierWithTransaction returns a new Querier with the provided transaction
func (*SQLStore) GetQuerierWithTransaction(tx *sql.Tx) ExtendQuerier {
	return New(tx)
}

// Commit commits a transaction
func (*SQLStore) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

// Rollback rolls back a transaction
func (*SQLStore) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// GetRuleEvaluationByProfileIdAndRuleType returns the rule evaluation for a given profile and its rule name
func (q *Queries) GetRuleEvaluationByProfileIdAndRuleType(
	ctx context.Context,
	profileID uuid.UUID,
	entityType NullEntities,
	entityID uuid.NullUUID,
	ruleName sql.NullString,
) (ListRuleEvaluationsByProfileIdRow, error) {
	params := ListRuleEvaluationsByProfileIdParams{
		ProfileID:  profileID,
		EntityType: entityType,
		EntityID:   entityID,
		RuleName:   ruleName,
	}
	res, err := q.ListRuleEvaluationsByProfileId(ctx, params)
	if err != nil {
		return ListRuleEvaluationsByProfileIdRow{}, err
	}

	// Single or no row expected
	switch len(res) {
	case 0:
		return ListRuleEvaluationsByProfileIdRow{}, nil
	case 1:
		return res[0], nil
	}
	return ListRuleEvaluationsByProfileIdRow{},
		fmt.Errorf("GetRuleEvaluationByProfileIdAndRuleType - expected 1 row, got %d", len(res))
}
