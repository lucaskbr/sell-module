package models

type ProductDepositMovement struct {
	ProductDepositID uint `json:"productDepositId"`
	Quantity uint `json:"quantity"`
}


type MovementProduct struct {
	Type string `json:"type"`
	DocumentID int `json:"documentId"`
	ProductDepositMovementList []ProductDepositMovement `json:"productDepositMovementList"`
}