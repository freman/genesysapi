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

// GetGamificationScorecardsUserAttendanceReader is a Reader for the GetGamificationScorecardsUserAttendance structure.
type GetGamificationScorecardsUserAttendanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationScorecardsUserAttendanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationScorecardsUserAttendanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationScorecardsUserAttendanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationScorecardsUserAttendanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationScorecardsUserAttendanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationScorecardsUserAttendanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationScorecardsUserAttendanceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationScorecardsUserAttendanceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationScorecardsUserAttendanceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationScorecardsUserAttendanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationScorecardsUserAttendanceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationScorecardsUserAttendanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationScorecardsUserAttendanceOK creates a GetGamificationScorecardsUserAttendanceOK with default headers values
func NewGetGamificationScorecardsUserAttendanceOK() *GetGamificationScorecardsUserAttendanceOK {
	return &GetGamificationScorecardsUserAttendanceOK{}
}

/*GetGamificationScorecardsUserAttendanceOK handles this case with default header values.

successful operation
*/
type GetGamificationScorecardsUserAttendanceOK struct {
	Payload *models.AttendanceStatusListing
}

func (o *GetGamificationScorecardsUserAttendanceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/attendance][%d] getGamificationScorecardsUserAttendanceOK  %+v", 200, o.Payload)
}

func (o *GetGamificationScorecardsUserAttendanceOK) GetPayload() *models.AttendanceStatusListing {
	return o.Payload
}

func (o *GetGamificationScorecardsUserAttendanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttendanceStatusListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserAttendanceBadRequest creates a GetGamificationScorecardsUserAttendanceBadRequest with default headers values
func NewGetGamificationScorecardsUserAttendanceBadRequest() *GetGamificationScorecardsUserAttendanceBadRequest {
	return &GetGamificationScorecardsUserAttendanceBadRequest{}
}

/*GetGamificationScorecardsUserAttendanceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationScorecardsUserAttendanceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserAttendanceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/attendance][%d] getGamificationScorecardsUserAttendanceBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationScorecardsUserAttendanceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserAttendanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserAttendanceUnauthorized creates a GetGamificationScorecardsUserAttendanceUnauthorized with default headers values
func NewGetGamificationScorecardsUserAttendanceUnauthorized() *GetGamificationScorecardsUserAttendanceUnauthorized {
	return &GetGamificationScorecardsUserAttendanceUnauthorized{}
}

/*GetGamificationScorecardsUserAttendanceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationScorecardsUserAttendanceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserAttendanceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/attendance][%d] getGamificationScorecardsUserAttendanceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationScorecardsUserAttendanceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserAttendanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserAttendanceForbidden creates a GetGamificationScorecardsUserAttendanceForbidden with default headers values
func NewGetGamificationScorecardsUserAttendanceForbidden() *GetGamificationScorecardsUserAttendanceForbidden {
	return &GetGamificationScorecardsUserAttendanceForbidden{}
}

/*GetGamificationScorecardsUserAttendanceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationScorecardsUserAttendanceForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserAttendanceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/attendance][%d] getGamificationScorecardsUserAttendanceForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationScorecardsUserAttendanceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserAttendanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserAttendanceNotFound creates a GetGamificationScorecardsUserAttendanceNotFound with default headers values
func NewGetGamificationScorecardsUserAttendanceNotFound() *GetGamificationScorecardsUserAttendanceNotFound {
	return &GetGamificationScorecardsUserAttendanceNotFound{}
}

/*GetGamificationScorecardsUserAttendanceNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationScorecardsUserAttendanceNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserAttendanceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/attendance][%d] getGamificationScorecardsUserAttendanceNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationScorecardsUserAttendanceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserAttendanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserAttendanceRequestEntityTooLarge creates a GetGamificationScorecardsUserAttendanceRequestEntityTooLarge with default headers values
func NewGetGamificationScorecardsUserAttendanceRequestEntityTooLarge() *GetGamificationScorecardsUserAttendanceRequestEntityTooLarge {
	return &GetGamificationScorecardsUserAttendanceRequestEntityTooLarge{}
}

/*GetGamificationScorecardsUserAttendanceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGamificationScorecardsUserAttendanceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserAttendanceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/attendance][%d] getGamificationScorecardsUserAttendanceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationScorecardsUserAttendanceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserAttendanceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserAttendanceUnsupportedMediaType creates a GetGamificationScorecardsUserAttendanceUnsupportedMediaType with default headers values
func NewGetGamificationScorecardsUserAttendanceUnsupportedMediaType() *GetGamificationScorecardsUserAttendanceUnsupportedMediaType {
	return &GetGamificationScorecardsUserAttendanceUnsupportedMediaType{}
}

/*GetGamificationScorecardsUserAttendanceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationScorecardsUserAttendanceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserAttendanceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/attendance][%d] getGamificationScorecardsUserAttendanceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationScorecardsUserAttendanceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserAttendanceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserAttendanceTooManyRequests creates a GetGamificationScorecardsUserAttendanceTooManyRequests with default headers values
func NewGetGamificationScorecardsUserAttendanceTooManyRequests() *GetGamificationScorecardsUserAttendanceTooManyRequests {
	return &GetGamificationScorecardsUserAttendanceTooManyRequests{}
}

/*GetGamificationScorecardsUserAttendanceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetGamificationScorecardsUserAttendanceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserAttendanceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/attendance][%d] getGamificationScorecardsUserAttendanceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationScorecardsUserAttendanceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserAttendanceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserAttendanceInternalServerError creates a GetGamificationScorecardsUserAttendanceInternalServerError with default headers values
func NewGetGamificationScorecardsUserAttendanceInternalServerError() *GetGamificationScorecardsUserAttendanceInternalServerError {
	return &GetGamificationScorecardsUserAttendanceInternalServerError{}
}

/*GetGamificationScorecardsUserAttendanceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationScorecardsUserAttendanceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserAttendanceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/attendance][%d] getGamificationScorecardsUserAttendanceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationScorecardsUserAttendanceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserAttendanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserAttendanceServiceUnavailable creates a GetGamificationScorecardsUserAttendanceServiceUnavailable with default headers values
func NewGetGamificationScorecardsUserAttendanceServiceUnavailable() *GetGamificationScorecardsUserAttendanceServiceUnavailable {
	return &GetGamificationScorecardsUserAttendanceServiceUnavailable{}
}

/*GetGamificationScorecardsUserAttendanceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationScorecardsUserAttendanceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserAttendanceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/attendance][%d] getGamificationScorecardsUserAttendanceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationScorecardsUserAttendanceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserAttendanceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserAttendanceGatewayTimeout creates a GetGamificationScorecardsUserAttendanceGatewayTimeout with default headers values
func NewGetGamificationScorecardsUserAttendanceGatewayTimeout() *GetGamificationScorecardsUserAttendanceGatewayTimeout {
	return &GetGamificationScorecardsUserAttendanceGatewayTimeout{}
}

/*GetGamificationScorecardsUserAttendanceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationScorecardsUserAttendanceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserAttendanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/attendance][%d] getGamificationScorecardsUserAttendanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationScorecardsUserAttendanceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserAttendanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
