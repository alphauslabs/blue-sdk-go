// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: pricing/v1/pricing.proto

/*
Package pricing is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pricing

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_Pricing_GetInfo_0(ctx context.Context, marshaler runtime.Marshaler, client PricingClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetInfoRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetInfo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Pricing_GetInfo_0(ctx context.Context, marshaler runtime.Marshaler, server PricingServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetInfoRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetInfo(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterPricingHandlerServer registers the http handlers for service Pricing to "mux".
// UnaryRPC     :call PricingServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterPricingHandlerFromEndpoint instead.
func RegisterPricingHandlerServer(ctx context.Context, mux *runtime.ServeMux, server PricingServer) error {

	mux.Handle("GET", pattern_Pricing_GetInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/blueapi.pricing.v1.Pricing/GetInfo", runtime.WithHTTPPathPattern("/v0/info"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Pricing_GetInfo_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Pricing_GetInfo_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterPricingHandlerFromEndpoint is same as RegisterPricingHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterPricingHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterPricingHandler(ctx, mux, conn)
}

// RegisterPricingHandler registers the http handlers for service Pricing to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterPricingHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterPricingHandlerClient(ctx, mux, NewPricingClient(conn))
}

// RegisterPricingHandlerClient registers the http handlers for service Pricing
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "PricingClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "PricingClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "PricingClient" to call the correct interceptors.
func RegisterPricingHandlerClient(ctx context.Context, mux *runtime.ServeMux, client PricingClient) error {

	mux.Handle("GET", pattern_Pricing_GetInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/blueapi.pricing.v1.Pricing/GetInfo", runtime.WithHTTPPathPattern("/v0/info"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Pricing_GetInfo_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Pricing_GetInfo_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Pricing_GetInfo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v0", "info"}, ""))
)

var (
	forward_Pricing_GetInfo_0 = runtime.ForwardResponseMessage
)