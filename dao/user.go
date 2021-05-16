package dao

import (
	"log"

	"github.com/SpicyChickenFLY/never-todo-backend/model"
	"github.com/lingdor/stackerror"
	"gorm.io/gorm"
)

// ==================== User ====================

// GetAllUsers get all Users
func GetAllUsers(tx *gorm.DB, users *model.Users) error {
	log.Println("GetAllUsers")
	result := tx.Where(&model.User{Deleted: false}).Find(users)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetUserByID get User by ID
func GetUserByID(tx *gorm.DB, User *model.User, UserID int) error {
	log.Printf("GetUserByID(UserID: %d)\n", UserID)
	result := tx.Where(&model.User{ID: UserID, Deleted: false}).First(&User)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}

// GetUserByNick get Users by Nick
func GetUserByNick(tx *gorm.DB, Users *model.Users, nick string) error {
	log.Printf("GetUserByNick(nick: %s)\n", nick)
	result := tx.Where(&model.User{Nick: nick, Deleted: false}).First(&Users)
	// defer result.Close()
	if err := result.Error; err != nil {
		return stackerror.New(err.Error())
	}
	return nil
}
