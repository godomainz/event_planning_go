package models

import (
	"event_planning_go/db"
	"event_planning_go/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES(?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	
	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	return err
}
