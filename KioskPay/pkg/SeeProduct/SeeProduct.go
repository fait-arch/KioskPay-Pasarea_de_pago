/*SeeProduct*/
package SeeProduct

import (
	"encoding/json"
	"log"
	"net/http"
)

// De API a JSON de productos
func SeeProduct() (map[string]interface{}, error) {
	//Llamado a la API de lista de productos
	resp, err := http.Get("http://localhost:8090/api/collections/Products/records")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer resp.Body.Close()

	// Decodifica el cuerpo de la respuesta como JSON
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln("Error al decodificar JSON:", err)
		return nil, err
	}

	return result, nil
}
