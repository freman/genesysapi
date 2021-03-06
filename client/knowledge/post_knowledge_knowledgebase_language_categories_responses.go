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

// PostKnowledgeKnowledgebaseLanguageCategoriesReader is a Reader for the PostKnowledgeKnowledgebaseLanguageCategories structure.
type PostKnowledgeKnowledgebaseLanguageCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostKnowledgeKnowledgebaseLanguageCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostKnowledgeKnowledgebaseLanguageCategoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostKnowledgeKnowledgebaseLanguageCategoriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostKnowledgeKnowledgebaseLanguageCategoriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostKnowledgeKnowledgebaseLanguageCategoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostKnowledgeKnowledgebaseLanguageCategoriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesOK creates a PostKnowledgeKnowledgebaseLanguageCategoriesOK with default headers values
func NewPostKnowledgeKnowledgebaseLanguageCategoriesOK() *PostKnowledgeKnowledgebaseLanguageCategoriesOK {
	return &PostKnowledgeKnowledgebaseLanguageCategoriesOK{}
}

/*PostKnowledgeKnowledgebaseLanguageCategoriesOK handles this case with default header values.

successful operation
*/
type PostKnowledgeKnowledgebaseLanguageCategoriesOK struct {
	Payload *models.KnowledgeExtendedCategory
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] postKnowledgeKnowledgebaseLanguageCategoriesOK  %+v", 200, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesOK) GetPayload() *models.KnowledgeExtendedCategory {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeExtendedCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesBadRequest creates a PostKnowledgeKnowledgebaseLanguageCategoriesBadRequest with default headers values
func NewPostKnowledgeKnowledgebaseLanguageCategoriesBadRequest() *PostKnowledgeKnowledgebaseLanguageCategoriesBadRequest {
	return &PostKnowledgeKnowledgebaseLanguageCategoriesBadRequest{}
}

/*PostKnowledgeKnowledgebaseLanguageCategoriesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostKnowledgeKnowledgebaseLanguageCategoriesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] postKnowledgeKnowledgebaseLanguageCategoriesBadRequest  %+v", 400, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesUnauthorized creates a PostKnowledgeKnowledgebaseLanguageCategoriesUnauthorized with default headers values
func NewPostKnowledgeKnowledgebaseLanguageCategoriesUnauthorized() *PostKnowledgeKnowledgebaseLanguageCategoriesUnauthorized {
	return &PostKnowledgeKnowledgebaseLanguageCategoriesUnauthorized{}
}

/*PostKnowledgeKnowledgebaseLanguageCategoriesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostKnowledgeKnowledgebaseLanguageCategoriesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] postKnowledgeKnowledgebaseLanguageCategoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesForbidden creates a PostKnowledgeKnowledgebaseLanguageCategoriesForbidden with default headers values
func NewPostKnowledgeKnowledgebaseLanguageCategoriesForbidden() *PostKnowledgeKnowledgebaseLanguageCategoriesForbidden {
	return &PostKnowledgeKnowledgebaseLanguageCategoriesForbidden{}
}

/*PostKnowledgeKnowledgebaseLanguageCategoriesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostKnowledgeKnowledgebaseLanguageCategoriesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] postKnowledgeKnowledgebaseLanguageCategoriesForbidden  %+v", 403, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesNotFound creates a PostKnowledgeKnowledgebaseLanguageCategoriesNotFound with default headers values
func NewPostKnowledgeKnowledgebaseLanguageCategoriesNotFound() *PostKnowledgeKnowledgebaseLanguageCategoriesNotFound {
	return &PostKnowledgeKnowledgebaseLanguageCategoriesNotFound{}
}

/*PostKnowledgeKnowledgebaseLanguageCategoriesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostKnowledgeKnowledgebaseLanguageCategoriesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] postKnowledgeKnowledgebaseLanguageCategoriesNotFound  %+v", 404, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge creates a PostKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge with default headers values
func NewPostKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge() *PostKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge {
	return &PostKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge{}
}

/*PostKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] postKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType creates a PostKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType with default headers values
func NewPostKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType() *PostKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType {
	return &PostKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType{}
}

/*PostKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] postKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests creates a PostKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests with default headers values
func NewPostKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests() *PostKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests {
	return &PostKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests{}
}

/*PostKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] postKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesInternalServerError creates a PostKnowledgeKnowledgebaseLanguageCategoriesInternalServerError with default headers values
func NewPostKnowledgeKnowledgebaseLanguageCategoriesInternalServerError() *PostKnowledgeKnowledgebaseLanguageCategoriesInternalServerError {
	return &PostKnowledgeKnowledgebaseLanguageCategoriesInternalServerError{}
}

/*PostKnowledgeKnowledgebaseLanguageCategoriesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostKnowledgeKnowledgebaseLanguageCategoriesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] postKnowledgeKnowledgebaseLanguageCategoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable creates a PostKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable with default headers values
func NewPostKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable() *PostKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable {
	return &PostKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable{}
}

/*PostKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] postKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout creates a PostKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout with default headers values
func NewPostKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout() *PostKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout {
	return &PostKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout{}
}

/*PostKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories][%d] postKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageCategoriesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
