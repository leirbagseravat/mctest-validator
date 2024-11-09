package services

import (
	"encoding/json"
	"fmt"
	"mctest-agent/internal/config"
	"mctest-agent/internal/mappers"
	"mctest-agent/internal/models/dto"
	"mctest-agent/internal/models/vo"
	"os/exec"
	"strings"
)

type reportService struct {
	mapper      mappers.ReportMapper
	threshoulds map[string]int
}

type ReportService interface {
	Generate(req *vo.Report) (res *dto.ReportResponse, err error)
}

func NewReportService(mapper mappers.ReportMapper, threshoulds *config.BanditThreshoulds) ReportService {
	return &reportService{mapper: mapper, threshoulds: buildThreshoulds(threshoulds)}
}

func buildThreshoulds(threshoulds *config.BanditThreshoulds) map[string]int {
	confidence, severity := threshoulds.Confidence, threshoulds.Severity
	return map[string]int{
		"HighSeverity":        severity.High,
		"MediumSeverity":      severity.Medium,
		"LowSeverity":         severity.Low,
		"UndefinedSeverity":   severity.Undefined,
		"HighConfidence":      confidence.High,
		"MediumConfidence":    confidence.Medium,
		"LowConfidence":       confidence.Low,
		"UndefinedConfidence": confidence.Undefined,
	}
}

func (s *reportService) Generate(req *vo.Report) (*dto.ReportResponse, error) {
	cmd := exec.Command("bandit", "-f", "json", "-r", req.File.Name())
	output, _ := cmd.Output()
	var report dto.BanditReportResponse
	err := json.Unmarshal(output, &report)
	if err != nil {
		fmt.Println("falhou unmarshal", err)
		return nil, err
	}

	message, status := s.validate(report)

	return s.mapper.MapToReportResponse(req.ID, message, status, report), nil
}

func (s *reportService) validate(report dto.BanditReportResponse) (string, string) {
	results := report.Metrics.Totals

	validationMessages := make([]string, 0)
	if !s.validateByFields(results.GetConfidenceFields(), s.threshoulds) {
		validationMessages = append(validationMessages, "confidence exceeds the threshoulds")
	}

	if !s.validateByFields(results.GetSeverityFields(), s.threshoulds) {
		validationMessages = append(validationMessages, "severity exceeds the threshoulds")
	}

	if len(report.Errors) > 0 {
		validationMessages = append(validationMessages, "your code has error messages")
	}

	if len(validationMessages) > 0 {
		return strings.Join(validationMessages, " and "), "ERROR"
	}

	return "Report Passed", "PASSED"
}

func (s *reportService) validateByFields(fields, threshoulds map[string]int) bool {
	for key, field := range fields {
		if field > threshoulds[key] {
			return false
		}
	}
	return true
}
