// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostRoutingPredictorsReader is a Reader for the PostRoutingPredictors structure.
type PostRoutingPredictorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRoutingPredictorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRoutingPredictorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostRoutingPredictorsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRoutingPredictorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRoutingPredictorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRoutingPredictorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRoutingPredictorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRoutingPredictorsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRoutingPredictorsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRoutingPredictorsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRoutingPredictorsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRoutingPredictorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRoutingPredictorsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRoutingPredictorsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRoutingPredictorsOK creates a PostRoutingPredictorsOK with default headers values
func NewPostRoutingPredictorsOK() *PostRoutingPredictorsOK {
	return &PostRoutingPredictorsOK{}
}

/*
PostRoutingPredictorsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostRoutingPredictorsOK struct {
	Payload *models.Predictor
}

// IsSuccess returns true when this post routing predictors o k response has a 2xx status code
func (o *PostRoutingPredictorsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post routing predictors o k response has a 3xx status code
func (o *PostRoutingPredictorsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors o k response has a 4xx status code
func (o *PostRoutingPredictorsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing predictors o k response has a 5xx status code
func (o *PostRoutingPredictorsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing predictors o k response a status code equal to that given
func (o *PostRoutingPredictorsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostRoutingPredictorsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsOK  %+v", 200, o.Payload)
}

func (o *PostRoutingPredictorsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsOK  %+v", 200, o.Payload)
}

func (o *PostRoutingPredictorsOK) GetPayload() *models.Predictor {
	return o.Payload
}

func (o *PostRoutingPredictorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Predictor)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingPredictorsCreated creates a PostRoutingPredictorsCreated with default headers values
func NewPostRoutingPredictorsCreated() *PostRoutingPredictorsCreated {
	return &PostRoutingPredictorsCreated{}
}

/*
PostRoutingPredictorsCreated describes a response with status code 201, with default header values.

Predictor created.
*/
type PostRoutingPredictorsCreated struct {
	Payload *models.Predictor
}

// IsSuccess returns true when this post routing predictors created response has a 2xx status code
func (o *PostRoutingPredictorsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post routing predictors created response has a 3xx status code
func (o *PostRoutingPredictorsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors created response has a 4xx status code
func (o *PostRoutingPredictorsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing predictors created response has a 5xx status code
func (o *PostRoutingPredictorsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing predictors created response a status code equal to that given
func (o *PostRoutingPredictorsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostRoutingPredictorsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsCreated  %+v", 201, o.Payload)
}

func (o *PostRoutingPredictorsCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsCreated  %+v", 201, o.Payload)
}

func (o *PostRoutingPredictorsCreated) GetPayload() *models.Predictor {
	return o.Payload
}

func (o *PostRoutingPredictorsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Predictor)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingPredictorsBadRequest creates a PostRoutingPredictorsBadRequest with default headers values
func NewPostRoutingPredictorsBadRequest() *PostRoutingPredictorsBadRequest {
	return &PostRoutingPredictorsBadRequest{}
}

/*
PostRoutingPredictorsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRoutingPredictorsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing predictors bad request response has a 2xx status code
func (o *PostRoutingPredictorsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing predictors bad request response has a 3xx status code
func (o *PostRoutingPredictorsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors bad request response has a 4xx status code
func (o *PostRoutingPredictorsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing predictors bad request response has a 5xx status code
func (o *PostRoutingPredictorsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing predictors bad request response a status code equal to that given
func (o *PostRoutingPredictorsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostRoutingPredictorsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingPredictorsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingPredictorsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingPredictorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingPredictorsUnauthorized creates a PostRoutingPredictorsUnauthorized with default headers values
func NewPostRoutingPredictorsUnauthorized() *PostRoutingPredictorsUnauthorized {
	return &PostRoutingPredictorsUnauthorized{}
}

/*
PostRoutingPredictorsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRoutingPredictorsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing predictors unauthorized response has a 2xx status code
func (o *PostRoutingPredictorsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing predictors unauthorized response has a 3xx status code
func (o *PostRoutingPredictorsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors unauthorized response has a 4xx status code
func (o *PostRoutingPredictorsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing predictors unauthorized response has a 5xx status code
func (o *PostRoutingPredictorsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing predictors unauthorized response a status code equal to that given
func (o *PostRoutingPredictorsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostRoutingPredictorsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingPredictorsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingPredictorsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingPredictorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingPredictorsForbidden creates a PostRoutingPredictorsForbidden with default headers values
func NewPostRoutingPredictorsForbidden() *PostRoutingPredictorsForbidden {
	return &PostRoutingPredictorsForbidden{}
}

/*
PostRoutingPredictorsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostRoutingPredictorsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing predictors forbidden response has a 2xx status code
func (o *PostRoutingPredictorsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing predictors forbidden response has a 3xx status code
func (o *PostRoutingPredictorsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors forbidden response has a 4xx status code
func (o *PostRoutingPredictorsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing predictors forbidden response has a 5xx status code
func (o *PostRoutingPredictorsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing predictors forbidden response a status code equal to that given
func (o *PostRoutingPredictorsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostRoutingPredictorsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingPredictorsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingPredictorsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingPredictorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingPredictorsNotFound creates a PostRoutingPredictorsNotFound with default headers values
func NewPostRoutingPredictorsNotFound() *PostRoutingPredictorsNotFound {
	return &PostRoutingPredictorsNotFound{}
}

/*
PostRoutingPredictorsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostRoutingPredictorsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing predictors not found response has a 2xx status code
func (o *PostRoutingPredictorsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing predictors not found response has a 3xx status code
func (o *PostRoutingPredictorsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors not found response has a 4xx status code
func (o *PostRoutingPredictorsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing predictors not found response has a 5xx status code
func (o *PostRoutingPredictorsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing predictors not found response a status code equal to that given
func (o *PostRoutingPredictorsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostRoutingPredictorsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingPredictorsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingPredictorsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingPredictorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingPredictorsRequestTimeout creates a PostRoutingPredictorsRequestTimeout with default headers values
func NewPostRoutingPredictorsRequestTimeout() *PostRoutingPredictorsRequestTimeout {
	return &PostRoutingPredictorsRequestTimeout{}
}

/*
PostRoutingPredictorsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRoutingPredictorsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing predictors request timeout response has a 2xx status code
func (o *PostRoutingPredictorsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing predictors request timeout response has a 3xx status code
func (o *PostRoutingPredictorsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors request timeout response has a 4xx status code
func (o *PostRoutingPredictorsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing predictors request timeout response has a 5xx status code
func (o *PostRoutingPredictorsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing predictors request timeout response a status code equal to that given
func (o *PostRoutingPredictorsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostRoutingPredictorsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRoutingPredictorsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRoutingPredictorsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingPredictorsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingPredictorsRequestEntityTooLarge creates a PostRoutingPredictorsRequestEntityTooLarge with default headers values
func NewPostRoutingPredictorsRequestEntityTooLarge() *PostRoutingPredictorsRequestEntityTooLarge {
	return &PostRoutingPredictorsRequestEntityTooLarge{}
}

/*
PostRoutingPredictorsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostRoutingPredictorsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing predictors request entity too large response has a 2xx status code
func (o *PostRoutingPredictorsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing predictors request entity too large response has a 3xx status code
func (o *PostRoutingPredictorsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors request entity too large response has a 4xx status code
func (o *PostRoutingPredictorsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing predictors request entity too large response has a 5xx status code
func (o *PostRoutingPredictorsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing predictors request entity too large response a status code equal to that given
func (o *PostRoutingPredictorsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostRoutingPredictorsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingPredictorsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingPredictorsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingPredictorsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingPredictorsUnsupportedMediaType creates a PostRoutingPredictorsUnsupportedMediaType with default headers values
func NewPostRoutingPredictorsUnsupportedMediaType() *PostRoutingPredictorsUnsupportedMediaType {
	return &PostRoutingPredictorsUnsupportedMediaType{}
}

/*
PostRoutingPredictorsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRoutingPredictorsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing predictors unsupported media type response has a 2xx status code
func (o *PostRoutingPredictorsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing predictors unsupported media type response has a 3xx status code
func (o *PostRoutingPredictorsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors unsupported media type response has a 4xx status code
func (o *PostRoutingPredictorsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing predictors unsupported media type response has a 5xx status code
func (o *PostRoutingPredictorsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing predictors unsupported media type response a status code equal to that given
func (o *PostRoutingPredictorsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostRoutingPredictorsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingPredictorsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingPredictorsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingPredictorsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingPredictorsTooManyRequests creates a PostRoutingPredictorsTooManyRequests with default headers values
func NewPostRoutingPredictorsTooManyRequests() *PostRoutingPredictorsTooManyRequests {
	return &PostRoutingPredictorsTooManyRequests{}
}

/*
PostRoutingPredictorsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRoutingPredictorsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing predictors too many requests response has a 2xx status code
func (o *PostRoutingPredictorsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing predictors too many requests response has a 3xx status code
func (o *PostRoutingPredictorsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors too many requests response has a 4xx status code
func (o *PostRoutingPredictorsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing predictors too many requests response has a 5xx status code
func (o *PostRoutingPredictorsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing predictors too many requests response a status code equal to that given
func (o *PostRoutingPredictorsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostRoutingPredictorsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingPredictorsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingPredictorsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingPredictorsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingPredictorsInternalServerError creates a PostRoutingPredictorsInternalServerError with default headers values
func NewPostRoutingPredictorsInternalServerError() *PostRoutingPredictorsInternalServerError {
	return &PostRoutingPredictorsInternalServerError{}
}

/*
PostRoutingPredictorsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRoutingPredictorsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing predictors internal server error response has a 2xx status code
func (o *PostRoutingPredictorsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing predictors internal server error response has a 3xx status code
func (o *PostRoutingPredictorsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors internal server error response has a 4xx status code
func (o *PostRoutingPredictorsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing predictors internal server error response has a 5xx status code
func (o *PostRoutingPredictorsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post routing predictors internal server error response a status code equal to that given
func (o *PostRoutingPredictorsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostRoutingPredictorsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingPredictorsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingPredictorsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingPredictorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingPredictorsServiceUnavailable creates a PostRoutingPredictorsServiceUnavailable with default headers values
func NewPostRoutingPredictorsServiceUnavailable() *PostRoutingPredictorsServiceUnavailable {
	return &PostRoutingPredictorsServiceUnavailable{}
}

/*
PostRoutingPredictorsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRoutingPredictorsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing predictors service unavailable response has a 2xx status code
func (o *PostRoutingPredictorsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing predictors service unavailable response has a 3xx status code
func (o *PostRoutingPredictorsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors service unavailable response has a 4xx status code
func (o *PostRoutingPredictorsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing predictors service unavailable response has a 5xx status code
func (o *PostRoutingPredictorsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post routing predictors service unavailable response a status code equal to that given
func (o *PostRoutingPredictorsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostRoutingPredictorsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingPredictorsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingPredictorsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingPredictorsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingPredictorsGatewayTimeout creates a PostRoutingPredictorsGatewayTimeout with default headers values
func NewPostRoutingPredictorsGatewayTimeout() *PostRoutingPredictorsGatewayTimeout {
	return &PostRoutingPredictorsGatewayTimeout{}
}

/*
PostRoutingPredictorsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostRoutingPredictorsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing predictors gateway timeout response has a 2xx status code
func (o *PostRoutingPredictorsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing predictors gateway timeout response has a 3xx status code
func (o *PostRoutingPredictorsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing predictors gateway timeout response has a 4xx status code
func (o *PostRoutingPredictorsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing predictors gateway timeout response has a 5xx status code
func (o *PostRoutingPredictorsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post routing predictors gateway timeout response a status code equal to that given
func (o *PostRoutingPredictorsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostRoutingPredictorsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingPredictorsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/predictors][%d] postRoutingPredictorsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingPredictorsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingPredictorsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
