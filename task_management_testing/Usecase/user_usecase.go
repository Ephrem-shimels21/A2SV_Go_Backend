package usecase

import (
	"context"

	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain/dtos"
	infrastructure "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Infrastructure"
)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (userUse *UserUsecase) RegisterUser(c context.Context, registerDto dtos.RegisterDto) (*domain.User, error) {
	hashedPassword, err := infrastructure.HashPassword(registerDto.Password)

	if err != nil {
		return &domain.User{}, err
	}
	println("in usecase")

	return userUse.userRepository.RegisterUser(c, registerDto, hashedPassword)
}

func (userUse *UserUsecase) Login(c context.Context, registerDto dtos.RegisterDto) (string, error) {
	existingUser, err := userUse.userRepository.FindUser(c, registerDto)

	if err != nil {
		return "", err
	}

	token, err := infrastructure.GenerateJWT(existingUser)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (userUse *UserUsecase) PromoteUser(c context.Context, promoteDto dtos.PromoteDto) error {
	return userUse.userRepository.PromoteUser(c, promoteDto)
}
