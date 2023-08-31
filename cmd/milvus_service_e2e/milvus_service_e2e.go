package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/yuan8180/delinkcious/pkg/milvus_client"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Run some tests with the client
	cli, err := milvus_client.NewClient("localhost:9090")
	check(err)

	collections, err := cli.ListCollections()
	check(err)
	log.Print("1list collections:")
	for _, collection := range collections {
		log.Print(collection)
	}

	collExists, err := cli.HasCollection("milvus01")
	check(err)
	log.Print("1whether the collection is exits:", collExists)

	err = cli.CreateCollection("milvus01")
	check(err)

	collections, err := cli.ListCollections()
	check(err)
	log.Print("2list collections:")
	for _, collection := range collections {
		log.Print(collection)
	}

	err = cli.DropCollection("milvus01")
	check(err)

	collExists, err := cli.HasCollection("milvus01")
	check(err)
	log.Print("2whether the collection is exits:", collExists)

}
