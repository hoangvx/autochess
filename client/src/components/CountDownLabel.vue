<template>
  <div class="ht-countDownWrapper">
    <span class="ht-countDownNumber">
      {{ count }}
    </span>
    <span class="ht-countDownLabel">
      {{ type }}
    </span>
  </div>
</template>

<script>
export default {
  name: 'CountDownLabel',
  props: {
    count: Number,
    type: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  div.ht-countDownWrapper {
    margin: 0 5px;
    display: inline-block;
  }

  div.ht-countDownWrapper span {
    display: block;
    color: white;
  }

  span.ht-countDownNumber {
    background:rgba(40,50,61,.5);
    height: 120px;
    width: 120px;
    text-align: center;
    vertical-align: middle;
    line-height: 120px;
  }

  span.ht-countDownLabel {
    font-size: .25em;
    text-align: left;
  }

  /* IPX Portrait and Landscape */
  @media only screen 
    and (min-device-width: 375px) 
    and (max-device-width: 812px) 
    and (-webkit-min-device-pixel-ratio: 3) { 
    span.ht-countDownNumber {
      height: 60px;
      width: 60px;
      line-height: 60px;
      font-size: 25px;
    }
    span.ht-countDownLabel {
      font-size: 8px;
    }
  }
</style>
