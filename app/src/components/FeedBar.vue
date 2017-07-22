<template>
  <!-- Main container -->
  <nav class="level">
    <!-- Left side -->
    <div class="level-left">
    </div>
  
    <!-- Right side -->
    <div class="level-right">
      <p class="level-item">
        <a class="button is-success" v-on:click.once="refreshFeed">Refresh</a>
      </p>
      <p class="level-item">
        <a class="button is-success">New</a>
      </p>
    </div>
  </nav>
</template>

<script>
import Vue from 'vue'
export default Vue.component('feed-bar', {
  methods: {
    refreshFeed() {
      this.$store.dispatch('refreshFeed')
    }
  },
})
</script>
