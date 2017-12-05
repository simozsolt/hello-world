package redisPoc

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
)

func GetPricelistRowsFromCsv(filePath string, pricelistId string) (PricelistData, error) {
	pData := PricelistData{
		PricelistId: pricelistId,
	}

	f, e := os.Open(filePath)
	if e != nil {
		return pData, e
	}
	defer f.Close()

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ';'

	indexToName := make(map[int]string)

	first := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return pData, err
		}

		// todo check required fields: prefix, rate ...
		// first line, build indexToName map
		if first == true {
			first = false
			for i, n := range record {
				indexToName[i] = n
			}

			continue
		}

		// convert coma to point in float value, after that convert to float
		price, _ := strconv.ParseFloat(
			strings.Replace(record[11], ",", ".", -1), //todo const name to value
			64,
		)

		// create optional map from row
		rawRowData := make(map[string]interface{})
		for i, v := range record {
			rawRowData[indexToName[i]] = v
		}

		pData.Data = append(
			pData.Data,
			PricelistRow{
				Prefix:      record[6], //todo const name to value
				Rate:        price,
				Description: record[0], //todo const name to value
				CountryCode: record[3], //todo const name to value
				Optional:    rawRowData,
			})
	}

	return pData, nil
}
