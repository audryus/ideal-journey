package cassandra

import (
	"ideal-journey/clients/logger"
	"ideal-journey/config"

	"github.com/gocql/gocql"
)

func NewClient() *gocql.Session {
	var err error
	cluster := gocql.NewCluster(config.GetConfig().Cassandra.Host)
	cluster.Keyspace = config.GetConfig().Cassandra.Keyspace
	cluster.Consistency = gocql.Quorum
	var session *gocql.Session

	if session, err = cluster.CreateSession(); err != nil {
		logger.Error("[CASSANDRA] Off-line.", err)
		panic(err)
	}

	logger.Info("[CASSANDRA] On-line.")
	return session
}
