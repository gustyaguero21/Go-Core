package svc

import (
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	message := "Some message"
	err := "Some error"

	// Crear un error usando la función NewError
	e := NewError(message, err)

	// Verificar que el tipo de error sea del tipo esperado (error)
	if _, ok := e.(*Error); !ok {
		t.Fatalf("Expected error of type *Error, got %T", e)
	}

	// Verificar que el mensaje sea el esperado
	expectedMessage := fmt.Sprintf("%s, %s", message, err)
	if e.Error() != expectedMessage {
		t.Errorf("Expected error message: %s, but got: %s", expectedMessage, e.Error())
	}
}

func TestErrorMethod(t *testing.T) {
	message := "Test message"
	err := "Test error"
	errorStruct := &Error{
		Message: message,
		Err:     err,
	}

	// Verificar que el método Error devuelva el formato correcto
	expectedError := fmt.Sprintf("%s, %s", message, err)
	if errorStruct.Error() != expectedError {
		t.Errorf("Expected error: %s, but got: %s", expectedError, errorStruct.Error())
	}
}
