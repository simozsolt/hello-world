package main

import (
    "fmt"
    "strings"
    "github.com/go-redis/redis"
)

type PricelistRow struct {
    Prefix string
    Rate float32
    Description string
}

func getPricelistRows() []PricelistRow  {
    pricelistRows := []PricelistRow {
        {Prefix: "36", Rate: 2.0, Description: "Landline"},
        {Prefix: "361", Rate: 1.0, Description: "Budapest"},
        {Prefix: "3620", Rate: 11.0, Description: "Telenor"},
        {Prefix: "3630", Rate: 12.0, Description: "T-Mobile"},
        {Prefix: "3670", Rate: 13.0, Description: "Vodafone"},
    }

    return pricelistRows
}

func printPricelistRows(pricelistRows []PricelistRow) {
    fmt.Println("Pricelist rows:")
    for index, value := range pricelistRows {
        fmt.Printf("Index: %d, %+v\n", index, value)
    }
}

func getKey(pricelistId string, countryCode string, sufix string) string {
    s := []string{"pricelists:", pricelistId, "/countries:", countryCode, "/", sufix}
    return strings.Join(s, "")
}

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

func loadPricelist(client *redis.Client, pricelistId string, countryCode string, pricelistRows []PricelistRow) {
    fmt.Println("Load pricelist into Redis")

    hmset := func(client *redis.Client, key string, pricelistRow PricelistRow) *redis.StringCmd {
/*        args := make([]string, 8)
        args[0] = "hmset"
        args[1] = key
        args[2] = "Prefix"
        args[3] = pricelistRow.Prefix
        args[4] = "Rate"
        args[5] = string(pricelistRow.Rate)
        args[6] = "Description"
        args[7] = pricelistRow.Description

        cmd := redis.NewStatusCmd(args)
*/
        //fmt.Printf("%+v", cmd)
        //client.Process(cmd)
        cmd := redis.NewStringCmd("hmset", key, "Prefix", pricelistRow.Prefix, "Rate", pricelistRow.Rate, "Description", pricelistRow.Description)
        client.Process(cmd)
        return cmd
    }

    for index, row := range pricelistRows {
        fmt.Println("Insert:", index)
        key := getKey(pricelistId, countryCode, row.Prefix)
        _, err := hmset(client, key, row).Result()
        if err != nil {
            fmt.Println("Error on hmset")
        }
    }
}

func main() {
    client := getClient("localhost:6379", "", 0)
    resetDb(client)

    pricelistRows := getPricelistRows()
    printPricelistRows(pricelistRows)

    loadPricelist(client, "telekom", "hu", pricelistRows)

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

}

/*
init db

set valami1 ertek1
set valami2 ertek2
hmset pricelists elso telekom masodik upc

 */
