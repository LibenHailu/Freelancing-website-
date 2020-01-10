package service

import (
	"fmt"

	"github.com/LibenHailu/fjobs/api/entity"
	"github.com/LibenHailu/fjobs/api/user"
)

type UserService struct {
	userRepo user.UserRepository
}

func NewUserService(repo user.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}
func (us *UserService) Users() ([]entity.User, []error) {
	users, errs := us.userRepo.Users()
	if len(errs) > 0 {
		return nil, errs
	}

	return users, nil
}

func (us *UserService) User(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.User(user)
	fmt.Println(errs)
	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// UserByID returns a single application user by its id
func (us *UserService) UserByID(id uint) (*entity.User, []error) {
	usr, errs := us.userRepo.UserByID(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// UpdateUser updates application user
func (us *UserService) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.UpdateUser(user)

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// DeleteUser deletes a single application user
func (us *UserService) DeleteUser(id uint) (*entity.User, []error) {
	usr, errs := us.userRepo.DeleteUser(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// StoreUser will insert a new application user
func (us *UserService) StoreUser(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.StoreUser(user)
	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

func (us *UserService) RecommendedJobs(id uint) ([]entity.Job, []error) {

	job, errs := us.userRepo.RecommendedJobs(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return job, nil
}
