package social_graph_manager

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	milvus "github.com/milvus-io/milvus-sdk-go/v2/client"
	entity "github.com/milvus-io/milvus-sdk-go/v2/entity"
)

type MilvusStore struct {
	ctx context.Context
	c   milvus.Client
}

func NewMilvusStore(milvusAddr string) (store *MilvusStore, err error) {
	milvusCtx := context.Background()
	c, err := milvus.NewClient(milvusCtx, milvus.Config{
		Address: milvusAddr,
	})
	if err != nil {
		log.Fatalf("failed to connect to milvus, err: %v", err)
		return
	}

	store = &MilvusStore{milvusCtx, c}
	return
}

func (s *MilvusStore) CreateCollection(collname string) (err error) {
	err = s.DropCollection(collname)
	if err != nil {
		return
	}

	// define collection schema
	schema := entity.NewSchema().WithName(collname).WithDescription("this is the basic example collection").
		// currently primary key field is compulsory, and only int64 is allowed
		WithField(entity.NewField().WithName("ID").WithDataType(entity.FieldTypeInt64).WithIsPrimaryKey(true).WithIsAutoID(false)).
		// also the vector field is needed
		WithField(entity.NewField().WithName("embeddings").WithDataType(entity.FieldTypeFloatVector).WithDim(128))

	err = s.c.CreateCollection(s.ctx, schema, entity.DefaultShardNumber)
	if err != nil {
		log.Fatal("failed to create collection:", err.Error())
	}
	return
}

func (s *MilvusStore) DropCollection(collname string) (err error) {
	fmt.Println("start drop")
	collExists, err := s.HasCollection(collname)
	if !collExists {
		log.Println("no specified collection was not found")
		return
	}
	err = s.c.DropCollection(s.ctx, collname)
	if err != nil {
		log.Fatalf("failed to drop collection, err: %v", err)
		return
	}

	return
}

func (s *MilvusStore) ListCollections() (collections []string, err error) {
	colls, err := s.c.ListCollections(s.ctx)
	if err != nil {
		log.Fatal("failed to list collections:", err.Error())
	}
	for i, collection := range colls {
		collections[i] = collection.Name
	}

	return
}

func (s *MilvusStore) HasCollection(collname string) (collExists bool, err error) {
	collExists, err = s.c.HasCollection(s.ctx, collname)
	if err != nil {
		log.Fatalf("failed to check collection exists, err: %v", err)
		return
	}

	return
}
