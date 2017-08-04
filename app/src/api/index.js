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
    r = r.send('limit=100').send(`search=${searchMode}`)
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
  markReadMulti({ newsIds }, errCb, successCb) {
    const formattedNewsIds = formatNewsIds(newsIds)
    request.post(`${__API__}/read`)
      .send(`news_id=${formattedNewsIds}`)
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
      .send(`news_id=${newsId}`)
      .send(`classification=${classification}`)
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
      if (err) {
        errCb(err.response.text)
        return
      }
      successCb(res.text)
    })
  },
  addFeed(url, errCb, successCb) {
    request.put(`${__API__}/feeds`)
      .send(`url=${url}`)
      .end((err, res) => {
        if (err) {
          errCb(err.response.text)
          return
        }
        successCb(res.text)
      })
  },
  deleteFeed(id, errCb, successCb) {
    request.delete(`${__API__}/feeds/${id}`)
      .end((err, res) => {
        if (err) {
          errCb(err.response.text)
          return
        }
        successCb(res.text)
      })
  },
  saveNewsItem({ newsId, feedId }, errCb, successCb) {
    request.post(`${__API__}/feeds/${feedId}/news/${newsId}`)
      .send('action=toggle-save')
      .end((err, res) => {
        if (err) {
          errCb(err.response.text)
          return
        }
        successCb(res.text)
      })
  },

}

function formatNewsIds(newsIds) {
  let res = ''
  newsIds.forEach(x => { res += `${x},` })
  return res.substr(0, res.length - 1)
}
