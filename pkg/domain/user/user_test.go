package user_test

import (
	"errors"
	"testing"
	"userapi-ddd-go/pkg/domain/user"
	"userapi-ddd-go/pkg/domain/user/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUser_NewUser(t *testing.T) {
	// Build our needed testcase data struct
	type testCase struct {
		name        string
		email       string
		password    string
		expectedErr error
	}
	// Create new test cases
	testCases := []testCase{
		{
			name:        "Empty Email validation",
			email:       "",
			password:    "",
			expectedErr: errors.New("empty email"),
		}, {
			name:        "Valid Email",
			email:       "test@test.com",
			password:    "",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new User
			_, err := user.NewUser(tc.email, tc.password)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestUser_GetUser(t *testing.T) {
	addedUser := user.User{
		Email: "test@test.com",
	}

	repo := &mocks.UserRepository{}
	repo.On("GetUser", mock.AnythingOfType("uuid.UUID")).
		Return(addedUser, nil).
		Once()

	service := user.NewUserService(repo)

	_, err := service.GetUser(uuid.New())

	assert.Equal(t, nil, err)

}
