package dto

import "time"

type CreateReportRequest struct {
	ID        string `form:"id" binding:"required"`
	UpdatedAt string `form:"updated_at" binding:"required"`
}

type BanditReportResponse struct {
	Errors      []interface{}  `json:"errors"`
	GeneratedAt *time.Time     `json:"generated_at"`
	Metrics     MetricsRequest `json:"metrics"`
	Results     []interface{}  `json:"results"`
}

type MetricsRequest struct {
	Totals TotalsRequest `json:"_totals"`
}

type TotalsRequest struct {
	HighConfidence      int `json:"CONFIDENCE.HIGH"`
	MediumConfidence    int `json:"CONFIDENCE.MEDIUM"`
	LowConfidence       int `json:"CONFIDENCE.LOW"`
	UndefinedConfidence int `json:"CONFIDENCE.UNDEFINED"`
	HighSeverity        int `json:"SEVERITY.HIGH"`
	MediumSeverity      int `json:"SEVERITY.MEDIUM"`
	LowSeverity         int `json:"SEVERITY.LOW"`
	UndefinedSeverity   int `json:"SEVERITY.UNDEFINED"`
	Loc                 int `json:"loc"`
	Nosec               int `json:"nosec"`
	SkippedTests        int `json:"skipped_tests"`
}

type ReportResponse struct {
	QuestionID   string               `json: "question_id"`
	CreatedAt    string               `json: "created_at"`
	Message      string               `json:"message"`
	Status       string               `json:"status"`
	BanditReport BanditReportResponse `json:"report_response"`
}

func (t *TotalsRequest) GetConfidenceFields() map[string]int {
	return map[string]int{
		"HighConfidence":      t.HighConfidence,
		"MediumConfidence":    t.MediumConfidence,
		"LowConfidence":       t.LowConfidence,
		"UndefinedConfidence": t.UndefinedConfidence,
	}
}

func (t *TotalsRequest) GetSeverityFields() map[string]int {
	return map[string]int{
		"HighSeverity":      t.HighSeverity,
		"MediumSeverity":    t.MediumSeverity,
		"LowSeverity":       t.LowSeverity,
		"UndefinedSeverity": t.UndefinedSeverity,
	}
}
