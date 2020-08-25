// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostGroupsReader is a Reader for the PostGroups structure.
type PostGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostGroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostGroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostGroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGroupsOK creates a PostGroupsOK with default headers values
func NewPostGroupsOK() *PostGroupsOK {
	return &PostGroupsOK{}
}

/*PostGroupsOK handles this case with default header values.

successful operation
*/
type PostGroupsOK struct {
	Payload *models.Group
}

func (o *PostGroupsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups][%d] postGroupsOK  %+v", 200, o.Payload)
}

func (o *PostGroupsOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *PostGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsBadRequest creates a PostGroupsBadRequest with default headers values
func NewPostGroupsBadRequest() *PostGroupsBadRequest {
	return &PostGroupsBadRequest{}
}

/*PostGroupsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGroupsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups][%d] postGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PostGroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsUnauthorized creates a PostGroupsUnauthorized with default headers values
func NewPostGroupsUnauthorized() *PostGroupsUnauthorized {
	return &PostGroupsUnauthorized{}
}

/*PostGroupsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGroupsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups][%d] postGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsForbidden creates a PostGroupsForbidden with default headers values
func NewPostGroupsForbidden() *PostGroupsForbidden {
	return &PostGroupsForbidden{}
}

/*PostGroupsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostGroupsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostGroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups][%d] postGroupsForbidden  %+v", 403, o.Payload)
}

func (o *PostGroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsNotFound creates a PostGroupsNotFound with default headers values
func NewPostGroupsNotFound() *PostGroupsNotFound {
	return &PostGroupsNotFound{}
}

/*PostGroupsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostGroupsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostGroupsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups][%d] postGroupsNotFound  %+v", 404, o.Payload)
}

func (o *PostGroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsRequestEntityTooLarge creates a PostGroupsRequestEntityTooLarge with default headers values
func NewPostGroupsRequestEntityTooLarge() *PostGroupsRequestEntityTooLarge {
	return &PostGroupsRequestEntityTooLarge{}
}

/*PostGroupsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostGroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostGroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups][%d] postGroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsUnsupportedMediaType creates a PostGroupsUnsupportedMediaType with default headers values
func NewPostGroupsUnsupportedMediaType() *PostGroupsUnsupportedMediaType {
	return &PostGroupsUnsupportedMediaType{}
}

/*PostGroupsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostGroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups][%d] postGroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsTooManyRequests creates a PostGroupsTooManyRequests with default headers values
func NewPostGroupsTooManyRequests() *PostGroupsTooManyRequests {
	return &PostGroupsTooManyRequests{}
}

/*PostGroupsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostGroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups][%d] postGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsInternalServerError creates a PostGroupsInternalServerError with default headers values
func NewPostGroupsInternalServerError() *PostGroupsInternalServerError {
	return &PostGroupsInternalServerError{}
}

/*PostGroupsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGroupsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups][%d] postGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsServiceUnavailable creates a PostGroupsServiceUnavailable with default headers values
func NewPostGroupsServiceUnavailable() *PostGroupsServiceUnavailable {
	return &PostGroupsServiceUnavailable{}
}

/*PostGroupsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostGroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups][%d] postGroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsGatewayTimeout creates a PostGroupsGatewayTimeout with default headers values
func NewPostGroupsGatewayTimeout() *PostGroupsGatewayTimeout {
	return &PostGroupsGatewayTimeout{}
}

/*PostGroupsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostGroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostGroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups][%d] postGroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}