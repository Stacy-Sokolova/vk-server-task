package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http/httptest"
	"testing"
	"time"
	"vk-server-task/internal/models"
	"vk-server-task/internal/service"
	service_mocks "vk-server-task/internal/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestHandler_create(t *testing.T) {
	type input struct {
		userId int
		ad     service.CreateRequest
	}

	type mockBehavior func(r *service_mocks.MockAds, input input)

	tests := []struct {
		name                 string
		inputBody            string
		input                input
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"title":"title","description":"description","image_url":"https://www.example.com/images/photo.jpg ","price":1000}`,
			input: input{
				userId: 1,
				ad: service.CreateRequest{
					Title:       "title",
					Description: "description",
					ImageURL:    "https://www.example.com/images/photo.jpg ",
					Price:       1000,
				},
			},
			mockBehavior: func(r *service_mocks.MockAds, input input) {
				r.EXPECT().Create(context.Background(), input.userId, input.ad).Return(
					&models.Ads{
						Id:          1,
						UserId:      1,
						Title:       "title",
						Description: "description",
						ImageURL:    "https://www.example.com/images/photo.jpg ",
						Price:       1000,
						CreatedAt: time.Date(
							2025, 07, 23, 15, 22, 15, 6004947, time.UTC)},
					nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"created ad":{"id":1,"title":"title","description":"description","image_url":"https://www.example.com/images/photo.jpg ","price":1000,"created_at":"2025-07-23T15:22:15.006004947Z"}}`,
		},
		{
			name:      "failed to create ad",
			inputBody: `{"title":"title","description":"description","image_url":"https://www.example.com/images/photo.jpg ","price":1000}`,
			input: input{
				userId: 1,
				ad: service.CreateRequest{
					Title:       "title",
					Description: "description",
					ImageURL:    "https://www.example.com/images/photo.jpg ",
					Price:       1000,
				},
			},
			mockBehavior: func(r *service_mocks.MockAds, input input) {
				r.EXPECT().Create(context.Background(), input.userId, input.ad).Return(nil, errors.New("some internal error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"failed to create ad"}`,
		},
		{
			name:                 "invalid request data",
			inputBody:            `{"title":"title","description":"description","image_url":"incorrect url","price":1000}`,
			input:                input{},
			mockBehavior:         func(r *service_mocks.MockAds, input input) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid request"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockAds(c)
			test.mockBehavior(repo, test.input)

			services := &service.Service{Ads: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/create", func(c *gin.Context) {
				c.Set("userId", 1)
			}, handler.Create)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/create",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
