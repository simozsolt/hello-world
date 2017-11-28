package redisPoc

func GetPricelistRows() PricelistData {
	pData := PricelistData{
		Data: []PricelistRow{
			{Prefix: "36", Rate: 2.0, Description: "Landline", Optional: map[string]interface{}{"elso": "ertek1", "masodik": "masodik2"}},
			{Prefix: "361", Rate: 1.0, Description: "Budapest"},
			{Prefix: "3620", Rate: 11.0, Description: "Telenor"},
			{Prefix: "3630", Rate: 12.0, Description: "T-Mobile"},
			{Prefix: "3670", Rate: 13.0, Description: "Vodafone"},
		},
		PricelistId: "telekom",
		CountryCode: "hu",
	}

	return pData
}
