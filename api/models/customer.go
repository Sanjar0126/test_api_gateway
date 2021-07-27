package models

//CreateCustomerModel ...
type CreateCustomerModel struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//UpdateCustomerModel ...
type UpdateCustomerModel struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
}

//GetCustomerModel ...
type GetCustomerModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	IsBlocked   bool   `json:"is_blocked"`
	CreatedAt   string `json:"created_at"`
	FcmToken    string `json:"fcm_token"`
	TgChatId    string `json:"tg_chat_id"`
	DateOfBirth string `json:"date_of_birth"`
}

//GetAllCustomersModel ...
type GetAllCustomersModel struct {
	Count     int                `json:"count"`
	Customers []GetCustomerModel `json:"customers"`
}

//CheckCustomerLoginRequest ...
type CustomerLoginRequest struct {
	Phone string `json:"phone"`
	Tag   string `json:"tag"`
}

//CheckCustomerLoginResponse ...
type CustomerLoginResponse struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmCustomerLoginRequest ...
type ConfirmCustomerLoginRequest struct {
	Code        string `json:"code"`
	Phone       string `json:"phone"`
	FcmToken    string `json:"fcm_token"`
	TgChatId    string `json:"tg_chat_id"`
	BotLanguage string `json:"bot_language"`
}

//ConfirmCustomerLoginResponse ...
type ConfirmCustomerLoginResponse struct {
	ID           string `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	IsBlocked    bool   `json:"is_blocked"`
	CreatedAt    string `json:"created_at"`
}

//SearchByPhoneResponse ...
type SearchByPhoneResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
}

// RegisterModel ...
type RegisterModel struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

//RegisterConfirmModel ...
type RegisterConfirmModel struct {
	Phone       string `json:"phone"`
	Code        string `json:"code"`
	FcmToken    string `json:"fcm_token"`
	TgChatId    string `json:"tg_chat_id"`
	BotLanguage string `json:"bot_language"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type BotLanguageModel struct {
	Lang string `json:"lang"`
}
