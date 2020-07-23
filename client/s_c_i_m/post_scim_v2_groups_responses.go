// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostScimV2GroupsReader is a Reader for the PostScimV2Groups structure.
type PostScimV2GroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostScimV2GroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostScimV2GroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostScimV2GroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostScimV2GroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostScimV2GroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostScimV2GroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostScimV2GroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostScimV2GroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostScimV2GroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostScimV2GroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostScimV2GroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostScimV2GroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostScimV2GroupsOK creates a PostScimV2GroupsOK with default headers values
func NewPostScimV2GroupsOK() *PostScimV2GroupsOK {
	return &PostScimV2GroupsOK{}
}

/*PostScimV2GroupsOK handles this case with default header values.

successful operation
*/
type PostScimV2GroupsOK struct {
	Payload *models.ScimV2Group
}

func (o *PostScimV2GroupsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/groups][%d] postScimV2GroupsOK  %+v", 200, o.Payload)
}

func (o *PostScimV2GroupsOK) GetPayload() *models.ScimV2Group {
	return o.Payload
}

func (o *PostScimV2GroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimV2Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2GroupsBadRequest creates a PostScimV2GroupsBadRequest with default headers values
func NewPostScimV2GroupsBadRequest() *PostScimV2GroupsBadRequest {
	return &PostScimV2GroupsBadRequest{}
}

/*PostScimV2GroupsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostScimV2GroupsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2GroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/groups][%d] postScimV2GroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PostScimV2GroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2GroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2GroupsUnauthorized creates a PostScimV2GroupsUnauthorized with default headers values
func NewPostScimV2GroupsUnauthorized() *PostScimV2GroupsUnauthorized {
	return &PostScimV2GroupsUnauthorized{}
}

/*PostScimV2GroupsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostScimV2GroupsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2GroupsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/groups][%d] postScimV2GroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostScimV2GroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2GroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2GroupsForbidden creates a PostScimV2GroupsForbidden with default headers values
func NewPostScimV2GroupsForbidden() *PostScimV2GroupsForbidden {
	return &PostScimV2GroupsForbidden{}
}

/*PostScimV2GroupsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostScimV2GroupsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2GroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/groups][%d] postScimV2GroupsForbidden  %+v", 403, o.Payload)
}

func (o *PostScimV2GroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2GroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2GroupsNotFound creates a PostScimV2GroupsNotFound with default headers values
func NewPostScimV2GroupsNotFound() *PostScimV2GroupsNotFound {
	return &PostScimV2GroupsNotFound{}
}

/*PostScimV2GroupsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostScimV2GroupsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2GroupsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/groups][%d] postScimV2GroupsNotFound  %+v", 404, o.Payload)
}

func (o *PostScimV2GroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2GroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2GroupsRequestEntityTooLarge creates a PostScimV2GroupsRequestEntityTooLarge with default headers values
func NewPostScimV2GroupsRequestEntityTooLarge() *PostScimV2GroupsRequestEntityTooLarge {
	return &PostScimV2GroupsRequestEntityTooLarge{}
}

/*PostScimV2GroupsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostScimV2GroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2GroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/groups][%d] postScimV2GroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostScimV2GroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2GroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2GroupsUnsupportedMediaType creates a PostScimV2GroupsUnsupportedMediaType with default headers values
func NewPostScimV2GroupsUnsupportedMediaType() *PostScimV2GroupsUnsupportedMediaType {
	return &PostScimV2GroupsUnsupportedMediaType{}
}

/*PostScimV2GroupsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostScimV2GroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2GroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/groups][%d] postScimV2GroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostScimV2GroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2GroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2GroupsTooManyRequests creates a PostScimV2GroupsTooManyRequests with default headers values
func NewPostScimV2GroupsTooManyRequests() *PostScimV2GroupsTooManyRequests {
	return &PostScimV2GroupsTooManyRequests{}
}

/*PostScimV2GroupsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostScimV2GroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2GroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/groups][%d] postScimV2GroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostScimV2GroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2GroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2GroupsInternalServerError creates a PostScimV2GroupsInternalServerError with default headers values
func NewPostScimV2GroupsInternalServerError() *PostScimV2GroupsInternalServerError {
	return &PostScimV2GroupsInternalServerError{}
}

/*PostScimV2GroupsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostScimV2GroupsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2GroupsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/groups][%d] postScimV2GroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostScimV2GroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2GroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2GroupsServiceUnavailable creates a PostScimV2GroupsServiceUnavailable with default headers values
func NewPostScimV2GroupsServiceUnavailable() *PostScimV2GroupsServiceUnavailable {
	return &PostScimV2GroupsServiceUnavailable{}
}

/*PostScimV2GroupsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostScimV2GroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2GroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/groups][%d] postScimV2GroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostScimV2GroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2GroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2GroupsGatewayTimeout creates a PostScimV2GroupsGatewayTimeout with default headers values
func NewPostScimV2GroupsGatewayTimeout() *PostScimV2GroupsGatewayTimeout {
	return &PostScimV2GroupsGatewayTimeout{}
}

/*PostScimV2GroupsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostScimV2GroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2GroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/groups][%d] postScimV2GroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostScimV2GroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2GroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
