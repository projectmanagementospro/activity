package service

import (
	"activity/models/domain"
	"activity/models/web"
	"activity/repository"

	"github.com/mashingan/smapping"
)

type ActivityService interface {
	All() []domain.Activity
	Create(request web.ActivityRequest) (domain.Activity, error)
	FindById(id uint) (domain.Activity, error)
	Update(request web.ActivityUpdateRequest) (domain.Activity, error)
	Delete(id uint) error
}

type activityService struct {
	activityRepository repository.ActivityRepository
}

func NewActivityService(activityRepo repository.ActivityRepository) ActivityService {
	return &activityService{activityRepository: activityRepo}
}

func (as *activityService) All() []domain.Activity {
	return as.activityRepository.All()
}

func (as *activityService) Create(request web.ActivityRequest) (domain.Activity, error) {
	activity := domain.Activity{}

	err := smapping.FillStruct(&activity, smapping.MapFields(&request))
	if err != nil {

		return activity, err
	}

	return as.activityRepository.Create(activity), nil
}

func (as *activityService) Update(request web.ActivityUpdateRequest) (domain.Activity, error) {
	activity := domain.Activity{}
	res, err := as.activityRepository.FindById(request.ID)
	if err != nil {
		return activity, err
	}
	err = smapping.FillStruct(&activity, smapping.MapFields(&request))
	if err != nil {
		return activity, err
	}
	activity.User_id = res.User_id

	return as.activityRepository.Update(activity), nil
}

func (as *activityService) FindById(id uint) (domain.Activity, error) {
	activity, err := as.activityRepository.FindById(id)
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (as *activityService) Delete(id uint) error {
	activity, err := as.activityRepository.FindById(id)
	if err != nil {
		return err
	}
	as.activityRepository.Delete(activity)
	return nil
}
