package machineapis

import (
  "github.com/gin-gonic/gin"
  "goblog/machinecommon"
  "strconv"
)

func Delmachine(c *gin.Context) {
  id, _ := strconv.Atoi(c.Query("id"))
  res := make(map[string]interface{})
  data, err := machinecommon.Delmachine(id)

  res["res"] = data
  res["err"] = err
  c.JSON(200, res)
}

func Getmachinelist(c *gin.Context) {

  res := make(map[string]interface{})
  r, err := machinecommon.Machinelist()

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
