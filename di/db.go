package di

import "app/adapter/mysql"

func InjectUserDB() *mysql.DB {
	return mysql.UserDB()
}
