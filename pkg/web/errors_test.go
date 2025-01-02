package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	// Configurar Gin en modo test
	gin.SetMode(gin.TestMode)

	// Crear un contexto de prueba
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	// Ejecutar la función bajo prueba
	status := http.StatusBadRequest
	errMessage := "Invalid input"
	NewError(ctx, status, errMessage)

	// Validar el código de estado
	assert.Equal(t, status, w.Code)

	// Validar el cuerpo de la respuesta
	expectedResponse := Error{
		Status: status,
		Err:    errMessage,
	}
	var actualResponse Error
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err) // Asegurarse de que el JSON se decodifica sin errores
	assert.Equal(t, expectedResponse, actualResponse)
}
