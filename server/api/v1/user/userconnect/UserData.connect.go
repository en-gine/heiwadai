// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/user/UserData.proto

package userconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	user "server/api/v1/user"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// UserDataControllerName is the fully-qualified name of the UserDataController service.
	UserDataControllerName = "server.user.UserDataController"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserDataControllerRegisterProcedure is the fully-qualified name of the UserDataController's
	// Register RPC.
	UserDataControllerRegisterProcedure = "/server.user.UserDataController/Register"
	// UserDataControllerUpdateProcedure is the fully-qualified name of the UserDataController's Update
	// RPC.
	UserDataControllerUpdateProcedure = "/server.user.UserDataController/Update"
)

// UserDataControllerClient is a client for the server.user.UserDataController service.
type UserDataControllerClient interface {
	Register(context.Context, *connect_go.Request[user.UserRegisterRequest]) (*connect_go.Response[user.UserDataResponse], error)
	Update(context.Context, *connect_go.Request[user.UserUpdateDataRequest]) (*connect_go.Response[user.UserDataResponse], error)
}

// NewUserDataControllerClient constructs a client for the server.user.UserDataController service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserDataControllerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserDataControllerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userDataControllerClient{
		register: connect_go.NewClient[user.UserRegisterRequest, user.UserDataResponse](
			httpClient,
			baseURL+UserDataControllerRegisterProcedure,
			opts...,
		),
		update: connect_go.NewClient[user.UserUpdateDataRequest, user.UserDataResponse](
			httpClient,
			baseURL+UserDataControllerUpdateProcedure,
			opts...,
		),
	}
}

// userDataControllerClient implements UserDataControllerClient.
type userDataControllerClient struct {
	register *connect_go.Client[user.UserRegisterRequest, user.UserDataResponse]
	update   *connect_go.Client[user.UserUpdateDataRequest, user.UserDataResponse]
}

// Register calls server.user.UserDataController.Register.
func (c *userDataControllerClient) Register(ctx context.Context, req *connect_go.Request[user.UserRegisterRequest]) (*connect_go.Response[user.UserDataResponse], error) {
	return c.register.CallUnary(ctx, req)
}

// Update calls server.user.UserDataController.Update.
func (c *userDataControllerClient) Update(ctx context.Context, req *connect_go.Request[user.UserUpdateDataRequest]) (*connect_go.Response[user.UserDataResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// UserDataControllerHandler is an implementation of the server.user.UserDataController service.
type UserDataControllerHandler interface {
	Register(context.Context, *connect_go.Request[user.UserRegisterRequest]) (*connect_go.Response[user.UserDataResponse], error)
	Update(context.Context, *connect_go.Request[user.UserUpdateDataRequest]) (*connect_go.Response[user.UserDataResponse], error)
}

// NewUserDataControllerHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserDataControllerHandler(svc UserDataControllerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	userDataControllerRegisterHandler := connect_go.NewUnaryHandler(
		UserDataControllerRegisterProcedure,
		svc.Register,
		opts...,
	)
	userDataControllerUpdateHandler := connect_go.NewUnaryHandler(
		UserDataControllerUpdateProcedure,
		svc.Update,
		opts...,
	)
	return "/server.user.UserDataController/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserDataControllerRegisterProcedure:
			userDataControllerRegisterHandler.ServeHTTP(w, r)
		case UserDataControllerUpdateProcedure:
			userDataControllerUpdateHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserDataControllerHandler returns CodeUnimplemented from all methods.
type UnimplementedUserDataControllerHandler struct{}

func (UnimplementedUserDataControllerHandler) Register(context.Context, *connect_go.Request[user.UserRegisterRequest]) (*connect_go.Response[user.UserDataResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.user.UserDataController.Register is not implemented"))
}

func (UnimplementedUserDataControllerHandler) Update(context.Context, *connect_go.Request[user.UserUpdateDataRequest]) (*connect_go.Response[user.UserDataResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.user.UserDataController.Update is not implemented"))
}