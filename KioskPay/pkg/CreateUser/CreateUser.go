package CreateUser

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// User es una estructura que representa la información de un usuario.
type User struct {
	FullName string `form:"fullName"` // Nombre completo del usuario
	CI       string `form:"ci"`       // Número de cédula de identidad
	Address  string `form:"address"`  // Dirección del usuario
	City     string `form:"city"`     // Ciudad del usuario
	Country  string `form:"country"`  // País del usuario
	Email    string `form:"email"`    // Correo electrónico del usuario
	Phone    string `form:"phone"`    // Número de teléfono del usuario
}

// CreateUser es una función que maneja la creación de un nuevo usuario.
// CreateUser es una función que maneja la creación de un nuevo usuario y devuelve el JSON del usuario.
func CreateUser(c *gin.Context) string {
	var user User

	// Intenta vincular los datos de la solicitud HTTP al objeto User
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println(err)
		return ""
	}

	// Llama a la función saveUser para almacenar la información del usuario y obtener el JSON del usuario.
	userJSON := SaveUser(user)

	// Devuelve el JSON del usuario como cadena de texto
	return userJSON
}

// saveUser es una función que guarda la información del usuario y devuelve el JSON del usuario.
func SaveUser(user User) string {
	// Convierte la estructura User a formato JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Imprime el JSON del usuario en la terminal (puedes comentar o eliminar esta línea si no es necesario)

	// Devuelve el JSON del usuario como cadena de texto
	return string(userJSON)
}
