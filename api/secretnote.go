package secretnote

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/codescalersinternships/secret-note-api-spa-nabil/internal/db/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type createNoteRequest struct {
	UserID     string `json:"userid" binding:"required"`
	Text       string `json:"text" binding:"required"`
	MaxRemDays int32  `json:"noteremvisits" binding:"required"`
	ExpireDate string `json:"expiredat" binding:"required"`
}

func (server *Server) CreateNote(ctx *gin.Context) {
	var req createNoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02 15:04:05"
	expireTime, err := time.Parse(layout, req.ExpireDate)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		ctx.Writer.Write([]byte("date isn't parsing"))
		return
	}
	id, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		ctx.Writer.Write([]byte("uuid isn't parsing"))
		return
	}
	note := db.Note{
		Text:          req.Text,
		NoteRemVisits: req.MaxRemDays,
		ExpireAt:      expireTime,
		UserID:        id,
	}
	note.CreateNote(server.store)
	fmt.Fprint(ctx.Writer, note)
}

type getNoteRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) GetNote(ctx *gin.Context) {
	var req getNoteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	note := db.Note{}
	id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		ctx.Writer.Write([]byte("date isn't parsing"))
		return
	}
	result := note.FindByID(id, server.store)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Fprint(ctx.Writer, "note is expired or doesn't exist")
		return
	}
	note.NoteRemVisits -= 1
	note.Update(server.store)
	if note.NoteRemVisits <= 0 || note.ExpireAt.Before(time.Now()) {
		server.store.Delete(&note)
	}
	fmt.Fprint(ctx.Writer, note)
}

type getAllNoteRequest struct {
	ID string `uri:"userid" binding:"required"`
}

func (server *Server) GetAllNotes(ctx *gin.Context) {
	var req getAllNoteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		ctx.Writer.Write([]byte("id isn't parsing"))
		return
	}
	user := db.User{
		ID: id,
	}
	notes := user.FindAllUserNotes(server.store)
	for _, note := range notes {
		fmt.Fprint(ctx.Writer, note)
	}
}
