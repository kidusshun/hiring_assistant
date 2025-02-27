package resumes

import (
	"time"

	"github.com/google/uuid"
)

type ResumeService interface {
	StoreResumeService(userEmail string, request CreateResumesPayload) ([]Resume, error)
	GetResumesService(userEmail string, jobPostingID uuid.UUID) ([]Resume, error)
}

type ResumeStore interface {
	AddResumes(resume []Resume) ([]Resume, error)
	GetResumesByJobPostingID(jobPostingID uuid.UUID) ([]Resume, error)
}

type ResumeStatus string

const (
	StatusPending   ResumeStatus = "pending"
	StatusProcessed ResumeStatus = "processed"
)

type Resume struct {
	ID            	uuid.UUID    	`json:"id"`
	JobPostingID  	uuid.UUID    	`json:"job_posting_id"`
	ApplicantName 	string       	`json:"applicant_name"`
	ApplicantEmail 	string      	`json:"applicant_email"`
	ResumePath    	string       	`json:"resume_file_path"`
	ResumeText		string       	`json:"resume_text"`
	Status        	ResumeStatus 	`json:"status"`
	CreatedAt 		time.Time 		`json:"created_at"`
	UpdatedAt 		time.Time		`json:"updated_at"`
}


type JobResumePayload struct {
	URL          string    `json:"url" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	MimeType     string    `json:"mimeType" validate:"required"`
	Text 		 string	   `json:"text"`
}

type CreateResumesPayload struct {
	JobPostingID uuid.UUID `json:"job_posting_id" validate:"required"`
	Resumes      []JobResumePayload `json:"resumes" validate:"required"`
}