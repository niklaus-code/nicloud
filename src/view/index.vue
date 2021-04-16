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
                <router-link :to="{name:'details', params:{id: item.id}}">
                <div class="item" style="height:200px;background-color:#f9f9f9">
                    <!--
                    <div class="layui-fluid">
                        <div class="layui-row" style="height:30px">
                        <div class="layui-col-xs12 layui-col-sm4 layui-col-md5">
                        <div class="img"><img v-bind:src=item.img alt=""></div>
                    </div>
                    -->
                    <div class="layui-col-xs12 layui-col-sm8 layui-col-md7">
                        <div class="item-cont" >
                            <h4 style="line-height:2.4">
                                {{item["title"]}}
                                <button style="margin-left:3px" :class="classtype[item['category_name']]" class="layui-btn-radius layui-btn-xs">
                                    {{item["category_name"]}}
                                	<!-- <button style="margin-left:3px" class="layui-btn-radius layui-btn-xs layui-btn-danger new-icon">{{item["category_name"]}} -->
                                </button>
                            </h4>
                            <!-- <font>{{item["category_name"]}}</font> -->
                            <font v-html=item.content>{{item["content"]}}</font>
                        </div>
                    </div>
                </div>
            </router-link>
            </div>
        </div>
        <div>
            <div id="demo" style="text-align: center;"></div>
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
        offset:10,
        classtype: {
            buttonclasssize: 'layui-btn-xs',
            存储: 'layui-btn',
            数据库: 'layui-btn-normal',
            虚拟化: 'layui-btn-warm',
            杂文: 'layui-btn-danger'
            }
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
            this.data = response.data.bloglist
            this.totalnumber = response.data.totalnumber
        })
    }
  }
}

</script>
<style scoped>
a {color: black}
img {max-height: 300px; width: auto}
.item-cont {max-height:20px}
.content {
    padding: 0;
}
</style>
