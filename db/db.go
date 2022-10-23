package db

import (
	"fmt"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

var Notes *base.Base
var Users *base.Base


// var ctx = context.Background()

func InitDB() {
  d, err := deta.New(deta.WithProjectKey(DetaProjectKey))
  if err != nil {
    fmt.Println("failed to init new Deta instance: ", err)
  }

  Notes, _ = base.New(d, "notes")
  Users, _ = base.New(d, "users")


  fmt.Println("connected to deta :))")
}