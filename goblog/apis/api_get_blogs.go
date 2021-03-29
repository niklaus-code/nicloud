package apis

import "github.com/gin-gonic/gin"
import "goblog/common"
import "strconv"


func Get_blog(c *gin.Context) {

    pagenumber := c.Param("pagenumber")
    startpage,err := strconv.Atoi(pagenumber)
    if err != nil {
        return
    }
    offset := 8
    bloglist := common.BlogGet(startpage, offset)

    res := make(map[string]interface{})
    res["totalnumber"] = 5
    res["bloglist"] = bloglist

    c.JSON(200,  res)
    }
