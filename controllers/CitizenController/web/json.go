package web

//berisi struct req / res controllers dan konversi ke domain sepertinya..

//response
type ResponseRegister struct {
	Name  string
	Email string
	NIK   string
}

type RequestRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	NIK      string `json:"nik"`
}

type CitizenLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
