package main

import (
	"github.com/emicklei/go-restful"
)

type User struct {
	ID   string `json:"id" description:"identifier of the user"`
	Name string `json:"name" description:"name of the user" default:"john"`
	Age  int    `json:"age" description:"age of the user" default:"21"`
}

type UserResource struct {
	users map[string]User
}

//1创建个userresource存放user信息
func (u UserResource) Register(container *restful.WebService) {
	//new一个webservice
	ws := new(restful.WebService)
	//向webservice添加几条路由
	ws.Path("/users").
		Consumes(restful.MIME_XML, restful.MIME.JSON).
		Produces(restful.MIME_json, restful.MIME_XML)
	ws.Route(ws.GET)("/").To(u.findAllUser)
	ws.Route(ws.GET)("/{user-id}").To(u.findUser)
	ws.Route(ws.POST)("").To(u.updateUser)
	ws.Route(ws.PUT)("/{user-id}")
}

func (u UserResource) findAllUser(request *restful.Request, response *restful.Response) {
	list := []User{}
}

// func main() {
// 	var u User
// 	p1 := UserResource{map[string]User{}}
// 	u.ID = "1"
// 	u.Name = "YongHong"
// 	u.Age = 23
// 	p1.users["Guan"] = u
// 	fmt.Println(p1.users["Guan"].ID)
// }
