package secretnote

import (
	"fmt"
	"net/http"
	"time"

	"github.com/codescalersinternships/secret-note-api-spa-nabil/api/util"
	db "github.com/codescalersinternships/secret-note-api-spa-nabil/internal/db/models"
	"github.com/gin-gonic/gin"
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
	hashedPass, err := util.HashPassword(req.Password_Hashed)
	if err != nil {
		fmt.Fprint(ctx.Writer, "cann't hash password", err)
		return
	}
	arg := db.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPass,
	}
	user := server.store.Create(&arg)
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
	arg := db.User{
		Email: req.Email,
	}
	server.store.First(&arg)
	if util.CheckPassword(req.Password_Hashed, arg.Password) != nil {
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
