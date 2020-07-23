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

// PostRoutingSkillsReader is a Reader for the PostRoutingSkills structure.
type PostRoutingSkillsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRoutingSkillsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRoutingSkillsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRoutingSkillsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRoutingSkillsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRoutingSkillsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRoutingSkillsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostRoutingSkillsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRoutingSkillsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRoutingSkillsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRoutingSkillsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRoutingSkillsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRoutingSkillsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRoutingSkillsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRoutingSkillsOK creates a PostRoutingSkillsOK with default headers values
func NewPostRoutingSkillsOK() *PostRoutingSkillsOK {
	return &PostRoutingSkillsOK{}
}

/*PostRoutingSkillsOK handles this case with default header values.

successful operation
*/
type PostRoutingSkillsOK struct {
	Payload *models.RoutingSkill
}

func (o *PostRoutingSkillsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/skills][%d] postRoutingSkillsOK  %+v", 200, o.Payload)
}

func (o *PostRoutingSkillsOK) GetPayload() *models.RoutingSkill {
	return o.Payload
}

func (o *PostRoutingSkillsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoutingSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingSkillsBadRequest creates a PostRoutingSkillsBadRequest with default headers values
func NewPostRoutingSkillsBadRequest() *PostRoutingSkillsBadRequest {
	return &PostRoutingSkillsBadRequest{}
}

/*PostRoutingSkillsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRoutingSkillsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingSkillsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/skills][%d] postRoutingSkillsBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingSkillsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingSkillsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingSkillsUnauthorized creates a PostRoutingSkillsUnauthorized with default headers values
func NewPostRoutingSkillsUnauthorized() *PostRoutingSkillsUnauthorized {
	return &PostRoutingSkillsUnauthorized{}
}

/*PostRoutingSkillsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRoutingSkillsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingSkillsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/skills][%d] postRoutingSkillsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingSkillsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingSkillsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingSkillsForbidden creates a PostRoutingSkillsForbidden with default headers values
func NewPostRoutingSkillsForbidden() *PostRoutingSkillsForbidden {
	return &PostRoutingSkillsForbidden{}
}

/*PostRoutingSkillsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostRoutingSkillsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingSkillsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/skills][%d] postRoutingSkillsForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingSkillsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingSkillsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingSkillsNotFound creates a PostRoutingSkillsNotFound with default headers values
func NewPostRoutingSkillsNotFound() *PostRoutingSkillsNotFound {
	return &PostRoutingSkillsNotFound{}
}

/*PostRoutingSkillsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostRoutingSkillsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingSkillsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/skills][%d] postRoutingSkillsNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingSkillsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingSkillsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingSkillsConflict creates a PostRoutingSkillsConflict with default headers values
func NewPostRoutingSkillsConflict() *PostRoutingSkillsConflict {
	return &PostRoutingSkillsConflict{}
}

/*PostRoutingSkillsConflict handles this case with default header values.

Conflict
*/
type PostRoutingSkillsConflict struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingSkillsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/skills][%d] postRoutingSkillsConflict  %+v", 409, o.Payload)
}

func (o *PostRoutingSkillsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingSkillsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingSkillsRequestEntityTooLarge creates a PostRoutingSkillsRequestEntityTooLarge with default headers values
func NewPostRoutingSkillsRequestEntityTooLarge() *PostRoutingSkillsRequestEntityTooLarge {
	return &PostRoutingSkillsRequestEntityTooLarge{}
}

/*PostRoutingSkillsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostRoutingSkillsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingSkillsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/skills][%d] postRoutingSkillsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingSkillsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingSkillsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingSkillsUnsupportedMediaType creates a PostRoutingSkillsUnsupportedMediaType with default headers values
func NewPostRoutingSkillsUnsupportedMediaType() *PostRoutingSkillsUnsupportedMediaType {
	return &PostRoutingSkillsUnsupportedMediaType{}
}

/*PostRoutingSkillsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRoutingSkillsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingSkillsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/skills][%d] postRoutingSkillsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingSkillsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingSkillsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingSkillsTooManyRequests creates a PostRoutingSkillsTooManyRequests with default headers values
func NewPostRoutingSkillsTooManyRequests() *PostRoutingSkillsTooManyRequests {
	return &PostRoutingSkillsTooManyRequests{}
}

/*PostRoutingSkillsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostRoutingSkillsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingSkillsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/skills][%d] postRoutingSkillsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingSkillsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingSkillsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingSkillsInternalServerError creates a PostRoutingSkillsInternalServerError with default headers values
func NewPostRoutingSkillsInternalServerError() *PostRoutingSkillsInternalServerError {
	return &PostRoutingSkillsInternalServerError{}
}

/*PostRoutingSkillsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRoutingSkillsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingSkillsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/skills][%d] postRoutingSkillsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingSkillsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingSkillsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingSkillsServiceUnavailable creates a PostRoutingSkillsServiceUnavailable with default headers values
func NewPostRoutingSkillsServiceUnavailable() *PostRoutingSkillsServiceUnavailable {
	return &PostRoutingSkillsServiceUnavailable{}
}

/*PostRoutingSkillsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRoutingSkillsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingSkillsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/skills][%d] postRoutingSkillsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingSkillsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingSkillsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingSkillsGatewayTimeout creates a PostRoutingSkillsGatewayTimeout with default headers values
func NewPostRoutingSkillsGatewayTimeout() *PostRoutingSkillsGatewayTimeout {
	return &PostRoutingSkillsGatewayTimeout{}
}

/*PostRoutingSkillsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostRoutingSkillsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingSkillsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/skills][%d] postRoutingSkillsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingSkillsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingSkillsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
