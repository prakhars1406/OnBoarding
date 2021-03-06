package factory

import (
	"github.com/prakhars1406/OnBoarding/utility"
	"github.com/prakhars1406/OnBoarding/config"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	dataStoreClient MongoClient          = nil
)

func MongoConnector() MongoClient {
	if dataStoreClient == nil {
		dataStore := utility.GetDataStore()
		if dataStore == config.MONGO {
			mongoServer := utility.GetMongoServer()
			dialInfo := mgo.DialInfo{
				Addrs:     []string{mongoServer},
				Direct:    false,
				Timeout:   1 * time.Second,
				FailFast:  false,
				PoolLimit: 100,
			}
			session, err := mgo.DialWithInfo(&dialInfo)
			if err != nil {
				return nil
			}
			dataStoreClient = &MongoClientImpl{mongoServer: mongoServer, session: session}
		}

	}
	return dataStoreClient
}
