package mango

import (
	"github.com/globalsign/mgo"
)

type Database struct {
	*mgo.Database
	*mgo.Session
}

func NewDatabase(d *mgo.Database, s *mgo.Session) *Database {
	s.SetSafe(&mgo.Safe{})
	s.SetPoolLimit(20000)
	return &Database{
		Database: d,
		Session:  s,
	}
}

func (d *Database) DB(name string) *mgo.Database {
	return d.Session.DB(name)
}

func (d *Database) C(name string) *mgo.Collection {
	return d.Database.C(name)
}
