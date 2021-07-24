package repository_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/gami/layered_arch_example/adapter/mysql"
	"gopkg.in/testfixtures.v2"
)

var (
	db *mysql.DB
)

const (
	localTimezone      = "Asia/Tokyo"
	localTimeDiffHours = 9
)

func init() {
	// set localtime to JST for test-fixtures
	loc, err := time.LoadLocation(localTimezone)
	if err != nil {
		loc = time.FixedZone(localTimezone, localTimeDiffHours*60*60)
	}

	time.Local = loc
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()
	if code == 0 {
		tearDown()
	}

	os.Exit(code)
}

func setup() {
	rand.Seed(time.Now().UnixNano())

	db = mysql.UserDB()

	tearDown() // clear DB at first

	loadFixtures(db, "simple")
}

func tearDown() {
	truncateDB(db, []string{
		"users",
	})
}

func truncateDB(db *mysql.DB, tables []string) {
	for _, table := range tables {
		_, err := db.Query(fmt.Sprintf("DELETE FROM %v", table))
		if err != nil {
			panic(err)
		}

		_, err = db.Query(fmt.Sprintf("ALTER TABLE %v AUTO_INCREMENT = 1", table))
		if err != nil {
			panic(err)
		}
	}
}

func loadFixtures(db *mysql.DB, name string) {
	folder := fmt.Sprintf("testdata/fixture/%v", name)
	fixtures, err := testfixtures.NewFolder(db.DB, &testfixtures.MySQL{}, folder)

	if err != nil {
		log.Fatal(err)
	}

	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}
