package secretnote

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/codescalersinternships/secret-note-api-spa-nabil/backend/internal/db/models"
	"github.com/gin-gonic/gin"
)

type signUpUserRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password_Hashed string `json:"password" binding:"required"`
}

func (server Server) SignUpUser(ctx *gin.Context) {
	var req signUpUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	hashedPass, err := HashPassword(req.Password_Hashed)
	if err != nil {
		fmt.Fprint(ctx.Writer, "cann't hash password", err)
		return
	}
	user := db.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPass,
	}
	err = user.CreateUser(server.store)
	if err != nil {
		fmt.Fprint(ctx.Writer, "cann't create user", err)
		return
	}
	fmt.Fprint(ctx.Writer, user)
}

type signInUserRequest struct {
	Email           string `json:"email" binding:"required"`
	Password_Hashed string `json:"password" binding:"required"`
}

func (server *Server) SignInUser(ctx *gin.Context) {
	var req signInUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	arg := db.User{}
	err := arg.FindByEmail(req.Email, server.store)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(ctx.Writer, "email doesn't exist", err)
		return
	}
	if CheckPassword(req.Password_Hashed, arg.Password) != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(ctx.Writer, "password is wrong")
		return
	}
	accessToken, err := server.tokenMaker.CreateToken(arg.Email, 24*time.Second)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(ctx.Writer, "cann't make token", err)
		return
	}
	res := struct {
		AccessToken string `json:"acces_token"`
		ID          string `json:"id"`
	}{
		AccessToken: accessToken,
		ID:          arg.ID.String(),
	}

	ctx.JSON(http.StatusOK, res)
}
