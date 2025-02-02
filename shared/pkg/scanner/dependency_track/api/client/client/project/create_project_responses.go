// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cisco-open/kubei/shared/pkg/scanner/dependency_track/api/client/models"
)

// CreateProjectReader is a Reader for the CreateProject structure.
type CreateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateProjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProjectCreated creates a CreateProjectCreated with default headers values
func NewCreateProjectCreated() *CreateProjectCreated {
	return &CreateProjectCreated{}
}

/* CreateProjectCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateProjectCreated struct {
	Payload *models.Project
}

func (o *CreateProjectCreated) Error() string {
	return fmt.Sprintf("[PUT /v1/project][%d] createProjectCreated  %+v", 201, o.Payload)
}
func (o *CreateProjectCreated) GetPayload() *models.Project {
	return o.Payload
}

func (o *CreateProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectUnauthorized creates a CreateProjectUnauthorized with default headers values
func NewCreateProjectUnauthorized() *CreateProjectUnauthorized {
	return &CreateProjectUnauthorized{}
}

/* CreateProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateProjectUnauthorized struct {
}

func (o *CreateProjectUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/project][%d] createProjectUnauthorized ", 401)
}

func (o *CreateProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectConflict creates a CreateProjectConflict with default headers values
func NewCreateProjectConflict() *CreateProjectConflict {
	return &CreateProjectConflict{}
}

/* CreateProjectConflict describes a response with status code 409, with default header values.

A project with the specified name already exists
*/
type CreateProjectConflict struct {
}

func (o *CreateProjectConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/project][%d] createProjectConflict ", 409)
}

func (o *CreateProjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
