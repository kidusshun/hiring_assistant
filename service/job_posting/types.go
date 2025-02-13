package jobposting

import (
	"time"

	"github.com/google/uuid"
)

type JobPostingService interface {
	CreateJobPosting(userEmail string, payload CreateJobPostingPayload) (*JobPosting, error)
	GetJobPostings(userEmail string, limit, offset int) ([]*JobPosting, error)
}
type JobPostingStore interface {
	CreateJobPosting(jobPosting *JobPosting) (*JobPosting, error)
	GetJobPostings(userID uuid.UUID, limit, offset int) ([]*JobPosting, error)
}

type JobPosting struct {
	ID 				uuid.UUID 	`json:"id"`
	UserID 			uuid.UUID 	`json:"user_id"`
	Title 			string 		`json:"title"`
	Description 	string 		`json:"description"`
	Location 		string 		`json:"location"`
	Department 		string 		`json:"department"`
	EmploymentType 	string 		`json:"employment_type"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
}

type CreateJobPostingPayload struct {
	Title 			string `json:"title"`
	Description 	string `json:"description"`
	Location 		string `json:"location"`
	Department 		string `json:"department"`
	EmploymentType 	string `json:"employment_type"`
}