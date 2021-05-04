package endpoints

import (
	"context"

	"github.com/potatowhite/watermark-service/pkg/watermark"

	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	AddDocumentEndpoint   endpoint.Endpoint
	StatusEndpoint        endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	watermarkEndpoint     endpoint.Endpoint
}

func NewEndpointSet(svc watermark.Service) Set {
	return Set{
		GetEndpoint:           MakeGetEndpoint(svc),
		AddDocumentEndpoint:   MakeAddDocumentEndPoint(svc),
		StatusEndpoint:        MakeStatusEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceEndpoint(svc),
		WatermarkEndpoint:     MakeWatermarkEndpoint(svc),
	}
}

func MakeGetEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		docs, err := svc.Get(ctx, req.Filters...)

		if err != nil {
			return GetResponse{docs, err.Error()}, nil
		}

		return GetResponse{docs, ""}, nil
	}
}

func MakeStatusEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StatusRequest)
		status, err := svc.Status(ctx, req.TicketId)

		if err != nil {
			return StatusResponse{Status: status, Err: err.Error()}, nil
		}

		return StatusResponse{Status: status, Err: ""}, nil

	}
}

func MakeAddDocumentEndPoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(AddDocumentRequest)
		ticketId, err = svc.AddDocument(ctx, req.Document)

		if err != nil {
			return AddDocumentResponse{TicketId: ticketId, Err: err.Error()}, nil
		}
		return AddDocumentResponse{TicketId: ticketId, Err: ""}, nil

	}
}

func MakeWatermarkEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(WatermarkRequest)
		code, err = svc.Watermark(ctx, req.TicketId, req.Mark)
	
		if err != nil {
			return WatermarkResponse{Code: code, Err: err.Error()}, nil
		}
		return WatermarkResponse{Code: code, Err: ""}, nil
	}
}


func MakeServiceStatusEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context), request interface{}) (interface{}, error) {

		_ := request.(ServiceStatusRequest)
		code, err = svc.ServiceStatus(ctx)
	
		if err != nil {
			return ServiceStatusResponse{Code: code, Err :err.Error()};
		}
	
		return ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}



func (s *Set) Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest{Filters: filters})
	if err != nil {
		return []internal.Document{}, err
	}

	getResp := resp.(GetResponse)
	if getResp.Err != "" {
		return []internal.Document(), errors.New(getResp.Err)
	}

	return getResp.Documents, nil

}

func (S *Set) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest())
	svcStatusResp := resp.(ServiceStatusResponse)

	if err != nil {
		return svcStatusResp.Code, err
	}

	if svcStatusResp.Err != "" {
		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
	}

	return svcStatusResp.Code, nil
}

func (S *Set) AddDocument(ctx context.Context, doc *internal.Document) (string, error) {
	resp, err := s.AddDocumentEndpoint(ctx, AddDocumentRequest{Document: doc})

	if err != nil {
		return "", err
	}

	adResp := resp.(AddDocumentResponse)
	if(adResp.Err != "") {
		return "",errors.New(adResp.Err)
	}

	return adResp.TicketId, nil
}

func (s *Set) Status(ctx context.Context, ticketId string) (internal.Status, error){
	resp, err := s.StatusEndpoint(ctx, StatusRequest{TicketId: ticketId})

	if err != nil {
		return internal.Failed, err
	}

	stsResp := resp.(StatusResponse)
	if stsResp.Err != "" {
		return internal.Failed, errors.New(stsResp.Err)
	}

	return stsResp.Status, nil
}

func (s *Set) Watermark(ctx context.Context, ticketId, mark string) (int, error) {
	resp, err := s.WatermarkEndpoint(ctx, WatermarkRequest{TicketId: ticketId, Mark: mark})
	
	wmResp := resp.(WatermarkResponse)
	if wmResp.Err != nil {
		return wmResp.Code, err
	}

	if wmResp.Err != "" {
		return wmResp.Code, errors.New(wmResp.Err)
	}

	return wmResp.Code, nil
}


var logger log.Logger

func init(){
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}