package mango

import (
	"fmt"
	"healthsystem/oneforall/common/errors"

	"github.com/globalsign/mgo"
)

func Connect(ip, name, username, password string) (*Database, error) {
	url := fmt.Sprintf("mongodb://%s:%s@%s:28018/%s", username, password, ip, name)
	dbSession, err := mgo.Dial(url)
	if err != nil {
		return nil, errors.ErrorLog(errors.ErrCodeMongoDB, err)
	}
	return NewDatabase(dbSession.DB(name), dbSession), err
}
