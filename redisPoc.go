package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/simozsolt/hello-world/redisPoc"
	"log"
)

func getClient(host string, pwd string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pwd, // no password set
		DB:       db,  // use default DB
	})

	return client
}

func resetDb(client *redis.Client) {
	cmd := redis.NewStringCmd("FLUSHALL")
	client.Process(cmd)
}

func main() {
	client := getClient("localhost:6379", "", 0)
	resetDb(client)

	filePath := "/home/simo/Work/go/src/github.com/simozsolt/hello-world/gts_nat.csv"
	pData, error := redisPoc.GetPricelistRowsFromCsv(filePath, "gts_nat")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("%s \n", pData)

	pData.InsertToDb(client)

	filePath2 := "/home/simo/Work/go/src/github.com/simozsolt/hello-world/gts_int.csv"
	pData2, error2 := redisPoc.GetPricelistRowsFromCsv(filePath2, "gts_int")
	if error2 != nil {
		log.Fatal(error)
	}
	fmt.Printf("%s \n", pData2)

	pData2.InsertToDb(client)
}

/*
https://github.com/go-redis/redis/blob/master/commands.go#L904

https://gobyexample.com/
*/
