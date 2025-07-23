package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http/httptest"
	"testing"
	"vk-server-task/internal/models"
	"vk-server-task/internal/service"
	service_mocks "vk-server-task/internal/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestHandler_register(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockAuth, user registerRequest)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            registerRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"login": "username", "password": "0123456"}`,
			inputUser: registerRequest{
				Login:    "username",
				Password: "0123456",
			},
			mockBehavior: func(r *service_mocks.MockAuth, user registerRequest) {
				r.EXPECT().CreateUser(context.Background(), user.Login, user.Password).Return(&models.User{Id: 1, Login: "username", Password: "0123456"}, "someToken", nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"token":"someToken","user":{"id":1,"login":"username"}}`,
		},
		{
			name:                 "invalid register information",
			inputBody:            `{"login": "username", "password": "1234"}`,
			inputUser:            registerRequest{},
			mockBehavior:         func(r *service_mocks.MockAuth, user registerRequest) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid request"}`,
		},
		{
			name:      "registration failed",
			inputBody: `{"login": "username", "password": "0123456"}`,
			inputUser: registerRequest{
				Login:    "username",
				Password: "0123456",
			},
			mockBehavior: func(r *service_mocks.MockAuth, user registerRequest) {
				r.EXPECT().CreateUser(context.Background(), user.Login, user.Password).Return(nil, "", errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockAuth(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Auth: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/register", handler.Register)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/register",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_login(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockAuth, user loginRequest)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            loginRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"login": "username", "password": "0123456"}`,
			inputUser: loginRequest{
				Login:    "username",
				Password: "0123456",
			},
			mockBehavior: func(r *service_mocks.MockAuth, user loginRequest) {
				r.EXPECT().LoginUser(context.Background(), user.Login, user.Password).Return(&models.User{Id: 1, Login: "username", Password: "0123456"}, "someToken", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"token":"someToken","user":{"id":1,"login":"username"}}`,
		},
		{
			name:                 "invalid login information",
			inputBody:            `{"login": "username", "password": 0123456}`,
			inputUser:            loginRequest{},
			mockBehavior:         func(r *service_mocks.MockAuth, user loginRequest) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid request"}`,
		},
		{
			name:      "login failed",
			inputBody: `{"login": "username", "password": "0123456"}`,
			inputUser: loginRequest{
				Login:    "username",
				Password: "0123456",
			},
			mockBehavior: func(r *service_mocks.MockAuth, user loginRequest) {
				r.EXPECT().LoginUser(context.Background(), user.Login, user.Password).Return(nil, "", errors.New("something went wrong"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockAuth(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Auth: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/login", handler.Login)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/login",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
