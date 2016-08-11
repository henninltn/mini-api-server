package db

import (
	"errors"
	"unicode/utf8"
)

type Word struct {
	Title    string `bson:"title"    json:"title"`
	Contents string `bson:"contents" json:"contents"`
}

type User struct {
	Id       string `bson:"id"       json:"id"`
	Password string `bson:"password" json:"password"`
}

func (w *Word) isValid() error {
	if w.Title == "" {
		return errors.New("InvalidMemberError at Word.Title")
	}
	return nil
}

func (u *User) isValid() error {
	if u.Id == "" {
		return errors.New("InvalidMemberError at User.Id")
	}
	if utf8.RuneCountInString(u.Password) < 8 {
		return errors.New("InvalidMemberError at User.Password")
	}
	return nil
}
