package models

type SMS struct {
	ApiKey    string `json:"apikey"`
	PartnerID string `json:"partnerID"`
	Shortcode string `json:"shortcode"`
	Message   string `json:"message"`
	Mobile    string `json:"mobile"`
}

/*type AutoGenerated struct {
	Count   int       `json:"count"`
	Smslist []Smslist `json:"smslist"`
}*/

type Smslist struct {
	PartnerID   string `json:"partnerID"`
	Apikey      string `json:"apikey"`
	PassType    string `json:"pass_type"`
	Clientsmsid int    `json:"clientsmsid"`
	Mobile      string `json:"mobile"`
	Message     string `json:"message"`
	Shortcode   string `json:"shortcode"`
}
