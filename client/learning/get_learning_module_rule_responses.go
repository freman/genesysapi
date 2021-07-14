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

/*GetLearningModuleRuleOK handles this case with default header values.

successful operation
*/
type GetLearningModuleRuleOK struct {
	Payload *models.LearningModuleRule
}

func (o *GetLearningModuleRuleOK) Error() string {
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

/*GetLearningModuleRuleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLearningModuleRuleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleRuleBadRequest) Error() string {
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

/*GetLearningModuleRuleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLearningModuleRuleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleRuleUnauthorized) Error() string {
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

/*GetLearningModuleRuleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLearningModuleRuleForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleRuleForbidden) Error() string {
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

/*GetLearningModuleRuleNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLearningModuleRuleNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleRuleNotFound) Error() string {
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

/*GetLearningModuleRuleRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLearningModuleRuleRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleRuleRequestTimeout) Error() string {
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

/*GetLearningModuleRuleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLearningModuleRuleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleRuleRequestEntityTooLarge) Error() string {
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

/*GetLearningModuleRuleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLearningModuleRuleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleRuleUnsupportedMediaType) Error() string {
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

/*GetLearningModuleRuleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLearningModuleRuleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleRuleTooManyRequests) Error() string {
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

/*GetLearningModuleRuleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLearningModuleRuleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleRuleInternalServerError) Error() string {
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

/*GetLearningModuleRuleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLearningModuleRuleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleRuleServiceUnavailable) Error() string {
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

/*GetLearningModuleRuleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLearningModuleRuleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleRuleGatewayTimeout) Error() string {
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
