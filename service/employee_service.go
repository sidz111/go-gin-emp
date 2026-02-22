package service

import (
	"context"
	"fmt"

	"github.com/sidz111/go-gin-emp-crud/model"
	"github.com/sidz111/go-gin-emp-crud/repository"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, emp *model.Employee) error {
	err := s.ValidateEmployee(emp)
	if err != nil {
		return err
	}
	if err := s.repo.Create(ctx, emp); err != nil {
		return fmt.Errorf("service: failed to create employee %w", err)
	}
	return nil
}

func (s *EmployeeService) GetEmployeeById(ctx context.Context, id int) (*model.Employee, error) {
	if id <= 0 {
		return nil, fmt.Errorf("id must be positive")
	}
	emp, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to fetch employee %w", err)
	}
	return emp, nil
}

func (s *EmployeeService) GetAllEmployees(ctx context.Context) ([]model.Employee, error) {
	employees, err := s.repo.GetAllEmployees(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to fetch employees %w", err)
	}
	return employees, nil
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, emp *model.Employee) error {
	if emp.ID <= 0 {
		return fmt.Errorf("invalid employee id")
	}
	if err := s.ValidateEmployee(emp); err != nil {
		return fmt.Errorf("invalid data %w", err)
	}
	if err := s.repo.UpdateEmployee(ctx, emp); err != nil {
		return fmt.Errorf("service: failed to update employee %w", err)
	}
	return nil
}

func (s *EmployeeService) DeleteEmployeeById(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("id must be positive")
	}
	if err := s.repo.DeleteEmployeeByID(ctx, id); err != nil {
		return fmt.Errorf("service: failed to delete employee %w", err)
	}
	return nil
}

func (s *EmployeeService) ValidateEmployee(emp *model.Employee) error {
	if emp.Name == "" {
		return fmt.Errorf("name required")
	}
	if emp.Address == "" {
		return fmt.Errorf("address required")
	}
	if emp.Salary <= 0 {
		return fmt.Errorf("salary must be Positive")
	}
	return nil
}
