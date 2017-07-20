<template>
  <section class="section">
    <news-bar></news-bar>
    <div class="columns">
      <news-menu class="column is-3" @toggle-source="toggleSource" :sources="sources"></news-menu>
      <div class="column is-9">
        <news-item v-for="newsItem in news" :key="newsItem.Id" :news="newsItem" :isDebug="isDebug" :fav="favIcon[newsItem.FeedId]" :currentNewsId="currentNewsId" @read="markRead"></news-item>
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
export default {
  props: ['isDebug', 'sources'],
  name: 'news',
  components: [
    NewsItem,
    NewsBar,
    NewsMenu
  ],
  methods: {
    onChangePage(page) {
      this.page = page
    },
    markRead(news) {
      if (news.read) {
        this.currentNewsId = news.Id
      }
      request.post(`${__API__}/feeds/${news.FeedId}/news/${news.Id}`)
        .send('action=read') // sending string automatically makes it form URL encoded
        .end((err, res) => {
          if (err) {
            console.log(err)
            return
          }
          news.read = true
          this.currentNewsId = news.Id
        })
    },
    loadNews() {
      request.get(`${__API__}/news`)
        .end((err, res) => {
          if (err) {
            console.log(err)
            return
          }
          this.allNews = JSON.parse(res.text)
          this.news = this.allNews
          this.total = this.news.length
        })
    },
    toggleSource(id) {
      this.filteredSource = id
    }

  },
  watch: {
    filteredSource() {
      this.news = this.allNews.filter(x => x.FeedId === this.filteredSource)
    }
  },
  data() {
    return {
      page: 1,
      total: 100,
      itemsPerPage: 10,
      allNews: [],
      favIcon: {},
      filteredSource: null,
      news: [],
      currentNewsId: 0,
    }
  },
  created() {
    this.loadNews()
  },
  updated() {
    this.sources.forEach(el => {
      this.favIcon[el.Id] = el.FavIcon
    }, this)
  }
}
</script>
