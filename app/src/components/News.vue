<template>
  <section class="section">
    <news-bar></news-bar>
    <div class="columns">
      <div class="column is-12">
        <news-item v-for="feed in news" :key="feed.Id" :news="feed" :isDebug="isDebug" @read="markRead"></news-item>
        <pagination :total="total" :page="page" :items-per-page="itemsPerPage" @change-page="onChangePage"></pagination>
      </div>
    </div>
  </section>
</template>

<script>
var __API__ = '/api'
import Pagination from './Pagination.vue'
import NewsItem from './NewsItem.vue'
import NewsBar from './NewsBar.vue'
import request from 'superagent'
export default {
  props: ['isDebug'],
  name: 'news',
  components: [
    Pagination,
    NewsItem,
    NewsBar
  ],
  methods: {
    onChangePage(page) {
      this.page = page
    },
    markRead(news) {
      request.post(`${__API__}/feeds/${news.FeedId}/news/${news.Id}`)
        .send('action=read') // sending string automatically makes it form URL encoded
        .end((err, res) => {
          if (err) {
            console.log(err)
            return
          }
          news.read = true
          console.log(res)
        })
    },
    loadNews() {
      this.$http.get(`${__API__}/news`).then(
        response => {
          this.news = JSON.parse(response.body)
          this.total = this.news.length
        },
        error => {
          console.log(error)
        }
      ).finally(() => {
      })
    }
  },
  data() {
    return {
      page: 1,
      total: 100,
      itemsPerPage: 10,
      news: []
    }
  },
  created() {
    this.loadNews()
  },
}
</script>
