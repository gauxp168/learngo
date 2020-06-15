package controllers

import (
	"container/list"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"simqo.com/mygospace/learngo/example/sample_im/models"
	"time"
)

type Subscriber struct {
	Name string
	Conn *websocket.Conn
}

var (
	// websocket
	subscribe = make(chan Subscriber, 10)
	unsubscribe = make(chan string, 10)
	publish = make(chan models.Event, 10)

	// longpolling
	waitingList = list.New()
	subscribers = list.New()
)

func Join(user string, conn *websocket.Conn)  {
	subscribe <- Subscriber{Name:user,Conn:conn}
}

func newEvent(ep models.EventType, user string, msg string) models.Event {
	return models.Event{ep, user, int(time.Now().Unix()), msg}
}

func Leave(user string)  {
	unsubscribe <- user
}

type Subscription struct {
	Archive []models.Event
	New <-chan models.Event
}

func chatRoom()  {
	for  {
		select {
		case sub := <- subscribe:
			if !isUserExist(subscribers, sub.Name) {
				subscribers.PushBack(sub)
				publish <- newEvent(models.EVENT_JOIN,sub.Name,"")
				beego.Info("new join")
			} else {
				beego.Info("old join")
			}
		case event := <-publish:
			for ch := waitingList.Back(); ch != niil; ch.Prev() {
				ch.Value.(chan bool) <- true
				waitingList.Remove(ch)
			}
			broadcastWebSocket(event)
			models.NewArchive(event)
			if event.Type == models.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
			}
		case unsub := <- unsubscribe:
			for sub := subscribers.Front(); sub != nil; sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					subscribers.Remove(sub)
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Info("close websocket conn")
					}
					publish <- newEvent(models.EVENT_LEAVE, unsub, "")
					break
				}
			}
		}
	}
}

func init()  {
	go chatRoom()
}

func isUserExist(subscriber *list.List, user string) bool {
	for sub := subscriber.Front(); sub != nil; sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return  true
		}
	}
	return  false
}