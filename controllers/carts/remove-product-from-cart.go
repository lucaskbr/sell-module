package carts

// import (
// 	"github.com/gofiber/fiber/v2"
// )

// // RemoveProductFromCart - DELETE
// func RemoveProductFromCart(c *fiber.Ctx) error {

// 	var UUID = c.Params("uuid")

// 	_, err := repository.FindCartByUUID(UUID)

// 	if err != nil {
// 		fmt.Println(err)
// 		return c.SendStatus(500)
// 	}

// 	var pdmList []productDepositMovement

// 	var pdm productDepositMovement

// 	pdm.ProductDepositID = body.ProductDepositID
// 	pdm.Quantity = body.Quantity

// 	pdmList = append(pdmList, pdm)

// 	var mp movementProduct

// 	mp.DocumentID = 0
// 	mp.Type = "PICKUP_RESERVATION"
// 	mp.ProductDepositMovementList = pdmList

// 	jsonReq, err := json.Marshal(mp)

// 	if err != nil {
// 		fmt.Println(err)
// 		return c.SendStatus(500)
// 	}

// 	resp, err := http.Post("http://localhost:8080/product-movement", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))

// 	if err != nil {
// 		fmt.Println(err)
// 		return c.SendStatus(500)
// 	}

// 	defer resp.Body.Close()

// 	bodyBytes, _ := ioutil.ReadAll(resp.Body)

// 	// Convert response body to string
// 	bodyString := string(bodyBytes)
// 	fmt.Println("API Response as String:\n" + bodyString)

// 	_, err = repository.RemoveProductFromCart(UUID, body.ProductDepositID, body.Quantity);

// 	if err != nil {
// 		return c.SendStatus(500)
// 	}

// 	return c.SendStatus(204)
// }