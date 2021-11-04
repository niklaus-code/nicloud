package apis

import "github.com/gin-gonic/gin"
import "goblog/common"


func Get_thoughts(c *gin.Context) {
    r, err := common.ThoughtsGet()
    res := make(map[string]interface{})
    res["res"] = r
    res["err"] = err
    c.JSON(200,  res)
    }
