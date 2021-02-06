package basic

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "sync"
  "fmt"
  "github.com/micro/go-micro/util/log"
)

var (
  inited	bool
  db		*sql.DB
  m		sync.RWMutex
)

func Init() {
  m.Lock()
  defer m.Unlock()

  var err error

  if inited {
    err = fmt.Errorf("[Init] db has inited")
    log.Logf(err.Error())
    return
  }

  dsn := "micro-login:tF#262420228@tcp(mysql:3306)/user?charset=utf8"
  db, err = sql.Open("mysql", dsn)
  if err != nil {
     panic(err.Error())
  }

  inited = true
}

func GetDB() *sql.DB {
  if !inited {
    panic("DB does not init")
  }
  return db
}
