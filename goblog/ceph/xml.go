package ceph

import (
  "fmt"
  "strings"

  "github.com/beevik/etree"
  "github.com/go-ini/ini"
  "goblog/dbs"
)

type Vms_os struct {
  Osname string
  Snapimage string
  Xml string
}

func getxml(osname string) (string, error) {
  db, err := db.NicloudDb()
  if err != nil {
    return "", err
  }

  var v []*Vms_os
  db.Where("osname=?", osname).Find(&v)
  return v[0].Xml, nil
}

func Xml(vcpu int, vmem int, uuid string, mac string, image_name string, osname string) (string, error) {
	var cfg, _ = ini.Load("conf/setting.ini")
	var port = cfg.Section("ceph").Key("port").String()
  var ceph_secret = cfg.Section("ceph").Key("ceph_secret").String()

	ips := strings.Split(cfg.Section("ceph").Key("ip").String(), ",")
	br := cfg.Section("bridge").Key("br").String()

	f, err := getxml(osname)
	if err != nil {
    return "", err
  }

	doc := etree.NewDocument()
  err = doc.ReadFromString(f)
	if err != nil {
		return "", err
	}

	cephsecret := doc.FindElement("./domain/devices/disk/auth/secret")
	cephsecret.CreateAttr("uuid", ceph_secret)

	cpu := doc.FindElement("./domain/vcpu")
	cpu.CreateText(fmt.Sprintf("%d", vcpu))

	id := doc.FindElement("./domain/uuid")
	id.CreateText(uuid)

	name := doc.FindElement("./domain/name")
	name.CreateText(uuid)

	mem := doc.FindElement("./domain/memory")
	mem.CreateText(fmt.Sprintf("%d", vmem))

	currentMemory := doc.FindElement("./domain/currentMemory")
	currentMemory.CreateText(fmt.Sprintf("%d", vmem))

	bridge := doc.FindElement("./domain/devices/interface/source")
	bridge.CreateAttr("bridge", fmt.Sprintf("%s", br))

  macaddr := doc.FindElement("./domain/devices/interface/mac")
  macaddr.CreateAttr("address", fmt.Sprintf("%s", mac))

	for _, e := range doc.FindElements("./domain/devices[1]/*") {
		if e.Tag == "disk" {
			for _, v := range e.ChildElements() {
				if v.Tag == "source" {
					v.CreateAttr("name", fmt.Sprintf("vm/%s", image_name))

					for ip_k, ip := range ips {
						v.CreateElement("host")
						v.ChildElements()[ip_k].CreateAttr("name", string(ip))
						v.ChildElements()[ip_k].CreateAttr("port", port)
					}
				}
			}
		}
	}
	doc.Indent(2)
	var docstring string
	docstring, err = doc.WriteToString()
	if err != nil {
		return "", err
	}

	return docstring, nil
}
