package main

import (
	"fmt"

	"example.com/tax-calculator/filemanager"
	"example.com/tax-calculator/prices"
)

func main() {
	taxRates := []float64{0.05, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	for index, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)

		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		go priceJob.Process(doneChans[index], errorChans[index])
		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done")
		}

	}
}
