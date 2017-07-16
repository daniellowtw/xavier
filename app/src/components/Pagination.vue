<template>
  <nav class="pagination">
    <a class="pagination-previous" v-show="showPreviousButton" @click="goPreviousPage()"><<</a>
    <a class="pagination-next" v-show="showNextButton" @click="goNextPage()">>></a>
    <ul class="pagination-list">
      <li>
        <a class="pagination-link" v-show="showPreviousButton" @click="goFirstPage()">1</a>
      </li>
      <li>
        <span class="pagination-ellipsis" v-show="showPreviousButton">...</span>
      </li>
      <li>
        <a class="pagination-link" v-show="showPreviousButton" @click="goPage(page-1)">{{page-1}}</a>
      </li>
      <li>
        <span class="pagination-link is-current">{{page}}</span>
      </li>
      <li>
        <a class="pagination-link" v-show="showNextButton" @click="goPage(page+1)">{{page+1}}</a>
      </li>
      <li>
        <span class="pagination-ellipsis" v-show="showNextButton">...</span>
      </li>
      <li>
        <a class="pagination-link" v-show="showNextButton" @click="goLastPage()">{{totalPages}}</a>
      </li>
    </ul>
  </nav>
</template>
<script>
import Vue from 'vue'
export default Vue.component('pagination', {
  props: ['total', 'page', 'itemsPerPage'],
  computed: {
    totalPages: function () {
      return Math.ceil(this.total / this.itemsPerPage) || 1
    },
    showNextButton: function () {
      return this.page !== this.totalPages
    },
    showPreviousButton: function () {
      return this.page !== 1
    }
  },
  methods: {
    goNextPage: function () {
      this.$emit('change-page', this.page + 1)
    },
    goPreviousPage: function () {
      this.$emit('change-page', this.page - 1)
    },
    goFirstPage: function () {
      this.$emit('change-page', 1)
    },
    goLastPage: function () {
      this.$emit('change-page', this.totalPages)
    },
    goPage: function (page) {
      this.$emit('change-page', page)
    }
  }
})
</script>