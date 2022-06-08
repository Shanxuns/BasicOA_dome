package task

import (
	"BasicOA/response"
	"BasicOA/serve/regular"
	"BasicOA/serve/user"
	_type "BasicOA/type"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func WebSocketTask(c *gin.Context) {
	// 创建User结构体
	var (
		User _type.User
		err  error
		err1 error
	)
	//升级get请求为webSocket协议
	if !websocket.IsWebSocketUpgrade(c.Request) {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: "请用webSocket协议"})
		return
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {

		}
	}(ws)
	// 写入并验证参数
	User.Email, err = response.GetHeader(c, regular.Email, "email")
	if err != nil {
		User.Email, _, err1 = response.GetQuery(c, regular.Email, "email")
		err = err1
	}
	if err != nil && err1 != nil {
		if err = ws.WriteMessage(websocket.TextMessage, []byte(err1.Error())); err != nil {
			return
		}
		return
	}
	User.Password, err = response.GetHeader(c, regular.Md5, "password")
	if err != nil {
		User.Password, _, err1 = response.GetQuery(c, regular.Md5, "password")
		err = err1
	}
	if err != nil && err1 != nil {
		if err = ws.WriteMessage(websocket.TextMessage, []byte(err1.Error())); err != nil {
			return
		}
		return
	}
	if err, User = user.User(User); err != nil {
		if err = ws.WriteMessage(websocket.TextMessage, []byte(err.Error())); err != nil {
			return
		}
		return
	}
	clients[ws] = User
	for {
		var Respond _type.Respond
		Respond = <-Channel
		for conn, v := range clients {
			if Respond.Code == 3 {
				if User.Id == Respond.Task {
					if err = conn.WriteJSON(Respond); err != nil {
						delete(clients, conn)
					}
					break
				}
				continue
			}
			if v.Sector == Respond.Task || Respond.Task == 0 {
				if err = conn.WriteJSON(Respond); err != nil {
					delete(clients, conn)
				}
			}
		}
	}
}
