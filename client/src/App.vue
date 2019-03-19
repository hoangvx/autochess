<template>
  <div id="app">
    <el-container>
      <el-header>
        <el-menu mode="horizontal" id="ht-top_menu">
          <el-menu-item index="1">
            <router-link to="/">Home</router-link>
          </el-menu-item>
          <el-menu-item index="2">
            <router-link to="/about">About</router-link>
          </el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <router-view/>
      </el-main>
    </el-container>
  </div>
</template>

<style>
#ht-top_menu {
  background: transparent;
  border: none;
}

#ht-top_menu .el-menu-item {
  float: right;
}
</style>
