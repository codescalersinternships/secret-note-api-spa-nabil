package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"` // Standard field for the primary key
	Name      string    `gorm:"not null" json:"name"`                                     // A regular string field that can't be null
	Email     string    `gorm:"not null;unique" json:"email"`                             // A pointer to a string, allowing for null values
	Password  string    `gorm:"not null" json:"password"`                                 // A regular string field that can't be null
	CreatedAt time.Time `json:"createdat"`                                                // Automatically managed by GORM for creation time
	UpdatedAt time.Time `json:"updatedat"`                                                // Automatically managed by GORM for update time
}

type Note struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Text          string    `gorm:"not null" json:"text"`
	NoteRemVisits int32     `gorm:"default:1" json:"noteremvisits"`
	UserID        uuid.UUID `gorm:"type:uuid;not null" json:"userid"`
	User          User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ExpireAt      time.Time `json:"expiredat"`
	CreatedAt     time.Time `json:"createdat"` // Automatically managed by GORM for creation time
}

func (user *User) CreateUser(db Store) error{
	return db.CreateNewUser(user)
}

func (user *User) FindByEmail(email string, db Store) error {
	return db.GetUserByEmail(email, user)
}
func (user *User) FindAllUserNotes(db Store) ([]Note,error) {
	var notes = []Note{}
	err := db.GetNotesByUser(user,&notes)
	return notes, err
}
func (note *Note) CreateNote(db Store) error{
	return db.CreateNewNote(note)
}
func (note *Note) FindByID(id uuid.UUID, db Store) error {
	return db.GetNoteByID(id, note)
}
func (note *Note) Update(db Store) error{
	return db.UpdateNote(note)
}

func (note *Note) Delete(db Store) error{
	return db.DeleteNote(note)
}
type Store interface{
	
	CreateNewUser(user *User) error
	GetUserByEmail(email string, user *User) error
	CreateNewNote(note *Note) error
	GetNoteByID(id uuid.UUID, note *Note) error
	GetNotesByUser(user *User, notes *[]Note) error
	UpdateNote(note *Note) error
	DeleteNote(note *Note) error
}
type SqlStore struct{
	GormStore      *gorm.DB
}
func (db *SqlStore) CreateNewUser(user *User) error{
	query :=db.GormStore.Create(user)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
func (db *SqlStore) GetUserByEmail(email string, user *User) error{
	query := db.GormStore.Where("email = ?", email).First(&user)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (db *SqlStore) CreateNewNote(note *Note) error{
	query := db.GormStore.Create(note)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (db *SqlStore) GetNoteByID(id uuid.UUID, note *Note) error{
	query := db.GormStore.Where("id = ?", id).First(&note)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (db *SqlStore) GetNotesByUser(user *User, notes *[]Note) error{
	query := db.GormStore.Where("user_id = ?", user.ID).Find(&notes)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
func (db *SqlStore) UpdateNote(note *Note) error{
	query := db.GormStore.Save(&note)
	if query != nil {
		return query.Error
	}
	return nil
}
func (db *SqlStore) DeleteNote(note *Note) error{
	query := db.GormStore.Delete(&note)
	if query != nil {
		return query.Error
	}
	return nil
}
type MockStore struct {
	notes          []*Note
	users          []*User
}
func (mockStore *MockStore) CreateNewUser(user *User) error {
	for _, u := range mockStore.users {
		if u.Email == user.Email {
			return fmt.Errorf("user with email %s already exists", user.Email)
		}
	}
	var err error
	user.ID, err = uuid.NewUUID()
	if err != nil {
		return err
	}
	mockStore.users = append(mockStore.users, user)
	return nil
}


func (mockStore *MockStore) GetUserByEmail(email string, user *User) error {
	for _, curUser := range mockStore.users {
		if curUser.Email == email {
			*user = *curUser // Copy the value into user
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func (mockStore *MockStore) CreateNewNote(note *Note) error {
	var err error
	note.ID, err = uuid.NewUUID()
	if err != nil {
		return err
	}
	mockStore.notes = append(mockStore.notes, note)
	return nil
}

func (mockStore *MockStore) GetNoteByID(id uuid.UUID, note *Note) error {
	for _, n := range mockStore.notes {
		if n.ID == id {
			*note = *n 
			return nil
		}
	}
	return fmt.Errorf("note not found")
}

func (mockStore *MockStore) GetNotesByUser(user *User, notes *[]Note) error {
	var userNotes []Note
	for _, note := range mockStore.notes {
		if note.UserID == user.ID {
			userNotes = append(userNotes, *note)
		}
	}
	*notes = userNotes
	return nil
}
func (mockStore *MockStore) UpdateNote(note *Note) error{
	for i, n := range mockStore.notes {
		if n.ID == note.ID {
			mockStore.notes[i].Text = note.Text
			mockStore.notes[i].ExpireAt = note.ExpireAt
			mockStore.notes[i].NoteRemVisits = note.NoteRemVisits
			return nil
		}
	}
	return errors.New("note not found")
}

func (mockStore *MockStore) DeleteNote(note *Note) error {
	for i, n := range mockStore.notes {
		if n.ID == note.ID {
			mockStore.notes = append(mockStore.notes[:i], mockStore.notes[i+1:]...)
			return nil
		}
	}
	return errors.New("note not found")
}