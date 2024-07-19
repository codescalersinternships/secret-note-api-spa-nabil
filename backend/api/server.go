package secretnote

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	store      *gorm.DB
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
func NewServer(store *gorm.DB) *Server {
	tokenMaker := NewJWTMaker(RandString(32))
	r := gin.Default()
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.router = r
	server.router.Use(CORSMiddleware())
	server.router.HandleMethodNotAllowed = true
	server.router.POST("/signin", server.SignInUser)
	server.router.POST("/signup", server.SignUpUser)
	server.router.GET("/note/:id", server.GetNote)
	server.router.GET("/:userid", server.GetAllNotes)
	server.router.POST("/create", server.CreateNote).Use(ProtectedHandler(server.tokenMaker))
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
