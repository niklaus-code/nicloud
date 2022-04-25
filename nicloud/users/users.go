package users

import (
  "github.com/dgrijalva/jwt-go"
  db "nicloud/dbs"
  "nicloud/vmerror"
  "reflect"
  "strconv"
  "time"
)

type Vms_users struct {
  Id int `gorm:"primary_key;AUTO_INCREMENT"`
  Username string `gorm:"unique;" json:"Username" validate:"required"`
  Passwd string `json:"Passwd" validate:"required"`
  Email string  `jgorm:"unique;" son:"Email" validate:"email"`
  Role int  `json:"Role" validate:"oneof=1 2"`
  Mobile string `gorm:"unique;" json:"Mobile" validate:"len=11" validate:"startswith=1"`
  Create_time time.Time
}

func DelUser(userid int) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  r := &Vms_users{}
  errdb := dbs.Where("id=?", userid).Delete(r)
  if errdb.Error != nil {
    return errdb.Error
  }

  return nil
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

func mapuser(obj []Vms_users) []map[string]interface{}  {
  var mapc []map[string]interface{}

  for _, v := range obj {
    c := make(map[string]interface{})
    m := reflect.TypeOf(v)
    n := reflect.ValueOf(v)
    for i := 0; i < m.NumField(); i++ {
      c[m.Field(i).Name] = n.Field(i).Interface()
    }

    role, err := GetRoleByRoleId(v.Role)
    if err != nil {
      c["Role"] = "nil"
    } else {
      c["Role"] = role.Rolename
    }
    mapc = append(mapc, c)
  }
  return mapc
}

func GetUsers() ([]map[string]interface{}, error){
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  var u []Vms_users
  dberr := dbs.Find(&u)
  if dberr.Error != nil {
    return nil, dberr.Error
  }
  return mapuser(u), nil
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

func CheckPWD(username string, passwd string) (string, string, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return "", "", vmerror.Error{Message: "数据库连接错误"}
  }
  u := &Vms_users{}
  dberr := dbs.Where("username=?", username).First(u).Error
  if dberr != nil {
    return "", "", dberr
  }
  if u.Passwd == passwd {
    uid := strconv.Itoa(u.Id)
    token, err := createtoken(u.Username, uid)
    if err != nil {
      return "", "", err
    }
    return token, username, err
  } else {
    return "", "", vmerror.Error{Message: "账号或者密码错误"}
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
  Rolename string `gorm:"unique;"`
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

func ChangePasswd(username string, oldpasswd string, newpasswd string)  error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  _, username, err = CheckPWD(username, oldpasswd)
  if err != nil {
    return err
  }

  if len(username) == 0 {
    return vmerror.Error{Message: "修改密码失败"}
  }

  dberr := dbs.Model(&Vms_users{}).Where("username = ? or email = ?", username, username).Update("passwd", newpasswd).Error
  if dberr != nil {
    return dberr
  }

  return nil
}
