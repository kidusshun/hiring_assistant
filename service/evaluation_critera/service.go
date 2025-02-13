package evaluationcritera

import (
	"errors"

	jobposting "github.com/kidusshun/hiring_assistant/service/job_posting"
	"github.com/kidusshun/hiring_assistant/service/user"
)

type Service struct {
	criteriaStore EvaluationCriteriaStore
	userStore user.UserStore
	jobPostingStore jobposting.JobPostingStore
}


func NewService(criteriaStore EvaluationCriteriaStore, userStore user.UserStore, jobPostingStore jobposting.JobPostingStore) *Service {
	return &Service{
		criteriaStore: criteriaStore,
		userStore: userStore,
		jobPostingStore: jobPostingStore,
	}
}


func (s *Service) AddEvaluationCriteria(email string, payload []CreateCriteriaPayload) ([]*EvaluationCriteria, error) {
	user, err := s.userStore.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	var evlautionCriterias []*EvaluationCriteria

	for _, criteria := range payload {
		jobPosting, err := s.jobPostingStore.GetJobPostingByID(criteria.JobPostingID)
		if err != nil {
			return nil, err
		}
		if jobPosting.UserID != user.ID {
			return nil, errors.New("user does not own the job posting")
		}
		evaluationCriteria, err := s.criteriaStore.CreateEvaluationCriteria(criteria.JobPostingID, criteria.CriteriaName, criteria.Description, criteria.Weight)
		if err != nil {
			return nil, err
		}
		evlautionCriterias = append(evlautionCriterias, evaluationCriteria)
	}

	return evlautionCriterias, nil
}


func (s *Service) GetEvaluationCritera(email string, payload GetEvaluationCriteriaPayload) ([]*EvaluationCriteria, error) {
	user, err := s.userStore.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	jobPosting, err := s.jobPostingStore.GetJobPostingByID(payload.JobPostingID)

	if err != nil {
		return nil, err
	}

	if jobPosting.UserID != user.ID {
		return nil, errors.New("user does not own the job posting")
	}

	evaluationCriterias, err := s.criteriaStore.GetEvaluationCriteriaByJobPostingID(payload.JobPostingID)
	if err != nil {
		return nil, err
	}

	return evaluationCriterias, nil
}