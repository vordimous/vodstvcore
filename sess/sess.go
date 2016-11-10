package sess

import (
	"esvodsCore/util"

	"github.com/gin-gonic/contrib/sessions"
	_ "github.com/jinzhu/gorm/dialects/postgres" //import postgres
)

const (
	//sessHost ...
	sessHost = "localdocker"
	//sessUser ...
	sessUser = "esvodsapi"
	//sessPassword ...
	sessPassword = "esvodsapi"
	//sessName ...
	sessName = "esvods"
)

//Init ...
func Init() sessions.RedisStore {

	//connect to reddis
	s, err := sessions.NewRedisStore(10, "tcp", sessHost+":6379", "", []byte("secret"))
	util.CheckErr(err, "Cannot connect to RedisStore")

	//save connection
	return s

}
