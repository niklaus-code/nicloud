<template>
<div>
    <headd></headd>
  <div class="content">
    <div class="cont w1000">
      <div class="title">
        <span class="layui-breadcrumb" lay-separator="|">
          <a href="javascript:;" class="active">设计文章</a>
          <a href="javascript:;">前端文章</a>
          <a href="javascript:;">旅游杂记</a>
        </span>
      </div>
      <div class="list-item" v-for="item in data">
        <div class="item">
          <div class="layui-fluid">
            <div class="layui-row">
              <div class="layui-col-xs12 layui-col-sm4 layui-col-md5">
                <div class="img"><img v-bind:src=item.img alt=""></div>
              </div>
              <div class="layui-col-xs12 layui-col-sm8 layui-col-md7">
                <div class="item-cont" >
                  <h3><router-link :to="{path:'/details', query:{id: item.id,signature: item.signature}}">{{item["title"]}}<button class="layui-btn layui-btn-danger new-icon">new</button></router-link></h3>
                  <h5>{{item["category_name"]}}</h5>
                   <h v-html=item.content>{{item["content"]}}</h>
                </div>
            </div>
            </div>
           </div>
        </div>
      </div>
      <div id="demo" style="text-align: center;"></div>
    </div>
  </div>

    <div style="text-align:center">
       <div class="layui-box layui-laypage layui-laypage-default" id="layui-laypage-1">
            <span @click="start--, reduc(), get_blog()" class="layui-laypage-prev" data-page="1">
                上一页
            </span>
            <span @click="change_page(1), get_blog()" class="layui-laypage-prev" data-page="1">
                First
            </span>
            <span data-page="1">第 {{start}} 页 / 共 {{totalnumber}} 页</span>
            <span @click="change_page(totalnumber), get_blog()" class="layui-laypage-prev" data-page="1">
                Last
            </span>
            <span @click="start++, add(), get_blog()" class="layui-laypage-next" data-page="3">
                下一页
            </span>
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
      start: 1,
      totalnumber: '',
      offset:10
    }
  },

  components: {
    foot, headd
  },

  mounted: function () {
      this.get_blog()
  },

  methods: {
    reduc: function () {
            if (this.start <= 1) {
                this.start=1
                }
            },

    add: function () {
            if (this.start > this.totalnumber) {
                this.start=this.totalnumber
                }
            },

    change_page: function (page) {
        this.start = page
    },

    get_blog: function () {
      var apiurl = `/api/blog/get_blog/get_blog/${this.start}`
      this.$http.get(apiurl).then(response => {
      //this.$http.get('/api/blog/get_blog/get_blog', {params: {start:this.start, offset:this.offset}}).then(response => {
        this.data = response.data.bloglist
         this.totalnumber = response.data.totalnumber
        })
    },

    top_article: function (id) {
      this.$http.post('/api/blog/get_blog/top_article_list', {article_id: id, action: "insert"}).then(response => {
        this.data = response.data
      })
    }
  }
}

</script>


<style scoped>
img {max-height: 300px; width: auto}
.item-cont {max-height:20px}
.item-cont p {overflow:hidden}
.content {
    padding: 0px 0 160px 0;
}
</style>
