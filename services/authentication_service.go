package services

import (
	"errors"
	"microservices/pkg/hashpass"
	"microservices/repositories"

	jwtPkg "microservices/pkg/jwt"
)

type AuthenticationService interface {
	Register(user UserRequest) error
	Login(userLogin UserLogin) (string, error)
}

type authenticationService struct {
	authenRepository repositories.UserRepository
}

func NewAuthenicationServices(authenRepository repositories.UserRepository) AuthenticationService {
	return &authenticationService{authenRepository: authenRepository}
}

// สมัครสมาชิก
func (s *authenticationService) Register(userReq UserRequest) error {

	//เข้ารหัส
	hashPass, err := hashpass.HashPassword(userReq.Password)
	if err != nil {
		return err
	}
	user := repositories.User{
		Username:  userReq.Username,
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Phone:     userReq.Phone,
		Email:     userReq.Email,
		Password:  hashPass,
	}

	if err := s.authenRepository.CreateUser(user); err != nil {
		return err
	}
	return nil
}

// เข้าสู่ระบบ
func (s *authenticationService) Login(userLogin UserLogin) (string, error) {
	//ค้นหายูส
	userDB, err := s.authenRepository.FindUserByUsername(userLogin.Username)
	if err != nil {
		return "", nil
	}
	//เช็ค password
	checked := hashpass.CheckPasswordHash(userLogin.Password, userDB.Password)
	if !checked {
		return "", errors.New("invalid password")
	}

	//สร้าง token timeout 30 minute
	token := jwtPkg.GenerateToken(userDB.ID, 30)

	return token, nil
}
