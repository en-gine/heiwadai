package user

import (
	"context"

	shared "server/api/v1/shared"
	"server/api/v1/user"
	userv1 "server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	"server/controller/util"
	"server/core/entity"
	usecase "server/core/usecase/user"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
)

type UserDataController struct {
	usecase usecase.UserDataUsecase
}

var _ userv1connect.UserDataControllerClient = &UserDataController{}

func NewUserDataController(userusecase *usecase.UserDataUsecase) *UserDataController {
	return &UserDataController{
		usecase: *userusecase,
	}
}

func (u *UserDataController) Update(ctx context.Context, req *connect.Request[userv1.UserUpdateDataRequest]) (*connect.Response[userv1.UserDataResponse], error) {
	msg := req.Msg

	user, domainErr := u.usecase.Update(
		uuid.MustParse(msg.ID),
		msg.FirstName,
		msg.LastName,
		msg.FirstNameKana,
		msg.LastNameKana,
		msg.CompanyName,
		util.TimeStampPtrToTimePtr(msg.BirthDate),
		msg.ZipCode,
		int(msg.Prefecture),
		msg.City,
		msg.Address,
		msg.Tel,
		msg.Mail,
		msg.AcceptMail,
	)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	res := UserEntityToResponse(user)
	return connect.NewResponse(res), nil
}

func UserEntityToResponse(user *entity.User) *user.UserDataResponse {
	return &userv1.UserDataResponse{
		ID:            user.ID.String(),
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		FirstNameKana: user.FirstNameKana,
		LastNameKana:  user.LastNameKana,
		CompanyName:   user.CompanyName,
		BirthDate:     util.TimePtrToTimeStampPtr(user.BirthDate),
		ZipCode:       user.ZipCode,
		Prefecture:    shared.Prefecture(user.Prefecture.ToInt()),
		City:          user.City,
		Address:       user.Address,
		Tel:           user.Tel,
		Mail:          user.Mail,
		AcceptMail:    user.AcceptMail,
	}
}
