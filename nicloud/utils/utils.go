package utils

import (
  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
  uuid "github.com/satori/go.uuid"
  "strings"
)

func Createuuid() string {
  /*create uuid*/
  u := uuid.NewV4().String()
  return u
}

func ParseToken(token string) (*jwt.StandardClaims, error) {
  jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
    return []byte("nicloud"), nil
  })
  if err == nil && jwtToken != nil {
    if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
      return claim, nil
    }
  }
  return nil, err
}

//token auth middlehandle func
func Tokenauth() gin.HandlerFunc {
  return func(c *gin.Context) {
    token := c.Request.Header.Get("token")
    if len(token) == 0 {
      c.JSON(400, "无法认证，重新登录")
    }

    _, err := ParseToken(strings.Fields(token)[0])
    if err != nil {
      c.JSON(200, "无法认证，重新登录")
    }
  }
}
