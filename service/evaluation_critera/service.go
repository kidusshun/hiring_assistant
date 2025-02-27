package evaluationcritera

import (
	"errors"
	"log"
	"strconv"

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


func (s *Service) AddEvaluationCriteria(email string, payload CreateCriteriaPayload) ([]*EvaluationCriteria, error) {
	user, err := s.userStore.GetUserByEmail(email)
	if err != nil {
		log.Println("1", err)
		return nil, err
	}
	jobPosting, err := s.jobPostingStore.GetJobPostingByID(payload.JobPostingID)
	if err != nil {
		log.Println("2", err)
		return nil, err
	}
	if jobPosting.UserID != user.ID {
		log.Println("3", err)
		return nil, errors.New("user does not own the job posting")
	}
	
	var evlautionCriterias []*EvaluationCriteria
	
	for _, criteria := range payload.Criterias {
		f64, err := strconv.ParseFloat(criteria.Weight, 32)
		
		if err != nil {
			return nil, err
		}
		evaluationCriteria, err := s.criteriaStore.CreateEvaluationCriteria(payload.JobPostingID, criteria.CriteriaName, criteria.Description, float32(f64))
		if err != nil {
			log.Println("3", err)
			return nil, err
		}
		evlautionCriterias = append(evlautionCriterias, evaluationCriteria)
	}

	return evlautionCriterias, nil
}


func (s *Service) GetEvaluationCriteria(email string, payload GetEvaluationCriteriaPayload) ([]*EvaluationCriteria, error) {
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