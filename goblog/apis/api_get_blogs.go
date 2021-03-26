package apis

import "github.com/gin-gonic/gin"
import "github.com/niklaus-code/gogogo/common"
import "strconv"


func Get_blog(c *gin.Context) {

    pagenumber := c.Param("pagenumber")
    page_number,err := strconv.Atoi(pagenumber)
    if err != nil {
        return 
    }
    startpage := (page_number-1)*8
    offset := 8
    bloglist := common.BlogGet(startpage, offset)

    res := make(map[string]interface{})
    res["totalnumber"] = 5
    res["bloglist"] = bloglist

    c.JSON(200,  res)
    }
