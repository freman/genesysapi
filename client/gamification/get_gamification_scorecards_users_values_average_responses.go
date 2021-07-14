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

// GetGamificationScorecardsUsersValuesAverageReader is a Reader for the GetGamificationScorecardsUsersValuesAverage structure.
type GetGamificationScorecardsUsersValuesAverageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationScorecardsUsersValuesAverageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationScorecardsUsersValuesAverageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationScorecardsUsersValuesAverageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationScorecardsUsersValuesAverageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationScorecardsUsersValuesAverageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationScorecardsUsersValuesAverageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationScorecardsUsersValuesAverageRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationScorecardsUsersValuesAverageRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationScorecardsUsersValuesAverageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationScorecardsUsersValuesAverageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationScorecardsUsersValuesAverageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationScorecardsUsersValuesAverageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationScorecardsUsersValuesAverageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationScorecardsUsersValuesAverageOK creates a GetGamificationScorecardsUsersValuesAverageOK with default headers values
func NewGetGamificationScorecardsUsersValuesAverageOK() *GetGamificationScorecardsUsersValuesAverageOK {
	return &GetGamificationScorecardsUsersValuesAverageOK{}
}

/*GetGamificationScorecardsUsersValuesAverageOK handles this case with default header values.

successful operation
*/
type GetGamificationScorecardsUsersValuesAverageOK struct {
	Payload *models.SingleWorkdayAverageValues
}

func (o *GetGamificationScorecardsUsersValuesAverageOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/values/average][%d] getGamificationScorecardsUsersValuesAverageOK  %+v", 200, o.Payload)
}

func (o *GetGamificationScorecardsUsersValuesAverageOK) GetPayload() *models.SingleWorkdayAverageValues {
	return o.Payload
}

func (o *GetGamificationScorecardsUsersValuesAverageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SingleWorkdayAverageValues)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUsersValuesAverageBadRequest creates a GetGamificationScorecardsUsersValuesAverageBadRequest with default headers values
func NewGetGamificationScorecardsUsersValuesAverageBadRequest() *GetGamificationScorecardsUsersValuesAverageBadRequest {
	return &GetGamificationScorecardsUsersValuesAverageBadRequest{}
}

/*GetGamificationScorecardsUsersValuesAverageBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationScorecardsUsersValuesAverageBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUsersValuesAverageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/values/average][%d] getGamificationScorecardsUsersValuesAverageBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationScorecardsUsersValuesAverageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUsersValuesAverageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUsersValuesAverageUnauthorized creates a GetGamificationScorecardsUsersValuesAverageUnauthorized with default headers values
func NewGetGamificationScorecardsUsersValuesAverageUnauthorized() *GetGamificationScorecardsUsersValuesAverageUnauthorized {
	return &GetGamificationScorecardsUsersValuesAverageUnauthorized{}
}

/*GetGamificationScorecardsUsersValuesAverageUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationScorecardsUsersValuesAverageUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUsersValuesAverageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/values/average][%d] getGamificationScorecardsUsersValuesAverageUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationScorecardsUsersValuesAverageUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUsersValuesAverageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUsersValuesAverageForbidden creates a GetGamificationScorecardsUsersValuesAverageForbidden with default headers values
func NewGetGamificationScorecardsUsersValuesAverageForbidden() *GetGamificationScorecardsUsersValuesAverageForbidden {
	return &GetGamificationScorecardsUsersValuesAverageForbidden{}
}

/*GetGamificationScorecardsUsersValuesAverageForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationScorecardsUsersValuesAverageForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUsersValuesAverageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/values/average][%d] getGamificationScorecardsUsersValuesAverageForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationScorecardsUsersValuesAverageForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUsersValuesAverageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUsersValuesAverageNotFound creates a GetGamificationScorecardsUsersValuesAverageNotFound with default headers values
func NewGetGamificationScorecardsUsersValuesAverageNotFound() *GetGamificationScorecardsUsersValuesAverageNotFound {
	return &GetGamificationScorecardsUsersValuesAverageNotFound{}
}

/*GetGamificationScorecardsUsersValuesAverageNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationScorecardsUsersValuesAverageNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUsersValuesAverageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/values/average][%d] getGamificationScorecardsUsersValuesAverageNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationScorecardsUsersValuesAverageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUsersValuesAverageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUsersValuesAverageRequestTimeout creates a GetGamificationScorecardsUsersValuesAverageRequestTimeout with default headers values
func NewGetGamificationScorecardsUsersValuesAverageRequestTimeout() *GetGamificationScorecardsUsersValuesAverageRequestTimeout {
	return &GetGamificationScorecardsUsersValuesAverageRequestTimeout{}
}

/*GetGamificationScorecardsUsersValuesAverageRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationScorecardsUsersValuesAverageRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUsersValuesAverageRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/values/average][%d] getGamificationScorecardsUsersValuesAverageRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationScorecardsUsersValuesAverageRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUsersValuesAverageRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUsersValuesAverageRequestEntityTooLarge creates a GetGamificationScorecardsUsersValuesAverageRequestEntityTooLarge with default headers values
func NewGetGamificationScorecardsUsersValuesAverageRequestEntityTooLarge() *GetGamificationScorecardsUsersValuesAverageRequestEntityTooLarge {
	return &GetGamificationScorecardsUsersValuesAverageRequestEntityTooLarge{}
}

/*GetGamificationScorecardsUsersValuesAverageRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGamificationScorecardsUsersValuesAverageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUsersValuesAverageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/values/average][%d] getGamificationScorecardsUsersValuesAverageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationScorecardsUsersValuesAverageRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUsersValuesAverageRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUsersValuesAverageUnsupportedMediaType creates a GetGamificationScorecardsUsersValuesAverageUnsupportedMediaType with default headers values
func NewGetGamificationScorecardsUsersValuesAverageUnsupportedMediaType() *GetGamificationScorecardsUsersValuesAverageUnsupportedMediaType {
	return &GetGamificationScorecardsUsersValuesAverageUnsupportedMediaType{}
}

/*GetGamificationScorecardsUsersValuesAverageUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationScorecardsUsersValuesAverageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUsersValuesAverageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/values/average][%d] getGamificationScorecardsUsersValuesAverageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationScorecardsUsersValuesAverageUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUsersValuesAverageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUsersValuesAverageTooManyRequests creates a GetGamificationScorecardsUsersValuesAverageTooManyRequests with default headers values
func NewGetGamificationScorecardsUsersValuesAverageTooManyRequests() *GetGamificationScorecardsUsersValuesAverageTooManyRequests {
	return &GetGamificationScorecardsUsersValuesAverageTooManyRequests{}
}

/*GetGamificationScorecardsUsersValuesAverageTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationScorecardsUsersValuesAverageTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUsersValuesAverageTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/values/average][%d] getGamificationScorecardsUsersValuesAverageTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationScorecardsUsersValuesAverageTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUsersValuesAverageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUsersValuesAverageInternalServerError creates a GetGamificationScorecardsUsersValuesAverageInternalServerError with default headers values
func NewGetGamificationScorecardsUsersValuesAverageInternalServerError() *GetGamificationScorecardsUsersValuesAverageInternalServerError {
	return &GetGamificationScorecardsUsersValuesAverageInternalServerError{}
}

/*GetGamificationScorecardsUsersValuesAverageInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationScorecardsUsersValuesAverageInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUsersValuesAverageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/values/average][%d] getGamificationScorecardsUsersValuesAverageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationScorecardsUsersValuesAverageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUsersValuesAverageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUsersValuesAverageServiceUnavailable creates a GetGamificationScorecardsUsersValuesAverageServiceUnavailable with default headers values
func NewGetGamificationScorecardsUsersValuesAverageServiceUnavailable() *GetGamificationScorecardsUsersValuesAverageServiceUnavailable {
	return &GetGamificationScorecardsUsersValuesAverageServiceUnavailable{}
}

/*GetGamificationScorecardsUsersValuesAverageServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationScorecardsUsersValuesAverageServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUsersValuesAverageServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/values/average][%d] getGamificationScorecardsUsersValuesAverageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationScorecardsUsersValuesAverageServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUsersValuesAverageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUsersValuesAverageGatewayTimeout creates a GetGamificationScorecardsUsersValuesAverageGatewayTimeout with default headers values
func NewGetGamificationScorecardsUsersValuesAverageGatewayTimeout() *GetGamificationScorecardsUsersValuesAverageGatewayTimeout {
	return &GetGamificationScorecardsUsersValuesAverageGatewayTimeout{}
}

/*GetGamificationScorecardsUsersValuesAverageGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationScorecardsUsersValuesAverageGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUsersValuesAverageGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/values/average][%d] getGamificationScorecardsUsersValuesAverageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationScorecardsUsersValuesAverageGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUsersValuesAverageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
