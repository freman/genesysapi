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

// GetGamificationScorecardsPointsTrendsReader is a Reader for the GetGamificationScorecardsPointsTrends structure.
type GetGamificationScorecardsPointsTrendsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationScorecardsPointsTrendsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationScorecardsPointsTrendsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationScorecardsPointsTrendsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationScorecardsPointsTrendsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationScorecardsPointsTrendsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationScorecardsPointsTrendsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationScorecardsPointsTrendsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationScorecardsPointsTrendsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationScorecardsPointsTrendsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationScorecardsPointsTrendsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationScorecardsPointsTrendsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationScorecardsPointsTrendsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationScorecardsPointsTrendsOK creates a GetGamificationScorecardsPointsTrendsOK with default headers values
func NewGetGamificationScorecardsPointsTrendsOK() *GetGamificationScorecardsPointsTrendsOK {
	return &GetGamificationScorecardsPointsTrendsOK{}
}

/*GetGamificationScorecardsPointsTrendsOK handles this case with default header values.

successful operation
*/
type GetGamificationScorecardsPointsTrendsOK struct {
	Payload *models.WorkdayPointsTrend
}

func (o *GetGamificationScorecardsPointsTrendsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/points/trends][%d] getGamificationScorecardsPointsTrendsOK  %+v", 200, o.Payload)
}

func (o *GetGamificationScorecardsPointsTrendsOK) GetPayload() *models.WorkdayPointsTrend {
	return o.Payload
}

func (o *GetGamificationScorecardsPointsTrendsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkdayPointsTrend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsPointsTrendsBadRequest creates a GetGamificationScorecardsPointsTrendsBadRequest with default headers values
func NewGetGamificationScorecardsPointsTrendsBadRequest() *GetGamificationScorecardsPointsTrendsBadRequest {
	return &GetGamificationScorecardsPointsTrendsBadRequest{}
}

/*GetGamificationScorecardsPointsTrendsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationScorecardsPointsTrendsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsPointsTrendsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/points/trends][%d] getGamificationScorecardsPointsTrendsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationScorecardsPointsTrendsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsPointsTrendsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsPointsTrendsUnauthorized creates a GetGamificationScorecardsPointsTrendsUnauthorized with default headers values
func NewGetGamificationScorecardsPointsTrendsUnauthorized() *GetGamificationScorecardsPointsTrendsUnauthorized {
	return &GetGamificationScorecardsPointsTrendsUnauthorized{}
}

/*GetGamificationScorecardsPointsTrendsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationScorecardsPointsTrendsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsPointsTrendsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/points/trends][%d] getGamificationScorecardsPointsTrendsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationScorecardsPointsTrendsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsPointsTrendsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsPointsTrendsForbidden creates a GetGamificationScorecardsPointsTrendsForbidden with default headers values
func NewGetGamificationScorecardsPointsTrendsForbidden() *GetGamificationScorecardsPointsTrendsForbidden {
	return &GetGamificationScorecardsPointsTrendsForbidden{}
}

/*GetGamificationScorecardsPointsTrendsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationScorecardsPointsTrendsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsPointsTrendsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/points/trends][%d] getGamificationScorecardsPointsTrendsForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationScorecardsPointsTrendsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsPointsTrendsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsPointsTrendsNotFound creates a GetGamificationScorecardsPointsTrendsNotFound with default headers values
func NewGetGamificationScorecardsPointsTrendsNotFound() *GetGamificationScorecardsPointsTrendsNotFound {
	return &GetGamificationScorecardsPointsTrendsNotFound{}
}

/*GetGamificationScorecardsPointsTrendsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationScorecardsPointsTrendsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsPointsTrendsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/points/trends][%d] getGamificationScorecardsPointsTrendsNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationScorecardsPointsTrendsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsPointsTrendsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsPointsTrendsRequestEntityTooLarge creates a GetGamificationScorecardsPointsTrendsRequestEntityTooLarge with default headers values
func NewGetGamificationScorecardsPointsTrendsRequestEntityTooLarge() *GetGamificationScorecardsPointsTrendsRequestEntityTooLarge {
	return &GetGamificationScorecardsPointsTrendsRequestEntityTooLarge{}
}

/*GetGamificationScorecardsPointsTrendsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGamificationScorecardsPointsTrendsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsPointsTrendsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/points/trends][%d] getGamificationScorecardsPointsTrendsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationScorecardsPointsTrendsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsPointsTrendsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsPointsTrendsUnsupportedMediaType creates a GetGamificationScorecardsPointsTrendsUnsupportedMediaType with default headers values
func NewGetGamificationScorecardsPointsTrendsUnsupportedMediaType() *GetGamificationScorecardsPointsTrendsUnsupportedMediaType {
	return &GetGamificationScorecardsPointsTrendsUnsupportedMediaType{}
}

/*GetGamificationScorecardsPointsTrendsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationScorecardsPointsTrendsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsPointsTrendsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/points/trends][%d] getGamificationScorecardsPointsTrendsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationScorecardsPointsTrendsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsPointsTrendsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsPointsTrendsTooManyRequests creates a GetGamificationScorecardsPointsTrendsTooManyRequests with default headers values
func NewGetGamificationScorecardsPointsTrendsTooManyRequests() *GetGamificationScorecardsPointsTrendsTooManyRequests {
	return &GetGamificationScorecardsPointsTrendsTooManyRequests{}
}

/*GetGamificationScorecardsPointsTrendsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetGamificationScorecardsPointsTrendsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsPointsTrendsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/points/trends][%d] getGamificationScorecardsPointsTrendsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationScorecardsPointsTrendsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsPointsTrendsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsPointsTrendsInternalServerError creates a GetGamificationScorecardsPointsTrendsInternalServerError with default headers values
func NewGetGamificationScorecardsPointsTrendsInternalServerError() *GetGamificationScorecardsPointsTrendsInternalServerError {
	return &GetGamificationScorecardsPointsTrendsInternalServerError{}
}

/*GetGamificationScorecardsPointsTrendsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationScorecardsPointsTrendsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsPointsTrendsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/points/trends][%d] getGamificationScorecardsPointsTrendsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationScorecardsPointsTrendsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsPointsTrendsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsPointsTrendsServiceUnavailable creates a GetGamificationScorecardsPointsTrendsServiceUnavailable with default headers values
func NewGetGamificationScorecardsPointsTrendsServiceUnavailable() *GetGamificationScorecardsPointsTrendsServiceUnavailable {
	return &GetGamificationScorecardsPointsTrendsServiceUnavailable{}
}

/*GetGamificationScorecardsPointsTrendsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationScorecardsPointsTrendsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsPointsTrendsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/points/trends][%d] getGamificationScorecardsPointsTrendsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationScorecardsPointsTrendsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsPointsTrendsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsPointsTrendsGatewayTimeout creates a GetGamificationScorecardsPointsTrendsGatewayTimeout with default headers values
func NewGetGamificationScorecardsPointsTrendsGatewayTimeout() *GetGamificationScorecardsPointsTrendsGatewayTimeout {
	return &GetGamificationScorecardsPointsTrendsGatewayTimeout{}
}

/*GetGamificationScorecardsPointsTrendsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationScorecardsPointsTrendsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsPointsTrendsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/points/trends][%d] getGamificationScorecardsPointsTrendsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationScorecardsPointsTrendsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsPointsTrendsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
