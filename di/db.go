package di

import "github.com/gami/layered_arch_example/mysql"

func InjectUserDB() *mysql.DB {
	return mysql.UserDB()
}
