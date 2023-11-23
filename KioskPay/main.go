package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// FUNCIONES ASINCRONICAS
//
//	Estructura de los datos del formulario
type DatosFormulario struct {
	NombreCompleto string `json:"fullName"`
	Direccion      string `json:"address"`
	Ciudad         string `json:"city"`
	Region         string `json:"state"`
	CodigoPostal   string `json:"zipCode"`
	Pais           string `json:"country"`
	Email          string `json:"email"`
	Celular        string `json:"phone"`
}

// Leer ruta del archivo de texto
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

// CONEXION CON EL INDEX
func handler(w http.ResponseWriter, r *http.Request) {

	//Accion de ingreso de datos del formulario
	if r.Method == "POST" {
		// Manejar la solicitud POST
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al analizar los datos del formulario", http.StatusInternalServerError)
			return
		}

		// Crear una instancia de DatosFormulario y asignar los valores
		datos := DatosFormulario{
			NombreCompleto: r.FormValue("fullName"),
			Direccion:      r.FormValue("address"),
			Ciudad:         r.FormValue("city"),
			Region:         r.FormValue("state"),
			CodigoPostal:   r.FormValue("zipCode"),
			Pais:           r.FormValue("country"),
			Email:          r.FormValue("email"),
			Celular:        r.FormValue("phone"),
		}
		// Conformar uso corecto de la estructura a JSON
		respuestaJSON, err := json.Marshal(datos)
		if err != nil {
			http.Error(w, "Error al convertir los datos a JSON", http.StatusInternalServerError)
			return
		}
		// Imprir en la terminal el JSON
		fmt.Println("Datos del formulario:")
		fmt.Printf("%+v\n", datos)

		// Responder al cliente con la respuesta JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respuestaJSON)
		return
	}

	// Manejar la solicitud GET (cargar el formulario)
	htmlContent, err := readFile("templates/index.html")
	if err != nil {
		http.Error(w, "Error al leer el archivo HTML :,( )", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.New("index").Parse(htmlContent)
	if err != nil {
		http.Error(w, "Error al analizar el archivo HTML :( )", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// MAIN
func main() {
	http.HandleFunc("/", handler)
	// Manejador para los archivos estáticos (CSS,HTML y JavaScript)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/templates/img/", http.StripPrefix("/templates/img/", http.FileServer(http.Dir("templates/img"))))

	//Abrir puerto local 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
