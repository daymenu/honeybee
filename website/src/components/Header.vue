<template>
  <div class="hb-header">
    <h1>{{ header }}</h1>
  </div>
</template>

<script>
export default {
  name: 'Header',
  props: {
    header: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>