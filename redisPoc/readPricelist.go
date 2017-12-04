package redisPoc

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//func GetPricelistRows() PricelistData {
//	pData := PricelistData{
//		Data: []PricelistRow{
//			{Prefix: "36", Rate: 2.0, Description: "Landline", Optional: map[string]interface{}{"elso": "ertek1", "masodik": "masodik2"}},
//			{Prefix: "361", Rate: 1.0, Description: "Budapest"},
//			{Prefix: "3620", Rate: 11.0, Description: "Telenor"},
//			{Prefix: "3630", Rate: 12.0, Description: "T-Mobile"},
//			{Prefix: "3670", Rate: 13.0, Description: "Vodafone"},
//		},
//		PricelistId: "telekom",
//		CountryCode: "hu",
//	}
//
//	return pData
//}

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
