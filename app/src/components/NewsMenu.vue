<template>
  <aside class="menu">
    <p class="menu-label">
      Feeds
    </p>
    <ul class="menu-list">
      <li v-for="(source, index) in sources" :key="source.Id">
        <a v-bind:class="{'primary':selectedSources[index] !== 0}" v-on:click="toggleSource(source, index)">{{source.Title}}
          <span class="icon">({{source.UnreadCount}})</span>
        </a>
      </li>
    </ul>
  </aside>
</template>

<script>
import Vue from 'vue'
export default Vue.component('news-menu', {
  props: ['isDebug', 'sources', 'selectedSources'],
  data() {
    return {
      isLoading: false,
    }
  },
  methods: {
    toggleSource(source, index) {
      this.$emit('toggle-source', source.Id, index)
    }
  },
})
</script>
<style>
.primary {
  color: #fff;
  background: #00c4a7;
}
</style>