package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"mf.com/go/crypto/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	// fmt.Println(res, err)
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		dataBytes, ioErr := io.ReadAll(res.Body)
		if ioErr != nil {
			return nil, ioErr
		}
		jerr := json.Unmarshal(dataBytes, &response)
		if jerr != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := datatypes.Rate{
		Currency: currency,
		Price:    response.Bid,
	}
	return &rate, nil
}
