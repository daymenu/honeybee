<template>
  <div id="app">
    <el-button @click="send">发送</el-button>
  </div>
</template>

<script>

import {Connect, SendMsg} from './api'

export default {
  name: 'app',
  components: {
    
  },
  created() {
    console.log(process.env)
    Connect()
  },
  methods: {
    send: function () {
      SendMsg("hello")
    }
  }
}
</script>

<style>

</style>
