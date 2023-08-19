package cmd_test

import (
	"bytes"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rzeAkbari/GoGo/gin/server"
	"github.com/rzeAkbari/GoGo/gin/store"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestService(t *testing.T) {
	testCases := []struct {
		name      string
		username  string
		status    int
		hasCookie bool
	}{
		{
			name:      "valid credentials",
			username:  "rze",
			status:    http.StatusOK,
			hasCookie: true,
		},
		{
			name:      "invalid credentials",
			username:  "invalid",
			status:    http.StatusForbidden,
			hasCookie: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("JWT_SIGN_KEY", "AllYourBase")

			dir, _ := os.Getwd()
			store, err := store.NewStore(dir + "/../data/test/gin.db")
			if err != nil {
				t.Fatalf("error connecting to db: %v", err)
			}
			defer store.DB.Close()

			s := server.NewServer(store)

			ginEngine := s.SetupRouter()
			body := map[string]string{
				"userName": tc.username,
				"password": "password",
			}
			bb, _ := json.Marshal(body)
			request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(bb))
			response := httptest.NewRecorder()

			ginEngine.ServeHTTP(response, request)

			hasCookie := response.Header().Get("Set-Cookie") != ""

			assert.Equal(t, tc.status, response.Code)
			assert.Equal(t, hasCookie, tc.hasCookie)
		})
	}

}
