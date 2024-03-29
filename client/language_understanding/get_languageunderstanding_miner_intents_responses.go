// Code generated by go-swagger; DO NOT EDIT.

package language_understanding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLanguageunderstandingMinerIntentsReader is a Reader for the GetLanguageunderstandingMinerIntents structure.
type GetLanguageunderstandingMinerIntentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguageunderstandingMinerIntentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguageunderstandingMinerIntentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguageunderstandingMinerIntentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguageunderstandingMinerIntentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguageunderstandingMinerIntentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguageunderstandingMinerIntentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLanguageunderstandingMinerIntentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguageunderstandingMinerIntentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguageunderstandingMinerIntentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguageunderstandingMinerIntentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguageunderstandingMinerIntentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguageunderstandingMinerIntentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguageunderstandingMinerIntentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguageunderstandingMinerIntentsOK creates a GetLanguageunderstandingMinerIntentsOK with default headers values
func NewGetLanguageunderstandingMinerIntentsOK() *GetLanguageunderstandingMinerIntentsOK {
	return &GetLanguageunderstandingMinerIntentsOK{}
}

/*
GetLanguageunderstandingMinerIntentsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLanguageunderstandingMinerIntentsOK struct {
	Payload *models.MinedIntentsListing
}

// IsSuccess returns true when this get languageunderstanding miner intents o k response has a 2xx status code
func (o *GetLanguageunderstandingMinerIntentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get languageunderstanding miner intents o k response has a 3xx status code
func (o *GetLanguageunderstandingMinerIntentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner intents o k response has a 4xx status code
func (o *GetLanguageunderstandingMinerIntentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languageunderstanding miner intents o k response has a 5xx status code
func (o *GetLanguageunderstandingMinerIntentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner intents o k response a status code equal to that given
func (o *GetLanguageunderstandingMinerIntentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLanguageunderstandingMinerIntentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsOK) GetPayload() *models.MinedIntentsListing {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MinedIntentsListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentsBadRequest creates a GetLanguageunderstandingMinerIntentsBadRequest with default headers values
func NewGetLanguageunderstandingMinerIntentsBadRequest() *GetLanguageunderstandingMinerIntentsBadRequest {
	return &GetLanguageunderstandingMinerIntentsBadRequest{}
}

/*
GetLanguageunderstandingMinerIntentsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguageunderstandingMinerIntentsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner intents bad request response has a 2xx status code
func (o *GetLanguageunderstandingMinerIntentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner intents bad request response has a 3xx status code
func (o *GetLanguageunderstandingMinerIntentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner intents bad request response has a 4xx status code
func (o *GetLanguageunderstandingMinerIntentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner intents bad request response has a 5xx status code
func (o *GetLanguageunderstandingMinerIntentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner intents bad request response a status code equal to that given
func (o *GetLanguageunderstandingMinerIntentsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLanguageunderstandingMinerIntentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentsUnauthorized creates a GetLanguageunderstandingMinerIntentsUnauthorized with default headers values
func NewGetLanguageunderstandingMinerIntentsUnauthorized() *GetLanguageunderstandingMinerIntentsUnauthorized {
	return &GetLanguageunderstandingMinerIntentsUnauthorized{}
}

/*
GetLanguageunderstandingMinerIntentsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguageunderstandingMinerIntentsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner intents unauthorized response has a 2xx status code
func (o *GetLanguageunderstandingMinerIntentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner intents unauthorized response has a 3xx status code
func (o *GetLanguageunderstandingMinerIntentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner intents unauthorized response has a 4xx status code
func (o *GetLanguageunderstandingMinerIntentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner intents unauthorized response has a 5xx status code
func (o *GetLanguageunderstandingMinerIntentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner intents unauthorized response a status code equal to that given
func (o *GetLanguageunderstandingMinerIntentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLanguageunderstandingMinerIntentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentsForbidden creates a GetLanguageunderstandingMinerIntentsForbidden with default headers values
func NewGetLanguageunderstandingMinerIntentsForbidden() *GetLanguageunderstandingMinerIntentsForbidden {
	return &GetLanguageunderstandingMinerIntentsForbidden{}
}

/*
GetLanguageunderstandingMinerIntentsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguageunderstandingMinerIntentsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner intents forbidden response has a 2xx status code
func (o *GetLanguageunderstandingMinerIntentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner intents forbidden response has a 3xx status code
func (o *GetLanguageunderstandingMinerIntentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner intents forbidden response has a 4xx status code
func (o *GetLanguageunderstandingMinerIntentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner intents forbidden response has a 5xx status code
func (o *GetLanguageunderstandingMinerIntentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner intents forbidden response a status code equal to that given
func (o *GetLanguageunderstandingMinerIntentsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLanguageunderstandingMinerIntentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentsNotFound creates a GetLanguageunderstandingMinerIntentsNotFound with default headers values
func NewGetLanguageunderstandingMinerIntentsNotFound() *GetLanguageunderstandingMinerIntentsNotFound {
	return &GetLanguageunderstandingMinerIntentsNotFound{}
}

/*
GetLanguageunderstandingMinerIntentsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetLanguageunderstandingMinerIntentsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner intents not found response has a 2xx status code
func (o *GetLanguageunderstandingMinerIntentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner intents not found response has a 3xx status code
func (o *GetLanguageunderstandingMinerIntentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner intents not found response has a 4xx status code
func (o *GetLanguageunderstandingMinerIntentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner intents not found response has a 5xx status code
func (o *GetLanguageunderstandingMinerIntentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner intents not found response a status code equal to that given
func (o *GetLanguageunderstandingMinerIntentsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLanguageunderstandingMinerIntentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentsRequestTimeout creates a GetLanguageunderstandingMinerIntentsRequestTimeout with default headers values
func NewGetLanguageunderstandingMinerIntentsRequestTimeout() *GetLanguageunderstandingMinerIntentsRequestTimeout {
	return &GetLanguageunderstandingMinerIntentsRequestTimeout{}
}

/*
GetLanguageunderstandingMinerIntentsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLanguageunderstandingMinerIntentsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner intents request timeout response has a 2xx status code
func (o *GetLanguageunderstandingMinerIntentsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner intents request timeout response has a 3xx status code
func (o *GetLanguageunderstandingMinerIntentsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner intents request timeout response has a 4xx status code
func (o *GetLanguageunderstandingMinerIntentsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner intents request timeout response has a 5xx status code
func (o *GetLanguageunderstandingMinerIntentsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner intents request timeout response a status code equal to that given
func (o *GetLanguageunderstandingMinerIntentsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetLanguageunderstandingMinerIntentsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentsRequestEntityTooLarge creates a GetLanguageunderstandingMinerIntentsRequestEntityTooLarge with default headers values
func NewGetLanguageunderstandingMinerIntentsRequestEntityTooLarge() *GetLanguageunderstandingMinerIntentsRequestEntityTooLarge {
	return &GetLanguageunderstandingMinerIntentsRequestEntityTooLarge{}
}

/*
GetLanguageunderstandingMinerIntentsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetLanguageunderstandingMinerIntentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner intents request entity too large response has a 2xx status code
func (o *GetLanguageunderstandingMinerIntentsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner intents request entity too large response has a 3xx status code
func (o *GetLanguageunderstandingMinerIntentsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner intents request entity too large response has a 4xx status code
func (o *GetLanguageunderstandingMinerIntentsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner intents request entity too large response has a 5xx status code
func (o *GetLanguageunderstandingMinerIntentsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner intents request entity too large response a status code equal to that given
func (o *GetLanguageunderstandingMinerIntentsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetLanguageunderstandingMinerIntentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentsUnsupportedMediaType creates a GetLanguageunderstandingMinerIntentsUnsupportedMediaType with default headers values
func NewGetLanguageunderstandingMinerIntentsUnsupportedMediaType() *GetLanguageunderstandingMinerIntentsUnsupportedMediaType {
	return &GetLanguageunderstandingMinerIntentsUnsupportedMediaType{}
}

/*
GetLanguageunderstandingMinerIntentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguageunderstandingMinerIntentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner intents unsupported media type response has a 2xx status code
func (o *GetLanguageunderstandingMinerIntentsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner intents unsupported media type response has a 3xx status code
func (o *GetLanguageunderstandingMinerIntentsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner intents unsupported media type response has a 4xx status code
func (o *GetLanguageunderstandingMinerIntentsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner intents unsupported media type response has a 5xx status code
func (o *GetLanguageunderstandingMinerIntentsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner intents unsupported media type response a status code equal to that given
func (o *GetLanguageunderstandingMinerIntentsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetLanguageunderstandingMinerIntentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentsTooManyRequests creates a GetLanguageunderstandingMinerIntentsTooManyRequests with default headers values
func NewGetLanguageunderstandingMinerIntentsTooManyRequests() *GetLanguageunderstandingMinerIntentsTooManyRequests {
	return &GetLanguageunderstandingMinerIntentsTooManyRequests{}
}

/*
GetLanguageunderstandingMinerIntentsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLanguageunderstandingMinerIntentsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner intents too many requests response has a 2xx status code
func (o *GetLanguageunderstandingMinerIntentsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner intents too many requests response has a 3xx status code
func (o *GetLanguageunderstandingMinerIntentsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner intents too many requests response has a 4xx status code
func (o *GetLanguageunderstandingMinerIntentsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner intents too many requests response has a 5xx status code
func (o *GetLanguageunderstandingMinerIntentsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner intents too many requests response a status code equal to that given
func (o *GetLanguageunderstandingMinerIntentsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetLanguageunderstandingMinerIntentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentsInternalServerError creates a GetLanguageunderstandingMinerIntentsInternalServerError with default headers values
func NewGetLanguageunderstandingMinerIntentsInternalServerError() *GetLanguageunderstandingMinerIntentsInternalServerError {
	return &GetLanguageunderstandingMinerIntentsInternalServerError{}
}

/*
GetLanguageunderstandingMinerIntentsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguageunderstandingMinerIntentsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner intents internal server error response has a 2xx status code
func (o *GetLanguageunderstandingMinerIntentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner intents internal server error response has a 3xx status code
func (o *GetLanguageunderstandingMinerIntentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner intents internal server error response has a 4xx status code
func (o *GetLanguageunderstandingMinerIntentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languageunderstanding miner intents internal server error response has a 5xx status code
func (o *GetLanguageunderstandingMinerIntentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get languageunderstanding miner intents internal server error response a status code equal to that given
func (o *GetLanguageunderstandingMinerIntentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLanguageunderstandingMinerIntentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentsServiceUnavailable creates a GetLanguageunderstandingMinerIntentsServiceUnavailable with default headers values
func NewGetLanguageunderstandingMinerIntentsServiceUnavailable() *GetLanguageunderstandingMinerIntentsServiceUnavailable {
	return &GetLanguageunderstandingMinerIntentsServiceUnavailable{}
}

/*
GetLanguageunderstandingMinerIntentsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguageunderstandingMinerIntentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner intents service unavailable response has a 2xx status code
func (o *GetLanguageunderstandingMinerIntentsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner intents service unavailable response has a 3xx status code
func (o *GetLanguageunderstandingMinerIntentsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner intents service unavailable response has a 4xx status code
func (o *GetLanguageunderstandingMinerIntentsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languageunderstanding miner intents service unavailable response has a 5xx status code
func (o *GetLanguageunderstandingMinerIntentsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get languageunderstanding miner intents service unavailable response a status code equal to that given
func (o *GetLanguageunderstandingMinerIntentsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetLanguageunderstandingMinerIntentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerIntentsGatewayTimeout creates a GetLanguageunderstandingMinerIntentsGatewayTimeout with default headers values
func NewGetLanguageunderstandingMinerIntentsGatewayTimeout() *GetLanguageunderstandingMinerIntentsGatewayTimeout {
	return &GetLanguageunderstandingMinerIntentsGatewayTimeout{}
}

/*
GetLanguageunderstandingMinerIntentsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetLanguageunderstandingMinerIntentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner intents gateway timeout response has a 2xx status code
func (o *GetLanguageunderstandingMinerIntentsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner intents gateway timeout response has a 3xx status code
func (o *GetLanguageunderstandingMinerIntentsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner intents gateway timeout response has a 4xx status code
func (o *GetLanguageunderstandingMinerIntentsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languageunderstanding miner intents gateway timeout response has a 5xx status code
func (o *GetLanguageunderstandingMinerIntentsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get languageunderstanding miner intents gateway timeout response a status code equal to that given
func (o *GetLanguageunderstandingMinerIntentsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetLanguageunderstandingMinerIntentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/intents][%d] getLanguageunderstandingMinerIntentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingMinerIntentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerIntentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
