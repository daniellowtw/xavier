<template>
  <!-- Main container -->
  <nav id="feed-bar" class="level">
    <!-- Left side -->
    <div class="level-left">
    </div>
  
    <!-- Right side -->
    <div class="level-right">
      <p class="level-item">
        <a class="button is-success" v-on:click.once="refreshFeed">Refresh</a>
      </p>
      <p class="level-item">
        <a class="button is-success" v-on:click="showModel=true">New</a>
  
      </p>
    </div>
    <div v-bind:class="{'modal':true, 'is-active':showModel}">
      <div class="modal-background"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">Add feed</p>
          <button class="delete" v-on:click="showModel = false"></button>
        </header>
        <section class="modal-card-body">
          <div class="field">
            <div class="control">
              <input class="input" type="text" v-model="url" placeholder="URL">
            </div>
          </div>
        </section>
        <footer class="modal-card-foot">
          <a class="button is-success" v-on:click="addUrl()">Add</a>
          <a class="button" v-on:click="showModel = false">Cancel</a>
        </footer>
      </div>
    </div>
  </nav>
</template>
<style>
#feed-bar {
  padding-top: 10px;
}
</style>
</style>

<script>
import Vue from 'vue'
export default Vue.component('feed-bar', {
  data() {
    return {
      url: '',
      showModel: false,
    }
  },
  methods: {
    addUrl() {
      let t = this
      this.$store.dispatch('addFeed', this.url).then(function () {
        t.showModel = false
      })
    },
    refreshFeed() {
      this.$store.dispatch('refreshFeed')
    }
  },
})
</script>
