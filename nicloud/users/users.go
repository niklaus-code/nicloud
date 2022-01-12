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
  expiresTime := time.Now().Unix() + int64(86400)

  claims := jwt.StandardClaims{
    Audience:  "ADMIN",     // 受众
    ExpiresAt: expiresTime,       // 失效时间
    Id:        userid,   // 编号
    IssuedAt:  time.Now().Unix(), // 签发时间
    Issuer:    username,       // 签发人
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

func Login(username string, passwd string) (string, string, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return "", "", err
  }
  u := &Vms_users{}
  dbs.Where("username=?", username).First(u)
  if u.Passwd == passwd {
    uid := strconv.Itoa(u.Id)
    token, err := createtoken(u.Username, uid)
    if err != nil {
      return "", "", err
    }
    return token, username, err
  } else {
    return "", "", vmerror.Error{Message: "登陆失败"}
  }
}


func GetUsernameById(userid int) (string, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return "", err
  }
  u := &Vms_users{}
  errdb := dbs.Where("id=?", userid).First(u)
  if errdb.Error != nil {
    return "", errdb.Error
  }

  return u.Username, nil
}
