package term_app

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
	"github.com/daniellowtw/xavier/feed"
)

type State uint

const (
	SELECT_FEED State = iota
	SELECT_FEED_ITEM
)

type App struct {
	Client *feed.Client
	state  State
}

func (a *App) Run() error {
	r, err := readline.New("> ")// make this dynamic
	if err != nil {
		return err
	}
	a.Client.ListAllFeeds()
	for {
		line, err := r.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)

		if err := a.process(line); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) process(line string) error {
	var err error
	switch a.state {
	case SELECT_FEED:
		a.state, err = a.processSelectFeedId(line)
	case SELECT_FEED_ITEM:
		a.state, err = a.processSelectFeedId(line)
	default:
		fmt.Println("What are you talking about?")
	}
	return err
}

func (a *App) processSelectFeedId(line string) (nextState State, err error) {
	_, err = strconv.Atoi(line)
	if err != nil {
		a.Client.ListAllFeeds()
		return SELECT_FEED, err
	}
	nextState = SELECT_FEED_ITEM
	a.Client.ListFeedItems(line)
	return
}
