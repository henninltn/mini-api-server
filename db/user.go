package db

import "gopkg.in/mgo.v2/bson"

type UserMapper struct {
	User
}

func (um *UserMapper) Insert() error {
	if err := um.User.isValid(); err != nil {
		return err
	}
	c, session := connect()
	defer session.Close()
	err := c.Insert(&um.User)
	return err
}

func (um *UserMapper) Update(u *User) error {
	if err := um.User.isValid(); err != nil {
		return err
	}
	c, session := connect()
	defer session.Close()
	err := c.Update(bson.M{"id": um.User.Id}, u)
	return err
}

func (_ *UserMapper) Find(id string) (User, error) {
	c, session := connect()
	defer session.Close()
	res := User{}
	err := c.Find(bson.M{"id": id}).One(&res)
	return res, err
}

func (_ *UserMapper) FindAll() ([]User, error) {
	c, session := connect()
	defer session.Close()
	res := []User{}
	err := c.Find(nil).All(&res)
	return res, err
}
