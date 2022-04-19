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

// GetGroupReader is a Reader for the GetGroup structure.
type GetGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGroupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupOK creates a GetGroupOK with default headers values
func NewGetGroupOK() *GetGroupOK {
	return &GetGroupOK{}
}

/*GetGroupOK handles this case with default header values.

successful operation
*/
type GetGroupOK struct {
	Payload *models.Group
}

func (o *GetGroupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}][%d] getGroupOK  %+v", 200, o.Payload)
}

func (o *GetGroupOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *GetGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupBadRequest creates a GetGroupBadRequest with default headers values
func NewGetGroupBadRequest() *GetGroupBadRequest {
	return &GetGroupBadRequest{}
}

/*GetGroupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGroupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}][%d] getGroupBadRequest  %+v", 400, o.Payload)
}

func (o *GetGroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupUnauthorized creates a GetGroupUnauthorized with default headers values
func NewGetGroupUnauthorized() *GetGroupUnauthorized {
	return &GetGroupUnauthorized{}
}

/*GetGroupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGroupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}][%d] getGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupForbidden creates a GetGroupForbidden with default headers values
func NewGetGroupForbidden() *GetGroupForbidden {
	return &GetGroupForbidden{}
}

/*GetGroupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGroupForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}][%d] getGroupForbidden  %+v", 403, o.Payload)
}

func (o *GetGroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupNotFound creates a GetGroupNotFound with default headers values
func NewGetGroupNotFound() *GetGroupNotFound {
	return &GetGroupNotFound{}
}

/*GetGroupNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGroupNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}][%d] getGroupNotFound  %+v", 404, o.Payload)
}

func (o *GetGroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupRequestTimeout creates a GetGroupRequestTimeout with default headers values
func NewGetGroupRequestTimeout() *GetGroupRequestTimeout {
	return &GetGroupRequestTimeout{}
}

/*GetGroupRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGroupRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGroupRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}][%d] getGroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGroupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupRequestEntityTooLarge creates a GetGroupRequestEntityTooLarge with default headers values
func NewGetGroupRequestEntityTooLarge() *GetGroupRequestEntityTooLarge {
	return &GetGroupRequestEntityTooLarge{}
}

/*GetGroupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetGroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}][%d] getGroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupUnsupportedMediaType creates a GetGroupUnsupportedMediaType with default headers values
func NewGetGroupUnsupportedMediaType() *GetGroupUnsupportedMediaType {
	return &GetGroupUnsupportedMediaType{}
}

/*GetGroupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}][%d] getGroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupTooManyRequests creates a GetGroupTooManyRequests with default headers values
func NewGetGroupTooManyRequests() *GetGroupTooManyRequests {
	return &GetGroupTooManyRequests{}
}

/*GetGroupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGroupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGroupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}][%d] getGroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupInternalServerError creates a GetGroupInternalServerError with default headers values
func NewGetGroupInternalServerError() *GetGroupInternalServerError {
	return &GetGroupInternalServerError{}
}

/*GetGroupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGroupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}][%d] getGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupServiceUnavailable creates a GetGroupServiceUnavailable with default headers values
func NewGetGroupServiceUnavailable() *GetGroupServiceUnavailable {
	return &GetGroupServiceUnavailable{}
}

/*GetGroupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}][%d] getGroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupGatewayTimeout creates a GetGroupGatewayTimeout with default headers values
func NewGetGroupGatewayTimeout() *GetGroupGatewayTimeout {
	return &GetGroupGatewayTimeout{}
}

/*GetGroupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}][%d] getGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
