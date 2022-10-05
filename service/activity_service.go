package service

import (
	"activity/models/domain"
	"activity/models/web"
	"activity/repository"

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

	err := smapping.FillStruct(&activity, smapping.MapFields(&request))
	if err != nil {

		return activity, err
	}

	return s.activityRepository.Create(activity), nil
}

func (s *activityService) Update(b web.ActivityUpdateRequest) (domain.Activity, error) {
	activity := domain.Activity{}
	res, err := s.activityRepository.FindById(b.ID)
	if err != nil {
		return activity, err
	}
	err = smapping.FillStruct(&activity, smapping.MapFields(&b))
	if err != nil {
		return activity, err
	}
	activity.User_id = res.User_id

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
