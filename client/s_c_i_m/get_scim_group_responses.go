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

// GetScimGroupReader is a Reader for the GetScimGroup structure.
type GetScimGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScimGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScimGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetScimGroupNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetScimGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScimGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScimGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScimGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScimGroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScimGroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScimGroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScimGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScimGroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScimGroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScimGroupOK creates a GetScimGroupOK with default headers values
func NewGetScimGroupOK() *GetScimGroupOK {
	return &GetScimGroupOK{}
}

/*GetScimGroupOK handles this case with default header values.

successful operation
*/
type GetScimGroupOK struct {
	Payload *models.ScimV2Group
}

func (o *GetScimGroupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/groups/{groupId}][%d] getScimGroupOK  %+v", 200, o.Payload)
}

func (o *GetScimGroupOK) GetPayload() *models.ScimV2Group {
	return o.Payload
}

func (o *GetScimGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimV2Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimGroupNotModified creates a GetScimGroupNotModified with default headers values
func NewGetScimGroupNotModified() *GetScimGroupNotModified {
	return &GetScimGroupNotModified{}
}

/*GetScimGroupNotModified handles this case with default header values.

If-Non-Match header matches current version. No content returned.
*/
type GetScimGroupNotModified struct {
}

func (o *GetScimGroupNotModified) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/groups/{groupId}][%d] getScimGroupNotModified ", 304)
}

func (o *GetScimGroupNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScimGroupBadRequest creates a GetScimGroupBadRequest with default headers values
func NewGetScimGroupBadRequest() *GetScimGroupBadRequest {
	return &GetScimGroupBadRequest{}
}

/*GetScimGroupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScimGroupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetScimGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/groups/{groupId}][%d] getScimGroupBadRequest  %+v", 400, o.Payload)
}

func (o *GetScimGroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimGroupUnauthorized creates a GetScimGroupUnauthorized with default headers values
func NewGetScimGroupUnauthorized() *GetScimGroupUnauthorized {
	return &GetScimGroupUnauthorized{}
}

/*GetScimGroupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScimGroupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetScimGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/groups/{groupId}][%d] getScimGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScimGroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimGroupForbidden creates a GetScimGroupForbidden with default headers values
func NewGetScimGroupForbidden() *GetScimGroupForbidden {
	return &GetScimGroupForbidden{}
}

/*GetScimGroupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetScimGroupForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetScimGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/groups/{groupId}][%d] getScimGroupForbidden  %+v", 403, o.Payload)
}

func (o *GetScimGroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimGroupNotFound creates a GetScimGroupNotFound with default headers values
func NewGetScimGroupNotFound() *GetScimGroupNotFound {
	return &GetScimGroupNotFound{}
}

/*GetScimGroupNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetScimGroupNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetScimGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/groups/{groupId}][%d] getScimGroupNotFound  %+v", 404, o.Payload)
}

func (o *GetScimGroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimGroupRequestEntityTooLarge creates a GetScimGroupRequestEntityTooLarge with default headers values
func NewGetScimGroupRequestEntityTooLarge() *GetScimGroupRequestEntityTooLarge {
	return &GetScimGroupRequestEntityTooLarge{}
}

/*GetScimGroupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetScimGroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetScimGroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/groups/{groupId}][%d] getScimGroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScimGroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimGroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimGroupUnsupportedMediaType creates a GetScimGroupUnsupportedMediaType with default headers values
func NewGetScimGroupUnsupportedMediaType() *GetScimGroupUnsupportedMediaType {
	return &GetScimGroupUnsupportedMediaType{}
}

/*GetScimGroupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScimGroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetScimGroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/groups/{groupId}][%d] getScimGroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScimGroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimGroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimGroupTooManyRequests creates a GetScimGroupTooManyRequests with default headers values
func NewGetScimGroupTooManyRequests() *GetScimGroupTooManyRequests {
	return &GetScimGroupTooManyRequests{}
}

/*GetScimGroupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetScimGroupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetScimGroupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/groups/{groupId}][%d] getScimGroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScimGroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimGroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimGroupInternalServerError creates a GetScimGroupInternalServerError with default headers values
func NewGetScimGroupInternalServerError() *GetScimGroupInternalServerError {
	return &GetScimGroupInternalServerError{}
}

/*GetScimGroupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScimGroupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetScimGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/groups/{groupId}][%d] getScimGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScimGroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimGroupServiceUnavailable creates a GetScimGroupServiceUnavailable with default headers values
func NewGetScimGroupServiceUnavailable() *GetScimGroupServiceUnavailable {
	return &GetScimGroupServiceUnavailable{}
}

/*GetScimGroupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScimGroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetScimGroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/groups/{groupId}][%d] getScimGroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScimGroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimGroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimGroupGatewayTimeout creates a GetScimGroupGatewayTimeout with default headers values
func NewGetScimGroupGatewayTimeout() *GetScimGroupGatewayTimeout {
	return &GetScimGroupGatewayTimeout{}
}

/*GetScimGroupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetScimGroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetScimGroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/groups/{groupId}][%d] getScimGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScimGroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimGroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}