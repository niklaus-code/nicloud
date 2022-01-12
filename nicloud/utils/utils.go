package utils

import (
  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
  uuid "github.com/satori/go.uuid"
  "strings"
  "strconv"
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
      res["err"] = "认证失败，请重新登陆"
      c.Abort()
      c.JSON(400, res)
    }

    _, err := ParseToken(strings.Fields(token)[0])
    if err != nil {
      res["err"] = "认证过期，请重新登陆"
      c.Abort()
      c.JSON(200, res)
    }
  }
}
