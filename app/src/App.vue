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
import request from 'superagent'
import { mapState } from 'vuex'
var __API__ = '/api'

export default {
  name: 'app',
  components: {
    Feed,
    NavBar,
    News
  },
  computed: mapState({
    sourcesList: 'sourcesList',
  }),
  data() {
    return {
      mode: 'feed'
    }
  },
  methods: {
    changeMode(mode) {
      this.mode = mode
    },
    loadSources() {
      request.get(`${__API__}/feeds`)
        .end((err, res) => {
          if (err) {
            console.log(err)
            return
          }
          this.$store.commit('updateSources', JSON.parse(res.text))
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
