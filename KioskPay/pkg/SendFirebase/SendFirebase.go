/*Envio a Firebase*/

package SendFirebase

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// User struct to represent the user data
type User struct {
	FullName string `json:"FullName"`
	CI       string `json:"CI"`
	Address  string `json:"Address"`
	City     string `json:"City"`
	Country  string `json:"Country"`
	Email    string `json:"Email"`
	Phone    string `json:"Phone"`
}

func CreateUser(userData string) {
	//-----------------FIREBASE----------------
	//  LEVANTAMIENTO DE CONECCION CON FIREBASE
	//-----------------------------------------
	// Leer el contenido del archivo de credenciales de Firebase
	data, err := ioutil.ReadFile("KEY.json")
	if err != nil {
		log.Fatalf("error al leer la key file: %v\n", err)
	}

	// Crear una opción de configuración para la autenticación de Firebase
	opt := option.WithCredentialsJSON(data)

	// Inicializar una nueva aplicación Firebase
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error levantar app: %v\n", err)
	}

	// Obtener un cliente Firestore a partir de la aplicación Firebase
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}
	defer client.Close() // Asegurar el cierre del cliente al final de la función

	//---------------------------------
	//  ENVIO A USERS A FIREBASE
	//---------------------------------

	// Decodificar el JSON de entrada en una estructura User
	var user User
	err = json.Unmarshal([]byte(userData), &user)
	if err != nil {
		log.Fatalf("error decoding JSON data: %v\n", err)
	}

	// Agregar un documento a la colección "Users" en Firestore
	_, _, err = client.Collection("Users").Add(context.Background(), map[string]interface{}{
		"FullName": user.FullName,
		"CI":       user.CI,
		"Address":  user.Address,
		"City":     user.City,
		"Country":  user.Country,
		"Email":    user.Email,
		"Phone":    user.Phone,
	})
	if err != nil {
		log.Fatalf("error adding document: %v\n", err)
	}
	// Imprimir un mensaje indicando que el documento se ha agregado con éxito
	log.Println("Document added successfully.")
}
