package ceph

import (
	"fmt"
	"strings"

	"github.com/beevik/etree"
	"github.com/go-ini/ini"
)

func Xml(vcpu int, vmem int, uuid string, mac string) (string, error) {
	var cfg, _ = ini.Load("conf/setting.ini")
	var port = cfg.Section("ceph").Key("port").String()

	ips := strings.Split(cfg.Section("ceph").Key("ip").String(), ",")
	br := cfg.Section("bridge").Key("br").String()

	f := "/home/ysman/niklaus-blog/goblog/conf/vm.xml"

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(f); err != nil {
		return doc.Text(), err
	}

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

	fmt.Println("!!!!!!!!!!")
	fmt.Println(br)

	bridge := doc.FindElement("./domain/devices/interface/source")
	bridge.CreateAttr("bridge", fmt.Sprintf("%s", br))

  macaddr := doc.FindElement("./domain/devices/interface/mac")
  macaddr.CreateAttr("address", fmt.Sprintf("%s", mac))

	for _, e := range doc.FindElements("./domain/devices[1]/*") {
		if e.Tag == "disk" {
			for _, v := range e.ChildElements() {
				if v.Tag == "source" {
					v.CreateAttr("name", "vm/x_20210806095906_d90834be30f346f58e02e656927c9122")

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
	docstring, err := doc.WriteToString()
	if err != nil {
		return "", err
	}

	return docstring, nil
}
