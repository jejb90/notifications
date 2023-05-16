package ctx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockGateway struct {
	UserID  string
	Message string
	Sent    bool
}

func (mg *MockGateway) Send(userID, message string) {
	mg.UserID = userID
	mg.Message = message
	mg.Sent = true
}

func TestHandler_Send(t *testing.T) {
	gateway := &MockGateway{}
	handler := NewHandler(gateway)

	// Caso de prueba 1: Envío exitoso sin límite de frecuencia
	err := handler.Send("status", "user1", "Mensaje de prueba")
	assert.NoError(t, err, "Se esperaba un envío exitoso sin límite de frecuencia")
	assert.Equal(t, "user1", gateway.UserID, "El ID de usuario no coincide")
	assert.Equal(t, "Mensaje de prueba", gateway.Message, "El mensaje no coincide")
	assert.True(t, gateway.Sent, "El mensaje no fue enviado")
}
