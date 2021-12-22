package web

type HealthFaRegisterRequest struct {
	Name      string `json:"facilitator_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Type      string `json:"type"`
}

type HealthFaRegisterResponse struct {
	Name      string `json:"facilitator_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Type      string `json:"type"`
}

type HealthFaLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
