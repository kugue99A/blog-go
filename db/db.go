package db

import "gorm.io/gorm"

type DB interface {
}

type dbClient struct {
  cli *gorm.DB
}

func NewDB(user string, password string, host string, port string, dbName: string) DB {
  var dns = user + ":" + password + "@tcp(" + host + ":" + post + ")/" + dbName + "?charset=utf8mb4&parseTime=True"
  return nil
}
