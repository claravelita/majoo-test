package dto

type (
	UserRequest struct {
		Name         string `json:"name" validate:"required"`
		UserName     string `json:"user_name" validate:"required"`
		Password     string `json:"password" validate:"required"`
		MerchantName string `json:"merchant_name" validate:"required"`
	}

	UserLoginRequest struct {
		UserName string `json:"user_name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	UserAccessResponse struct {
		Name         string `json:"name"`
		UserName     string `json:"user_name"`
		MerchantName string `json:"merchant_name"`
		AccessToken  string `json:"access_token"`
	}
)
