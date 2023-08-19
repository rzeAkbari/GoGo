package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Credentials struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Store interface {
	CheckAuth(credentials Credentials) (bool, error)
}

type Server struct {
	Store Store
}

func NewServer(store Store) *Server {
	return &Server{
		store,
	}
}

func (s *Server) handleLogin(c *gin.Context) {
	var signingKey = os.Getenv("JWT_SIGN_KEY")
	var credentials Credentials
	if c.BindJSON(&credentials) == nil {
		valid, _ := s.Store.CheckAuth(credentials)
		if valid {
			tenDays := time.Now().AddDate(0, 0, 10)
			claim := &jwt.RegisteredClaims{
				Issuer:    "rze",
				Audience:  jwt.ClaimStrings{"admin"},
				Subject:   credentials.UserName,
				ExpiresAt: &jwt.NumericDate{Time: tenDays},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

			jwt, err := token.SignedString([]byte(signingKey))
			if err != nil {
				c.String(http.StatusInternalServerError, "internal server error")
			}

			c.SetCookie("jwt", jwt, int(tenDays.UnixMilli()), "/", "localhost", true, false)
			c.String(http.StatusOK, jwt)
		}
		c.String(http.StatusForbidden, "unauthorised")
	}
}

func (s *Server) SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/login", s.handleLogin)

	return r
}
