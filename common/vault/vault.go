package vault

import (
	"healthsystem/oneforall/common/errors"
)

type credential struct {
	ID       string
	Password string
}

// Credentials for thirdparty
var thirdPartyCredential = map[string]credential{
	"mongodb": {
		ID:       "oneforall-test-user",
		Password: "oneforall-test-password",
	},
}

func ServiceCredential(serviceName string) (*credential, error) {
	if len(serviceName) == 0 {
		return nil, errors.ErrorLog(errors.ErrCodeEmptyServiceName)
	}
	credValue, ok := thirdPartyCredential[serviceName]
	if !ok {
		return nil, errors.ErrorLog(errors.ErrCodeInvalideServiceName, serviceName)
	}
	return &credValue, nil
}
