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

// PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateReader is a Reader for the PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerate structure.
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateBadRequest creates a PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateBadRequest with default headers values
func NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateBadRequest() *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateBadRequest {
	return &PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateBadRequest{}
}

/*PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementManagementunitWeekShorttermforecastsGenerateBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnauthorized creates a PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnauthorized with default headers values
func NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnauthorized() *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnauthorized {
	return &PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnauthorized{}
}

/*PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateForbidden creates a PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateForbidden with default headers values
func NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateForbidden() *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateForbidden {
	return &PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateForbidden{}
}

/*PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementManagementunitWeekShorttermforecastsGenerateForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateNotFound creates a PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateNotFound with default headers values
func NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateNotFound() *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateNotFound {
	return &PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateNotFound{}
}

/*PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementManagementunitWeekShorttermforecastsGenerateNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGone creates a PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGone with default headers values
func NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGone() *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGone {
	return &PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGone{}
}

/*PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGone handles this case with default header values.

Gone
*/
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGone struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGone) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGone  %+v", 410, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateRequestEntityTooLarge creates a PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateRequestEntityTooLarge() *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateRequestEntityTooLarge {
	return &PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateRequestEntityTooLarge{}
}

/*PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementManagementunitWeekShorttermforecastsGenerateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnsupportedMediaType creates a PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnsupportedMediaType() *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnsupportedMediaType {
	return &PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnsupportedMediaType{}
}

/*PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateTooManyRequests creates a PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateTooManyRequests with default headers values
func NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateTooManyRequests() *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateTooManyRequests {
	return &PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateTooManyRequests{}
}

/*PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementManagementunitWeekShorttermforecastsGenerateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateInternalServerError creates a PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateInternalServerError with default headers values
func NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateInternalServerError() *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateInternalServerError {
	return &PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateInternalServerError{}
}

/*PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementManagementunitWeekShorttermforecastsGenerateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateServiceUnavailable creates a PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateServiceUnavailable with default headers values
func NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateServiceUnavailable() *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateServiceUnavailable {
	return &PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateServiceUnavailable{}
}

/*PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementManagementunitWeekShorttermforecastsGenerateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGatewayTimeout creates a PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGatewayTimeout with default headers values
func NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGatewayTimeout() *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGatewayTimeout {
	return &PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGatewayTimeout{}
}

/*PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateDefault creates a PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateDefault with default headers values
func NewPostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateDefault(code int) *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateDefault {
	return &PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateDefault{
		_statusCode: code,
	}
}

/*PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateDefault handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateDefault struct {
	_statusCode int
}

// Code gets the status code for the post workforcemanagement managementunit week shorttermforecasts generate default response
func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateDefault) Code() int {
	return o._statusCode
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementManagementunitWeekShorttermforecastsGenerate default ", o._statusCode)
}

func (o *PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}