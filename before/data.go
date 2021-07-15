package main

type User struct {
	Id       string
	Name     string
	LastName string
}


func FindUser(id string) (*User, error) {

	users := make([]User, 1)

	users = append(users, User{
		Id:       "1",
		Name:     "Emre",
		LastName: "Tiryaki",
	})

	for _, user := range users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, ErrDocumentNotFound

}
