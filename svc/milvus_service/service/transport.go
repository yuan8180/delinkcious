package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	om "github.com/yuan8180/delinkcious/pkg/object_model"
)

type getByCollnameRequest struct {
	Collname string `json:"collname"`
}

type createCollectionResponse struct {
	Err string `json:"err"`
}

type dropCollectionResponse struct {
	Err string `json:"err"`
}

type listCollectionsResponse struct {
	collections []string `json:"collections"`
	Err         string   `json:"err"`
}

type hasCollectionResponse struct {
	collExists bool   `json:"collExists"`
	Err        string `json:"err"`
}

func decodeCreateCollectionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getByCollnameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeDropCollectionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getByCollnameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeListCollectionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

func decodeHasCollectionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getByCollnameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func makeCreateCollectionEndpoint(svc om.MilvusManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getByCollnameRequest)
		err := svc.CreateCollection(req.Collname)
		res := createCollectionResponse{}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

func makeDropCollectionEndpoint(svc om.MilvusManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getByCollnameRequest)
		err := svc.DropCollection(req.Collname)
		res := dropCollectionResponse{}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

func makeListCollectionsEndpoint(svc om.MilvusManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		collections, err := svc.ListCollections()
		res := listCollectionsResponse{collections: collections}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

func makeHasCollectionEndpoint(svc om.MilvusManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getByCollnameRequest)
		collExists, err := svc.HasCollection(req.Collname)
		res := hasCollectionResponse{collExists: collExists}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}
