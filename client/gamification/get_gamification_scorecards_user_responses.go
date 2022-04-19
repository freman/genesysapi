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

// GetGamificationScorecardsUserReader is a Reader for the GetGamificationScorecardsUser structure.
type GetGamificationScorecardsUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationScorecardsUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationScorecardsUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationScorecardsUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationScorecardsUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationScorecardsUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationScorecardsUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationScorecardsUserRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationScorecardsUserRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationScorecardsUserUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationScorecardsUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationScorecardsUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationScorecardsUserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationScorecardsUserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationScorecardsUserOK creates a GetGamificationScorecardsUserOK with default headers values
func NewGetGamificationScorecardsUserOK() *GetGamificationScorecardsUserOK {
	return &GetGamificationScorecardsUserOK{}
}

/*GetGamificationScorecardsUserOK handles this case with default header values.

successful operation
*/
type GetGamificationScorecardsUserOK struct {
	Payload *models.WorkdayMetricListing
}

func (o *GetGamificationScorecardsUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}][%d] getGamificationScorecardsUserOK  %+v", 200, o.Payload)
}

func (o *GetGamificationScorecardsUserOK) GetPayload() *models.WorkdayMetricListing {
	return o.Payload
}

func (o *GetGamificationScorecardsUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkdayMetricListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserBadRequest creates a GetGamificationScorecardsUserBadRequest with default headers values
func NewGetGamificationScorecardsUserBadRequest() *GetGamificationScorecardsUserBadRequest {
	return &GetGamificationScorecardsUserBadRequest{}
}

/*GetGamificationScorecardsUserBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationScorecardsUserBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}][%d] getGamificationScorecardsUserBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationScorecardsUserBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserUnauthorized creates a GetGamificationScorecardsUserUnauthorized with default headers values
func NewGetGamificationScorecardsUserUnauthorized() *GetGamificationScorecardsUserUnauthorized {
	return &GetGamificationScorecardsUserUnauthorized{}
}

/*GetGamificationScorecardsUserUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationScorecardsUserUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}][%d] getGamificationScorecardsUserUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationScorecardsUserUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserForbidden creates a GetGamificationScorecardsUserForbidden with default headers values
func NewGetGamificationScorecardsUserForbidden() *GetGamificationScorecardsUserForbidden {
	return &GetGamificationScorecardsUserForbidden{}
}

/*GetGamificationScorecardsUserForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationScorecardsUserForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}][%d] getGamificationScorecardsUserForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationScorecardsUserForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserNotFound creates a GetGamificationScorecardsUserNotFound with default headers values
func NewGetGamificationScorecardsUserNotFound() *GetGamificationScorecardsUserNotFound {
	return &GetGamificationScorecardsUserNotFound{}
}

/*GetGamificationScorecardsUserNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationScorecardsUserNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}][%d] getGamificationScorecardsUserNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationScorecardsUserNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserRequestTimeout creates a GetGamificationScorecardsUserRequestTimeout with default headers values
func NewGetGamificationScorecardsUserRequestTimeout() *GetGamificationScorecardsUserRequestTimeout {
	return &GetGamificationScorecardsUserRequestTimeout{}
}

/*GetGamificationScorecardsUserRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationScorecardsUserRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}][%d] getGamificationScorecardsUserRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationScorecardsUserRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserRequestEntityTooLarge creates a GetGamificationScorecardsUserRequestEntityTooLarge with default headers values
func NewGetGamificationScorecardsUserRequestEntityTooLarge() *GetGamificationScorecardsUserRequestEntityTooLarge {
	return &GetGamificationScorecardsUserRequestEntityTooLarge{}
}

/*GetGamificationScorecardsUserRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetGamificationScorecardsUserRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}][%d] getGamificationScorecardsUserRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationScorecardsUserRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserUnsupportedMediaType creates a GetGamificationScorecardsUserUnsupportedMediaType with default headers values
func NewGetGamificationScorecardsUserUnsupportedMediaType() *GetGamificationScorecardsUserUnsupportedMediaType {
	return &GetGamificationScorecardsUserUnsupportedMediaType{}
}

/*GetGamificationScorecardsUserUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationScorecardsUserUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}][%d] getGamificationScorecardsUserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationScorecardsUserUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserTooManyRequests creates a GetGamificationScorecardsUserTooManyRequests with default headers values
func NewGetGamificationScorecardsUserTooManyRequests() *GetGamificationScorecardsUserTooManyRequests {
	return &GetGamificationScorecardsUserTooManyRequests{}
}

/*GetGamificationScorecardsUserTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationScorecardsUserTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}][%d] getGamificationScorecardsUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationScorecardsUserTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserInternalServerError creates a GetGamificationScorecardsUserInternalServerError with default headers values
func NewGetGamificationScorecardsUserInternalServerError() *GetGamificationScorecardsUserInternalServerError {
	return &GetGamificationScorecardsUserInternalServerError{}
}

/*GetGamificationScorecardsUserInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationScorecardsUserInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}][%d] getGamificationScorecardsUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationScorecardsUserInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserServiceUnavailable creates a GetGamificationScorecardsUserServiceUnavailable with default headers values
func NewGetGamificationScorecardsUserServiceUnavailable() *GetGamificationScorecardsUserServiceUnavailable {
	return &GetGamificationScorecardsUserServiceUnavailable{}
}

/*GetGamificationScorecardsUserServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationScorecardsUserServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}][%d] getGamificationScorecardsUserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationScorecardsUserServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserGatewayTimeout creates a GetGamificationScorecardsUserGatewayTimeout with default headers values
func NewGetGamificationScorecardsUserGatewayTimeout() *GetGamificationScorecardsUserGatewayTimeout {
	return &GetGamificationScorecardsUserGatewayTimeout{}
}

/*GetGamificationScorecardsUserGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationScorecardsUserGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}][%d] getGamificationScorecardsUserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationScorecardsUserGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
