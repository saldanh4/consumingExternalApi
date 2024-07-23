package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct que recebe todos os campos retornados da API externa
type ExternalApiLayout struct {
	Abbreviation string `json:"abbreviation"`
	Client_ip    string `json:"client_ip"`
	Datetime     string `json:"datetime"`
	Day_of_week  int    `json:"day_of_week"`
	Day_of_year  int    `json:"day_of_year"`
	Dst          bool   `json:"dst"`
	Dst_from     string `json:"dst_from"`
	Dst_offset   int    `json:"dst_offset"`
	Dst_until    string `json:"dst_until"`
	Raw_offset   int64  `json:"raw_offset"`
	Timezone     string `json:"timezone"`
	Unixtime     int64  `json:"unixtime"`
	Utc_datetime string `json:"utc_datetime"`
	Utc_offset   string `json:"utc_offset"`
	Week_number  int8   `json:"week_number"`
}

// Função que faz a chamada Get na API externa, trata os dados e retorna a hora de São Paulo
func getExternalApi(c *gin.Context) {
	var horaSp ExternalApiLayout
	url := "https://worldtimeapi.org/api/timezone/America/Sao_Paulo"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&horaSp)
	if err != nil {
		fmt.Println("Erro ao coletar a resposta: ", err)
		return
	}

	c.IndentedJSON(http.StatusOK, horaSp.Datetime)
}

func main() {
	router := gin.Default()
	router.GET("/", getExternalApi)

	router.Run(":8080")

}
