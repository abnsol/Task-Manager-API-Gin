package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

// AuthMiddlewareSuite defines the test suite for auth middleware
type AuthMiddlewareSuite struct {
	suite.Suite
	router         *gin.Engine
	originalSecret string
}

// SetupTest runs before each test in the suite
func (s *AuthMiddlewareSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.router = gin.New()
	s.originalSecret = os.Getenv("jwt_secret")
	os.Setenv("jwt_secret", "a-very-secure-and-not-at-all-hardcoded-secret")
}

// TearDownTest runs after each test in the suite
func (s *AuthMiddlewareSuite) TearDownTest() {
	os.Setenv("jwt_secret", s.originalSecret)
}

// TestAuthMiddlewareSuite runs the entire test suite
func TestAuthMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareSuite))
}

// createTestToken is a helper to generate a JWT for testing
func (s *AuthMiddlewareSuite) createTestToken(role string) string {
	secret := []byte(os.Getenv("jwt_secret"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": role,
		"exp":  time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString(secret)
	return tokenString
}

// TestAuthMiddleware tests the main authentication middleware
func (s *AuthMiddlewareSuite) TestAuthMiddleware() {
	s.router.GET("/test", AuthMiddleware(), func(c *gin.Context) {
		_, exists := c.Get("claims")
		s.True(exists)
		c.Status(http.StatusOK)
	})

	s.Run("valid jwt", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		token := s.createTestToken("user")
		req.Header.Set("Authorization", "Bearer "+token)
		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusOK, w.Code)
	})
}

// TestRoleMiddleware tests the role-based access control middleware
func (s *AuthMiddlewareSuite) TestRoleMiddleware() {
	successHandler := func(c *gin.Context) { c.Status(http.StatusOK) }
	s.router.GET("/admin", AuthMiddleware(), RoleMiddleware("admin"), successHandler)

	s.Run("access granted for correct role", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/admin", nil)
		token := s.createTestToken("admin")
		req.Header.Set("Authorization", "Bearer "+token)
		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusOK, w.Code)
	})

	s.Run("access forbidden for wrong role", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/admin", nil)
		token := s.createTestToken("user")
		req.Header.Set("Authorization", "Bearer "+token)
		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusForbidden, w.Code)
	})
}
