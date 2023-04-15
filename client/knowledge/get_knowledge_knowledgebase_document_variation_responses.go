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

// GetKnowledgeKnowledgebaseDocumentVariationReader is a Reader for the GetKnowledgeKnowledgebaseDocumentVariation structure.
type GetKnowledgeKnowledgebaseDocumentVariationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeKnowledgebaseDocumentVariationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeKnowledgebaseDocumentVariationOK creates a GetKnowledgeKnowledgebaseDocumentVariationOK with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationOK() *GetKnowledgeKnowledgebaseDocumentVariationOK {
	return &GetKnowledgeKnowledgebaseDocumentVariationOK{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetKnowledgeKnowledgebaseDocumentVariationOK struct {
	Payload *models.DocumentVariation
}

// IsSuccess returns true when this get knowledge knowledgebase document variation o k response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get knowledge knowledgebase document variation o k response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variation o k response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document variation o k response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variation o k response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationOK) GetPayload() *models.DocumentVariation {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentVariation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationBadRequest creates a GetKnowledgeKnowledgebaseDocumentVariationBadRequest with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationBadRequest() *GetKnowledgeKnowledgebaseDocumentVariationBadRequest {
	return &GetKnowledgeKnowledgebaseDocumentVariationBadRequest{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebaseDocumentVariationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variation bad request response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variation bad request response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variation bad request response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variation bad request response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variation bad request response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationUnauthorized creates a GetKnowledgeKnowledgebaseDocumentVariationUnauthorized with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationUnauthorized() *GetKnowledgeKnowledgebaseDocumentVariationUnauthorized {
	return &GetKnowledgeKnowledgebaseDocumentVariationUnauthorized{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebaseDocumentVariationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variation unauthorized response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variation unauthorized response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variation unauthorized response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variation unauthorized response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variation unauthorized response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationForbidden creates a GetKnowledgeKnowledgebaseDocumentVariationForbidden with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationForbidden() *GetKnowledgeKnowledgebaseDocumentVariationForbidden {
	return &GetKnowledgeKnowledgebaseDocumentVariationForbidden{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebaseDocumentVariationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variation forbidden response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variation forbidden response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variation forbidden response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variation forbidden response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variation forbidden response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationNotFound creates a GetKnowledgeKnowledgebaseDocumentVariationNotFound with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationNotFound() *GetKnowledgeKnowledgebaseDocumentVariationNotFound {
	return &GetKnowledgeKnowledgebaseDocumentVariationNotFound{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebaseDocumentVariationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variation not found response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variation not found response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variation not found response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variation not found response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variation not found response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationRequestTimeout creates a GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationRequestTimeout() *GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout {
	return &GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variation request timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variation request timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variation request timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variation request timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variation request timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge creates a GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge() *GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variation request entity too large response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variation request entity too large response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variation request entity too large response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variation request entity too large response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variation request entity too large response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType creates a GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType() *GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType {
	return &GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variation unsupported media type response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variation unsupported media type response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variation unsupported media type response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variation unsupported media type response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variation unsupported media type response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationTooManyRequests creates a GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationTooManyRequests() *GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests {
	return &GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variation too many requests response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variation too many requests response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variation too many requests response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variation too many requests response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variation too many requests response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationInternalServerError creates a GetKnowledgeKnowledgebaseDocumentVariationInternalServerError with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationInternalServerError() *GetKnowledgeKnowledgebaseDocumentVariationInternalServerError {
	return &GetKnowledgeKnowledgebaseDocumentVariationInternalServerError{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebaseDocumentVariationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variation internal server error response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variation internal server error response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variation internal server error response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document variation internal server error response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase document variation internal server error response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable creates a GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable() *GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable {
	return &GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variation service unavailable response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variation service unavailable response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variation service unavailable response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document variation service unavailable response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase document variation service unavailable response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout creates a GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout() *GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout {
	return &GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variation gateway timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variation gateway timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variation gateway timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document variation gateway timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase document variation gateway timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}][%d] getKnowledgeKnowledgebaseDocumentVariationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
