package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/DaffaKhayru/gofiber-starter-pack"
	"github.com/DaffaKhayru/gofiber-starter-pack/models"
	"github.com/DaffaKhayru/gofiber-starter-pack/util"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestSignup(t *testing.T) {
	// hashing password
	hashPassword, err := util.HashPassword("daffa123");

	// if error while hashing password
	if err != nil {
		fmt.Println("Error when hashing password")
	}

	user := models.User{
		Username: "daffakhayru",
		Email: "daffakhayru@gmail.com",
		Password: hashPassword,
	}

	body, _ := json.Marshal(user)


	req := httptest.NewRequest(http.MethodPost, "/api/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.App.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}