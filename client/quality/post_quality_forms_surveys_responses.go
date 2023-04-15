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

// PostQualityFormsSurveysReader is a Reader for the PostQualityFormsSurveys structure.
type PostQualityFormsSurveysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualityFormsSurveysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualityFormsSurveysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualityFormsSurveysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualityFormsSurveysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualityFormsSurveysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualityFormsSurveysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostQualityFormsSurveysRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostQualityFormsSurveysConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualityFormsSurveysRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualityFormsSurveysUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualityFormsSurveysTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualityFormsSurveysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualityFormsSurveysServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualityFormsSurveysGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualityFormsSurveysOK creates a PostQualityFormsSurveysOK with default headers values
func NewPostQualityFormsSurveysOK() *PostQualityFormsSurveysOK {
	return &PostQualityFormsSurveysOK{}
}

/*
PostQualityFormsSurveysOK describes a response with status code 200, with default header values.

successful operation
*/
type PostQualityFormsSurveysOK struct {
	Payload *models.SurveyForm
}

// IsSuccess returns true when this post quality forms surveys o k response has a 2xx status code
func (o *PostQualityFormsSurveysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post quality forms surveys o k response has a 3xx status code
func (o *PostQualityFormsSurveysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys o k response has a 4xx status code
func (o *PostQualityFormsSurveysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality forms surveys o k response has a 5xx status code
func (o *PostQualityFormsSurveysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms surveys o k response a status code equal to that given
func (o *PostQualityFormsSurveysOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostQualityFormsSurveysOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysOK  %+v", 200, o.Payload)
}

func (o *PostQualityFormsSurveysOK) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysOK  %+v", 200, o.Payload)
}

func (o *PostQualityFormsSurveysOK) GetPayload() *models.SurveyForm {
	return o.Payload
}

func (o *PostQualityFormsSurveysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SurveyForm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysBadRequest creates a PostQualityFormsSurveysBadRequest with default headers values
func NewPostQualityFormsSurveysBadRequest() *PostQualityFormsSurveysBadRequest {
	return &PostQualityFormsSurveysBadRequest{}
}

/*
PostQualityFormsSurveysBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualityFormsSurveysBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms surveys bad request response has a 2xx status code
func (o *PostQualityFormsSurveysBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms surveys bad request response has a 3xx status code
func (o *PostQualityFormsSurveysBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys bad request response has a 4xx status code
func (o *PostQualityFormsSurveysBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms surveys bad request response has a 5xx status code
func (o *PostQualityFormsSurveysBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms surveys bad request response a status code equal to that given
func (o *PostQualityFormsSurveysBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostQualityFormsSurveysBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityFormsSurveysBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityFormsSurveysBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysUnauthorized creates a PostQualityFormsSurveysUnauthorized with default headers values
func NewPostQualityFormsSurveysUnauthorized() *PostQualityFormsSurveysUnauthorized {
	return &PostQualityFormsSurveysUnauthorized{}
}

/*
PostQualityFormsSurveysUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualityFormsSurveysUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms surveys unauthorized response has a 2xx status code
func (o *PostQualityFormsSurveysUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms surveys unauthorized response has a 3xx status code
func (o *PostQualityFormsSurveysUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys unauthorized response has a 4xx status code
func (o *PostQualityFormsSurveysUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms surveys unauthorized response has a 5xx status code
func (o *PostQualityFormsSurveysUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms surveys unauthorized response a status code equal to that given
func (o *PostQualityFormsSurveysUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostQualityFormsSurveysUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityFormsSurveysUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityFormsSurveysUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysForbidden creates a PostQualityFormsSurveysForbidden with default headers values
func NewPostQualityFormsSurveysForbidden() *PostQualityFormsSurveysForbidden {
	return &PostQualityFormsSurveysForbidden{}
}

/*
PostQualityFormsSurveysForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostQualityFormsSurveysForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms surveys forbidden response has a 2xx status code
func (o *PostQualityFormsSurveysForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms surveys forbidden response has a 3xx status code
func (o *PostQualityFormsSurveysForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys forbidden response has a 4xx status code
func (o *PostQualityFormsSurveysForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms surveys forbidden response has a 5xx status code
func (o *PostQualityFormsSurveysForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms surveys forbidden response a status code equal to that given
func (o *PostQualityFormsSurveysForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostQualityFormsSurveysForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityFormsSurveysForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityFormsSurveysForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysNotFound creates a PostQualityFormsSurveysNotFound with default headers values
func NewPostQualityFormsSurveysNotFound() *PostQualityFormsSurveysNotFound {
	return &PostQualityFormsSurveysNotFound{}
}

/*
PostQualityFormsSurveysNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostQualityFormsSurveysNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms surveys not found response has a 2xx status code
func (o *PostQualityFormsSurveysNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms surveys not found response has a 3xx status code
func (o *PostQualityFormsSurveysNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys not found response has a 4xx status code
func (o *PostQualityFormsSurveysNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms surveys not found response has a 5xx status code
func (o *PostQualityFormsSurveysNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms surveys not found response a status code equal to that given
func (o *PostQualityFormsSurveysNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostQualityFormsSurveysNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityFormsSurveysNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityFormsSurveysNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysRequestTimeout creates a PostQualityFormsSurveysRequestTimeout with default headers values
func NewPostQualityFormsSurveysRequestTimeout() *PostQualityFormsSurveysRequestTimeout {
	return &PostQualityFormsSurveysRequestTimeout{}
}

/*
PostQualityFormsSurveysRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostQualityFormsSurveysRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms surveys request timeout response has a 2xx status code
func (o *PostQualityFormsSurveysRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms surveys request timeout response has a 3xx status code
func (o *PostQualityFormsSurveysRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys request timeout response has a 4xx status code
func (o *PostQualityFormsSurveysRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms surveys request timeout response has a 5xx status code
func (o *PostQualityFormsSurveysRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms surveys request timeout response a status code equal to that given
func (o *PostQualityFormsSurveysRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostQualityFormsSurveysRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityFormsSurveysRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityFormsSurveysRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysConflict creates a PostQualityFormsSurveysConflict with default headers values
func NewPostQualityFormsSurveysConflict() *PostQualityFormsSurveysConflict {
	return &PostQualityFormsSurveysConflict{}
}

/*
PostQualityFormsSurveysConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostQualityFormsSurveysConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms surveys conflict response has a 2xx status code
func (o *PostQualityFormsSurveysConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms surveys conflict response has a 3xx status code
func (o *PostQualityFormsSurveysConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys conflict response has a 4xx status code
func (o *PostQualityFormsSurveysConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms surveys conflict response has a 5xx status code
func (o *PostQualityFormsSurveysConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms surveys conflict response a status code equal to that given
func (o *PostQualityFormsSurveysConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostQualityFormsSurveysConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysConflict  %+v", 409, o.Payload)
}

func (o *PostQualityFormsSurveysConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysConflict  %+v", 409, o.Payload)
}

func (o *PostQualityFormsSurveysConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysRequestEntityTooLarge creates a PostQualityFormsSurveysRequestEntityTooLarge with default headers values
func NewPostQualityFormsSurveysRequestEntityTooLarge() *PostQualityFormsSurveysRequestEntityTooLarge {
	return &PostQualityFormsSurveysRequestEntityTooLarge{}
}

/*
PostQualityFormsSurveysRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostQualityFormsSurveysRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms surveys request entity too large response has a 2xx status code
func (o *PostQualityFormsSurveysRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms surveys request entity too large response has a 3xx status code
func (o *PostQualityFormsSurveysRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys request entity too large response has a 4xx status code
func (o *PostQualityFormsSurveysRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms surveys request entity too large response has a 5xx status code
func (o *PostQualityFormsSurveysRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms surveys request entity too large response a status code equal to that given
func (o *PostQualityFormsSurveysRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostQualityFormsSurveysRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityFormsSurveysRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityFormsSurveysRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysUnsupportedMediaType creates a PostQualityFormsSurveysUnsupportedMediaType with default headers values
func NewPostQualityFormsSurveysUnsupportedMediaType() *PostQualityFormsSurveysUnsupportedMediaType {
	return &PostQualityFormsSurveysUnsupportedMediaType{}
}

/*
PostQualityFormsSurveysUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualityFormsSurveysUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms surveys unsupported media type response has a 2xx status code
func (o *PostQualityFormsSurveysUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms surveys unsupported media type response has a 3xx status code
func (o *PostQualityFormsSurveysUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys unsupported media type response has a 4xx status code
func (o *PostQualityFormsSurveysUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms surveys unsupported media type response has a 5xx status code
func (o *PostQualityFormsSurveysUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms surveys unsupported media type response a status code equal to that given
func (o *PostQualityFormsSurveysUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostQualityFormsSurveysUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityFormsSurveysUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityFormsSurveysUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysTooManyRequests creates a PostQualityFormsSurveysTooManyRequests with default headers values
func NewPostQualityFormsSurveysTooManyRequests() *PostQualityFormsSurveysTooManyRequests {
	return &PostQualityFormsSurveysTooManyRequests{}
}

/*
PostQualityFormsSurveysTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostQualityFormsSurveysTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms surveys too many requests response has a 2xx status code
func (o *PostQualityFormsSurveysTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms surveys too many requests response has a 3xx status code
func (o *PostQualityFormsSurveysTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys too many requests response has a 4xx status code
func (o *PostQualityFormsSurveysTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality forms surveys too many requests response has a 5xx status code
func (o *PostQualityFormsSurveysTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality forms surveys too many requests response a status code equal to that given
func (o *PostQualityFormsSurveysTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostQualityFormsSurveysTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityFormsSurveysTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityFormsSurveysTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysInternalServerError creates a PostQualityFormsSurveysInternalServerError with default headers values
func NewPostQualityFormsSurveysInternalServerError() *PostQualityFormsSurveysInternalServerError {
	return &PostQualityFormsSurveysInternalServerError{}
}

/*
PostQualityFormsSurveysInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualityFormsSurveysInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms surveys internal server error response has a 2xx status code
func (o *PostQualityFormsSurveysInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms surveys internal server error response has a 3xx status code
func (o *PostQualityFormsSurveysInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys internal server error response has a 4xx status code
func (o *PostQualityFormsSurveysInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality forms surveys internal server error response has a 5xx status code
func (o *PostQualityFormsSurveysInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality forms surveys internal server error response a status code equal to that given
func (o *PostQualityFormsSurveysInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostQualityFormsSurveysInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityFormsSurveysInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityFormsSurveysInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysServiceUnavailable creates a PostQualityFormsSurveysServiceUnavailable with default headers values
func NewPostQualityFormsSurveysServiceUnavailable() *PostQualityFormsSurveysServiceUnavailable {
	return &PostQualityFormsSurveysServiceUnavailable{}
}

/*
PostQualityFormsSurveysServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualityFormsSurveysServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms surveys service unavailable response has a 2xx status code
func (o *PostQualityFormsSurveysServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms surveys service unavailable response has a 3xx status code
func (o *PostQualityFormsSurveysServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys service unavailable response has a 4xx status code
func (o *PostQualityFormsSurveysServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality forms surveys service unavailable response has a 5xx status code
func (o *PostQualityFormsSurveysServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality forms surveys service unavailable response a status code equal to that given
func (o *PostQualityFormsSurveysServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostQualityFormsSurveysServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityFormsSurveysServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityFormsSurveysServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysGatewayTimeout creates a PostQualityFormsSurveysGatewayTimeout with default headers values
func NewPostQualityFormsSurveysGatewayTimeout() *PostQualityFormsSurveysGatewayTimeout {
	return &PostQualityFormsSurveysGatewayTimeout{}
}

/*
PostQualityFormsSurveysGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostQualityFormsSurveysGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality forms surveys gateway timeout response has a 2xx status code
func (o *PostQualityFormsSurveysGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality forms surveys gateway timeout response has a 3xx status code
func (o *PostQualityFormsSurveysGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality forms surveys gateway timeout response has a 4xx status code
func (o *PostQualityFormsSurveysGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality forms surveys gateway timeout response has a 5xx status code
func (o *PostQualityFormsSurveysGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality forms surveys gateway timeout response a status code equal to that given
func (o *PostQualityFormsSurveysGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostQualityFormsSurveysGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityFormsSurveysGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityFormsSurveysGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
