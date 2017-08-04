import Vuex from 'vuex'
import Vue from 'vue'
import api from '../api'
// TODO it shouldn't be here!
import swal from 'sweetalert2'

const debug = process.env.NODE_ENV !== 'production'
Vue.use(Vuex)

export default new Vuex.Store({
  // initial state
  state: {
    isDebug: false,
    mode: 'feed',
    sources: [],
    news: [],
    currentNewsId: 0,
    currentNewsClassification: 0,
    isCurrentNewsSaved: false,
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
      selectedNews.Classification = classification
    },
    saveCurrentNews(state, { newsId, isSaved }) {
      let selectedNews = state.news.find(x => x.Id === newsId)
      selectedNews.IsSaved = isSaved
    },
    // TODO: think of a better way for alerting.
    notify(state, { title, body, type }) {
      swal(title, body, type)
    },
    changeMode(state, mode) {
      state.mode = mode
    }
  },
  actions: {
    loadSources({ commit }) {
      console.log('foooo')
      api.loadSources(console.log, x => commit('updateSources', x))
    },
    loadNews({ commit }, { selectedSources, searchMode }) {
      api.loadNews({
        selectedSources,
        searchMode,
      }, console.log, x => commit('updateNews', x))
    },
    markReadMulti({ commit }, news) {
      console.log(news)
      api.markReadMulti({ newsIds: news.map(x => x.Id) }, console.log, () => {
        news.forEach(p => {
          commit('readNews', { newsId: p.Id, feedId: p.FeedId })
        })
      })
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
    },
    addFeed({ commit, dispatch }, url) {
      return new Promise((resolve, reject) => {
        api.addFeed(url, x => commit('notify', { title: 'Error', body: x, type: 'error' }), () => {
          dispatch('loadSources')
          resolve()
        })
      })
    },
    deleteFeed({ commit, dispatch }, feedId) {
      api.deleteFeed(feedId, x => commit('notify', { title: 'Error', body: x, type: 'error' }), () => {
        dispatch('loadSources')
      })
    },
    saveNewsItem({ commit }, { newsId, feedId }) {
      api.saveNewsItem({ newsId, feedId }, x => commit('notify', { title: 'Error', body: x, type: 'error' }), (x) => {
        let isSaved = (x === 'true')
        commit('saveCurrentNews', { newsId, isSaved })
      })
    },
  },
  strict: debug,
})
