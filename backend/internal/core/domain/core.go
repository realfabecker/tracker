package domain

type Config struct {
	AppPort           string `env:"APP_PORT"`
	AppName           string `env:"APP_NAME"`
	AppHost           string `env:"APP_HOST"`
	CognitoClientId   string `env:"COGNITO_CLIENT_ID"`
	CognitoJwkUrl     string `env:"COGNITO_JWK_URL"`
	DynamoDBTableName string `env:"DYNAMODB_TABLE_NAME"`
}

type EmptyResponseDTO struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message,omitempty" example:"Operação realizada com sucesso"`
	Code    int    `json:"code,omitempty" example:"200"`
} //	@name	EmptyResponseDTO

type ResponseDTO[T any] struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message,omitempty" example:"Operação realizada com sucesso"`
	Code    int    `json:"code,omitempty" example:"200"`
	Data    *T     `json:"data,omitempty"`
} //	@name	ResponseDTO

type PagedDTO[T interface{}] struct {
	PageCount int32  `json:"page_count" example:"10"`
	Items     []T    `json:"items"`
	PageToken string `json:"page_token,omitempty"`
	HasMore   bool   `json:"has_more" example:"false"`
} //	@name	PagedDTO

type PagedDTOQuery struct {
	Limit     int32  `query:"limit" validate:"required,max=50"`
	PageToken string `query:"page_token"`
} //	@name	PagedDTOQuery
