package sqlc

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-aws-micro/util"
	"log"
	"os"
	"testing"
)

/*const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:admin@localhost:5432/simple_bank?sslmode=disable"
)
*/
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can not load config:", err)
	}
	testDB, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
