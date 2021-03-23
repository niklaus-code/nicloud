<template>
<div>
    <headd></headd>

  <div class="content whisper-content leacots-content details-content">
    <div class="cont w1000">
      <div class="whisper-list">
        <div class="item-box">
          <div class="review-version">
              <div class="form-box">
                <div class="article-cont">
                  <div class="title">
                    <h3>{{data["title"]}}</h3>
                    <p class="cont-info">
                        <span class="data" style="color:green">
                            {{data["category_name"]}}
                        </span>
                        <span class="types">
                            {{data["create_time"]}}
                        </span></p>
                  </div>
                 <div style="background-color:#f9f9f9;padding:5px 10px" v-html="data.content"></div>
                </div>
                <div class="form">
                  <form class="layui-form" action="">
                    <div class="layui-form-item layui-form-text">
                      <div class="layui-input-block">
                        <textarea name="desc" placeholder="既然来了，就说几句" class="layui-textarea"></textarea>
                      </div>
                    </div>
                    <div class="layui-form-item">
                      <div class="layui-input-block" style="text-align: right;">
                        <button class="layui-btn definite">確定</button>
                      </div>
                    </div>
                  </form>
                </div>
              </div>
              <div class="volume">
                全部留言 <span>10</span>
              </div>
              <div class="list-cont">
                
                <div class="cont">
                  <div class="img">
                    <img height="40px" width="40px" src="data:image/jpeg;base64,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" alt="图片Base64编码" style="max-width:90%; max-height:2000px;"/>
                  </div>
                  <div class="text">
                    <p class="tit"><span class="name">史塔克</span><span class="data">2018/06/06</span></p>
                    <p class="ct">凛冬将至</p>
                  </div>
                </div>
              </div>
          </div>
        </div>
      </div>
      <div id="demo" style="text-align: center;"></div>
    </div>
  </div>
<foot></foot>
  </div>
</template>
<script>

import foot from '@/components/footer'
import headd from '@/components/head'

export default {
    data () {
        return {
            data: '',
            category: ''
        }
    },

    components: {
        foot, headd
    },

    mounted: function () {
           var id = this.$route.params.id
           if (!id) {
                this.blogid = window.sessionStorage.getItem('blogid')
            } else {
                this.blogid = id
                window.sessionStorage.setItem('blogid', this.blogid)
            }
        this.blog()
    },

    methods: {
        blog: function () {
            var apiurl = `/api/blog/get_blog/get_blog_by_id/${this.blogid}`
            this.$http.get(apiurl).then(response => {
            this.data = response.data.res
            })
        }
    }
  }
</script>
<style>
    .details-content .article-cont p {
    padding:30px 0 0 5px
}
</style>
