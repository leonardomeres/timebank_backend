package models

/* These structs are only used for swagger documentation */

type GenericErrorResponse struct {
	Error string `json:"error" example:"Operation failed"`
}

type GenericUnauthorizedResponse struct {
	Error string `json:"error" example:"Unauthorized"`
}

type UserRegistrationResponse struct {
	Message string `json:"message" example:"User created successfully"`
}

type UserRegistrationErrorResponse struct {
	Error string `json:"error" example:"Email already registered"`
}

type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type LoginFailResponse struct {
	Error string `json:"error" example:"Invalid password credentials"`
}

type SkillCreationResponse struct {
	Response string `json:"response" example:"Skill created successfully"`
}

type SkillFailCreationResponse struct {
	Error string `json:"error" example:"Failed to create skill"`
}

type SkillExistsResponse struct {
	Error string `json:"error" example:"Skill already exists"`
}

type GetProfileResponse struct {
	ID      uint     `json:"id" example:"1"`
	Email   string   `json:"email" example:"bob@example.com.br"`
	Name    string   `json:"name" example:"Bob"`
	Balance float64  `json:"balance" example:"4.0"`
	Skills  []string `json:"skills" example:"[\"Programming\", \"Design\"]"`
}
