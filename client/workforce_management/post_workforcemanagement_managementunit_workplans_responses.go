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

// PostWorkforcemanagementManagementunitWorkplansReader is a Reader for the PostWorkforcemanagementManagementunitWorkplans structure.
type PostWorkforcemanagementManagementunitWorkplansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementManagementunitWorkplansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementManagementunitWorkplansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementManagementunitWorkplansBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementManagementunitWorkplansUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementManagementunitWorkplansForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementManagementunitWorkplansNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementManagementunitWorkplansRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementManagementunitWorkplansUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementManagementunitWorkplansTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementManagementunitWorkplansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementManagementunitWorkplansServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementManagementunitWorkplansGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementManagementunitWorkplansOK creates a PostWorkforcemanagementManagementunitWorkplansOK with default headers values
func NewPostWorkforcemanagementManagementunitWorkplansOK() *PostWorkforcemanagementManagementunitWorkplansOK {
	return &PostWorkforcemanagementManagementunitWorkplansOK{}
}

/*PostWorkforcemanagementManagementunitWorkplansOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementManagementunitWorkplansOK struct {
	Payload *models.WorkPlan
}

func (o *PostWorkforcemanagementManagementunitWorkplansOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans][%d] postWorkforcemanagementManagementunitWorkplansOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplansOK) GetPayload() *models.WorkPlan {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplansBadRequest creates a PostWorkforcemanagementManagementunitWorkplansBadRequest with default headers values
func NewPostWorkforcemanagementManagementunitWorkplansBadRequest() *PostWorkforcemanagementManagementunitWorkplansBadRequest {
	return &PostWorkforcemanagementManagementunitWorkplansBadRequest{}
}

/*PostWorkforcemanagementManagementunitWorkplansBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementManagementunitWorkplansBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplansBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans][%d] postWorkforcemanagementManagementunitWorkplansBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplansBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplansBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplansUnauthorized creates a PostWorkforcemanagementManagementunitWorkplansUnauthorized with default headers values
func NewPostWorkforcemanagementManagementunitWorkplansUnauthorized() *PostWorkforcemanagementManagementunitWorkplansUnauthorized {
	return &PostWorkforcemanagementManagementunitWorkplansUnauthorized{}
}

/*PostWorkforcemanagementManagementunitWorkplansUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementManagementunitWorkplansUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplansUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans][%d] postWorkforcemanagementManagementunitWorkplansUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplansUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplansUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplansForbidden creates a PostWorkforcemanagementManagementunitWorkplansForbidden with default headers values
func NewPostWorkforcemanagementManagementunitWorkplansForbidden() *PostWorkforcemanagementManagementunitWorkplansForbidden {
	return &PostWorkforcemanagementManagementunitWorkplansForbidden{}
}

/*PostWorkforcemanagementManagementunitWorkplansForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementManagementunitWorkplansForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplansForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans][%d] postWorkforcemanagementManagementunitWorkplansForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplansForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplansForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplansNotFound creates a PostWorkforcemanagementManagementunitWorkplansNotFound with default headers values
func NewPostWorkforcemanagementManagementunitWorkplansNotFound() *PostWorkforcemanagementManagementunitWorkplansNotFound {
	return &PostWorkforcemanagementManagementunitWorkplansNotFound{}
}

/*PostWorkforcemanagementManagementunitWorkplansNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementManagementunitWorkplansNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplansNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans][%d] postWorkforcemanagementManagementunitWorkplansNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplansNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplansNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplansRequestEntityTooLarge creates a PostWorkforcemanagementManagementunitWorkplansRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementManagementunitWorkplansRequestEntityTooLarge() *PostWorkforcemanagementManagementunitWorkplansRequestEntityTooLarge {
	return &PostWorkforcemanagementManagementunitWorkplansRequestEntityTooLarge{}
}

/*PostWorkforcemanagementManagementunitWorkplansRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostWorkforcemanagementManagementunitWorkplansRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplansRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans][%d] postWorkforcemanagementManagementunitWorkplansRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplansRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplansRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplansUnsupportedMediaType creates a PostWorkforcemanagementManagementunitWorkplansUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementManagementunitWorkplansUnsupportedMediaType() *PostWorkforcemanagementManagementunitWorkplansUnsupportedMediaType {
	return &PostWorkforcemanagementManagementunitWorkplansUnsupportedMediaType{}
}

/*PostWorkforcemanagementManagementunitWorkplansUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementManagementunitWorkplansUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplansUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans][%d] postWorkforcemanagementManagementunitWorkplansUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplansUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplansUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplansTooManyRequests creates a PostWorkforcemanagementManagementunitWorkplansTooManyRequests with default headers values
func NewPostWorkforcemanagementManagementunitWorkplansTooManyRequests() *PostWorkforcemanagementManagementunitWorkplansTooManyRequests {
	return &PostWorkforcemanagementManagementunitWorkplansTooManyRequests{}
}

/*PostWorkforcemanagementManagementunitWorkplansTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostWorkforcemanagementManagementunitWorkplansTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplansTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans][%d] postWorkforcemanagementManagementunitWorkplansTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplansTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplansTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplansInternalServerError creates a PostWorkforcemanagementManagementunitWorkplansInternalServerError with default headers values
func NewPostWorkforcemanagementManagementunitWorkplansInternalServerError() *PostWorkforcemanagementManagementunitWorkplansInternalServerError {
	return &PostWorkforcemanagementManagementunitWorkplansInternalServerError{}
}

/*PostWorkforcemanagementManagementunitWorkplansInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementManagementunitWorkplansInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplansInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans][%d] postWorkforcemanagementManagementunitWorkplansInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplansInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplansInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplansServiceUnavailable creates a PostWorkforcemanagementManagementunitWorkplansServiceUnavailable with default headers values
func NewPostWorkforcemanagementManagementunitWorkplansServiceUnavailable() *PostWorkforcemanagementManagementunitWorkplansServiceUnavailable {
	return &PostWorkforcemanagementManagementunitWorkplansServiceUnavailable{}
}

/*PostWorkforcemanagementManagementunitWorkplansServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementManagementunitWorkplansServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplansServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans][%d] postWorkforcemanagementManagementunitWorkplansServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplansServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplansServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplansGatewayTimeout creates a PostWorkforcemanagementManagementunitWorkplansGatewayTimeout with default headers values
func NewPostWorkforcemanagementManagementunitWorkplansGatewayTimeout() *PostWorkforcemanagementManagementunitWorkplansGatewayTimeout {
	return &PostWorkforcemanagementManagementunitWorkplansGatewayTimeout{}
}

/*PostWorkforcemanagementManagementunitWorkplansGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementManagementunitWorkplansGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplansGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans][%d] postWorkforcemanagementManagementunitWorkplansGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplansGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplansGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
