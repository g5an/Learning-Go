package main

import "fmt"

type User struct {
	ID   string `json:"id" description:"identifier of the user"`
	Name string `json:"name" description:"name of the user" default:"john"`
	Age  int    `json:"age" description:"age of the user" default:"21"`
}
type UserResource struct {
	// normally one would use DAO (data access object)
	users map[string]User
}

func main() {
	// var u User
	// u.ID = "st"
	// u.Name = "nm"
	// u.Age = 11
	// list := []User{}
	// list = append(list, u)
	// fmt.Printf("%T\n", list)
	// fmt.Printf("%v\n", list)
	yonghong := User{ID: "1"}
	fmt.Println(yonghong)
	yonghong.Name = "sst"
	yonghong.Age = 11
	ren := UserResource{users: make(map[string]User)}
	ren.users[yonghong.ID] = yonghong
	fmt.Println(yonghong)
	fmt.Println(ren.users)
}
