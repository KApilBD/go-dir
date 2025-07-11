package datatypes

type Rate struct {
	Currency string
	Price    float64 //json parser works with 64
}
