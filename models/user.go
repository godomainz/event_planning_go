package models

import (
	"errors"
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

func (user User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return errors.New("Invalid Credentials")
	}

	isValidPassword := utils.CheckPasswordHash(user.Password, retrievedPassword)
	if !isValidPassword {
		return errors.New("Invalid Credentials")
	}

	return nil
}
