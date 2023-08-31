package service

import (
	"errors"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	sgm "github.com/yuan8180/delinkcious/pkg/milvus_manager"
)

var (
	// return when an expected path variable is missing.
	BadRoutingError = errors.New("inconsistent mapping between route and handler")
)

func Run() {
	store, err := sgm.NewMilvusStore("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	svc, err := sgm.NewMilvusManager(store)
	if err != nil {
		log.Fatal(err)
	}

	createCollectionHandler := httptransport.NewServer(
		makeCreateCollectionEndpoint(svc),
		decodeCreateCollectionRequest,
		encodeResponse,
	)

	dropCollectionHandler := httptransport.NewServer(
		makeDropCollectionEndpoint(svc),
		decodeDropCollectionRequest,
		encodeResponse,
	)

	listCollectionsHandler := httptransport.NewServer(
		makeListCollectionsEndpoint(svc),
		decodeListCollectionsRequest,
		encodeResponse,
	)

	hasCollectionHandler := httptransport.NewServer(
		makeHasCollectionEndpoint(svc),
		decodeHasCollectionRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path("/createCollection/{collname}").Handler(createCollectionHandler)
	r.Methods("POST").Path("/dropCollection/{collname}").Handler(dropCollectionHandler)
	r.Methods("GET").Path("/listCollections").Handler(listCollectionsHandler)
	r.Methods("GET").Path("/hasCollection/{collname}").Handler(hasCollectionHandler)

	log.Println("Listening on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", r))
}
