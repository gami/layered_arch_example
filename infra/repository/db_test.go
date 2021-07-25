package repository_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"app/adapter/mysql"

	"github.com/go-testfixtures/testfixtures/v3"
)

var (
	testDB *mysql.DB
)

func TestMain(m *testing.M) {
	setup()

	code := m.Run()
	if code == 0 {
		log.Println(code)
	}

	os.Exit(code)
}

func setup() {
	testDB = mysql.UserDB()
	loadFixtures(testDB.DB, "simple")
}

func loadFixtures(db *sql.DB, name string) {
	folder := fmt.Sprintf("testdata/fixture/%v", name)
	loader, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(folder),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := loader.Load(); err != nil {
		log.Fatal(err)
	}
}
