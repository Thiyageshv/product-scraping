package cassandra

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

type CasDb struct {
	Cursor     *Query
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

func TestCasaConfig() *CasaConfig {
	var nodes []Node
	nodes = append(nodes, Node{Ip: "", Name: "ob1"})
	return &CasaConfig{
		Cluster:     nodes,
		Keyspace:    "hello_world",
		Consistency: gocql.Quorum,
		DateFormat:  "2006-01-02",
	}
}

func (c *CasDb) GetCurrentTime() (time.Time, string, int) {
	time_seen := time.Now().UTC()
	date_seen := time_seen.Format(fmt.Sprintf("%s", c.DateFormat))
	hour_seen := time_seen.Hour()
	return time_seen, date_seen, hour_seen
}

func (c *CasDb) GetTimeUUID() gocql.UUID {
	return gocql.TimeUUID()
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
	cursor := SetQueryBook()
	return &CasDb{Cursor: cursor, Session: session, DateFormat: config.DateFormat}, nil
}

func main() {
	// connect to the cluster
	//cluster := gocql.NewCluster("10.33.172.128", "10.33.172.129", "10.33.172.130", "10.33.172.131")
	config := TestCasaConfig()
	//hosts := getHostString(config)
	//cluster := gocql.NewCluster(hosts...)
	//cluster.Keyspace = "hello_world"
	//cluster.Consistency = gocql.Quorum
	//session, _ := cluster.CreateSession()
	//defer session.Close()
	//get_container_details(session)
	casdb, _ := InitializeCasaConn(config)
	//err1 := casdb.Insert_event("test", "testtote2")
	//if err1 != nil {
	//        log.Println(err1)
	//}
	results, err := casdb.Get_tote_scanned("test", "testtote2")
	if err != nil {
		log.Println(err)
	}
	for _, result := range results {
		log.Println(result.Transaction_status)
	}
}