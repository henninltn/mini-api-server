package db

import "gopkg.in/mgo.v2/bson"

type WordMapper struct {
	Word
}

func (wm *WordMapper) Insert() error {
	if err := wm.Word.isValid(); err != nil {
		return err
	}
	c, session := connect()
	defer session.Close()
	err := c.Insert(&wm.Word)
	return err
}

func (wm *WordMapper) Update(w *Word) error {
	if err := wm.Word.isValid(); err != nil {
		return err
	}
	c, session := connect()
	defer session.Close()
	err := c.Update(bson.M{"title": wm.Word.Title}, w)
	return err
}

func (_ *WordMapper) Find(title string) (Word, error) {
	c, session := connect()
	defer session.Close()
	res := Word{}
	err := c.Find(bson.M{"title": title}).One(&res)
	return res, err
}

func (_ *WordMapper) FindAll() ([]Word, error) {
	c, session := connect()
	defer session.Close()
	res := []Word{}
	err := c.Find(nil).All(&res)
	return res, err
}
