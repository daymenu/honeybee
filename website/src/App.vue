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
    Connect()
  },
  methods: {
    send: function () {
      console.log('ddd')
      SendMsg("hello")
    }
  }
}
</script>

<style>

</style>
