<template>
  <section class="section">
    <news-bar :searchMode="searchMode" @changeMode="onChangeMode"></news-bar>
    <div class="columns">
      <news-menu class="column is-3" @toggle-source="chooseNewsSource" :selectedSources="selectedSources"></news-menu>
      <div class="column is-9">
        <news-item v-for="newsItem in news" :key="newsItem.Id" :news="newsItem" :currentNewsId="currentNewsId" @read="markRead"></news-item>
      </div>
    </div>
  </section>
</template>

<script>
var __API__ = '/api'
import NewsItem from './NewsItem.vue'
import NewsBar from './NewsBar.vue'
import NewsMenu from './NewsMenu.vue'
import request from 'superagent'
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
  }),
  methods: {
    onChangeMode(mode) {
      this.searchMode = mode
      this.loadNews()
    },
    onChangePage(page) {
      this.page = page
    },
    markRead(news) {
      if (news.Read) {
        this.currentNewsId = news.Id
      } else {
        this.$store.commit('readNews', news.FeedId)
        // this.sources.filter(x => x.Id === news.FeedId)[0].UnreadCount--
      }
      request.post(`${__API__}/feeds/${news.FeedId}/news/${news.Id}`)
        .send('action=read') // sending string automatically makes it form URL encoded
        .end((err, res) => {
          if (err) {
            console.log(err)
            return
          }
          news.Read = true
          this.currentNewsId = news.Id
        })
    },
    loadNews() {
      let selectedIds = this.selectedSources.filter(x => x !== 0)
      let r = request.post(`${__API__}/news`)
      r = r.send('limit=100')
      if (this.searchMode === 'unread') {
        r = r.send('search=unread')
      }
      if (selectedIds.length !== 0) {
        r = r.send(`ids=${selectedIds.join(',')}`)
      }
      r.end((err, res) => {
        if (err) {
          console.log(err)
          return
        }
        this.news = JSON.parse(res.text)
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
      news: [],
      currentNewsId: 0,
      searchMode: 'unread',
      selectedSources: [],
    }
  },
  created() {
    this.loadNews()
  },
}
</script>
