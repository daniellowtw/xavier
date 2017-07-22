import Vuex from 'vuex'
import Vue from 'vue'
import api from '../api'
// TODO it shouldn't be here!
import swal from 'sweetalert2'

const debug = process.env.NODE_ENV !== 'production'
Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isDebug: false,
    sources: [],
    news: [],
    currentNewsId: 0,
    currentNewsClassification: 0,
  },
  mutations: {
    updateSources(state, sources) {
      state.sources = sources
    },
    updateNews(state, news) {
      state.news = news
    },
    toggleDebug(state) {
      state.isDebug = !state.isDebug
    },
    readNews(state, { feedID, newsId }) {
      let selectedNews = state.news.find(x => x.Id === newsId)
      if (!selectedNews) {
        return
      }
      state.currentNewsId = newsId
      state.currentNewsClassification = selectedNews.classification
      if (selectedNews.Read) {
        state.currentNewsId = newsId
        return
      }
      selectedNews.Read = true
      let feed = state.sources.find(x => x.Id === feedID)
      if (feed) {
        feed.UnreadCount--
      }
    },
    classifyNews(state, { newsId, classification }) {
      let selectedNews = state.news.find(x => x.Id === newsId)
      selectedNews.classification = classification
      state.currentNewsClassification = classification
    },
    // TODO: think of a better way for alerting.
    notify(state, { title, body, type }) {
      swal(title, body, type)
    }
  },
  actions: {
    loadSources({ commit }) {
      api.loadSources(console.log, x => commit('updateSources', x))
    },
    loadNews({ commit }, { selectedSources, searchMode }) {
      api.loadNews({
        selectedSources,
        searchMode,
      }, console.log, x => commit('updateNews', x))
    },
    markRead({ commit }, { newsId, feedId }) {
      api.markRead({
        newsId,
        feedId,
      }, console.log, () => {
        commit('readNews', { feedId, newsId })
      })
    },
    classify({ commit }, { newsId, feedId, classification }) {
      api.classify({
        newsId, classification
      }, console.log, () => {
        commit('classifyNews', { newsId, classification })
      })
    },
    refreshFeed({ commit }) {
      api.refreshFeed(console.log, x => {
        commit('notify', { title: 'Updated feeds', body: x, type: 'success' })
      })
    }
  },
  strict: debug,
})
