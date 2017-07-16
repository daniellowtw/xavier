<template>
  <section class="section">
      <div class="columns">
        <div class="column is-12">
          <table class="table is-narrow is-bordered">
            <thead>
              <th>Name</th>
              <th>Description</th>
              <th>Unread/Total</th>
              <th>LastUpdated</th>
              <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="feed in feeds" :key="feed.Id">
                <td>
                  <a :href="feed.UrlSource">{{feed.Title}}</a>
                </td>
                <td>{{feed.Description}}</td>
                <td>{{feed.UnreadCount}}/{{feed.TotalCount}}</td>
                <td>{{feed.LastUpdated}}</td>
                <td class="is-icon">
                  <a href="#">
                    <i class="fa fa-refresh"></i>
                  </a>
                  <a href="#" @click.prevent="editFeed(feed)">
                    <i class="fa fa-edit"></i>
                  </a>
                  <a href="#" @click.prevent="removeFeed(feed)">
                    <i class="fa fa-trash"></i>
                  </a>
                </td>
              </tr>
            </tbody>
          </table>
          <pagination :total="total" :page="page" :items-per-page="itemsPerPage" @change-page="onChangePage"></pagination>
        </div>
    </div>
  </section>
</template>

<script>
var __API__ = '/api/feeds'
import Pagination from './Pagination.vue'
export default {
  name: 'feed',
  components: [
    Pagination
  ],
  methods: {
    onChangePage(page) {
      this.page = page
    },
    loadFeeds() {
      this.$http.get(__API__).then(
        response => {
          this.feeds = JSON.parse(response.body)
          this.total = this.feeds.length
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
      feeds: []
    }
  },
  created() {
    this.loadFeeds()
  },
}
</script>
