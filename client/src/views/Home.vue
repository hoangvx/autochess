<template>
  <el-container id="ht-homepage" >
    <el-row type="flex" align="middle">
      <el-col id="ht-countdown">
        <el-row style="height: 200px"/>
        <el-row class="hidden-xs-only">
          <el-col>
            <CountDownLabel v-if="years" :count="years" type="YEARS" />
            <CountDownLabel v-if="months" :count="months" type="MONTHS" />
            <CountDownLabel :count="days" type="DAYS" />
            <CountDownLabel :count="hours" type="HOURS" />
            <CountDownLabel :count="minutes" type="MINUTES" />
            <CountDownLabel :count="seconds" type="SECONDS" />
          </el-col>
        </el-row>
        <ImageSlide />
      </el-col>
    </el-row>
  </el-container>
</template>

// <el-menu-item index="1">
//             <router-link to="/">Home</router-link>
//           </el-menu-item>
//           <el-menu-item index="2">
//             <router-link to="/about">About</router-link>
//           </el-menu-item>

<style>
  html, body,
  #app,
  #app section,
  #app section main,
  #ht-homepage {
    height: 100%;
  }

  #app section main {
    padding: 0;
    margin: 0;
  }

  #ht-homepage {
    background-image: url('https://images.unsplash.com/photo-1515934751635-c81c6bc9a2d8?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=3300&q=80');
    background-repeat: no-repeat;
    background-position: center;
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
  }

  #ht-homepage:before {
    content: ' ';
    position: absolute;
    width: 100%;
    height: 100%;
    background-image: linear-gradient(to bottom right,#002f4b,#dc4225);
    opacity: .5;
  }

  #ht-homepage .el-row {
    width: 100%;
  }

  #ht-countdown {
    text-align: center;
    font-size: 4em;
  }
</style>

<script>
import moment from 'moment'
// components
import CountDownLabel from '../components/CountDownLabel'
import ImageSlide from '../components/ImageSlide'

export default {
  name: 'home',
  data: function() {
    return {
      startDate: moment("2019-07-26"),
      duration: null
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
    })
  },
  components: {
    CountDownLabel,
    ImageSlide
  }
}
</script>

