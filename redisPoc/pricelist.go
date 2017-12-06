package redisPoc

import (
	"fmt"
	"github.com/go-redis/redis"
	"strings"
	"time"
	"log"
	"errors"
)

type PricelistRow struct {
	Prefix      string
	Rate        float64
	CountryCode string
	Description string

	Optional map[string]interface{}
}

func (pRow PricelistRow) String() string {
	//return fmt.Sprintf("Prefix: %s Rate: %f Description: %s %+v", pRow.Prefix, pRow.Rate, pRow.Description, pRow.Optional)
	return fmt.Sprintf("Prefix: %s Rate: %f Country: %s Description: %s", pRow.Prefix, pRow.Rate, pRow.CountryCode, pRow.Description)
}

type PricelistData struct {
	Data        []PricelistRow
	PricelistId string
}

func (pData PricelistData) String() string {
	var finalOutput []string
	for index, value := range pData.Data {
		finalOutput = append(finalOutput, fmt.Sprintf("Index: %d %s", index, value))
	}
	return strings.Join(finalOutput, "\n")
}

func (p PricelistData) getKey(prefix string, countryCode string) string {
	s := []string{"pricelists:", p.PricelistId, "/countries:", countryCode, "/", prefix}
	return strings.Join(s, "")
}

func (p PricelistData) getKeyPricelistsCountryList() string {
	s := []string{"pricelists:", p.PricelistId, "/countries"}
	return strings.Join(s, "")
}

func (p PricelistData) InsertToDb(c *redis.Client) {
	sadd := func(client *redis.Client, key string, value string) *redis.IntCmd {
		// sadd key value
		// smembers key
		var args []interface{}
		args = append(args, "sadd")
		args = append(args, key)
		args = append(args, value)

		cmd := redis.NewIntCmd(args...)
		c.Process(cmd)
		return cmd
	}

	hmset := func(client *redis.Client, key string, pricelistRow PricelistRow) *redis.StatusCmd {
		// hmset key keyVal1 val keyVal2 val2
		// hgetall key
		// hget key keyVal1
		var args []interface{}

		args = append(args, "hmset")
		args = append(args, key)
		args = append(args, "Prefix")
		args = append(args, pricelistRow.Prefix)
		args = append(args, "Rate")
		args = append(args, pricelistRow.Rate)
		args = append(args, "Description")
		args = append(args, pricelistRow.Description)

		for index, value := range pricelistRow.Optional {
			args = append(args, index)
			args = append(args, value)
		}

		cmd := redis.NewStatusCmd(args...)
		client.Process(cmd)
		return cmd
	}

	_, err := sadd(c, "pricelists", p.PricelistId).Result()
	if err != nil {
		log.Println("Error on add pricelistId to pricelists list")
	}

	for _, row := range p.Data {
		key := p.getKey(row.Prefix, row.CountryCode)
		_, err := hmset(c, key, row).Result()
		if err != nil {
			log.Println("Error on hmset")
		}

		key2 := p.getKeyPricelistsCountryList()
		_, err = sadd(c, key2, row.CountryCode).Result()
		if err != nil {
			log.Println("Error on add country to pricelists country list")
		}
	}
}

func LookUp(client *redis.Client, pricelistId string, countryCode string, prefix string, searchLength int) (error) {
	hGetAll := func(client *redis.Client, key string) *redis.StringStringMapCmd {
		cmd := redis.NewStringStringMapCmd("hgetall", key)
		client.Process(cmd)
		return cmd
	}

	pData := PricelistData{
		PricelistId: pricelistId,
	}

	actualPrefix := ""
	for searchLength > 0 {
		runes := []rune(prefix)
		actualPrefix = string(runes[0:searchLength])

		key := pData.getKey(actualPrefix, countryCode)

		startTime := time.Now()
		v, err := hGetAll(client, key).Result()
		delta := time.Since(startTime)

		// 1 sec = 1 000 000 microsec
		log.Printf("Actual Prefix: %s; length: %d; QueryTime: %s\n", actualPrefix, searchLength, delta)

		log.Printf("%f query/sec\n", 1/delta.Seconds())
		log.Printf("Key: %s\n", key)

		if err != nil {
			return errors.New(fmt.Sprintf("Error: %+v\n", err))
		}
		if len(v) > 0 {
			log.Printf("Result: %+v\n", v)
			break
		} else {
			log.Printf("Not found\n\n")
		}

		searchLength--
	}

	return nil
}
