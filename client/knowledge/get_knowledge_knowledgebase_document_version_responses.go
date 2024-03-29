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

// GetKnowledgeKnowledgebaseDocumentVersionReader is a Reader for the GetKnowledgeKnowledgebaseDocumentVersion structure.
type GetKnowledgeKnowledgebaseDocumentVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeKnowledgebaseDocumentVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeKnowledgebaseDocumentVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeKnowledgebaseDocumentVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeKnowledgebaseDocumentVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeKnowledgebaseDocumentVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeKnowledgebaseDocumentVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeKnowledgebaseDocumentVersionRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeKnowledgebaseDocumentVersionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeKnowledgebaseDocumentVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeKnowledgebaseDocumentVersionOK creates a GetKnowledgeKnowledgebaseDocumentVersionOK with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVersionOK() *GetKnowledgeKnowledgebaseDocumentVersionOK {
	return &GetKnowledgeKnowledgebaseDocumentVersionOK{}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionOK describes a response with status code 200, with default header values.

successful operation
*/
type GetKnowledgeKnowledgebaseDocumentVersionOK struct {
	Payload *models.KnowledgeDocumentVersion
}

// IsSuccess returns true when this get knowledge knowledgebase document version o k response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get knowledge knowledgebase document version o k response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document version o k response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document version o k response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document version o k response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionOK) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionOK) GetPayload() *models.KnowledgeDocumentVersion {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeDocumentVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVersionBadRequest creates a GetKnowledgeKnowledgebaseDocumentVersionBadRequest with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVersionBadRequest() *GetKnowledgeKnowledgebaseDocumentVersionBadRequest {
	return &GetKnowledgeKnowledgebaseDocumentVersionBadRequest{}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebaseDocumentVersionBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document version bad request response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document version bad request response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document version bad request response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document version bad request response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document version bad request response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVersionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVersionUnauthorized creates a GetKnowledgeKnowledgebaseDocumentVersionUnauthorized with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVersionUnauthorized() *GetKnowledgeKnowledgebaseDocumentVersionUnauthorized {
	return &GetKnowledgeKnowledgebaseDocumentVersionUnauthorized{}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebaseDocumentVersionUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document version unauthorized response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document version unauthorized response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document version unauthorized response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document version unauthorized response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document version unauthorized response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVersionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVersionForbidden creates a GetKnowledgeKnowledgebaseDocumentVersionForbidden with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVersionForbidden() *GetKnowledgeKnowledgebaseDocumentVersionForbidden {
	return &GetKnowledgeKnowledgebaseDocumentVersionForbidden{}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebaseDocumentVersionForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document version forbidden response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document version forbidden response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document version forbidden response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document version forbidden response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document version forbidden response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVersionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVersionNotFound creates a GetKnowledgeKnowledgebaseDocumentVersionNotFound with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVersionNotFound() *GetKnowledgeKnowledgebaseDocumentVersionNotFound {
	return &GetKnowledgeKnowledgebaseDocumentVersionNotFound{}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebaseDocumentVersionNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document version not found response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document version not found response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document version not found response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document version not found response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document version not found response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVersionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVersionRequestTimeout creates a GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVersionRequestTimeout() *GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout {
	return &GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout{}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document version request timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document version request timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document version request timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document version request timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document version request timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge creates a GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge() *GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge{}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document version request entity too large response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document version request entity too large response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document version request entity too large response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document version request entity too large response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document version request entity too large response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType creates a GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType() *GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType {
	return &GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType{}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document version unsupported media type response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document version unsupported media type response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document version unsupported media type response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document version unsupported media type response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document version unsupported media type response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVersionTooManyRequests creates a GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVersionTooManyRequests() *GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests {
	return &GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests{}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document version too many requests response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document version too many requests response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document version too many requests response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document version too many requests response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document version too many requests response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVersionInternalServerError creates a GetKnowledgeKnowledgebaseDocumentVersionInternalServerError with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVersionInternalServerError() *GetKnowledgeKnowledgebaseDocumentVersionInternalServerError {
	return &GetKnowledgeKnowledgebaseDocumentVersionInternalServerError{}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebaseDocumentVersionInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document version internal server error response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document version internal server error response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document version internal server error response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document version internal server error response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase document version internal server error response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVersionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable creates a GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable() *GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable {
	return &GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable{}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document version service unavailable response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document version service unavailable response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document version service unavailable response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document version service unavailable response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase document version service unavailable response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout creates a GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout() *GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout {
	return &GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout{}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document version gateway timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document version gateway timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document version gateway timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document version gateway timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase document version gateway timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}][%d] getKnowledgeKnowledgebaseDocumentVersionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVersionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
