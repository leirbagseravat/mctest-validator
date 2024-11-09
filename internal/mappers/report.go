package mappers

import "mctest-agent/internal/models/dto"

type ReportMapper interface {
	MapToReportResponse(id string, message, status string, report dto.BanditReportResponse) *dto.ReportResponse
}

type reportMapper struct{}

func NewReportMapper() ReportMapper {
	return &reportMapper{}
}

func (*reportMapper) MapToReportResponse(id string, message, status string, report dto.BanditReportResponse) *dto.ReportResponse {
	return &dto.ReportResponse{
		QuestionID:   id,
		CreatedAt:    report.GeneratedAt.String(),
		Message:      message,
		Status:       status,
		BanditReport: report,
	}
}
