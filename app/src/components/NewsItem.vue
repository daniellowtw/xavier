<template>
  <div class="card">
    <header class="card-header clickable">
      <p class="card-header-title" v-on:click="toggleShow" v-bind:class="{read: (news.read===true)}">
        <span class="icon">
          <i>
            <img v-bind:src="fav">
          </i>
        </span> {{news.Title}}
      </p>
      <a class="card-header-icon">
        <span class="icon">
          <i class="fa fa-angle-down" v-on:click="toggleShow"></i>
        </span>
      </a>
    </header>
    <div class="card-content" v-show="show">
      <div v-show="isDebug">
        {{news}}
      </div>
      <div class="content" v-html="news.Description">
      </div>
      <div class="content" v-html="news.Content">
      </div>
    </div>
    <footer class="card-footer" v-show="show">
      <a class="card-footer-item">Save</a>
      <a class="card-footer-item">Edit</a>
      <a class="card-footer-item">Delete</a>
    </footer>
  </div>
</template>
<script>
import Vue from 'vue'
export default Vue.component('news-item', {
  props: ['news', 'isDebug', 'fav'],
  data() {
    return {
      show: false,
      content: '',
    }
  },
  methods: {
    toggleShow() {
      this.$emit('read', this.news)
      this.show = !this.show
    },
  },
})
</script>
<style>
.clickable {
  cursor: pointer;
}

.read {
  font-size: small;
}
</style>