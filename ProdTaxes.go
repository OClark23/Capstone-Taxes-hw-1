package checkout

const (
	wicExempt    = "Wic Eligible food"
	clothingExem = "Clothing"
	fur          = "fur"
)

type Item struct {
	Name  string
	Price float64
	Type  string
}

type TaxRateMap map[string]float64

var taxRates = TaxRateMap{
	"NJ": 0.066,   // 6.6%
	"PA": 0.06,    // 6%
	"DE": 0,       // no sales tax
}

func Checkout(state string, items []Item) (float64, error) {
	var total float64
	taxRate, ok := taxRates[state]
	if !ok {
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
