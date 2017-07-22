import request from 'superagent'

var __API__ = '/api'
export default {
  loadSources(errCb, successCb) {
    request.get(`${__API__}/feeds`)
      .end((err, res) => {
        if (err) {
          errCb(err)
          return
        }
        successCb(JSON.parse(res.text))
      })
  },
  loadNews({ selectedSources, searchMode }, errCb, successCb) {
    let selectedIds = selectedSources.filter(x => x !== 0)
    let r = request.post(`${__API__}/news`)
    r = r.send('limit=100')
    if (searchMode === 'unread') {
      r = r.send('search=unread')
    }
    if (selectedIds.length !== 0) {
      r = r.send(`ids=${selectedIds.join(',')}`)
    }
    r.end((err, res) => {
      if (err) {
        errCb(err)
        return
      }
      successCb(JSON.parse(res.text))
    })
  },
  markRead({ newsId, feedId }, errCb, successCb) {
    request.post(`${__API__}/feeds/${feedId}/news/${newsId}`)
      .send('action=read') // sending string automatically makes it form URL encoded
      .end((err, res) => {
        if (err) {
          errCb(err)
          return
        }
        successCb()
      })
  },
  classify({ newsId, classification }, errCb, successCb) {
    request.post(`${__API__}/learn`)
      .send(`news_id=${newsId}`) // sending string automatically makes it form URL encoded
      .send(`classification=${classification}`) // sending string automatically makes it form URL encoded
      .end((err, res) => {
        if (err) {
          console.log(err)
          return
        }
        successCb()
      })
  },
  refreshFeed(errCb, successCb) {
    request.post(`${__API__}/feeds`).end((err, res) => {
      this.isLoading = false
      if (err) {
        errCb(err)
        return
      }
      successCb(res.text)
    })
  }
}
