package http

import (
	"fmt"
	"os/exec"

	"mctest-agent/internal/config"
	"mctest-agent/internal/controllers"
	"mctest-agent/internal/mappers"
	"mctest-agent/internal/services"

	"github.com/gin-gonic/gin"
)

const baseDSN = "%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4"

func Init() error {
	cmd := exec.Command("pip3", "install", "bandit")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to install bandit" + err.Error())
		return err
	}

	threshoulds := config.GetBanditThresholds()
	mapper := mappers.NewReportMapper()
	service := services.NewReportService(mapper, threshoulds)
	controller := controllers.NewReportController(service)

	r := gin.Default()

	r.POST("/mctest/validator/question/:id", controller.GenerateReport)

	r.Run(":8080")

	return nil
}
