package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"simqo.com/mygospace/learngo/example/sample_im/models"
)

type WebSocketController struct {
	baseController
}

func (this *WebSocketController) Get()  {
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}
	this.TplName = "websocket.tpl"
	this.Data["UserName"] = uname
	this.Data["isWebSocket"] = true

}

func (this *WebSocketController) Join()  {
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _,ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "", 400)
		return
	}else if err != nil {
		beego.Error("")
		return
	}
	Join(uname, ws)
	defer Leave(uname)

	for{
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}

func broadcastWebSocket(event models.Event)  {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}
	for sub := subscribers.Front(); sub != nil; sub.Next() {
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if err := ws.WriteMessage(websocket.TextMessage, data); err != nil {
				unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}
