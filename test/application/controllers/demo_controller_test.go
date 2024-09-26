package controllers

import (
	"github.com/example/application-api/src/application/controllers"
	"github.com/example/application-api/src/domain/models"
	"github.com/example/application-api/src/domain/usecases/mock"
	"encoding/json"
	"errors"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"go.uber.org/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestAppController(t *testing.T) {
	t.Run("Should return new demo controller", func(t *testing.T) {
		result := controllers.NewDemoController()
		assert.NotNil(t, result)
		assert.NotNil(t, result.Routes())
	})
}

func TestGetAllDemo(t *testing.T) {
	ctrl := gomock.NewController(t)
	demoGetAll := mock.NewMockIDemoGetAll(ctrl)
	demoGetByID := mock.NewMockIDemoGetByID(ctrl)
	restController := controllers.DemoController{DemoGetAll: demoGetAll, DemoGetByID: demoGetByID}
	defer ctrl.Finish()

	t.Run("Should get all demo", func(t *testing.T) {
		expected := []models.Demo{
			{ID: 1, Name: "demo test api 1"},
			{ID: 2, Name: "demo test api 2"},
		}

		demoGetAll.EXPECT().Execute(gomock.Any()).Return(expected, nil)

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodGet,
			Url:    "/demos",
			Path:   "/demos",
		}, restController.GetAll)

		var result []models.Demo
		assert.NoError(t, resp.DecodeBody(&result))
		assert.NotNil(t, result)
		assert.Equal(t, expected, result)
	})

	t.Run("Should return error when occurred error in GetPathParam", func(t *testing.T) {
		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodGet,
			Url:    "/demos/abc",
			Path:   "/demos/{id}",
		}, restController.GetById)

		assert.Equal(t, resp.StatusCode(), http.StatusBadRequest)
	})

	t.Run("Should return error when occurred error in GetById", func(t *testing.T) {
		errMock := errors.New("mock error")

		demoGetByID.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(nil, errMock)

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodGet,
			Url:    "/demos/999",
			Path:   "/demos/{id}",
		}, restController.GetById)

		assert.Equal(t, resp.StatusCode(), http.StatusInternalServerError)

		body, err := io.ReadAll(resp.RawBody())
		assert.NoError(t, err)
		assert.NotNil(t, body)

		var result restserver.Error
		err = json.Unmarshal(body, &result)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, errMock.Error(), result.Error)
	})

	t.Run("Should get demo by id", func(t *testing.T) {
		expected := models.Demo{ID: 1, Name: "demo test api 1"}
		demoGetByID.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(&expected, nil)

		resp := restserver.NewRequestTest(&restserver.RequestTest{
			Method: http.MethodGet,
			Url:    "/demos/1",
			Path:   "/demos/{id}",
		}, restController.GetById)

		assert.Equal(t, resp.StatusCode(), http.StatusOK)

		body, err := io.ReadAll(resp.RawBody())
		assert.NoError(t, err)
		assert.NotNil(t, body)

		var result models.Demo
		err = json.Unmarshal(body, &result)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expected, result)
	})
}
