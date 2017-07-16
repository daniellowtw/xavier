<template>
  <!-- Main container -->
  <nav class="level">
    <!-- Left side -->
    <div class="level-left">
    </div>
  
    <!-- Right side -->
    <div class="level-right">
      <p class="level-item">
        <a class="button is-success" v-on:click="refreshFeed" :disabled="isLoading">Refresh</a>
      </p>
      <p class="level-item">
        <a class="button is-success">New</a>
      </p>
    </div>
  </nav>
</template>

<script>
var __API__ = '/api/feeds'
import request from 'superagent'
import swal from 'sweetalert2'
import Vue from 'vue'
export default Vue.component('feed-bar', {
  data() {
    return {
      isLoading: false
    }
  },
  methods: {
    refreshFeed() {
      if (this.isLoading) return
      request.post(__API__).end((err, res) => {
        this.isLoading = false
        if (err) {
          console.log(err)
          return
        }
        swal('Updated feeds', 'Successfully polled items', 'success')
      })
      this.isLoading = true
    }
  },
})
</script>
