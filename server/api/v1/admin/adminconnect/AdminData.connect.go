// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/admin/AdminData.proto

package adminconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
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
	// AdminDataControllerName is the fully-qualified name of the AdminDataController service.
	AdminDataControllerName = "server.admin.AdminDataController"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AdminDataControllerUpdateProcedure is the fully-qualified name of the AdminDataController's
	// Update RPC.
	AdminDataControllerUpdateProcedure = "/server.admin.AdminDataController/Update"
)

// AdminDataControllerClient is a client for the server.admin.AdminDataController service.
type AdminDataControllerClient interface {
	Update(context.Context, *connect_go.Request[admin.AdminUpdateDataRequest]) (*connect_go.Response[admin.AdminDataResponse], error)
}

// NewAdminDataControllerClient constructs a client for the server.admin.AdminDataController
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAdminDataControllerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AdminDataControllerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &adminDataControllerClient{
		update: connect_go.NewClient[admin.AdminUpdateDataRequest, admin.AdminDataResponse](
			httpClient,
			baseURL+AdminDataControllerUpdateProcedure,
			opts...,
		),
	}
}

// adminDataControllerClient implements AdminDataControllerClient.
type adminDataControllerClient struct {
	update *connect_go.Client[admin.AdminUpdateDataRequest, admin.AdminDataResponse]
}

// Update calls server.admin.AdminDataController.Update.
func (c *adminDataControllerClient) Update(ctx context.Context, req *connect_go.Request[admin.AdminUpdateDataRequest]) (*connect_go.Response[admin.AdminDataResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// AdminDataControllerHandler is an implementation of the server.admin.AdminDataController service.
type AdminDataControllerHandler interface {
	Update(context.Context, *connect_go.Request[admin.AdminUpdateDataRequest]) (*connect_go.Response[admin.AdminDataResponse], error)
}

// NewAdminDataControllerHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAdminDataControllerHandler(svc AdminDataControllerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	adminDataControllerUpdateHandler := connect_go.NewUnaryHandler(
		AdminDataControllerUpdateProcedure,
		svc.Update,
		opts...,
	)
	return "/server.admin.AdminDataController/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AdminDataControllerUpdateProcedure:
			adminDataControllerUpdateHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAdminDataControllerHandler returns CodeUnimplemented from all methods.
type UnimplementedAdminDataControllerHandler struct{}

func (UnimplementedAdminDataControllerHandler) Update(context.Context, *connect_go.Request[admin.AdminUpdateDataRequest]) (*connect_go.Response[admin.AdminDataResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AdminDataController.Update is not implemented"))
}
