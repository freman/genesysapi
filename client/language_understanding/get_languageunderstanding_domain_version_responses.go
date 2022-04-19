// Code generated by go-swagger; DO NOT EDIT.

package language_understanding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLanguageunderstandingDomainVersionReader is a Reader for the GetLanguageunderstandingDomainVersion structure.
type GetLanguageunderstandingDomainVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguageunderstandingDomainVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguageunderstandingDomainVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguageunderstandingDomainVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguageunderstandingDomainVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguageunderstandingDomainVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguageunderstandingDomainVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLanguageunderstandingDomainVersionRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguageunderstandingDomainVersionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguageunderstandingDomainVersionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguageunderstandingDomainVersionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguageunderstandingDomainVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguageunderstandingDomainVersionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguageunderstandingDomainVersionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguageunderstandingDomainVersionOK creates a GetLanguageunderstandingDomainVersionOK with default headers values
func NewGetLanguageunderstandingDomainVersionOK() *GetLanguageunderstandingDomainVersionOK {
	return &GetLanguageunderstandingDomainVersionOK{}
}

/*GetLanguageunderstandingDomainVersionOK handles this case with default header values.

Retrieved the specified NLU Domain Version
*/
type GetLanguageunderstandingDomainVersionOK struct {
	Payload *models.NluDomainVersion
}

func (o *GetLanguageunderstandingDomainVersionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] getLanguageunderstandingDomainVersionOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionOK) GetPayload() *models.NluDomainVersion {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NluDomainVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionBadRequest creates a GetLanguageunderstandingDomainVersionBadRequest with default headers values
func NewGetLanguageunderstandingDomainVersionBadRequest() *GetLanguageunderstandingDomainVersionBadRequest {
	return &GetLanguageunderstandingDomainVersionBadRequest{}
}

/*GetLanguageunderstandingDomainVersionBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguageunderstandingDomainVersionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] getLanguageunderstandingDomainVersionBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionUnauthorized creates a GetLanguageunderstandingDomainVersionUnauthorized with default headers values
func NewGetLanguageunderstandingDomainVersionUnauthorized() *GetLanguageunderstandingDomainVersionUnauthorized {
	return &GetLanguageunderstandingDomainVersionUnauthorized{}
}

/*GetLanguageunderstandingDomainVersionUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguageunderstandingDomainVersionUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] getLanguageunderstandingDomainVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionForbidden creates a GetLanguageunderstandingDomainVersionForbidden with default headers values
func NewGetLanguageunderstandingDomainVersionForbidden() *GetLanguageunderstandingDomainVersionForbidden {
	return &GetLanguageunderstandingDomainVersionForbidden{}
}

/*GetLanguageunderstandingDomainVersionForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguageunderstandingDomainVersionForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] getLanguageunderstandingDomainVersionForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionNotFound creates a GetLanguageunderstandingDomainVersionNotFound with default headers values
func NewGetLanguageunderstandingDomainVersionNotFound() *GetLanguageunderstandingDomainVersionNotFound {
	return &GetLanguageunderstandingDomainVersionNotFound{}
}

/*GetLanguageunderstandingDomainVersionNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLanguageunderstandingDomainVersionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] getLanguageunderstandingDomainVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionRequestTimeout creates a GetLanguageunderstandingDomainVersionRequestTimeout with default headers values
func NewGetLanguageunderstandingDomainVersionRequestTimeout() *GetLanguageunderstandingDomainVersionRequestTimeout {
	return &GetLanguageunderstandingDomainVersionRequestTimeout{}
}

/*GetLanguageunderstandingDomainVersionRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLanguageunderstandingDomainVersionRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] getLanguageunderstandingDomainVersionRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionRequestEntityTooLarge creates a GetLanguageunderstandingDomainVersionRequestEntityTooLarge with default headers values
func NewGetLanguageunderstandingDomainVersionRequestEntityTooLarge() *GetLanguageunderstandingDomainVersionRequestEntityTooLarge {
	return &GetLanguageunderstandingDomainVersionRequestEntityTooLarge{}
}

/*GetLanguageunderstandingDomainVersionRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetLanguageunderstandingDomainVersionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] getLanguageunderstandingDomainVersionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionUnsupportedMediaType creates a GetLanguageunderstandingDomainVersionUnsupportedMediaType with default headers values
func NewGetLanguageunderstandingDomainVersionUnsupportedMediaType() *GetLanguageunderstandingDomainVersionUnsupportedMediaType {
	return &GetLanguageunderstandingDomainVersionUnsupportedMediaType{}
}

/*GetLanguageunderstandingDomainVersionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguageunderstandingDomainVersionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] getLanguageunderstandingDomainVersionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionTooManyRequests creates a GetLanguageunderstandingDomainVersionTooManyRequests with default headers values
func NewGetLanguageunderstandingDomainVersionTooManyRequests() *GetLanguageunderstandingDomainVersionTooManyRequests {
	return &GetLanguageunderstandingDomainVersionTooManyRequests{}
}

/*GetLanguageunderstandingDomainVersionTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLanguageunderstandingDomainVersionTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] getLanguageunderstandingDomainVersionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionInternalServerError creates a GetLanguageunderstandingDomainVersionInternalServerError with default headers values
func NewGetLanguageunderstandingDomainVersionInternalServerError() *GetLanguageunderstandingDomainVersionInternalServerError {
	return &GetLanguageunderstandingDomainVersionInternalServerError{}
}

/*GetLanguageunderstandingDomainVersionInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguageunderstandingDomainVersionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] getLanguageunderstandingDomainVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionServiceUnavailable creates a GetLanguageunderstandingDomainVersionServiceUnavailable with default headers values
func NewGetLanguageunderstandingDomainVersionServiceUnavailable() *GetLanguageunderstandingDomainVersionServiceUnavailable {
	return &GetLanguageunderstandingDomainVersionServiceUnavailable{}
}

/*GetLanguageunderstandingDomainVersionServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguageunderstandingDomainVersionServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] getLanguageunderstandingDomainVersionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionGatewayTimeout creates a GetLanguageunderstandingDomainVersionGatewayTimeout with default headers values
func NewGetLanguageunderstandingDomainVersionGatewayTimeout() *GetLanguageunderstandingDomainVersionGatewayTimeout {
	return &GetLanguageunderstandingDomainVersionGatewayTimeout{}
}

/*GetLanguageunderstandingDomainVersionGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLanguageunderstandingDomainVersionGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] getLanguageunderstandingDomainVersionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
