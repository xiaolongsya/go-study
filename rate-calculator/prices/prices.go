package prices

import (
	"fmt"

	"example.com/tax-calculator/conversion"
	"example.com/tax-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	TaxRate           float64             `json:"tax_rate"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()
	if err != nil {
		errorChan <- err
		return
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	doneChan <- true

}

func NewTaxIncludedPriceJob(taxRate float64, iom iomanager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}
