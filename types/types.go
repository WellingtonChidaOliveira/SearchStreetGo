package types

type Address struct {
	ZipCode      string `json:"cep"`
	Street       string `json:"logradouro"`
	Complement   string `json:"complemento"`
	Unit         string `json:"unidade"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
	IbgeCode     string `json:"ibge"`
	GiaCode      string `json:"gia"`
	AreaCode     string `json:"ddd"`
	SiafiCode    string `json:"siafi"`
}
