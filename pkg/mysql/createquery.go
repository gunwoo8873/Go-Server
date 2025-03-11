package mysql

// Create user in the database
type Users struct {
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserRole     int    `json:"user_role"`
}

func CreateQurey() {

}
