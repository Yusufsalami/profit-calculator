package main

import (
	"fmt"

	"example.com/price/filemanager"
	"example.com/price/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}



	for _, taxRate := range taxRates{
		fm := filemanager.New("price.txt",fmt.Sprintf("result_%.0f.json", taxRate * 100))

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
	


}