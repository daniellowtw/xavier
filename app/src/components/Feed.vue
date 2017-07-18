<template>
  <section class="section">
    <feed-bar @refresh-feed="loadFeeds"></feed-bar>
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
                <span class="icon">
                  <i>
                    <img v-bind:src="feed.FavIcon">
                  </i>
                </span>
                <a :href="feed.UrlSource">{{feed.Title}}</a>
              </td>
              <td>{{feed.Description}}
                <div v-show="isDebug">
                  {{feed}}
                </div>
              </td>
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
import Pagination from './Pagination.vue'
import FeedBar from './FeedBar.vue'
export default {
  props: ['isDebug', 'feeds'],
  name: 'feed',
  components: [
    Pagination,
    FeedBar
  ],
  methods: {
    onChangePage(page) {
      this.page = page
    },
  },
  data() {
    return {
      page: 1,
      total: 100,
      itemsPerPage: 10,
    }
  },
  created() {
    this.loadFeeds()
  },
}
</script>
