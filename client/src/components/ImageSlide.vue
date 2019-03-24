<template>
  <el-carousel class="hidden-xs-only" :interval="4000" height="300px" indicator-position="none">
    <el-carousel-item v-for="item in 6" :key="item">
      <img src="https://images.unsplash.com/photo-1494859632785-32abaf5241f4?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1978&q=80" />
    </el-carousel-item>
  </el-carousel>
</template>

<script>
export default {
  name: 'ImageSlide'
}
</script>

<style>
  .el-carousel {
    width: 100%;
  }

  .el-carousel__item .el-carousel__mask {
    background-color: transparent;
  }

  .el-carousel__item img {
    opacity: 0.9;
    object-fit: contain;
  }
</style>