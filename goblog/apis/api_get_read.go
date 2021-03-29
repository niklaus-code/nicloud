package apis

import "github.com/gin-gonic/gin"
import "goblog/common"


func Get_read(c *gin.Context) {
    res := common.ReadGet()
    c.JSON(200,  res)
    }
