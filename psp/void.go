package psp

type VoidRequest struct {
	AuthorizeId string `json:"authorizeId"`
	Amount      Amount `json:"amount"`
}

type VoidResponse struct {
}
