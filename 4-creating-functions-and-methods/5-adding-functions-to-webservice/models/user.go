package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)
/*
func that allows me to return the
users stored in model layer
 */
func GetUsers() []*User {
	return users
}
/*
func that allows me to add users
into user collection
*/
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}