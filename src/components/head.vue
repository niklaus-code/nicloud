<template>
<div class="header">
    <div class="menu-btn">
      <div class="menu"></div>
    </div>
    <h1 class="logo">
      <a href="index.html">
        <span>MYBLOG</span>
        <img src="../../static/imges/logo.png">
      </a>
    </h1>
    <div class="nav">
      <!-- <router-link :to="{path: '/'}">文章</router-link>
      <router-link :to="{path: '/whisper'}">微语</router-link>
      <router-link :to="{path: '/read'}">书</router-link>
      <a>关于</a>     <ul class="layui-nav header-down-nav">-->

    <!--
      <a class="layui-nav-item"><router-link :to="{path: '/'}">文章</route-link></a>
      <a class="layui-nav-item"><router-link :to="{path: '/whisper'}">微语</route-link></a>
      <a class="layui-nav-item"><router-link :to="{path: '/read'}">读书</route-link></a>
      <a class="layui-nav-item"><router-link :to="{path: '/'}">关于</route-link></a>
-->
      <router-link :to="{path: '/whisper'}">微语</router-link>
      <router-link :to="{path: '/blog'}">关于技术</router-link>
      <router-link :to="{path: '/read'}">读书&旅行</router-link>
      <router-link :to="{path: '/about'}">关于</router-link>
    </div> 

        <ul class="layui-nav header-down-nav">
       <li class="layui-nav-item"><a>文章</a></li>
       <li class="layui-nav-item"><a>微语</a></li>
       <li class="layui-nav-item"><a>读书</a></li>
       <li class="layui-nav-item"><a class="active">关于</a></li>
    </ul>

    <p class="welcome-text">
      欢迎来到<span class="name">&nbsp <strong><router-link :to="{path: '/login'}">NiKlaus</router-link></strong> &nbsp</span>的博客~
    </p>

  </div>

</template>
