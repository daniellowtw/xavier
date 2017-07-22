<template>
  <section class="section">
    <news-bar :searchMode="searchMode" @changeMode="onChangeMode"></news-bar>
    <div class="columns">
      <news-menu class="column is-3" @toggle-source="chooseNewsSource" :selectedSources="selectedSources"></news-menu>
      <div class="column is-9">
        <news-item v-for="newsItem in news" :key="newsItem.Id" :news="newsItem"></news-item>
      </div>
    </div>
  </section>
</template>

<script>
import NewsItem from './NewsItem.vue'
import NewsBar from './NewsBar.vue'
import NewsMenu from './NewsMenu.vue'
import { mapState } from 'vuex'

export default {
  name: 'news',
  components: [
    NewsItem,
    NewsBar,
    NewsMenu
  ],
  computed: mapState({
    sources: 'sources',
    news: 'news'
  }),
  methods: {
    onChangeMode(mode) {
      this.searchMode = mode
      this.loadNews()
    },
    loadNews() {
      this.$store.dispatch('loadNews', {
        selectedSources: this.selectedSources,
        searchMode: this.searchMode
      })
    },
    // TODO: Dead code
    toggleSource(id, index) {
      let newVal = (this.selectedSources[index] === 0) ? id : 0
      this.selectedSources = this.selectedSources.slice(0, index).concat([newVal]).concat(this.selectedSources.slice(index + 1))
      this.loadNews()
    },
    chooseNewsSource(id, index) {
      this.selectedSources = (new Array(this.sources.length)).fill(0)
      this.selectedSources[index] = id
      this.loadNews()
    }

  },
  watch: {
    sources() {
      this.selectedSources = (new Array(this.sources.length)).fill(0)
    }
  },
  data() {
    return {
      searchMode: 'unread',
      selectedSources: [],
    }
  },
  created() {
    this.loadNews()
  },
}
</script>
