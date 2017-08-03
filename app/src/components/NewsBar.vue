<template>
  <!-- Main container -->
  <nav id="news-bar" class="level">
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
      <div class="level-item">
        <a class="button is-success" v-on:click="classifyAllAsDislike()">Mark all unclassified as dislike</a>
      </div>
      <div class="level-item">
        <a class="button is-success" v-on:click="markAllAsRead()">Mark all as read</a>
      </div>
    </div>
  </nav>
</template>
<style>
#news-bar {
  padding-top: 10px;
}
</style>

<script>
import Vue from 'vue'
import { mapState } from 'vuex'
export default Vue.component('news-bar', {
  props: ['searchMode'],
  computed: mapState({
    isDebug: 'isDebug',
    news: 'news',
  }),
  data() {
    return {
      isLoading: false
    }
  },
  methods: {
    changeMode(v) {
      this.$emit('changeMode', v)
    },
    classifyAllAsDislike() {
      let d = this.$store.dispatch
      this.news.filter(x => x.Classification === 0).forEach(x => d('classify', { newsId: x.Id, feedId: x.FeedId, classification: 2 }))
    },
    markAllAsRead() {
      this.news.forEach(x => this.$store.dispatch('markRead', { newsId: x.Id, feedId: x.FeedId }))
    }
  }
})
</script>
