// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// EmployeesFilterEmploymentStatus - Employment status to filter on
type EmployeesFilterEmploymentStatus string

const (
	EmployeesFilterEmploymentStatusActive     EmployeesFilterEmploymentStatus = "active"
	EmployeesFilterEmploymentStatusInactive   EmployeesFilterEmploymentStatus = "inactive"
	EmployeesFilterEmploymentStatusTerminated EmployeesFilterEmploymentStatus = "terminated"
	EmployeesFilterEmploymentStatusOther      EmployeesFilterEmploymentStatus = "other"
)

func (e EmployeesFilterEmploymentStatus) ToPointer() *EmployeesFilterEmploymentStatus {
	return &e
}
func (e *EmployeesFilterEmploymentStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "inactive":
		fallthrough
	case "terminated":
		fallthrough
	case "other":
		*e = EmployeesFilterEmploymentStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EmployeesFilterEmploymentStatus: %v", v)
	}
}

type EmployeesFilter struct {
	// Company ID to filter on
	CompanyID *string `queryParam:"name=company_id"`
	// Email to filter on
	Email *string `queryParam:"name=email"`
	// First Name to filter on
	FirstName *string `queryParam:"name=first_name"`
	// Job title to filter on
	Title *string `queryParam:"name=title"`
	// Last Name to filter on
	LastName *string `queryParam:"name=last_name"`
	// Manager id to filter on
	ManagerID *string `queryParam:"name=manager_id"`
	// Employment status to filter on
	EmploymentStatus *EmployeesFilterEmploymentStatus `queryParam:"name=employment_status"`
	// Employee number to filter on
	EmployeeNumber *string `queryParam:"name=employee_number"`
	// ID of the department to filter on
	DepartmentID *string `queryParam:"name=department_id"`
	// City to filter on
	City *string `queryParam:"name=city"`
	// Country to filter on
	Country *string `queryParam:"name=country"`
}

func (o *EmployeesFilter) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *EmployeesFilter) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *EmployeesFilter) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *EmployeesFilter) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *EmployeesFilter) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *EmployeesFilter) GetManagerID() *string {
	if o == nil {
		return nil
	}
	return o.ManagerID
}

func (o *EmployeesFilter) GetEmploymentStatus() *EmployeesFilterEmploymentStatus {
	if o == nil {
		return nil
	}
	return o.EmploymentStatus
}

func (o *EmployeesFilter) GetEmployeeNumber() *string {
	if o == nil {
		return nil
	}
	return o.EmployeeNumber
}

func (o *EmployeesFilter) GetDepartmentID() *string {
	if o == nil {
		return nil
	}
	return o.DepartmentID
}

func (o *EmployeesFilter) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *EmployeesFilter) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}
