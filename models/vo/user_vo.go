package vo

type AddUserVO struct {
	Account   string `json:"account"`
	Activated bool   `json:"activated"`
	Role      uint   `json:"role"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	SendEmail bool   `json:"send_email"`
}
