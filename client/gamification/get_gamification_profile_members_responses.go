// Code generated by go-swagger; DO NOT EDIT.

package gamification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetGamificationProfileMembersReader is a Reader for the GetGamificationProfileMembers structure.
type GetGamificationProfileMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationProfileMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationProfileMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationProfileMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationProfileMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationProfileMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationProfileMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationProfileMembersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationProfileMembersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationProfileMembersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationProfileMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationProfileMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationProfileMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationProfileMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationProfileMembersOK creates a GetGamificationProfileMembersOK with default headers values
func NewGetGamificationProfileMembersOK() *GetGamificationProfileMembersOK {
	return &GetGamificationProfileMembersOK{}
}

/*GetGamificationProfileMembersOK handles this case with default header values.

successful operation
*/
type GetGamificationProfileMembersOK struct {
	Payload *models.MemberListing
}

func (o *GetGamificationProfileMembersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{performanceProfileId}/members][%d] getGamificationProfileMembersOK  %+v", 200, o.Payload)
}

func (o *GetGamificationProfileMembersOK) GetPayload() *models.MemberListing {
	return o.Payload
}

func (o *GetGamificationProfileMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMembersBadRequest creates a GetGamificationProfileMembersBadRequest with default headers values
func NewGetGamificationProfileMembersBadRequest() *GetGamificationProfileMembersBadRequest {
	return &GetGamificationProfileMembersBadRequest{}
}

/*GetGamificationProfileMembersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationProfileMembersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{performanceProfileId}/members][%d] getGamificationProfileMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationProfileMembersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMembersUnauthorized creates a GetGamificationProfileMembersUnauthorized with default headers values
func NewGetGamificationProfileMembersUnauthorized() *GetGamificationProfileMembersUnauthorized {
	return &GetGamificationProfileMembersUnauthorized{}
}

/*GetGamificationProfileMembersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationProfileMembersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{performanceProfileId}/members][%d] getGamificationProfileMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationProfileMembersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMembersForbidden creates a GetGamificationProfileMembersForbidden with default headers values
func NewGetGamificationProfileMembersForbidden() *GetGamificationProfileMembersForbidden {
	return &GetGamificationProfileMembersForbidden{}
}

/*GetGamificationProfileMembersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationProfileMembersForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{performanceProfileId}/members][%d] getGamificationProfileMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationProfileMembersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMembersNotFound creates a GetGamificationProfileMembersNotFound with default headers values
func NewGetGamificationProfileMembersNotFound() *GetGamificationProfileMembersNotFound {
	return &GetGamificationProfileMembersNotFound{}
}

/*GetGamificationProfileMembersNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationProfileMembersNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{performanceProfileId}/members][%d] getGamificationProfileMembersNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationProfileMembersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMembersRequestTimeout creates a GetGamificationProfileMembersRequestTimeout with default headers values
func NewGetGamificationProfileMembersRequestTimeout() *GetGamificationProfileMembersRequestTimeout {
	return &GetGamificationProfileMembersRequestTimeout{}
}

/*GetGamificationProfileMembersRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationProfileMembersRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMembersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{performanceProfileId}/members][%d] getGamificationProfileMembersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationProfileMembersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMembersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMembersRequestEntityTooLarge creates a GetGamificationProfileMembersRequestEntityTooLarge with default headers values
func NewGetGamificationProfileMembersRequestEntityTooLarge() *GetGamificationProfileMembersRequestEntityTooLarge {
	return &GetGamificationProfileMembersRequestEntityTooLarge{}
}

/*GetGamificationProfileMembersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGamificationProfileMembersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMembersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{performanceProfileId}/members][%d] getGamificationProfileMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationProfileMembersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMembersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMembersUnsupportedMediaType creates a GetGamificationProfileMembersUnsupportedMediaType with default headers values
func NewGetGamificationProfileMembersUnsupportedMediaType() *GetGamificationProfileMembersUnsupportedMediaType {
	return &GetGamificationProfileMembersUnsupportedMediaType{}
}

/*GetGamificationProfileMembersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationProfileMembersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMembersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{performanceProfileId}/members][%d] getGamificationProfileMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationProfileMembersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMembersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMembersTooManyRequests creates a GetGamificationProfileMembersTooManyRequests with default headers values
func NewGetGamificationProfileMembersTooManyRequests() *GetGamificationProfileMembersTooManyRequests {
	return &GetGamificationProfileMembersTooManyRequests{}
}

/*GetGamificationProfileMembersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationProfileMembersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{performanceProfileId}/members][%d] getGamificationProfileMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationProfileMembersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMembersInternalServerError creates a GetGamificationProfileMembersInternalServerError with default headers values
func NewGetGamificationProfileMembersInternalServerError() *GetGamificationProfileMembersInternalServerError {
	return &GetGamificationProfileMembersInternalServerError{}
}

/*GetGamificationProfileMembersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationProfileMembersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{performanceProfileId}/members][%d] getGamificationProfileMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationProfileMembersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMembersServiceUnavailable creates a GetGamificationProfileMembersServiceUnavailable with default headers values
func NewGetGamificationProfileMembersServiceUnavailable() *GetGamificationProfileMembersServiceUnavailable {
	return &GetGamificationProfileMembersServiceUnavailable{}
}

/*GetGamificationProfileMembersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationProfileMembersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{performanceProfileId}/members][%d] getGamificationProfileMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationProfileMembersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMembersGatewayTimeout creates a GetGamificationProfileMembersGatewayTimeout with default headers values
func NewGetGamificationProfileMembersGatewayTimeout() *GetGamificationProfileMembersGatewayTimeout {
	return &GetGamificationProfileMembersGatewayTimeout{}
}

/*GetGamificationProfileMembersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationProfileMembersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{performanceProfileId}/members][%d] getGamificationProfileMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationProfileMembersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
