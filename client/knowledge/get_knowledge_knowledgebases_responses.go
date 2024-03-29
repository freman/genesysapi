// Code generated by go-swagger; DO NOT EDIT.

package knowledge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetKnowledgeKnowledgebasesReader is a Reader for the GetKnowledgeKnowledgebases structure.
type GetKnowledgeKnowledgebasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeKnowledgebasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeKnowledgebasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeKnowledgebasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeKnowledgebasesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeKnowledgebasesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeKnowledgebasesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeKnowledgebasesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeKnowledgebasesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeKnowledgebasesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeKnowledgebasesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeKnowledgebasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeKnowledgebasesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeKnowledgebasesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeKnowledgebasesOK creates a GetKnowledgeKnowledgebasesOK with default headers values
func NewGetKnowledgeKnowledgebasesOK() *GetKnowledgeKnowledgebasesOK {
	return &GetKnowledgeKnowledgebasesOK{}
}

/*
GetKnowledgeKnowledgebasesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetKnowledgeKnowledgebasesOK struct {
	Payload *models.KnowledgeBaseListing
}

// IsSuccess returns true when this get knowledge knowledgebases o k response has a 2xx status code
func (o *GetKnowledgeKnowledgebasesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get knowledge knowledgebases o k response has a 3xx status code
func (o *GetKnowledgeKnowledgebasesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebases o k response has a 4xx status code
func (o *GetKnowledgeKnowledgebasesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebases o k response has a 5xx status code
func (o *GetKnowledgeKnowledgebasesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebases o k response a status code equal to that given
func (o *GetKnowledgeKnowledgebasesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKnowledgeKnowledgebasesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesOK) GetPayload() *models.KnowledgeBaseListing {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeBaseListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebasesBadRequest creates a GetKnowledgeKnowledgebasesBadRequest with default headers values
func NewGetKnowledgeKnowledgebasesBadRequest() *GetKnowledgeKnowledgebasesBadRequest {
	return &GetKnowledgeKnowledgebasesBadRequest{}
}

/*
GetKnowledgeKnowledgebasesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebasesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebases bad request response has a 2xx status code
func (o *GetKnowledgeKnowledgebasesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebases bad request response has a 3xx status code
func (o *GetKnowledgeKnowledgebasesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebases bad request response has a 4xx status code
func (o *GetKnowledgeKnowledgebasesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebases bad request response has a 5xx status code
func (o *GetKnowledgeKnowledgebasesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebases bad request response a status code equal to that given
func (o *GetKnowledgeKnowledgebasesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKnowledgeKnowledgebasesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebasesUnauthorized creates a GetKnowledgeKnowledgebasesUnauthorized with default headers values
func NewGetKnowledgeKnowledgebasesUnauthorized() *GetKnowledgeKnowledgebasesUnauthorized {
	return &GetKnowledgeKnowledgebasesUnauthorized{}
}

/*
GetKnowledgeKnowledgebasesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebasesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebases unauthorized response has a 2xx status code
func (o *GetKnowledgeKnowledgebasesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebases unauthorized response has a 3xx status code
func (o *GetKnowledgeKnowledgebasesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebases unauthorized response has a 4xx status code
func (o *GetKnowledgeKnowledgebasesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebases unauthorized response has a 5xx status code
func (o *GetKnowledgeKnowledgebasesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebases unauthorized response a status code equal to that given
func (o *GetKnowledgeKnowledgebasesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKnowledgeKnowledgebasesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebasesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebasesForbidden creates a GetKnowledgeKnowledgebasesForbidden with default headers values
func NewGetKnowledgeKnowledgebasesForbidden() *GetKnowledgeKnowledgebasesForbidden {
	return &GetKnowledgeKnowledgebasesForbidden{}
}

/*
GetKnowledgeKnowledgebasesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebasesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebases forbidden response has a 2xx status code
func (o *GetKnowledgeKnowledgebasesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebases forbidden response has a 3xx status code
func (o *GetKnowledgeKnowledgebasesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebases forbidden response has a 4xx status code
func (o *GetKnowledgeKnowledgebasesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebases forbidden response has a 5xx status code
func (o *GetKnowledgeKnowledgebasesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebases forbidden response a status code equal to that given
func (o *GetKnowledgeKnowledgebasesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKnowledgeKnowledgebasesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebasesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebasesNotFound creates a GetKnowledgeKnowledgebasesNotFound with default headers values
func NewGetKnowledgeKnowledgebasesNotFound() *GetKnowledgeKnowledgebasesNotFound {
	return &GetKnowledgeKnowledgebasesNotFound{}
}

/*
GetKnowledgeKnowledgebasesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebasesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebases not found response has a 2xx status code
func (o *GetKnowledgeKnowledgebasesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebases not found response has a 3xx status code
func (o *GetKnowledgeKnowledgebasesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebases not found response has a 4xx status code
func (o *GetKnowledgeKnowledgebasesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebases not found response has a 5xx status code
func (o *GetKnowledgeKnowledgebasesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebases not found response a status code equal to that given
func (o *GetKnowledgeKnowledgebasesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKnowledgeKnowledgebasesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebasesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebasesRequestTimeout creates a GetKnowledgeKnowledgebasesRequestTimeout with default headers values
func NewGetKnowledgeKnowledgebasesRequestTimeout() *GetKnowledgeKnowledgebasesRequestTimeout {
	return &GetKnowledgeKnowledgebasesRequestTimeout{}
}

/*
GetKnowledgeKnowledgebasesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeKnowledgebasesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebases request timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebasesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebases request timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebasesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebases request timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebasesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebases request timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebasesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebases request timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebasesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetKnowledgeKnowledgebasesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebasesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebasesRequestEntityTooLarge creates a GetKnowledgeKnowledgebasesRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebasesRequestEntityTooLarge() *GetKnowledgeKnowledgebasesRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebasesRequestEntityTooLarge{}
}

/*
GetKnowledgeKnowledgebasesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetKnowledgeKnowledgebasesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebases request entity too large response has a 2xx status code
func (o *GetKnowledgeKnowledgebasesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebases request entity too large response has a 3xx status code
func (o *GetKnowledgeKnowledgebasesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebases request entity too large response has a 4xx status code
func (o *GetKnowledgeKnowledgebasesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebases request entity too large response has a 5xx status code
func (o *GetKnowledgeKnowledgebasesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebases request entity too large response a status code equal to that given
func (o *GetKnowledgeKnowledgebasesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetKnowledgeKnowledgebasesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebasesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebasesUnsupportedMediaType creates a GetKnowledgeKnowledgebasesUnsupportedMediaType with default headers values
func NewGetKnowledgeKnowledgebasesUnsupportedMediaType() *GetKnowledgeKnowledgebasesUnsupportedMediaType {
	return &GetKnowledgeKnowledgebasesUnsupportedMediaType{}
}

/*
GetKnowledgeKnowledgebasesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebasesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebases unsupported media type response has a 2xx status code
func (o *GetKnowledgeKnowledgebasesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebases unsupported media type response has a 3xx status code
func (o *GetKnowledgeKnowledgebasesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebases unsupported media type response has a 4xx status code
func (o *GetKnowledgeKnowledgebasesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebases unsupported media type response has a 5xx status code
func (o *GetKnowledgeKnowledgebasesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebases unsupported media type response a status code equal to that given
func (o *GetKnowledgeKnowledgebasesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetKnowledgeKnowledgebasesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebasesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebasesTooManyRequests creates a GetKnowledgeKnowledgebasesTooManyRequests with default headers values
func NewGetKnowledgeKnowledgebasesTooManyRequests() *GetKnowledgeKnowledgebasesTooManyRequests {
	return &GetKnowledgeKnowledgebasesTooManyRequests{}
}

/*
GetKnowledgeKnowledgebasesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeKnowledgebasesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebases too many requests response has a 2xx status code
func (o *GetKnowledgeKnowledgebasesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebases too many requests response has a 3xx status code
func (o *GetKnowledgeKnowledgebasesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebases too many requests response has a 4xx status code
func (o *GetKnowledgeKnowledgebasesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebases too many requests response has a 5xx status code
func (o *GetKnowledgeKnowledgebasesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebases too many requests response a status code equal to that given
func (o *GetKnowledgeKnowledgebasesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetKnowledgeKnowledgebasesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebasesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebasesInternalServerError creates a GetKnowledgeKnowledgebasesInternalServerError with default headers values
func NewGetKnowledgeKnowledgebasesInternalServerError() *GetKnowledgeKnowledgebasesInternalServerError {
	return &GetKnowledgeKnowledgebasesInternalServerError{}
}

/*
GetKnowledgeKnowledgebasesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebasesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebases internal server error response has a 2xx status code
func (o *GetKnowledgeKnowledgebasesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebases internal server error response has a 3xx status code
func (o *GetKnowledgeKnowledgebasesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebases internal server error response has a 4xx status code
func (o *GetKnowledgeKnowledgebasesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebases internal server error response has a 5xx status code
func (o *GetKnowledgeKnowledgebasesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebases internal server error response a status code equal to that given
func (o *GetKnowledgeKnowledgebasesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKnowledgeKnowledgebasesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebasesServiceUnavailable creates a GetKnowledgeKnowledgebasesServiceUnavailable with default headers values
func NewGetKnowledgeKnowledgebasesServiceUnavailable() *GetKnowledgeKnowledgebasesServiceUnavailable {
	return &GetKnowledgeKnowledgebasesServiceUnavailable{}
}

/*
GetKnowledgeKnowledgebasesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebasesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebases service unavailable response has a 2xx status code
func (o *GetKnowledgeKnowledgebasesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebases service unavailable response has a 3xx status code
func (o *GetKnowledgeKnowledgebasesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebases service unavailable response has a 4xx status code
func (o *GetKnowledgeKnowledgebasesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebases service unavailable response has a 5xx status code
func (o *GetKnowledgeKnowledgebasesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebases service unavailable response a status code equal to that given
func (o *GetKnowledgeKnowledgebasesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetKnowledgeKnowledgebasesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebasesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebasesGatewayTimeout creates a GetKnowledgeKnowledgebasesGatewayTimeout with default headers values
func NewGetKnowledgeKnowledgebasesGatewayTimeout() *GetKnowledgeKnowledgebasesGatewayTimeout {
	return &GetKnowledgeKnowledgebasesGatewayTimeout{}
}

/*
GetKnowledgeKnowledgebasesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebasesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebases gateway timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebasesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebases gateway timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebasesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebases gateway timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebasesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebases gateway timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebasesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebases gateway timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebasesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetKnowledgeKnowledgebasesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases][%d] getKnowledgeKnowledgebasesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebasesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebasesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
