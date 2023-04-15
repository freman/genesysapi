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

// DeleteQualityFormsSurveyReader is a Reader for the DeleteQualityFormsSurvey structure.
type DeleteQualityFormsSurveyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteQualityFormsSurveyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteQualityFormsSurveyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteQualityFormsSurveyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteQualityFormsSurveyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteQualityFormsSurveyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteQualityFormsSurveyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteQualityFormsSurveyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteQualityFormsSurveyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteQualityFormsSurveyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteQualityFormsSurveyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteQualityFormsSurveyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteQualityFormsSurveyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteQualityFormsSurveyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteQualityFormsSurveyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteQualityFormsSurveyNoContent creates a DeleteQualityFormsSurveyNoContent with default headers values
func NewDeleteQualityFormsSurveyNoContent() *DeleteQualityFormsSurveyNoContent {
	return &DeleteQualityFormsSurveyNoContent{}
}

/*
DeleteQualityFormsSurveyNoContent describes a response with status code 204, with default header values.

Operation was successful.
*/
type DeleteQualityFormsSurveyNoContent struct {
}

// IsSuccess returns true when this delete quality forms survey no content response has a 2xx status code
func (o *DeleteQualityFormsSurveyNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete quality forms survey no content response has a 3xx status code
func (o *DeleteQualityFormsSurveyNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey no content response has a 4xx status code
func (o *DeleteQualityFormsSurveyNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete quality forms survey no content response has a 5xx status code
func (o *DeleteQualityFormsSurveyNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete quality forms survey no content response a status code equal to that given
func (o *DeleteQualityFormsSurveyNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteQualityFormsSurveyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyNoContent ", 204)
}

func (o *DeleteQualityFormsSurveyNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyNoContent ", 204)
}

func (o *DeleteQualityFormsSurveyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteQualityFormsSurveyBadRequest creates a DeleteQualityFormsSurveyBadRequest with default headers values
func NewDeleteQualityFormsSurveyBadRequest() *DeleteQualityFormsSurveyBadRequest {
	return &DeleteQualityFormsSurveyBadRequest{}
}

/*
DeleteQualityFormsSurveyBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteQualityFormsSurveyBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete quality forms survey bad request response has a 2xx status code
func (o *DeleteQualityFormsSurveyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete quality forms survey bad request response has a 3xx status code
func (o *DeleteQualityFormsSurveyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey bad request response has a 4xx status code
func (o *DeleteQualityFormsSurveyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete quality forms survey bad request response has a 5xx status code
func (o *DeleteQualityFormsSurveyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete quality forms survey bad request response a status code equal to that given
func (o *DeleteQualityFormsSurveyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteQualityFormsSurveyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteQualityFormsSurveyBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteQualityFormsSurveyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityFormsSurveyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityFormsSurveyUnauthorized creates a DeleteQualityFormsSurveyUnauthorized with default headers values
func NewDeleteQualityFormsSurveyUnauthorized() *DeleteQualityFormsSurveyUnauthorized {
	return &DeleteQualityFormsSurveyUnauthorized{}
}

/*
DeleteQualityFormsSurveyUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteQualityFormsSurveyUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete quality forms survey unauthorized response has a 2xx status code
func (o *DeleteQualityFormsSurveyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete quality forms survey unauthorized response has a 3xx status code
func (o *DeleteQualityFormsSurveyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey unauthorized response has a 4xx status code
func (o *DeleteQualityFormsSurveyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete quality forms survey unauthorized response has a 5xx status code
func (o *DeleteQualityFormsSurveyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete quality forms survey unauthorized response a status code equal to that given
func (o *DeleteQualityFormsSurveyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteQualityFormsSurveyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteQualityFormsSurveyUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteQualityFormsSurveyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityFormsSurveyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityFormsSurveyForbidden creates a DeleteQualityFormsSurveyForbidden with default headers values
func NewDeleteQualityFormsSurveyForbidden() *DeleteQualityFormsSurveyForbidden {
	return &DeleteQualityFormsSurveyForbidden{}
}

/*
DeleteQualityFormsSurveyForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteQualityFormsSurveyForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete quality forms survey forbidden response has a 2xx status code
func (o *DeleteQualityFormsSurveyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete quality forms survey forbidden response has a 3xx status code
func (o *DeleteQualityFormsSurveyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey forbidden response has a 4xx status code
func (o *DeleteQualityFormsSurveyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete quality forms survey forbidden response has a 5xx status code
func (o *DeleteQualityFormsSurveyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete quality forms survey forbidden response a status code equal to that given
func (o *DeleteQualityFormsSurveyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteQualityFormsSurveyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyForbidden  %+v", 403, o.Payload)
}

func (o *DeleteQualityFormsSurveyForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyForbidden  %+v", 403, o.Payload)
}

func (o *DeleteQualityFormsSurveyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityFormsSurveyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityFormsSurveyNotFound creates a DeleteQualityFormsSurveyNotFound with default headers values
func NewDeleteQualityFormsSurveyNotFound() *DeleteQualityFormsSurveyNotFound {
	return &DeleteQualityFormsSurveyNotFound{}
}

/*
DeleteQualityFormsSurveyNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteQualityFormsSurveyNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete quality forms survey not found response has a 2xx status code
func (o *DeleteQualityFormsSurveyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete quality forms survey not found response has a 3xx status code
func (o *DeleteQualityFormsSurveyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey not found response has a 4xx status code
func (o *DeleteQualityFormsSurveyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete quality forms survey not found response has a 5xx status code
func (o *DeleteQualityFormsSurveyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete quality forms survey not found response a status code equal to that given
func (o *DeleteQualityFormsSurveyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteQualityFormsSurveyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyNotFound  %+v", 404, o.Payload)
}

func (o *DeleteQualityFormsSurveyNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyNotFound  %+v", 404, o.Payload)
}

func (o *DeleteQualityFormsSurveyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityFormsSurveyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityFormsSurveyRequestTimeout creates a DeleteQualityFormsSurveyRequestTimeout with default headers values
func NewDeleteQualityFormsSurveyRequestTimeout() *DeleteQualityFormsSurveyRequestTimeout {
	return &DeleteQualityFormsSurveyRequestTimeout{}
}

/*
DeleteQualityFormsSurveyRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteQualityFormsSurveyRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete quality forms survey request timeout response has a 2xx status code
func (o *DeleteQualityFormsSurveyRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete quality forms survey request timeout response has a 3xx status code
func (o *DeleteQualityFormsSurveyRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey request timeout response has a 4xx status code
func (o *DeleteQualityFormsSurveyRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete quality forms survey request timeout response has a 5xx status code
func (o *DeleteQualityFormsSurveyRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete quality forms survey request timeout response a status code equal to that given
func (o *DeleteQualityFormsSurveyRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteQualityFormsSurveyRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteQualityFormsSurveyRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteQualityFormsSurveyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityFormsSurveyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityFormsSurveyConflict creates a DeleteQualityFormsSurveyConflict with default headers values
func NewDeleteQualityFormsSurveyConflict() *DeleteQualityFormsSurveyConflict {
	return &DeleteQualityFormsSurveyConflict{}
}

/*
DeleteQualityFormsSurveyConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteQualityFormsSurveyConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete quality forms survey conflict response has a 2xx status code
func (o *DeleteQualityFormsSurveyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete quality forms survey conflict response has a 3xx status code
func (o *DeleteQualityFormsSurveyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey conflict response has a 4xx status code
func (o *DeleteQualityFormsSurveyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete quality forms survey conflict response has a 5xx status code
func (o *DeleteQualityFormsSurveyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete quality forms survey conflict response a status code equal to that given
func (o *DeleteQualityFormsSurveyConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteQualityFormsSurveyConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyConflict  %+v", 409, o.Payload)
}

func (o *DeleteQualityFormsSurveyConflict) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyConflict  %+v", 409, o.Payload)
}

func (o *DeleteQualityFormsSurveyConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityFormsSurveyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityFormsSurveyRequestEntityTooLarge creates a DeleteQualityFormsSurveyRequestEntityTooLarge with default headers values
func NewDeleteQualityFormsSurveyRequestEntityTooLarge() *DeleteQualityFormsSurveyRequestEntityTooLarge {
	return &DeleteQualityFormsSurveyRequestEntityTooLarge{}
}

/*
DeleteQualityFormsSurveyRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteQualityFormsSurveyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete quality forms survey request entity too large response has a 2xx status code
func (o *DeleteQualityFormsSurveyRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete quality forms survey request entity too large response has a 3xx status code
func (o *DeleteQualityFormsSurveyRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey request entity too large response has a 4xx status code
func (o *DeleteQualityFormsSurveyRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete quality forms survey request entity too large response has a 5xx status code
func (o *DeleteQualityFormsSurveyRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete quality forms survey request entity too large response a status code equal to that given
func (o *DeleteQualityFormsSurveyRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteQualityFormsSurveyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteQualityFormsSurveyRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteQualityFormsSurveyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityFormsSurveyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityFormsSurveyUnsupportedMediaType creates a DeleteQualityFormsSurveyUnsupportedMediaType with default headers values
func NewDeleteQualityFormsSurveyUnsupportedMediaType() *DeleteQualityFormsSurveyUnsupportedMediaType {
	return &DeleteQualityFormsSurveyUnsupportedMediaType{}
}

/*
DeleteQualityFormsSurveyUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteQualityFormsSurveyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete quality forms survey unsupported media type response has a 2xx status code
func (o *DeleteQualityFormsSurveyUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete quality forms survey unsupported media type response has a 3xx status code
func (o *DeleteQualityFormsSurveyUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey unsupported media type response has a 4xx status code
func (o *DeleteQualityFormsSurveyUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete quality forms survey unsupported media type response has a 5xx status code
func (o *DeleteQualityFormsSurveyUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete quality forms survey unsupported media type response a status code equal to that given
func (o *DeleteQualityFormsSurveyUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteQualityFormsSurveyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteQualityFormsSurveyUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteQualityFormsSurveyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityFormsSurveyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityFormsSurveyTooManyRequests creates a DeleteQualityFormsSurveyTooManyRequests with default headers values
func NewDeleteQualityFormsSurveyTooManyRequests() *DeleteQualityFormsSurveyTooManyRequests {
	return &DeleteQualityFormsSurveyTooManyRequests{}
}

/*
DeleteQualityFormsSurveyTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteQualityFormsSurveyTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete quality forms survey too many requests response has a 2xx status code
func (o *DeleteQualityFormsSurveyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete quality forms survey too many requests response has a 3xx status code
func (o *DeleteQualityFormsSurveyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey too many requests response has a 4xx status code
func (o *DeleteQualityFormsSurveyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete quality forms survey too many requests response has a 5xx status code
func (o *DeleteQualityFormsSurveyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete quality forms survey too many requests response a status code equal to that given
func (o *DeleteQualityFormsSurveyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteQualityFormsSurveyTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteQualityFormsSurveyTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteQualityFormsSurveyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityFormsSurveyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityFormsSurveyInternalServerError creates a DeleteQualityFormsSurveyInternalServerError with default headers values
func NewDeleteQualityFormsSurveyInternalServerError() *DeleteQualityFormsSurveyInternalServerError {
	return &DeleteQualityFormsSurveyInternalServerError{}
}

/*
DeleteQualityFormsSurveyInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteQualityFormsSurveyInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete quality forms survey internal server error response has a 2xx status code
func (o *DeleteQualityFormsSurveyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete quality forms survey internal server error response has a 3xx status code
func (o *DeleteQualityFormsSurveyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey internal server error response has a 4xx status code
func (o *DeleteQualityFormsSurveyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete quality forms survey internal server error response has a 5xx status code
func (o *DeleteQualityFormsSurveyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete quality forms survey internal server error response a status code equal to that given
func (o *DeleteQualityFormsSurveyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteQualityFormsSurveyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteQualityFormsSurveyInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteQualityFormsSurveyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityFormsSurveyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityFormsSurveyServiceUnavailable creates a DeleteQualityFormsSurveyServiceUnavailable with default headers values
func NewDeleteQualityFormsSurveyServiceUnavailable() *DeleteQualityFormsSurveyServiceUnavailable {
	return &DeleteQualityFormsSurveyServiceUnavailable{}
}

/*
DeleteQualityFormsSurveyServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteQualityFormsSurveyServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete quality forms survey service unavailable response has a 2xx status code
func (o *DeleteQualityFormsSurveyServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete quality forms survey service unavailable response has a 3xx status code
func (o *DeleteQualityFormsSurveyServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey service unavailable response has a 4xx status code
func (o *DeleteQualityFormsSurveyServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete quality forms survey service unavailable response has a 5xx status code
func (o *DeleteQualityFormsSurveyServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete quality forms survey service unavailable response a status code equal to that given
func (o *DeleteQualityFormsSurveyServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteQualityFormsSurveyServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteQualityFormsSurveyServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteQualityFormsSurveyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityFormsSurveyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityFormsSurveyGatewayTimeout creates a DeleteQualityFormsSurveyGatewayTimeout with default headers values
func NewDeleteQualityFormsSurveyGatewayTimeout() *DeleteQualityFormsSurveyGatewayTimeout {
	return &DeleteQualityFormsSurveyGatewayTimeout{}
}

/*
DeleteQualityFormsSurveyGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteQualityFormsSurveyGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete quality forms survey gateway timeout response has a 2xx status code
func (o *DeleteQualityFormsSurveyGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete quality forms survey gateway timeout response has a 3xx status code
func (o *DeleteQualityFormsSurveyGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete quality forms survey gateway timeout response has a 4xx status code
func (o *DeleteQualityFormsSurveyGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete quality forms survey gateway timeout response has a 5xx status code
func (o *DeleteQualityFormsSurveyGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete quality forms survey gateway timeout response a status code equal to that given
func (o *DeleteQualityFormsSurveyGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteQualityFormsSurveyGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteQualityFormsSurveyGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/forms/surveys/{formId}][%d] deleteQualityFormsSurveyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteQualityFormsSurveyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityFormsSurveyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
