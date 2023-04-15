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

// GetKnowledgeKnowledgebaseLanguageCategoriesReader is a Reader for the GetKnowledgeKnowledgebaseLanguageCategories structure.
type GetKnowledgeKnowledgebaseLanguageCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeKnowledgebaseLanguageCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeKnowledgebaseLanguageCategoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeKnowledgebaseLanguageCategoriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeKnowledgebaseLanguageCategoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeKnowledgebaseLanguageCategoriesOK creates a GetKnowledgeKnowledgebaseLanguageCategoriesOK with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesOK() *GetKnowledgeKnowledgebaseLanguageCategoriesOK {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesOK{}
}

/*
GetKnowledgeKnowledgebaseLanguageCategoriesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesOK struct {
	Payload *models.CategoryListing
}

// IsSuccess returns true when this get knowledge knowledgebase language categories o k response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get knowledge knowledgebase language categories o k response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language categories o k response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase language categories o k response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language categories o k response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesOK) GetPayload() *models.CategoryListing {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageCategoriesBadRequest creates a GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesBadRequest() *GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest{}
}

/*
GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language categories bad request response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language categories bad request response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language categories bad request response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language categories bad request response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language categories bad request response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized creates a GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized() *GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized{}
}

/*
GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language categories unauthorized response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language categories unauthorized response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language categories unauthorized response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language categories unauthorized response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language categories unauthorized response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageCategoriesForbidden creates a GetKnowledgeKnowledgebaseLanguageCategoriesForbidden with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesForbidden() *GetKnowledgeKnowledgebaseLanguageCategoriesForbidden {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesForbidden{}
}

/*
GetKnowledgeKnowledgebaseLanguageCategoriesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language categories forbidden response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language categories forbidden response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language categories forbidden response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language categories forbidden response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language categories forbidden response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageCategoriesNotFound creates a GetKnowledgeKnowledgebaseLanguageCategoriesNotFound with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesNotFound() *GetKnowledgeKnowledgebaseLanguageCategoriesNotFound {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesNotFound{}
}

/*
GetKnowledgeKnowledgebaseLanguageCategoriesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language categories not found response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language categories not found response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language categories not found response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language categories not found response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language categories not found response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout creates a GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout() *GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout{}
}

/*
GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language categories request timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language categories request timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language categories request timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language categories request timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language categories request timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge creates a GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge() *GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge{}
}

/*
GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language categories request entity too large response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language categories request entity too large response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language categories request entity too large response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language categories request entity too large response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language categories request entity too large response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType creates a GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType() *GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType{}
}

/*
GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language categories unsupported media type response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language categories unsupported media type response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language categories unsupported media type response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language categories unsupported media type response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language categories unsupported media type response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests creates a GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests() *GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests{}
}

/*
GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language categories too many requests response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language categories too many requests response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language categories too many requests response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language categories too many requests response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language categories too many requests response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError creates a GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError() *GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError{}
}

/*
GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language categories internal server error response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language categories internal server error response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language categories internal server error response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase language categories internal server error response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase language categories internal server error response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable creates a GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable() *GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable{}
}

/*
GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language categories service unavailable response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language categories service unavailable response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language categories service unavailable response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase language categories service unavailable response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase language categories service unavailable response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout creates a GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout() *GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout{}
}

/*
GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language categories gateway timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language categories gateway timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language categories gateway timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase language categories gateway timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase language categories gateway timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] getKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
