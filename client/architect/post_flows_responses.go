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

// PostFlowsReader is a Reader for the PostFlows structure.
type PostFlowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFlowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostFlowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostFlowsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostFlowsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostFlowsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostFlowsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostFlowsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostFlowsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostFlowsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostFlowsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostFlowsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostFlowsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostFlowsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostFlowsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFlowsOK creates a PostFlowsOK with default headers values
func NewPostFlowsOK() *PostFlowsOK {
	return &PostFlowsOK{}
}

/*PostFlowsOK handles this case with default header values.

successful operation
*/
type PostFlowsOK struct {
	Payload *models.Flow
}

func (o *PostFlowsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsOK  %+v", 200, o.Payload)
}

func (o *PostFlowsOK) GetPayload() *models.Flow {
	return o.Payload
}

func (o *PostFlowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Flow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsBadRequest creates a PostFlowsBadRequest with default headers values
func NewPostFlowsBadRequest() *PostFlowsBadRequest {
	return &PostFlowsBadRequest{}
}

/*PostFlowsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostFlowsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsBadRequest  %+v", 400, o.Payload)
}

func (o *PostFlowsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsUnauthorized creates a PostFlowsUnauthorized with default headers values
func NewPostFlowsUnauthorized() *PostFlowsUnauthorized {
	return &PostFlowsUnauthorized{}
}

/*PostFlowsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostFlowsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostFlowsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsForbidden creates a PostFlowsForbidden with default headers values
func NewPostFlowsForbidden() *PostFlowsForbidden {
	return &PostFlowsForbidden{}
}

/*PostFlowsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostFlowsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsForbidden  %+v", 403, o.Payload)
}

func (o *PostFlowsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsNotFound creates a PostFlowsNotFound with default headers values
func NewPostFlowsNotFound() *PostFlowsNotFound {
	return &PostFlowsNotFound{}
}

/*PostFlowsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostFlowsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsNotFound  %+v", 404, o.Payload)
}

func (o *PostFlowsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsMethodNotAllowed creates a PostFlowsMethodNotAllowed with default headers values
func NewPostFlowsMethodNotAllowed() *PostFlowsMethodNotAllowed {
	return &PostFlowsMethodNotAllowed{}
}

/*PostFlowsMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type PostFlowsMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostFlowsMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsConflict creates a PostFlowsConflict with default headers values
func NewPostFlowsConflict() *PostFlowsConflict {
	return &PostFlowsConflict{}
}

/*PostFlowsConflict handles this case with default header values.

Conflict
*/
type PostFlowsConflict struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsConflict  %+v", 409, o.Payload)
}

func (o *PostFlowsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsRequestEntityTooLarge creates a PostFlowsRequestEntityTooLarge with default headers values
func NewPostFlowsRequestEntityTooLarge() *PostFlowsRequestEntityTooLarge {
	return &PostFlowsRequestEntityTooLarge{}
}

/*PostFlowsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostFlowsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostFlowsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsUnsupportedMediaType creates a PostFlowsUnsupportedMediaType with default headers values
func NewPostFlowsUnsupportedMediaType() *PostFlowsUnsupportedMediaType {
	return &PostFlowsUnsupportedMediaType{}
}

/*PostFlowsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostFlowsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostFlowsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsTooManyRequests creates a PostFlowsTooManyRequests with default headers values
func NewPostFlowsTooManyRequests() *PostFlowsTooManyRequests {
	return &PostFlowsTooManyRequests{}
}

/*PostFlowsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostFlowsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostFlowsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsInternalServerError creates a PostFlowsInternalServerError with default headers values
func NewPostFlowsInternalServerError() *PostFlowsInternalServerError {
	return &PostFlowsInternalServerError{}
}

/*PostFlowsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostFlowsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFlowsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsServiceUnavailable creates a PostFlowsServiceUnavailable with default headers values
func NewPostFlowsServiceUnavailable() *PostFlowsServiceUnavailable {
	return &PostFlowsServiceUnavailable{}
}

/*PostFlowsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostFlowsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostFlowsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsGatewayTimeout creates a PostFlowsGatewayTimeout with default headers values
func NewPostFlowsGatewayTimeout() *PostFlowsGatewayTimeout {
	return &PostFlowsGatewayTimeout{}
}

/*PostFlowsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostFlowsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows][%d] postFlowsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostFlowsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
