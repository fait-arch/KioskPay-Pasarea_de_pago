/*Server con golang*/

package main

import (
	//CreateProduct "KioskPay/pkg"
	//CreateUser "KioskPay/pkg/CreateUser"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {

	//----------POCKET BASE---------
	//  LEVANTAMIENTO DE POCKET BASE
	//------------------------------
	// Ruta al ejecutable
	executablePath := ".\\db\\pocketbase.exe"
	// Argumentos del comando
	args := []string{"serve"}
	// Configurar el comando
	cmd := exec.Command(executablePath, args...)
	// Configurar la salida estándar y de errores para ver los resultados
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Ejecutar el comando en una goroutine
	go func() {
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error al ejecutar el comando: %v\n", err)
		}
	}()

	//--------------------GIN TONIC--------------------
	//  LEVANTAMIENTO DE UN SERVIDOR WEB SIMPLE CON GIN
	//-------------------------------------------------
	//
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// Rutas estáticas para servir archivos estáticos y publicos (HTML, CSS y JavaScript)
	router.Static("/static", "./public_web/static")
	router.StaticFS("/img", http.Dir("./public_web/templates/img"))
	router.LoadHTMLGlob("public_web/templates/*.html")
	// Ruta para la página principal
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

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
	defer client.Close() //Asegurar el cierre del cliente al final de la función

	/*
			//
			//  LEER DATOS DE FORMULARIO DE USERS
			//
			// Ruta POST para manejar la entrada del formulario
			router.POST("/", func(c *gin.Context) {
				CreateUser.CreateUser(c)
				// Redirige de vuelta a la página principal después de enviar
				c.Redirect(http.StatusFound, "/")
			})

			//---------------------------------
			//  ENVIO A USERS A FIREBASE
			//---------------------------------
			// Agregar un documento a la colección "Users" en Firestore
			_, _, err = client.Collection("Users").Add(context.Background(), map[string]interface{}{
				"user_Name": "texto_prueba_002",
			})
			if err != nil {
				log.Fatalf("error adding document: %v\n", err)
			}
			// Imprimir un mensaje indicando que el documento se ha agregado con éxito
			log.Println("Document added successfully.")



		   //---------------------------------
		   //  ENVIO A PRODUCTS A POCKET BASE
		   //---------------------------------
		   product := &CreateProduct.Product{
		       ProductImg:        "",
		       ProductName:       "",
		       ProductPrice:      "",
		       ProductDescripion: "",
		   }


		   resp, err := CreateProduct.CreateProduct(product)
		   if err != nil {
		       fmt.Println("Error:", err)
		       return
		   }
		   defer resp.Body.Close()
		   fmt.Println(product)
		   fmt.Println("Product created successfully")
	*/

	//
	//INCIAR :8080
	//
	// Configurar la confianza en los proxies
	router.ForwardedByClientIP = true
	// Habilitar el puerto 8080
	fmt.Println("🅢 🅔 🅡 🅥 🅔 🅡   🅡 🅤 🅝 🏃‍♂️🏃‍♂️")
	router.Run(":8080")

}
