package cart

import (
	"ecom/types"
	"fmt"
)

func getCartItemIDs(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for product %d", item.ProductID)
		}
		productIds[i] = item.ProductID
	}
	return productIds, nil
}

func (h *Handler) createOrder(products []types.Product, items []types.CartItem, userID int) (int, float64, error) {
	productMap := make(map[int]types.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}
	// check if all products are actually in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}
	// calculate total price
	totalPrice := calculateTotalPrice(items, productMap)
	// reduce quantity of products in db
	for _, item := range items {
		product := productMap[item.ProductID]
		item.Quantity = product.Quantity - item.Quantity

		if err := h.productStore.UpdateProductQuantity(item); err != nil {
			return 0, 0, err
		}
	}
	// create order
	orderID, err := h.store.CreateOrder(&types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "Mulago, Lagos, Nigeria",
	})
	if err != nil {
		return 0, 0, err
	}
	// create order items
	for _, item := range items {
		_, err2 := h.store.CreateOrderItem(&types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
		if err2 != nil {
			return 0, 0, err2
		}
	}
	return orderID, totalPrice, nil
}

func calculateTotalPrice(cartItems []types.CartItem, productMap map[int]types.Product) float64 {
	var totalPrice float64
	for _, item := range cartItems {
		product := productMap[item.ProductID]
		totalPrice += product.Price * float64(item.Quantity)
	}
	return totalPrice
}

func checkIfCartIsInStock(cartItems []types.CartItem, productMap map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := productMap[item.ProductID]

		if !ok {
			return fmt.Errorf("product %d not available in the store", item.ProductID)
		}

		if item.Quantity > product.Quantity {
			return fmt.Errorf("product %d is not available in the quantity requested", item.ProductID)
		}
	}
	return nil
}
