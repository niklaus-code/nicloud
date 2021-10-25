package ceph

import (
  "github.com/ceph/go-ceph/rados"
  rbd "github.com/ceph/go-ceph/rbd"
  "time"
)

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


func Rm_image(uuid string) (error) {
  conn, err := CephConn()
  if err != nil {
    return err
  }

  ioctx, err := conn.OpenIOContext("vm")
  if err != nil {
    return err
  }

  img := rbd.GetImage(ioctx, uuid)
  Archive_img := "x_"+(time.Now().Format("200601021504"))+uuid
  img.Rename(Archive_img)

  return nil
}

func RbdClone(id string) (string, error) {

  conn, err := CephConn()
  if err != nil {
    return "", err
  }

  ioctx, _ := conn.OpenIOContext("vm")
  img := rbd.GetImage(ioctx, "0000_demo_centos7")
  _, e := img.Clone("20210806_095737", ioctx, id, rbd.RbdFeatureLayering, 0)

  if e != nil {
    return "", e
  }
  return id, nil
}
