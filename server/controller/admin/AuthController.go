package admin

import (
	"context"
	"errors"

	"server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	"server/controller"
	usecase "server/core/usecase/admin"
	"server/infrastructure/logger"
	"server/router"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthController struct {
	authUseCase usecase.AuthUsecase
}

var _ adminv1connect.AuthControllerClient = &AuthController{}

func NewAuthController(authUsecase *usecase.AuthUsecase) *AuthController {
	return &AuthController{
		authUseCase: *authUsecase,
	}
}

func (ac *AuthController) SignOut(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	if ctx.Value(router.TokenKey) == nil {
		return connect.NewResponse(&emptypb.Empty{}), nil
	}
	token := ctx.Value(router.TokenKey).(string)
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ログインが必要です。"))
	}
	err := ac.authUseCase.SignOut(token)
	if err != nil {
		logger.Error(err.Error())
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("サインアウトに失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) ResetPasswordMail(ctx context.Context, req *connect.Request[admin.ResetPasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.ResetPasswordMail(msg.Email)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("リセットメールの送信に失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) UpdatePassword(ctx context.Context, req *connect.Request[admin.UpdatePasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	if ctx.Value(router.TokenKey) == nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("トークンが必要です。"))
	}

	token := ctx.Value(router.TokenKey).(string)
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ログインが必要です。"))
	}
	err := ac.authUseCase.UpdatePassword(msg.Password, token)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("パスワードの変更に失敗しました。\nネットワークの問題や同じパスワードに変更した、などの理由が考えられます。"))
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) UpdateEmail(ctx context.Context, req *connect.Request[admin.UpdateEmailRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	if ctx.Value(router.TokenKey) == nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("トークンが必要です。"))
	}
	token := ctx.Value(router.TokenKey).(string)
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ログインが必要です。"))
	}
	err := ac.authUseCase.UpdateEmail(msg.Mail, token)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("メールアドレスの変更に失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) Refresh(ctx context.Context, req *connect.Request[admin.AdminRefreshTokenRequest]) (*connect.Response[admin.AdminAuthTokenResponse], error) {
	msg := req.Msg
	token := msg.AccessToken
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ログインが必要です。"))
	}
	tkn, err := ac.authUseCase.Refresh(token, msg.RefreshToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("トークンの再取得に失敗しました。"))
	}
	return connect.NewResponse(&admin.AdminAuthTokenResponse{
		AccessToken:  tkn.AccessToken,
		ExpiresIn:    int64(*tkn.ExpiresIn),
		RefreshToken: *tkn.RefreshToken,
	}), nil
}

func (ac *AuthController) Register(ctx context.Context, req *connect.Request[admin.AdminRegisterRequest]) (*connect.Response[admin.AdminRegisterResponse], error) {
	msg := req.Msg
	storeID, err := uuid.Parse(req.Msg.BelongStoreID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("所属店舗のUUIDが正しい形式ではありません。"))
	}
	data, domaiErr := ac.authUseCase.Register(
		msg.Name,
		storeID,
		msg.Mail,
	)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	return connect.NewResponse(&admin.AdminRegisterResponse{
		ID: data.ID.String(),
	}), nil
}

func (ac *AuthController) ResendInviteMail(ctx context.Context, req *connect.Request[admin.ResendInviteRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg

	domaiErr := ac.authUseCase.ReInviteMail(
		msg.Mail,
	)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}
