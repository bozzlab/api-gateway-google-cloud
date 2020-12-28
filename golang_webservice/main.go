package main

import (  
  "github.com/gofiber/fiber/v2"
  "github.com/google/uuid"
  "github.com/gofiber/fiber/v2/middleware/logger"
)

type InfoService struct {
    ServiceName string `json:"service_name"`
    Language string `json:"language"`
    Framework string `json:"framework"`
  }

type NotFoundMsg struct {
  Message string `json:"message"`
}

type TransactionObj struct {
  ClientID string `json:"client_id"`
  ProductName string `json:"product_name"`
  Provider string `json:"provider"`
  Amount float64 `json:"amount"`
  Descriptiion string `json:"description"`
  Currency string `json:"currency"`
}

type TransactionSuccess struct {
  Message string `json:"message"`
  TransactionID string `json:"transaction_id"`
}

type TransactionStatusObj struct {
  Status string `json:"status"`
}

var TransactionDataTemp = make(map[string]*TransactionObj)
var TransactionDataStatus = make(map[string]string)


func main() {
  app := fiber.New(fiber.Config{
    ErrorHandler: func(c *fiber.Ctx, err error) error {
        return c.Status(500).SendString(err.Error())
    },})

  app.Get("/info", func(c *fiber.Ctx) error {
      return c.Status(fiber.StatusOK).JSON(InfoService{
                    ServiceName: "Payment API",
                    Language:  "Golang",
                    Framework : "Fiber",
        })
    })

  app.Post("/payment", func(c *fiber.Ctx) error {
    tID := uuid.New()
    t := new(TransactionObj)
    if err := c.BodyParser(t); err != nil {
      return err
    }

    TransactionDataTemp[tID.String()] = t
    TransactionDataStatus[tID.String()] = "Pending"

    return  c.Status(fiber.StatusOK).JSON(TransactionSuccess{
        Message : "Success",
        TransactionID : tID.String(),
      })

    })
  app.Patch("/payment/status", func(c *fiber.Ctx) error {
    KeyID := c.Query("id")
    if TransactionDataTemp[KeyID] == nil {
        return  c.Status(fiber.StatusOK).JSON(NotFoundMsg{Message : "Not Found"})
      }
    
    s := new(TransactionStatusObj)
    if err := c.BodyParser(s); err != nil {
        return err
      }
    TransactionDataStatus[KeyID] = s.Status

    return  c.Status(fiber.StatusOK).JSON(TransactionStatusObj{Status : TransactionDataStatus[KeyID]})
    })

  app.Get("/payment", func(c *fiber.Ctx) error {
    KeyID := c.Query("id")
    if TransactionDataTemp[KeyID] == nil {
        return  c.Status(fiber.StatusOK).JSON(NotFoundMsg{Message : "Not Found"})
      }
    return  c.Status(fiber.StatusOK).JSON(TransactionDataTemp[KeyID])
    })

  app.Get("/payment/status", func(c *fiber.Ctx) error {
    KeyID := c.Query("id")
    if TransactionDataTemp[KeyID] == nil {
        return  c.Status(fiber.StatusOK).JSON(NotFoundMsg{Message : "Not Found"})
      }
    return  c.Status(fiber.StatusOK).JSON(TransactionStatusObj{Status : TransactionDataStatus[KeyID]})
    })

  app.Use(func(c *fiber.Ctx) error {
    return c.Status(fiber.StatusNotFound).JSON(NotFoundMsg{Message : "Not Found"})
  })
  app.Use(logger.New(logger.Config{
    Format:     "${pid} ${status} - ${method} ${path}\n",
    TimeFormat: "02-Jan-2006",
    TimeZone:   "Asia/Bangkok",
}))
  app.Listen(":8080")

}