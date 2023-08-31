package social_graph_client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
	om "github.com/yuan8180/delinkcious/pkg/object_model"
)

func NewClient(baseURL string) (om.MilvusManager, error) {
	// Quickly sanitize the instance string.
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "http://" + baseURL
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	createCollectionEndpoint := httptransport.NewClient(
		"POST",
		copyURL(u, "/createCollection"),
		encodeGetByCollnameRequest,
		decodeSimpleResponse).Endpoint()

	dropCollectionEndpoint := httptransport.NewClient(
		"POST",
		copyURL(u, "/dropCollection"),
		encodeGetByCollnameRequest,
		decodeSimpleResponse).Endpoint()

	listCollectionsEndpoint := httptransport.NewClient(
		"GET",
		copyURL(u, "/listCollections"),
		encodeHTTPGenericRequest,
		decodeListCollectionsResponse).Endpoint()

	hasCollectionEndpoint := httptransport.NewClient(
		"GET",
		copyURL(u, "/hasCollection"),
		encodeGetByCollnameRequest,
		decodeHasCollectionResponse).Endpoint()

	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return EndpointSet{
		CreateCollectionEndpoint: createCollectionEndpoint,
		DropCollectionEndpoint:   dropCollectionEndpoint,
		ListCollectionsEndpoint:  listCollectionsEndpoint,
		HasCollectionEndpoint:    hasCollectionEndpoint,
	}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

// encodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// JSON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// Extract the collname from the incmoing request and add it to the path
func encodeGetByCollnameRequest(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(getByCollNameRequest)
	collname := url.QueryEscape(r.Collname)
	req.URL.Path += "/" + collname
	return encodeHTTPGenericRequest(ctx, req, request)
}
