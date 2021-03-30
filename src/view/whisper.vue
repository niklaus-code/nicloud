<template>
<div>
    <headd></headd>
  <div class="content whisper-content">
    <div class="cont">
      <div class="whisper-list" v-for="blog in data">
        <div class="item-box">
          <div class="item">
            <div class="whisper-title">
              <i class="layui-icon layui-icon-date"></i><span class="date">{{blog["create_time"]}}</span>
            </div>
            <div v-html="blog.content"></div>
            <!--
            <div class="op-list">
                <p class="like"><i @click="point_like(blog['id'])" class="layui-icon layui-icon-praise"></i>
                    <span v-if="pointlike < 0 "> {{blog["like_number"]}}</span>
                    <span v-else>{{pointlike}}</span>
                </p> 
            </div>
                -->
          </div>
          <div class="review-version layui-hide">
              <div class="form">
                <img src="../../static/imges/niklaus.jpg">
                <form class="layui-form" action="">
                  <div class="layui-form-item layui-form-text">
                    <div class="layui-input-block">
                      <textarea name="desc" class="layui-textarea"></textarea>
                    </div>
                  </div>
                  <div class="layui-form-item">
                    <div class="layui-input-block" style="text-align: right;">
                      <button class="layui-btn definite">確定</button>
                    </div>
                  </div>
                </form>
              </div>
              <div class="list-cont">
                <div class="cont">
                  <div class="img">
                    <img src="../../static/imges/niklaus.jpg" alt="" height="20px" width="20px">
                  </div>
                  <div class="text">
                    <p class="tit"><span class="name">吳亦凡</span><span class="data">2018/06/06</span></p>
                    <p class="ct"> eee</p>
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
      pointlike: -1
    }
  },
  components: {
    foot, headd
  },
  mounted: function () {
    this.blog()
    },
  methods: {
    blog: function (id, sig) {
      this.$http.get('/api/blog/get_blog/get_blog_thoughts', {params: {category: 5}}).then(response => {
        this.data = response.data
      })
    },
    point_like: function (id) {
      this.$http.post('/api/blog/get_blog/point_like', {id: id}).then(response => {
        this.blog()
      })
    }
  }
}
</script>
<style>
.content {
    padding: 60px 0 160px 0;
}
</style>
