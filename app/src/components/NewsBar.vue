<template>
  <!-- Main container -->
  <nav class="level">
    <!-- Left side -->
    <div class="level-left">
    </div>
  
    <!-- Right side -->
    <div class="level-right">
    </div>
  </nav>
</template>

<script>
import Vue from 'vue'
export default Vue.component('news-bar', {
  props: ['isDebug'],
  data() {
    return {
      isLoading: false
    }
  },
})
</script>
