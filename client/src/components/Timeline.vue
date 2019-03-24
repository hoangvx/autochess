<template>
  <div id="ht-timeline">
    <ul>
      <li v-for="story in stories" :key="story.id">
        <div class="ht-timeline_badge">
          <font-awesome-icon icon="heart" />
        </div>
        <div class="ht-timeline_content">
          <div class="ht-timeline_heading">
            <h2>{{ story.date }}</h2>
          </div>
          <div class="ht-timeline_body">
            <img :src="story.image">
            <h2>{{ story.title }}</h2>
            <p>{{ story.content }}</p>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
import { library } from '@fortawesome/fontawesome-svg-core'
import { faHeart } from '@fortawesome/free-solid-svg-icons'
// components
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
library.add(faHeart)

import stories from '../data/stories'

export default {
  name: 'Timeline',
  data: function() {
    return {
      stories
    }
  },
  mounted: function() {
    require('../assets/images/trang.jpg')
    this.$nextTick(function() {
    })
  },
  components: {
    FontAwesomeIcon
  }
}
</script>

<style>
  #ht-timeline ul {
    list-style: none;
    padding: 20px 0 20px;
    position: relative;
    margin-top: 0;
    margin-bottom: 10px;
  }

  #ht-timeline ul::before {
    top: 0;
    bottom: 0;
    position: absolute;
    content: " ";
    width: 3px;
    background-color: rgba(0, 0, 0, 0.09);
    left: 50%;
    margin-left: -1.5px;
  }

  #ht-timeline ul li:first-child {
    margin-top: -27px;
  }

  #ht-timeline ul li {
    margin-top: -25%;
    position: relative;
    margin-bottom: 20px;
  }

  #ht-timeline ul li::before,
  #ht-timeline ul li::after {
    content: " ";
    display: table;
  }

  #ht-timeline ul li::after {
    clear: both;
  }

  #ht-timeline ul li > .ht-timeline_badge {
    position: absolute;
    top: 1%;
    left: 50%;
    color: #fff;
    width: 50px;
    height: 50px;
    line-height: 50px;
    font-size: 1.4em;
    text-align: center;
    margin-left: -25px;
    background-color: #f9667e;
    z-index: 100;
    border-top-right-radius: 50%;
    border-top-left-radius: 50%;
    border-bottom-right-radius: 50%;
    border-bottom-left-radius: 50%;
  }

  #ht-timeline ul li > .ht-timeline_content {
    top: -10px;
    width: calc(47% - 60px);
    float: left;
    border-radius: 2px;
    position: relative;
    background: transparent;
    border: solid 3px #69655a;
    border-radius: 5px;
    padding: 20px;
  }

  #ht-timeline ul li:nth-child(even) > .ht-timeline_content {
    float: right;
  }

  .ht-timeline_content::before{
    position: absolute;
    top: 24px;
    right: -18px;
    display: inline-block;
    border-top: 17px solid transparent;
    border-left: 17px solid#69655a;
    border-right: 0 solid#69655a;
    border-bottom: 18px solid transparent;
    content: " ";
  }

  .ht-timeline_content::after {
    position: absolute;
    top: 27px;
    right: -14px;
    display: inline-block;
    border-top: 14px solid transparent;
    border-left: 14px solid #EEEEEE;
    border-right: 0 solid #EEEEEE;
    border-bottom: 14px solid transparent;
    content: " ";
  }

  #ht-timeline ul li:nth-child(even) > .ht-timeline_content::before {
    border-left-width: 0;
    border-right-width: 17px;
    left: -18px;
    right: auto;
  }

  #ht-timeline ul li:nth-child(even) > .ht-timeline_content::after {
    border-left-width: 0;
    border-right-width: 14px;
    left: -14px;
    right: auto;
  }

  .ht-timeline_heading h2 {
    text-align: right;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    margin-bottom: 10px;
  }

  #ht-timeline ul li:nth-child(even) .ht-timeline_heading h2 {
    text-align: left;
  }

  .ht-timeline_body img {
    display: block;
    max-width: 100%;
    height: auto;
    margin: 20px auto;
    border-radius: 5px;
  }

  .ht-timeline_body h2 {
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    text-align: center;
  }

  .ht-timeline_body p {
    text-align: center;
  }

   /* IPX Portrait and Landscape */
  @media only screen 
    and (min-device-width: 375px) 
    and (max-device-width: 812px) 
    and (-webkit-min-device-pixel-ratio: 3) { 
    #ht-timeline ul::before {
      left: 30px;
    }

    #ht-timeline ul li {
      margin-top: 0px;
    }

    #ht-timeline ul li > .ht-timeline_badge {
      left: 5px;
      margin-left: 0;
      top: 16px;
    }

    #ht-timeline ul li > .ht-timeline_content {
      width: -webkit-calc(100% - 120px);
      width: calc(100% - 120px);
      float: right;
      top: 0;
    }

    .ht-timeline_content::before{
      border-left-width: 0;
      border-right-width: 17px;
      left: -18px;
      right: auto;
    }

    .ht-timeline_content::after {
      border-left-width: 0;
      border-right-width: 14px;
      left: -14px;
      right: auto;
    }
  }
</style>