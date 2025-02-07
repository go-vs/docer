package main

import "github.com/go-vs/docer"

type UserProfile struct {
	Fullname string `json:"fullname"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type Pet struct {
	ID      uint         `json:"id"`
	Name    string       `json:"name"`
	Type    string       `json:"type"`
	Profile *UserProfile `json:"profile"`
}

type User struct {
	ID       uint         `json:"id"`
	Username string       `json:"username"`
	Profile  *UserProfile `json:"profile"`
	Pets     []*Pet       `json:"pets"`
}

type Query struct {
	Type string `query:"type"`
}

type ResponseMeta struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

type Response struct {
	Meta ResponseMeta `json:"meta"`
	Data string       `json:"data"`
}

func main() {
	doc := docer.New()
	doc.HasBody(User{}, "json")
	doc.HasQuery(Query{}, "query")
	doc.HasResponse(Response{}, "json")
	if err := doc.JSON("doc.json"); err != nil {
		panic(err)
	}
	if err := doc.Generate("doc.md"); err != nil {
		panic(err)
	}
}
