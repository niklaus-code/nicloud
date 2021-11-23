package users

import (
  "github.com/dgrijalva/jwt-go"
  db "nicloud/dbs"
  "nicloud/vmerror"
  "strconv"
  "time"
)

type Vms_users struct {
  Id int
  Username string
  Passwd string
  Email string
  Admin int
}

func createtoken(username string, userid string) (string, error) {
  expiresTime := time.Now().Unix() + int64(10)

  claims := jwt.StandardClaims{
    Audience:  username,     // 受众
    ExpiresAt: expiresTime,       // 失效时间
    Id:        userid,   // 编号
    IssuedAt:  time.Now().Unix(), // 签发时间
    Issuer:    "gin hello",       // 签发人
    NotBefore: time.Now().Unix(), // 生效时间
    Subject:   "login",           // 主题
  }
  tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  var jwtSecret = []byte("nicloud")
  token, err := tokenClaims.SignedString(jwtSecret)
  if err != nil {
    return "", err
  }
  return token, nil
}

func Login(username string, passwd string) (string, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return "", err
  }
  u := &Vms_users{}
  dbs.Where("username=?", username).First(u)
  if u.Passwd == passwd {
    token, err := createtoken(u.Username, strconv.Itoa(u.Id))
    if err != nil {
      return "", err
    }
    return token, err
  } else {
    return "", vmerror.Error{Message: "Authentication failed"}
  }
}
