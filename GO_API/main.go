package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type employe struct {
	ID           string `json:"id"`
	Nama         string `json:"nama"`
	Divisi       string `json:"divisi"`
	Kodekaryawan int    `json:"kodekaryawan"`
}

var employes = []employe{
	{ID: "1", Nama: "Ikbal Ardiansah", Divisi: "Web Developer", Kodekaryawan: 123},
	{ID: "2", Nama: "Ari Julianto", Divisi: "SPV Web Developer", Kodekaryawan: 456},
	{ID: "3", Nama: "Steavany Deasy", Divisi: "Admin Web Developer", Kodekaryawan: 789},
}

func getEmployes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employes)
}

func employeById(c *gin.Context) {
	id := c.Param("id")
	employe, err := getEmployesById(id)

	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, employe)
}

func CheckEmploye(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Tidak ada data Karyawan"})
		return
	}

	employe, err := getEmployesById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Tidak ada data Karyawan"})
		return
	}
	if employe.Kodekaryawan <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Tidak ada data Karyawan"})
		return
	}
	employe.Kodekaryawan -= 1
	c.IndentedJSON(http.StatusOK, employe)

}

func getEmployesById(id string) (*employe, error) {
	for i, b := range employes {
		if b.ID == id {
			return &employes[i], nil
		}
	}
	return nil, errors.New("Tidak ada data Karyawan")
}

func addEmployes(c *gin.Context) {
	var newEmployes employe

	if err := c.BindJSON(&newEmployes); err != nil {
		return
	}
	employes = append(employes, newEmployes)
	c.IndentedJSON(http.StatusCreated, newEmployes)
}

func main() {
	router := gin.Default()
	router.GET("/employes", getEmployes)
	router.GET("/employes/:id", employeById)
	router.POST("/employes", addEmployes)
	router.PATCH("/checkemploye", CheckEmploye)
	router.Run("localhost:8080")
}
