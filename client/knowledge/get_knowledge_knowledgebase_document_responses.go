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

// GetKnowledgeKnowledgebaseDocumentReader is a Reader for the GetKnowledgeKnowledgebaseDocument structure.
type GetKnowledgeKnowledgebaseDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeKnowledgebaseDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeKnowledgebaseDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeKnowledgebaseDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeKnowledgebaseDocumentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeKnowledgebaseDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeKnowledgebaseDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeKnowledgebaseDocumentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeKnowledgebaseDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeKnowledgebaseDocumentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeKnowledgebaseDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeKnowledgebaseDocumentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeKnowledgebaseDocumentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeKnowledgebaseDocumentOK creates a GetKnowledgeKnowledgebaseDocumentOK with default headers values
func NewGetKnowledgeKnowledgebaseDocumentOK() *GetKnowledgeKnowledgebaseDocumentOK {
	return &GetKnowledgeKnowledgebaseDocumentOK{}
}

/*
GetKnowledgeKnowledgebaseDocumentOK describes a response with status code 200, with default header values.

successful operation
*/
type GetKnowledgeKnowledgebaseDocumentOK struct {
	Payload *models.KnowledgeDocumentResponse
}

// IsSuccess returns true when this get knowledge knowledgebase document o k response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get knowledge knowledgebase document o k response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document o k response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document o k response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document o k response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKnowledgeKnowledgebaseDocumentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentOK) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentOK) GetPayload() *models.KnowledgeDocumentResponse {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentBadRequest creates a GetKnowledgeKnowledgebaseDocumentBadRequest with default headers values
func NewGetKnowledgeKnowledgebaseDocumentBadRequest() *GetKnowledgeKnowledgebaseDocumentBadRequest {
	return &GetKnowledgeKnowledgebaseDocumentBadRequest{}
}

/*
GetKnowledgeKnowledgebaseDocumentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebaseDocumentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document bad request response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document bad request response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document bad request response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document bad request response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document bad request response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKnowledgeKnowledgebaseDocumentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentUnauthorized creates a GetKnowledgeKnowledgebaseDocumentUnauthorized with default headers values
func NewGetKnowledgeKnowledgebaseDocumentUnauthorized() *GetKnowledgeKnowledgebaseDocumentUnauthorized {
	return &GetKnowledgeKnowledgebaseDocumentUnauthorized{}
}

/*
GetKnowledgeKnowledgebaseDocumentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebaseDocumentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document unauthorized response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document unauthorized response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document unauthorized response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document unauthorized response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document unauthorized response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKnowledgeKnowledgebaseDocumentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentForbidden creates a GetKnowledgeKnowledgebaseDocumentForbidden with default headers values
func NewGetKnowledgeKnowledgebaseDocumentForbidden() *GetKnowledgeKnowledgebaseDocumentForbidden {
	return &GetKnowledgeKnowledgebaseDocumentForbidden{}
}

/*
GetKnowledgeKnowledgebaseDocumentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebaseDocumentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document forbidden response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document forbidden response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document forbidden response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document forbidden response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document forbidden response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKnowledgeKnowledgebaseDocumentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentNotFound creates a GetKnowledgeKnowledgebaseDocumentNotFound with default headers values
func NewGetKnowledgeKnowledgebaseDocumentNotFound() *GetKnowledgeKnowledgebaseDocumentNotFound {
	return &GetKnowledgeKnowledgebaseDocumentNotFound{}
}

/*
GetKnowledgeKnowledgebaseDocumentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebaseDocumentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document not found response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document not found response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document not found response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document not found response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document not found response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKnowledgeKnowledgebaseDocumentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentRequestTimeout creates a GetKnowledgeKnowledgebaseDocumentRequestTimeout with default headers values
func NewGetKnowledgeKnowledgebaseDocumentRequestTimeout() *GetKnowledgeKnowledgebaseDocumentRequestTimeout {
	return &GetKnowledgeKnowledgebaseDocumentRequestTimeout{}
}

/*
GetKnowledgeKnowledgebaseDocumentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeKnowledgebaseDocumentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document request timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document request timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document request timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document request timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document request timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetKnowledgeKnowledgebaseDocumentRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge creates a GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge() *GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge{}
}

/*
GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document request entity too large response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document request entity too large response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document request entity too large response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document request entity too large response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document request entity too large response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentUnsupportedMediaType creates a GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType with default headers values
func NewGetKnowledgeKnowledgebaseDocumentUnsupportedMediaType() *GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType {
	return &GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType{}
}

/*
GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document unsupported media type response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document unsupported media type response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document unsupported media type response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document unsupported media type response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document unsupported media type response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentTooManyRequests creates a GetKnowledgeKnowledgebaseDocumentTooManyRequests with default headers values
func NewGetKnowledgeKnowledgebaseDocumentTooManyRequests() *GetKnowledgeKnowledgebaseDocumentTooManyRequests {
	return &GetKnowledgeKnowledgebaseDocumentTooManyRequests{}
}

/*
GetKnowledgeKnowledgebaseDocumentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeKnowledgebaseDocumentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document too many requests response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document too many requests response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document too many requests response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document too many requests response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document too many requests response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetKnowledgeKnowledgebaseDocumentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentInternalServerError creates a GetKnowledgeKnowledgebaseDocumentInternalServerError with default headers values
func NewGetKnowledgeKnowledgebaseDocumentInternalServerError() *GetKnowledgeKnowledgebaseDocumentInternalServerError {
	return &GetKnowledgeKnowledgebaseDocumentInternalServerError{}
}

/*
GetKnowledgeKnowledgebaseDocumentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebaseDocumentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document internal server error response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document internal server error response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document internal server error response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document internal server error response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase document internal server error response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKnowledgeKnowledgebaseDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentServiceUnavailable creates a GetKnowledgeKnowledgebaseDocumentServiceUnavailable with default headers values
func NewGetKnowledgeKnowledgebaseDocumentServiceUnavailable() *GetKnowledgeKnowledgebaseDocumentServiceUnavailable {
	return &GetKnowledgeKnowledgebaseDocumentServiceUnavailable{}
}

/*
GetKnowledgeKnowledgebaseDocumentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebaseDocumentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document service unavailable response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document service unavailable response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document service unavailable response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document service unavailable response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase document service unavailable response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetKnowledgeKnowledgebaseDocumentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentGatewayTimeout creates a GetKnowledgeKnowledgebaseDocumentGatewayTimeout with default headers values
func NewGetKnowledgeKnowledgebaseDocumentGatewayTimeout() *GetKnowledgeKnowledgebaseDocumentGatewayTimeout {
	return &GetKnowledgeKnowledgebaseDocumentGatewayTimeout{}
}

/*
GetKnowledgeKnowledgebaseDocumentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebaseDocumentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document gateway timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document gateway timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document gateway timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document gateway timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase document gateway timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetKnowledgeKnowledgebaseDocumentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] getKnowledgeKnowledgebaseDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
