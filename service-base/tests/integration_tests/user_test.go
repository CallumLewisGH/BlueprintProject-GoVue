package integration_tests

import (
	"context"
	"testing"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/contracts/requests"
	command "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/handlers/commands"
	query "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/handlers/queries"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestCreateUser_Success() {
	ctx := context.Background()

	testUser := requests.CreateUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		AuthId:   "test-auth-id",
	}

	createdUser, err := command.CreateUserCommand(ctx, testUser)
	s.NoError(err)
	s.NotNil(createdUser)
	s.Equal(testUser.Username, createdUser.Username)
	s.Equal(testUser.Email, createdUser.Email)
}

func (s *TestSuite) TestGetUserById_Success() {
	ctx := context.Background()

	// Create test user
	testUser := requests.CreateUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		AuthId:   "test-auth-id",
	}
	createdUser, err := command.CreateUserCommand(ctx, testUser)
	s.NoError(err)

	// Get user by ID
	foundUser, err := query.GetUserByIdQuery(ctx, createdUser.ID)
	s.NoError(err)
	s.Equal(createdUser.ID, foundUser.ID)
	s.Equal(testUser.Username, foundUser.Username)
	s.Equal(testUser.Email, foundUser.Email)
}

func (s *TestSuite) TestGetUserById_NotFound() {
	ctx := context.Background()
	nonExistentId := uuid.New()
	foundUser, err := query.GetUserByIdQuery(ctx, nonExistentId)
	s.Error(err)
	s.Contains(err.Error(), "record not found")
	s.Nil(foundUser)
}

func (s *TestSuite) TestUpdateUser_Success() {
	ctx := context.Background()
	// Create test user
	testUser := requests.CreateUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		AuthId:   "test-auth-id",
	}
	createdUser, err := command.CreateUserCommand(ctx, testUser)
	s.NoError(err)

	// Update user
	newUsername := "updated-username"
	newEmail := "updated@example.com"
	updateReq := requests.UpdateUserRequest{
		Username: &newUsername,
		Email:    &newEmail,
	}
	updatedUser, err := command.UpdateUserByIdCommand(ctx, createdUser.ID, updateReq)
	s.NoError(err)
	s.Equal(newUsername, updatedUser.Username)
	s.Equal(newEmail, updatedUser.Email)
}

func (s *TestSuite) TestUpdateUser_NotFound() {
	ctx := context.Background()
	nonExistentId := uuid.New()
	newUsername := "updated-username"
	newEmail := "updated@example.com"
	updateReq := requests.UpdateUserRequest{
		Username: &newUsername,
		Email:    &newEmail,
	}
	updatedUser, err := command.UpdateUserByIdCommand(ctx, nonExistentId, updateReq)
	s.Error(err)
	s.Contains(err.Error(), "record not found")
	s.Nil(updatedUser)
}

func (s *TestSuite) TestDeleteUser_Success() {
	ctx := context.Background()
	testUser := requests.CreateUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		AuthId:   "test-auth-id",
	}
	createdUser, err := command.CreateUserCommand(ctx, testUser)
	s.NoError(err)

	// Delete user
	deletedUser, err := command.DeleteUserByIdCommand(ctx, createdUser.ID)
	s.NoError(err)
	s.Equal(createdUser.ID, deletedUser.ID)

	// Verify user is deleted
	_, err = query.GetUserByIdQuery(ctx, createdUser.ID)
	s.Error(err)
	s.Contains(err.Error(), "record not found")
}

func (s *TestSuite) TestDeleteUser_NotFound() {
	ctx := context.Background()
	nonExistentId := uuid.New()
	deletedUser, err := command.DeleteUserByIdCommand(ctx, nonExistentId)
	s.Error(err)
	s.Contains(err.Error(), "record not found")
	s.Nil(deletedUser)
}

func (s *TestSuite) TestGetUserByAuthId_Success() {
	ctx := context.Background()
	testUser := requests.CreateUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		AuthId:   "test-auth-id",
	}

	_, err := command.CreateUserCommand(ctx, testUser)
	s.NoError(err)

	outTestUser, err := query.GetUserByAuthIdQuery(ctx, testUser.AuthId)

	s.NoError(err)
	s.Equal(outTestUser.Email, testUser.Email)
}

func (s *TestSuite) TestGetUserByAuthId_Fail() {
	ctx := context.Background()
	nonExistentAuthId := "non-existent-auth-id"
	outTestUser, err := query.GetUserByAuthIdQuery(ctx, nonExistentAuthId)

	s.Contains(err.Error(), "record not found")
	s.Nil(outTestUser)
}
