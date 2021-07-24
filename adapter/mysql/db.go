package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gami/layered_arch_example/config"
	"github.com/pkg/errors"

	// load mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *DB
)

const (
	// ConnMaxLifetimeSec is maximum amount of time a connection may be reused. The unit is second.
	// Expired connections may be closed lazily before reuse.
	ConnMaxLifetimeSec = 120

	// MaxOpenConns is the maximum number of open connections to the database.
	MaxOpenConns = 100

	// MaxIdleConns the maximum number of connections in the idle connection pool.
	// Here, MaxIdleConns is same as MaxOpenConns for perfomance. Idle conns will be killed when over ConnMaxLifetimeSec.
	MaxIdleConns = 100
)

const (
	//DefaultDBName is a real name of default DB
	DefaultDBName = "user"
)

// DB is a wrapped DB instance for User databases.
type DB struct {
	*sql.DB
	Name string
}

func init() {
	err := connectUserDB()
	if err != nil {
		log.Println(err)
		return
	}
}

func connect(connStr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	setupDBConns(db)
	return db, nil
}

func setupDBConns(db *sql.DB) {
	db.SetConnMaxLifetime(ConnMaxLifetimeSec * time.Second)
	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)
}

func connectString(database config.Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=Asia%%2FTokyo&charset=utf8&parseTime=true&time_zone=%%22Asia%%2FTokyo%%22",
		database.User,
		database.Password,
		database.Host,
		database.Port,
		database.Name,
	)
}

func connectUserDB() error {
	c, err := connect(connectString(config.GetConfig().Database))
	if err != nil {
		return errors.Wrap(err, "db connect failed")
	}
	db = &DB{
		DB:   c,
		Name: DefaultDBName,
	}
	return nil
}

// UserDB is a function to return user db connection.
func UserDB() *DB {
	return db
}
