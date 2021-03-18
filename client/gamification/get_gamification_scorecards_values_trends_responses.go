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

// GetGamificationScorecardsValuesTrendsReader is a Reader for the GetGamificationScorecardsValuesTrends structure.
type GetGamificationScorecardsValuesTrendsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationScorecardsValuesTrendsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationScorecardsValuesTrendsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationScorecardsValuesTrendsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationScorecardsValuesTrendsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationScorecardsValuesTrendsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationScorecardsValuesTrendsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationScorecardsValuesTrendsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationScorecardsValuesTrendsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationScorecardsValuesTrendsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationScorecardsValuesTrendsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationScorecardsValuesTrendsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationScorecardsValuesTrendsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationScorecardsValuesTrendsOK creates a GetGamificationScorecardsValuesTrendsOK with default headers values
func NewGetGamificationScorecardsValuesTrendsOK() *GetGamificationScorecardsValuesTrendsOK {
	return &GetGamificationScorecardsValuesTrendsOK{}
}

/*GetGamificationScorecardsValuesTrendsOK handles this case with default header values.

successful operation
*/
type GetGamificationScorecardsValuesTrendsOK struct {
	Payload *models.WorkdayValuesTrend
}

func (o *GetGamificationScorecardsValuesTrendsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/values/trends][%d] getGamificationScorecardsValuesTrendsOK  %+v", 200, o.Payload)
}

func (o *GetGamificationScorecardsValuesTrendsOK) GetPayload() *models.WorkdayValuesTrend {
	return o.Payload
}

func (o *GetGamificationScorecardsValuesTrendsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkdayValuesTrend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsValuesTrendsBadRequest creates a GetGamificationScorecardsValuesTrendsBadRequest with default headers values
func NewGetGamificationScorecardsValuesTrendsBadRequest() *GetGamificationScorecardsValuesTrendsBadRequest {
	return &GetGamificationScorecardsValuesTrendsBadRequest{}
}

/*GetGamificationScorecardsValuesTrendsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationScorecardsValuesTrendsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsValuesTrendsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/values/trends][%d] getGamificationScorecardsValuesTrendsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationScorecardsValuesTrendsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsValuesTrendsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsValuesTrendsUnauthorized creates a GetGamificationScorecardsValuesTrendsUnauthorized with default headers values
func NewGetGamificationScorecardsValuesTrendsUnauthorized() *GetGamificationScorecardsValuesTrendsUnauthorized {
	return &GetGamificationScorecardsValuesTrendsUnauthorized{}
}

/*GetGamificationScorecardsValuesTrendsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationScorecardsValuesTrendsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsValuesTrendsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/values/trends][%d] getGamificationScorecardsValuesTrendsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationScorecardsValuesTrendsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsValuesTrendsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsValuesTrendsForbidden creates a GetGamificationScorecardsValuesTrendsForbidden with default headers values
func NewGetGamificationScorecardsValuesTrendsForbidden() *GetGamificationScorecardsValuesTrendsForbidden {
	return &GetGamificationScorecardsValuesTrendsForbidden{}
}

/*GetGamificationScorecardsValuesTrendsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationScorecardsValuesTrendsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsValuesTrendsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/values/trends][%d] getGamificationScorecardsValuesTrendsForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationScorecardsValuesTrendsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsValuesTrendsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsValuesTrendsNotFound creates a GetGamificationScorecardsValuesTrendsNotFound with default headers values
func NewGetGamificationScorecardsValuesTrendsNotFound() *GetGamificationScorecardsValuesTrendsNotFound {
	return &GetGamificationScorecardsValuesTrendsNotFound{}
}

/*GetGamificationScorecardsValuesTrendsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationScorecardsValuesTrendsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsValuesTrendsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/values/trends][%d] getGamificationScorecardsValuesTrendsNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationScorecardsValuesTrendsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsValuesTrendsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsValuesTrendsRequestEntityTooLarge creates a GetGamificationScorecardsValuesTrendsRequestEntityTooLarge with default headers values
func NewGetGamificationScorecardsValuesTrendsRequestEntityTooLarge() *GetGamificationScorecardsValuesTrendsRequestEntityTooLarge {
	return &GetGamificationScorecardsValuesTrendsRequestEntityTooLarge{}
}

/*GetGamificationScorecardsValuesTrendsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGamificationScorecardsValuesTrendsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsValuesTrendsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/values/trends][%d] getGamificationScorecardsValuesTrendsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationScorecardsValuesTrendsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsValuesTrendsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsValuesTrendsUnsupportedMediaType creates a GetGamificationScorecardsValuesTrendsUnsupportedMediaType with default headers values
func NewGetGamificationScorecardsValuesTrendsUnsupportedMediaType() *GetGamificationScorecardsValuesTrendsUnsupportedMediaType {
	return &GetGamificationScorecardsValuesTrendsUnsupportedMediaType{}
}

/*GetGamificationScorecardsValuesTrendsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationScorecardsValuesTrendsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsValuesTrendsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/values/trends][%d] getGamificationScorecardsValuesTrendsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationScorecardsValuesTrendsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsValuesTrendsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsValuesTrendsTooManyRequests creates a GetGamificationScorecardsValuesTrendsTooManyRequests with default headers values
func NewGetGamificationScorecardsValuesTrendsTooManyRequests() *GetGamificationScorecardsValuesTrendsTooManyRequests {
	return &GetGamificationScorecardsValuesTrendsTooManyRequests{}
}

/*GetGamificationScorecardsValuesTrendsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetGamificationScorecardsValuesTrendsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsValuesTrendsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/values/trends][%d] getGamificationScorecardsValuesTrendsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationScorecardsValuesTrendsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsValuesTrendsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsValuesTrendsInternalServerError creates a GetGamificationScorecardsValuesTrendsInternalServerError with default headers values
func NewGetGamificationScorecardsValuesTrendsInternalServerError() *GetGamificationScorecardsValuesTrendsInternalServerError {
	return &GetGamificationScorecardsValuesTrendsInternalServerError{}
}

/*GetGamificationScorecardsValuesTrendsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationScorecardsValuesTrendsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsValuesTrendsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/values/trends][%d] getGamificationScorecardsValuesTrendsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationScorecardsValuesTrendsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsValuesTrendsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsValuesTrendsServiceUnavailable creates a GetGamificationScorecardsValuesTrendsServiceUnavailable with default headers values
func NewGetGamificationScorecardsValuesTrendsServiceUnavailable() *GetGamificationScorecardsValuesTrendsServiceUnavailable {
	return &GetGamificationScorecardsValuesTrendsServiceUnavailable{}
}

/*GetGamificationScorecardsValuesTrendsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationScorecardsValuesTrendsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsValuesTrendsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/values/trends][%d] getGamificationScorecardsValuesTrendsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationScorecardsValuesTrendsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsValuesTrendsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsValuesTrendsGatewayTimeout creates a GetGamificationScorecardsValuesTrendsGatewayTimeout with default headers values
func NewGetGamificationScorecardsValuesTrendsGatewayTimeout() *GetGamificationScorecardsValuesTrendsGatewayTimeout {
	return &GetGamificationScorecardsValuesTrendsGatewayTimeout{}
}

/*GetGamificationScorecardsValuesTrendsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationScorecardsValuesTrendsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsValuesTrendsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/values/trends][%d] getGamificationScorecardsValuesTrendsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationScorecardsValuesTrendsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsValuesTrendsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
