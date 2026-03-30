package cmdmanager

import "fmt"

type CMDManager struct {
}

func (fm CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER(0 to finish)")

	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scanln(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}
func (fm CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
