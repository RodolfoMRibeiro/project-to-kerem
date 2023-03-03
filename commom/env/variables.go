package env

import "os"

var (
	SECRET_KEY string = os.Getenv("SECRET_KEY")

	MONGO_URI             string = os.Getenv("MONGO_URI")
	MONGO_COLLECTION_NAME string = os.Getenv("MONGO_COLLECTION_NAME")
)
