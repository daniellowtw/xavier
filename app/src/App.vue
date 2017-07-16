<template>
  <div id="app" class="container">
    <nav-bar @change-mode="changeMode"></nav-bar>
    <feed :mode="mode" v-show="mode == 'feed'"></feed>
  </div>
</template>

<script>
import Feed from './components/Feed'
import NavBar from './components/NavBar'

export default {
  name: 'app',
  components: {
    Feed,
    NavBar
  },
  created() {
  },
  data() {
    return {
      mode: 'feed'
    }
  },
  methods: {
    changeMode(mode) {
      this.mode = mode
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
