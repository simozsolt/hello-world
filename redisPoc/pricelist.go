package redisPoc

import (
	"fmt"
	"github.com/go-redis/redis"
	"strings"
)

type PricelistRow struct {
	Prefix      string
	Rate        float32
	Description string

	Optional map[string]interface{}
}

func (pRow PricelistRow) String() string {
	return fmt.Sprintf("%s %f %s %+v", pRow.Prefix, pRow.Rate, pRow.Description, pRow.Optional)
}

type PricelistData struct {
	Data        []PricelistRow
	PricelistId string
	CountryCode string
}

func (pData PricelistData) String() string {
	var finalOutput []string
	for index, value := range pData.Data {
		finalOutput = append(finalOutput, fmt.Sprintf("Index: %d %s", index, value))
	}
	return strings.Join(finalOutput, "\n")
}

func (p PricelistData) getKey(prefix string) string {
	s := []string{"pricelists:", p.PricelistId, "/countries:", p.CountryCode, "/", prefix}
	return strings.Join(s, "")
}

func (p PricelistData) InsertToDb(c *redis.Client) {
	hmset := func(client *redis.Client, key string, pricelistRow PricelistRow) *redis.StatusCmd {
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

	for index, row := range p.Data {
		fmt.Println("Insert:", index)
		key := p.getKey(row.Prefix)
		_, err := hmset(c, key, row).Result()
		if err != nil {
			fmt.Println("Error on hmset")
		}
	}
}
