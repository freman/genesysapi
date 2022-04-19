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

// PatchLanguageunderstandingMinerDraftReader is a Reader for the PatchLanguageunderstandingMinerDraft structure.
type PatchLanguageunderstandingMinerDraftReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLanguageunderstandingMinerDraftReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLanguageunderstandingMinerDraftOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchLanguageunderstandingMinerDraftBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchLanguageunderstandingMinerDraftUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchLanguageunderstandingMinerDraftForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchLanguageunderstandingMinerDraftNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchLanguageunderstandingMinerDraftRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchLanguageunderstandingMinerDraftRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchLanguageunderstandingMinerDraftUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchLanguageunderstandingMinerDraftTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchLanguageunderstandingMinerDraftInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchLanguageunderstandingMinerDraftServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchLanguageunderstandingMinerDraftGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchLanguageunderstandingMinerDraftOK creates a PatchLanguageunderstandingMinerDraftOK with default headers values
func NewPatchLanguageunderstandingMinerDraftOK() *PatchLanguageunderstandingMinerDraftOK {
	return &PatchLanguageunderstandingMinerDraftOK{}
}

/*PatchLanguageunderstandingMinerDraftOK handles this case with default header values.

successful operation
*/
type PatchLanguageunderstandingMinerDraftOK struct {
	Payload *models.Draft
}

func (o *PatchLanguageunderstandingMinerDraftOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftOK  %+v", 200, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftOK) GetPayload() *models.Draft {
	return o.Payload
}

func (o *PatchLanguageunderstandingMinerDraftOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Draft)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLanguageunderstandingMinerDraftBadRequest creates a PatchLanguageunderstandingMinerDraftBadRequest with default headers values
func NewPatchLanguageunderstandingMinerDraftBadRequest() *PatchLanguageunderstandingMinerDraftBadRequest {
	return &PatchLanguageunderstandingMinerDraftBadRequest{}
}

/*PatchLanguageunderstandingMinerDraftBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchLanguageunderstandingMinerDraftBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchLanguageunderstandingMinerDraftBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLanguageunderstandingMinerDraftBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLanguageunderstandingMinerDraftUnauthorized creates a PatchLanguageunderstandingMinerDraftUnauthorized with default headers values
func NewPatchLanguageunderstandingMinerDraftUnauthorized() *PatchLanguageunderstandingMinerDraftUnauthorized {
	return &PatchLanguageunderstandingMinerDraftUnauthorized{}
}

/*PatchLanguageunderstandingMinerDraftUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchLanguageunderstandingMinerDraftUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchLanguageunderstandingMinerDraftUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLanguageunderstandingMinerDraftUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLanguageunderstandingMinerDraftForbidden creates a PatchLanguageunderstandingMinerDraftForbidden with default headers values
func NewPatchLanguageunderstandingMinerDraftForbidden() *PatchLanguageunderstandingMinerDraftForbidden {
	return &PatchLanguageunderstandingMinerDraftForbidden{}
}

/*PatchLanguageunderstandingMinerDraftForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchLanguageunderstandingMinerDraftForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchLanguageunderstandingMinerDraftForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftForbidden  %+v", 403, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLanguageunderstandingMinerDraftForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLanguageunderstandingMinerDraftNotFound creates a PatchLanguageunderstandingMinerDraftNotFound with default headers values
func NewPatchLanguageunderstandingMinerDraftNotFound() *PatchLanguageunderstandingMinerDraftNotFound {
	return &PatchLanguageunderstandingMinerDraftNotFound{}
}

/*PatchLanguageunderstandingMinerDraftNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchLanguageunderstandingMinerDraftNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchLanguageunderstandingMinerDraftNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftNotFound  %+v", 404, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLanguageunderstandingMinerDraftNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLanguageunderstandingMinerDraftRequestTimeout creates a PatchLanguageunderstandingMinerDraftRequestTimeout with default headers values
func NewPatchLanguageunderstandingMinerDraftRequestTimeout() *PatchLanguageunderstandingMinerDraftRequestTimeout {
	return &PatchLanguageunderstandingMinerDraftRequestTimeout{}
}

/*PatchLanguageunderstandingMinerDraftRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchLanguageunderstandingMinerDraftRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchLanguageunderstandingMinerDraftRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLanguageunderstandingMinerDraftRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLanguageunderstandingMinerDraftRequestEntityTooLarge creates a PatchLanguageunderstandingMinerDraftRequestEntityTooLarge with default headers values
func NewPatchLanguageunderstandingMinerDraftRequestEntityTooLarge() *PatchLanguageunderstandingMinerDraftRequestEntityTooLarge {
	return &PatchLanguageunderstandingMinerDraftRequestEntityTooLarge{}
}

/*PatchLanguageunderstandingMinerDraftRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchLanguageunderstandingMinerDraftRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchLanguageunderstandingMinerDraftRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLanguageunderstandingMinerDraftRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLanguageunderstandingMinerDraftUnsupportedMediaType creates a PatchLanguageunderstandingMinerDraftUnsupportedMediaType with default headers values
func NewPatchLanguageunderstandingMinerDraftUnsupportedMediaType() *PatchLanguageunderstandingMinerDraftUnsupportedMediaType {
	return &PatchLanguageunderstandingMinerDraftUnsupportedMediaType{}
}

/*PatchLanguageunderstandingMinerDraftUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchLanguageunderstandingMinerDraftUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchLanguageunderstandingMinerDraftUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLanguageunderstandingMinerDraftUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLanguageunderstandingMinerDraftTooManyRequests creates a PatchLanguageunderstandingMinerDraftTooManyRequests with default headers values
func NewPatchLanguageunderstandingMinerDraftTooManyRequests() *PatchLanguageunderstandingMinerDraftTooManyRequests {
	return &PatchLanguageunderstandingMinerDraftTooManyRequests{}
}

/*PatchLanguageunderstandingMinerDraftTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchLanguageunderstandingMinerDraftTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchLanguageunderstandingMinerDraftTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLanguageunderstandingMinerDraftTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLanguageunderstandingMinerDraftInternalServerError creates a PatchLanguageunderstandingMinerDraftInternalServerError with default headers values
func NewPatchLanguageunderstandingMinerDraftInternalServerError() *PatchLanguageunderstandingMinerDraftInternalServerError {
	return &PatchLanguageunderstandingMinerDraftInternalServerError{}
}

/*PatchLanguageunderstandingMinerDraftInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchLanguageunderstandingMinerDraftInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchLanguageunderstandingMinerDraftInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLanguageunderstandingMinerDraftInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLanguageunderstandingMinerDraftServiceUnavailable creates a PatchLanguageunderstandingMinerDraftServiceUnavailable with default headers values
func NewPatchLanguageunderstandingMinerDraftServiceUnavailable() *PatchLanguageunderstandingMinerDraftServiceUnavailable {
	return &PatchLanguageunderstandingMinerDraftServiceUnavailable{}
}

/*PatchLanguageunderstandingMinerDraftServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchLanguageunderstandingMinerDraftServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchLanguageunderstandingMinerDraftServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLanguageunderstandingMinerDraftServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLanguageunderstandingMinerDraftGatewayTimeout creates a PatchLanguageunderstandingMinerDraftGatewayTimeout with default headers values
func NewPatchLanguageunderstandingMinerDraftGatewayTimeout() *PatchLanguageunderstandingMinerDraftGatewayTimeout {
	return &PatchLanguageunderstandingMinerDraftGatewayTimeout{}
}

/*PatchLanguageunderstandingMinerDraftGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchLanguageunderstandingMinerDraftGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchLanguageunderstandingMinerDraftGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLanguageunderstandingMinerDraftGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
