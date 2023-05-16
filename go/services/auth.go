package services

type authResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthService interface {
	Login(email string, password string) (res authResponse, err error)
	Refresh(headers map[string]string) (res authResponse, err error)
	Verify(headers map[string]string) error
}
