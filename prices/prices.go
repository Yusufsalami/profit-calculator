package prices

import (
	
	"fmt"

	"example.com/price/conversion"
	"example.com/price/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager filemanager.FileManger `json:"-"`
	TaxRate           float64 `json:"tax_rate"`
	InputPrices       []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func(job *TaxIncludedPriceJob) LoadData(){
	lines, err := job.IOManager.ReadFile()
	if err != nil{
		fmt.Println(err)
		return
	}



	prices, err := conversion.StringToFloats(lines)
		if err!= nil {
			fmt.Println(err)
			return
		}
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {

	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrices := price * (1+ job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrices)
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult( job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManger, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}

}