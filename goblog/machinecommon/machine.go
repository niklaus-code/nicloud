package machinecommon

import (
  "github.com/jinzhu/gorm"
)

type Machineroom struct {
  Zichanmingcheng string
  Pinpai string
  Xinghao string
  Xuliehao string
  Zichanbiaoqian string
  Danwei string
  Suoshubumen string
  Zichanzerenbumen string
  Zerenren string
  Suoshujifang string
  Jigui string
  Jiguizichanbiaoqian string
  Weizhi string
  Gaodu string
  Shebeizhuangtai string
  Edinggonglv string
  Yongdiandengji string
  Guanliip string
  Yewuip string
  Beizhu string
  Status int
}


func mcdb() *gorm.DB {
  db, errDb := gorm.Open("mysql", "modis:modis@(127.0.0.1:3306)/nicloud?parseTime=true")
  if errDb != nil {
    return  nil
  }
  return db
}

func Machinelist() ([]*Machineroom, error)  {
  db := mcdb()
  var v []*Machineroom
  db.Find(&v)

  return v, nil
}

func Addmacine(zichangmingcheng string, pingpai string,  Xinghao string, Xuliehao string,
  zichanbiaoqian string, danwei string, suoshubumen string,  zichanzerenbumen string, zerenren string,
  suoshujifang string, jigui string,  jiguizichanbiaoqian string, weizhi string,  gaodu string,
  shebeizhuangtai string, edinggonglv string, yongdiandengji string, guanliip string, yewuip string,
  beizhu string, status int) error {

  m := &Machineroom{
    Zichanmingcheng: zichangmingcheng,
    Pinpai: pingpai,
    Xinghao: Xinghao,
    Xuliehao: Xuliehao,
    Zichanbiaoqian: zichanbiaoqian,
    Danwei: danwei,
    Suoshubumen: suoshubumen,
    Zichanzerenbumen: zichanzerenbumen,
    Zerenren: zerenren,
    Suoshujifang: suoshujifang,
    Jigui: jigui,
    Jiguizichanbiaoqian: jiguizichanbiaoqian,
    Weizhi: weizhi,
    Gaodu: gaodu,
    Shebeizhuangtai: shebeizhuangtai,
    Edinggonglv: edinggonglv,
    Yongdiandengji: yongdiandengji,
    Guanliip: guanliip,
    Yewuip: yewuip,
    Beizhu: beizhu,
    Status: status,
  }

  db := mcdb()
  err := db.Create(*m)
  if err != nil {
    return nil
  }

  db.NewRecord(&m)
  return nil
}
