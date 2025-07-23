package handler

import (
	"errors"
	"net/http/httptest"
	"testing"
	"vk-server-task/internal/service"
	service_mocks "vk-server-task/internal/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestHandler_middleware(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockAuth, token string)

	tests := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *service_mocks.MockAuth, token string) {
				r.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `1`,
		},
		{
			name:                 "empty header",
			headerName:           "",
			headerValue:          "Bearer token",
			token:                "token",
			mockBehavior:         func(r *service_mocks.MockAuth, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"empty auth header"}`,
		},
		{
			name:                 "invalid header",
			headerName:           "Authorization",
			headerValue:          "Bearr token",
			token:                "token",
			mockBehavior:         func(r *service_mocks.MockAuth, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "empty token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "token",
			mockBehavior:         func(r *service_mocks.MockAuth, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"token is empty"}`,
		},
		{
			name:        "parse error",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *service_mocks.MockAuth, token string) {
				r.EXPECT().ParseToken(token).Return(0, errors.New("parse token error"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"parse token error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockAuth(c)
			test.mockBehavior(repo, test.token)

			services := &service.Service{Auth: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.GET("/auth", handler.AuthMiddleware, func(c *gin.Context) {
				id, exists := c.Get("userId")
				if exists {
					c.String(200, "%d", id)
				}
			})

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/auth", nil)
			req.Header.Set(test.headerName, test.headerValue)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
