// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostUserRoutingskillsReader is a Reader for the PostUserRoutingskills structure.
type PostUserRoutingskillsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserRoutingskillsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserRoutingskillsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUserRoutingskillsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUserRoutingskillsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUserRoutingskillsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUserRoutingskillsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostUserRoutingskillsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostUserRoutingskillsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostUserRoutingskillsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserRoutingskillsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUserRoutingskillsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUserRoutingskillsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUserRoutingskillsOK creates a PostUserRoutingskillsOK with default headers values
func NewPostUserRoutingskillsOK() *PostUserRoutingskillsOK {
	return &PostUserRoutingskillsOK{}
}

/*PostUserRoutingskillsOK handles this case with default header values.

successful operation
*/
type PostUserRoutingskillsOK struct {
	Payload *models.UserRoutingSkill
}

func (o *PostUserRoutingskillsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routingskills][%d] postUserRoutingskillsOK  %+v", 200, o.Payload)
}

func (o *PostUserRoutingskillsOK) GetPayload() *models.UserRoutingSkill {
	return o.Payload
}

func (o *PostUserRoutingskillsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserRoutingSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutingskillsBadRequest creates a PostUserRoutingskillsBadRequest with default headers values
func NewPostUserRoutingskillsBadRequest() *PostUserRoutingskillsBadRequest {
	return &PostUserRoutingskillsBadRequest{}
}

/*PostUserRoutingskillsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostUserRoutingskillsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutingskillsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routingskills][%d] postUserRoutingskillsBadRequest  %+v", 400, o.Payload)
}

func (o *PostUserRoutingskillsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutingskillsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutingskillsUnauthorized creates a PostUserRoutingskillsUnauthorized with default headers values
func NewPostUserRoutingskillsUnauthorized() *PostUserRoutingskillsUnauthorized {
	return &PostUserRoutingskillsUnauthorized{}
}

/*PostUserRoutingskillsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostUserRoutingskillsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutingskillsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routingskills][%d] postUserRoutingskillsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUserRoutingskillsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutingskillsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutingskillsForbidden creates a PostUserRoutingskillsForbidden with default headers values
func NewPostUserRoutingskillsForbidden() *PostUserRoutingskillsForbidden {
	return &PostUserRoutingskillsForbidden{}
}

/*PostUserRoutingskillsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostUserRoutingskillsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutingskillsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routingskills][%d] postUserRoutingskillsForbidden  %+v", 403, o.Payload)
}

func (o *PostUserRoutingskillsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutingskillsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutingskillsNotFound creates a PostUserRoutingskillsNotFound with default headers values
func NewPostUserRoutingskillsNotFound() *PostUserRoutingskillsNotFound {
	return &PostUserRoutingskillsNotFound{}
}

/*PostUserRoutingskillsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostUserRoutingskillsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutingskillsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routingskills][%d] postUserRoutingskillsNotFound  %+v", 404, o.Payload)
}

func (o *PostUserRoutingskillsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutingskillsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutingskillsRequestEntityTooLarge creates a PostUserRoutingskillsRequestEntityTooLarge with default headers values
func NewPostUserRoutingskillsRequestEntityTooLarge() *PostUserRoutingskillsRequestEntityTooLarge {
	return &PostUserRoutingskillsRequestEntityTooLarge{}
}

/*PostUserRoutingskillsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostUserRoutingskillsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutingskillsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routingskills][%d] postUserRoutingskillsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUserRoutingskillsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutingskillsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutingskillsUnsupportedMediaType creates a PostUserRoutingskillsUnsupportedMediaType with default headers values
func NewPostUserRoutingskillsUnsupportedMediaType() *PostUserRoutingskillsUnsupportedMediaType {
	return &PostUserRoutingskillsUnsupportedMediaType{}
}

/*PostUserRoutingskillsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostUserRoutingskillsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutingskillsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routingskills][%d] postUserRoutingskillsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUserRoutingskillsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutingskillsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutingskillsTooManyRequests creates a PostUserRoutingskillsTooManyRequests with default headers values
func NewPostUserRoutingskillsTooManyRequests() *PostUserRoutingskillsTooManyRequests {
	return &PostUserRoutingskillsTooManyRequests{}
}

/*PostUserRoutingskillsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostUserRoutingskillsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutingskillsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routingskills][%d] postUserRoutingskillsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUserRoutingskillsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutingskillsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutingskillsInternalServerError creates a PostUserRoutingskillsInternalServerError with default headers values
func NewPostUserRoutingskillsInternalServerError() *PostUserRoutingskillsInternalServerError {
	return &PostUserRoutingskillsInternalServerError{}
}

/*PostUserRoutingskillsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostUserRoutingskillsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutingskillsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routingskills][%d] postUserRoutingskillsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUserRoutingskillsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutingskillsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutingskillsServiceUnavailable creates a PostUserRoutingskillsServiceUnavailable with default headers values
func NewPostUserRoutingskillsServiceUnavailable() *PostUserRoutingskillsServiceUnavailable {
	return &PostUserRoutingskillsServiceUnavailable{}
}

/*PostUserRoutingskillsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostUserRoutingskillsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutingskillsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routingskills][%d] postUserRoutingskillsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUserRoutingskillsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutingskillsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutingskillsGatewayTimeout creates a PostUserRoutingskillsGatewayTimeout with default headers values
func NewPostUserRoutingskillsGatewayTimeout() *PostUserRoutingskillsGatewayTimeout {
	return &PostUserRoutingskillsGatewayTimeout{}
}

/*PostUserRoutingskillsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostUserRoutingskillsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutingskillsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routingskills][%d] postUserRoutingskillsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUserRoutingskillsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutingskillsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
