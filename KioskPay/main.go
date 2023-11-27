package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// CONECTAR AL INDEX
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Manejar la solicitud POST
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al analizar los datos del formulario", http.StatusInternalServerError)
			return
		}

		// Realiza acciones con los datos, por ejemplo, imprímelos en la terminal
		fmt.Println("Datos del formulario:")
		fmt.Println("Nombre Completo:", r.FormValue("fullName"))
		fmt.Println("Dirección:", r.FormValue("address"))
		fmt.Println("Ciudad:", r.FormValue("city"))
		fmt.Println("Reguion:", r.FormValue("state"))
		fmt.Println("Codigo Postal:", r.FormValue("zipCode"))
		fmt.Println("Pais:", r.FormValue("country"))
		fmt.Println("Email:", r.FormValue("email"))
		fmt.Println("Celular:", r.FormValue("phone"))

		// Responder al cliente (puedes enviar un JSON u otro tipo de respuesta si es necesario)
		w.WriteHeader(http.StatusOK)
		return
	}

	// Manejar la solicitud GET (cargar el formulario)
	htmlContent, err := readFile("templates/index.html")
	if err != nil {
		http.Error(w, "Error al leer el archivo HTML", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("index").Parse(htmlContent)
	if err != nil {
		http.Error(w, "Error al analizar el archivo HTML", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func readFile(filename string) (string, error) {

	// Obtiene la ruta absoluta del archivo
	filePath, err := filepath.Abs(filename)
	//Mensaje de error
	if err != nil {
		return "", err
	}

	// Lee el contenido del archivo
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func main() {
	http.HandleFunc("/", handler)
	// Manejador para los archivos estáticos (CSS,HTML y JavaScript)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/templates/img/", http.StripPrefix("/templates/img/", http.FileServer(http.Dir("templates/img"))))

	//Abrir puerto local 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
