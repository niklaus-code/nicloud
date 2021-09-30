package machinecommon

import "github.com/jinzhu/gorm"

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
  Status string
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
