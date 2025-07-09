package data

import "fmt"

type distance float64
type distanceKm float64

func (miles distance) ToKm() distanceKm {
	return distanceKm(1.60934 * miles)
}

func (kms distanceKm) toMiles() distance {
	return distance(kms / 1.60934)
}

func Test() {
	d := distance(4.5)
	km := d.ToKm()
	miles := km.toMiles()
	fmt.Printf("Kms:: %.2f Miles:: %.2f ", km, miles)
}
