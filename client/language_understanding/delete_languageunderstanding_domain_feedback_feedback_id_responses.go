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

// DeleteLanguageunderstandingDomainFeedbackFeedbackIDReader is a Reader for the DeleteLanguageunderstandingDomainFeedbackFeedbackID structure.
type DeleteLanguageunderstandingDomainFeedbackFeedbackIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDNoContent creates a DeleteLanguageunderstandingDomainFeedbackFeedbackIDNoContent with default headers values
func NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDNoContent() *DeleteLanguageunderstandingDomainFeedbackFeedbackIDNoContent {
	return &DeleteLanguageunderstandingDomainFeedbackFeedbackIDNoContent{}
}

/*DeleteLanguageunderstandingDomainFeedbackFeedbackIDNoContent handles this case with default header values.

Feedback deleted successfully
*/
type DeleteLanguageunderstandingDomainFeedbackFeedbackIDNoContent struct {
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] deleteLanguageunderstandingDomainFeedbackFeedbackIdNoContent ", 204)
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDBadRequest creates a DeleteLanguageunderstandingDomainFeedbackFeedbackIDBadRequest with default headers values
func NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDBadRequest() *DeleteLanguageunderstandingDomainFeedbackFeedbackIDBadRequest {
	return &DeleteLanguageunderstandingDomainFeedbackFeedbackIDBadRequest{}
}

/*DeleteLanguageunderstandingDomainFeedbackFeedbackIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteLanguageunderstandingDomainFeedbackFeedbackIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] deleteLanguageunderstandingDomainFeedbackFeedbackIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized creates a DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized with default headers values
func NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized() *DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized {
	return &DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized{}
}

/*DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] deleteLanguageunderstandingDomainFeedbackFeedbackIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDForbidden creates a DeleteLanguageunderstandingDomainFeedbackFeedbackIDForbidden with default headers values
func NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDForbidden() *DeleteLanguageunderstandingDomainFeedbackFeedbackIDForbidden {
	return &DeleteLanguageunderstandingDomainFeedbackFeedbackIDForbidden{}
}

/*DeleteLanguageunderstandingDomainFeedbackFeedbackIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteLanguageunderstandingDomainFeedbackFeedbackIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] deleteLanguageunderstandingDomainFeedbackFeedbackIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDNotFound creates a DeleteLanguageunderstandingDomainFeedbackFeedbackIDNotFound with default headers values
func NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDNotFound() *DeleteLanguageunderstandingDomainFeedbackFeedbackIDNotFound {
	return &DeleteLanguageunderstandingDomainFeedbackFeedbackIDNotFound{}
}

/*DeleteLanguageunderstandingDomainFeedbackFeedbackIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteLanguageunderstandingDomainFeedbackFeedbackIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] deleteLanguageunderstandingDomainFeedbackFeedbackIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge creates a DeleteLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge with default headers values
func NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge() *DeleteLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge {
	return &DeleteLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge{}
}

/*DeleteLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] deleteLanguageunderstandingDomainFeedbackFeedbackIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType creates a DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType with default headers values
func NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType() *DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType {
	return &DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType{}
}

/*DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] deleteLanguageunderstandingDomainFeedbackFeedbackIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests creates a DeleteLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests with default headers values
func NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests() *DeleteLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests {
	return &DeleteLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests{}
}

/*DeleteLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] deleteLanguageunderstandingDomainFeedbackFeedbackIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError creates a DeleteLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError with default headers values
func NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError() *DeleteLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError {
	return &DeleteLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError{}
}

/*DeleteLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] deleteLanguageunderstandingDomainFeedbackFeedbackIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable creates a DeleteLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable with default headers values
func NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable() *DeleteLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable {
	return &DeleteLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable{}
}

/*DeleteLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] deleteLanguageunderstandingDomainFeedbackFeedbackIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout creates a DeleteLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout with default headers values
func NewDeleteLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout() *DeleteLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout {
	return &DeleteLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout{}
}

/*DeleteLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] deleteLanguageunderstandingDomainFeedbackFeedbackIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
