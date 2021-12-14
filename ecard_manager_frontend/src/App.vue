<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
  export default {
    name: 'App',
  }
</script>

<style>
  @import "./assets/css/main.css";
  @import "./assets/css/color-dark.css"; /*深色主题*/

  html, body {
    height: 100%;
  }

  /* 解决表头不对齐问题，样式 */
  body .el-table th.gutter {
    display: table-cell !important;
  }

  body .el-table colgroup.gutter {
    display: table-cell !important;
  }
  ::-webkit-scrollbar {
    width: 8px;
    height: 8px;
    background-color: #fff;
    border-radius: 5px;
  }
  ::-webkit-scrollbar-thumb {
    border-radius: 5px;
    -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
    background-color: rgba(0, 0, 0, 0.1);
  }
</style>
