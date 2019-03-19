<template>
  <el-container id="ht-homepage">
    
  </el-container>
</template>

<style>
html, body,
#app,
#app section,
#app section main,
#ht-homepage {
  height: 100%;
}

#app section main {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  padding: 0;
  margin: 0;
  z-index: -1;
}

#ht-homepage {
  background-image: url('https://images.unsplash.com/photo-1515934751635-c81c6bc9a2d8?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=3300&q=80');
  background-repeat: no-repeat;
  background-position: center;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;

  padding-top: 60px;
}

#ht-homepage:before {
  content: '';
	position: absolute;
	top: 0;
	right: 0;
	bottom: 0;
	left: 0;
	background-image: linear-gradient(to bottom right,#002f4b,#dc4225);
	opacity: .5;
}
</style>

<script>
export default {
  name: 'home',
  data: function() {
    return {
      message: 'Hello Vue!'
    }
  }
}
</script>

