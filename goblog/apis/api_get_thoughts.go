package apis

import "github.com/gin-gonic/gin"
import "goblog/common"


func Get_thoughts(c *gin.Context) {
    res := common.ThoughtsGet()
    c.JSON(200,  res)
    }
