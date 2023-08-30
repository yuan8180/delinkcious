package service

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	om "github.com/the-gigi/delinkcious/pkg/object_model"
	"net/http"
)

type followRequest struct {
	Followed string `json:"followed"`
	Follower string `json:"follower"`
}

type followResponse struct {
	Err string `json:"err"`
}

type unfollowRequest struct {
	Followed string `json:"followed"`
	Follower string `json:"follower"`
}

type unfollowResponse struct {
	Err string `json:"err"`
}

type getByUsernameRequest struct {
	Username string `json:"username"`
}

type getFollowersResponse struct {
	Followers map[string]bool `json:"followers"`
	Err       string          `json:"err"`
}

type getFollowingResponse struct {
	Following map[string]bool `json:"following"`
	Err       string          `json:"err"`
}

func decodeFollowRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request followRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUnfollowRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request unfollowRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetFollowingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getByUsernameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetFollowersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getByUsernameRequest
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
		req := request.(GetByCollectionNameRequest)
		err := svc.CreateCollection(req.Collectionname)
		res := createCollectionResponse{}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

func makeDropCollectionEndpoint(svc om.MilvusManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByCollectionNameRequest)
		err := svc.DropCollection(req.Collectionname)
		res := dropCollectionResponse{}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

func makeListCollectionEndpoint(svc om.MilvusManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(listCollectionRequest)
		collectionMap, err := svc.ListCollection{}
                res := listCollectionResponse{Collections: collectionMap}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

func makeHasCollectionEndpoint(svc om.MilvusManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByCollectionNameRequest)
		hasCollection, err := svc.HasCollection(req.Collectionname)
		res := hasCollectionResponse{HasCollection: hasCollection}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}
