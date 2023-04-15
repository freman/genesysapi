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

// GetKnowledgeKnowledgebaseDocumentVariationsReader is a Reader for the GetKnowledgeKnowledgebaseDocumentVariations structure.
type GetKnowledgeKnowledgebaseDocumentVariationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeKnowledgebaseDocumentVariationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeKnowledgebaseDocumentVariationsOK creates a GetKnowledgeKnowledgebaseDocumentVariationsOK with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationsOK() *GetKnowledgeKnowledgebaseDocumentVariationsOK {
	return &GetKnowledgeKnowledgebaseDocumentVariationsOK{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetKnowledgeKnowledgebaseDocumentVariationsOK struct {
	Payload *models.DocumentVariationListing
}

// IsSuccess returns true when this get knowledge knowledgebase document variations o k response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get knowledge knowledgebase document variations o k response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variations o k response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document variations o k response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variations o k response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsOK) GetPayload() *models.DocumentVariationListing {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentVariationListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationsBadRequest creates a GetKnowledgeKnowledgebaseDocumentVariationsBadRequest with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationsBadRequest() *GetKnowledgeKnowledgebaseDocumentVariationsBadRequest {
	return &GetKnowledgeKnowledgebaseDocumentVariationsBadRequest{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebaseDocumentVariationsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variations bad request response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variations bad request response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variations bad request response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variations bad request response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variations bad request response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationsUnauthorized creates a GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationsUnauthorized() *GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized {
	return &GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variations unauthorized response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variations unauthorized response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variations unauthorized response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variations unauthorized response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variations unauthorized response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationsForbidden creates a GetKnowledgeKnowledgebaseDocumentVariationsForbidden with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationsForbidden() *GetKnowledgeKnowledgebaseDocumentVariationsForbidden {
	return &GetKnowledgeKnowledgebaseDocumentVariationsForbidden{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebaseDocumentVariationsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variations forbidden response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variations forbidden response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variations forbidden response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variations forbidden response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variations forbidden response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationsNotFound creates a GetKnowledgeKnowledgebaseDocumentVariationsNotFound with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationsNotFound() *GetKnowledgeKnowledgebaseDocumentVariationsNotFound {
	return &GetKnowledgeKnowledgebaseDocumentVariationsNotFound{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebaseDocumentVariationsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variations not found response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variations not found response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variations not found response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variations not found response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variations not found response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout creates a GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout() *GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout {
	return &GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variations request timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variations request timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variations request timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variations request timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variations request timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge creates a GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge() *GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variations request entity too large response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variations request entity too large response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variations request entity too large response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variations request entity too large response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variations request entity too large response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType creates a GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType() *GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType {
	return &GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variations unsupported media type response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variations unsupported media type response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variations unsupported media type response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variations unsupported media type response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variations unsupported media type response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests creates a GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests() *GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests {
	return &GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variations too many requests response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variations too many requests response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variations too many requests response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase document variations too many requests response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase document variations too many requests response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationsInternalServerError creates a GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationsInternalServerError() *GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError {
	return &GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variations internal server error response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variations internal server error response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variations internal server error response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document variations internal server error response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase document variations internal server error response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable creates a GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable() *GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable {
	return &GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variations service unavailable response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variations service unavailable response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variations service unavailable response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document variations service unavailable response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase document variations service unavailable response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout creates a GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout with default headers values
func NewGetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout() *GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout {
	return &GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout{}
}

/*
GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase document variations gateway timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase document variations gateway timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase document variations gateway timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase document variations gateway timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase document variations gateway timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations][%d] getKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentVariationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
