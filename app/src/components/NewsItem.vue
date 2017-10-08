<template>
  <div class="card">
    <header class="card-header">
      <p class="card-header-title" v-bind:class="{read: (news.Read)}">
        <span class="icon clickable">
            <i v-if="news.IsSaved" v-on:click="save()" class="fa fa-star" aria-hidden="true"></i>
            <i v-else v-on:click="save()" class="fa fa-star-o" aria-hidden="true"></i>
        </span>

        <span class="icon">
          <i>
            <img v-bind:src="favIcon">
          </i>
        </span>
        <span v-on:click="toggleShow" class="clickable">
          {{news.Title}}
        </span>
        <span class="entry-separator">•</span>
        <a v-bind:href="news.LinkHref" class="sub" target="blank"> {{news.Published}}</a>
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
      <a class="card-footer-item" v-on:click="classify(2)" v-bind:class="{'is-primary': (news.Classification === 2)}">Don't like it</a>
      <a class="card-footer-item" v-on:click="save()" v-bind:class="{'is-primary': (news.IsSaved)}">Save</a>
      <a class="card-footer-item" v-on:click="classify(1)" v-bind:class="{'is-primary': (news.Classification === 1)}">Like it</a>
    </footer>
  </div>
</template>
<script>
import Vue from 'vue'
import { mapState } from 'vuex'
export default Vue.component('news-item', {
  props: ['news'],
  computed: {
    favIcon() { let t = this.sources.find(x => x.Id === this.news.FeedId); return t === undefined ? '' : t.FavIcon },
    ...mapState({
      isDebug: 'isDebug',
      sources: 'sources',
      currentNewsId: 'currentNewsId',
    }),
  },
  data() {
    return {
      show: false,
      content: '',
    }
  },
  methods: {
    toggleShow(evt) {
      var curr = evt.srcElement
      var max = 3
      while (max > 0 && curr !== undefined && curr.className !== 'card') {
        curr = curr.parentNode
        max -= 1
      }
      curr.scrollIntoView()
      this.show = !this.show
      this.markRead()
    },
    markRead() {
      this.$store.dispatch('markRead', { newsId: this.news.Id, feedId: this.news.FeedId })
    },
    classify(v) {
      this.$store.dispatch('classify', { newsId: this.news.Id, feedId: this.news.FeedId, classification: v })
    },
    save() {
      this.$store.dispatch('saveNewsItem', { newsId: this.news.Id, feedId: this.news.FeedId })
    }
  },
  watch: {
    currentNewsId(v) {
      if (this.news.Id !== v) {
        this.show = false
      }
    },
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

  .sub {
    display: inline;
    color: #aaa;
    align-self: center;
    font-size: 0.7em;
  }

  .entry-separator {
    margin-right: 5px;
    margin-left: 5px
  }
</style>
