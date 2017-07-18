<template>
  <aside class="menu">
    <p class="menu-label">
      Feeds
    </p>
    <ul class="menu-list">
      <li v-for="source in sources" :key="source.Id">
        <a class="has-shadow" v-on:click="toggleSource(source.Id)">{{source.Title}}
          <span class="icon">({{source.UnreadCount}})</span>
        </a>
      </li>
    </ul>
  </aside>
</template>

<script>
import Vue from 'vue'
export default Vue.component('news-menu', {
  props: ['isDebug', 'sources'],
  data() {
    return {
      isLoading: false,
    }
  },
  methods: {
    toggleSource(id) {
      console.log(id)
      this.$emit('toggle-source', id)
    }
  }
})
</script>
