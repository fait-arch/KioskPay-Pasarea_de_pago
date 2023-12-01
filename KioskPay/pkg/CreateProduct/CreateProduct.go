/* CreateProduct.go */
package CreateProduct

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Definimos la estructura JSON de Collections "products"
type Product struct {
	ProductImg        string `json:"product_Img"`
	ProductName       string `json:"product_Name"`
	ProductPrice      string `json:"product_Price"`
	ProductDescripion string `json:"product_descripion"`
}

func CreateProduct(product *Product) (*http.Response, error) {
	url := "http://localhost:8090/api/collections/Poducts/records"

	jsonData, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
