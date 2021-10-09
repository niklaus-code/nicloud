package machineapis

import (
  "github.com/gin-gonic/gin"
  "goblog/machinecommon"
  "strconv"
)

func Search(c *gin.Context) {
  res := make(map[string]interface{})
  content := c.Query("content")
  r, err := machinecommon.Searchmachine(content)

  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}

func Getpage(c *gin.Context) {
  res := make(map[string]interface{})
  totalnumber, pagenumber,  err := machinecommon.Allpage()

  res["pagenumber"] = pagenumber
  res["totalnumber"] = totalnumber
  res["err"] = err
  c.JSON(200, res)
}

func Delmachine(c *gin.Context) {
  id, _ := strconv.Atoi(c.Query("id"))
  start, _ := strconv.Atoi(c.Query("startpage"))
  offset, _ := strconv.Atoi(c.Query("offset"))
  res := make(map[string]interface{})
  data, err := machinecommon.Delmachine(id, start, offset)

  res["res"] = data
  res["err"] = err
  c.JSON(200, res)
}

func Getmachinelist(c *gin.Context) {
  start, _ := strconv.Atoi(c.Query("startpage"))
  offset, _ := strconv.Atoi(c.Query("offset"))

  res := make(map[string]interface{})
  r, err := machinecommon.Machinelist(start, offset)

  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}

func Addmachine(c *gin.Context) {

  zichanmingcheng := c.Query("zichanmingcheng")
  pingpai := c.Query("pinpai")
  Xinghao := c.Query("xinghao")
  Xuliehao := c.Query("xuliehao")
  zichanbiaoqian := c.Query("zichanbiaoqian")
  danwei := c.Query("danwei")
  suoshubumen := c.Query("suoshubumen")
  zichanzerenbumen := c.Query("zichanzerenbumen")
  zerenren := c.Query("zerenren")
  suoshujifang := c.Query("suoshujifang")
  jigui := c.Query("jigui")
  jiguizichanbiaoqian := c.Query("jiguizichanbiaoqian")
  weizhi := c.Query("weizhi")
  gaodu  := c.Query("gaodu")
  shebeizhuangtai, _ := strconv.Atoi(c.Query("shebeizhuangtai"))
  edinggonglv := c.Query("edinggonglv")
  yongdiandengji := c.Query("yongdiandengji")
  guanliip := c.Query("guanliip")
  yewuip := c.Query("yewuip")
  beizhu  := c.Query("beizhu ")


  res := make(map[string]interface{})
  err := machinecommon.Addmacine(zichanmingcheng, pingpai, Xinghao, Xuliehao, zichanbiaoqian,danwei,suoshubumen,
    zichanzerenbumen, zerenren,  suoshujifang,jiguizichanbiaoqian, weizhi, jigui, gaodu, shebeizhuangtai, edinggonglv, yongdiandengji,
    guanliip, yewuip, beizhu)

  res["res"] = true
  res["err"] = err
  c.JSON(200, res)
}
