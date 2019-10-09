package model

//User Object to store user data
type User struct {
	Object
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

//JWT Object to store the token
type JWT struct {
	Token string `json:"token"`
}

//CheckPassword to store check password
type CheckPassword struct {
	Password string `json:"password"`
}

//ChangePassword Object to store change password data
type ChangePassword struct {
	ID              int    `json:"id"`
	Password        string `json:"password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

//DataUser Object to store user data
type DataUser struct {
	Token     string `json:"auth_token"`
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	// JobID     int    `json:"job_id"`
	JobName   string `json:"job_name"`
	CompanyID int    `json:"company_id"`
	Image     string `json:"image_url"`
	Role      string `json:"role"`
}

//ChangePasswordResponse to store data object and meta object
type ChangePasswordResponse struct {
	Data User      `json:"data"`
	Meta MetaModel `json:"meta"`
}

//LoginResponse Object to store data object and meta object
type LoginResponse struct {
	Data DataUser  `json:"data"`
	Meta MetaModel `json:"meta"`
}

//MetaModel for string the status or info from API
type MetaModel struct {
	Object
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
}

type Object struct {
}
