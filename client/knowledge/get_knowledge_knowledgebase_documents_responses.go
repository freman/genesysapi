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

// GetKnowledgeKnowledgebaseDocumentsReader is a Reader for the GetKnowledgeKnowledgebaseDocuments structure.
type GetKnowledgeKnowledgebaseDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeKnowledgebaseDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeKnowledgebaseDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeKnowledgebaseDocumentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeKnowledgebaseDocumentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeKnowledgebaseDocumentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeKnowledgebaseDocumentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeKnowledgebaseDocumentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeKnowledgebaseDocumentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeKnowledgebaseDocumentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeKnowledgebaseDocumentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeKnowledgebaseDocumentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeKnowledgebaseDocumentsOK creates a GetKnowledgeKnowledgebaseDocumentsOK with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsOK() *GetKnowledgeKnowledgebaseDocumentsOK {
	return &GetKnowledgeKnowledgebaseDocumentsOK{}
}

/*
GetKnowledgeKnowledgebaseDocumentsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetKnowledgeKnowledgebaseDocumentsOK struct {
	Payload *models.KnowledgeDocumentResponseListing
}

// IsSuccess returns true when this get knowledge knowledgebase documents o k response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get knowledge knowledgebase documents o k response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase documents o k response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase documents o k response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase documents o k response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKnowledgeKnowledgebaseDocumentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsOK) GetPayload() *models.KnowledgeDocumentResponseListing {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeDocumentResponseListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsBadRequest creates a GetKnowledgeKnowledgebaseDocumentsBadRequest with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsBadRequest() *GetKnowledgeKnowledgebaseDocumentsBadRequest {
	return &GetKnowledgeKnowledgebaseDocumentsBadRequest{}
}

/*
GetKnowledgeKnowledgebaseDocumentsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebaseDocumentsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase documents bad request response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase documents bad request response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase documents bad request response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase documents bad request response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase documents bad request response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKnowledgeKnowledgebaseDocumentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsUnauthorized creates a GetKnowledgeKnowledgebaseDocumentsUnauthorized with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsUnauthorized() *GetKnowledgeKnowledgebaseDocumentsUnauthorized {
	return &GetKnowledgeKnowledgebaseDocumentsUnauthorized{}
}

/*
GetKnowledgeKnowledgebaseDocumentsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebaseDocumentsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase documents unauthorized response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase documents unauthorized response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase documents unauthorized response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase documents unauthorized response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase documents unauthorized response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsForbidden creates a GetKnowledgeKnowledgebaseDocumentsForbidden with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsForbidden() *GetKnowledgeKnowledgebaseDocumentsForbidden {
	return &GetKnowledgeKnowledgebaseDocumentsForbidden{}
}

/*
GetKnowledgeKnowledgebaseDocumentsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebaseDocumentsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase documents forbidden response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase documents forbidden response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase documents forbidden response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase documents forbidden response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase documents forbidden response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKnowledgeKnowledgebaseDocumentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsNotFound creates a GetKnowledgeKnowledgebaseDocumentsNotFound with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsNotFound() *GetKnowledgeKnowledgebaseDocumentsNotFound {
	return &GetKnowledgeKnowledgebaseDocumentsNotFound{}
}

/*
GetKnowledgeKnowledgebaseDocumentsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebaseDocumentsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase documents not found response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase documents not found response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase documents not found response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase documents not found response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase documents not found response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKnowledgeKnowledgebaseDocumentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsRequestTimeout creates a GetKnowledgeKnowledgebaseDocumentsRequestTimeout with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsRequestTimeout() *GetKnowledgeKnowledgebaseDocumentsRequestTimeout {
	return &GetKnowledgeKnowledgebaseDocumentsRequestTimeout{}
}

/*
GetKnowledgeKnowledgebaseDocumentsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeKnowledgebaseDocumentsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase documents request timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase documents request timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase documents request timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase documents request timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase documents request timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge creates a GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge() *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge{}
}

/*
GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase documents request entity too large response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase documents request entity too large response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase documents request entity too large response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase documents request entity too large response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase documents request entity too large response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType creates a GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType() *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType {
	return &GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType{}
}

/*
GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase documents unsupported media type response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase documents unsupported media type response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase documents unsupported media type response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase documents unsupported media type response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase documents unsupported media type response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsTooManyRequests creates a GetKnowledgeKnowledgebaseDocumentsTooManyRequests with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsTooManyRequests() *GetKnowledgeKnowledgebaseDocumentsTooManyRequests {
	return &GetKnowledgeKnowledgebaseDocumentsTooManyRequests{}
}

/*
GetKnowledgeKnowledgebaseDocumentsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeKnowledgebaseDocumentsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase documents too many requests response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase documents too many requests response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase documents too many requests response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase documents too many requests response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase documents too many requests response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetKnowledgeKnowledgebaseDocumentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsInternalServerError creates a GetKnowledgeKnowledgebaseDocumentsInternalServerError with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsInternalServerError() *GetKnowledgeKnowledgebaseDocumentsInternalServerError {
	return &GetKnowledgeKnowledgebaseDocumentsInternalServerError{}
}

/*
GetKnowledgeKnowledgebaseDocumentsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebaseDocumentsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase documents internal server error response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase documents internal server error response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase documents internal server error response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase documents internal server error response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase documents internal server error response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKnowledgeKnowledgebaseDocumentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsServiceUnavailable creates a GetKnowledgeKnowledgebaseDocumentsServiceUnavailable with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsServiceUnavailable() *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable {
	return &GetKnowledgeKnowledgebaseDocumentsServiceUnavailable{}
}

/*
GetKnowledgeKnowledgebaseDocumentsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebaseDocumentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase documents service unavailable response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase documents service unavailable response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase documents service unavailable response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase documents service unavailable response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase documents service unavailable response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsGatewayTimeout creates a GetKnowledgeKnowledgebaseDocumentsGatewayTimeout with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsGatewayTimeout() *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout {
	return &GetKnowledgeKnowledgebaseDocumentsGatewayTimeout{}
}

/*
GetKnowledgeKnowledgebaseDocumentsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebaseDocumentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase documents gateway timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase documents gateway timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase documents gateway timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase documents gateway timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase documents gateway timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
