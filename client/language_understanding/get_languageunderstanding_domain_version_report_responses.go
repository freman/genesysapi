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

// GetLanguageunderstandingDomainVersionReportReader is a Reader for the GetLanguageunderstandingDomainVersionReport structure.
type GetLanguageunderstandingDomainVersionReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguageunderstandingDomainVersionReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguageunderstandingDomainVersionReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguageunderstandingDomainVersionReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguageunderstandingDomainVersionReportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguageunderstandingDomainVersionReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguageunderstandingDomainVersionReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguageunderstandingDomainVersionReportRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguageunderstandingDomainVersionReportUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguageunderstandingDomainVersionReportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguageunderstandingDomainVersionReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguageunderstandingDomainVersionReportServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguageunderstandingDomainVersionReportGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguageunderstandingDomainVersionReportOK creates a GetLanguageunderstandingDomainVersionReportOK with default headers values
func NewGetLanguageunderstandingDomainVersionReportOK() *GetLanguageunderstandingDomainVersionReportOK {
	return &GetLanguageunderstandingDomainVersionReportOK{}
}

/*GetLanguageunderstandingDomainVersionReportOK handles this case with default header values.

Find quality report for NLU Domain Version.
*/
type GetLanguageunderstandingDomainVersionReportOK struct {
	Payload *models.NluDomainVersionQualityReport
}

func (o *GetLanguageunderstandingDomainVersionReportOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report][%d] getLanguageunderstandingDomainVersionReportOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionReportOK) GetPayload() *models.NluDomainVersionQualityReport {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NluDomainVersionQualityReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionReportBadRequest creates a GetLanguageunderstandingDomainVersionReportBadRequest with default headers values
func NewGetLanguageunderstandingDomainVersionReportBadRequest() *GetLanguageunderstandingDomainVersionReportBadRequest {
	return &GetLanguageunderstandingDomainVersionReportBadRequest{}
}

/*GetLanguageunderstandingDomainVersionReportBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguageunderstandingDomainVersionReportBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionReportBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report][%d] getLanguageunderstandingDomainVersionReportBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionReportBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionReportUnauthorized creates a GetLanguageunderstandingDomainVersionReportUnauthorized with default headers values
func NewGetLanguageunderstandingDomainVersionReportUnauthorized() *GetLanguageunderstandingDomainVersionReportUnauthorized {
	return &GetLanguageunderstandingDomainVersionReportUnauthorized{}
}

/*GetLanguageunderstandingDomainVersionReportUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguageunderstandingDomainVersionReportUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionReportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report][%d] getLanguageunderstandingDomainVersionReportUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionReportUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionReportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionReportForbidden creates a GetLanguageunderstandingDomainVersionReportForbidden with default headers values
func NewGetLanguageunderstandingDomainVersionReportForbidden() *GetLanguageunderstandingDomainVersionReportForbidden {
	return &GetLanguageunderstandingDomainVersionReportForbidden{}
}

/*GetLanguageunderstandingDomainVersionReportForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguageunderstandingDomainVersionReportForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionReportForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report][%d] getLanguageunderstandingDomainVersionReportForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionReportForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionReportNotFound creates a GetLanguageunderstandingDomainVersionReportNotFound with default headers values
func NewGetLanguageunderstandingDomainVersionReportNotFound() *GetLanguageunderstandingDomainVersionReportNotFound {
	return &GetLanguageunderstandingDomainVersionReportNotFound{}
}

/*GetLanguageunderstandingDomainVersionReportNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLanguageunderstandingDomainVersionReportNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionReportNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report][%d] getLanguageunderstandingDomainVersionReportNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionReportNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionReportRequestEntityTooLarge creates a GetLanguageunderstandingDomainVersionReportRequestEntityTooLarge with default headers values
func NewGetLanguageunderstandingDomainVersionReportRequestEntityTooLarge() *GetLanguageunderstandingDomainVersionReportRequestEntityTooLarge {
	return &GetLanguageunderstandingDomainVersionReportRequestEntityTooLarge{}
}

/*GetLanguageunderstandingDomainVersionReportRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLanguageunderstandingDomainVersionReportRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionReportRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report][%d] getLanguageunderstandingDomainVersionReportRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionReportRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionReportRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionReportUnsupportedMediaType creates a GetLanguageunderstandingDomainVersionReportUnsupportedMediaType with default headers values
func NewGetLanguageunderstandingDomainVersionReportUnsupportedMediaType() *GetLanguageunderstandingDomainVersionReportUnsupportedMediaType {
	return &GetLanguageunderstandingDomainVersionReportUnsupportedMediaType{}
}

/*GetLanguageunderstandingDomainVersionReportUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguageunderstandingDomainVersionReportUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionReportUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report][%d] getLanguageunderstandingDomainVersionReportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionReportUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionReportUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionReportTooManyRequests creates a GetLanguageunderstandingDomainVersionReportTooManyRequests with default headers values
func NewGetLanguageunderstandingDomainVersionReportTooManyRequests() *GetLanguageunderstandingDomainVersionReportTooManyRequests {
	return &GetLanguageunderstandingDomainVersionReportTooManyRequests{}
}

/*GetLanguageunderstandingDomainVersionReportTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetLanguageunderstandingDomainVersionReportTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionReportTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report][%d] getLanguageunderstandingDomainVersionReportTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionReportTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionReportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionReportInternalServerError creates a GetLanguageunderstandingDomainVersionReportInternalServerError with default headers values
func NewGetLanguageunderstandingDomainVersionReportInternalServerError() *GetLanguageunderstandingDomainVersionReportInternalServerError {
	return &GetLanguageunderstandingDomainVersionReportInternalServerError{}
}

/*GetLanguageunderstandingDomainVersionReportInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguageunderstandingDomainVersionReportInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionReportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report][%d] getLanguageunderstandingDomainVersionReportInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionReportInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionReportServiceUnavailable creates a GetLanguageunderstandingDomainVersionReportServiceUnavailable with default headers values
func NewGetLanguageunderstandingDomainVersionReportServiceUnavailable() *GetLanguageunderstandingDomainVersionReportServiceUnavailable {
	return &GetLanguageunderstandingDomainVersionReportServiceUnavailable{}
}

/*GetLanguageunderstandingDomainVersionReportServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguageunderstandingDomainVersionReportServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionReportServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report][%d] getLanguageunderstandingDomainVersionReportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionReportServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionReportServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionReportGatewayTimeout creates a GetLanguageunderstandingDomainVersionReportGatewayTimeout with default headers values
func NewGetLanguageunderstandingDomainVersionReportGatewayTimeout() *GetLanguageunderstandingDomainVersionReportGatewayTimeout {
	return &GetLanguageunderstandingDomainVersionReportGatewayTimeout{}
}

/*GetLanguageunderstandingDomainVersionReportGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLanguageunderstandingDomainVersionReportGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionReportGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report][%d] getLanguageunderstandingDomainVersionReportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionReportGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionReportGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}