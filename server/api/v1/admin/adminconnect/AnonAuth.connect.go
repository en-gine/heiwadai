// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/admin/AnonAuth.proto

package adminconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	admin "server/api/v1/admin"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// AnonAuthControllerName is the fully-qualified name of the AnonAuthController service.
	AnonAuthControllerName = "server.admin.AnonAuthController"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AnonAuthControllerSignUpProcedure is the fully-qualified name of the AnonAuthController's SignUp
	// RPC.
	AnonAuthControllerSignUpProcedure = "/server.admin.AnonAuthController/SignUp"
	// AnonAuthControllerSignInProcedure is the fully-qualified name of the AnonAuthController's SignIn
	// RPC.
	AnonAuthControllerSignInProcedure = "/server.admin.AnonAuthController/SignIn"
	// AnonAuthControllerResetPasswordMailProcedure is the fully-qualified name of the
	// AnonAuthController's ResetPasswordMail RPC.
	AnonAuthControllerResetPasswordMailProcedure = "/server.admin.AnonAuthController/ResetPasswordMail"
	// AnonAuthControllerSetNewPasswordProcedure is the fully-qualified name of the AnonAuthController's
	// SetNewPassword RPC.
	AnonAuthControllerSetNewPasswordProcedure = "/server.admin.AnonAuthController/SetNewPassword"
)

// AnonAuthControllerClient is a client for the server.admin.AnonAuthController service.
type AnonAuthControllerClient interface {
	// トークン不要
	SignUp(context.Context, *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[emptypb.Empty], error)
	SignIn(context.Context, *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[admin.AnonAuthTokenResponse], error)
	ResetPasswordMail(context.Context, *connect_go.Request[admin.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error)
	SetNewPassword(context.Context, *connect_go.Request[admin.SetNewPasswordRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewAnonAuthControllerClient constructs a client for the server.admin.AnonAuthController service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAnonAuthControllerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AnonAuthControllerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &anonAuthControllerClient{
		signUp: connect_go.NewClient[admin.AdminAuthRequest, emptypb.Empty](
			httpClient,
			baseURL+AnonAuthControllerSignUpProcedure,
			opts...,
		),
		signIn: connect_go.NewClient[admin.AdminAuthRequest, admin.AnonAuthTokenResponse](
			httpClient,
			baseURL+AnonAuthControllerSignInProcedure,
			opts...,
		),
		resetPasswordMail: connect_go.NewClient[admin.ResetPasswordRequest, emptypb.Empty](
			httpClient,
			baseURL+AnonAuthControllerResetPasswordMailProcedure,
			opts...,
		),
		setNewPassword: connect_go.NewClient[admin.SetNewPasswordRequest, emptypb.Empty](
			httpClient,
			baseURL+AnonAuthControllerSetNewPasswordProcedure,
			opts...,
		),
	}
}

// anonAuthControllerClient implements AnonAuthControllerClient.
type anonAuthControllerClient struct {
	signUp            *connect_go.Client[admin.AdminAuthRequest, emptypb.Empty]
	signIn            *connect_go.Client[admin.AdminAuthRequest, admin.AnonAuthTokenResponse]
	resetPasswordMail *connect_go.Client[admin.ResetPasswordRequest, emptypb.Empty]
	setNewPassword    *connect_go.Client[admin.SetNewPasswordRequest, emptypb.Empty]
}

// SignUp calls server.admin.AnonAuthController.SignUp.
func (c *anonAuthControllerClient) SignUp(ctx context.Context, req *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.signUp.CallUnary(ctx, req)
}

// SignIn calls server.admin.AnonAuthController.SignIn.
func (c *anonAuthControllerClient) SignIn(ctx context.Context, req *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[admin.AnonAuthTokenResponse], error) {
	return c.signIn.CallUnary(ctx, req)
}

// ResetPasswordMail calls server.admin.AnonAuthController.ResetPasswordMail.
func (c *anonAuthControllerClient) ResetPasswordMail(ctx context.Context, req *connect_go.Request[admin.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.resetPasswordMail.CallUnary(ctx, req)
}

// SetNewPassword calls server.admin.AnonAuthController.SetNewPassword.
func (c *anonAuthControllerClient) SetNewPassword(ctx context.Context, req *connect_go.Request[admin.SetNewPasswordRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.setNewPassword.CallUnary(ctx, req)
}

// AnonAuthControllerHandler is an implementation of the server.admin.AnonAuthController service.
type AnonAuthControllerHandler interface {
	// トークン不要
	SignUp(context.Context, *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[emptypb.Empty], error)
	SignIn(context.Context, *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[admin.AnonAuthTokenResponse], error)
	ResetPasswordMail(context.Context, *connect_go.Request[admin.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error)
	SetNewPassword(context.Context, *connect_go.Request[admin.SetNewPasswordRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewAnonAuthControllerHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAnonAuthControllerHandler(svc AnonAuthControllerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	anonAuthControllerSignUpHandler := connect_go.NewUnaryHandler(
		AnonAuthControllerSignUpProcedure,
		svc.SignUp,
		opts...,
	)
	anonAuthControllerSignInHandler := connect_go.NewUnaryHandler(
		AnonAuthControllerSignInProcedure,
		svc.SignIn,
		opts...,
	)
	anonAuthControllerResetPasswordMailHandler := connect_go.NewUnaryHandler(
		AnonAuthControllerResetPasswordMailProcedure,
		svc.ResetPasswordMail,
		opts...,
	)
	anonAuthControllerSetNewPasswordHandler := connect_go.NewUnaryHandler(
		AnonAuthControllerSetNewPasswordProcedure,
		svc.SetNewPassword,
		opts...,
	)
	return "/server.admin.AnonAuthController/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AnonAuthControllerSignUpProcedure:
			anonAuthControllerSignUpHandler.ServeHTTP(w, r)
		case AnonAuthControllerSignInProcedure:
			anonAuthControllerSignInHandler.ServeHTTP(w, r)
		case AnonAuthControllerResetPasswordMailProcedure:
			anonAuthControllerResetPasswordMailHandler.ServeHTTP(w, r)
		case AnonAuthControllerSetNewPasswordProcedure:
			anonAuthControllerSetNewPasswordHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAnonAuthControllerHandler returns CodeUnimplemented from all methods.
type UnimplementedAnonAuthControllerHandler struct{}

func (UnimplementedAnonAuthControllerHandler) SignUp(context.Context, *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AnonAuthController.SignUp is not implemented"))
}

func (UnimplementedAnonAuthControllerHandler) SignIn(context.Context, *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[admin.AnonAuthTokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AnonAuthController.SignIn is not implemented"))
}

func (UnimplementedAnonAuthControllerHandler) ResetPasswordMail(context.Context, *connect_go.Request[admin.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AnonAuthController.ResetPasswordMail is not implemented"))
}

func (UnimplementedAnonAuthControllerHandler) SetNewPassword(context.Context, *connect_go.Request[admin.SetNewPasswordRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AnonAuthController.SetNewPassword is not implemented"))
}
