package vault

import (
	"healthsystem/oneforall/common/errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestVault__Should_fail_with_empty_key(t *testing.T) {
	_, err := ServiceCredential("")
	errMessage := errors.ErrorLog(errors.ErrCodeEmptyServiceName)
	assert.EqualError(t, err, errMessage.Error())
}

func TestVault__Should_pass_with_correct_key(t *testing.T) {
	_, err := ServiceCredential("mongodb")
	assert.NoError(t, err)
}

func TestVault__Should_fail_with_invalid_key(t *testing.T) {
	_, err := ServiceCredential("mongod")
	errMessage := errors.ErrorLog(errors.ErrCodeInvalideServiceName, "mongod")
	assert.EqualError(t, err, errMessage.Error())
}
