<template>
  <div id="app" class="container">
    <nav-bar :isDebug="isDebug" @change-mode="changeMode" @toggleDebug="toggleDebug"></nav-bar>
    <feed :isDebug="isDebug" :mode="mode" :feeds="sourcesList" v-show="mode == 'feed'"></feed>
    <news :isDebug="isDebug" :mode="mode" :sources="sourcesList" v-show="mode == 'news'"></news>
  </div>
</template>

<script>
import Feed from './components/Feed'
import News from './components/News'
import NavBar from './components/NavBar'
import request from 'superagent'
var __API__ = '/api'

export default {
  name: 'app',
  components: {
    Feed,
    NavBar,
    News
  },
  data() {
    return {
      mode: 'feed',
      isDebug: false,
      sourcesList: []
    }
  },
  methods: {
    changeMode(mode) {
      this.mode = mode
    },
    toggleDebug(mode) {
      this.isDebug = !this.isDebug
    },
    loadSources() {
      request.get(`${__API__}/feeds`)
        .end((err, res) => {
          if (err) {
            console.log(err)
            return
          }
          this.sourcesList = JSON.parse(res.text)
        })
    }
  },
  created() {
    this.loadSources()
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
