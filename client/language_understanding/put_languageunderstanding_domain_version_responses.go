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

// PutLanguageunderstandingDomainVersionReader is a Reader for the PutLanguageunderstandingDomainVersion structure.
type PutLanguageunderstandingDomainVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLanguageunderstandingDomainVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutLanguageunderstandingDomainVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutLanguageunderstandingDomainVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutLanguageunderstandingDomainVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutLanguageunderstandingDomainVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutLanguageunderstandingDomainVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutLanguageunderstandingDomainVersionRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutLanguageunderstandingDomainVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutLanguageunderstandingDomainVersionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutLanguageunderstandingDomainVersionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutLanguageunderstandingDomainVersionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutLanguageunderstandingDomainVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutLanguageunderstandingDomainVersionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutLanguageunderstandingDomainVersionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutLanguageunderstandingDomainVersionOK creates a PutLanguageunderstandingDomainVersionOK with default headers values
func NewPutLanguageunderstandingDomainVersionOK() *PutLanguageunderstandingDomainVersionOK {
	return &PutLanguageunderstandingDomainVersionOK{}
}

/*PutLanguageunderstandingDomainVersionOK handles this case with default header values.

Updated the specified NLU Domain Version
*/
type PutLanguageunderstandingDomainVersionOK struct {
	Payload *models.NluDomainVersion
}

func (o *PutLanguageunderstandingDomainVersionOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionOK  %+v", 200, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionOK) GetPayload() *models.NluDomainVersion {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NluDomainVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLanguageunderstandingDomainVersionBadRequest creates a PutLanguageunderstandingDomainVersionBadRequest with default headers values
func NewPutLanguageunderstandingDomainVersionBadRequest() *PutLanguageunderstandingDomainVersionBadRequest {
	return &PutLanguageunderstandingDomainVersionBadRequest{}
}

/*PutLanguageunderstandingDomainVersionBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutLanguageunderstandingDomainVersionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutLanguageunderstandingDomainVersionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionBadRequest  %+v", 400, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLanguageunderstandingDomainVersionUnauthorized creates a PutLanguageunderstandingDomainVersionUnauthorized with default headers values
func NewPutLanguageunderstandingDomainVersionUnauthorized() *PutLanguageunderstandingDomainVersionUnauthorized {
	return &PutLanguageunderstandingDomainVersionUnauthorized{}
}

/*PutLanguageunderstandingDomainVersionUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutLanguageunderstandingDomainVersionUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutLanguageunderstandingDomainVersionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLanguageunderstandingDomainVersionForbidden creates a PutLanguageunderstandingDomainVersionForbidden with default headers values
func NewPutLanguageunderstandingDomainVersionForbidden() *PutLanguageunderstandingDomainVersionForbidden {
	return &PutLanguageunderstandingDomainVersionForbidden{}
}

/*PutLanguageunderstandingDomainVersionForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutLanguageunderstandingDomainVersionForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutLanguageunderstandingDomainVersionForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionForbidden  %+v", 403, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLanguageunderstandingDomainVersionNotFound creates a PutLanguageunderstandingDomainVersionNotFound with default headers values
func NewPutLanguageunderstandingDomainVersionNotFound() *PutLanguageunderstandingDomainVersionNotFound {
	return &PutLanguageunderstandingDomainVersionNotFound{}
}

/*PutLanguageunderstandingDomainVersionNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutLanguageunderstandingDomainVersionNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutLanguageunderstandingDomainVersionNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionNotFound  %+v", 404, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLanguageunderstandingDomainVersionRequestTimeout creates a PutLanguageunderstandingDomainVersionRequestTimeout with default headers values
func NewPutLanguageunderstandingDomainVersionRequestTimeout() *PutLanguageunderstandingDomainVersionRequestTimeout {
	return &PutLanguageunderstandingDomainVersionRequestTimeout{}
}

/*PutLanguageunderstandingDomainVersionRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutLanguageunderstandingDomainVersionRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutLanguageunderstandingDomainVersionRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLanguageunderstandingDomainVersionConflict creates a PutLanguageunderstandingDomainVersionConflict with default headers values
func NewPutLanguageunderstandingDomainVersionConflict() *PutLanguageunderstandingDomainVersionConflict {
	return &PutLanguageunderstandingDomainVersionConflict{}
}

/*PutLanguageunderstandingDomainVersionConflict handles this case with default header values.

Conflict
*/
type PutLanguageunderstandingDomainVersionConflict struct {
	Payload *models.ErrorBody
}

func (o *PutLanguageunderstandingDomainVersionConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionConflict  %+v", 409, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLanguageunderstandingDomainVersionRequestEntityTooLarge creates a PutLanguageunderstandingDomainVersionRequestEntityTooLarge with default headers values
func NewPutLanguageunderstandingDomainVersionRequestEntityTooLarge() *PutLanguageunderstandingDomainVersionRequestEntityTooLarge {
	return &PutLanguageunderstandingDomainVersionRequestEntityTooLarge{}
}

/*PutLanguageunderstandingDomainVersionRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutLanguageunderstandingDomainVersionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutLanguageunderstandingDomainVersionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLanguageunderstandingDomainVersionUnsupportedMediaType creates a PutLanguageunderstandingDomainVersionUnsupportedMediaType with default headers values
func NewPutLanguageunderstandingDomainVersionUnsupportedMediaType() *PutLanguageunderstandingDomainVersionUnsupportedMediaType {
	return &PutLanguageunderstandingDomainVersionUnsupportedMediaType{}
}

/*PutLanguageunderstandingDomainVersionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutLanguageunderstandingDomainVersionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutLanguageunderstandingDomainVersionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLanguageunderstandingDomainVersionTooManyRequests creates a PutLanguageunderstandingDomainVersionTooManyRequests with default headers values
func NewPutLanguageunderstandingDomainVersionTooManyRequests() *PutLanguageunderstandingDomainVersionTooManyRequests {
	return &PutLanguageunderstandingDomainVersionTooManyRequests{}
}

/*PutLanguageunderstandingDomainVersionTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutLanguageunderstandingDomainVersionTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutLanguageunderstandingDomainVersionTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLanguageunderstandingDomainVersionInternalServerError creates a PutLanguageunderstandingDomainVersionInternalServerError with default headers values
func NewPutLanguageunderstandingDomainVersionInternalServerError() *PutLanguageunderstandingDomainVersionInternalServerError {
	return &PutLanguageunderstandingDomainVersionInternalServerError{}
}

/*PutLanguageunderstandingDomainVersionInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutLanguageunderstandingDomainVersionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutLanguageunderstandingDomainVersionInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLanguageunderstandingDomainVersionServiceUnavailable creates a PutLanguageunderstandingDomainVersionServiceUnavailable with default headers values
func NewPutLanguageunderstandingDomainVersionServiceUnavailable() *PutLanguageunderstandingDomainVersionServiceUnavailable {
	return &PutLanguageunderstandingDomainVersionServiceUnavailable{}
}

/*PutLanguageunderstandingDomainVersionServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutLanguageunderstandingDomainVersionServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutLanguageunderstandingDomainVersionServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLanguageunderstandingDomainVersionGatewayTimeout creates a PutLanguageunderstandingDomainVersionGatewayTimeout with default headers values
func NewPutLanguageunderstandingDomainVersionGatewayTimeout() *PutLanguageunderstandingDomainVersionGatewayTimeout {
	return &PutLanguageunderstandingDomainVersionGatewayTimeout{}
}

/*PutLanguageunderstandingDomainVersionGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutLanguageunderstandingDomainVersionGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutLanguageunderstandingDomainVersionGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}][%d] putLanguageunderstandingDomainVersionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutLanguageunderstandingDomainVersionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutLanguageunderstandingDomainVersionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
