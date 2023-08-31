package milvus_client

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type SimpleResponse struct {
	Err string
}

func decodeSimpleResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp SimpleResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

type EndpointSet struct {
	CreateCollectionEndpoint endpoint.Endpoint
	DropCollectionEndpoint   endpoint.Endpoint
	ListCollectionsEndpoint  endpoint.Endpoint
	HasCollectionEndpoint    endpoint.Endpoint
}

type getByCollNameRequest struct {
	Collname string `json:"collname"`
}

func (s EndpointSet) CreateCollection(collname string) (err error) {
	resp, err := s.CreateCollectionEndpoint(context.Background(), getByCollNameRequest{Collname: collname})
	if err != nil {
		return err
	}
	response := resp.(SimpleResponse)

	if response.Err != "" {
		err = errors.New(response.Err)
	}
	return
}

func (s EndpointSet) DropCollection(collname string) (err error) {
	resp, err := s.DropCollectionEndpoint(context.Background(), getByCollNameRequest{Collname: collname})
	if err != nil {
		return err
	}
	response := resp.(SimpleResponse)
	if response.Err != "" {
		err = errors.New(response.Err)
	}
	return
}

type listCollectionsResponse struct {
	Collections []string `json:"collections"`
	Err         string   `json:"err"`
}

func decodeListCollectionsResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp listCollectionsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func (s EndpointSet) ListCollections() (collections []string, err error) {
	resp, err := s.ListCollectionsEndpoint(context.Background(), struct{}{})
	if err != nil {
		return
	}

	response := resp.(listCollectionsResponse)
	if response.Err != "" {
		err = errors.New(response.Err)
	}
	for i, collection := range response.Collections {
		collections[i] = collection
	}
	return
}

type hasCollectionResponse struct {
	CollExists bool   `json:"collExists"`
	Err        string `json:"err"`
}

func decodeHasCollectionResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp hasCollectionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func (s EndpointSet) HasCollection(collname string) (collExists bool, err error) {
	resp, err := s.HasCollectionEndpoint(context.Background(), getByCollNameRequest{Collname: collname})
	if err != nil {
		return
	}

	response := resp.(hasCollectionResponse)
	if response.Err != "" {
		err = errors.New(response.Err)
	}
	collExists = response.CollExists
	return
}
