<template>
  <div id="app" class="container">
    <nav-bar @change-mode="changeMode"></nav-bar>
    <feed :mode="mode" v-show="mode == 'feed'"></feed>
    <news :mode="mode" v-show="mode == 'news'"></news>
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
  data() {
    return {
      mode: 'feed'
    }
  },
  methods: {
    changeMode(mode) {
      this.mode = mode
    },
  },
  created() {
    this.$store.dispatch('loadSources')
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
