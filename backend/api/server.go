package secretnote

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/codescalersinternships/secret-note-api-spa-nabil/backend/docs"
	db "github.com/codescalersinternships/secret-note-api-spa-nabil/backend/internal/db/models"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/time/rate"
)


type Server struct {
	store      db.Store
	tokenMaker JWTMaker
	router     *gin.Engine
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func NewServer(store db.Store) *Server {
	tokenMaker := NewJWTMaker(RandString(32))
	r := gin.Default()
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.router = r
	server.router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.router.Use(CORSMiddleware())
	server.router.HandleMethodNotAllowed = true
	server.router.POST("/signin", RateLimiter(server.SignInUser, rateLimit, burst))
	server.router.POST("/signup", RateLimiter(server.SignUpUser, rateLimit, burst))
	server.router.GET("/note/:id", RateLimiter(server.GetNote, rateLimit, burst))
	server.router.GET("/:userid", RateLimiter(server.GetAllNotes, rateLimit, burst))
	server.router.POST("/create", RateLimiter(server.CreateNote, rateLimit, burst)).Use(ProtectedHandler(server.tokenMaker))
	return server
}

func (server *Server) Start(address string) error {
	err := server.router.Run(address)
	if err != nil {
		return err
	}
	return endless.ListenAndServe(":8090", server.router)
}
func ProtectedHandler(token JWTMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("authorization")

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		fields := strings.Fields(authorizationHeader)

		if len(fields) < 1 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		var accessToken string
		if len(fields) >= 2 {
			accessToken = fields[1]
		} else {
			accessToken = fields[0]
		}
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		ctx.Set("payload", payload)
		ctx.Next()
	}

}
const (
	rateLimit = 10
	burst     = 10
)

func RateLimiter(next func(c *gin.Context), r rate.Limit, b int) gin.HandlerFunc {
	limiter := rate.NewLimiter(r, b)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
		}
		next(c)
	}
}