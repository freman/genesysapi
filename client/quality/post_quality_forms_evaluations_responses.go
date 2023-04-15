// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostQualityFormsEvaluationsReader is a Reader for the PostQualityFormsEvaluations structure.
type PostQualityFormsEvaluationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualityFormsEvaluationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualityFormsEvaluationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualityFormsEvaluationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualityFormsEvaluationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualityFormsEvaluationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualityFormsEvaluationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostQualityFormsEvaluationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostQualityFormsEvaluationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualityFormsEvaluationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualityFormsEvaluationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualityFormsEvaluationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualityFormsEvaluationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualityFormsEvaluationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualityFormsEvaluationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualityFormsEvaluationsOK creates a PostQualityFormsEvaluationsOK with default headers values
func NewPostQualityFormsEvaluationsOK() *PostQualityFormsEvaluationsOK {
	return &PostQualityFormsEvaluationsOK{}
}

/*
PostQualityFormsEvaluationsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostQualityFormsEvaluationsOK struct {
	Payload *models.EvaluationForm
}

// IsSuccess returns true when this post quality forms evaluations o k response has a 2xx status code
func (o *PostQualityFormsEvaluationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post quality forms evaluations o k response has a 3xx status code
func (o *PostQualityFormsEvaluationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations o k response has a 4xx status code
func (o *PostQualityFormsEvaluationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality forms evaluations o k response has a 5xx status code
func (o *PostQualityFormsEvaluationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms evaluations o k response a status code equal to that given
func (o *PostQualityFormsEvaluationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostQualityFormsEvaluationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsOK  %+v", 200, o.Payload)
}

func (o *PostQualityFormsEvaluationsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsOK  %+v", 200, o.Payload)
}

func (o *PostQualityFormsEvaluationsOK) GetPayload() *models.EvaluationForm {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationForm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsBadRequest creates a PostQualityFormsEvaluationsBadRequest with default headers values
func NewPostQualityFormsEvaluationsBadRequest() *PostQualityFormsEvaluationsBadRequest {
	return &PostQualityFormsEvaluationsBadRequest{}
}

/*
PostQualityFormsEvaluationsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualityFormsEvaluationsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms evaluations bad request response has a 2xx status code
func (o *PostQualityFormsEvaluationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms evaluations bad request response has a 3xx status code
func (o *PostQualityFormsEvaluationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations bad request response has a 4xx status code
func (o *PostQualityFormsEvaluationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms evaluations bad request response has a 5xx status code
func (o *PostQualityFormsEvaluationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms evaluations bad request response a status code equal to that given
func (o *PostQualityFormsEvaluationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostQualityFormsEvaluationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityFormsEvaluationsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityFormsEvaluationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsUnauthorized creates a PostQualityFormsEvaluationsUnauthorized with default headers values
func NewPostQualityFormsEvaluationsUnauthorized() *PostQualityFormsEvaluationsUnauthorized {
	return &PostQualityFormsEvaluationsUnauthorized{}
}

/*
PostQualityFormsEvaluationsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualityFormsEvaluationsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms evaluations unauthorized response has a 2xx status code
func (o *PostQualityFormsEvaluationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms evaluations unauthorized response has a 3xx status code
func (o *PostQualityFormsEvaluationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations unauthorized response has a 4xx status code
func (o *PostQualityFormsEvaluationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms evaluations unauthorized response has a 5xx status code
func (o *PostQualityFormsEvaluationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms evaluations unauthorized response a status code equal to that given
func (o *PostQualityFormsEvaluationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostQualityFormsEvaluationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityFormsEvaluationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityFormsEvaluationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsForbidden creates a PostQualityFormsEvaluationsForbidden with default headers values
func NewPostQualityFormsEvaluationsForbidden() *PostQualityFormsEvaluationsForbidden {
	return &PostQualityFormsEvaluationsForbidden{}
}

/*
PostQualityFormsEvaluationsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostQualityFormsEvaluationsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms evaluations forbidden response has a 2xx status code
func (o *PostQualityFormsEvaluationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms evaluations forbidden response has a 3xx status code
func (o *PostQualityFormsEvaluationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations forbidden response has a 4xx status code
func (o *PostQualityFormsEvaluationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms evaluations forbidden response has a 5xx status code
func (o *PostQualityFormsEvaluationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms evaluations forbidden response a status code equal to that given
func (o *PostQualityFormsEvaluationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostQualityFormsEvaluationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityFormsEvaluationsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityFormsEvaluationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsNotFound creates a PostQualityFormsEvaluationsNotFound with default headers values
func NewPostQualityFormsEvaluationsNotFound() *PostQualityFormsEvaluationsNotFound {
	return &PostQualityFormsEvaluationsNotFound{}
}

/*
PostQualityFormsEvaluationsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostQualityFormsEvaluationsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms evaluations not found response has a 2xx status code
func (o *PostQualityFormsEvaluationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms evaluations not found response has a 3xx status code
func (o *PostQualityFormsEvaluationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations not found response has a 4xx status code
func (o *PostQualityFormsEvaluationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms evaluations not found response has a 5xx status code
func (o *PostQualityFormsEvaluationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms evaluations not found response a status code equal to that given
func (o *PostQualityFormsEvaluationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostQualityFormsEvaluationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityFormsEvaluationsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityFormsEvaluationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsRequestTimeout creates a PostQualityFormsEvaluationsRequestTimeout with default headers values
func NewPostQualityFormsEvaluationsRequestTimeout() *PostQualityFormsEvaluationsRequestTimeout {
	return &PostQualityFormsEvaluationsRequestTimeout{}
}

/*
PostQualityFormsEvaluationsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostQualityFormsEvaluationsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms evaluations request timeout response has a 2xx status code
func (o *PostQualityFormsEvaluationsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms evaluations request timeout response has a 3xx status code
func (o *PostQualityFormsEvaluationsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations request timeout response has a 4xx status code
func (o *PostQualityFormsEvaluationsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms evaluations request timeout response has a 5xx status code
func (o *PostQualityFormsEvaluationsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms evaluations request timeout response a status code equal to that given
func (o *PostQualityFormsEvaluationsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostQualityFormsEvaluationsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityFormsEvaluationsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityFormsEvaluationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsConflict creates a PostQualityFormsEvaluationsConflict with default headers values
func NewPostQualityFormsEvaluationsConflict() *PostQualityFormsEvaluationsConflict {
	return &PostQualityFormsEvaluationsConflict{}
}

/*
PostQualityFormsEvaluationsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostQualityFormsEvaluationsConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms evaluations conflict response has a 2xx status code
func (o *PostQualityFormsEvaluationsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms evaluations conflict response has a 3xx status code
func (o *PostQualityFormsEvaluationsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations conflict response has a 4xx status code
func (o *PostQualityFormsEvaluationsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms evaluations conflict response has a 5xx status code
func (o *PostQualityFormsEvaluationsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms evaluations conflict response a status code equal to that given
func (o *PostQualityFormsEvaluationsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostQualityFormsEvaluationsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsConflict  %+v", 409, o.Payload)
}

func (o *PostQualityFormsEvaluationsConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsConflict  %+v", 409, o.Payload)
}

func (o *PostQualityFormsEvaluationsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsRequestEntityTooLarge creates a PostQualityFormsEvaluationsRequestEntityTooLarge with default headers values
func NewPostQualityFormsEvaluationsRequestEntityTooLarge() *PostQualityFormsEvaluationsRequestEntityTooLarge {
	return &PostQualityFormsEvaluationsRequestEntityTooLarge{}
}

/*
PostQualityFormsEvaluationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostQualityFormsEvaluationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms evaluations request entity too large response has a 2xx status code
func (o *PostQualityFormsEvaluationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms evaluations request entity too large response has a 3xx status code
func (o *PostQualityFormsEvaluationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations request entity too large response has a 4xx status code
func (o *PostQualityFormsEvaluationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms evaluations request entity too large response has a 5xx status code
func (o *PostQualityFormsEvaluationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms evaluations request entity too large response a status code equal to that given
func (o *PostQualityFormsEvaluationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostQualityFormsEvaluationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityFormsEvaluationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityFormsEvaluationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsUnsupportedMediaType creates a PostQualityFormsEvaluationsUnsupportedMediaType with default headers values
func NewPostQualityFormsEvaluationsUnsupportedMediaType() *PostQualityFormsEvaluationsUnsupportedMediaType {
	return &PostQualityFormsEvaluationsUnsupportedMediaType{}
}

/*
PostQualityFormsEvaluationsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualityFormsEvaluationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms evaluations unsupported media type response has a 2xx status code
func (o *PostQualityFormsEvaluationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms evaluations unsupported media type response has a 3xx status code
func (o *PostQualityFormsEvaluationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations unsupported media type response has a 4xx status code
func (o *PostQualityFormsEvaluationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms evaluations unsupported media type response has a 5xx status code
func (o *PostQualityFormsEvaluationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms evaluations unsupported media type response a status code equal to that given
func (o *PostQualityFormsEvaluationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostQualityFormsEvaluationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityFormsEvaluationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityFormsEvaluationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsTooManyRequests creates a PostQualityFormsEvaluationsTooManyRequests with default headers values
func NewPostQualityFormsEvaluationsTooManyRequests() *PostQualityFormsEvaluationsTooManyRequests {
	return &PostQualityFormsEvaluationsTooManyRequests{}
}

/*
PostQualityFormsEvaluationsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostQualityFormsEvaluationsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms evaluations too many requests response has a 2xx status code
func (o *PostQualityFormsEvaluationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms evaluations too many requests response has a 3xx status code
func (o *PostQualityFormsEvaluationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations too many requests response has a 4xx status code
func (o *PostQualityFormsEvaluationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms evaluations too many requests response has a 5xx status code
func (o *PostQualityFormsEvaluationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms evaluations too many requests response a status code equal to that given
func (o *PostQualityFormsEvaluationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostQualityFormsEvaluationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityFormsEvaluationsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityFormsEvaluationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsInternalServerError creates a PostQualityFormsEvaluationsInternalServerError with default headers values
func NewPostQualityFormsEvaluationsInternalServerError() *PostQualityFormsEvaluationsInternalServerError {
	return &PostQualityFormsEvaluationsInternalServerError{}
}

/*
PostQualityFormsEvaluationsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualityFormsEvaluationsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms evaluations internal server error response has a 2xx status code
func (o *PostQualityFormsEvaluationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms evaluations internal server error response has a 3xx status code
func (o *PostQualityFormsEvaluationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations internal server error response has a 4xx status code
func (o *PostQualityFormsEvaluationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality forms evaluations internal server error response has a 5xx status code
func (o *PostQualityFormsEvaluationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality forms evaluations internal server error response a status code equal to that given
func (o *PostQualityFormsEvaluationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostQualityFormsEvaluationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityFormsEvaluationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityFormsEvaluationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsServiceUnavailable creates a PostQualityFormsEvaluationsServiceUnavailable with default headers values
func NewPostQualityFormsEvaluationsServiceUnavailable() *PostQualityFormsEvaluationsServiceUnavailable {
	return &PostQualityFormsEvaluationsServiceUnavailable{}
}

/*
PostQualityFormsEvaluationsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualityFormsEvaluationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms evaluations service unavailable response has a 2xx status code
func (o *PostQualityFormsEvaluationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms evaluations service unavailable response has a 3xx status code
func (o *PostQualityFormsEvaluationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations service unavailable response has a 4xx status code
func (o *PostQualityFormsEvaluationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality forms evaluations service unavailable response has a 5xx status code
func (o *PostQualityFormsEvaluationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality forms evaluations service unavailable response a status code equal to that given
func (o *PostQualityFormsEvaluationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostQualityFormsEvaluationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityFormsEvaluationsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityFormsEvaluationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsGatewayTimeout creates a PostQualityFormsEvaluationsGatewayTimeout with default headers values
func NewPostQualityFormsEvaluationsGatewayTimeout() *PostQualityFormsEvaluationsGatewayTimeout {
	return &PostQualityFormsEvaluationsGatewayTimeout{}
}

/*
PostQualityFormsEvaluationsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostQualityFormsEvaluationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms evaluations gateway timeout response has a 2xx status code
func (o *PostQualityFormsEvaluationsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms evaluations gateway timeout response has a 3xx status code
func (o *PostQualityFormsEvaluationsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms evaluations gateway timeout response has a 4xx status code
func (o *PostQualityFormsEvaluationsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality forms evaluations gateway timeout response has a 5xx status code
func (o *PostQualityFormsEvaluationsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality forms evaluations gateway timeout response a status code equal to that given
func (o *PostQualityFormsEvaluationsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostQualityFormsEvaluationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityFormsEvaluationsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityFormsEvaluationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
