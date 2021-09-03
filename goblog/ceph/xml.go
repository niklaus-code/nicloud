package ceph

import (
  "fmt"
  "github.com/beevik/etree"
)


func Xml(vcpu int, vmem int, uuid string) (string, error) {
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

  for _, e := range doc.FindElements("./domain/devices[1]/*") {
    if e.Tag == "disk" {
      for _, v := range e.ChildElements() {
        if v.Tag == "source" {
          v.CreateAttr("name", "vm/x_20210810befb2bf3b7c3453eb2bd99906a7f6c5f")

          v.CreateElement("host")
          v.ChildElements()[0].CreateAttr("name", "10.0.82.153")
          v.ChildElements()[0].CreateAttr("port", "6789")

          v.CreateElement("host")
          v.ChildElements()[1].CreateAttr("name", "10.0.82.152")
          v.ChildElements()[1].CreateAttr("port", "6789")
        }
      }
    }
  }
  doc.Indent(2)
  var docstring  string
  docstring, err := doc.WriteToString()
  if err != nil {
    return "", err
  }

  return docstring, nil
}
