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

// GetGroupProfileReader is a Reader for the GetGroupProfile structure.
type GetGroupProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGroupProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGroupProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGroupProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGroupProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGroupProfileRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGroupProfileUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGroupProfileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGroupProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGroupProfileServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGroupProfileGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupProfileOK creates a GetGroupProfileOK with default headers values
func NewGetGroupProfileOK() *GetGroupProfileOK {
	return &GetGroupProfileOK{}
}

/*GetGroupProfileOK handles this case with default header values.

successful operation
*/
type GetGroupProfileOK struct {
	Payload *models.GroupProfile
}

func (o *GetGroupProfileOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/profile][%d] getGroupProfileOK  %+v", 200, o.Payload)
}

func (o *GetGroupProfileOK) GetPayload() *models.GroupProfile {
	return o.Payload
}

func (o *GetGroupProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupProfileBadRequest creates a GetGroupProfileBadRequest with default headers values
func NewGetGroupProfileBadRequest() *GetGroupProfileBadRequest {
	return &GetGroupProfileBadRequest{}
}

/*GetGroupProfileBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGroupProfileBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGroupProfileBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/profile][%d] getGroupProfileBadRequest  %+v", 400, o.Payload)
}

func (o *GetGroupProfileBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupProfileUnauthorized creates a GetGroupProfileUnauthorized with default headers values
func NewGetGroupProfileUnauthorized() *GetGroupProfileUnauthorized {
	return &GetGroupProfileUnauthorized{}
}

/*GetGroupProfileUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGroupProfileUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGroupProfileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/profile][%d] getGroupProfileUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGroupProfileUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupProfileForbidden creates a GetGroupProfileForbidden with default headers values
func NewGetGroupProfileForbidden() *GetGroupProfileForbidden {
	return &GetGroupProfileForbidden{}
}

/*GetGroupProfileForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGroupProfileForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGroupProfileForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/profile][%d] getGroupProfileForbidden  %+v", 403, o.Payload)
}

func (o *GetGroupProfileForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupProfileNotFound creates a GetGroupProfileNotFound with default headers values
func NewGetGroupProfileNotFound() *GetGroupProfileNotFound {
	return &GetGroupProfileNotFound{}
}

/*GetGroupProfileNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGroupProfileNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGroupProfileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/profile][%d] getGroupProfileNotFound  %+v", 404, o.Payload)
}

func (o *GetGroupProfileNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupProfileRequestEntityTooLarge creates a GetGroupProfileRequestEntityTooLarge with default headers values
func NewGetGroupProfileRequestEntityTooLarge() *GetGroupProfileRequestEntityTooLarge {
	return &GetGroupProfileRequestEntityTooLarge{}
}

/*GetGroupProfileRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGroupProfileRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGroupProfileRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/profile][%d] getGroupProfileRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGroupProfileRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupProfileRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupProfileUnsupportedMediaType creates a GetGroupProfileUnsupportedMediaType with default headers values
func NewGetGroupProfileUnsupportedMediaType() *GetGroupProfileUnsupportedMediaType {
	return &GetGroupProfileUnsupportedMediaType{}
}

/*GetGroupProfileUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGroupProfileUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGroupProfileUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/profile][%d] getGroupProfileUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGroupProfileUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupProfileUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupProfileTooManyRequests creates a GetGroupProfileTooManyRequests with default headers values
func NewGetGroupProfileTooManyRequests() *GetGroupProfileTooManyRequests {
	return &GetGroupProfileTooManyRequests{}
}

/*GetGroupProfileTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetGroupProfileTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGroupProfileTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/profile][%d] getGroupProfileTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGroupProfileTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupProfileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupProfileInternalServerError creates a GetGroupProfileInternalServerError with default headers values
func NewGetGroupProfileInternalServerError() *GetGroupProfileInternalServerError {
	return &GetGroupProfileInternalServerError{}
}

/*GetGroupProfileInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGroupProfileInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGroupProfileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/profile][%d] getGroupProfileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGroupProfileInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupProfileServiceUnavailable creates a GetGroupProfileServiceUnavailable with default headers values
func NewGetGroupProfileServiceUnavailable() *GetGroupProfileServiceUnavailable {
	return &GetGroupProfileServiceUnavailable{}
}

/*GetGroupProfileServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGroupProfileServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGroupProfileServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/profile][%d] getGroupProfileServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGroupProfileServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupProfileServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupProfileGatewayTimeout creates a GetGroupProfileGatewayTimeout with default headers values
func NewGetGroupProfileGatewayTimeout() *GetGroupProfileGatewayTimeout {
	return &GetGroupProfileGatewayTimeout{}
}

/*GetGroupProfileGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGroupProfileGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGroupProfileGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/profile][%d] getGroupProfileGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGroupProfileGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupProfileGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
