package routes

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/superduperkevin/fiber-rest/database"
	"github.com/superduperkevin/fiber-rest/models"
)


type Order struct {
	ID uint `json:"id"`
	Product Product `json:"product"`
	User User `json:"user"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(orderModel models.Order, user User, product Product) Order {
	return Order{ID: orderModel.ID, Product: product, User: user, CreatedAt: orderModel.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)

	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product
		database.Database.Db.Find(&user, "id = ?", order.UserRefer)
		database.Database.Db.Find(&product, "id = ?", order.ProductRefer)
		responseUser := CreateResponseUser(user)
		responseProduct := CreateResponseProduct(product)
		responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
		responseOrders = append(responseOrders, responseOrder)
	}
	return c.Status(200).JSON(responseOrders)
}

func findOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("Order not found")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	
	var order models.Order

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product

	database.Database.Db.First(&user, order.UserRefer)
	database.Database.Db.First(&product, order.ProductRefer)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}

func DeleteOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil {
		return c.Status(400).JSON("User Not Found")
	}

	if err := findOrder(id, &order); err != nil {
		return c.Status(400).JSON("Order Not Found")
	}

	database.Database.Db.Delete(&order, id)

	return c.Status(200).SendString("Order Successfully Deleted")
}