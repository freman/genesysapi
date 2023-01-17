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

// GetLearningModuleRuleReader is a Reader for the GetLearningModuleRule structure.
type GetLearningModuleRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearningModuleRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearningModuleRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearningModuleRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLearningModuleRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearningModuleRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearningModuleRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLearningModuleRuleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLearningModuleRuleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLearningModuleRuleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLearningModuleRuleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLearningModuleRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLearningModuleRuleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLearningModuleRuleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearningModuleRuleOK creates a GetLearningModuleRuleOK with default headers values
func NewGetLearningModuleRuleOK() *GetLearningModuleRuleOK {
	return &GetLearningModuleRuleOK{}
}

/*
GetLearningModuleRuleOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLearningModuleRuleOK struct {
	Payload *models.LearningModuleRule
}

// IsSuccess returns true when this get learning module rule o k response has a 2xx status code
func (o *GetLearningModuleRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get learning module rule o k response has a 3xx status code
func (o *GetLearningModuleRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning module rule o k response has a 4xx status code
func (o *GetLearningModuleRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get learning module rule o k response has a 5xx status code
func (o *GetLearningModuleRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning module rule o k response a status code equal to that given
func (o *GetLearningModuleRuleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLearningModuleRuleOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleOK  %+v", 200, o.Payload)
}

func (o *GetLearningModuleRuleOK) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleOK  %+v", 200, o.Payload)
}

func (o *GetLearningModuleRuleOK) GetPayload() *models.LearningModuleRule {
	return o.Payload
}

func (o *GetLearningModuleRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LearningModuleRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleRuleBadRequest creates a GetLearningModuleRuleBadRequest with default headers values
func NewGetLearningModuleRuleBadRequest() *GetLearningModuleRuleBadRequest {
	return &GetLearningModuleRuleBadRequest{}
}

/*
GetLearningModuleRuleBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLearningModuleRuleBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning module rule bad request response has a 2xx status code
func (o *GetLearningModuleRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning module rule bad request response has a 3xx status code
func (o *GetLearningModuleRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning module rule bad request response has a 4xx status code
func (o *GetLearningModuleRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning module rule bad request response has a 5xx status code
func (o *GetLearningModuleRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning module rule bad request response a status code equal to that given
func (o *GetLearningModuleRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLearningModuleRuleBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearningModuleRuleBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearningModuleRuleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleRuleUnauthorized creates a GetLearningModuleRuleUnauthorized with default headers values
func NewGetLearningModuleRuleUnauthorized() *GetLearningModuleRuleUnauthorized {
	return &GetLearningModuleRuleUnauthorized{}
}

/*
GetLearningModuleRuleUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLearningModuleRuleUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning module rule unauthorized response has a 2xx status code
func (o *GetLearningModuleRuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning module rule unauthorized response has a 3xx status code
func (o *GetLearningModuleRuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning module rule unauthorized response has a 4xx status code
func (o *GetLearningModuleRuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning module rule unauthorized response has a 5xx status code
func (o *GetLearningModuleRuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning module rule unauthorized response a status code equal to that given
func (o *GetLearningModuleRuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLearningModuleRuleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLearningModuleRuleUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLearningModuleRuleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleRuleForbidden creates a GetLearningModuleRuleForbidden with default headers values
func NewGetLearningModuleRuleForbidden() *GetLearningModuleRuleForbidden {
	return &GetLearningModuleRuleForbidden{}
}

/*
GetLearningModuleRuleForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetLearningModuleRuleForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning module rule forbidden response has a 2xx status code
func (o *GetLearningModuleRuleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning module rule forbidden response has a 3xx status code
func (o *GetLearningModuleRuleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning module rule forbidden response has a 4xx status code
func (o *GetLearningModuleRuleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning module rule forbidden response has a 5xx status code
func (o *GetLearningModuleRuleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning module rule forbidden response a status code equal to that given
func (o *GetLearningModuleRuleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLearningModuleRuleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleForbidden  %+v", 403, o.Payload)
}

func (o *GetLearningModuleRuleForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleForbidden  %+v", 403, o.Payload)
}

func (o *GetLearningModuleRuleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleRuleNotFound creates a GetLearningModuleRuleNotFound with default headers values
func NewGetLearningModuleRuleNotFound() *GetLearningModuleRuleNotFound {
	return &GetLearningModuleRuleNotFound{}
}

/*
GetLearningModuleRuleNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetLearningModuleRuleNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning module rule not found response has a 2xx status code
func (o *GetLearningModuleRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning module rule not found response has a 3xx status code
func (o *GetLearningModuleRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning module rule not found response has a 4xx status code
func (o *GetLearningModuleRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning module rule not found response has a 5xx status code
func (o *GetLearningModuleRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning module rule not found response a status code equal to that given
func (o *GetLearningModuleRuleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLearningModuleRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetLearningModuleRuleNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetLearningModuleRuleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleRuleRequestTimeout creates a GetLearningModuleRuleRequestTimeout with default headers values
func NewGetLearningModuleRuleRequestTimeout() *GetLearningModuleRuleRequestTimeout {
	return &GetLearningModuleRuleRequestTimeout{}
}

/*
GetLearningModuleRuleRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLearningModuleRuleRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning module rule request timeout response has a 2xx status code
func (o *GetLearningModuleRuleRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning module rule request timeout response has a 3xx status code
func (o *GetLearningModuleRuleRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning module rule request timeout response has a 4xx status code
func (o *GetLearningModuleRuleRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning module rule request timeout response has a 5xx status code
func (o *GetLearningModuleRuleRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning module rule request timeout response a status code equal to that given
func (o *GetLearningModuleRuleRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetLearningModuleRuleRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLearningModuleRuleRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLearningModuleRuleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleRuleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleRuleRequestEntityTooLarge creates a GetLearningModuleRuleRequestEntityTooLarge with default headers values
func NewGetLearningModuleRuleRequestEntityTooLarge() *GetLearningModuleRuleRequestEntityTooLarge {
	return &GetLearningModuleRuleRequestEntityTooLarge{}
}

/*
GetLearningModuleRuleRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetLearningModuleRuleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning module rule request entity too large response has a 2xx status code
func (o *GetLearningModuleRuleRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning module rule request entity too large response has a 3xx status code
func (o *GetLearningModuleRuleRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning module rule request entity too large response has a 4xx status code
func (o *GetLearningModuleRuleRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning module rule request entity too large response has a 5xx status code
func (o *GetLearningModuleRuleRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning module rule request entity too large response a status code equal to that given
func (o *GetLearningModuleRuleRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetLearningModuleRuleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLearningModuleRuleRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLearningModuleRuleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleRuleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleRuleUnsupportedMediaType creates a GetLearningModuleRuleUnsupportedMediaType with default headers values
func NewGetLearningModuleRuleUnsupportedMediaType() *GetLearningModuleRuleUnsupportedMediaType {
	return &GetLearningModuleRuleUnsupportedMediaType{}
}

/*
GetLearningModuleRuleUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLearningModuleRuleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning module rule unsupported media type response has a 2xx status code
func (o *GetLearningModuleRuleUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning module rule unsupported media type response has a 3xx status code
func (o *GetLearningModuleRuleUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning module rule unsupported media type response has a 4xx status code
func (o *GetLearningModuleRuleUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning module rule unsupported media type response has a 5xx status code
func (o *GetLearningModuleRuleUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning module rule unsupported media type response a status code equal to that given
func (o *GetLearningModuleRuleUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetLearningModuleRuleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLearningModuleRuleUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLearningModuleRuleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleRuleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleRuleTooManyRequests creates a GetLearningModuleRuleTooManyRequests with default headers values
func NewGetLearningModuleRuleTooManyRequests() *GetLearningModuleRuleTooManyRequests {
	return &GetLearningModuleRuleTooManyRequests{}
}

/*
GetLearningModuleRuleTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLearningModuleRuleTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning module rule too many requests response has a 2xx status code
func (o *GetLearningModuleRuleTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning module rule too many requests response has a 3xx status code
func (o *GetLearningModuleRuleTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning module rule too many requests response has a 4xx status code
func (o *GetLearningModuleRuleTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get learning module rule too many requests response has a 5xx status code
func (o *GetLearningModuleRuleTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get learning module rule too many requests response a status code equal to that given
func (o *GetLearningModuleRuleTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetLearningModuleRuleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLearningModuleRuleTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLearningModuleRuleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleRuleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleRuleInternalServerError creates a GetLearningModuleRuleInternalServerError with default headers values
func NewGetLearningModuleRuleInternalServerError() *GetLearningModuleRuleInternalServerError {
	return &GetLearningModuleRuleInternalServerError{}
}

/*
GetLearningModuleRuleInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLearningModuleRuleInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning module rule internal server error response has a 2xx status code
func (o *GetLearningModuleRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning module rule internal server error response has a 3xx status code
func (o *GetLearningModuleRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning module rule internal server error response has a 4xx status code
func (o *GetLearningModuleRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get learning module rule internal server error response has a 5xx status code
func (o *GetLearningModuleRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get learning module rule internal server error response a status code equal to that given
func (o *GetLearningModuleRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLearningModuleRuleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLearningModuleRuleInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLearningModuleRuleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleRuleServiceUnavailable creates a GetLearningModuleRuleServiceUnavailable with default headers values
func NewGetLearningModuleRuleServiceUnavailable() *GetLearningModuleRuleServiceUnavailable {
	return &GetLearningModuleRuleServiceUnavailable{}
}

/*
GetLearningModuleRuleServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLearningModuleRuleServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning module rule service unavailable response has a 2xx status code
func (o *GetLearningModuleRuleServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning module rule service unavailable response has a 3xx status code
func (o *GetLearningModuleRuleServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning module rule service unavailable response has a 4xx status code
func (o *GetLearningModuleRuleServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get learning module rule service unavailable response has a 5xx status code
func (o *GetLearningModuleRuleServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get learning module rule service unavailable response a status code equal to that given
func (o *GetLearningModuleRuleServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetLearningModuleRuleServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLearningModuleRuleServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLearningModuleRuleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleRuleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleRuleGatewayTimeout creates a GetLearningModuleRuleGatewayTimeout with default headers values
func NewGetLearningModuleRuleGatewayTimeout() *GetLearningModuleRuleGatewayTimeout {
	return &GetLearningModuleRuleGatewayTimeout{}
}

/*
GetLearningModuleRuleGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetLearningModuleRuleGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get learning module rule gateway timeout response has a 2xx status code
func (o *GetLearningModuleRuleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get learning module rule gateway timeout response has a 3xx status code
func (o *GetLearningModuleRuleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get learning module rule gateway timeout response has a 4xx status code
func (o *GetLearningModuleRuleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get learning module rule gateway timeout response has a 5xx status code
func (o *GetLearningModuleRuleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get learning module rule gateway timeout response a status code equal to that given
func (o *GetLearningModuleRuleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetLearningModuleRuleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLearningModuleRuleGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/rule][%d] getLearningModuleRuleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLearningModuleRuleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleRuleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
