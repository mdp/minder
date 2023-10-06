// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: profile_status.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const getProfileStatusByIdAndProject = `-- name: GetProfileStatusByIdAndProject :one
SELECT p.id, p.name, ps.profile_status, ps.last_updated FROM profile_status ps
INNER JOIN profiles p ON p.id = ps.profile_id
WHERE p.id = $1 AND p.project_id = $2
`

type GetProfileStatusByIdAndProjectParams struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
}

type GetProfileStatusByIdAndProjectRow struct {
	ID            uuid.UUID       `json:"id"`
	Name          string          `json:"name"`
	ProfileStatus EvalStatusTypes `json:"profile_status"`
	LastUpdated   time.Time       `json:"last_updated"`
}

func (q *Queries) GetProfileStatusByIdAndProject(ctx context.Context, arg GetProfileStatusByIdAndProjectParams) (GetProfileStatusByIdAndProjectRow, error) {
	row := q.db.QueryRowContext(ctx, getProfileStatusByIdAndProject, arg.ID, arg.ProjectID)
	var i GetProfileStatusByIdAndProjectRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ProfileStatus,
		&i.LastUpdated,
	)
	return i, err
}

const getProfileStatusByNameAndProject = `-- name: GetProfileStatusByNameAndProject :one
SELECT p.id, p.name, ps.profile_status, ps.last_updated FROM profile_status ps
INNER JOIN profiles p ON p.id = ps.profile_id
WHERE p.name = $1 AND p.project_id = $2
`

type GetProfileStatusByNameAndProjectParams struct {
	Name      string    `json:"name"`
	ProjectID uuid.UUID `json:"project_id"`
}

type GetProfileStatusByNameAndProjectRow struct {
	ID            uuid.UUID       `json:"id"`
	Name          string          `json:"name"`
	ProfileStatus EvalStatusTypes `json:"profile_status"`
	LastUpdated   time.Time       `json:"last_updated"`
}

func (q *Queries) GetProfileStatusByNameAndProject(ctx context.Context, arg GetProfileStatusByNameAndProjectParams) (GetProfileStatusByNameAndProjectRow, error) {
	row := q.db.QueryRowContext(ctx, getProfileStatusByNameAndProject, arg.Name, arg.ProjectID)
	var i GetProfileStatusByNameAndProjectRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ProfileStatus,
		&i.LastUpdated,
	)
	return i, err
}

const getProfileStatusByProject = `-- name: GetProfileStatusByProject :many
SELECT p.id, p.name, ps.profile_status, ps.last_updated FROM profile_status ps
INNER JOIN profiles p ON p.id = ps.profile_id
WHERE p.project_id = $1
`

type GetProfileStatusByProjectRow struct {
	ID            uuid.UUID       `json:"id"`
	Name          string          `json:"name"`
	ProfileStatus EvalStatusTypes `json:"profile_status"`
	LastUpdated   time.Time       `json:"last_updated"`
}

func (q *Queries) GetProfileStatusByProject(ctx context.Context, projectID uuid.UUID) ([]GetProfileStatusByProjectRow, error) {
	rows, err := q.db.QueryContext(ctx, getProfileStatusByProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProfileStatusByProjectRow{}
	for rows.Next() {
		var i GetProfileStatusByProjectRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ProfileStatus,
			&i.LastUpdated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRuleEvaluationsByProfileId = `-- name: ListRuleEvaluationsByProfileId :many
WITH
   eval_details AS (
   SELECT
       rule_eval_id,
       status AS eval_status,
       details AS eval_details,
       last_updated AS eval_last_updated
   FROM rule_details_eval
   ),
   remediation_details AS (
       SELECT
           rule_eval_id,
           status AS rem_status,
           details AS rem_details,
           last_updated AS rem_last_updated
       FROM rule_details_remediate
   ),
   alert_details AS (
       SELECT
           rule_eval_id,
           status AS alert_status,
           details AS alert_details,
           last_updated AS alert_last_updated
       FROM rule_details_alert
   )

SELECT
    ed.eval_status,
    ed.eval_last_updated,
    ed.eval_details,
    rd.rem_status,
    rd.rem_details,
    rd.rem_last_updated,
    ad.alert_status,
    ad.alert_details,
    ad.alert_last_updated,
    res.repository_id,
    res.entity,
    repo.repo_name,
    repo.repo_owner,
    repo.provider,
    rt.name AS rule_type_name,
    rt.id AS rule_type_id
FROM rule_evaluations res
         LEFT JOIN eval_details ed ON ed.rule_eval_id = res.id
         LEFT JOIN remediation_details rd ON rd.rule_eval_id = res.id
         LEFT JOIN alert_details ad ON ad.rule_eval_id = res.id
         INNER JOIN repositories repo ON repo.id = res.repository_id
         INNER JOIN rule_type rt ON rt.id = res.rule_type_id
WHERE res.profile_id = $1 AND
    (
        CASE
            WHEN $2::entities = 'repository' AND res.repository_id = $3::UUID THEN true
            WHEN $2::entities  = 'artifact' AND res.artifact_id = $3::UUID THEN true
            WHEN $3::UUID IS NULL THEN true
            ELSE false
            END
        ) AND (rt.name = $4 OR $4 IS NULL)
`

type ListRuleEvaluationsByProfileIdParams struct {
	ProfileID  uuid.UUID      `json:"profile_id"`
	EntityType NullEntities   `json:"entity_type"`
	EntityID   uuid.NullUUID  `json:"entity_id"`
	RuleName   sql.NullString `json:"rule_name"`
}

type ListRuleEvaluationsByProfileIdRow struct {
	EvalStatus       NullEvalStatusTypes        `json:"eval_status"`
	EvalLastUpdated  sql.NullTime               `json:"eval_last_updated"`
	EvalDetails      sql.NullString             `json:"eval_details"`
	RemStatus        NullRemediationStatusTypes `json:"rem_status"`
	RemDetails       sql.NullString             `json:"rem_details"`
	RemLastUpdated   sql.NullTime               `json:"rem_last_updated"`
	AlertStatus      NullAlertStatusTypes       `json:"alert_status"`
	AlertDetails     sql.NullString             `json:"alert_details"`
	AlertLastUpdated sql.NullTime               `json:"alert_last_updated"`
	RepositoryID     uuid.NullUUID              `json:"repository_id"`
	Entity           Entities                   `json:"entity"`
	RepoName         string                     `json:"repo_name"`
	RepoOwner        string                     `json:"repo_owner"`
	Provider         string                     `json:"provider"`
	RuleTypeName     string                     `json:"rule_type_name"`
	RuleTypeID       uuid.UUID                  `json:"rule_type_id"`
}

func (q *Queries) ListRuleEvaluationsByProfileId(ctx context.Context, arg ListRuleEvaluationsByProfileIdParams) ([]ListRuleEvaluationsByProfileIdRow, error) {
	rows, err := q.db.QueryContext(ctx, listRuleEvaluationsByProfileId,
		arg.ProfileID,
		arg.EntityType,
		arg.EntityID,
		arg.RuleName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListRuleEvaluationsByProfileIdRow{}
	for rows.Next() {
		var i ListRuleEvaluationsByProfileIdRow
		if err := rows.Scan(
			&i.EvalStatus,
			&i.EvalLastUpdated,
			&i.EvalDetails,
			&i.RemStatus,
			&i.RemDetails,
			&i.RemLastUpdated,
			&i.AlertStatus,
			&i.AlertDetails,
			&i.AlertLastUpdated,
			&i.RepositoryID,
			&i.Entity,
			&i.RepoName,
			&i.RepoOwner,
			&i.Provider,
			&i.RuleTypeName,
			&i.RuleTypeID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const upsertRuleDetailsAlert = `-- name: UpsertRuleDetailsAlert :one
INSERT INTO rule_details_alert (
    rule_eval_id,
    status,
    details,
    last_updated
)
VALUES ($1, $2, $3, NOW())
ON CONFLICT(rule_eval_id)
    DO UPDATE SET
                  status = $2,
                  details = $3,
                  last_updated = NOW()
    WHERE rule_details_alert.rule_eval_id = $1
RETURNING id
`

type UpsertRuleDetailsAlertParams struct {
	RuleEvalID uuid.UUID        `json:"rule_eval_id"`
	Status     AlertStatusTypes `json:"status"`
	Details    string           `json:"details"`
}

func (q *Queries) UpsertRuleDetailsAlert(ctx context.Context, arg UpsertRuleDetailsAlertParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, upsertRuleDetailsAlert, arg.RuleEvalID, arg.Status, arg.Details)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const upsertRuleDetailsEval = `-- name: UpsertRuleDetailsEval :one
INSERT INTO rule_details_eval (
    rule_eval_id,
    status,
    details,
    last_updated
)
VALUES ($1, $2, $3, NOW())
 ON CONFLICT(rule_eval_id)
    DO UPDATE SET
           status = $2,
           details = $3,
           last_updated = NOW()
    WHERE rule_details_eval.rule_eval_id = $1
RETURNING id
`

type UpsertRuleDetailsEvalParams struct {
	RuleEvalID uuid.UUID       `json:"rule_eval_id"`
	Status     EvalStatusTypes `json:"status"`
	Details    string          `json:"details"`
}

func (q *Queries) UpsertRuleDetailsEval(ctx context.Context, arg UpsertRuleDetailsEvalParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, upsertRuleDetailsEval, arg.RuleEvalID, arg.Status, arg.Details)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const upsertRuleDetailsRemediate = `-- name: UpsertRuleDetailsRemediate :one
INSERT INTO rule_details_remediate (
    rule_eval_id,
    status,
    details,
    last_updated
)
VALUES ($1, $2, $3, NOW())
ON CONFLICT(rule_eval_id)
    DO UPDATE SET
                  status = $2,
                  details = $3,
                  last_updated = NOW()
    WHERE rule_details_remediate.rule_eval_id = $1
RETURNING id
`

type UpsertRuleDetailsRemediateParams struct {
	RuleEvalID uuid.UUID              `json:"rule_eval_id"`
	Status     RemediationStatusTypes `json:"status"`
	Details    string                 `json:"details"`
}

func (q *Queries) UpsertRuleDetailsRemediate(ctx context.Context, arg UpsertRuleDetailsRemediateParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, upsertRuleDetailsRemediate, arg.RuleEvalID, arg.Status, arg.Details)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const upsertRuleEvaluations = `-- name: UpsertRuleEvaluations :one
INSERT INTO rule_evaluations (
    profile_id, repository_id, artifact_id, rule_type_id, entity
) VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (profile_id, repository_id, COALESCE(artifact_id, '00000000-0000-0000-0000-000000000000'::UUID), entity, rule_type_id)
  DO UPDATE SET profile_id = $1
RETURNING id
`

type UpsertRuleEvaluationsParams struct {
	ProfileID    uuid.UUID     `json:"profile_id"`
	RepositoryID uuid.NullUUID `json:"repository_id"`
	ArtifactID   uuid.NullUUID `json:"artifact_id"`
	RuleTypeID   uuid.UUID     `json:"rule_type_id"`
	Entity       Entities      `json:"entity"`
}

func (q *Queries) UpsertRuleEvaluations(ctx context.Context, arg UpsertRuleEvaluationsParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, upsertRuleEvaluations,
		arg.ProfileID,
		arg.RepositoryID,
		arg.ArtifactID,
		arg.RuleTypeID,
		arg.Entity,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
