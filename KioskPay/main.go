/*Server con golang*/

package main

import (
	//CreateProduct "KioskPay/pkg"
	CreateUser "KioskPay/pkg/CreateUser"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {

	//------------------------------
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
	/*
	   //-------------------
	   //  ENVIO A PRODUCTS
	   //-------------------
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

	//--------------------------------------------
	//  CREACION DE UN SERVIDOR WEB SIMPLE CON GIN
	//--------------------------------------------

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

	//------------------------------------
	//  LEER DATOS DE FORMULARIO DE USERS
	//------------------------------------
	// Ruta POST para manejar la entrada del formulario
	router.POST("/", func(c *gin.Context) {
		CreateUser.CreateUser(c)
		c.Redirect(http.StatusFound, "/") // Redirige de vuelta a la página principal después de enviar
	})

	// Configurar la confianza en los proxies
	router.ForwardedByClientIP = true

	//
	//INCIAR :8080
	//
	// Habilitar el puerto 8080
	fmt.Println("🅢 🅔 🅡 🅥 🅔 🅡   🅡 🅤 🅝 🏃‍♂️🏃‍♂️")
	router.Run(":8080")

}
