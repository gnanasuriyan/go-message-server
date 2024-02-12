package models

type PaginationDto struct {
	Page  int `json:"page" query:"page"`
	Limit int `json:"limit" query:"limit"`
}

type LoginRequestDto struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type SuccessResponseDto struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
