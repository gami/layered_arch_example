package di

import "github.com/gami/layered_arch_example/adapter/mysql"

func InjectUserDB() *mysql.DB {
	return mysql.UserDB()
}
