package model


type User struct {
	ID string  `json:"id" form: "id" query: "id" `
	Name string `json:"name" form: "name" query "name" `
}


