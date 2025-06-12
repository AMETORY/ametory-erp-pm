package objects

type GetShopResponse struct {
	Shops []Shop `json:"shops"`
}

type Shop struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	SellerType string `json:"seller_type"`
	Cipher     string `json:"cipher"`
	Code       string `json:"code"`
}
