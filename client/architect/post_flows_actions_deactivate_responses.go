// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostFlowsActionsDeactivateReader is a Reader for the PostFlowsActionsDeactivate structure.
type PostFlowsActionsDeactivateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFlowsActionsDeactivateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostFlowsActionsDeactivateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostFlowsActionsDeactivateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostFlowsActionsDeactivateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostFlowsActionsDeactivateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostFlowsActionsDeactivateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostFlowsActionsDeactivateMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostFlowsActionsDeactivateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostFlowsActionsDeactivateGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostFlowsActionsDeactivateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostFlowsActionsDeactivateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostFlowsActionsDeactivateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostFlowsActionsDeactivateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostFlowsActionsDeactivateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostFlowsActionsDeactivateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFlowsActionsDeactivateOK creates a PostFlowsActionsDeactivateOK with default headers values
func NewPostFlowsActionsDeactivateOK() *PostFlowsActionsDeactivateOK {
	return &PostFlowsActionsDeactivateOK{}
}

/*PostFlowsActionsDeactivateOK handles this case with default header values.

successful operation
*/
type PostFlowsActionsDeactivateOK struct {
	Payload *models.Flow
}

func (o *PostFlowsActionsDeactivateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateOK  %+v", 200, o.Payload)
}

func (o *PostFlowsActionsDeactivateOK) GetPayload() *models.Flow {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Flow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateBadRequest creates a PostFlowsActionsDeactivateBadRequest with default headers values
func NewPostFlowsActionsDeactivateBadRequest() *PostFlowsActionsDeactivateBadRequest {
	return &PostFlowsActionsDeactivateBadRequest{}
}

/*PostFlowsActionsDeactivateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostFlowsActionsDeactivateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateBadRequest  %+v", 400, o.Payload)
}

func (o *PostFlowsActionsDeactivateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateUnauthorized creates a PostFlowsActionsDeactivateUnauthorized with default headers values
func NewPostFlowsActionsDeactivateUnauthorized() *PostFlowsActionsDeactivateUnauthorized {
	return &PostFlowsActionsDeactivateUnauthorized{}
}

/*PostFlowsActionsDeactivateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostFlowsActionsDeactivateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostFlowsActionsDeactivateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateForbidden creates a PostFlowsActionsDeactivateForbidden with default headers values
func NewPostFlowsActionsDeactivateForbidden() *PostFlowsActionsDeactivateForbidden {
	return &PostFlowsActionsDeactivateForbidden{}
}

/*PostFlowsActionsDeactivateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostFlowsActionsDeactivateForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateForbidden  %+v", 403, o.Payload)
}

func (o *PostFlowsActionsDeactivateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateNotFound creates a PostFlowsActionsDeactivateNotFound with default headers values
func NewPostFlowsActionsDeactivateNotFound() *PostFlowsActionsDeactivateNotFound {
	return &PostFlowsActionsDeactivateNotFound{}
}

/*PostFlowsActionsDeactivateNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostFlowsActionsDeactivateNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateNotFound  %+v", 404, o.Payload)
}

func (o *PostFlowsActionsDeactivateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateMethodNotAllowed creates a PostFlowsActionsDeactivateMethodNotAllowed with default headers values
func NewPostFlowsActionsDeactivateMethodNotAllowed() *PostFlowsActionsDeactivateMethodNotAllowed {
	return &PostFlowsActionsDeactivateMethodNotAllowed{}
}

/*PostFlowsActionsDeactivateMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type PostFlowsActionsDeactivateMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostFlowsActionsDeactivateMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateConflict creates a PostFlowsActionsDeactivateConflict with default headers values
func NewPostFlowsActionsDeactivateConflict() *PostFlowsActionsDeactivateConflict {
	return &PostFlowsActionsDeactivateConflict{}
}

/*PostFlowsActionsDeactivateConflict handles this case with default header values.

Conflict
*/
type PostFlowsActionsDeactivateConflict struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateConflict  %+v", 409, o.Payload)
}

func (o *PostFlowsActionsDeactivateConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateGone creates a PostFlowsActionsDeactivateGone with default headers values
func NewPostFlowsActionsDeactivateGone() *PostFlowsActionsDeactivateGone {
	return &PostFlowsActionsDeactivateGone{}
}

/*PostFlowsActionsDeactivateGone handles this case with default header values.

Gone
*/
type PostFlowsActionsDeactivateGone struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateGone) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateGone  %+v", 410, o.Payload)
}

func (o *PostFlowsActionsDeactivateGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateRequestEntityTooLarge creates a PostFlowsActionsDeactivateRequestEntityTooLarge with default headers values
func NewPostFlowsActionsDeactivateRequestEntityTooLarge() *PostFlowsActionsDeactivateRequestEntityTooLarge {
	return &PostFlowsActionsDeactivateRequestEntityTooLarge{}
}

/*PostFlowsActionsDeactivateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostFlowsActionsDeactivateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostFlowsActionsDeactivateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateUnsupportedMediaType creates a PostFlowsActionsDeactivateUnsupportedMediaType with default headers values
func NewPostFlowsActionsDeactivateUnsupportedMediaType() *PostFlowsActionsDeactivateUnsupportedMediaType {
	return &PostFlowsActionsDeactivateUnsupportedMediaType{}
}

/*PostFlowsActionsDeactivateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostFlowsActionsDeactivateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostFlowsActionsDeactivateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateTooManyRequests creates a PostFlowsActionsDeactivateTooManyRequests with default headers values
func NewPostFlowsActionsDeactivateTooManyRequests() *PostFlowsActionsDeactivateTooManyRequests {
	return &PostFlowsActionsDeactivateTooManyRequests{}
}

/*PostFlowsActionsDeactivateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostFlowsActionsDeactivateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostFlowsActionsDeactivateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateInternalServerError creates a PostFlowsActionsDeactivateInternalServerError with default headers values
func NewPostFlowsActionsDeactivateInternalServerError() *PostFlowsActionsDeactivateInternalServerError {
	return &PostFlowsActionsDeactivateInternalServerError{}
}

/*PostFlowsActionsDeactivateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostFlowsActionsDeactivateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFlowsActionsDeactivateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateServiceUnavailable creates a PostFlowsActionsDeactivateServiceUnavailable with default headers values
func NewPostFlowsActionsDeactivateServiceUnavailable() *PostFlowsActionsDeactivateServiceUnavailable {
	return &PostFlowsActionsDeactivateServiceUnavailable{}
}

/*PostFlowsActionsDeactivateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostFlowsActionsDeactivateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostFlowsActionsDeactivateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsDeactivateGatewayTimeout creates a PostFlowsActionsDeactivateGatewayTimeout with default headers values
func NewPostFlowsActionsDeactivateGatewayTimeout() *PostFlowsActionsDeactivateGatewayTimeout {
	return &PostFlowsActionsDeactivateGatewayTimeout{}
}

/*PostFlowsActionsDeactivateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostFlowsActionsDeactivateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsActionsDeactivateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/deactivate][%d] postFlowsActionsDeactivateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostFlowsActionsDeactivateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsDeactivateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}