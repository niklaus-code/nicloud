package apis

import "github.com/gin-gonic/gin"
import "goblog/common"


func Get_read(c *gin.Context) {
    r, err := common.ReadGet()
    res := make(map[string]interface{})
    res["res"] =r
    res["err"] = err
    c.JSON(200,  res)
    }
