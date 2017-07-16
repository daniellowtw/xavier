<template>
  <nav class="nav has-shadow">
    <div class="nav-left">
      <a v-on:click='changeMode("feed")' class='navbar-item is-tab'>Feed</a>
      <a v-on:click='changeMode("news")' class='navbar-item is-tab'>News</a>
    </div>
    <div class="nav-right">
      <a v-on:click='toggleDebug()' class='navbar-item is-active'>Debug: {{isDebug}}</a>
    </div>
  </nav>
</template>

<script>
export default {
  props: ['mode', 'isDebug'],
  name: 'nav-bar',
  methods: {
    changeMode: function (mode) {
      this.$emit('change-mode', mode)
    },
    toggleDebug() {
      this.$emit('toggleDebug')
    }
  },
}
</script>
