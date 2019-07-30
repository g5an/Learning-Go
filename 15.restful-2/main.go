package main

import (
	"net/http"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
)

type User struct {
	ID   string `json:"id" description:"identifier of the user"`
	Name string `json:"name" description:"name of the user" default:"john"`
	Age  int    `json:"age" description:"age of the user" default:"21"`
}
type UserResource struct {
	// normally one would use DAO (data access object)
	users map[string]User
}

func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME.JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)
	tags := []string{"users"}

	ws.Route(ws.GET("/").To(u.findAllUsers).
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]User{}).
		Returns(200, "OK", []User{}))

	return ws
}

func (u UserResource) findAllUsers(request *restful.Request, response *restful.Response) {
	list := []User{}
	for _, each := range u.users {
		list = append(list, each)
	}
	response.WriteEntity(list)
}

func main() {
	ws := new(restful.WebService)
	ws.Path("/users")
	ws.Route(ws.GET("/").To(u.findAllUsers).
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]User{}).
		Returns(200, "OK", []User{}))

	container := restful.NewContainer().Add(ws)
	http.ListenAndServe(":8089", container)

}
