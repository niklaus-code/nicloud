package vmapis

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "nicloud/vm"
  "nicloud/vmerror"
)

func Cpudetails(c *gin.Context) {
  uuid := c.Query("uuid")
  host := c.Query("host")

 var res = make(map[string]interface{})
 cpu, err := vm.Cpuinfo(host, uuid)
 if err != nil {
   res["err"] = vmerror.Error{Message: err.Error()}
   c.JSON(http.StatusInternalServerError, res)
   return
 }
 res["res"] = cpu
 res["err"] = nil
 c.JSON(http.StatusOK, res)
}

func Memdetails(c *gin.Context) {
  uuid := c.Query("uuid")
  host := c.Query("host")
 var res = make(map[string]interface{})
 m, err := vm.Meminfo(host, uuid)
 if err != nil {
   res["err"] = vmerror.Error{Message: err.Error()}
   c.JSON(http.StatusInternalServerError, res)
   return
 }

 res["res"] = m
 res["err"] = nil
 c.JSON(http.StatusOK, res)
}

func Netdetails(c *gin.Context) {
  uuid := c.Query("uuid")
  host := c.Query("host")
 var res = make(map[string]interface{})
 n, err := vm.Netinfo(host, uuid)
 if err != nil {
   res["err"] = vmerror.Error{Message: err.Error()}
   c.JSON(http.StatusInternalServerError, res)
   return
 }

 res["res"] = n
 res["err"] = nil
 c.JSON(http.StatusOK, res)
}

func Diskdetails(c *gin.Context) {
  uuid := c.Query("uuid")
  host := c.Query("host")
  var res = make(map[string]interface{})
  dio, err := vm.Diskinfo(host, uuid, "linux")
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
    c.JSON(http.StatusInternalServerError, res)
    return
  }

  res["res"] = dio
  res["err"] = nil
  c.JSON(http.StatusOK, res)
}
