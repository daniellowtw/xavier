<template>
  <div class="card">
    <header class="card-header clickable">
      <p class="card-header-title" v-on:click="toggleShow" v-bind:class="{read: (news.Read)}">
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
      <a class="card-footer-item" v-on:click="classify(-1)" v-bind:class="{'is-primary': (this.classification === -1)}">Don't like it</a>
      <a class="card-footer-item">Save</a>
      <a class="card-footer-item" v-on:click="classify(1)" v-bind:class="{'is-primary': (this.classification === 1)}">Like it</a>
    </footer>
  </div>
</template>
<script>
import Vue from 'vue'
import request from 'superagent'
var __API__ = '/api'
export default Vue.component('news-item', {
  props: ['news', 'isDebug', 'fav', 'currentNewsId'],
  data() {
    return {
      show: false,
      content: '',
      classification: 0,
    }
  },
  methods: {
    toggleShow() {
      this.show = !this.show
      this.$emit('read', this.news)
    },
    classify(v) {
      request.post(`${__API__}/learn`)
        .send(`news_id=${this.news.Id}`) // sending string automatically makes it form URL encoded
        .send(`classification=${v}`) // sending string automatically makes it form URL encoded
        .end((err, res) => {
          if (err) {
            console.log(err)
            return
          }
          this.classification = v
        }) // sending string automatically makes it form URL encoded
    }
  },
  watch: {
    currentNewsId(v) {
      console.log(v)
      if (this.news.Id !== v) {
        this.show = false
      }
    }
  }
})
</script>
<style>
.clickable {
  cursor: pointer;
}

.read {
  font-size: small;
}

.is-primary {
  color: #fff;
  background-color: #00d1b2;
}
</style>