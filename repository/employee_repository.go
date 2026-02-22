package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sidz111/go-gin-emp-crud/model"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(ctx context.Context, emp *model.Employee) error {
	query := "INSERT INTO emp(name, address, salary) values(?,?,?)"
	result, err := r.db.ExecContext(ctx, query, emp.Name, emp.Address, emp.Salary)
	if err != nil {
		return fmt.Errorf("Failed to Create Employee %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Failed to Save Employee %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("failed to insert a row")
	}
	return nil
}

func (r *EmployeeRepository) GetById(ctx context.Context, id int) (*model.Employee, error) {
	query := "SELECT id, name, address, salary FROM emp WHERE id=?"
	var emp model.Employee
	err := r.db.QueryRowContext(ctx, query, id).Scan(&emp.ID, &emp.Name, &emp.Address, &emp.Salary)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("employee Not Found")
		}
		return nil, fmt.Errorf("failed to fetch employee %w", err)
	}
	return &emp, nil
}

func (r *EmployeeRepository) GetAllEmployees(ctx context.Context) ([]model.Employee, error) {
	query := "SELECT id, name, address, salary FROM emp"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch employees %w", err)
	}
	defer rows.Close()
	var employees []model.Employee
	for rows.Next() {
		var emp model.Employee
		err := rows.Scan(&emp.ID, &emp.Name, &emp.Address, &emp.Salary)
		if err != nil {
			return nil, fmt.Errorf("failed to scanned employee %w", err)
		}
		employees = append(employees, emp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while itterating rows %w", err)
	}
	return employees, nil
}

func (r *EmployeeRepository) UpdateEmployee(ctx context.Context, emp *model.Employee) error {
	query := "UPDATE emp SET name=?, address=?, salary=? WHERE id=?"
	result, err := r.db.ExecContext(ctx, query, emp.Name, emp.Address, emp.Salary, emp.ID)
	if err != nil {
		return fmt.Errorf("failed to Execute Update employee %w", err)
	}
	row, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to update employee %w", err)
	}

	if row == 0 {
		return fmt.Errorf("invalid employee id")
	}
	return nil
}

func (r *EmployeeRepository) DeleteEmployeeByID(ctx context.Context, id int) error {
	query := "DELETE FROM emp WHERE id=?"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete employee %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected row in employee %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("employee not found")
	}
	return nil
}
