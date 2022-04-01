package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

func main(){
  client := &http.Client{}
  data := make(map[string]interface{})
  data["datasetid"] = "31f55d7067cc410fbd5c57bf047b9cc4"
  data["password"] = "a:x^UFB;_D"
  bytesData, _ := json.Marshal(data)

  req, _ := http.NewRequest("POST","http://10.0.85.134:5678/gsapi/ftpapi/dataset_certification",bytes.NewReader(bytesData))
  req.Header.Add("x-token", "ja4x3yZ2gl7Hk+h0WVK3cjyfn6amcYZwyVGSdu9pFoetso3Ph98aoyWOdo+m5126cxnWOcpaOLVKGwoIp3dsG5MYETakdIGhAhca8V7KBusX0pIKl72+o7oVeIUvjCl9qDA0Tw4jbr4ioozJWggrPnp3iuBotaMPgHZm066U2Ajd5CuG4ZqE5qY2y8SrUau8+xP5WcwaUsZAz5WjzYYUfWy21QN/1M2sZyjhejUilpuMpiCvxa4kcZx6zL6NLo3zto+QPL6ouUllbx8X8AoEm0OC7NdaSHJSVLAXgl8jkEc2SpqP4pwI2L16kyGRKGWzA9bFGNN0b0+2OyXmE1nh2Q+UsRUgiEC3jrUZd0dxphKY9EY+cbv26HtVZl/jiAUq/l2HvSYSNf2Av3bKsJ6YBI6V44kR148GEfoH9PVI5Dnlp2EfSLPG6M0AoaJW1HqIatvKkpkrAzRKvlDxPLRLAK7FiV6dZPzkLrvbihGZFL6Ef+MbdQuOU0aInE91VmbAPVCrW+z5ZyPYqv/CKq8jrvdcurZhSsZjxczplURgsIX/B0Hj97NOCWjxoBZuPWQe1w6snNUiNi1NHJosr4FiNPJgpZRiDgUOS2Qb+OjIYRXCh25wBX8oFRFEWhgi315Iq1OrhMrsa3OVMmAWoBNOsBChRHECPhFdPaduJIac3Sc=")
  resp, _ := client.Do(req)
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))
}
