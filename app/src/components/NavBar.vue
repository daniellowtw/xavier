<template>
  <nav class="nav has-shadow">
    <div class="nav-left">
      <a v-on:click='changeMode("feed")' class='navbar-item is-tab'>Feed</a>
      <a v-on:click='changeMode("news")' class='navbar-item is-tab'>News</a>
    </div>
    <div class="nav-right">
      <a v-on:click='toggleDebug()' class='navbar-item'>Debug: {{isDebug}}</a>
    </div>
  </nav>
</template>

<script>
import { mapState } from 'vuex'
export default {
  computed: mapState({
    isDebug: 'isDebug',
    mode: 'mode',
  }),
  name: 'nav-bar',
  methods: {
    changeMode: function (mode) {
      this.$store.commit('changeMode', mode)
    },
    toggleDebug() {
      this.$store.commit('toggleDebug')
    }
  },
}
</script>
