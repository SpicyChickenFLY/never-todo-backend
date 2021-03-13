package dao

import (
	"log"

	"github.com/SpicyChickenFLY/never-todo-backend/model"
	"github.com/lingdor/stackerror"
	"gorm.io/gorm"
)

// ==================== User ====================

// GetAllUsers is a func to get all Users
func GetAllUsers(tx *gorm.DB, users *model.Users) error {
	log.Println("GetAllUsers")
	result := tx.Where(&model.User{Deleted: false}).Find(users)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetUserByID is a func to get Users by ID
func GetUserByID(tx *gorm.DB, Users *model.Users, UserID int) error {
	log.Printf("GetUserByID(UserID: %d)\n", UserID)
	result := tx.Where(&model.User{ID: UserID, Deleted: false}).First(&Users)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetUserByNick is a func to get Users by Nick
func GetUserByNick(tx *gorm.DB, Users *model.Users, nick string) error {
	log.Printf("GetUserByNick(nick: %d)\n", nick)
	result := tx.Where(&model.User{Nick: nick, Deleted: false}).First(&Users)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}
