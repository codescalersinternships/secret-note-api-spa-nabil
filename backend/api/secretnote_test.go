package secretnote

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/codescalersinternships/secret-note-api-spa-nabil/backend/internal/db/models"
	"github.com/stretchr/testify/assert"
)
func Test_createNote(t *testing.T) {
	mockStore := &db.MockStore{}
	server := NewServer(mockStore)
	user := db.User{
		Name: "tester",
		Email: "test@gmail.com",
		Password: "testerpass123",
	}
	server.store.CreateNewUser(&user)
	
	fmt.Println(user.ID.String())
	tests := []struct {
		name         string
		requestBody  string
		expectedCode int
	}{
		{
			name:         "Valid Request",
			requestBody:  `{"userid":"`+user.ID.String()+`","text":"This is a note","noteremvisits":5,"expiredat":"2024-12-31 23:59:59"}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Missing Text",
			requestBody:  `{"userid":"`+user.ID.String()+`","noteremvisits":5,"expiredat":"2024-12-31 23:59:59"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Missing Expire Date",
			requestBody:  `{"userid":"`+user.ID.String()+`","text":"This is a note","noteremvisits":5}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Missing Remaining Visits",
			requestBody:  `{"userid":"`+user.ID.String()+`","text":"This is a note","expiredat":"2024-12-31 23:59:59"}`,
			expectedCode: http.StatusBadRequest,
		},
	}
	for _,test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, _ := http.NewRequest("POST", "/create", bytes.NewBufferString(test.requestBody))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()

			server.router.ServeHTTP(response, request)
			fmt.Println(test.requestBody)
			assert.Equal(t, test.expectedCode, response.Code)
		})
	}
}