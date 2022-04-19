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

// GetGamificationLeaderboardReader is a Reader for the GetGamificationLeaderboard structure.
type GetGamificationLeaderboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationLeaderboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationLeaderboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationLeaderboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationLeaderboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationLeaderboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationLeaderboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationLeaderboardRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationLeaderboardRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationLeaderboardUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationLeaderboardTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationLeaderboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationLeaderboardServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationLeaderboardGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationLeaderboardOK creates a GetGamificationLeaderboardOK with default headers values
func NewGetGamificationLeaderboardOK() *GetGamificationLeaderboardOK {
	return &GetGamificationLeaderboardOK{}
}

/*GetGamificationLeaderboardOK handles this case with default header values.

successful operation
*/
type GetGamificationLeaderboardOK struct {
	Payload *models.Leaderboard
}

func (o *GetGamificationLeaderboardOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/leaderboard][%d] getGamificationLeaderboardOK  %+v", 200, o.Payload)
}

func (o *GetGamificationLeaderboardOK) GetPayload() *models.Leaderboard {
	return o.Payload
}

func (o *GetGamificationLeaderboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Leaderboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationLeaderboardBadRequest creates a GetGamificationLeaderboardBadRequest with default headers values
func NewGetGamificationLeaderboardBadRequest() *GetGamificationLeaderboardBadRequest {
	return &GetGamificationLeaderboardBadRequest{}
}

/*GetGamificationLeaderboardBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationLeaderboardBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationLeaderboardBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/leaderboard][%d] getGamificationLeaderboardBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationLeaderboardBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationLeaderboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationLeaderboardUnauthorized creates a GetGamificationLeaderboardUnauthorized with default headers values
func NewGetGamificationLeaderboardUnauthorized() *GetGamificationLeaderboardUnauthorized {
	return &GetGamificationLeaderboardUnauthorized{}
}

/*GetGamificationLeaderboardUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationLeaderboardUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationLeaderboardUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/leaderboard][%d] getGamificationLeaderboardUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationLeaderboardUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationLeaderboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationLeaderboardForbidden creates a GetGamificationLeaderboardForbidden with default headers values
func NewGetGamificationLeaderboardForbidden() *GetGamificationLeaderboardForbidden {
	return &GetGamificationLeaderboardForbidden{}
}

/*GetGamificationLeaderboardForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationLeaderboardForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationLeaderboardForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/leaderboard][%d] getGamificationLeaderboardForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationLeaderboardForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationLeaderboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationLeaderboardNotFound creates a GetGamificationLeaderboardNotFound with default headers values
func NewGetGamificationLeaderboardNotFound() *GetGamificationLeaderboardNotFound {
	return &GetGamificationLeaderboardNotFound{}
}

/*GetGamificationLeaderboardNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationLeaderboardNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationLeaderboardNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/leaderboard][%d] getGamificationLeaderboardNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationLeaderboardNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationLeaderboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationLeaderboardRequestTimeout creates a GetGamificationLeaderboardRequestTimeout with default headers values
func NewGetGamificationLeaderboardRequestTimeout() *GetGamificationLeaderboardRequestTimeout {
	return &GetGamificationLeaderboardRequestTimeout{}
}

/*GetGamificationLeaderboardRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationLeaderboardRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationLeaderboardRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/leaderboard][%d] getGamificationLeaderboardRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationLeaderboardRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationLeaderboardRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationLeaderboardRequestEntityTooLarge creates a GetGamificationLeaderboardRequestEntityTooLarge with default headers values
func NewGetGamificationLeaderboardRequestEntityTooLarge() *GetGamificationLeaderboardRequestEntityTooLarge {
	return &GetGamificationLeaderboardRequestEntityTooLarge{}
}

/*GetGamificationLeaderboardRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetGamificationLeaderboardRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationLeaderboardRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/leaderboard][%d] getGamificationLeaderboardRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationLeaderboardRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationLeaderboardRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationLeaderboardUnsupportedMediaType creates a GetGamificationLeaderboardUnsupportedMediaType with default headers values
func NewGetGamificationLeaderboardUnsupportedMediaType() *GetGamificationLeaderboardUnsupportedMediaType {
	return &GetGamificationLeaderboardUnsupportedMediaType{}
}

/*GetGamificationLeaderboardUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationLeaderboardUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationLeaderboardUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/leaderboard][%d] getGamificationLeaderboardUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationLeaderboardUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationLeaderboardUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationLeaderboardTooManyRequests creates a GetGamificationLeaderboardTooManyRequests with default headers values
func NewGetGamificationLeaderboardTooManyRequests() *GetGamificationLeaderboardTooManyRequests {
	return &GetGamificationLeaderboardTooManyRequests{}
}

/*GetGamificationLeaderboardTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationLeaderboardTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationLeaderboardTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/leaderboard][%d] getGamificationLeaderboardTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationLeaderboardTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationLeaderboardTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationLeaderboardInternalServerError creates a GetGamificationLeaderboardInternalServerError with default headers values
func NewGetGamificationLeaderboardInternalServerError() *GetGamificationLeaderboardInternalServerError {
	return &GetGamificationLeaderboardInternalServerError{}
}

/*GetGamificationLeaderboardInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationLeaderboardInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationLeaderboardInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/leaderboard][%d] getGamificationLeaderboardInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationLeaderboardInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationLeaderboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationLeaderboardServiceUnavailable creates a GetGamificationLeaderboardServiceUnavailable with default headers values
func NewGetGamificationLeaderboardServiceUnavailable() *GetGamificationLeaderboardServiceUnavailable {
	return &GetGamificationLeaderboardServiceUnavailable{}
}

/*GetGamificationLeaderboardServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationLeaderboardServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationLeaderboardServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/leaderboard][%d] getGamificationLeaderboardServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationLeaderboardServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationLeaderboardServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationLeaderboardGatewayTimeout creates a GetGamificationLeaderboardGatewayTimeout with default headers values
func NewGetGamificationLeaderboardGatewayTimeout() *GetGamificationLeaderboardGatewayTimeout {
	return &GetGamificationLeaderboardGatewayTimeout{}
}

/*GetGamificationLeaderboardGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationLeaderboardGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationLeaderboardGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/leaderboard][%d] getGamificationLeaderboardGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationLeaderboardGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationLeaderboardGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
