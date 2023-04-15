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

// GetGamificationScorecardsUserPointsTrendsReader is a Reader for the GetGamificationScorecardsUserPointsTrends structure.
type GetGamificationScorecardsUserPointsTrendsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationScorecardsUserPointsTrendsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationScorecardsUserPointsTrendsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationScorecardsUserPointsTrendsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationScorecardsUserPointsTrendsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationScorecardsUserPointsTrendsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationScorecardsUserPointsTrendsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationScorecardsUserPointsTrendsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationScorecardsUserPointsTrendsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationScorecardsUserPointsTrendsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationScorecardsUserPointsTrendsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationScorecardsUserPointsTrendsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationScorecardsUserPointsTrendsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationScorecardsUserPointsTrendsOK creates a GetGamificationScorecardsUserPointsTrendsOK with default headers values
func NewGetGamificationScorecardsUserPointsTrendsOK() *GetGamificationScorecardsUserPointsTrendsOK {
	return &GetGamificationScorecardsUserPointsTrendsOK{}
}

/*
GetGamificationScorecardsUserPointsTrendsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetGamificationScorecardsUserPointsTrendsOK struct {
	Payload *models.WorkdayPointsTrend
}

// IsSuccess returns true when this get gamification scorecards user points trends o k response has a 2xx status code
func (o *GetGamificationScorecardsUserPointsTrendsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gamification scorecards user points trends o k response has a 3xx status code
func (o *GetGamificationScorecardsUserPointsTrendsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user points trends o k response has a 4xx status code
func (o *GetGamificationScorecardsUserPointsTrendsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification scorecards user points trends o k response has a 5xx status code
func (o *GetGamificationScorecardsUserPointsTrendsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user points trends o k response a status code equal to that given
func (o *GetGamificationScorecardsUserPointsTrendsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGamificationScorecardsUserPointsTrendsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsOK  %+v", 200, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsOK  %+v", 200, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsOK) GetPayload() *models.WorkdayPointsTrend {
	return o.Payload
}

func (o *GetGamificationScorecardsUserPointsTrendsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkdayPointsTrend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserPointsTrendsBadRequest creates a GetGamificationScorecardsUserPointsTrendsBadRequest with default headers values
func NewGetGamificationScorecardsUserPointsTrendsBadRequest() *GetGamificationScorecardsUserPointsTrendsBadRequest {
	return &GetGamificationScorecardsUserPointsTrendsBadRequest{}
}

/*
GetGamificationScorecardsUserPointsTrendsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationScorecardsUserPointsTrendsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user points trends bad request response has a 2xx status code
func (o *GetGamificationScorecardsUserPointsTrendsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user points trends bad request response has a 3xx status code
func (o *GetGamificationScorecardsUserPointsTrendsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user points trends bad request response has a 4xx status code
func (o *GetGamificationScorecardsUserPointsTrendsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user points trends bad request response has a 5xx status code
func (o *GetGamificationScorecardsUserPointsTrendsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user points trends bad request response a status code equal to that given
func (o *GetGamificationScorecardsUserPointsTrendsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetGamificationScorecardsUserPointsTrendsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserPointsTrendsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserPointsTrendsUnauthorized creates a GetGamificationScorecardsUserPointsTrendsUnauthorized with default headers values
func NewGetGamificationScorecardsUserPointsTrendsUnauthorized() *GetGamificationScorecardsUserPointsTrendsUnauthorized {
	return &GetGamificationScorecardsUserPointsTrendsUnauthorized{}
}

/*
GetGamificationScorecardsUserPointsTrendsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationScorecardsUserPointsTrendsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user points trends unauthorized response has a 2xx status code
func (o *GetGamificationScorecardsUserPointsTrendsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user points trends unauthorized response has a 3xx status code
func (o *GetGamificationScorecardsUserPointsTrendsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user points trends unauthorized response has a 4xx status code
func (o *GetGamificationScorecardsUserPointsTrendsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user points trends unauthorized response has a 5xx status code
func (o *GetGamificationScorecardsUserPointsTrendsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user points trends unauthorized response a status code equal to that given
func (o *GetGamificationScorecardsUserPointsTrendsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGamificationScorecardsUserPointsTrendsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserPointsTrendsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserPointsTrendsForbidden creates a GetGamificationScorecardsUserPointsTrendsForbidden with default headers values
func NewGetGamificationScorecardsUserPointsTrendsForbidden() *GetGamificationScorecardsUserPointsTrendsForbidden {
	return &GetGamificationScorecardsUserPointsTrendsForbidden{}
}

/*
GetGamificationScorecardsUserPointsTrendsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationScorecardsUserPointsTrendsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user points trends forbidden response has a 2xx status code
func (o *GetGamificationScorecardsUserPointsTrendsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user points trends forbidden response has a 3xx status code
func (o *GetGamificationScorecardsUserPointsTrendsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user points trends forbidden response has a 4xx status code
func (o *GetGamificationScorecardsUserPointsTrendsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user points trends forbidden response has a 5xx status code
func (o *GetGamificationScorecardsUserPointsTrendsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user points trends forbidden response a status code equal to that given
func (o *GetGamificationScorecardsUserPointsTrendsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGamificationScorecardsUserPointsTrendsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserPointsTrendsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserPointsTrendsNotFound creates a GetGamificationScorecardsUserPointsTrendsNotFound with default headers values
func NewGetGamificationScorecardsUserPointsTrendsNotFound() *GetGamificationScorecardsUserPointsTrendsNotFound {
	return &GetGamificationScorecardsUserPointsTrendsNotFound{}
}

/*
GetGamificationScorecardsUserPointsTrendsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetGamificationScorecardsUserPointsTrendsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user points trends not found response has a 2xx status code
func (o *GetGamificationScorecardsUserPointsTrendsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user points trends not found response has a 3xx status code
func (o *GetGamificationScorecardsUserPointsTrendsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user points trends not found response has a 4xx status code
func (o *GetGamificationScorecardsUserPointsTrendsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user points trends not found response has a 5xx status code
func (o *GetGamificationScorecardsUserPointsTrendsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user points trends not found response a status code equal to that given
func (o *GetGamificationScorecardsUserPointsTrendsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetGamificationScorecardsUserPointsTrendsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserPointsTrendsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserPointsTrendsRequestTimeout creates a GetGamificationScorecardsUserPointsTrendsRequestTimeout with default headers values
func NewGetGamificationScorecardsUserPointsTrendsRequestTimeout() *GetGamificationScorecardsUserPointsTrendsRequestTimeout {
	return &GetGamificationScorecardsUserPointsTrendsRequestTimeout{}
}

/*
GetGamificationScorecardsUserPointsTrendsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationScorecardsUserPointsTrendsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user points trends request timeout response has a 2xx status code
func (o *GetGamificationScorecardsUserPointsTrendsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user points trends request timeout response has a 3xx status code
func (o *GetGamificationScorecardsUserPointsTrendsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user points trends request timeout response has a 4xx status code
func (o *GetGamificationScorecardsUserPointsTrendsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user points trends request timeout response has a 5xx status code
func (o *GetGamificationScorecardsUserPointsTrendsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user points trends request timeout response a status code equal to that given
func (o *GetGamificationScorecardsUserPointsTrendsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetGamificationScorecardsUserPointsTrendsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserPointsTrendsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge creates a GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge with default headers values
func NewGetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge() *GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge {
	return &GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge{}
}

/*
GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user points trends request entity too large response has a 2xx status code
func (o *GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user points trends request entity too large response has a 3xx status code
func (o *GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user points trends request entity too large response has a 4xx status code
func (o *GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user points trends request entity too large response has a 5xx status code
func (o *GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user points trends request entity too large response a status code equal to that given
func (o *GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserPointsTrendsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserPointsTrendsUnsupportedMediaType creates a GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType with default headers values
func NewGetGamificationScorecardsUserPointsTrendsUnsupportedMediaType() *GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType {
	return &GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType{}
}

/*
GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user points trends unsupported media type response has a 2xx status code
func (o *GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user points trends unsupported media type response has a 3xx status code
func (o *GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user points trends unsupported media type response has a 4xx status code
func (o *GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user points trends unsupported media type response has a 5xx status code
func (o *GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user points trends unsupported media type response a status code equal to that given
func (o *GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserPointsTrendsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserPointsTrendsTooManyRequests creates a GetGamificationScorecardsUserPointsTrendsTooManyRequests with default headers values
func NewGetGamificationScorecardsUserPointsTrendsTooManyRequests() *GetGamificationScorecardsUserPointsTrendsTooManyRequests {
	return &GetGamificationScorecardsUserPointsTrendsTooManyRequests{}
}

/*
GetGamificationScorecardsUserPointsTrendsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationScorecardsUserPointsTrendsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user points trends too many requests response has a 2xx status code
func (o *GetGamificationScorecardsUserPointsTrendsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user points trends too many requests response has a 3xx status code
func (o *GetGamificationScorecardsUserPointsTrendsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user points trends too many requests response has a 4xx status code
func (o *GetGamificationScorecardsUserPointsTrendsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user points trends too many requests response has a 5xx status code
func (o *GetGamificationScorecardsUserPointsTrendsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user points trends too many requests response a status code equal to that given
func (o *GetGamificationScorecardsUserPointsTrendsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetGamificationScorecardsUserPointsTrendsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserPointsTrendsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserPointsTrendsInternalServerError creates a GetGamificationScorecardsUserPointsTrendsInternalServerError with default headers values
func NewGetGamificationScorecardsUserPointsTrendsInternalServerError() *GetGamificationScorecardsUserPointsTrendsInternalServerError {
	return &GetGamificationScorecardsUserPointsTrendsInternalServerError{}
}

/*
GetGamificationScorecardsUserPointsTrendsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationScorecardsUserPointsTrendsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user points trends internal server error response has a 2xx status code
func (o *GetGamificationScorecardsUserPointsTrendsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user points trends internal server error response has a 3xx status code
func (o *GetGamificationScorecardsUserPointsTrendsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user points trends internal server error response has a 4xx status code
func (o *GetGamificationScorecardsUserPointsTrendsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification scorecards user points trends internal server error response has a 5xx status code
func (o *GetGamificationScorecardsUserPointsTrendsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification scorecards user points trends internal server error response a status code equal to that given
func (o *GetGamificationScorecardsUserPointsTrendsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGamificationScorecardsUserPointsTrendsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserPointsTrendsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserPointsTrendsServiceUnavailable creates a GetGamificationScorecardsUserPointsTrendsServiceUnavailable with default headers values
func NewGetGamificationScorecardsUserPointsTrendsServiceUnavailable() *GetGamificationScorecardsUserPointsTrendsServiceUnavailable {
	return &GetGamificationScorecardsUserPointsTrendsServiceUnavailable{}
}

/*
GetGamificationScorecardsUserPointsTrendsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationScorecardsUserPointsTrendsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user points trends service unavailable response has a 2xx status code
func (o *GetGamificationScorecardsUserPointsTrendsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user points trends service unavailable response has a 3xx status code
func (o *GetGamificationScorecardsUserPointsTrendsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user points trends service unavailable response has a 4xx status code
func (o *GetGamificationScorecardsUserPointsTrendsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification scorecards user points trends service unavailable response has a 5xx status code
func (o *GetGamificationScorecardsUserPointsTrendsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification scorecards user points trends service unavailable response a status code equal to that given
func (o *GetGamificationScorecardsUserPointsTrendsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetGamificationScorecardsUserPointsTrendsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserPointsTrendsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserPointsTrendsGatewayTimeout creates a GetGamificationScorecardsUserPointsTrendsGatewayTimeout with default headers values
func NewGetGamificationScorecardsUserPointsTrendsGatewayTimeout() *GetGamificationScorecardsUserPointsTrendsGatewayTimeout {
	return &GetGamificationScorecardsUserPointsTrendsGatewayTimeout{}
}

/*
GetGamificationScorecardsUserPointsTrendsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetGamificationScorecardsUserPointsTrendsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user points trends gateway timeout response has a 2xx status code
func (o *GetGamificationScorecardsUserPointsTrendsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user points trends gateway timeout response has a 3xx status code
func (o *GetGamificationScorecardsUserPointsTrendsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user points trends gateway timeout response has a 4xx status code
func (o *GetGamificationScorecardsUserPointsTrendsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification scorecards user points trends gateway timeout response has a 5xx status code
func (o *GetGamificationScorecardsUserPointsTrendsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification scorecards user points trends gateway timeout response a status code equal to that given
func (o *GetGamificationScorecardsUserPointsTrendsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetGamificationScorecardsUserPointsTrendsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/points/trends][%d] getGamificationScorecardsUserPointsTrendsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationScorecardsUserPointsTrendsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserPointsTrendsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
