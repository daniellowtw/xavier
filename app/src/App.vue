<template>
  <div id="app" class="container">
    <nav-bar :isDebug="isDebug" @change-mode="changeMode" @toggleDebug="toggleDebug"></nav-bar>
    <feed :isDebug="isDebug" :mode="mode" v-show="mode == 'feed'"></feed>
    <news :isDebug="isDebug" :mode="mode" v-show="mode == 'news'"></news>
  </div>
</template>

<script>
import Feed from './components/Feed'
import News from './components/News'
import NavBar from './components/NavBar'

export default {
  name: 'app',
  components: {
    Feed,
    NavBar,
    News
  },
  created() {
  },
  data() {
    return {
      mode: 'feed',
      isDebug: false
    }
  },
  methods: {
    changeMode(mode) {
      this.mode = mode
    },
    toggleDebug(mode) {
      this.isDebug = !this.isDebug
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
