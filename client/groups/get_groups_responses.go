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

// GetGroupsReader is a Reader for the GetGroups structure.
type GetGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupsOK creates a GetGroupsOK with default headers values
func NewGetGroupsOK() *GetGroupsOK {
	return &GetGroupsOK{}
}

/*GetGroupsOK handles this case with default header values.

successful operation
*/
type GetGroupsOK struct {
	Payload *models.GroupEntityListing
}

func (o *GetGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsOK  %+v", 200, o.Payload)
}

func (o *GetGroupsOK) GetPayload() *models.GroupEntityListing {
	return o.Payload
}

func (o *GetGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsBadRequest creates a GetGroupsBadRequest with default headers values
func NewGetGroupsBadRequest() *GetGroupsBadRequest {
	return &GetGroupsBadRequest{}
}

/*GetGroupsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGroupsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsUnauthorized creates a GetGroupsUnauthorized with default headers values
func NewGetGroupsUnauthorized() *GetGroupsUnauthorized {
	return &GetGroupsUnauthorized{}
}

/*GetGroupsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGroupsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsForbidden creates a GetGroupsForbidden with default headers values
func NewGetGroupsForbidden() *GetGroupsForbidden {
	return &GetGroupsForbidden{}
}

/*GetGroupsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGroupsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetGroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsNotFound creates a GetGroupsNotFound with default headers values
func NewGetGroupsNotFound() *GetGroupsNotFound {
	return &GetGroupsNotFound{}
}

/*GetGroupsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGroupsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsNotFound  %+v", 404, o.Payload)
}

func (o *GetGroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsRequestEntityTooLarge creates a GetGroupsRequestEntityTooLarge with default headers values
func NewGetGroupsRequestEntityTooLarge() *GetGroupsRequestEntityTooLarge {
	return &GetGroupsRequestEntityTooLarge{}
}

/*GetGroupsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsUnsupportedMediaType creates a GetGroupsUnsupportedMediaType with default headers values
func NewGetGroupsUnsupportedMediaType() *GetGroupsUnsupportedMediaType {
	return &GetGroupsUnsupportedMediaType{}
}

/*GetGroupsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsTooManyRequests creates a GetGroupsTooManyRequests with default headers values
func NewGetGroupsTooManyRequests() *GetGroupsTooManyRequests {
	return &GetGroupsTooManyRequests{}
}

/*GetGroupsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetGroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsInternalServerError creates a GetGroupsInternalServerError with default headers values
func NewGetGroupsInternalServerError() *GetGroupsInternalServerError {
	return &GetGroupsInternalServerError{}
}

/*GetGroupsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGroupsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsServiceUnavailable creates a GetGroupsServiceUnavailable with default headers values
func NewGetGroupsServiceUnavailable() *GetGroupsServiceUnavailable {
	return &GetGroupsServiceUnavailable{}
}

/*GetGroupsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsGatewayTimeout creates a GetGroupsGatewayTimeout with default headers values
func NewGetGroupsGatewayTimeout() *GetGroupsGatewayTimeout {
	return &GetGroupsGatewayTimeout{}
}

/*GetGroupsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
