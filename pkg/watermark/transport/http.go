package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/potatowhite/watermark-service/internal/util"
	"github.com/potatowhite/watermark=service/pkg/endpoints"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	m := http.NewServeMux()

	m.Handle("/healthz", httptransport.NewServer(
		ep.ServiceStatusEndpoint,
		decodeHTTPServiceStatusReqeust,
		encodeResponse,
	))
	
	m.Handle("/status". httptransport.NewServer(
		ep.StatusEndpoint,
		decodeHTTPStatusRequest,
		encodeResponse,
	))

	m.Handle("/addDocument". httptransport.NewServer(
		ep.AddDocumentEndpoint,
		decodeHTTPAddDocumentRequest,
		encodeResponse,
	))


	m.Handle("/get". httptransport.NewServer(
		ep.GetEndpoint,
		decodeHTTPGetRequest,
		encodeResponse,
	))


	m.Handle("/watermark". httptransport.NewServer(
		ep.WatermarkEndpoint,
		decodeHTTPWatermarkRequest,
		encodeResponse,
	))

	return m;
}


func decodeHTTPGetRequest(_ context.Context, r *http.Request) (interface{}, error){
	var req endpoint.GetRequest
	
	if r.ContentLength == 0 {
		logger.Log("Get request with no body")
		return req, nil
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeHTTPStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.StatusRequest
	err := json.NewDecoder(r.body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeHTTPWatermarkRequest(_ context.Context, r *http.Request) (interface{}, error){
	var req endpoint.WatermarkRequest
	err := json.NewDecoder(r.body).Decode(&req)

	if err != nil{
		return nil, err
	}

	return req, nil
}

func decodeHTTPAddDocumentRequest(_ context.Context, r *http.Request) (interface{}, error){
	var req endpoint.AddDocumentRequest
	err := json.NewDecoder(r.body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeHTTPServiceStatusRequest(_ context.Context, _ *httphttp.Request) (interface{}, error){
	var req endpoint.ServiceStatusRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response; ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch err {
	case util.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case util.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
			w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

var logger log.Logger

func init() {
	logger = log.NewLoggerfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}