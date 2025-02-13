package evaluationcritera

import (
	"time"

	"github.com/google/uuid"
)

type EvaluationCriteriaService interface {
	AddEvaluationCriteria(email string, payload []CreateCriteriaPayload) ([]*EvaluationCriteria, error)
	GetEvaluationCriteria(email string, payload GetEvaluationCriteriaPayload) ([]*EvaluationCriteria, error)
}
type EvaluationCriteriaStore interface {
	CreateEvaluationCriteria(jobPostingID uuid.UUID, criteriaName, description string, weight float32) (*EvaluationCriteria, error)
	GetEvaluationCriteriaByJobPostingID(jobPostingID uuid.UUID) ([]*EvaluationCriteria, error)
}


type EvaluationCriteria struct {
	ID          	int    		`json:"id"`
	JobPostingID 	uuid.UUID 	`json:"json_posting_id"`
	CriteriaName 	string 		`json:"criteria_name"`
	Description 	string 		`json:"description"`
	Weight			float32		`json:"weight"`
	CreatedAt   	time.Time 	`json:"created_at"`
	UpdatedAt   	time.Time 	`json:"updated_at"`
}

type CreateCriteriaPayload struct {
	JobPostingID 	uuid.UUID 	`json:"json_posting_id"`
	CriteriaName 	string 		`json:"criteria_name"`
	Description 	string 		`json:"description"`
	Weight			float32		`json:"weight"`
}

type GetEvaluationCriteriaPayload struct {
	JobPostingID 	uuid.UUID 	`json:"json_posting_id"`
}