package main

import (
	CreateUser "KioskPay/pkg/CreateUser"
	SendFirebase "KioskPay/pkg/SendFirebase"

	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
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

	//
	//
	//

	//
	//  LEER DATOS DE FORMULARIO DE USERS
	//
	// Ruta POST para manejar la entrada del formulario
	router.POST("/", func(c *gin.Context) {

		userJSON := CreateUser.CreateUser(c)
		// Redirige de vuelta a la página principal después de enviar el formulario
		c.Redirect(http.StatusFound, "/")
		fmt.Println(userJSON)
		SendFirebase.CreateUser(userJSON)

	})
	/*
		// Llama a la función SeeProduct()
		resultSeeProduct, err := SeeProduct.SeeProduct()
		if err != nil {
			log.Println("Error al obtener el producto:", err)
		} else {
			log.Printf("Resultado JSON: %+v\n", resultSeeProduct)
		}
		print(resultSeeProduct)
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
