package dto

type UserDTO struct {
	UserID uint   `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}
type UserResponse struct {
	Response
	UserDTO
}
