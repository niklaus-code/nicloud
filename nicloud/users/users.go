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
  Username string `json:"Username" validate:"required"`
  Passwd string `json:"Passwd" validate:"required"`
  Email string  `json:"Email" validate:"email"`
  Role int  `json:"Role" validate:"oneof=1 2"`
  Mobile string `json:"Mobile" validate:"len=11" validate:"startswith=1"`
  Create_time time.Time
}

func Getrolebyrolename(rolename string) (*Vms_roles, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  r := &Vms_roles{}
  errdb := dbs.Where("rolename=?", rolename).First(r)
  if errdb.Error != nil {
    return nil, errdb.Error
  }

  return r, err
}

func Createuser(username string, passwd string, email string, roleid int, mobile string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  u := &Vms_users{
    Username: username,
    Passwd: passwd,
    Email: email,
    Role: roleid,
    Mobile: mobile,
    Create_time: time.Now(),
  }
  errdb := dbs.Create(u)
  if errdb.Error != nil {
    return errdb.Error
  }

  return nil
}

func GetUsers() ([]*Vms_users, error){
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var u []*Vms_users
  dberr := dbs.Find(&u)
  if dberr.Error != nil {
    return nil, dberr.Error
  }
  return u, nil
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
    return "", "", vmerror.Error{Message: "数据库连接错误"}
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


func GetUserByUserID(userid int) (*Vms_users, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  u := &Vms_users{}
  errdb := dbs.Where("id=?", userid).First(u)
  if errdb.Error != nil {
    return nil, errdb.Error
  }

  return u, nil
}

type Vms_roles struct {
  Id int
  Rolename string
}

func GetRoleByRoleId(roleid int) (*Vms_roles, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  r := &Vms_roles{}
  errdb := dbs.Where("id=?", roleid).First(r)
  if errdb.Error != nil {
    return nil, errdb.Error
  }

  return r, nil
}

func GetrAllRoles() ([]*Vms_roles,error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }

  var r []*Vms_roles
  errdb := dbs.Find(&r)
  if errdb.Error != nil {
    return nil, errdb.Error
  }

  return r, nil
}
