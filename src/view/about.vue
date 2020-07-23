<template>
<div>
    <headd></headd>
  <div class="about-content">
    <div class="w1000">
      <div class="item info">
        <div class="title">
          <h3>我的介绍</h3>
        </div>
        <div class="cont">
          <img src="http://116.85.50.170/static/imges/niklaus.jpg">
          <div class="per-info">
            <p>
              <span class="name">NiKlaus</span><br/>
              <span class="Career">写代码的</span><br/>
              <span class="interest">喜欢  健身，垒代码， 做饭，美剧~</span>
            </p>
          </div>
        </div>
      </div>
      <div class="item tool">
        <div class="title">
          <!--<h3>我的技能</h3-->
        </div>
        <!-- <div class="layui-row">
            <img style="float:left" src="https://manyushuai.site/static/imges/skill.png">
        </div -->
      </div>
      <div class="item contact">
        <!-- div class="title">
          <h3>联系方式</h3>
        </di> -->
      <!--  <div class="cont">
          <img style="margin:0" src="https://manyushuai.site/static/imges/wx_qrcode.jpg">
          <div class="text">
            <p class="WeChat" style="margin-top:-10px">Email：ysman@manyushuai.site<span></span></p>
            <p class="WeChat">site：https://manyushuai.site<span></span></p>
            <p class="WeChat">公众号：满宇帅的分享<span></span></p>
          </div>
        </div> -->
      </div>
    </div>
  </div>
  <foot></foot>
</div>
</template>
<script>
import headd from '@/components/head'
import foot from '@/components/footer'

export default {
  data () {
    return {
      data: '',
    }
  },

  components: {
    headd, foot
  }
}
</script>

