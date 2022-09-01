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

// GetGamificationScorecardsProfileMetricUsersValuesTrendsReader is a Reader for the GetGamificationScorecardsProfileMetricUsersValuesTrends structure.
type GetGamificationScorecardsProfileMetricUsersValuesTrendsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationScorecardsProfileMetricUsersValuesTrendsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationScorecardsProfileMetricUsersValuesTrendsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationScorecardsProfileMetricUsersValuesTrendsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationScorecardsProfileMetricUsersValuesTrendsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationScorecardsProfileMetricUsersValuesTrendsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationScorecardsProfileMetricUsersValuesTrendsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationScorecardsProfileMetricUsersValuesTrendsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationScorecardsProfileMetricUsersValuesTrendsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationScorecardsProfileMetricUsersValuesTrendsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationScorecardsProfileMetricUsersValuesTrendsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationScorecardsProfileMetricUsersValuesTrendsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationScorecardsProfileMetricUsersValuesTrendsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationScorecardsProfileMetricUsersValuesTrendsOK creates a GetGamificationScorecardsProfileMetricUsersValuesTrendsOK with default headers values
func NewGetGamificationScorecardsProfileMetricUsersValuesTrendsOK() *GetGamificationScorecardsProfileMetricUsersValuesTrendsOK {
	return &GetGamificationScorecardsProfileMetricUsersValuesTrendsOK{}
}

/*GetGamificationScorecardsProfileMetricUsersValuesTrendsOK handles this case with default header values.

successful operation
*/
type GetGamificationScorecardsProfileMetricUsersValuesTrendsOK struct {
	Payload *models.MetricValueTrendAverage
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends][%d] getGamificationScorecardsProfileMetricUsersValuesTrendsOK  %+v", 200, o.Payload)
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsOK) GetPayload() *models.MetricValueTrendAverage {
	return o.Payload
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetricValueTrendAverage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsProfileMetricUsersValuesTrendsBadRequest creates a GetGamificationScorecardsProfileMetricUsersValuesTrendsBadRequest with default headers values
func NewGetGamificationScorecardsProfileMetricUsersValuesTrendsBadRequest() *GetGamificationScorecardsProfileMetricUsersValuesTrendsBadRequest {
	return &GetGamificationScorecardsProfileMetricUsersValuesTrendsBadRequest{}
}

/*GetGamificationScorecardsProfileMetricUsersValuesTrendsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationScorecardsProfileMetricUsersValuesTrendsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends][%d] getGamificationScorecardsProfileMetricUsersValuesTrendsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsProfileMetricUsersValuesTrendsUnauthorized creates a GetGamificationScorecardsProfileMetricUsersValuesTrendsUnauthorized with default headers values
func NewGetGamificationScorecardsProfileMetricUsersValuesTrendsUnauthorized() *GetGamificationScorecardsProfileMetricUsersValuesTrendsUnauthorized {
	return &GetGamificationScorecardsProfileMetricUsersValuesTrendsUnauthorized{}
}

/*GetGamificationScorecardsProfileMetricUsersValuesTrendsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationScorecardsProfileMetricUsersValuesTrendsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends][%d] getGamificationScorecardsProfileMetricUsersValuesTrendsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsProfileMetricUsersValuesTrendsForbidden creates a GetGamificationScorecardsProfileMetricUsersValuesTrendsForbidden with default headers values
func NewGetGamificationScorecardsProfileMetricUsersValuesTrendsForbidden() *GetGamificationScorecardsProfileMetricUsersValuesTrendsForbidden {
	return &GetGamificationScorecardsProfileMetricUsersValuesTrendsForbidden{}
}

/*GetGamificationScorecardsProfileMetricUsersValuesTrendsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationScorecardsProfileMetricUsersValuesTrendsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends][%d] getGamificationScorecardsProfileMetricUsersValuesTrendsForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsProfileMetricUsersValuesTrendsNotFound creates a GetGamificationScorecardsProfileMetricUsersValuesTrendsNotFound with default headers values
func NewGetGamificationScorecardsProfileMetricUsersValuesTrendsNotFound() *GetGamificationScorecardsProfileMetricUsersValuesTrendsNotFound {
	return &GetGamificationScorecardsProfileMetricUsersValuesTrendsNotFound{}
}

/*GetGamificationScorecardsProfileMetricUsersValuesTrendsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationScorecardsProfileMetricUsersValuesTrendsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends][%d] getGamificationScorecardsProfileMetricUsersValuesTrendsNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsProfileMetricUsersValuesTrendsRequestTimeout creates a GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestTimeout with default headers values
func NewGetGamificationScorecardsProfileMetricUsersValuesTrendsRequestTimeout() *GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestTimeout {
	return &GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestTimeout{}
}

/*GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends][%d] getGamificationScorecardsProfileMetricUsersValuesTrendsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsProfileMetricUsersValuesTrendsRequestEntityTooLarge creates a GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestEntityTooLarge with default headers values
func NewGetGamificationScorecardsProfileMetricUsersValuesTrendsRequestEntityTooLarge() *GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestEntityTooLarge {
	return &GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestEntityTooLarge{}
}

/*GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends][%d] getGamificationScorecardsProfileMetricUsersValuesTrendsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsProfileMetricUsersValuesTrendsUnsupportedMediaType creates a GetGamificationScorecardsProfileMetricUsersValuesTrendsUnsupportedMediaType with default headers values
func NewGetGamificationScorecardsProfileMetricUsersValuesTrendsUnsupportedMediaType() *GetGamificationScorecardsProfileMetricUsersValuesTrendsUnsupportedMediaType {
	return &GetGamificationScorecardsProfileMetricUsersValuesTrendsUnsupportedMediaType{}
}

/*GetGamificationScorecardsProfileMetricUsersValuesTrendsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationScorecardsProfileMetricUsersValuesTrendsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends][%d] getGamificationScorecardsProfileMetricUsersValuesTrendsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsProfileMetricUsersValuesTrendsTooManyRequests creates a GetGamificationScorecardsProfileMetricUsersValuesTrendsTooManyRequests with default headers values
func NewGetGamificationScorecardsProfileMetricUsersValuesTrendsTooManyRequests() *GetGamificationScorecardsProfileMetricUsersValuesTrendsTooManyRequests {
	return &GetGamificationScorecardsProfileMetricUsersValuesTrendsTooManyRequests{}
}

/*GetGamificationScorecardsProfileMetricUsersValuesTrendsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationScorecardsProfileMetricUsersValuesTrendsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends][%d] getGamificationScorecardsProfileMetricUsersValuesTrendsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsProfileMetricUsersValuesTrendsInternalServerError creates a GetGamificationScorecardsProfileMetricUsersValuesTrendsInternalServerError with default headers values
func NewGetGamificationScorecardsProfileMetricUsersValuesTrendsInternalServerError() *GetGamificationScorecardsProfileMetricUsersValuesTrendsInternalServerError {
	return &GetGamificationScorecardsProfileMetricUsersValuesTrendsInternalServerError{}
}

/*GetGamificationScorecardsProfileMetricUsersValuesTrendsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationScorecardsProfileMetricUsersValuesTrendsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends][%d] getGamificationScorecardsProfileMetricUsersValuesTrendsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsProfileMetricUsersValuesTrendsServiceUnavailable creates a GetGamificationScorecardsProfileMetricUsersValuesTrendsServiceUnavailable with default headers values
func NewGetGamificationScorecardsProfileMetricUsersValuesTrendsServiceUnavailable() *GetGamificationScorecardsProfileMetricUsersValuesTrendsServiceUnavailable {
	return &GetGamificationScorecardsProfileMetricUsersValuesTrendsServiceUnavailable{}
}

/*GetGamificationScorecardsProfileMetricUsersValuesTrendsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationScorecardsProfileMetricUsersValuesTrendsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends][%d] getGamificationScorecardsProfileMetricUsersValuesTrendsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsProfileMetricUsersValuesTrendsGatewayTimeout creates a GetGamificationScorecardsProfileMetricUsersValuesTrendsGatewayTimeout with default headers values
func NewGetGamificationScorecardsProfileMetricUsersValuesTrendsGatewayTimeout() *GetGamificationScorecardsProfileMetricUsersValuesTrendsGatewayTimeout {
	return &GetGamificationScorecardsProfileMetricUsersValuesTrendsGatewayTimeout{}
}

/*GetGamificationScorecardsProfileMetricUsersValuesTrendsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationScorecardsProfileMetricUsersValuesTrendsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends][%d] getGamificationScorecardsProfileMetricUsersValuesTrendsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsProfileMetricUsersValuesTrendsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}