package cassandra

import (
	"strings"

	"github.com/gocql/gocql"
	"github.com/laithrafid/bookstore_utils-go/config_utils"
	"github.com/laithrafid/bookstore_utils-go/logger_utils"
)

var (
	session *gocql.Session
)

func init() {
	config, confErr := config_utils.LoadConfig(".")
	if confErr != nil {
		logger_utils.Error("cannot load config of application:", confErr)
	}
	hosts := strings.Split(config.CassDBnodes, ",")
	// Connect to Cassandra cluster:
	cluster := gocql.NewCluster(config.CassDBSource)
	cluster.Keyspace = config.CassDBKeyspace
	cluster.Consistency = gocql.Quorum
	cluster.Hosts = hosts

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		logger_utils.Error("cannot create connect to cassandra", err)
	}
}

func GetSession() *gocql.Session {
	return session
}
