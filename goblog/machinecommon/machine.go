package machinecommon

import (
  "github.com/jinzhu/gorm"
)

type Machineroom struct {
  Id  int
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
  Shebeizhuangtai int
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

func Delmachine(id int, start int, offset int) ([]*Machineroom, error)  {
  db := mcdb()
  db.Model(&Machineroom{}).Where("id=?", id).Update("status", 0)

  return Machinelist(start, offset)
}

func Machinelist(startpage int, offset int) ([]*Machineroom, error)  {
  offsetpage := (startpage-1)*offset

  db := mcdb()
  var v []*Machineroom
  db.Where("status=1").Order("suoshujifang").Limit(offset).Offset(offsetpage).Find(&v)

  return v, nil
}

func Allpage() (int, int, error)  {
  db := mcdb()
  var v []*Machineroom
  db.Where("status=1").Find(&v)
  allpage := len(v)/100+1
  return len(v), allpage, nil
}

func Addmacine(zichangmingcheng string, pingpai string,  Xinghao string, Xuliehao string,
  zichanbiaoqian string, danwei string, suoshubumen string,  zichanzerenbumen string, zerenren string,
  suoshujifang string, jigui string,  jiguizichanbiaoqian string, weizhi string,  gaodu string,
  shebeizhuangtai int, edinggonglv string, yongdiandengji string, guanliip string, yewuip string,
  beizhu string) error {

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
    Status: 1,
  }

  db := mcdb()
  err := db.Create(&m)
  if err != nil {
    return nil
  }

  db.NewRecord(&m)
  return nil
}
