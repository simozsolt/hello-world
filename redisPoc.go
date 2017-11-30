package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/simozsolt/hello-world/redisPoc"
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
	pData := redisPoc.GetPricelistRowsFromCsv(filePath, "gts_nat")
	fmt.Printf("%s \n", pData)

	pData.InsertToDb(client)

	filePath2 := "/home/simo/Work/go/src/github.com/simozsolt/hello-world/gts_int.csv"
	pData2 := redisPoc.GetPricelistRowsFromCsv(filePath2, "gts_int")
	fmt.Printf("%s \n", pData2)

	pData2.InsertToDb(client)
}

/*
https://github.com/go-redis/redis/blob/master/commands.go#L904

https://gobyexample.com/
*/
