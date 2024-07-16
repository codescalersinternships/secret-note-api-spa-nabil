package secretnote

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/codescalersinternships/secret-note-api-spa-nabil/internal/db/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type signUpUserRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password_Hashed string `json:"password" binding:"required"`
}

func (server *Server) SignUpUser(ctx *gin.Context) {
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
	user.CreateUser(server.store)
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
	result := arg.FindByEmail(req.Email, server.store)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Fprint(ctx.Writer, "email doesn't exist")
		return
	}
	if CheckPassword(req.Password_Hashed, arg.Password) != nil {
		fmt.Fprint(ctx.Writer, "password is wrong")
		return
	}
	accessToken, err := server.tokenMaker.CreateToken(arg.Email, 24*time.Hour)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(ctx.Writer, "cann't make token", err)
		return
	}
	res := struct {
		AccessToken string `json:"acces_token"`
	}{
		AccessToken: accessToken,
	}

	ctx.JSON(http.StatusOK, res)
}
