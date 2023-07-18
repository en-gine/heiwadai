package user

import (
	"server/core/entity"
	"server/core/errors"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	"server/core/infra/types"
	"time"
)

type AuthUsecase struct {
	userRepository repository.IUserRepository
	userQuery      queryservice.IUserQueryService
	authAction     action.IAuthAction
}

func NewAuthUsecase(userRepository repository.IUserRepository, userQuery queryservice.IUserQueryService, authAction action.IAuthAction) *AuthUsecase {
	return &AuthUsecase{
		userRepository: userRepository,
		userQuery:      userQuery,
		authAction:     authAction,
	}
}

func (u *AuthUsecase) Register(
	FirstName string,
	LastName string,
	FirstNameKana string,
	LastNameKana string,
	CompanyName *string,
	BirthDate time.Time,
	ZipCode *string,
	Prefecture string,
	City *string,
	Address *string,
	Tel *string,
	Mail string,
	AcceptMail bool, // メルマガ配信可
	AcceptTerm bool, // 利用規約に同意
) (*entity.User, *errors.DomainError) {

	if AcceptTerm == false {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "利用規約に同意してください")
	}

	existUser, err := u.userQuery.GetByMail(Mail)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "ユーザーの検索に失敗しました")
	}

	if existUser != nil {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "既に登録されているメールアドレスです")
	}

	prefecture, domainErr := entity.StringToPrefecture(Prefecture)
	if domainErr != nil {
		return nil, domainErr
	}

	// 招待メール送信
	newId, err := u.authAction.InviteUserByEmail(Mail)

	userData := entity.CreateUser(
		newId,
		FirstName,
		LastName,
		FirstNameKana,
		LastNameKana,
		CompanyName,
		BirthDate,
		ZipCode,
		prefecture,
		City,
		Address,
		Tel,
		Mail,
		AcceptMail,
	)

	err = u.userRepository.Save(userData)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return userData, nil
}

func (u *AuthUsecase) SignUp(
	Mail string,
	Password string,
) *errors.DomainError {
	err := u.authAction.SignUp(Mail, Password)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AuthUsecase) SignIn(
	Mail string,
	Password string,
) (*types.Token, *errors.DomainError) {
	token, err := u.authAction.SignIn(Mail, Password)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return token, nil
}

func (u *AuthUsecase) ResetPasswordMail(
	Mail string,
) *errors.DomainError {
	err := u.authAction.ResetPasswordMail(Mail)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AuthUsecase) UpdatePassword(
	Password string,
	Token string,
) *errors.DomainError {
	err := u.authAction.UpdatePassword(Password, Token)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AuthUsecase) UpdateEmail(
	Mail string,
	Token string,
) *errors.DomainError {
	err := u.authAction.UpdateEmail(Mail, Token)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}