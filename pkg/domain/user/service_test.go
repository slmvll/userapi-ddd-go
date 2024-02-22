package user_test

import (
	"errors"
	"testing"
	"userapi-ddd-go/pkg/domain/user"
	"userapi-ddd-go/pkg/domain/user/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUser_AddUser(t *testing.T) {
	// Build our needed testcase data struct
	type testCase struct {
		name        string
		email       string
		expectedErr error
	}
	// Create new test cases
	testCases := []testCase{
		{
			name:        "Empty Email validation",
			email:       "",
			expectedErr: errors.New("empty email"),
		}, {
			name:        "Valid Email",
			email:       "test@test.com",
			expectedErr: nil,
		},
	}

	repo := &mocks.UserRepository{}
	repo.On("AddUser", mock.AnythingOfType("user.User")).
		Return(nil).
		Once()

	service := user.NewUserService(repo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Create a new User
			user := user.User{
				Email: tc.email,
			}
			_, err := service.AddUser(user)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestTask_GetAllUsers(t *testing.T) {
	addedUser := user.User{
		Email: "test@test.com",
	}

	repo := &mocks.UserRepository{}
	repo.On("GetAllUsers").
		Return([]user.User{addedUser}, nil).
		Once()

	service := user.NewUserService(repo)

	_, err := service.GetAllUsers()

	assert.Equal(t, nil, err)
}
