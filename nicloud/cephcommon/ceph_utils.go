package cephcommon

import (
  "github.com/ceph/go-ceph/rados"
  rbd "github.com/ceph/go-ceph/rbd"
  "nicloud/dbs"
  "nicloud/vmerror"
  "time"
)

type Vms_Ceph struct {
  Uuid string
  Pool string `json:"Pool" validate:"required"`
  Datacenter string `json:"Datacenter" validate:"required"`
  Ceph_secret string  `json:"Ceph_secret" validate:"required"`
  Ips string  `json:"Ips" validate:"required"`
  Port string `json:"Port" validate:"required"`
  Comment string
  Status int8
}

func Delete(uuid string) error {
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }

  dberr := dbs.Where("uuid=?", uuid).Delete(Vms_Ceph{})
  if dberr.Error != nil {
    return dberr.Error
  }

  return nil
}

func Add(name string, pool string, datacenter string, ceph_secret string, ips string, port string, comment  string) error {
  c := &Vms_Ceph{
    Uuid: name,
    Pool: pool,
    Datacenter: datacenter,
    Ceph_secret: ceph_secret,
    Ips: ips,
    Port: port,
    Comment: comment,
    Status: 1,
  }
  dbs, err := db.NicloudDb()
  if err != nil {
    return err
  }
  errdb := dbs.Create(c)
  if errdb.Error != nil {
    return errdb.Error
  }
  dbs.NewRecord(c)
  return nil
}

func Get()([]*Vms_Ceph, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := []*Vms_Ceph{}
  dbs.Find(&c)
  return c, nil
}

func Getpool(datacenter string, storage string)([]*Vms_Ceph, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := []*Vms_Ceph{}
  dbs.Where("datacenter=? and uuid=?", datacenter, storage).Find(&c)
  return c, nil
}

func Cephinfobyname(datacenter string, storage string)(*Vms_Ceph, error) {
  dbs, err := db.NicloudDb()
  if err != nil {
    return nil, err
  }
  c := &Vms_Ceph{}
  dbs.Where("datacenter=? and uuid=?", datacenter, storage).First(c)
  return c, nil
}

func CephConn() (*rados.Conn, error) {
  conn, err := rados.NewConn()
  if err != nil {
    return nil, err
  }
  err = conn.ReadDefaultConfigFile()
  if err != nil {
    return nil, err
  }
  err = conn.Connect()
  if err != nil {
    return nil, err
  }

  return conn, nil
}

func Rm_image(uuid string, pool string) (error) {
  conn, err := CephConn()
  if err != nil {
    return err
  }

  ioctx, err := conn.OpenIOContext(pool)
  if err != nil {
    return err
  }

  img := rbd.GetImage(ioctx, uuid)
  Archive_img := "x_"+(time.Now().Format("200601021504"))+uuid
  img.Rename(Archive_img)

  return nil
}

func RbdClone(id string, cephblock string, snap string, pool string) (string, error) {

  conn, err := CephConn()
  if err != nil {
    return "", vmerror.Error{"ceph 连接失败"}
  }

  ioctx, _ := conn.OpenIOContext(pool)
  imag := rbd.GetImage(ioctx, cephblock)

  _, err = imag.Clone(snap, ioctx, id, rbd.FeatureLayering, 12)
  if err != nil {
    return "", err
  }
  return id, nil
}

func Createcephblock(uuid string, contain int, pool string) error {
  conn, err := CephConn()
  if err != nil {
    return err
  }

  ioctx, _ := conn.OpenIOContext(pool)
  _, err = rbd.Create(ioctx, uuid, uint64(1024*1024*1024*contain), 0)
  if err != nil {
    return err
  }
  return nil
}

func Changename(uuid string) {

}
