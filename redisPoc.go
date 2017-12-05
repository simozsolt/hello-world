package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/simozsolt/hello-world/redisPoc"
	"log"
	"os"
	"strconv"
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

func resetDbProcess(client *redis.Client) {
	resetDb(client)

	filePath := "/home/simo/Work/go/src/github.com/simozsolt/hello-world/gts_nat.csv"
	pData, error := redisPoc.GetPricelistRowsFromCsv(filePath, "gts_nat")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("PricelistId: %s; RowCount: %d\n", pData.PricelistId, len(pData.Data))
	//fmt.Printf("%s \n", pData)

	pData.InsertToDb(client)

	filePath2 := "/home/simo/Work/go/src/github.com/simozsolt/hello-world/gts_int.csv"
	pData2, error2 := redisPoc.GetPricelistRowsFromCsv(filePath2, "gts_int")
	if error2 != nil {
		log.Fatal(error)
	}
	fmt.Printf("PricelistId: %s; RowCount: %d\n", pData2.PricelistId, len(pData2.Data))
	//fmt.Printf("%s \n", pData2)

	pData2.InsertToDb(client)
}

func getArgs() (resetDbParam bool, prefix string, searchLength int) {
	resetDbParam = false
	if os.Args[1] == "true" {
		resetDbParam = true
	}

	prefix = os.Args[2]
	searchLength, _ = strconv.Atoi(os.Args[3])

	return
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Required params: resetDb(true|false) prefix(phoneNr) searchLength(int)")
		return
	}

	resetDbParam, prefix, searchLength := getArgs()
	fmt.Printf("ResetDb: %+v; Prefix: %s; SearchLength: %d\n", resetDbParam, prefix, searchLength)

	client := getClient("localhost:6379", "", 0)
	if resetDbParam {
		resetDbProcess(client)
	}

	redisPoc.LookUp(client, "gts_nat", "hu", prefix, searchLength)
}

/*
https://github.com/go-redis/redis/blob/master/commands.go#L904

https://gobyexample.com/
*/
