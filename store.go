package model

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Image struct {
	Name string `bson:"name,omitempty" json:"name"`
}

type Store struct {
	Session    *mgo.Session
	dial       string
	dbName     string
	collection string
}

func (s *Store) Close() {
	s.Session.Close()
}

func (s *Store) Insert(model *Image) error {
	err := s.Session.DB(s.dbName).C(s.collection).Insert(model)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) All() ([]Image, error) {

	model := []Image{}
	err := s.Session.DB(s.dbName).C(s.collection).Find(bson.M{}).All(&model)

	if err != nil {
		return nil, err
	}
	return model, nil
}

func NewStore(dial string, dbName string, collection string) (*Store, error) {
	store := &Store{}
	store.Session, _ = mgo.Dial(dial)
	store.dial = dial
	store.dbName = dbName
	store.collection = collection

	return store, nil
}
