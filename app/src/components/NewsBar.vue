<template>
  <!-- Main container -->
  <nav class="level">
    <!-- Left side -->
    <div class="level-left">
      <div class="field has-addons">
        <p class="control">
          <a class="button" v-on:click="changeMode('saved')" v-bind:class="{'is-primary':this.searchMode === 'saved'}">
            <span>Saved</span>
          </a>
        </p>
        <p class="control">
          <a class="button" v-on:click="changeMode('unread')" v-bind:class="{'is-primary':this.searchMode === 'unread'}">
            <span>Unread</span>
          </a>
        </p>
        <p class="control">
          <a class="button" v-on:click="changeMode('all')" v-bind:class="{'is-primary':this.searchMode === 'all'}">
            <span>All</span>
          </a>
        </p>
      </div>
    </div>
  
    <!-- Right side -->
    <div class="level-right">
    </div>
  </nav>
</template>

<script>
import Vue from 'vue'
import { mapState } from 'vuex'
export default Vue.component('news-bar', {
  props: ['searchMode'],
  computed: mapState({
    isDebug: 'isDebug'
  }),
  data() {
    return {
      isLoading: false
    }
  },
  methods: {
    changeMode(v) {
      this.$emit('changeMode', v)
    }
  }
})
</script>
