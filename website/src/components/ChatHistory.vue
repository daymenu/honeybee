<template>
  <div class="hb-chat-history">
      <ul>
          <li v-for="chat in chats" :key="chat.id">{{chat.msg}}</li>
      </ul>
  </div>
</template>

<script>
export default {
  name: 'ChatHistory',
  props: {
    chats: Array
  }
}
</script>

<style scoped>
</style>