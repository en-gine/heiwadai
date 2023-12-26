// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/admin/Coupon.proto

package adminconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	admin "server/api/v1/admin"
	shared "server/api/v1/shared"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// AdminCouponControllerName is the fully-qualified name of the AdminCouponController service.
	AdminCouponControllerName = "server.admin.AdminCouponController"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AdminCouponControllerGetUserCouponListProcedure is the fully-qualified name of the
	// AdminCouponController's GetUserCouponList RPC.
	AdminCouponControllerGetUserCouponListProcedure = "/server.admin.AdminCouponController/GetUserCouponList"
	// AdminCouponControllerCreateCustomCouponProcedure is the fully-qualified name of the
	// AdminCouponController's CreateCustomCoupon RPC.
	AdminCouponControllerCreateCustomCouponProcedure = "/server.admin.AdminCouponController/CreateCustomCoupon"
	// AdminCouponControllerSaveCustomCouponProcedure is the fully-qualified name of the
	// AdminCouponController's SaveCustomCoupon RPC.
	AdminCouponControllerSaveCustomCouponProcedure = "/server.admin.AdminCouponController/SaveCustomCoupon"
	// AdminCouponControllerGetCustomCouponByIDProcedure is the fully-qualified name of the
	// AdminCouponController's GetCustomCouponByID RPC.
	AdminCouponControllerGetCustomCouponByIDProcedure = "/server.admin.AdminCouponController/GetCustomCouponByID"
	// AdminCouponControllerGetCustomCouponListProcedure is the fully-qualified name of the
	// AdminCouponController's GetCustomCouponList RPC.
	AdminCouponControllerGetCustomCouponListProcedure = "/server.admin.AdminCouponController/GetCustomCouponList"
	// AdminCouponControllerAttachCustomCouponToAllUserProcedure is the fully-qualified name of the
	// AdminCouponController's AttachCustomCouponToAllUser RPC.
	AdminCouponControllerAttachCustomCouponToAllUserProcedure = "/server.admin.AdminCouponController/AttachCustomCouponToAllUser"
)

// AdminCouponControllerClient is a client for the server.admin.AdminCouponController service.
type AdminCouponControllerClient interface {
	GetUserCouponList(context.Context, *connect_go.Request[admin.UserIDRequest]) (*connect_go.Response[admin.UserAttachedCouponsResponse], error)
	CreateCustomCoupon(context.Context, *connect_go.Request[admin.CreateCustomCouponRequest]) (*connect_go.Response[emptypb.Empty], error)
	SaveCustomCoupon(context.Context, *connect_go.Request[admin.CouponIDRequest]) (*connect_go.Response[emptypb.Empty], error)
	GetCustomCouponByID(context.Context, *connect_go.Request[admin.CouponIDRequest]) (*connect_go.Response[shared.Coupon], error)
	GetCustomCouponList(context.Context, *connect_go.Request[shared.Pager]) (*connect_go.Response[admin.CouponListResponse], error)
	AttachCustomCouponToAllUser(context.Context, *connect_go.Request[admin.CouponIDRequest]) (*connect_go.Response[admin.AffectedCountResponse], error)
}

// NewAdminCouponControllerClient constructs a client for the server.admin.AdminCouponController
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAdminCouponControllerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AdminCouponControllerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &adminCouponControllerClient{
		getUserCouponList: connect_go.NewClient[admin.UserIDRequest, admin.UserAttachedCouponsResponse](
			httpClient,
			baseURL+AdminCouponControllerGetUserCouponListProcedure,
			opts...,
		),
		createCustomCoupon: connect_go.NewClient[admin.CreateCustomCouponRequest, emptypb.Empty](
			httpClient,
			baseURL+AdminCouponControllerCreateCustomCouponProcedure,
			opts...,
		),
		saveCustomCoupon: connect_go.NewClient[admin.CouponIDRequest, emptypb.Empty](
			httpClient,
			baseURL+AdminCouponControllerSaveCustomCouponProcedure,
			opts...,
		),
		getCustomCouponByID: connect_go.NewClient[admin.CouponIDRequest, shared.Coupon](
			httpClient,
			baseURL+AdminCouponControllerGetCustomCouponByIDProcedure,
			opts...,
		),
		getCustomCouponList: connect_go.NewClient[shared.Pager, admin.CouponListResponse](
			httpClient,
			baseURL+AdminCouponControllerGetCustomCouponListProcedure,
			opts...,
		),
		attachCustomCouponToAllUser: connect_go.NewClient[admin.CouponIDRequest, admin.AffectedCountResponse](
			httpClient,
			baseURL+AdminCouponControllerAttachCustomCouponToAllUserProcedure,
			opts...,
		),
	}
}

// adminCouponControllerClient implements AdminCouponControllerClient.
type adminCouponControllerClient struct {
	getUserCouponList           *connect_go.Client[admin.UserIDRequest, admin.UserAttachedCouponsResponse]
	createCustomCoupon          *connect_go.Client[admin.CreateCustomCouponRequest, emptypb.Empty]
	saveCustomCoupon            *connect_go.Client[admin.CouponIDRequest, emptypb.Empty]
	getCustomCouponByID         *connect_go.Client[admin.CouponIDRequest, shared.Coupon]
	getCustomCouponList         *connect_go.Client[shared.Pager, admin.CouponListResponse]
	attachCustomCouponToAllUser *connect_go.Client[admin.CouponIDRequest, admin.AffectedCountResponse]
}

// GetUserCouponList calls server.admin.AdminCouponController.GetUserCouponList.
func (c *adminCouponControllerClient) GetUserCouponList(ctx context.Context, req *connect_go.Request[admin.UserIDRequest]) (*connect_go.Response[admin.UserAttachedCouponsResponse], error) {
	return c.getUserCouponList.CallUnary(ctx, req)
}

// CreateCustomCoupon calls server.admin.AdminCouponController.CreateCustomCoupon.
func (c *adminCouponControllerClient) CreateCustomCoupon(ctx context.Context, req *connect_go.Request[admin.CreateCustomCouponRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.createCustomCoupon.CallUnary(ctx, req)
}

// SaveCustomCoupon calls server.admin.AdminCouponController.SaveCustomCoupon.
func (c *adminCouponControllerClient) SaveCustomCoupon(ctx context.Context, req *connect_go.Request[admin.CouponIDRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.saveCustomCoupon.CallUnary(ctx, req)
}

// GetCustomCouponByID calls server.admin.AdminCouponController.GetCustomCouponByID.
func (c *adminCouponControllerClient) GetCustomCouponByID(ctx context.Context, req *connect_go.Request[admin.CouponIDRequest]) (*connect_go.Response[shared.Coupon], error) {
	return c.getCustomCouponByID.CallUnary(ctx, req)
}

// GetCustomCouponList calls server.admin.AdminCouponController.GetCustomCouponList.
func (c *adminCouponControllerClient) GetCustomCouponList(ctx context.Context, req *connect_go.Request[shared.Pager]) (*connect_go.Response[admin.CouponListResponse], error) {
	return c.getCustomCouponList.CallUnary(ctx, req)
}

// AttachCustomCouponToAllUser calls server.admin.AdminCouponController.AttachCustomCouponToAllUser.
func (c *adminCouponControllerClient) AttachCustomCouponToAllUser(ctx context.Context, req *connect_go.Request[admin.CouponIDRequest]) (*connect_go.Response[admin.AffectedCountResponse], error) {
	return c.attachCustomCouponToAllUser.CallUnary(ctx, req)
}

// AdminCouponControllerHandler is an implementation of the server.admin.AdminCouponController
// service.
type AdminCouponControllerHandler interface {
	GetUserCouponList(context.Context, *connect_go.Request[admin.UserIDRequest]) (*connect_go.Response[admin.UserAttachedCouponsResponse], error)
	CreateCustomCoupon(context.Context, *connect_go.Request[admin.CreateCustomCouponRequest]) (*connect_go.Response[emptypb.Empty], error)
	SaveCustomCoupon(context.Context, *connect_go.Request[admin.CouponIDRequest]) (*connect_go.Response[emptypb.Empty], error)
	GetCustomCouponByID(context.Context, *connect_go.Request[admin.CouponIDRequest]) (*connect_go.Response[shared.Coupon], error)
	GetCustomCouponList(context.Context, *connect_go.Request[shared.Pager]) (*connect_go.Response[admin.CouponListResponse], error)
	AttachCustomCouponToAllUser(context.Context, *connect_go.Request[admin.CouponIDRequest]) (*connect_go.Response[admin.AffectedCountResponse], error)
}

// NewAdminCouponControllerHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAdminCouponControllerHandler(svc AdminCouponControllerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	adminCouponControllerGetUserCouponListHandler := connect_go.NewUnaryHandler(
		AdminCouponControllerGetUserCouponListProcedure,
		svc.GetUserCouponList,
		opts...,
	)
	adminCouponControllerCreateCustomCouponHandler := connect_go.NewUnaryHandler(
		AdminCouponControllerCreateCustomCouponProcedure,
		svc.CreateCustomCoupon,
		opts...,
	)
	adminCouponControllerSaveCustomCouponHandler := connect_go.NewUnaryHandler(
		AdminCouponControllerSaveCustomCouponProcedure,
		svc.SaveCustomCoupon,
		opts...,
	)
	adminCouponControllerGetCustomCouponByIDHandler := connect_go.NewUnaryHandler(
		AdminCouponControllerGetCustomCouponByIDProcedure,
		svc.GetCustomCouponByID,
		opts...,
	)
	adminCouponControllerGetCustomCouponListHandler := connect_go.NewUnaryHandler(
		AdminCouponControllerGetCustomCouponListProcedure,
		svc.GetCustomCouponList,
		opts...,
	)
	adminCouponControllerAttachCustomCouponToAllUserHandler := connect_go.NewUnaryHandler(
		AdminCouponControllerAttachCustomCouponToAllUserProcedure,
		svc.AttachCustomCouponToAllUser,
		opts...,
	)
	return "/server.admin.AdminCouponController/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AdminCouponControllerGetUserCouponListProcedure:
			adminCouponControllerGetUserCouponListHandler.ServeHTTP(w, r)
		case AdminCouponControllerCreateCustomCouponProcedure:
			adminCouponControllerCreateCustomCouponHandler.ServeHTTP(w, r)
		case AdminCouponControllerSaveCustomCouponProcedure:
			adminCouponControllerSaveCustomCouponHandler.ServeHTTP(w, r)
		case AdminCouponControllerGetCustomCouponByIDProcedure:
			adminCouponControllerGetCustomCouponByIDHandler.ServeHTTP(w, r)
		case AdminCouponControllerGetCustomCouponListProcedure:
			adminCouponControllerGetCustomCouponListHandler.ServeHTTP(w, r)
		case AdminCouponControllerAttachCustomCouponToAllUserProcedure:
			adminCouponControllerAttachCustomCouponToAllUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAdminCouponControllerHandler returns CodeUnimplemented from all methods.
type UnimplementedAdminCouponControllerHandler struct{}

func (UnimplementedAdminCouponControllerHandler) GetUserCouponList(context.Context, *connect_go.Request[admin.UserIDRequest]) (*connect_go.Response[admin.UserAttachedCouponsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AdminCouponController.GetUserCouponList is not implemented"))
}

func (UnimplementedAdminCouponControllerHandler) CreateCustomCoupon(context.Context, *connect_go.Request[admin.CreateCustomCouponRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AdminCouponController.CreateCustomCoupon is not implemented"))
}

func (UnimplementedAdminCouponControllerHandler) SaveCustomCoupon(context.Context, *connect_go.Request[admin.CouponIDRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AdminCouponController.SaveCustomCoupon is not implemented"))
}

func (UnimplementedAdminCouponControllerHandler) GetCustomCouponByID(context.Context, *connect_go.Request[admin.CouponIDRequest]) (*connect_go.Response[shared.Coupon], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AdminCouponController.GetCustomCouponByID is not implemented"))
}

func (UnimplementedAdminCouponControllerHandler) GetCustomCouponList(context.Context, *connect_go.Request[shared.Pager]) (*connect_go.Response[admin.CouponListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AdminCouponController.GetCustomCouponList is not implemented"))
}

func (UnimplementedAdminCouponControllerHandler) AttachCustomCouponToAllUser(context.Context, *connect_go.Request[admin.CouponIDRequest]) (*connect_go.Response[admin.AffectedCountResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AdminCouponController.AttachCustomCouponToAllUser is not implemented"))
}
