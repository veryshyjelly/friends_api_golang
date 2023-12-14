package database

import (
	"devjudge/go-in-docker/modals"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Connect() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("connections.db"), &gorm.Config{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&modals.User{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&modals.FriendRequest{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&modals.Friend{})
	return db, err
}

func CreateUserDB(db *gorm.DB, username string) error {
	var user modals.User
	db.Where("username = ?", username).First(&user)
	if user.Id != 0 {
		return fmt.Errorf("user already present")
	}
	log.Println("creating user: ", username)
	user = modals.User{Username: username}
	db.Create(&user)
	return nil
}

func SendFriendRequestDB(db *gorm.DB, fromUser, toUser string) error {
	if fromUser == toUser {
		return fmt.Errorf("cannot send friend request to self")
	}
	var fromUserModal, toUserModal modals.User

	db.Where("username = ?", fromUser).First(&fromUserModal)
	db.Where("username = ?", toUser).First(&toUserModal)

	if fromUserModal.Id == 0 {
		return fmt.Errorf("cannot find from user")
	}
	if toUserModal.Id == 0 {
		return fmt.Errorf("cannot find to user")
	}

	var friendRequest modals.FriendRequest
	db.Where("from_user = ? AND to_user = ?", fromUserModal.Id, toUserModal.Id).First(&friendRequest)

	if friendRequest.Id != 0 {
		return fmt.Errorf("friend request already present")
	}

	var reverseRequest modals.FriendRequest
	db.Where("from_user = ? AND to_user = ?", toUserModal.Id, fromUserModal.Id).First(&reverseRequest)

	if reverseRequest.Id != 0 {
		var toFriend, fromFriend = modals.Friend{MainUser: fromUserModal.Id, FriendUser: toUserModal.Id}, modals.Friend{MainUser: toUserModal.Id, FriendUser: fromUserModal.Id}
		db.Create(&toFriend)
		db.Create(&fromFriend)
	}

	friendRequest = modals.FriendRequest{FromUser: fromUserModal.Id, ToUser: toUserModal.Id}
	db.Create(&friendRequest)
	return nil
}

func GetAllFriendsDB(db *gorm.DB, username string) ([]string, error) {
	var friends = make([]string, 0)
	var user modals.User
	db.Model(&modals.User{}).Preload("Friends").Where("username = ?", username).First(&user)
	if user.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}
	for _, f := range user.Friends {
		var friend = modals.User{Id: f.FriendUser}
		db.First(&friend)
		friends = append(friends, friend.Username)
	}
	return friends, nil
}
