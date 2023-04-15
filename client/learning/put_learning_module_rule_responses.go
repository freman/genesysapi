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

// PutLearningModuleRuleReader is a Reader for the PutLearningModuleRule structure.
type PutLearningModuleRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLearningModuleRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutLearningModuleRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutLearningModuleRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutLearningModuleRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutLearningModuleRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutLearningModuleRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutLearningModuleRuleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutLearningModuleRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutLearningModuleRuleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutLearningModuleRuleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutLearningModuleRuleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutLearningModuleRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutLearningModuleRuleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutLearningModuleRuleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutLearningModuleRuleOK creates a PutLearningModuleRuleOK with default headers values
func NewPutLearningModuleRuleOK() *PutLearningModuleRuleOK {
	return &PutLearningModuleRuleOK{}
}

/*
PutLearningModuleRuleOK describes a response with status code 200, with default header values.

successful operation
*/
type PutLearningModuleRuleOK struct {
	Payload *models.LearningModuleRule
}

// IsSuccess returns true when this put learning module rule o k response has a 2xx status code
func (o *PutLearningModuleRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put learning module rule o k response has a 3xx status code
func (o *PutLearningModuleRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule o k response has a 4xx status code
func (o *PutLearningModuleRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put learning module rule o k response has a 5xx status code
func (o *PutLearningModuleRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put learning module rule o k response a status code equal to that given
func (o *PutLearningModuleRuleOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutLearningModuleRuleOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleOK  %+v", 200, o.Payload)
}

func (o *PutLearningModuleRuleOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleOK  %+v", 200, o.Payload)
}

func (o *PutLearningModuleRuleOK) GetPayload() *models.LearningModuleRule {
	return o.Payload
}

func (o *PutLearningModuleRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LearningModuleRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearningModuleRuleBadRequest creates a PutLearningModuleRuleBadRequest with default headers values
func NewPutLearningModuleRuleBadRequest() *PutLearningModuleRuleBadRequest {
	return &PutLearningModuleRuleBadRequest{}
}

/*
PutLearningModuleRuleBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutLearningModuleRuleBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put learning module rule bad request response has a 2xx status code
func (o *PutLearningModuleRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put learning module rule bad request response has a 3xx status code
func (o *PutLearningModuleRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule bad request response has a 4xx status code
func (o *PutLearningModuleRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put learning module rule bad request response has a 5xx status code
func (o *PutLearningModuleRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put learning module rule bad request response a status code equal to that given
func (o *PutLearningModuleRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutLearningModuleRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleBadRequest  %+v", 400, o.Payload)
}

func (o *PutLearningModuleRuleBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleBadRequest  %+v", 400, o.Payload)
}

func (o *PutLearningModuleRuleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLearningModuleRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearningModuleRuleUnauthorized creates a PutLearningModuleRuleUnauthorized with default headers values
func NewPutLearningModuleRuleUnauthorized() *PutLearningModuleRuleUnauthorized {
	return &PutLearningModuleRuleUnauthorized{}
}

/*
PutLearningModuleRuleUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutLearningModuleRuleUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put learning module rule unauthorized response has a 2xx status code
func (o *PutLearningModuleRuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put learning module rule unauthorized response has a 3xx status code
func (o *PutLearningModuleRuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule unauthorized response has a 4xx status code
func (o *PutLearningModuleRuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put learning module rule unauthorized response has a 5xx status code
func (o *PutLearningModuleRuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put learning module rule unauthorized response a status code equal to that given
func (o *PutLearningModuleRuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutLearningModuleRuleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *PutLearningModuleRuleUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *PutLearningModuleRuleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLearningModuleRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearningModuleRuleForbidden creates a PutLearningModuleRuleForbidden with default headers values
func NewPutLearningModuleRuleForbidden() *PutLearningModuleRuleForbidden {
	return &PutLearningModuleRuleForbidden{}
}

/*
PutLearningModuleRuleForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutLearningModuleRuleForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put learning module rule forbidden response has a 2xx status code
func (o *PutLearningModuleRuleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put learning module rule forbidden response has a 3xx status code
func (o *PutLearningModuleRuleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule forbidden response has a 4xx status code
func (o *PutLearningModuleRuleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put learning module rule forbidden response has a 5xx status code
func (o *PutLearningModuleRuleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put learning module rule forbidden response a status code equal to that given
func (o *PutLearningModuleRuleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutLearningModuleRuleForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleForbidden  %+v", 403, o.Payload)
}

func (o *PutLearningModuleRuleForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleForbidden  %+v", 403, o.Payload)
}

func (o *PutLearningModuleRuleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLearningModuleRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearningModuleRuleNotFound creates a PutLearningModuleRuleNotFound with default headers values
func NewPutLearningModuleRuleNotFound() *PutLearningModuleRuleNotFound {
	return &PutLearningModuleRuleNotFound{}
}

/*
PutLearningModuleRuleNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutLearningModuleRuleNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put learning module rule not found response has a 2xx status code
func (o *PutLearningModuleRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put learning module rule not found response has a 3xx status code
func (o *PutLearningModuleRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule not found response has a 4xx status code
func (o *PutLearningModuleRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put learning module rule not found response has a 5xx status code
func (o *PutLearningModuleRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put learning module rule not found response a status code equal to that given
func (o *PutLearningModuleRuleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutLearningModuleRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleNotFound  %+v", 404, o.Payload)
}

func (o *PutLearningModuleRuleNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleNotFound  %+v", 404, o.Payload)
}

func (o *PutLearningModuleRuleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLearningModuleRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearningModuleRuleRequestTimeout creates a PutLearningModuleRuleRequestTimeout with default headers values
func NewPutLearningModuleRuleRequestTimeout() *PutLearningModuleRuleRequestTimeout {
	return &PutLearningModuleRuleRequestTimeout{}
}

/*
PutLearningModuleRuleRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutLearningModuleRuleRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put learning module rule request timeout response has a 2xx status code
func (o *PutLearningModuleRuleRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put learning module rule request timeout response has a 3xx status code
func (o *PutLearningModuleRuleRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule request timeout response has a 4xx status code
func (o *PutLearningModuleRuleRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put learning module rule request timeout response has a 5xx status code
func (o *PutLearningModuleRuleRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put learning module rule request timeout response a status code equal to that given
func (o *PutLearningModuleRuleRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutLearningModuleRuleRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutLearningModuleRuleRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutLearningModuleRuleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLearningModuleRuleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearningModuleRuleConflict creates a PutLearningModuleRuleConflict with default headers values
func NewPutLearningModuleRuleConflict() *PutLearningModuleRuleConflict {
	return &PutLearningModuleRuleConflict{}
}

/*
PutLearningModuleRuleConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutLearningModuleRuleConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put learning module rule conflict response has a 2xx status code
func (o *PutLearningModuleRuleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put learning module rule conflict response has a 3xx status code
func (o *PutLearningModuleRuleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule conflict response has a 4xx status code
func (o *PutLearningModuleRuleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put learning module rule conflict response has a 5xx status code
func (o *PutLearningModuleRuleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put learning module rule conflict response a status code equal to that given
func (o *PutLearningModuleRuleConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PutLearningModuleRuleConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleConflict  %+v", 409, o.Payload)
}

func (o *PutLearningModuleRuleConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleConflict  %+v", 409, o.Payload)
}

func (o *PutLearningModuleRuleConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLearningModuleRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearningModuleRuleRequestEntityTooLarge creates a PutLearningModuleRuleRequestEntityTooLarge with default headers values
func NewPutLearningModuleRuleRequestEntityTooLarge() *PutLearningModuleRuleRequestEntityTooLarge {
	return &PutLearningModuleRuleRequestEntityTooLarge{}
}

/*
PutLearningModuleRuleRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutLearningModuleRuleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put learning module rule request entity too large response has a 2xx status code
func (o *PutLearningModuleRuleRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put learning module rule request entity too large response has a 3xx status code
func (o *PutLearningModuleRuleRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule request entity too large response has a 4xx status code
func (o *PutLearningModuleRuleRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put learning module rule request entity too large response has a 5xx status code
func (o *PutLearningModuleRuleRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put learning module rule request entity too large response a status code equal to that given
func (o *PutLearningModuleRuleRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutLearningModuleRuleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutLearningModuleRuleRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutLearningModuleRuleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLearningModuleRuleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearningModuleRuleUnsupportedMediaType creates a PutLearningModuleRuleUnsupportedMediaType with default headers values
func NewPutLearningModuleRuleUnsupportedMediaType() *PutLearningModuleRuleUnsupportedMediaType {
	return &PutLearningModuleRuleUnsupportedMediaType{}
}

/*
PutLearningModuleRuleUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutLearningModuleRuleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put learning module rule unsupported media type response has a 2xx status code
func (o *PutLearningModuleRuleUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put learning module rule unsupported media type response has a 3xx status code
func (o *PutLearningModuleRuleUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule unsupported media type response has a 4xx status code
func (o *PutLearningModuleRuleUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put learning module rule unsupported media type response has a 5xx status code
func (o *PutLearningModuleRuleUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put learning module rule unsupported media type response a status code equal to that given
func (o *PutLearningModuleRuleUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutLearningModuleRuleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutLearningModuleRuleUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutLearningModuleRuleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLearningModuleRuleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearningModuleRuleTooManyRequests creates a PutLearningModuleRuleTooManyRequests with default headers values
func NewPutLearningModuleRuleTooManyRequests() *PutLearningModuleRuleTooManyRequests {
	return &PutLearningModuleRuleTooManyRequests{}
}

/*
PutLearningModuleRuleTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutLearningModuleRuleTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put learning module rule too many requests response has a 2xx status code
func (o *PutLearningModuleRuleTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put learning module rule too many requests response has a 3xx status code
func (o *PutLearningModuleRuleTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule too many requests response has a 4xx status code
func (o *PutLearningModuleRuleTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put learning module rule too many requests response has a 5xx status code
func (o *PutLearningModuleRuleTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put learning module rule too many requests response a status code equal to that given
func (o *PutLearningModuleRuleTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutLearningModuleRuleTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutLearningModuleRuleTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutLearningModuleRuleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLearningModuleRuleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearningModuleRuleInternalServerError creates a PutLearningModuleRuleInternalServerError with default headers values
func NewPutLearningModuleRuleInternalServerError() *PutLearningModuleRuleInternalServerError {
	return &PutLearningModuleRuleInternalServerError{}
}

/*
PutLearningModuleRuleInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutLearningModuleRuleInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put learning module rule internal server error response has a 2xx status code
func (o *PutLearningModuleRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put learning module rule internal server error response has a 3xx status code
func (o *PutLearningModuleRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule internal server error response has a 4xx status code
func (o *PutLearningModuleRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put learning module rule internal server error response has a 5xx status code
func (o *PutLearningModuleRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put learning module rule internal server error response a status code equal to that given
func (o *PutLearningModuleRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutLearningModuleRuleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *PutLearningModuleRuleInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *PutLearningModuleRuleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLearningModuleRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearningModuleRuleServiceUnavailable creates a PutLearningModuleRuleServiceUnavailable with default headers values
func NewPutLearningModuleRuleServiceUnavailable() *PutLearningModuleRuleServiceUnavailable {
	return &PutLearningModuleRuleServiceUnavailable{}
}

/*
PutLearningModuleRuleServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutLearningModuleRuleServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put learning module rule service unavailable response has a 2xx status code
func (o *PutLearningModuleRuleServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put learning module rule service unavailable response has a 3xx status code
func (o *PutLearningModuleRuleServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule service unavailable response has a 4xx status code
func (o *PutLearningModuleRuleServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put learning module rule service unavailable response has a 5xx status code
func (o *PutLearningModuleRuleServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put learning module rule service unavailable response a status code equal to that given
func (o *PutLearningModuleRuleServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutLearningModuleRuleServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutLearningModuleRuleServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutLearningModuleRuleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLearningModuleRuleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearningModuleRuleGatewayTimeout creates a PutLearningModuleRuleGatewayTimeout with default headers values
func NewPutLearningModuleRuleGatewayTimeout() *PutLearningModuleRuleGatewayTimeout {
	return &PutLearningModuleRuleGatewayTimeout{}
}

/*
PutLearningModuleRuleGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutLearningModuleRuleGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put learning module rule gateway timeout response has a 2xx status code
func (o *PutLearningModuleRuleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put learning module rule gateway timeout response has a 3xx status code
func (o *PutLearningModuleRuleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put learning module rule gateway timeout response has a 4xx status code
func (o *PutLearningModuleRuleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put learning module rule gateway timeout response has a 5xx status code
func (o *PutLearningModuleRuleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put learning module rule gateway timeout response a status code equal to that given
func (o *PutLearningModuleRuleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutLearningModuleRuleGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutLearningModuleRuleGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/learning/modules/{moduleId}/rule][%d] putLearningModuleRuleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutLearningModuleRuleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLearningModuleRuleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
