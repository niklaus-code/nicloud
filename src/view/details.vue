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
                    <p class="cont-info"><span class="data">{{data["create_time"]}}</span><span class="types">{{data["category"]}}</span></p>
                  </div>
                 <div style="background-color:#f9f9f9" v-html="data.content"></div>
                 <!-- <div class="btn-box">
                    <button class="layui-btn layui-btn-primary">上一篇</button>
                    <button class="layui-btn layui-btn-primary">下一篇</button>
                  </div> -->
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
                    <img src="../../static/imges/header.png" alt="">
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
    this.category = this.$route.query.category
    var blogid = this.$route.query.id
    this.blog_id = blogid
    var signature = this.$route.query.signature
    this.blog(blogid, signature)
    },

   methods: {
    blog: function (id, sig) {
      this.$http.post('/api/xianyu/get_user/get_blog_by_id', {id: id, category: this.category}).then(response => {
        this.blog_id = id
        this.blog_like = response.data.res.like_number
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
