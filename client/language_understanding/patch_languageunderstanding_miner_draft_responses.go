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

/*
PatchLanguageunderstandingMinerDraftOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchLanguageunderstandingMinerDraftOK struct {
	Payload *models.Draft
}

// IsSuccess returns true when this patch languageunderstanding miner draft o k response has a 2xx status code
func (o *PatchLanguageunderstandingMinerDraftOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch languageunderstanding miner draft o k response has a 3xx status code
func (o *PatchLanguageunderstandingMinerDraftOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch languageunderstanding miner draft o k response has a 4xx status code
func (o *PatchLanguageunderstandingMinerDraftOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch languageunderstanding miner draft o k response has a 5xx status code
func (o *PatchLanguageunderstandingMinerDraftOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch languageunderstanding miner draft o k response a status code equal to that given
func (o *PatchLanguageunderstandingMinerDraftOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchLanguageunderstandingMinerDraftOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftOK  %+v", 200, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftOK) String() string {
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

/*
PatchLanguageunderstandingMinerDraftBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchLanguageunderstandingMinerDraftBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch languageunderstanding miner draft bad request response has a 2xx status code
func (o *PatchLanguageunderstandingMinerDraftBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch languageunderstanding miner draft bad request response has a 3xx status code
func (o *PatchLanguageunderstandingMinerDraftBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch languageunderstanding miner draft bad request response has a 4xx status code
func (o *PatchLanguageunderstandingMinerDraftBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch languageunderstanding miner draft bad request response has a 5xx status code
func (o *PatchLanguageunderstandingMinerDraftBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch languageunderstanding miner draft bad request response a status code equal to that given
func (o *PatchLanguageunderstandingMinerDraftBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchLanguageunderstandingMinerDraftBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftBadRequest) String() string {
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

/*
PatchLanguageunderstandingMinerDraftUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchLanguageunderstandingMinerDraftUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch languageunderstanding miner draft unauthorized response has a 2xx status code
func (o *PatchLanguageunderstandingMinerDraftUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch languageunderstanding miner draft unauthorized response has a 3xx status code
func (o *PatchLanguageunderstandingMinerDraftUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch languageunderstanding miner draft unauthorized response has a 4xx status code
func (o *PatchLanguageunderstandingMinerDraftUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch languageunderstanding miner draft unauthorized response has a 5xx status code
func (o *PatchLanguageunderstandingMinerDraftUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch languageunderstanding miner draft unauthorized response a status code equal to that given
func (o *PatchLanguageunderstandingMinerDraftUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchLanguageunderstandingMinerDraftUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftUnauthorized) String() string {
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

/*
PatchLanguageunderstandingMinerDraftForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchLanguageunderstandingMinerDraftForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch languageunderstanding miner draft forbidden response has a 2xx status code
func (o *PatchLanguageunderstandingMinerDraftForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch languageunderstanding miner draft forbidden response has a 3xx status code
func (o *PatchLanguageunderstandingMinerDraftForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch languageunderstanding miner draft forbidden response has a 4xx status code
func (o *PatchLanguageunderstandingMinerDraftForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch languageunderstanding miner draft forbidden response has a 5xx status code
func (o *PatchLanguageunderstandingMinerDraftForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch languageunderstanding miner draft forbidden response a status code equal to that given
func (o *PatchLanguageunderstandingMinerDraftForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchLanguageunderstandingMinerDraftForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftForbidden  %+v", 403, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftForbidden) String() string {
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

/*
PatchLanguageunderstandingMinerDraftNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchLanguageunderstandingMinerDraftNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch languageunderstanding miner draft not found response has a 2xx status code
func (o *PatchLanguageunderstandingMinerDraftNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch languageunderstanding miner draft not found response has a 3xx status code
func (o *PatchLanguageunderstandingMinerDraftNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch languageunderstanding miner draft not found response has a 4xx status code
func (o *PatchLanguageunderstandingMinerDraftNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch languageunderstanding miner draft not found response has a 5xx status code
func (o *PatchLanguageunderstandingMinerDraftNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch languageunderstanding miner draft not found response a status code equal to that given
func (o *PatchLanguageunderstandingMinerDraftNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchLanguageunderstandingMinerDraftNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftNotFound  %+v", 404, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftNotFound) String() string {
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

/*
PatchLanguageunderstandingMinerDraftRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchLanguageunderstandingMinerDraftRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch languageunderstanding miner draft request timeout response has a 2xx status code
func (o *PatchLanguageunderstandingMinerDraftRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch languageunderstanding miner draft request timeout response has a 3xx status code
func (o *PatchLanguageunderstandingMinerDraftRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch languageunderstanding miner draft request timeout response has a 4xx status code
func (o *PatchLanguageunderstandingMinerDraftRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch languageunderstanding miner draft request timeout response has a 5xx status code
func (o *PatchLanguageunderstandingMinerDraftRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch languageunderstanding miner draft request timeout response a status code equal to that given
func (o *PatchLanguageunderstandingMinerDraftRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchLanguageunderstandingMinerDraftRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftRequestTimeout) String() string {
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

/*
PatchLanguageunderstandingMinerDraftRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchLanguageunderstandingMinerDraftRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch languageunderstanding miner draft request entity too large response has a 2xx status code
func (o *PatchLanguageunderstandingMinerDraftRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch languageunderstanding miner draft request entity too large response has a 3xx status code
func (o *PatchLanguageunderstandingMinerDraftRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch languageunderstanding miner draft request entity too large response has a 4xx status code
func (o *PatchLanguageunderstandingMinerDraftRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch languageunderstanding miner draft request entity too large response has a 5xx status code
func (o *PatchLanguageunderstandingMinerDraftRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch languageunderstanding miner draft request entity too large response a status code equal to that given
func (o *PatchLanguageunderstandingMinerDraftRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchLanguageunderstandingMinerDraftRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftRequestEntityTooLarge) String() string {
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

/*
PatchLanguageunderstandingMinerDraftUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchLanguageunderstandingMinerDraftUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch languageunderstanding miner draft unsupported media type response has a 2xx status code
func (o *PatchLanguageunderstandingMinerDraftUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch languageunderstanding miner draft unsupported media type response has a 3xx status code
func (o *PatchLanguageunderstandingMinerDraftUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch languageunderstanding miner draft unsupported media type response has a 4xx status code
func (o *PatchLanguageunderstandingMinerDraftUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch languageunderstanding miner draft unsupported media type response has a 5xx status code
func (o *PatchLanguageunderstandingMinerDraftUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch languageunderstanding miner draft unsupported media type response a status code equal to that given
func (o *PatchLanguageunderstandingMinerDraftUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchLanguageunderstandingMinerDraftUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftUnsupportedMediaType) String() string {
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

/*
PatchLanguageunderstandingMinerDraftTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchLanguageunderstandingMinerDraftTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch languageunderstanding miner draft too many requests response has a 2xx status code
func (o *PatchLanguageunderstandingMinerDraftTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch languageunderstanding miner draft too many requests response has a 3xx status code
func (o *PatchLanguageunderstandingMinerDraftTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch languageunderstanding miner draft too many requests response has a 4xx status code
func (o *PatchLanguageunderstandingMinerDraftTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch languageunderstanding miner draft too many requests response has a 5xx status code
func (o *PatchLanguageunderstandingMinerDraftTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch languageunderstanding miner draft too many requests response a status code equal to that given
func (o *PatchLanguageunderstandingMinerDraftTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchLanguageunderstandingMinerDraftTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftTooManyRequests) String() string {
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

/*
PatchLanguageunderstandingMinerDraftInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchLanguageunderstandingMinerDraftInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch languageunderstanding miner draft internal server error response has a 2xx status code
func (o *PatchLanguageunderstandingMinerDraftInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch languageunderstanding miner draft internal server error response has a 3xx status code
func (o *PatchLanguageunderstandingMinerDraftInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch languageunderstanding miner draft internal server error response has a 4xx status code
func (o *PatchLanguageunderstandingMinerDraftInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch languageunderstanding miner draft internal server error response has a 5xx status code
func (o *PatchLanguageunderstandingMinerDraftInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch languageunderstanding miner draft internal server error response a status code equal to that given
func (o *PatchLanguageunderstandingMinerDraftInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchLanguageunderstandingMinerDraftInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftInternalServerError) String() string {
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

/*
PatchLanguageunderstandingMinerDraftServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchLanguageunderstandingMinerDraftServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch languageunderstanding miner draft service unavailable response has a 2xx status code
func (o *PatchLanguageunderstandingMinerDraftServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch languageunderstanding miner draft service unavailable response has a 3xx status code
func (o *PatchLanguageunderstandingMinerDraftServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch languageunderstanding miner draft service unavailable response has a 4xx status code
func (o *PatchLanguageunderstandingMinerDraftServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch languageunderstanding miner draft service unavailable response has a 5xx status code
func (o *PatchLanguageunderstandingMinerDraftServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch languageunderstanding miner draft service unavailable response a status code equal to that given
func (o *PatchLanguageunderstandingMinerDraftServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchLanguageunderstandingMinerDraftServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftServiceUnavailable) String() string {
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

/*
PatchLanguageunderstandingMinerDraftGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchLanguageunderstandingMinerDraftGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch languageunderstanding miner draft gateway timeout response has a 2xx status code
func (o *PatchLanguageunderstandingMinerDraftGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch languageunderstanding miner draft gateway timeout response has a 3xx status code
func (o *PatchLanguageunderstandingMinerDraftGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch languageunderstanding miner draft gateway timeout response has a 4xx status code
func (o *PatchLanguageunderstandingMinerDraftGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch languageunderstanding miner draft gateway timeout response has a 5xx status code
func (o *PatchLanguageunderstandingMinerDraftGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch languageunderstanding miner draft gateway timeout response a status code equal to that given
func (o *PatchLanguageunderstandingMinerDraftGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchLanguageunderstandingMinerDraftGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/languageunderstanding/miners/{minerId}/drafts/{draftId}][%d] patchLanguageunderstandingMinerDraftGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchLanguageunderstandingMinerDraftGatewayTimeout) String() string {
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
