package secretnote

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/codescalersinternships/secret-note-api-spa-nabil/backend/internal/db/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ErrorResponse struct {
    Message string `json:"message"`
}


type createNoteRequest struct {
	UserID     string `json:"userid" binding:"required"`
	Text       string `json:"text" binding:"required"`
	MaxRemDays int32  `json:"noteremvisits" binding:"required"`
	ExpireDate string `json:"expiredat" binding:"required"`
}

// @Summary Create a new note
// @Description Create a new note with unique URL
// @Accept json
// @Produce json
// @Param body body createNoteRequest true "Note data"
// @Success 200 {object} db.Note
// @Failure 400 {object} ErrorResponse
// @Router /create [post]
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
		fmt.Fprint(ctx.Writer, "date isn't parsing")
		return
	}
	id, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(ctx.Writer, "uuid isn't parsing")
		return
	}
	note := db.Note{
		Text:          req.Text,
		NoteRemVisits: req.MaxRemDays,
		ExpireAt:      expireTime,
		UserID:        id,
	}
	err = note.CreateNote(server.store)
	if err != nil {
		fmt.Fprint(ctx.Writer, "can't create note")
		return
	}
	fmt.Fprint(ctx.Writer, note)
}

type getNoteRequest struct {
	ID string `uri:"id" binding:"required"`
}
// @Summary Retrieve a specific note
// @Description Get a note by its ID Using the link parsing
// @Accept json
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} db.Note
// @Failure 400 {object} ErrorResponse
// @Router /note/{id} [get]
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
		fmt.Fprint(ctx.Writer, "date isn't parsing")
		return
	}
	err = note.FindByID(id, server.store)
	if err != nil {
		fmt.Fprint(ctx.Writer, "note is expired or doesn't exist")
		return
	}
	note.NoteRemVisits -= 1
	note.Update(server.store)
	if note.NoteRemVisits <= 0 || note.ExpireAt.Before(time.Now()) {
		err = note.Delete(server.store)
		if err != nil {
			fmt.Fprint(ctx.Writer, "Can't delete note")
			return
		}
	}
	ctx.JSON(http.StatusOK, note)
}

type getAllNoteRequest struct {
	ID string `uri:"userid" binding:"required"`
}

// @Summary Retrieve all notes for a user
// @Description Get all notes for the specified user
// @Accept json
// @Produce json
// @Param userid path string true "User ID"
// @Success 200 {array} db.Note
// @Failure 400 {object} ErrorResponse
// @Router /notes/{userid} [get]
func (server *Server) GetAllNotes(ctx *gin.Context) {
	var req getAllNoteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(ctx.Writer, "id isn't parsing")
		return
	}
	user := db.User{
		ID: id,
	}
	notes,err := user.FindAllUserNotes(server.store)
	if err != nil {
		fmt.Fprint(ctx.Writer, "can't get user notes")
		return
	}
	ctx.JSON(http.StatusOK, notes)
}
