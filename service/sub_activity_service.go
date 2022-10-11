package service

import (
	"activity/models/domain"
	"activity/models/web"
	"activity/repository"

	"github.com/mashingan/smapping"
)

type SubActivityService interface {
	All() []domain.SubActivity
	Create(request web.SubActivityRequest) (domain.SubActivity, error)
	FindById(id uint) (domain.SubActivity, error)
	Update(request web.SubActivityUpdateRequest) (domain.SubActivity, error)
	Delete(id uint) error
}

type subactivityService struct {
	subactivityRepository repository.SubActivityRepository
}

func NewSubActivityService(subactivityRepo repository.SubActivityRepository) SubActivityService {
	return &subactivityService{subactivityRepository: subactivityRepo}
}

func (sas *subactivityService) All() []domain.SubActivity {
	return sas.subactivityRepository.All()
}

func (sas *subactivityService) Create(request web.SubActivityRequest) (domain.SubActivity, error) {
	subactivity := domain.SubActivity{}

	err := smapping.FillStruct(&subactivity, smapping.MapFields(&request))
	if err != nil {

		return subactivity, err
	}
	_, err = sas.subactivityRepository.IsDReportExist(request.ActivityId)
	if err != nil {
		return subactivity, err
	}

	return sas.subactivityRepository.Create(subactivity), nil
}

func (sas *subactivityService) Update(request web.SubActivityUpdateRequest) (domain.SubActivity, error) {
	subactivity := domain.SubActivity{}
	res, err := sas.subactivityRepository.FindById(request.ID)
	if err != nil {
		return subactivity, err
	}
	err = smapping.FillStruct(&subactivity, smapping.MapFields(&request))
	if err != nil {
		return subactivity, err
	}
	subactivity.User_id = res.User_id

	return sas.subactivityRepository.Update(subactivity), nil
}

func (sas *subactivityService) FindById(id uint) (domain.SubActivity, error) {
	subactivity, err := sas.subactivityRepository.FindById(id)
	if err != nil {
		return subactivity, err
	}
	return subactivity, nil
}

func (sas *subactivityService) Delete(id uint) error {
	subactivity, err := sas.subactivityRepository.FindById(id)
	if err != nil {
		return err
	}
	sas.subactivityRepository.Delete(subactivity)
	return nil
}
