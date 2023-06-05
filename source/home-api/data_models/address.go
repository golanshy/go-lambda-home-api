package data_models

type Address struct {
	FirstLine  string `json:"first_line,omitempty"`
	SecondLine string `json:"second_line,omitempty"`
	Town       string `json:"town,omitempty"`
	Postcode   string `json:"postcode,omitempty"`
	County     string `json:"county,omitempty"`
	Country    string `json:"country,omitempty"`
}
