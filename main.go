package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	bearer string
	port   = flag.String("port", ":3000", "Port to listen on")
	prod   = flag.Bool("prod", false, "Enable prefork in Production")
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getUsers(location string) []byte {
	URL := "https://api.github.com/search/users?q=location:" + location

	req, req_err := http.NewRequest("GET", URL, nil)

	if req_err != nil {
		fmt.Println(req_err)
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, client_err := client.Do(req)
	if client_err != nil {
		fmt.Println(client_err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func getRepos(language string) []byte {
	URL := "https://api.github.com/search/repositories?q=language:" + language

	req, req_err := http.NewRequest("GET", URL, nil)

	if req_err != nil {
		fmt.Println(req_err)
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, client_err := client.Do(req)
	if client_err != nil {
		fmt.Println(client_err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func main() {
	flag.Parse()
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowMethods:  "GET",
		ExposeHeaders: "Content-Type,Authorization,Accept",
	}))
	bearer = "Bearer " + getEnv("ACCESS_TOKEN", "")

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Github Search")
	})

	v1.Get("/users", func(c *fiber.Ctx) error {
		c.Set("Content-type", "application/json; charset=utf-8")
		return c.Send(getUsers(""))
	})

	v1.Get("/users/location::location?", func(c *fiber.Ctx) error {
		c.Set("Content-type", "application/json; charset=utf-8")
		return c.Send(getUsers(c.Params("location")))
	})

	v1.Get("/repos/:language", func(c *fiber.Ctx) error {
		c.Set("Content-type", "application/json; charset=utf-8")
		return c.Send(getRepos(c.Params("language")))
	})

	app.Listen(":3000")
}
