package service

import (
	"os"
	"time"
	"log"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"

	"github.com/architecture-template/echo-ddd/config/key"
	"github.com/architecture-template/echo-ddd/auth/presentation/parameter"
	"github.com/architecture-template/echo-ddd/domain/model"
	"github.com/architecture-template/echo-ddd/domain/repository"
)

type UserService interface {
	FindByEmail(email string) (*model.User, error)
	RegisterUser(userParam *parameter.RegisterUser) (*model.User, error)
	LoginUser(userParam *parameter.LoginUser) (*model.User, error)
}

type userService struct {
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewUserService(
	userRepository repository.UserRepository,
	transactionRepository repository.TransactionRepository,
	) UserService {
	return &userService{
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

// FindByUserKey user_keyから取得する
func (u *userService) FindByEmail(email string) (*model.User, error){
	result, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// RegisterUser ユーザー登録
func (u *userService) RegisterUser(userParam *parameter.RegisterUser) (*model.User, error) {
	// transaction
	tx, err := u.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := u.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := u.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	userKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userParam.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	
	userModel := &model.User{}
	userModel.UserKey = userKey
	userModel.UserName = userParam.UserName
	userModel.Email = userParam.Email
	userModel.Password = string(hashedPassword)
	userModel.Token = "nil"

	result, err := u.userRepository.Insert(userModel, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// LoginUser ログイン
func (u *userService) LoginUser(userParam *parameter.LoginUser) (*model.User, error) {
	// transaction
	tx, err := u.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := u.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := u.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	user, err := u.userRepository.FindByEmail(userParam.Email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userParam.Password))
	if err != nil {
		return nil, err
	}

	baseToken := jwt.New(jwt.SigningMethodHS256)
	claims := baseToken.Claims.(jwt.MapClaims)
	claims["user_key"] = user.UserKey
	claims["user_name"] = user.UserName
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token, err := baseToken.SignedString([]byte(os.Getenv("AUTH_SECRET")))
	if err != nil {
		return nil, err
	}
	user.Token = token

	result, err := u.userRepository.Update(user, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
