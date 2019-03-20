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
    font-size: .3em;
    text-align: left;
  }
</style>
