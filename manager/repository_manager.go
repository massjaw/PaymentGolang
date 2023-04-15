package manager

import "Merchant-Bank/repository"

type RepoManager interface {
	UserRepo() repository.UserRepo
}

type repositoryManager struct {
	infraManager InfraManager
}

func (c *repositoryManager) UserRepo() repository.UserRepo {
	return repository.NewUserRepository(c.infraManager.DbConn())
}

func NewRepoManager(manager InfraManager) RepoManager {
	return &repositoryManager{
		infraManager: manager,
	}
}
