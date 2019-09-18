package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	util "product-scraping/lib_utilities"
)

type CasDb struct {
	Session    *gocql.Session
	DateFormat string
}

type Node struct {
	Ip   string
	Name string
}

type CasaConfig struct {
	Cluster     []Node
	Keyspace    string
	Consistency gocql.Consistency
	DateFormat  string
}

func LocalCasaConfig() *CasaConfig {
	var nodes []Node
	nodes = append(nodes, Node{Ip: util.GetHostIP()})
	return &CasaConfig{
		Cluster:     nodes,
		Keyspace:    "scraper",
		Consistency: gocql.Quorum,
		DateFormat:  "2006-01-02",
	}
}


func GetTimeUUID() gocql.UUID {
	return gocql.TimeUUID()
}

func prepareQuery(query string, args ...interface{}) string {
	fullquery := fmt.Sprintf(query, args...)
	return fullquery
}

func getHostString(config *CasaConfig) []string {
	var hosts []string
	for _, node := range config.Cluster {
		if node.Name != "" {
			hosts = append(hosts, node.Name)
		} else {
			hosts = append(hosts, node.Ip)
		}
	}
	return hosts
}

func InitializeCasaConn(config *CasaConfig) (*CasDb, error) {
	hosts := getHostString(config)
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = config.Keyspace
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	//defer session.Close()
	return &CasDb{Session: session, DateFormat: config.DateFormat}, nil
}