package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", getCep)

	e.Logger.Fatal(e.Start(":8888"))
}

type cepResponse struct {
	Cep        string `json:"cep"`
	Uf         string `json:"uf"`
	Cidade     string `json:"cidade"`
	Bairro     string `json:"bairro"`
	Logradouro string `json:"logradouro"`
}

func getCep(c echo.Context) error {

	res := cepResponse{
		Cep:        "65081264",
		Uf:         "MA",
		Cidade:     "São Luís",
		Bairro:     "Vila Embratel",
		Logradouro: "2ª Travessa Seis de Abril",
	}

	return c.JSON(http.StatusOK, res)
}
