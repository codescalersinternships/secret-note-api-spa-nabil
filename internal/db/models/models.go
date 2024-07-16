package db

import (
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

func (user *User) CreateUser(db *gorm.DB) {
	db.Create(user)
}

func (user *User) FindByEmail(email string, db *gorm.DB) *gorm.DB {
	return db.Where("email = ?", email).First(&user)
}
func (user *User) FindAllUserNotes(db *gorm.DB) []Note {
	notes := []Note{}
	db.Where("user_id = ?", user.ID).Find(&notes)
	return notes
}
func (note *Note) CreateNote(db *gorm.DB) {
	db.Create(note)
}
func (note *Note) FindByID(id uuid.UUID, db *gorm.DB) *gorm.DB {
	return db.Where("id = ?", id).First(&note)
}

func (note *Note) Update(db *gorm.DB) {
	db.Save(&note)
}

func (note *Note) Delete(db *gorm.DB) {
	db.Delete(&note)
}
