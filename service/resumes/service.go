package resumes

import (
	"errors"
	"log"

	"github.com/google/uuid"
	jobposting "github.com/kidusshun/hiring_assistant/service/job_posting"
	"github.com/kidusshun/hiring_assistant/service/llmclient"
	"github.com/kidusshun/hiring_assistant/service/user"
)


type service struct {
	userStore user.UserStore
	jobPostingStore jobposting.JobPostingStore
	resumeStore ResumeStore
}


func NewService(resumeStore ResumeStore, userStore user.UserStore, jobPostingStore jobposting.JobPostingStore) *service {
	return &service{
		resumeStore: resumeStore,
		userStore: userStore,
		jobPostingStore: jobPostingStore,
	}
}


func (s *service) StoreResumeService(userEmail string, resumes CreateResumesPayload) ([]Resume, error) {
	storedUser, err := s.userStore.GetUserByEmail(userEmail)

	if err != nil {
		return nil, err
	}
	storedJobPosting, err := s.jobPostingStore.GetJobPostingByID(resumes.JobPostingID)
	if err != nil {
		return nil, err
	}

	if storedJobPosting.UserID != storedUser.ID {
		return nil, errors.New("user does not own the job posting")
	}

	processedResumes, err := processResumes(resumes.Resumes, resumes.JobPostingID)

	if err != nil {
		return nil, err
	}


	storedResumes, err := s.resumeStore.AddResumes(processedResumes)
	if err != nil {
		return nil, err
	}

	return storedResumes, nil
}


func processResumes(resumes []JobResumePayload, jobPostingID uuid.UUID) ([]Resume, error) {
	var processedResumes []Resume
	for _, resume := range resumes {

		applicant, err := llmclient.ParseEmailAndName(resume.Text)
		if err != nil {
			log.Println("llm parse email name error", err)
			return nil, err
		}
		resume := Resume{
			JobPostingID: jobPostingID,
			ApplicantEmail: applicant.Email,
			ApplicantName: applicant.Name,
			ResumePath: resume.URL,
			ResumeText: resume.Text,
			Status: StatusPending,
		}
		processedResumes = append(processedResumes, resume)
	}

	return processedResumes, nil
}


