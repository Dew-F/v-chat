package auth

import (
	"context"
	"errors"

	"github.com/Dew-F/v-chat/internal/id"
	"github.com/Dew-F/v-chat/internal/user"
)

type Service struct {
	users *user.Repository
	jwt   *JWTManager
}

func NewService(
	users *user.Repository,
	jwt *JWTManager,
) *Service {
	return &Service{
		users: users,
		jwt:   jwt,
	}
}

type RegisterInput struct {
	Username string
	Email    string
	Password string
}

func (s *Service) Register(
	ctx context.Context,
	input RegisterInput,
) (*user.User, error) {

	if err := ValidateRegisterInput(input); err != nil {
		return nil, err
	}

	passwordHash, err := HashPassword(input.Password)

	if err != nil {
		return nil, err
	}

	newUser := &user.User{
		ID:           id.New(),
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: passwordHash,
	}

	if err := s.users.Create(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	Token string
	User  *user.User
}

func (s *Service) Login(
	ctx context.Context,
	input LoginInput,
) (*LoginOutput, error) {
	u, err := s.users.FindByEmail(ctx, input.Email)

	if err != nil {
		return nil, err
	}

	if !VerifyPassword(
		input.Password,
		u.PasswordHash,
	) {
		return nil,
			errors.New("invalid credentials")
	}

	token, err := s.jwt.CreateToken(u.ID)

	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		Token: token,
		User:  u,
	}, nil
}
