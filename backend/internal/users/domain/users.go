package domain

//  User model infor
//	@Description	User information
type User struct {
	PK   string `dynamodbav:"PK" json:"-"`
	SK   string `dynamodbav:"SK" json:"-"`
	Type string `dynamodbav:"Type" json:"-"`
	Id   string `dynamodbav:"UserId" json:"id" example:"realfabecker@gmail.com"`
	Name string `dynamodbav:"Name" json:"name" example:"Rafael Becker"`
} // @name User

// UserToken
type UserToken struct {
	RefreshToken string `json:"RefreshToken"`
	AccesToken   string `json:"AccessToken"`
} // @name UserToken

// WalletLoginDTOQ
type WalletLoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
} // @name WalletLoginDTO
