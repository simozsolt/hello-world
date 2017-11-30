package main

import (
	"github.com/simozsolt/hello-world/redisPoc"
	"fmt"
)



func main() {
	filePath := "/home/simo/Work/go/src/github.com/simozsolt/hello-world/gts_nat.csv"

	pData := redisPoc.GetPricelistRowsFromCsv(filePath)
	fmt.Printf("%s\n", pData)
}
