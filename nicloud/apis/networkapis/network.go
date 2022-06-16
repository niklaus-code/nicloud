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

  validate := validator.New()
  err := validate.Struct(v)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  err = networks.AddVlan(datacenter, vlan, bridge, network, prefix, gateway)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func Get(c *gin.Context) {
  vlans, err := networks.Getvlan()
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, vlans)
}

func Getvlanbydatacenter(c *gin.Context) {
  datacenter := c.Query("datacenter")

  vlans, err := networks.Getvlanbydatacenter(datacenter)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, vlans)
}

func Delete(c *gin.Context) {
  vlan := c.Query("vlan")

  err := networks.DeleteVlan(vlan)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}


func Deleteip(c *gin.Context) {
  vlan := c.Query("vlan")
  ip := c.Query("ip")
  ipcheck := vm.GetVmByIp(ip)

  if ipcheck.Ip != "" {
    vmerror.SERVERERROR(c, vmerror.Error{Message: "ip已被占用"})
    return
  }
  err := networks.Deleteip(ip, vlan)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func CreateIp(c *gin.Context) {
  startip := c.Query("startip")
  endip := c.Query("endip")
  vlan := c.Query("vlan")
  prefix, err := strconv.Atoi(c.Query("prefix"))
  gateway := c.Query("gateway")

  if err != nil {
    vmerror.REQUESTERROR(c, err)
    return
  }

  if prefix >= 32 || prefix < 8 {
    vmerror.REQUESTERROR(c, vmerror.Error{"参数格式错误"})
    return
  }

  err = networks.Createip(startip, endip, vlan, prefix, gateway)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func UpIp(c *gin.Context) {
  ipv4 := c.Query("ipv4")
  vlan := c.Query("vlan")

  err := networks.OpIP(ipv4, vlan, 0)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func DownIp(c *gin.Context) {
  ipv4 := c.Query("ipv4")
  vlan := c.Query("vlan")

  err := networks.OpIP(ipv4, vlan, 1)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, nil)
}

func GetIplist(c *gin.Context) {
  vlan := c.Query("vlan")
  iplist, err:= networks.IPlist(vlan)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }
  vmerror.SUCCESS(c, iplist)
}

func GetallIp(c *gin.Context) {
  vlan := c.Query("vlan")
  iplist, err := networks.AllIP(vlan)
  if err != nil {
    vmerror.SERVERERROR(c, err)
    return
  }

  vmerror.SUCCESS(c, iplist)
}

func DownloadExcel(c *gin.Context) {
  vlan := c.Query("vlan")

  c.Header("Content-Type", "application/octet-stream")
  c.Header("Content-Disposition", "attachment; filename=IPLists.txt")
  c.Header("Content-Transfer-Encoding", "binary")
  file, _ := os.OpenFile("test2.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
  defer file.Close()

  ipliststr, err := networks.Downloadips(vlan)
  if err != nil {
    con := bytes.NewReader([]byte(err.Error()))
    http.ServeContent(c.Writer, c.Request, "vlan"+vlan, time.Now(), con)
    return
  }
  con := bytes.NewReader([]byte(ipliststr))
  http.ServeContent(c.Writer, c.Request, "vlan"+vlan, time.Now(), con)
}
