package manager

import "Merchant-Bank/repository"

type RepoManager interface {
	UserRepo() repository.UserRepo
	PaymentRepo() repository.PaymentRepo
	WalletRepo() repository.WalletRepo
}

type repositoryManager struct {
	infraManager InfraManager
}

func (c *repositoryManager) UserRepo() repository.UserRepo {
	return repository.NewUserRepository(c.infraManager.DbConn())
}

func (c *repositoryManager) PaymentRepo() repository.PaymentRepo {
	return repository.NewPaymentRepository(c.infraManager.DbConn())
}

func (c *repositoryManager) WalletRepo() repository.WalletRepo {
	return repository.NewWalletRepository(c.infraManager.DbConn())
}

func NewRepoManager(manager InfraManager) RepoManager {
	return &repositoryManager{
		infraManager: manager,
	}
}
