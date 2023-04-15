// Code generated by go-swagger; DO NOT EDIT.

package learning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostLearningAssessmentsScoringReader is a Reader for the PostLearningAssessmentsScoring structure.
type PostLearningAssessmentsScoringReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearningAssessmentsScoringReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLearningAssessmentsScoringOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearningAssessmentsScoringBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLearningAssessmentsScoringUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearningAssessmentsScoringForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLearningAssessmentsScoringNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLearningAssessmentsScoringRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLearningAssessmentsScoringRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLearningAssessmentsScoringUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLearningAssessmentsScoringTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLearningAssessmentsScoringInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLearningAssessmentsScoringServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLearningAssessmentsScoringGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearningAssessmentsScoringOK creates a PostLearningAssessmentsScoringOK with default headers values
func NewPostLearningAssessmentsScoringOK() *PostLearningAssessmentsScoringOK {
	return &PostLearningAssessmentsScoringOK{}
}

/*
PostLearningAssessmentsScoringOK describes a response with status code 200, with default header values.

successful operation
*/
type PostLearningAssessmentsScoringOK struct {
	Payload *models.AssessmentScoringSet
}

// IsSuccess returns true when this post learning assessments scoring o k response has a 2xx status code
func (o *PostLearningAssessmentsScoringOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post learning assessments scoring o k response has a 3xx status code
func (o *PostLearningAssessmentsScoringOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assessments scoring o k response has a 4xx status code
func (o *PostLearningAssessmentsScoringOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post learning assessments scoring o k response has a 5xx status code
func (o *PostLearningAssessmentsScoringOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assessments scoring o k response a status code equal to that given
func (o *PostLearningAssessmentsScoringOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostLearningAssessmentsScoringOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringOK  %+v", 200, o.Payload)
}

func (o *PostLearningAssessmentsScoringOK) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringOK  %+v", 200, o.Payload)
}

func (o *PostLearningAssessmentsScoringOK) GetPayload() *models.AssessmentScoringSet {
	return o.Payload
}

func (o *PostLearningAssessmentsScoringOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssessmentScoringSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssessmentsScoringBadRequest creates a PostLearningAssessmentsScoringBadRequest with default headers values
func NewPostLearningAssessmentsScoringBadRequest() *PostLearningAssessmentsScoringBadRequest {
	return &PostLearningAssessmentsScoringBadRequest{}
}

/*
PostLearningAssessmentsScoringBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLearningAssessmentsScoringBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assessments scoring bad request response has a 2xx status code
func (o *PostLearningAssessmentsScoringBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assessments scoring bad request response has a 3xx status code
func (o *PostLearningAssessmentsScoringBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assessments scoring bad request response has a 4xx status code
func (o *PostLearningAssessmentsScoringBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assessments scoring bad request response has a 5xx status code
func (o *PostLearningAssessmentsScoringBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assessments scoring bad request response a status code equal to that given
func (o *PostLearningAssessmentsScoringBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostLearningAssessmentsScoringBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearningAssessmentsScoringBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearningAssessmentsScoringBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssessmentsScoringBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssessmentsScoringUnauthorized creates a PostLearningAssessmentsScoringUnauthorized with default headers values
func NewPostLearningAssessmentsScoringUnauthorized() *PostLearningAssessmentsScoringUnauthorized {
	return &PostLearningAssessmentsScoringUnauthorized{}
}

/*
PostLearningAssessmentsScoringUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLearningAssessmentsScoringUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assessments scoring unauthorized response has a 2xx status code
func (o *PostLearningAssessmentsScoringUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assessments scoring unauthorized response has a 3xx status code
func (o *PostLearningAssessmentsScoringUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assessments scoring unauthorized response has a 4xx status code
func (o *PostLearningAssessmentsScoringUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assessments scoring unauthorized response has a 5xx status code
func (o *PostLearningAssessmentsScoringUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assessments scoring unauthorized response a status code equal to that given
func (o *PostLearningAssessmentsScoringUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostLearningAssessmentsScoringUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLearningAssessmentsScoringUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLearningAssessmentsScoringUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssessmentsScoringUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssessmentsScoringForbidden creates a PostLearningAssessmentsScoringForbidden with default headers values
func NewPostLearningAssessmentsScoringForbidden() *PostLearningAssessmentsScoringForbidden {
	return &PostLearningAssessmentsScoringForbidden{}
}

/*
PostLearningAssessmentsScoringForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostLearningAssessmentsScoringForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assessments scoring forbidden response has a 2xx status code
func (o *PostLearningAssessmentsScoringForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assessments scoring forbidden response has a 3xx status code
func (o *PostLearningAssessmentsScoringForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assessments scoring forbidden response has a 4xx status code
func (o *PostLearningAssessmentsScoringForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assessments scoring forbidden response has a 5xx status code
func (o *PostLearningAssessmentsScoringForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assessments scoring forbidden response a status code equal to that given
func (o *PostLearningAssessmentsScoringForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostLearningAssessmentsScoringForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringForbidden  %+v", 403, o.Payload)
}

func (o *PostLearningAssessmentsScoringForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringForbidden  %+v", 403, o.Payload)
}

func (o *PostLearningAssessmentsScoringForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssessmentsScoringForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssessmentsScoringNotFound creates a PostLearningAssessmentsScoringNotFound with default headers values
func NewPostLearningAssessmentsScoringNotFound() *PostLearningAssessmentsScoringNotFound {
	return &PostLearningAssessmentsScoringNotFound{}
}

/*
PostLearningAssessmentsScoringNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostLearningAssessmentsScoringNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assessments scoring not found response has a 2xx status code
func (o *PostLearningAssessmentsScoringNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assessments scoring not found response has a 3xx status code
func (o *PostLearningAssessmentsScoringNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assessments scoring not found response has a 4xx status code
func (o *PostLearningAssessmentsScoringNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assessments scoring not found response has a 5xx status code
func (o *PostLearningAssessmentsScoringNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assessments scoring not found response a status code equal to that given
func (o *PostLearningAssessmentsScoringNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostLearningAssessmentsScoringNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringNotFound  %+v", 404, o.Payload)
}

func (o *PostLearningAssessmentsScoringNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringNotFound  %+v", 404, o.Payload)
}

func (o *PostLearningAssessmentsScoringNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssessmentsScoringNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssessmentsScoringRequestTimeout creates a PostLearningAssessmentsScoringRequestTimeout with default headers values
func NewPostLearningAssessmentsScoringRequestTimeout() *PostLearningAssessmentsScoringRequestTimeout {
	return &PostLearningAssessmentsScoringRequestTimeout{}
}

/*
PostLearningAssessmentsScoringRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLearningAssessmentsScoringRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assessments scoring request timeout response has a 2xx status code
func (o *PostLearningAssessmentsScoringRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assessments scoring request timeout response has a 3xx status code
func (o *PostLearningAssessmentsScoringRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assessments scoring request timeout response has a 4xx status code
func (o *PostLearningAssessmentsScoringRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assessments scoring request timeout response has a 5xx status code
func (o *PostLearningAssessmentsScoringRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assessments scoring request timeout response a status code equal to that given
func (o *PostLearningAssessmentsScoringRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostLearningAssessmentsScoringRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLearningAssessmentsScoringRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLearningAssessmentsScoringRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssessmentsScoringRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssessmentsScoringRequestEntityTooLarge creates a PostLearningAssessmentsScoringRequestEntityTooLarge with default headers values
func NewPostLearningAssessmentsScoringRequestEntityTooLarge() *PostLearningAssessmentsScoringRequestEntityTooLarge {
	return &PostLearningAssessmentsScoringRequestEntityTooLarge{}
}

/*
PostLearningAssessmentsScoringRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostLearningAssessmentsScoringRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assessments scoring request entity too large response has a 2xx status code
func (o *PostLearningAssessmentsScoringRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assessments scoring request entity too large response has a 3xx status code
func (o *PostLearningAssessmentsScoringRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assessments scoring request entity too large response has a 4xx status code
func (o *PostLearningAssessmentsScoringRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assessments scoring request entity too large response has a 5xx status code
func (o *PostLearningAssessmentsScoringRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assessments scoring request entity too large response a status code equal to that given
func (o *PostLearningAssessmentsScoringRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostLearningAssessmentsScoringRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLearningAssessmentsScoringRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLearningAssessmentsScoringRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssessmentsScoringRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssessmentsScoringUnsupportedMediaType creates a PostLearningAssessmentsScoringUnsupportedMediaType with default headers values
func NewPostLearningAssessmentsScoringUnsupportedMediaType() *PostLearningAssessmentsScoringUnsupportedMediaType {
	return &PostLearningAssessmentsScoringUnsupportedMediaType{}
}

/*
PostLearningAssessmentsScoringUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLearningAssessmentsScoringUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assessments scoring unsupported media type response has a 2xx status code
func (o *PostLearningAssessmentsScoringUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assessments scoring unsupported media type response has a 3xx status code
func (o *PostLearningAssessmentsScoringUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assessments scoring unsupported media type response has a 4xx status code
func (o *PostLearningAssessmentsScoringUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assessments scoring unsupported media type response has a 5xx status code
func (o *PostLearningAssessmentsScoringUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assessments scoring unsupported media type response a status code equal to that given
func (o *PostLearningAssessmentsScoringUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostLearningAssessmentsScoringUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLearningAssessmentsScoringUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLearningAssessmentsScoringUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssessmentsScoringUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssessmentsScoringTooManyRequests creates a PostLearningAssessmentsScoringTooManyRequests with default headers values
func NewPostLearningAssessmentsScoringTooManyRequests() *PostLearningAssessmentsScoringTooManyRequests {
	return &PostLearningAssessmentsScoringTooManyRequests{}
}

/*
PostLearningAssessmentsScoringTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLearningAssessmentsScoringTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assessments scoring too many requests response has a 2xx status code
func (o *PostLearningAssessmentsScoringTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assessments scoring too many requests response has a 3xx status code
func (o *PostLearningAssessmentsScoringTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assessments scoring too many requests response has a 4xx status code
func (o *PostLearningAssessmentsScoringTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post learning assessments scoring too many requests response has a 5xx status code
func (o *PostLearningAssessmentsScoringTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post learning assessments scoring too many requests response a status code equal to that given
func (o *PostLearningAssessmentsScoringTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostLearningAssessmentsScoringTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLearningAssessmentsScoringTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLearningAssessmentsScoringTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssessmentsScoringTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssessmentsScoringInternalServerError creates a PostLearningAssessmentsScoringInternalServerError with default headers values
func NewPostLearningAssessmentsScoringInternalServerError() *PostLearningAssessmentsScoringInternalServerError {
	return &PostLearningAssessmentsScoringInternalServerError{}
}

/*
PostLearningAssessmentsScoringInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLearningAssessmentsScoringInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assessments scoring internal server error response has a 2xx status code
func (o *PostLearningAssessmentsScoringInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assessments scoring internal server error response has a 3xx status code
func (o *PostLearningAssessmentsScoringInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assessments scoring internal server error response has a 4xx status code
func (o *PostLearningAssessmentsScoringInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post learning assessments scoring internal server error response has a 5xx status code
func (o *PostLearningAssessmentsScoringInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post learning assessments scoring internal server error response a status code equal to that given
func (o *PostLearningAssessmentsScoringInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostLearningAssessmentsScoringInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLearningAssessmentsScoringInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLearningAssessmentsScoringInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssessmentsScoringInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssessmentsScoringServiceUnavailable creates a PostLearningAssessmentsScoringServiceUnavailable with default headers values
func NewPostLearningAssessmentsScoringServiceUnavailable() *PostLearningAssessmentsScoringServiceUnavailable {
	return &PostLearningAssessmentsScoringServiceUnavailable{}
}

/*
PostLearningAssessmentsScoringServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLearningAssessmentsScoringServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assessments scoring service unavailable response has a 2xx status code
func (o *PostLearningAssessmentsScoringServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assessments scoring service unavailable response has a 3xx status code
func (o *PostLearningAssessmentsScoringServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assessments scoring service unavailable response has a 4xx status code
func (o *PostLearningAssessmentsScoringServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post learning assessments scoring service unavailable response has a 5xx status code
func (o *PostLearningAssessmentsScoringServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post learning assessments scoring service unavailable response a status code equal to that given
func (o *PostLearningAssessmentsScoringServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostLearningAssessmentsScoringServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLearningAssessmentsScoringServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLearningAssessmentsScoringServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssessmentsScoringServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssessmentsScoringGatewayTimeout creates a PostLearningAssessmentsScoringGatewayTimeout with default headers values
func NewPostLearningAssessmentsScoringGatewayTimeout() *PostLearningAssessmentsScoringGatewayTimeout {
	return &PostLearningAssessmentsScoringGatewayTimeout{}
}

/*
PostLearningAssessmentsScoringGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostLearningAssessmentsScoringGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post learning assessments scoring gateway timeout response has a 2xx status code
func (o *PostLearningAssessmentsScoringGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post learning assessments scoring gateway timeout response has a 3xx status code
func (o *PostLearningAssessmentsScoringGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post learning assessments scoring gateway timeout response has a 4xx status code
func (o *PostLearningAssessmentsScoringGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post learning assessments scoring gateway timeout response has a 5xx status code
func (o *PostLearningAssessmentsScoringGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post learning assessments scoring gateway timeout response a status code equal to that given
func (o *PostLearningAssessmentsScoringGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostLearningAssessmentsScoringGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLearningAssessmentsScoringGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/learning/assessments/scoring][%d] postLearningAssessmentsScoringGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLearningAssessmentsScoringGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssessmentsScoringGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
