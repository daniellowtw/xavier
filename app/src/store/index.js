import Vuex from 'vuex'
import Vue from 'vue'

const debug = process.env.NODE_ENV !== 'production'
Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isDebug: false,
    sources: [],
  },
  mutations: {
    updateSources(state, sources) {
      state.sources = sources
    },
    toggleDebug(state) {
      state.isDebug = !state.isDebug
    },
    readNews(state, feedID) {
      let feed = state.sources.find(x => x.Id === feedID)
      if (feed) {
        feed.UnreadCount--
      }
    }
  },
  strict: debug,
})
