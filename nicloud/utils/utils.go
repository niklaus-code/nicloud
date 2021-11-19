package utils

import (
  uuid "github.com/satori/go.uuid"
)

func Createuuid() string {
  /*create uuid*/
  u := uuid.NewV4().String()
  return u
}

