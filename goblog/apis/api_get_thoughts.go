package apis

import "github.com/gin-gonic/gin"
import "github.com/niklaus-code/gogogo/common"


func Get_thoughts(c *gin.Context) {
    res := common.ThoughtsGet()
    c.JSON(200,  res)
    }
