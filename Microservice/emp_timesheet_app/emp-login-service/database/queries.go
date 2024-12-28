package database

var LoginQueries = struct {
	GetUserPassword string
}{
	GetUserPassword: "SELECT password FROM Employee WHERE username = ?",
}
