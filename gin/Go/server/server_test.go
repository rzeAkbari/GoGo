package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rzeAkbari/GoGo/gin/server"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	var signKey = "AllYourBase"
	os.Setenv("JWT_SIGN_KEY", signKey)
	os.Exit(m.Run())
}

type spy struct{}

func (s *spy) CheckAuth(credentials server.Credentials) (bool, error) {
	return true, nil
}

func TestServerHandlers(t *testing.T) {
	t.Run("handles success login user password", func(t *testing.T) {
		spyStore := spy{}
		s := server.NewServer(&spyStore)
		ginEngine := s.SetupRouter()
		body := map[string]string{
			"userName": "user",
			"password": "pass",
		}
		bb, _ := json.Marshal(body)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(bb))
		response := httptest.NewRecorder()

		ginEngine.ServeHTTP(response, request)

		token, _ := jwt.Parse(response.Body.String(), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SIGN_KEY")), nil
		})

		claims := token.Claims
		issuer, _ := claims.GetIssuer()
		subject, _ := claims.GetSubject()

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "rze", issuer)
		assert.Equal(t, "user", subject)
	})
}
