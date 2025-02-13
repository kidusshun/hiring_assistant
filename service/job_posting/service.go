package jobposting

import (
	"database/sql"

	"github.com/kidusshun/hiring_assistant/service/user"
)


type service struct {
	userStore user.UserStore
	jobPostingStore JobPostingStore
}


func NewService(userStore user.UserStore, jobPostingStore JobPostingStore) *service {
	return &service{
		userStore: userStore,
		jobPostingStore: jobPostingStore,
	}
}


func (s *service) CreateJobPosting(userEmail string, payload CreateJobPostingPayload) (*JobPosting, error) {
	user, err := s.userStore.GetUserByEmail(userEmail)

	if err != nil {
		return nil, err
	}

	jobPosting := &JobPosting{
		UserID: user.ID,
		Title: payload.Title,
		Description: payload.Description,
		Location: payload.Location,
		Department: payload.Department,
		EmploymentType: payload.EmploymentType,
	}

	createdJobPosting, err := s.jobPostingStore.CreateJobPosting(jobPosting)

	if err != nil {
		return nil, err
	}

	return createdJobPosting, nil
}


func (s *service) GetJobPostings(userEmail string, limit, offset int) ([]*JobPosting, error) {
	user, err := s.userStore.GetUserByEmail(userEmail)

	if err != nil {
		return nil, err
	}

	jobPostings, err := s.jobPostingStore.GetJobPostings(user.ID, limit, offset)

	if err != nil {
		if err == sql.ErrNoRows{
			return []*JobPosting{}, nil
		}
		return nil, err
	}
	return jobPostings, nil
}