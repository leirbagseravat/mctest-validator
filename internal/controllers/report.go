package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"mctest-agent/internal/models/vo"
	"mctest-agent/internal/services"

	"github.com/gin-gonic/gin"
)

type reportController struct {
	service services.ReportService
}

type ReportController interface {
	GenerateReport(c *gin.Context)
}

func NewReportController(service services.ReportService) ReportController {
	return &reportController{service: service}
}

func (controller *reportController) GenerateReport(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Print(" morre arquivo")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tempFile, err := os.CreateTemp("", "temp-*.py")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = io.Copy(tempFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = file.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := os.Chmod(tempFile.Name(), 0700); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falhou ao dar permissão de execução."})
		return
	}

	res, err := controller.service.Generate(&vo.Report{
		ID:        c.Param("id"),
		CreatedAt: time.Now(),
		File:      tempFile,
	})

	if err != nil {
		fileErr := os.RemoveAll(tempFile.Name())
		if fileErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = os.RemoveAll(tempFile.Name())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(res)
	c.JSON(http.StatusCreated, gin.H{"result": res})
}
