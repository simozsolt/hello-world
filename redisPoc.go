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

	pData := redisPoc.GetPricelistRows()
	fmt.Printf("%s \n", pData)

	pData.InsertToDb(client)

	//

}

/*
   Get := func(client *redis.Client, key string) *redis.StringCmd {
	   cmd := redis.NewStringCmd("get", key)
	   client.Process(cmd)
	   return cmd
   }

   HgetAll := func (client *redis.Client, key string) *redis.StringStringMapCmd {
	   cmd := redis.NewStringStringMapCmd("hgetall", key)
	   client.Process(cmd)
	   return cmd
   }

   v, err := Get(client, "not_exists").Result()

   if (err != nil) {
	   fmt.Printf("%s\n", err)
   }
   fmt.Printf("%v\n", v)

   v2, err2 := HgetAll(client, "kulcs").Result();
   if (err2 != nil) {
	   fmt.Printf("%s\n", err2)
   }
   fmt.Printf("%v\n\n", v2)

   for index, value := range v2 {
	   fmt.Println(index, value)
   }*/

/*
https://github.com/go-redis/redis/blob/master/commands.go#L904

https://gobyexample.com/
*/
