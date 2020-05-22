# coding=utf-8
import requests
import qrcode


class qr(object):
    def __init__(self):
        self.img_path = "/var/www/mysite/html/static/img/qr_img/qr_"
        self.img_url = "http://vue.manyushuai.site/#/blog?"

    def make_pic(self, id, signature):
        try:
            img = self.img_path + str(id) + ".png"
            url = self.img_url+"id="+str(id)+"&signature="+signature
            qr_code = qrcode.make(url)
            print url
            qr_code.save(img)
        except:
            return False


res = requests.post("http://vue.manyushuai.site/api/xianyu/get_user/get_blog")

qr_obj = qr()

for item in res.json()["name"]:
    try:
        qr_obj.make_pic(item[0], item[5])
    except:
        print "error id="+item[0]
