package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

type user struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age string `json:"age"`
}

func (u *user) getUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *user) updateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (u *user) deleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (u *user) createUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func getUsers(db *sql.DB, start, count int) ([]user, error) {
	return nil, errors.New("Not implemented")
}

