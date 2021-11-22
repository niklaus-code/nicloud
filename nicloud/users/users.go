package users

import (
  db "nicloud/dbs"
  "nicloud/vmerror"
)

type Vms_users struct {
  Id int
  Username string
  Passwd string
  Email string
  Admin int
}

func Login(username string, passwd string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  u := &Vms_users{}
  dbs.Where("username=?", username).First(u)
  if u.Passwd == passwd {
    return nil
  } else {
    return vmerror.Error{Message: "Authentication failed"}
  }

}
