package domain

//	User model infor
//
// @Description	User information
type User struct {
	PK   string `dynamodbav:"PK" json:"-"`
	SK   string `dynamodbav:"SK" json:"-"`
	Type string `dynamodbav:"Type" json:"-"`
	Id   string `dynamodbav:"UserId" json:"id" example:"realfabecker@gmail.com"`
	Name string `dynamodbav:"Name" json:"name" example:"Rafael Becker"`
} // @name	User

// UserToken
type UserToken struct {
	RefreshToken  *string `json:"RefreshToken,omitempty"`
	AccesToken    *string `json:"AccessToken,omitempty"`
	AuthChallenge *string `json:"AuthChallange,omitempty"`
	AuthSession   *string `json:"AuthSession,omitempty"`
} // @name	UserToken

// UserLoginDTOQ
type UserLoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
} // @name	UserLoginDTO

// UserLoginChangeDTO
type UserLoginChangeDTO struct {
	Email       string `json:"email" validate:"required,email"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
	Session     string `json:"session" validate:"required"`
} // @name	UserLoginChangeDTO
