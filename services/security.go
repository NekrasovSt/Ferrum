package services

import (
	"Ferrum/data"
	"Ferrum/dto"
	"github.com/google/uuid"
)

type SecurityService interface {
	Validate(tokenIssueData *dto.TokenGenerationData, realm *data.Realm) *data.OperationError
	CheckCredentials(tokenIssueData *dto.TokenGenerationData, realm *data.Realm) *data.OperationError
	GetCurrentUser(realm *data.Realm, userName string) *data.User
	StartOrUpdateSession(realm string, userId uuid.UUID, duration int) uuid.UUID
	AssignTokens(realm string, userId uuid.UUID, accessToken *string, refreshToken *string)
	GetSession(realm string, userId uuid.UUID) *data.UserSession
	GetSessionByAccessToken(realm string, token *string) *data.UserSession
	IsSessionExpired(realm string, userId uuid.UUID) bool
}
