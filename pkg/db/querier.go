// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccessToken(ctx context.Context, arg CreateAccessTokenParams) (AccessToken, error)
	CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error)
	CreateOrganisation(ctx context.Context, arg CreateOrganisationParams) (Organisation, error)
	CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccessToken(ctx context.Context, organisationID int32) error
	DeleteGroup(ctx context.Context, id int32) error
	DeleteOrganisation(ctx context.Context, id int32) error
	DeleteRole(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	GetAccessTokenByOrganisationID(ctx context.Context, organisationID int32) (AccessToken, error)
	GetGroupByID(ctx context.Context, id int32) (Group, error)
	GetGroupByName(ctx context.Context, name string) (Group, error)
	GetOrganisation(ctx context.Context, id int32) (Organisation, error)
	GetOrganisationByName(ctx context.Context, name string) (Organisation, error)
	GetOrganisationForUpdate(ctx context.Context, name string) (Organisation, error)
	GetRoleByID(ctx context.Context, id int32) (Role, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int32) (User, error)
	GetUserByUserName(ctx context.Context, username string) (User, error)
	ListGroups(ctx context.Context, arg ListGroupsParams) ([]Group, error)
	ListOrganisations(ctx context.Context, arg ListOrganisationsParams) ([]Organisation, error)
	ListRoles(ctx context.Context, arg ListRolesParams) ([]Role, error)
	ListUsers(ctx context.Context) ([]User, error)
	UpdateAccessToken(ctx context.Context, arg UpdateAccessTokenParams) (AccessToken, error)
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) (Group, error)
	UpdateOrganisation(ctx context.Context, arg UpdateOrganisationParams) (Organisation, error)
	UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)