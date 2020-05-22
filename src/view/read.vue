<template>
<div>
<headd></headd>
  <div id="layer-photos-demo" class="album-content w1000">
  <!-- <div class="album-content w1000"  id="layer-photos-demo" class="layer-photos-demo"> -->
    <div class="img-info">
      <img src="http://116.85.50.170/static/imges/reading/guangxi/074a03e9899178f74d8d5a17898ed09.jpg" alt="">
      <div class="title">
        <p class="data">拍摄于<span>2018年10月1日</span></p>
        <p class="text">广西北海——银滩</p>
      </div>
    </div>
    <div class="img-list">
      <div class="layui-fluid" style="padding:0">
        <div class="layui-row layui-col-space30 space">
          <div class="layui-col-xs12 layui-col-sm4 layui-col-md4" v-for="blog in data" v-bind:key="blog">
            <div class="item">
                <router-link :to="{path:'/details', query:{id: blog.id, category:'read'}}">
              <img v-bind:src=blog.image>
                </router-link>
              <div class="cont-text">
                <router-link :to="{path:'/details', query:{id: blog.id, category:'read'}}">
                <div class="data"> {{blog["title"]}}</div>
                </router-link>
                <p class="addstatics"><i class="layui-icon layui-icon-location"></i><span>{{blog["author"]}}</span></p>
                <p class="briefly">{{blog["beginning"]}}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div id="demo"></div>
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
      data: ''
    }
  },

  components: {
    foot, headd
  },

  mounted: function () {
    this.read()
  },

  methods: {
    read: function (id, sig) {
      this.$http.get('/api/blog/get_user/get_blog_read').then(response => {
        this.data = response.data
      })
    }
  }
}
</script>
