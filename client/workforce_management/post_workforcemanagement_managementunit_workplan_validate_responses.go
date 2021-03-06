// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostWorkforcemanagementManagementunitWorkplanValidateReader is a Reader for the PostWorkforcemanagementManagementunitWorkplanValidate structure.
type PostWorkforcemanagementManagementunitWorkplanValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementManagementunitWorkplanValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementManagementunitWorkplanValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementManagementunitWorkplanValidateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementManagementunitWorkplanValidateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementManagementunitWorkplanValidateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementManagementunitWorkplanValidateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementManagementunitWorkplanValidateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementManagementunitWorkplanValidateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementManagementunitWorkplanValidateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementManagementunitWorkplanValidateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementManagementunitWorkplanValidateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementManagementunitWorkplanValidateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementManagementunitWorkplanValidateOK creates a PostWorkforcemanagementManagementunitWorkplanValidateOK with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanValidateOK() *PostWorkforcemanagementManagementunitWorkplanValidateOK {
	return &PostWorkforcemanagementManagementunitWorkplanValidateOK{}
}

/*PostWorkforcemanagementManagementunitWorkplanValidateOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementManagementunitWorkplanValidateOK struct {
	Payload *models.ValidateWorkPlanResponse
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate][%d] postWorkforcemanagementManagementunitWorkplanValidateOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateOK) GetPayload() *models.ValidateWorkPlanResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidateWorkPlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanValidateBadRequest creates a PostWorkforcemanagementManagementunitWorkplanValidateBadRequest with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanValidateBadRequest() *PostWorkforcemanagementManagementunitWorkplanValidateBadRequest {
	return &PostWorkforcemanagementManagementunitWorkplanValidateBadRequest{}
}

/*PostWorkforcemanagementManagementunitWorkplanValidateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementManagementunitWorkplanValidateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate][%d] postWorkforcemanagementManagementunitWorkplanValidateBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanValidateUnauthorized creates a PostWorkforcemanagementManagementunitWorkplanValidateUnauthorized with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanValidateUnauthorized() *PostWorkforcemanagementManagementunitWorkplanValidateUnauthorized {
	return &PostWorkforcemanagementManagementunitWorkplanValidateUnauthorized{}
}

/*PostWorkforcemanagementManagementunitWorkplanValidateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementManagementunitWorkplanValidateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate][%d] postWorkforcemanagementManagementunitWorkplanValidateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanValidateForbidden creates a PostWorkforcemanagementManagementunitWorkplanValidateForbidden with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanValidateForbidden() *PostWorkforcemanagementManagementunitWorkplanValidateForbidden {
	return &PostWorkforcemanagementManagementunitWorkplanValidateForbidden{}
}

/*PostWorkforcemanagementManagementunitWorkplanValidateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementManagementunitWorkplanValidateForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate][%d] postWorkforcemanagementManagementunitWorkplanValidateForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanValidateNotFound creates a PostWorkforcemanagementManagementunitWorkplanValidateNotFound with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanValidateNotFound() *PostWorkforcemanagementManagementunitWorkplanValidateNotFound {
	return &PostWorkforcemanagementManagementunitWorkplanValidateNotFound{}
}

/*PostWorkforcemanagementManagementunitWorkplanValidateNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementManagementunitWorkplanValidateNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate][%d] postWorkforcemanagementManagementunitWorkplanValidateNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanValidateRequestEntityTooLarge creates a PostWorkforcemanagementManagementunitWorkplanValidateRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanValidateRequestEntityTooLarge() *PostWorkforcemanagementManagementunitWorkplanValidateRequestEntityTooLarge {
	return &PostWorkforcemanagementManagementunitWorkplanValidateRequestEntityTooLarge{}
}

/*PostWorkforcemanagementManagementunitWorkplanValidateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostWorkforcemanagementManagementunitWorkplanValidateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate][%d] postWorkforcemanagementManagementunitWorkplanValidateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanValidateUnsupportedMediaType creates a PostWorkforcemanagementManagementunitWorkplanValidateUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanValidateUnsupportedMediaType() *PostWorkforcemanagementManagementunitWorkplanValidateUnsupportedMediaType {
	return &PostWorkforcemanagementManagementunitWorkplanValidateUnsupportedMediaType{}
}

/*PostWorkforcemanagementManagementunitWorkplanValidateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementManagementunitWorkplanValidateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate][%d] postWorkforcemanagementManagementunitWorkplanValidateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanValidateTooManyRequests creates a PostWorkforcemanagementManagementunitWorkplanValidateTooManyRequests with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanValidateTooManyRequests() *PostWorkforcemanagementManagementunitWorkplanValidateTooManyRequests {
	return &PostWorkforcemanagementManagementunitWorkplanValidateTooManyRequests{}
}

/*PostWorkforcemanagementManagementunitWorkplanValidateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostWorkforcemanagementManagementunitWorkplanValidateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate][%d] postWorkforcemanagementManagementunitWorkplanValidateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanValidateInternalServerError creates a PostWorkforcemanagementManagementunitWorkplanValidateInternalServerError with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanValidateInternalServerError() *PostWorkforcemanagementManagementunitWorkplanValidateInternalServerError {
	return &PostWorkforcemanagementManagementunitWorkplanValidateInternalServerError{}
}

/*PostWorkforcemanagementManagementunitWorkplanValidateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementManagementunitWorkplanValidateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate][%d] postWorkforcemanagementManagementunitWorkplanValidateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanValidateServiceUnavailable creates a PostWorkforcemanagementManagementunitWorkplanValidateServiceUnavailable with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanValidateServiceUnavailable() *PostWorkforcemanagementManagementunitWorkplanValidateServiceUnavailable {
	return &PostWorkforcemanagementManagementunitWorkplanValidateServiceUnavailable{}
}

/*PostWorkforcemanagementManagementunitWorkplanValidateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementManagementunitWorkplanValidateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate][%d] postWorkforcemanagementManagementunitWorkplanValidateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanValidateGatewayTimeout creates a PostWorkforcemanagementManagementunitWorkplanValidateGatewayTimeout with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanValidateGatewayTimeout() *PostWorkforcemanagementManagementunitWorkplanValidateGatewayTimeout {
	return &PostWorkforcemanagementManagementunitWorkplanValidateGatewayTimeout{}
}

/*PostWorkforcemanagementManagementunitWorkplanValidateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementManagementunitWorkplanValidateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate][%d] postWorkforcemanagementManagementunitWorkplanValidateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanValidateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
