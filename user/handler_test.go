package user

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"study/db"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestGetHandler(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err)
	repo := NewRepository(database)
	service := NewService(repo)
	handler := NewHandler(service)
	app := fiber.New()
	app.Get("/users/:id", handler.Get)
	id, err := repo.Create(Model{Name: "test", Email: "test@gmail.com"})
	assert.Nil(t, err)
	req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", id), nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetHandler_ReturnsErrorFromService(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err)
	repo := NewRepository(database)
	service := &MockService{}
	service.On("Get", mock.AnythingOfType("uint")).Return(nil, errors.New("fena bir hata"))
	handler := NewHandler(service)
	app := fiber.New()
	app.Get("/users/:id", handler.Get)
	id, err := repo.Create(Model{Name: "test", Email: "test@gmail.com"})
	assert.Nil(t, err)
	req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", id), nil)
	resp, err := app.Test(req)
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	fmt.Print(string(body))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
