package apis

import "github.com/gin-gonic/gin"
import "goblog/common"
import "strconv"


func Get_blog_by_id(c *gin.Context) {

    id := c.Param("id")
    blogid, err := strconv.Atoi(id)
    if err != nil {
        return
    }
    bloginfo := common.BlogGetById(blogid)

    res := make(map[string]interface{})
    res["res"] = bloginfo

    c.JSON(200,  res)
    }
