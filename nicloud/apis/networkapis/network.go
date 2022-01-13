package networkapis

import (
  "bytes"
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "net/http"
  "nicloud/networks"
  "nicloud/vm"
  "nicloud/vmerror"
  "os"
  "strconv"
  "time"
)

func Add(c *gin.Context) {
  vlan := c.PostForm("vlan")
  bridge := c.PostForm("bridge")
  network := c.PostForm("network")
  prefix, _ := strconv.Atoi(c.PostForm("prefix"))
  gateway := c.PostForm("gateway")
  datacenter := c.PostForm("datacenter")

  v := networks.Vms_vlans{
    Vlan: vlan,
    Bridge: bridge,
    Network: network,
    Prefix: prefix,
    Gateway: gateway,
    Datacenter: datacenter,
  }

  res := make(map[string]interface{})
  validate := validator.New()
  err := validate.Struct(v)
  if err != nil {
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(400, res)
    return
  }

  err = networks.AddVlan(datacenter, vlan, bridge, network, prefix, gateway)
  if err != nil {
    res["err"] = err
    c.JSON(200, res)
    return
  }
  res["err"] = nil
  c.JSON(200, res)
}

func Get(c *gin.Context) {
  res := make(map[string]interface{})

  vlans, err := networks.Getvlan()
  res["res"] = vlans
  res["err"] = err
  c.JSON(200, res)
}

func Getvlanbydatacenter(c *gin.Context) {
  datacenter := c.Query("datacenter")
  res := make(map[string]interface{})
  vlans, err := networks.Getvlanbydatacenter(datacenter)
  res["res"] = vlans
  res["err"] = err
  c.JSON(200, res)
}

func Delete(c *gin.Context) {
  vlan := c.Query("vlan")

  res := make(map[string]interface{})

  err := networks.DeleteVlan(vlan)
  res["err"] = err
  c.JSON(200, res)
}


func Deleteip(c *gin.Context) {
 vlan := c.Query("vlan")
 ip := c.Query("ip")
 res := make(map[string]interface{})
 ipcheck := vm.GetVmByIp(ip)

 if ipcheck.Ip != "" {
   res["err"] = vmerror.Error{Message: "ip已被占用"}
 } else {
   res["err"] = networks.Deleteip(ip, vlan)
 }
 c.JSON(200, res)
}

func CreateIp(c *gin.Context) {
  startip := c.Query("startip")
  endip := c.Query("endip")
  vlan := c.Query("vlan")
  prefix, err := strconv.Atoi(c.Query("prefix"))
  gateway := c.Query("gateway")

  res := make(map[string]interface{})
  if err != nil {
    res["err"] = err
    c.JSON(400, res)
  }

  if prefix >= 32 || prefix < 8 {
    res["err"] = err
    c.JSON(400, vmerror.Error{"参数错误"})
  }

  err = networks.Createip(startip, endip, vlan, prefix, gateway)
  res["err"] = err
  c.JSON(200, res)
}

func UpIp(c *gin.Context) {
  ipv4 := c.Query("ipv4")
  vlan := c.Query("vlan")
  res := make(map[string]interface{})

  err := networks.OpIP(ipv4, vlan, 0)
  res["err"] = err
  c.JSON(200, res)
}

func DownIp(c *gin.Context) {
  ipv4 := c.Query("ipv4")
  vlan := c.Query("vlan")
  res := make(map[string]interface{})

  err := networks.OpIP(ipv4, vlan, 1)
  res["err"] = err
  c.JSON(200, res)
}

func GetIplist(c *gin.Context) {
  vlan := c.Query("vlan")
  iplist := networks.IPlist(vlan)
  res := make(map[string]interface{})
  res["res"] = iplist

  c.JSON(200, res)
}

func GetallIp(c *gin.Context) {
  vlan := c.Query("vlan")
  iplist := networks.AllIP(vlan)
  res := make(map[string]interface{})
  res["res"] = iplist

  c.JSON(200, res)
}

func DownloadExcel(c *gin.Context) {
  vlan := c.Query("vlan")

  c.Header("Content-Type", "application/octet-stream")
  c.Header("Content-Disposition", "attachment; filename=IPLists.txt")
  c.Header("Content-Transfer-Encoding", "binary")
  file, _ := os.OpenFile("test2.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
  defer file.Close()

  ipliststr := networks.Downloadips(vlan)
  con := bytes.NewReader([]byte(ipliststr))
  http.ServeContent(c.Writer, c.Request, "asd", time.Now(), con)
}
