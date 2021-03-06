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

/*GetKnowledgeKnowledgebaseLanguageCategoriesOK handles this case with default header values.

successful operation
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesOK struct {
	Payload *models.CategoryListing
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesOK) Error() string {
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

/*GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesBadRequest) Error() string {
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

/*GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) Error() string {
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

/*GetKnowledgeKnowledgebaseLanguageCategoriesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesForbidden) Error() string {
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

/*GetKnowledgeKnowledgebaseLanguageCategoriesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesNotFound) Error() string {
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

// NewGetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge creates a GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge() *GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge{}
}

/*GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) Error() string {
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

/*GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) Error() string {
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

/*GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) Error() string {
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

/*GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) Error() string {
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

/*GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) Error() string {
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

/*GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) Error() string {
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
