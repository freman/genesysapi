// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetJourneyOutcomesPredictorsReader is a Reader for the GetJourneyOutcomesPredictors structure.
type GetJourneyOutcomesPredictorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJourneyOutcomesPredictorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJourneyOutcomesPredictorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJourneyOutcomesPredictorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetJourneyOutcomesPredictorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJourneyOutcomesPredictorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJourneyOutcomesPredictorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetJourneyOutcomesPredictorsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetJourneyOutcomesPredictorsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetJourneyOutcomesPredictorsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetJourneyOutcomesPredictorsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJourneyOutcomesPredictorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetJourneyOutcomesPredictorsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetJourneyOutcomesPredictorsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJourneyOutcomesPredictorsOK creates a GetJourneyOutcomesPredictorsOK with default headers values
func NewGetJourneyOutcomesPredictorsOK() *GetJourneyOutcomesPredictorsOK {
	return &GetJourneyOutcomesPredictorsOK{}
}

/*
GetJourneyOutcomesPredictorsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetJourneyOutcomesPredictorsOK struct {
	Payload *models.OutcomePredictorListing
}

// IsSuccess returns true when this get journey outcomes predictors o k response has a 2xx status code
func (o *GetJourneyOutcomesPredictorsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get journey outcomes predictors o k response has a 3xx status code
func (o *GetJourneyOutcomesPredictorsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcomes predictors o k response has a 4xx status code
func (o *GetJourneyOutcomesPredictorsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey outcomes predictors o k response has a 5xx status code
func (o *GetJourneyOutcomesPredictorsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcomes predictors o k response a status code equal to that given
func (o *GetJourneyOutcomesPredictorsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetJourneyOutcomesPredictorsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsOK  %+v", 200, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsOK  %+v", 200, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsOK) GetPayload() *models.OutcomePredictorListing {
	return o.Payload
}

func (o *GetJourneyOutcomesPredictorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutcomePredictorListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomesPredictorsBadRequest creates a GetJourneyOutcomesPredictorsBadRequest with default headers values
func NewGetJourneyOutcomesPredictorsBadRequest() *GetJourneyOutcomesPredictorsBadRequest {
	return &GetJourneyOutcomesPredictorsBadRequest{}
}

/*
GetJourneyOutcomesPredictorsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetJourneyOutcomesPredictorsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcomes predictors bad request response has a 2xx status code
func (o *GetJourneyOutcomesPredictorsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcomes predictors bad request response has a 3xx status code
func (o *GetJourneyOutcomesPredictorsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcomes predictors bad request response has a 4xx status code
func (o *GetJourneyOutcomesPredictorsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcomes predictors bad request response has a 5xx status code
func (o *GetJourneyOutcomesPredictorsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcomes predictors bad request response a status code equal to that given
func (o *GetJourneyOutcomesPredictorsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetJourneyOutcomesPredictorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomesPredictorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomesPredictorsUnauthorized creates a GetJourneyOutcomesPredictorsUnauthorized with default headers values
func NewGetJourneyOutcomesPredictorsUnauthorized() *GetJourneyOutcomesPredictorsUnauthorized {
	return &GetJourneyOutcomesPredictorsUnauthorized{}
}

/*
GetJourneyOutcomesPredictorsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetJourneyOutcomesPredictorsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcomes predictors unauthorized response has a 2xx status code
func (o *GetJourneyOutcomesPredictorsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcomes predictors unauthorized response has a 3xx status code
func (o *GetJourneyOutcomesPredictorsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcomes predictors unauthorized response has a 4xx status code
func (o *GetJourneyOutcomesPredictorsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcomes predictors unauthorized response has a 5xx status code
func (o *GetJourneyOutcomesPredictorsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcomes predictors unauthorized response a status code equal to that given
func (o *GetJourneyOutcomesPredictorsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetJourneyOutcomesPredictorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomesPredictorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomesPredictorsForbidden creates a GetJourneyOutcomesPredictorsForbidden with default headers values
func NewGetJourneyOutcomesPredictorsForbidden() *GetJourneyOutcomesPredictorsForbidden {
	return &GetJourneyOutcomesPredictorsForbidden{}
}

/*
GetJourneyOutcomesPredictorsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetJourneyOutcomesPredictorsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcomes predictors forbidden response has a 2xx status code
func (o *GetJourneyOutcomesPredictorsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcomes predictors forbidden response has a 3xx status code
func (o *GetJourneyOutcomesPredictorsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcomes predictors forbidden response has a 4xx status code
func (o *GetJourneyOutcomesPredictorsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcomes predictors forbidden response has a 5xx status code
func (o *GetJourneyOutcomesPredictorsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcomes predictors forbidden response a status code equal to that given
func (o *GetJourneyOutcomesPredictorsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetJourneyOutcomesPredictorsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomesPredictorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomesPredictorsNotFound creates a GetJourneyOutcomesPredictorsNotFound with default headers values
func NewGetJourneyOutcomesPredictorsNotFound() *GetJourneyOutcomesPredictorsNotFound {
	return &GetJourneyOutcomesPredictorsNotFound{}
}

/*
GetJourneyOutcomesPredictorsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetJourneyOutcomesPredictorsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcomes predictors not found response has a 2xx status code
func (o *GetJourneyOutcomesPredictorsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcomes predictors not found response has a 3xx status code
func (o *GetJourneyOutcomesPredictorsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcomes predictors not found response has a 4xx status code
func (o *GetJourneyOutcomesPredictorsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcomes predictors not found response has a 5xx status code
func (o *GetJourneyOutcomesPredictorsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcomes predictors not found response a status code equal to that given
func (o *GetJourneyOutcomesPredictorsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetJourneyOutcomesPredictorsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomesPredictorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomesPredictorsRequestTimeout creates a GetJourneyOutcomesPredictorsRequestTimeout with default headers values
func NewGetJourneyOutcomesPredictorsRequestTimeout() *GetJourneyOutcomesPredictorsRequestTimeout {
	return &GetJourneyOutcomesPredictorsRequestTimeout{}
}

/*
GetJourneyOutcomesPredictorsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetJourneyOutcomesPredictorsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcomes predictors request timeout response has a 2xx status code
func (o *GetJourneyOutcomesPredictorsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcomes predictors request timeout response has a 3xx status code
func (o *GetJourneyOutcomesPredictorsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcomes predictors request timeout response has a 4xx status code
func (o *GetJourneyOutcomesPredictorsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcomes predictors request timeout response has a 5xx status code
func (o *GetJourneyOutcomesPredictorsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcomes predictors request timeout response a status code equal to that given
func (o *GetJourneyOutcomesPredictorsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetJourneyOutcomesPredictorsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomesPredictorsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomesPredictorsRequestEntityTooLarge creates a GetJourneyOutcomesPredictorsRequestEntityTooLarge with default headers values
func NewGetJourneyOutcomesPredictorsRequestEntityTooLarge() *GetJourneyOutcomesPredictorsRequestEntityTooLarge {
	return &GetJourneyOutcomesPredictorsRequestEntityTooLarge{}
}

/*
GetJourneyOutcomesPredictorsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetJourneyOutcomesPredictorsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcomes predictors request entity too large response has a 2xx status code
func (o *GetJourneyOutcomesPredictorsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcomes predictors request entity too large response has a 3xx status code
func (o *GetJourneyOutcomesPredictorsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcomes predictors request entity too large response has a 4xx status code
func (o *GetJourneyOutcomesPredictorsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcomes predictors request entity too large response has a 5xx status code
func (o *GetJourneyOutcomesPredictorsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcomes predictors request entity too large response a status code equal to that given
func (o *GetJourneyOutcomesPredictorsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetJourneyOutcomesPredictorsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomesPredictorsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomesPredictorsUnsupportedMediaType creates a GetJourneyOutcomesPredictorsUnsupportedMediaType with default headers values
func NewGetJourneyOutcomesPredictorsUnsupportedMediaType() *GetJourneyOutcomesPredictorsUnsupportedMediaType {
	return &GetJourneyOutcomesPredictorsUnsupportedMediaType{}
}

/*
GetJourneyOutcomesPredictorsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetJourneyOutcomesPredictorsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcomes predictors unsupported media type response has a 2xx status code
func (o *GetJourneyOutcomesPredictorsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcomes predictors unsupported media type response has a 3xx status code
func (o *GetJourneyOutcomesPredictorsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcomes predictors unsupported media type response has a 4xx status code
func (o *GetJourneyOutcomesPredictorsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcomes predictors unsupported media type response has a 5xx status code
func (o *GetJourneyOutcomesPredictorsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcomes predictors unsupported media type response a status code equal to that given
func (o *GetJourneyOutcomesPredictorsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetJourneyOutcomesPredictorsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomesPredictorsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomesPredictorsTooManyRequests creates a GetJourneyOutcomesPredictorsTooManyRequests with default headers values
func NewGetJourneyOutcomesPredictorsTooManyRequests() *GetJourneyOutcomesPredictorsTooManyRequests {
	return &GetJourneyOutcomesPredictorsTooManyRequests{}
}

/*
GetJourneyOutcomesPredictorsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetJourneyOutcomesPredictorsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcomes predictors too many requests response has a 2xx status code
func (o *GetJourneyOutcomesPredictorsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcomes predictors too many requests response has a 3xx status code
func (o *GetJourneyOutcomesPredictorsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcomes predictors too many requests response has a 4xx status code
func (o *GetJourneyOutcomesPredictorsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey outcomes predictors too many requests response has a 5xx status code
func (o *GetJourneyOutcomesPredictorsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey outcomes predictors too many requests response a status code equal to that given
func (o *GetJourneyOutcomesPredictorsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetJourneyOutcomesPredictorsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomesPredictorsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomesPredictorsInternalServerError creates a GetJourneyOutcomesPredictorsInternalServerError with default headers values
func NewGetJourneyOutcomesPredictorsInternalServerError() *GetJourneyOutcomesPredictorsInternalServerError {
	return &GetJourneyOutcomesPredictorsInternalServerError{}
}

/*
GetJourneyOutcomesPredictorsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetJourneyOutcomesPredictorsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcomes predictors internal server error response has a 2xx status code
func (o *GetJourneyOutcomesPredictorsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcomes predictors internal server error response has a 3xx status code
func (o *GetJourneyOutcomesPredictorsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcomes predictors internal server error response has a 4xx status code
func (o *GetJourneyOutcomesPredictorsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey outcomes predictors internal server error response has a 5xx status code
func (o *GetJourneyOutcomesPredictorsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey outcomes predictors internal server error response a status code equal to that given
func (o *GetJourneyOutcomesPredictorsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetJourneyOutcomesPredictorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomesPredictorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomesPredictorsServiceUnavailable creates a GetJourneyOutcomesPredictorsServiceUnavailable with default headers values
func NewGetJourneyOutcomesPredictorsServiceUnavailable() *GetJourneyOutcomesPredictorsServiceUnavailable {
	return &GetJourneyOutcomesPredictorsServiceUnavailable{}
}

/*
GetJourneyOutcomesPredictorsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetJourneyOutcomesPredictorsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcomes predictors service unavailable response has a 2xx status code
func (o *GetJourneyOutcomesPredictorsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcomes predictors service unavailable response has a 3xx status code
func (o *GetJourneyOutcomesPredictorsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcomes predictors service unavailable response has a 4xx status code
func (o *GetJourneyOutcomesPredictorsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey outcomes predictors service unavailable response has a 5xx status code
func (o *GetJourneyOutcomesPredictorsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey outcomes predictors service unavailable response a status code equal to that given
func (o *GetJourneyOutcomesPredictorsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetJourneyOutcomesPredictorsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomesPredictorsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomesPredictorsGatewayTimeout creates a GetJourneyOutcomesPredictorsGatewayTimeout with default headers values
func NewGetJourneyOutcomesPredictorsGatewayTimeout() *GetJourneyOutcomesPredictorsGatewayTimeout {
	return &GetJourneyOutcomesPredictorsGatewayTimeout{}
}

/*
GetJourneyOutcomesPredictorsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetJourneyOutcomesPredictorsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey outcomes predictors gateway timeout response has a 2xx status code
func (o *GetJourneyOutcomesPredictorsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey outcomes predictors gateway timeout response has a 3xx status code
func (o *GetJourneyOutcomesPredictorsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey outcomes predictors gateway timeout response has a 4xx status code
func (o *GetJourneyOutcomesPredictorsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey outcomes predictors gateway timeout response has a 5xx status code
func (o *GetJourneyOutcomesPredictorsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey outcomes predictors gateway timeout response a status code equal to that given
func (o *GetJourneyOutcomesPredictorsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetJourneyOutcomesPredictorsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/predictors][%d] getJourneyOutcomesPredictorsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyOutcomesPredictorsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomesPredictorsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
