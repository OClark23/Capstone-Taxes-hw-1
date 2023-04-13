package checkout

const (
	njTaxRate    = 0.066   // 6.6%
	paTaxRate    = 0.06    // 6%
	wicExempt    = "Wic Eligible food"
	clothingExem = "Clothing"
	fur          = "fur"
)

type Item struct {
	Name string
	Price float64
	Type string
}

func Checkout(state string, items []Item) (float64, error) {
	var total float64
	var taxRate float64
	switch state {
	case "NJ":
		taxRate = njTaxRate
	case "PA":
		taxRate = paTaxRate
	case "DE":
		taxRate = 0
	default:
		return 0, fmt.Errorf("state %s not supported", state)
	}
	for _, item := range items {
		var taxable bool
		if item.Type == wicExempt {
			continue
		} else if item.Type == clothingExem {
			if !strings.Contains(strings.ToLower(item.Name), fur) {
				continue
			}
			taxable = true
		} else {
			taxable = true
		}
		if taxable {
			total += item.Price * (1 + taxRate)
		} else {
			total += item.Price
		}
	}
	return total, nil
}
