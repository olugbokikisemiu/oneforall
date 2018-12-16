package main

import (
	"flag"
	"healthsystem/oneforall/common/flags"
	"healthsystem/oneforall/common/mango"
	"healthsystem/oneforall/common/vault"

	log "github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	mgoCred, err := vault.ServiceCredential("mongodb")
	if err != nil {
		log.Error(err)
	}

	if _, err := mango.Connect(
		*flags.MongoURL,
		*flags.MongoDB,
		mgoCred.ID,
		mgoCred.Password,
	); err != nil {
		log.Error(err)
	} else {
		log.Info("Conneted to DB")
	}

}
