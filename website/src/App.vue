<template>
  <div id="app">
    <Header header="Honeybee"></Header>
    <Chat-History :chats="chats"></Chat-History>
    <el-button @click="send">发送</el-button>
  </div>
</template>

<script>
import Header from './components/Header'
import ChatHistory from './components/ChatHistory'

import {Connect, SendMsg} from './api'

export default {
  name: 'app',
  components: {
    Header,
    ChatHistory
  },
  created() {
    console.log(process.env)
    Connect()
  },
  data(){
    return {
      chats:[
        {id:'1', msg:'ha'},
        {id:'2', msg:'hi'}
      ]
    }
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
