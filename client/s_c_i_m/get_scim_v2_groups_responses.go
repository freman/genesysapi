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

// GetScimV2GroupsReader is a Reader for the GetScimV2Groups structure.
type GetScimV2GroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScimV2GroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScimV2GroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScimV2GroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScimV2GroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScimV2GroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScimV2GroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScimV2GroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScimV2GroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScimV2GroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScimV2GroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScimV2GroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScimV2GroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScimV2GroupsOK creates a GetScimV2GroupsOK with default headers values
func NewGetScimV2GroupsOK() *GetScimV2GroupsOK {
	return &GetScimV2GroupsOK{}
}

/*GetScimV2GroupsOK handles this case with default header values.

successful operation
*/
type GetScimV2GroupsOK struct {
	Payload *models.ScimGroupListResponse
}

func (o *GetScimV2GroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/groups][%d] getScimV2GroupsOK  %+v", 200, o.Payload)
}

func (o *GetScimV2GroupsOK) GetPayload() *models.ScimGroupListResponse {
	return o.Payload
}

func (o *GetScimV2GroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimGroupListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2GroupsBadRequest creates a GetScimV2GroupsBadRequest with default headers values
func NewGetScimV2GroupsBadRequest() *GetScimV2GroupsBadRequest {
	return &GetScimV2GroupsBadRequest{}
}

/*GetScimV2GroupsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScimV2GroupsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2GroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/groups][%d] getScimV2GroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetScimV2GroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2GroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2GroupsUnauthorized creates a GetScimV2GroupsUnauthorized with default headers values
func NewGetScimV2GroupsUnauthorized() *GetScimV2GroupsUnauthorized {
	return &GetScimV2GroupsUnauthorized{}
}

/*GetScimV2GroupsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScimV2GroupsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2GroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/groups][%d] getScimV2GroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScimV2GroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2GroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2GroupsForbidden creates a GetScimV2GroupsForbidden with default headers values
func NewGetScimV2GroupsForbidden() *GetScimV2GroupsForbidden {
	return &GetScimV2GroupsForbidden{}
}

/*GetScimV2GroupsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetScimV2GroupsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2GroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/groups][%d] getScimV2GroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetScimV2GroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2GroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2GroupsNotFound creates a GetScimV2GroupsNotFound with default headers values
func NewGetScimV2GroupsNotFound() *GetScimV2GroupsNotFound {
	return &GetScimV2GroupsNotFound{}
}

/*GetScimV2GroupsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetScimV2GroupsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2GroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/groups][%d] getScimV2GroupsNotFound  %+v", 404, o.Payload)
}

func (o *GetScimV2GroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2GroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2GroupsRequestEntityTooLarge creates a GetScimV2GroupsRequestEntityTooLarge with default headers values
func NewGetScimV2GroupsRequestEntityTooLarge() *GetScimV2GroupsRequestEntityTooLarge {
	return &GetScimV2GroupsRequestEntityTooLarge{}
}

/*GetScimV2GroupsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetScimV2GroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2GroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/groups][%d] getScimV2GroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScimV2GroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2GroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2GroupsUnsupportedMediaType creates a GetScimV2GroupsUnsupportedMediaType with default headers values
func NewGetScimV2GroupsUnsupportedMediaType() *GetScimV2GroupsUnsupportedMediaType {
	return &GetScimV2GroupsUnsupportedMediaType{}
}

/*GetScimV2GroupsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScimV2GroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2GroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/groups][%d] getScimV2GroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScimV2GroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2GroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2GroupsTooManyRequests creates a GetScimV2GroupsTooManyRequests with default headers values
func NewGetScimV2GroupsTooManyRequests() *GetScimV2GroupsTooManyRequests {
	return &GetScimV2GroupsTooManyRequests{}
}

/*GetScimV2GroupsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetScimV2GroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2GroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/groups][%d] getScimV2GroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScimV2GroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2GroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2GroupsInternalServerError creates a GetScimV2GroupsInternalServerError with default headers values
func NewGetScimV2GroupsInternalServerError() *GetScimV2GroupsInternalServerError {
	return &GetScimV2GroupsInternalServerError{}
}

/*GetScimV2GroupsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScimV2GroupsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2GroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/groups][%d] getScimV2GroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScimV2GroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2GroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2GroupsServiceUnavailable creates a GetScimV2GroupsServiceUnavailable with default headers values
func NewGetScimV2GroupsServiceUnavailable() *GetScimV2GroupsServiceUnavailable {
	return &GetScimV2GroupsServiceUnavailable{}
}

/*GetScimV2GroupsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScimV2GroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2GroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/groups][%d] getScimV2GroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScimV2GroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2GroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2GroupsGatewayTimeout creates a GetScimV2GroupsGatewayTimeout with default headers values
func NewGetScimV2GroupsGatewayTimeout() *GetScimV2GroupsGatewayTimeout {
	return &GetScimV2GroupsGatewayTimeout{}
}

/*GetScimV2GroupsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetScimV2GroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2GroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/groups][%d] getScimV2GroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScimV2GroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2GroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
