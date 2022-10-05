package service

import (
	"activity/models/domain"
	"activity/models/web"
	"activity/repository"
	"fmt"

	"github.com/mashingan/smapping"
)

type ActivityService interface {
	All() []domain.Activity
	Create(b web.ActivityRequest) (domain.Activity, error)
	FindById(id uint) (domain.Activity, error)
	Update(b web.ActivityUpdateRequest) (domain.Activity, error)
	Delete(id uint) error
}

type activityService struct {
	activityRepository repository.ActivityRepository
}

func NewActivityService(activityRepository repository.ActivityRepository) ActivityService {
	return &activityService{activityRepository: activityRepository}
}

func (s *activityService) All() []domain.Activity {
	return s.activityRepository.All()
}

func (s *activityService) Create(request web.ActivityRequest) (domain.Activity, error) {
	activity := domain.Activity{}

	//time.Date(request.StartDate.Year(), request.StartDate.Month(), request.StartDate.Day(), 0, 0, 0, 0, time.Local)
	//projectcharter.StartDate = utils.ConvertDate(request.StartDate)
	//projectcharter.EndDate = utils.ConvertDate(request.EndDate)
	//request.StartDate = nil
	//request.EndDate = nil

	err := smapping.FillStruct(&activity, smapping.MapFields(&request))
	if err != nil {

		return activity, err
	}
	fmt.Println(activity)
	// _, err = s.projectcharterRepository.IsDuplicateEmail(request.Email)
	// if err != nil {
	// 	return projectcharter, err
	// }
	return s.activityRepository.Create(activity), nil
}

func (s *activityService) Update(b web.ActivityUpdateRequest) (domain.Activity, error) {
	activity := domain.Activity{}
	_, err := s.activityRepository.FindById(b.ID)
	if err != nil {
		return activity, err
	}
	err = smapping.FillStruct(&activity, smapping.MapFields(&b))
	if err != nil {
		return activity, err
	}
	//projectcharter.ID = res.ID

	return s.activityRepository.Update(activity), nil
}

func (s *activityService) FindById(id uint) (domain.Activity, error) {
	activity, err := s.activityRepository.FindById(id)
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (s *activityService) Delete(id uint) error {
	activity, err := s.activityRepository.FindById(id)
	if err != nil {
		return err
	}
	s.activityRepository.Delete(activity)
	return nil
}
