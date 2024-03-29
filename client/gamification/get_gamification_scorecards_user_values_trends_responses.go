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

// GetGamificationScorecardsUserValuesTrendsReader is a Reader for the GetGamificationScorecardsUserValuesTrends structure.
type GetGamificationScorecardsUserValuesTrendsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationScorecardsUserValuesTrendsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationScorecardsUserValuesTrendsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationScorecardsUserValuesTrendsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationScorecardsUserValuesTrendsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationScorecardsUserValuesTrendsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationScorecardsUserValuesTrendsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationScorecardsUserValuesTrendsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationScorecardsUserValuesTrendsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationScorecardsUserValuesTrendsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationScorecardsUserValuesTrendsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationScorecardsUserValuesTrendsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationScorecardsUserValuesTrendsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationScorecardsUserValuesTrendsOK creates a GetGamificationScorecardsUserValuesTrendsOK with default headers values
func NewGetGamificationScorecardsUserValuesTrendsOK() *GetGamificationScorecardsUserValuesTrendsOK {
	return &GetGamificationScorecardsUserValuesTrendsOK{}
}

/*
GetGamificationScorecardsUserValuesTrendsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetGamificationScorecardsUserValuesTrendsOK struct {
	Payload *models.WorkdayValuesTrend
}

// IsSuccess returns true when this get gamification scorecards user values trends o k response has a 2xx status code
func (o *GetGamificationScorecardsUserValuesTrendsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gamification scorecards user values trends o k response has a 3xx status code
func (o *GetGamificationScorecardsUserValuesTrendsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user values trends o k response has a 4xx status code
func (o *GetGamificationScorecardsUserValuesTrendsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification scorecards user values trends o k response has a 5xx status code
func (o *GetGamificationScorecardsUserValuesTrendsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user values trends o k response a status code equal to that given
func (o *GetGamificationScorecardsUserValuesTrendsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGamificationScorecardsUserValuesTrendsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsOK  %+v", 200, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsOK  %+v", 200, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsOK) GetPayload() *models.WorkdayValuesTrend {
	return o.Payload
}

func (o *GetGamificationScorecardsUserValuesTrendsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkdayValuesTrend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserValuesTrendsBadRequest creates a GetGamificationScorecardsUserValuesTrendsBadRequest with default headers values
func NewGetGamificationScorecardsUserValuesTrendsBadRequest() *GetGamificationScorecardsUserValuesTrendsBadRequest {
	return &GetGamificationScorecardsUserValuesTrendsBadRequest{}
}

/*
GetGamificationScorecardsUserValuesTrendsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationScorecardsUserValuesTrendsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user values trends bad request response has a 2xx status code
func (o *GetGamificationScorecardsUserValuesTrendsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user values trends bad request response has a 3xx status code
func (o *GetGamificationScorecardsUserValuesTrendsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user values trends bad request response has a 4xx status code
func (o *GetGamificationScorecardsUserValuesTrendsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user values trends bad request response has a 5xx status code
func (o *GetGamificationScorecardsUserValuesTrendsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user values trends bad request response a status code equal to that given
func (o *GetGamificationScorecardsUserValuesTrendsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetGamificationScorecardsUserValuesTrendsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserValuesTrendsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserValuesTrendsUnauthorized creates a GetGamificationScorecardsUserValuesTrendsUnauthorized with default headers values
func NewGetGamificationScorecardsUserValuesTrendsUnauthorized() *GetGamificationScorecardsUserValuesTrendsUnauthorized {
	return &GetGamificationScorecardsUserValuesTrendsUnauthorized{}
}

/*
GetGamificationScorecardsUserValuesTrendsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationScorecardsUserValuesTrendsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user values trends unauthorized response has a 2xx status code
func (o *GetGamificationScorecardsUserValuesTrendsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user values trends unauthorized response has a 3xx status code
func (o *GetGamificationScorecardsUserValuesTrendsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user values trends unauthorized response has a 4xx status code
func (o *GetGamificationScorecardsUserValuesTrendsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user values trends unauthorized response has a 5xx status code
func (o *GetGamificationScorecardsUserValuesTrendsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user values trends unauthorized response a status code equal to that given
func (o *GetGamificationScorecardsUserValuesTrendsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGamificationScorecardsUserValuesTrendsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserValuesTrendsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserValuesTrendsForbidden creates a GetGamificationScorecardsUserValuesTrendsForbidden with default headers values
func NewGetGamificationScorecardsUserValuesTrendsForbidden() *GetGamificationScorecardsUserValuesTrendsForbidden {
	return &GetGamificationScorecardsUserValuesTrendsForbidden{}
}

/*
GetGamificationScorecardsUserValuesTrendsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationScorecardsUserValuesTrendsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user values trends forbidden response has a 2xx status code
func (o *GetGamificationScorecardsUserValuesTrendsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user values trends forbidden response has a 3xx status code
func (o *GetGamificationScorecardsUserValuesTrendsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user values trends forbidden response has a 4xx status code
func (o *GetGamificationScorecardsUserValuesTrendsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user values trends forbidden response has a 5xx status code
func (o *GetGamificationScorecardsUserValuesTrendsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user values trends forbidden response a status code equal to that given
func (o *GetGamificationScorecardsUserValuesTrendsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGamificationScorecardsUserValuesTrendsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserValuesTrendsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserValuesTrendsNotFound creates a GetGamificationScorecardsUserValuesTrendsNotFound with default headers values
func NewGetGamificationScorecardsUserValuesTrendsNotFound() *GetGamificationScorecardsUserValuesTrendsNotFound {
	return &GetGamificationScorecardsUserValuesTrendsNotFound{}
}

/*
GetGamificationScorecardsUserValuesTrendsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetGamificationScorecardsUserValuesTrendsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user values trends not found response has a 2xx status code
func (o *GetGamificationScorecardsUserValuesTrendsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user values trends not found response has a 3xx status code
func (o *GetGamificationScorecardsUserValuesTrendsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user values trends not found response has a 4xx status code
func (o *GetGamificationScorecardsUserValuesTrendsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user values trends not found response has a 5xx status code
func (o *GetGamificationScorecardsUserValuesTrendsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user values trends not found response a status code equal to that given
func (o *GetGamificationScorecardsUserValuesTrendsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetGamificationScorecardsUserValuesTrendsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserValuesTrendsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserValuesTrendsRequestTimeout creates a GetGamificationScorecardsUserValuesTrendsRequestTimeout with default headers values
func NewGetGamificationScorecardsUserValuesTrendsRequestTimeout() *GetGamificationScorecardsUserValuesTrendsRequestTimeout {
	return &GetGamificationScorecardsUserValuesTrendsRequestTimeout{}
}

/*
GetGamificationScorecardsUserValuesTrendsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationScorecardsUserValuesTrendsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user values trends request timeout response has a 2xx status code
func (o *GetGamificationScorecardsUserValuesTrendsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user values trends request timeout response has a 3xx status code
func (o *GetGamificationScorecardsUserValuesTrendsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user values trends request timeout response has a 4xx status code
func (o *GetGamificationScorecardsUserValuesTrendsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user values trends request timeout response has a 5xx status code
func (o *GetGamificationScorecardsUserValuesTrendsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user values trends request timeout response a status code equal to that given
func (o *GetGamificationScorecardsUserValuesTrendsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetGamificationScorecardsUserValuesTrendsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserValuesTrendsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge creates a GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge with default headers values
func NewGetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge() *GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge {
	return &GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge{}
}

/*
GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user values trends request entity too large response has a 2xx status code
func (o *GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user values trends request entity too large response has a 3xx status code
func (o *GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user values trends request entity too large response has a 4xx status code
func (o *GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user values trends request entity too large response has a 5xx status code
func (o *GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user values trends request entity too large response a status code equal to that given
func (o *GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserValuesTrendsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserValuesTrendsUnsupportedMediaType creates a GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType with default headers values
func NewGetGamificationScorecardsUserValuesTrendsUnsupportedMediaType() *GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType {
	return &GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType{}
}

/*
GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user values trends unsupported media type response has a 2xx status code
func (o *GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user values trends unsupported media type response has a 3xx status code
func (o *GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user values trends unsupported media type response has a 4xx status code
func (o *GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user values trends unsupported media type response has a 5xx status code
func (o *GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user values trends unsupported media type response a status code equal to that given
func (o *GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserValuesTrendsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserValuesTrendsTooManyRequests creates a GetGamificationScorecardsUserValuesTrendsTooManyRequests with default headers values
func NewGetGamificationScorecardsUserValuesTrendsTooManyRequests() *GetGamificationScorecardsUserValuesTrendsTooManyRequests {
	return &GetGamificationScorecardsUserValuesTrendsTooManyRequests{}
}

/*
GetGamificationScorecardsUserValuesTrendsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationScorecardsUserValuesTrendsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user values trends too many requests response has a 2xx status code
func (o *GetGamificationScorecardsUserValuesTrendsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user values trends too many requests response has a 3xx status code
func (o *GetGamificationScorecardsUserValuesTrendsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user values trends too many requests response has a 4xx status code
func (o *GetGamificationScorecardsUserValuesTrendsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification scorecards user values trends too many requests response has a 5xx status code
func (o *GetGamificationScorecardsUserValuesTrendsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification scorecards user values trends too many requests response a status code equal to that given
func (o *GetGamificationScorecardsUserValuesTrendsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetGamificationScorecardsUserValuesTrendsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserValuesTrendsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserValuesTrendsInternalServerError creates a GetGamificationScorecardsUserValuesTrendsInternalServerError with default headers values
func NewGetGamificationScorecardsUserValuesTrendsInternalServerError() *GetGamificationScorecardsUserValuesTrendsInternalServerError {
	return &GetGamificationScorecardsUserValuesTrendsInternalServerError{}
}

/*
GetGamificationScorecardsUserValuesTrendsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationScorecardsUserValuesTrendsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user values trends internal server error response has a 2xx status code
func (o *GetGamificationScorecardsUserValuesTrendsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user values trends internal server error response has a 3xx status code
func (o *GetGamificationScorecardsUserValuesTrendsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user values trends internal server error response has a 4xx status code
func (o *GetGamificationScorecardsUserValuesTrendsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification scorecards user values trends internal server error response has a 5xx status code
func (o *GetGamificationScorecardsUserValuesTrendsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification scorecards user values trends internal server error response a status code equal to that given
func (o *GetGamificationScorecardsUserValuesTrendsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGamificationScorecardsUserValuesTrendsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserValuesTrendsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserValuesTrendsServiceUnavailable creates a GetGamificationScorecardsUserValuesTrendsServiceUnavailable with default headers values
func NewGetGamificationScorecardsUserValuesTrendsServiceUnavailable() *GetGamificationScorecardsUserValuesTrendsServiceUnavailable {
	return &GetGamificationScorecardsUserValuesTrendsServiceUnavailable{}
}

/*
GetGamificationScorecardsUserValuesTrendsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationScorecardsUserValuesTrendsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user values trends service unavailable response has a 2xx status code
func (o *GetGamificationScorecardsUserValuesTrendsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user values trends service unavailable response has a 3xx status code
func (o *GetGamificationScorecardsUserValuesTrendsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user values trends service unavailable response has a 4xx status code
func (o *GetGamificationScorecardsUserValuesTrendsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification scorecards user values trends service unavailable response has a 5xx status code
func (o *GetGamificationScorecardsUserValuesTrendsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification scorecards user values trends service unavailable response a status code equal to that given
func (o *GetGamificationScorecardsUserValuesTrendsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetGamificationScorecardsUserValuesTrendsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserValuesTrendsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserValuesTrendsGatewayTimeout creates a GetGamificationScorecardsUserValuesTrendsGatewayTimeout with default headers values
func NewGetGamificationScorecardsUserValuesTrendsGatewayTimeout() *GetGamificationScorecardsUserValuesTrendsGatewayTimeout {
	return &GetGamificationScorecardsUserValuesTrendsGatewayTimeout{}
}

/*
GetGamificationScorecardsUserValuesTrendsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetGamificationScorecardsUserValuesTrendsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification scorecards user values trends gateway timeout response has a 2xx status code
func (o *GetGamificationScorecardsUserValuesTrendsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification scorecards user values trends gateway timeout response has a 3xx status code
func (o *GetGamificationScorecardsUserValuesTrendsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification scorecards user values trends gateway timeout response has a 4xx status code
func (o *GetGamificationScorecardsUserValuesTrendsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification scorecards user values trends gateway timeout response has a 5xx status code
func (o *GetGamificationScorecardsUserValuesTrendsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification scorecards user values trends gateway timeout response a status code equal to that given
func (o *GetGamificationScorecardsUserValuesTrendsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetGamificationScorecardsUserValuesTrendsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/values/trends][%d] getGamificationScorecardsUserValuesTrendsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationScorecardsUserValuesTrendsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserValuesTrendsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
