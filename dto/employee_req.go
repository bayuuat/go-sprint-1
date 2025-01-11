package dto

import (
	"errors"
	"net/url"
)

type EmployeeReq struct {
	IdentityNumber   string `json:"identityNumber" validate:"required,min=5,max=33"`
	Name             string `json:"name" validate:"required,min=4,max=33"`
	EmployeeImageUri string `json:"employeeImageUri" validate:"required,uri"`
	Gender           string `json:"gender" validate:"required,oneof=male female"`
	DepartmentID     string `json:"departmentId" validate:"required"`
}

type EmployeeFilter struct {
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
	Name           string `json:"name"`
	UserId         string `json:"user_id"`
	IdentityNumber string `json:"identity_number"`
	Gender         string `json:"gender"`
	DepartmentID   string `json:"department_id"`
}

func (req *EmployeeReq) Validate() (result map[string]interface{}, err error) {
	result = make(map[string]interface{})

	if result, err = req.ValidateIdentityNumber(result); err != nil {
		return nil, err
	}

	if result, err = req.ValidateName(result); err != nil {
		return nil, err
	}

	if result, err = req.ValidateEmployeeImageUri(result); err != nil {
		return nil, err
	}

	if result, err = req.ValidateGender(result); err != nil {
		return nil, err
	}

	return result, nil
}

func (req *EmployeeReq) ValidateIdentityNumber(temp map[string]interface{}) (result map[string]interface{}, err error) {
	result = temp
	if len(req.IdentityNumber) >= 5 && len(req.IdentityNumber) <= 33 {
		result["identity_number"] = req.IdentityNumber
	} else if req.IdentityNumber != "" {
		err = errors.New("bad request")
	}
	return

}

func (req *EmployeeReq) ValidateName(temp map[string]interface{}) (result map[string]interface{}, err error) {
	result = temp
	if len(req.Name) >= 4 && len(req.Name) <= 33 {
		result["name"] = req.Name
	} else if req.Name != "" {
		err = errors.New("bad request")
	}
	return
}

func (req *EmployeeReq) ValidateEmployeeImageUri(temp map[string]interface{}) (result map[string]interface{}, err error) {
	result = temp

	if req.EmployeeImageUri == "" {
		return result, nil
	}

	_, err = url.ParseRequestURI(req.EmployeeImageUri)
	if err == nil {
		result["employee_image_uri"] = req.EmployeeImageUri
	} else {
		err = errors.New("bad request")
	}
	return
}

func (req *EmployeeReq) ValidateGender(temp map[string]interface{}) (result map[string]interface{}, err error) {
	result = temp

	if req.Gender == "male" || req.Gender == "female" {
		result["gender"] = req.Gender
	} else if req.Gender != "" {
		err = errors.New("bad request")
	}
	return
}
