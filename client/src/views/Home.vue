<template>
  <div id="ht-wrapper">
    <!-- <transition name="slide">
      <div id="ht-spin" v-if="loading">
        <el-row type="flex" justify="center" align="middle">
          <el-row>
            <Spin />
          </el-row>
        </el-row>
      </div>
    </transition> -->
    <div id="ht-top">
      <div id="ht-cover"></div>
      <div id="ht-cover_info">
        <el-container>
          <el-header height="80px">
            <el-menu mode="horizontal">
              <el-menu-item index="1">About</el-menu-item>
              <el-menu-item index="1">Galery</el-menu-item>
            </el-menu>
            <img src="../assets/images/logo2.png"/>
          </el-header>
          <el-main>
            <el-row type="flex" justify="center" align="bottom">
              <div id="ht-cover_name">
                <h1>Xuan Hoang</h1>
                <h1>&</h1>
                <h1>Khanh Trang</h1>
              </div>
            </el-row>
          </el-main>
          <el-footer height="150px">
            <el-row type="flex" align="middle">
              <el-col id="ht-countdown">
                <CountDownLabel v-if="years" :count="years" type="YEARS" />
                <CountDownLabel v-if="months" :count="months" type="MONTHS" />
                <CountDownLabel :count="days" type="DAYS" />
                <CountDownLabel :count="hours" type="HOURS" />
                <CountDownLabel :count="minutes" type="MINUTES" />
                <CountDownLabel :count="seconds" type="SECONDS" />
              </el-col>
            </el-row>
          </el-footer>
        </el-container>
      </div>
      <!-- <el-row type="flex" align="bottom" id="ht-info">
        <el-col>
          <el-row type="flex" align="middle">
            <el-col id="ht-countdown">
              <CountDownLabel v-if="years" :count="years" type="YEARS" />
              <CountDownLabel v-if="months" :count="months" type="MONTHS" />
              <CountDownLabel :count="days" type="DAYS" />
              <CountDownLabel :count="hours" type="HOURS" />
              <CountDownLabel :count="minutes" type="MINUTES" />
              <CountDownLabel :count="seconds" type="SECONDS" />
            </el-col>
          </el-row>
        </el-col>
      </el-row> -->
    </div>
    <div id="ht-content">
      <el-container id="ht-sec1" direction="vertical">
        <el-row type="flex" justify="center">
          <el-row>
            <el-col class="ht-title">
              <h2>The Couple</h2>
              <img src="../assets/pattern/6.png" />
            </el-col>
          </el-row>
        </el-row>
        <el-row :gutter="10" class="ht-bride">
          <el-col :xs="24" :lg="12" :md="12" :sm="12">
            <div class="ht-about">
              <h2>Bride</h2>
              <h4>Trang</h4>
              <p>It is a long established fact that a reader will 
                be distracted by the readable content of a page when looking at its layout. 
                The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters,
                as opposed to using 'Content here, content here', making it look like readable English</p>
            </div>
          </el-col>
          <el-col :xs="24" :lg="12" :md="12" :sm="12">
            <img class="ht-avatar" src="../assets/images/trang.jpg">
          </el-col>
        </el-row>
        <el-row  type="flex" justify="center">
          <div class="ht-heart">
            <font-awesome-icon icon="heartbeat" size="4x"/>
          </div>
        </el-row>
        <el-row :gutter="10" class="ht-groom">
          <el-col :xs="24" :lg="12" :md="12" :sm="12">
            <img class="ht-avatar" src="../assets/images/hoang.jpg">
          </el-col>
          <el-col :xs="24" :lg="12" :md="12" :sm="12">
            <div class="ht-about">
              <h2>Groom</h2>
              <h4>Hoang</h4>
              <p>It is a long established fact that a reader will 
                be distracted by the readable content of a page when looking at its layout. 
                The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters,
                as opposed to using 'Content here, content here', making it look like readable English</p>
            </div>
          </el-col>
        </el-row>
      </el-container>
      <el-container id="ht-stories" direction="vertical">
        <el-row type="flex" justify="center">
          <el-row>
            <el-col class="ht-title">
              <h2>The Stories</h2>
              <img src="../assets/pattern/6.png" />
            </el-col>
          </el-row>
        </el-row>
        <el-row type="flex" justify="center">
          <p class="ht-story">
            <em>
              It is a long established fact that a reader will 
              be distracted by the readable content of a page when looking at its layout. 
              The point of
            </em>
          </p>
        </el-row>
        <el-row>
          <Timeline />
        </el-row>
      </el-container>
    </div>
  </div>
</template>

<style>
  @import url('https://fonts.googleapis.com/css?family=Alex+Brush');

  h1, h2, h3, h4, h5 {
    font-family: "Alex Brush", cursive;
    margin: 0;
  }

  html, body, 
  #ht-wrapper, #ht-top {
    padding: 0;
    margin: 0;
    height: 100%;
  }

  #ht-cover_info {
    height: 100%;
    width: 100%;
    top: 0;
    left: 0;
    position: absolute;
  }

  #ht-spin {
    background: white;
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    z-index: 2;
  }

  #ht-spin,
  #ht-spin > .el-row {
    height: 100%;
  }

  .slide-fade-enter-active {
    transition: all .3s ease;
  }
  .slide-fade-leave-active {
    transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
  }
  .slide-fade-enter, .slide-fade-leave-to
  /* Trước 2.1.8 thì dùng .slide-fade-leave-active */ {
    transform: translateX(10px);
    opacity: 0;
  }

  #ht-cover {
    height: 100%;
    background-image: url(../assets/images/cover.jpg);
    background-repeat: no-repeat;
    background-position: center;
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
    position: relative;
  }

  #ht-cover:before {
    content: ' ';
    position: absolute;
    width: 100%;
    height: 100%;
    background-image: linear-gradient(to bottom right,#002f4b,black);
    opacity: .5;
  }

  #ht-info {
    height: 100%;
  }

  #ht-countdown {
    text-align: center;
  }

  #ht-content {
    color: #69655a;
    background: url(../assets/pattern/4.png);
    padding: 0 50px;
  }

  div.ht-title {
    text-align: center;
    padding: 50px 0 10px 0;
  }

  div.ht-title h2 {
    font-size: 60px;
  }

  div.ht-title img {
    text-align: center;
    object-fit: cover;
    width: 80%;
  }

  div.ht-section_content {
    width: 100%;
  }

  img.ht-avatar {
    height: 220px;
    width: 220px;
    object-fit: cover;
    border-radius: 10px;
    border: 5px solid #f8f8f8;
    box-shadow: 0 2px 24px rgba(0, 0, 0, 0.2);
    -webkit-box-shadow: 0 2px 24px rgba(0, 0, 0, 0.2);
  }

  div.ht-bride img {
    margin-left: 20px;
  }

  div.ht-groom {
    text-align: right;
  }

  div.ht-groom img {
    margin-right: 20px;
  }

  div.ht-bride div.ht-about {
    margin-right: 20px;
    padding-left: 20px;
    text-align: right;
  }

  div.ht-groom div.ht-about {
    margin-left: 20px;
    padding-right: 20px;
    text-align: left;
  }

  div.ht-about h2 {
    font-size: 40px;
  }

  div.ht-about h4 {
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    font-size: 20px;
  }

  div.ht-heart {
    position: relative;
    margin: 10px 0;
  }

  div.ht-heart::before,
  div.ht-heart::after {
    content: "";
    height: 1px;
    position: absolute;
    top: 45%;
    width: 150px;
    background-color: #69655a;
  }
  div.ht-heart::before {
    left: -140px;
  }

  div.ht-heart::after {
    right: -140px;
  }

  p.ht-story {
    text-align: center;
    max-width: 400px;
    margin-bottom: 60px;
  }

  /* Top Section */
  div#ht-cover_info > section {
    height: 100%;
  }

  header {
    position: relative;
  }

  header img {
    height: 200%;
    object-fit: cover;
    padding: 0;
  }

  header ul.el-menu.el-menu--horizontal {
    border-color: #edcc59;
  }

  header ul.el-menu {
    position: absolute;
    width: 100%;
    background: transparent;
    height: 80px;
    width: calc(100% - 230px);
    margin: 0;
    right: 50px;
  }

  header ul.el-menu > .el-menu-item {
    height: 80px;color: #edcc59;
    line-height: 80px;
    float: right;
  }  

  main {
    height: calc(100% - 150px - 80px);
  }

  main > .el-row {
    height: 100%;
  }

  #ht-cover_name {
    text-align: center;
    font-size: 3em;
    color: white;
  }

  /* IPX Portrait and Landscape */
  @media only screen 
    and (min-device-width: 375px) 
    and (max-device-width: 812px) 
    and (-webkit-min-device-pixel-ratio: 3) { 
    #ht-content {
      padding: 0 10px;
    }
    #ht-cover_name {
    font-size: 2em;
  }
  }
</style>

<script>
import moment from 'moment'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faHeartbeat } from '@fortawesome/free-solid-svg-icons'
// components
import CountDownLabel from '../components/CountDownLabel'
import Timeline from '../components/Timeline'
// import Spin from '../components/Spin'

import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
library.add(faHeartbeat)

export default {
  name: 'home',
  data: function() {
    return {
      startDate: moment("2019-07-26"),
      duration: null,
      loading: true
    }
  },
  computed: {
    durationNotNull: function() { return this.duration != null },
    years: function() { 
      if (this.durationNotNull) {
        return this.duration.years()
      }
      return 0
    },
    months: function() { 
      if (this.durationNotNull) {
        return this.duration.months()
      }
      return 0
    },
    days: function() { 
      if (this.durationNotNull) {
        return this.duration.days()
      }
      return 0
    },
    hours: function() { 
      if (this.durationNotNull) {
        return this.duration.hours()
      }
      return 0
    },
    minutes: function() { 
      if (this.durationNotNull) {
        return this.duration.minutes()
      }
      return 0
    },
    seconds: function() { 
      if (this.durationNotNull) {
        return this.duration.seconds()
      }
      return 0
    },
  },
  mounted: function() {
    this.$nextTick(function () {
      let self = this
      let today = moment()
      let diff = self.startDate - today
      self.duration = moment.duration(diff, 'milliseconds')
      setInterval(function() {
        self.duration = moment.duration(self.duration.as('milliseconds') - 1000, 'milliseconds')
      }, 1000)

      setTimeout(function() {
        self.loading = false
      }, 5000);
    })
  },
  components: {
    CountDownLabel,
    // ImageSlide,
    FontAwesomeIcon,
    Timeline,
    // Spin
  }
}
</script>

