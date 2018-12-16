package mango

import (
	"healthsystem/oneforall/common/errors"
	"os"
	"testing"

	"healthsystem/oneforall/common/vault"

	"github.com/stretchr/testify/assert"
)

func TestMain(M *testing.M) {
	os.Exit(M.Run())
}

func TestMango__Should_connect_successfully(t *testing.T) {
	mongoCred, mgoErr := vault.ServiceCredential("mongodb")
	assert.NoError(t, mgoErr)
	_, err := Connect("localhost", "admin", mongoCred.ID, mongoCred.Password)
	assert.NoError(t, err)
}

func TestMango__Should_fail_if_dbName_is_wrong(t *testing.T) {
	mongoCred, mgoErr := vault.ServiceCredential("mongodb")
	assert.NoError(t, mgoErr)
	_, err := Connect("localhost", "test", mongoCred.ID, mongoCred.Password)
	errMessage := errors.ErrorLog(errors.ErrCodeMongoDB, "server returned error on SASL authentication step: Authentication failed.").Error()
	assert.EqualError(t, err, errMessage)
}
