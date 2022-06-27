package utils

import (
  "crypto/hmac"
  "crypto/sha256"
  "encoding/hex"
  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
  uuid "github.com/satori/go.uuid"
  "net/http"
  "nicloud/users"
  "nicloud/vmerror"
  "strconv"
  "strings"
)

func Createuuid() string {
  /*create uuid*/
  u := uuid.NewV4().String()
  return u
}

func ParseToken(token string) (int, error) {
  jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
    return []byte("nicloud"), nil
  })
  if err == nil && jwtToken != nil {
    if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
      uid, err := strconv.Atoi(claim.Id)
      if err != nil {
        return 0, err
      }
      return uid, nil
    }
  }
  return 0, err
}

//token auth middlehandle func
func Tokenauth() gin.HandlerFunc {
  res := make(map[string]interface{})
  return func(c *gin.Context) {
    token := c.Request.Header.Get("token")
    if len(token) == 0 {
      res["err"] = vmerror.Error{Message: "认证失败，请重新登陆"}
      c.Abort()
      c.JSON(http.StatusBadRequest, res)
      return
    }

    _, err := ParseToken(strings.Fields(token)[0])
    if err != nil {
      res["err"] = vmerror.Error{Message: "认证过期，请重新登陆"}
      c.Abort()
      c.JSON(http.StatusOK, res)
      return
    }
  }
}

func RoleAuth() gin.HandlerFunc {
  res := make(map[string]interface{})
  return func(c *gin.Context) {
    token := c.Request.Header.Get("token")
    userid, err := ParseToken(strings.Fields(token)[0])
    if err != nil {
      res["err"] = vmerror.Error{Message: "认证过期，请重新登陆"}
      c.Abort()
      c.JSON(http.StatusOK, res)
      return
      }

    user, err := users.GetUserByUserID(userid)
    if err != nil {
      res["err"] = vmerror.Error{Message: "认证失败"}
      c.Abort()
      c.JSON(http.StatusOK, res)
      return
    }

    role, err := users.GetRoleByRoleId(user.Role)
    if err != nil {
      res["err"] = vmerror.Error{Message: "没有权限，请联系管理员"}
      c.Abort()
      c.JSON(http.StatusOK, res)
      return
    }

    if role.Rolename != "admin" {
      res["err"] = vmerror.Error{Message: "没有权限，请联系管理员"}
      c.Abort()
      c.JSON(http.StatusOK, res)
      return
    }
  }
}


func Encryption(data string) string {
  secret := "nicloud"

  // Create a new HMAC by defining the hash type and the key (as byte array)
  h := hmac.New(sha256.New, []byte(secret))

  // Write Data to it
  h.Write([]byte(data))

  // Get result and encode as hexadecimal string
  sha := hex.EncodeToString(h.Sum(nil))

  return sha
}
