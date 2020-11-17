package models

type address struct {
	Cep string `json:"cep"`
	Number int `json:"number"`
}

type user struct {
	Email string `json:"email"`
	Name string `json:"name"`
	Lastname string `json:"lastName"`
	Cpf int `json:"cpf"`
	PhoneNumber string `json:"phoneNumber"`
}

type CheckoutBody struct {
	CartUUID string `json:"cartUUID"`
	User user `json:"user"`
 	Address address `json:"address"`
}
