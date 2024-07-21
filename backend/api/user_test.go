package secretnote

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	db "github.com/codescalersinternships/secret-note-api-spa-nabil/backend/internal/db/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_SignUpUser(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  string
		expectedCode int
	}{
		{
			name:         "Valid Request",
			requestBody:  `{"name":"John Doe","email":"john.doe@example.com","password":"securepassword"}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Missing Name",
			requestBody:  `{"email":"john.doe@example.com","password":"securepassword"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Missing Email",
			requestBody:  `{"name":"John Doe","password":"securepassword"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Missing Password",
			requestBody:  `{"name":"John Doe","email":"john.doe@example.com"}`,
			expectedCode: http.StatusBadRequest,
		},
	}
	mockStore := &db.MockStore{}
	server := NewServer(mockStore)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, _ := http.NewRequest("POST", "/signup", bytes.NewBufferString(test.requestBody))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			server.router.ServeHTTP(response, request)
			assert.Equal(t, test.expectedCode, response.Code)
		})
	}
}


func TestSignInUser(t *testing.T) {

	mockStore := &db.MockStore{}
	server := NewServer(mockStore)
	hashedPassword, _ := HashPassword("securepassword")

	user := db.User{
		ID: uuid.New(),
		Name: "john",
		Email: "john.doe@example.com",
		Password: hashedPassword,
		CreatedAt: time.Now(),
	}
	server.store.CreateNewUser(&user)
	tests := []struct {
		name         string
		requestBody  string
		expectedCode int
	}{
		{
			name:         "Valid Sign In",
			requestBody:  `{"email":"john.doe@example.com","password":"securepassword"}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Invalid Email",
			requestBody:  `{"email":"invalid@example.com","password":"securepassword"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Incorrect Password",
			requestBody:  `{"email":"john.doe@example.com","password":"wrongpassword"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Missing Email",
			requestBody:  `{"password":"securepassword"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Missing Password",
			requestBody:  `{"email":"john.doe@example.com"}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, _ := http.NewRequest("POST", "/signin", bytes.NewBufferString(test.requestBody))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()

			server.router.ServeHTTP(response, request)

			assert.Equal(t, test.expectedCode, response.Code)
		})
	}
}