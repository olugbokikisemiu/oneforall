package flags

import "flag"

var (
	MongoURL = flag.String("mongourl", "", "URL for test mongodb")
	MongoDB  = flag.String("mongodb", "", "authentication database for mongodb")
)
