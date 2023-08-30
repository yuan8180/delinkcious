package service

import (
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	sgm "github.com/the-gigi/delinkcious/pkg/milvus_manager"
	"log"
	"net/http"
)

var (
	// return when an expected path variable is missing.
	BadRoutingError = errors.New("inconsistent mapping between route and handler")
)

func Run() {
	store, err := sgm.NewMilvusStore("localhost", 27017)
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

	listCollectionHandler := httptransport.NewServer(
		makeListCollectionEndpoint(svc),
		decodeListCollectionRequest,
		encodeResponse,
	)

	hasCollectionHandler := httptransport.NewServer(
		makeHasCollectionEndpoint(svc),
		decodeHasCollectionRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path("/createCollection/{collectionname}").Handler(createCollectionHandler)
	r.Methods("POST").Path("/dropCollection/{collectionname}").Handler(dropCollectionHandler)
	r.Methods("GET").Path("/listCollection").Handler(listCollectionHandler)
	r.Methods("GET").Path("/hasCollection/{collectionname}").Handler(hasCollectionHandler)

	log.Println("Listening on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", r))
}
