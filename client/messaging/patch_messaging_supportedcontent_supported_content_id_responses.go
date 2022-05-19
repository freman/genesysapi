// Code generated by go-swagger; DO NOT EDIT.

package messaging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchMessagingSupportedcontentSupportedContentIDReader is a Reader for the PatchMessagingSupportedcontentSupportedContentID structure.
type PatchMessagingSupportedcontentSupportedContentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchMessagingSupportedcontentSupportedContentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchMessagingSupportedcontentSupportedContentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchMessagingSupportedcontentSupportedContentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchMessagingSupportedcontentSupportedContentIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchMessagingSupportedcontentSupportedContentIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchMessagingSupportedcontentSupportedContentIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchMessagingSupportedcontentSupportedContentIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchMessagingSupportedcontentSupportedContentIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchMessagingSupportedcontentSupportedContentIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchMessagingSupportedcontentSupportedContentIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchMessagingSupportedcontentSupportedContentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchMessagingSupportedcontentSupportedContentIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchMessagingSupportedcontentSupportedContentIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchMessagingSupportedcontentSupportedContentIDOK creates a PatchMessagingSupportedcontentSupportedContentIDOK with default headers values
func NewPatchMessagingSupportedcontentSupportedContentIDOK() *PatchMessagingSupportedcontentSupportedContentIDOK {
	return &PatchMessagingSupportedcontentSupportedContentIDOK{}
}

/*PatchMessagingSupportedcontentSupportedContentIDOK handles this case with default header values.

successful operation
*/
type PatchMessagingSupportedcontentSupportedContentIDOK struct {
	Payload *models.SupportedContent
}

func (o *PatchMessagingSupportedcontentSupportedContentIDOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/messaging/supportedcontent/{supportedContentId}][%d] patchMessagingSupportedcontentSupportedContentIdOK  %+v", 200, o.Payload)
}

func (o *PatchMessagingSupportedcontentSupportedContentIDOK) GetPayload() *models.SupportedContent {
	return o.Payload
}

func (o *PatchMessagingSupportedcontentSupportedContentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportedContent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMessagingSupportedcontentSupportedContentIDBadRequest creates a PatchMessagingSupportedcontentSupportedContentIDBadRequest with default headers values
func NewPatchMessagingSupportedcontentSupportedContentIDBadRequest() *PatchMessagingSupportedcontentSupportedContentIDBadRequest {
	return &PatchMessagingSupportedcontentSupportedContentIDBadRequest{}
}

/*PatchMessagingSupportedcontentSupportedContentIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchMessagingSupportedcontentSupportedContentIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchMessagingSupportedcontentSupportedContentIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/messaging/supportedcontent/{supportedContentId}][%d] patchMessagingSupportedcontentSupportedContentIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchMessagingSupportedcontentSupportedContentIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchMessagingSupportedcontentSupportedContentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMessagingSupportedcontentSupportedContentIDUnauthorized creates a PatchMessagingSupportedcontentSupportedContentIDUnauthorized with default headers values
func NewPatchMessagingSupportedcontentSupportedContentIDUnauthorized() *PatchMessagingSupportedcontentSupportedContentIDUnauthorized {
	return &PatchMessagingSupportedcontentSupportedContentIDUnauthorized{}
}

/*PatchMessagingSupportedcontentSupportedContentIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchMessagingSupportedcontentSupportedContentIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchMessagingSupportedcontentSupportedContentIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/messaging/supportedcontent/{supportedContentId}][%d] patchMessagingSupportedcontentSupportedContentIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchMessagingSupportedcontentSupportedContentIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchMessagingSupportedcontentSupportedContentIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMessagingSupportedcontentSupportedContentIDForbidden creates a PatchMessagingSupportedcontentSupportedContentIDForbidden with default headers values
func NewPatchMessagingSupportedcontentSupportedContentIDForbidden() *PatchMessagingSupportedcontentSupportedContentIDForbidden {
	return &PatchMessagingSupportedcontentSupportedContentIDForbidden{}
}

/*PatchMessagingSupportedcontentSupportedContentIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchMessagingSupportedcontentSupportedContentIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchMessagingSupportedcontentSupportedContentIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/messaging/supportedcontent/{supportedContentId}][%d] patchMessagingSupportedcontentSupportedContentIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchMessagingSupportedcontentSupportedContentIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchMessagingSupportedcontentSupportedContentIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMessagingSupportedcontentSupportedContentIDNotFound creates a PatchMessagingSupportedcontentSupportedContentIDNotFound with default headers values
func NewPatchMessagingSupportedcontentSupportedContentIDNotFound() *PatchMessagingSupportedcontentSupportedContentIDNotFound {
	return &PatchMessagingSupportedcontentSupportedContentIDNotFound{}
}

/*PatchMessagingSupportedcontentSupportedContentIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchMessagingSupportedcontentSupportedContentIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchMessagingSupportedcontentSupportedContentIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/messaging/supportedcontent/{supportedContentId}][%d] patchMessagingSupportedcontentSupportedContentIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchMessagingSupportedcontentSupportedContentIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchMessagingSupportedcontentSupportedContentIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMessagingSupportedcontentSupportedContentIDRequestTimeout creates a PatchMessagingSupportedcontentSupportedContentIDRequestTimeout with default headers values
func NewPatchMessagingSupportedcontentSupportedContentIDRequestTimeout() *PatchMessagingSupportedcontentSupportedContentIDRequestTimeout {
	return &PatchMessagingSupportedcontentSupportedContentIDRequestTimeout{}
}

/*PatchMessagingSupportedcontentSupportedContentIDRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchMessagingSupportedcontentSupportedContentIDRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchMessagingSupportedcontentSupportedContentIDRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/messaging/supportedcontent/{supportedContentId}][%d] patchMessagingSupportedcontentSupportedContentIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchMessagingSupportedcontentSupportedContentIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchMessagingSupportedcontentSupportedContentIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMessagingSupportedcontentSupportedContentIDRequestEntityTooLarge creates a PatchMessagingSupportedcontentSupportedContentIDRequestEntityTooLarge with default headers values
func NewPatchMessagingSupportedcontentSupportedContentIDRequestEntityTooLarge() *PatchMessagingSupportedcontentSupportedContentIDRequestEntityTooLarge {
	return &PatchMessagingSupportedcontentSupportedContentIDRequestEntityTooLarge{}
}

/*PatchMessagingSupportedcontentSupportedContentIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchMessagingSupportedcontentSupportedContentIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchMessagingSupportedcontentSupportedContentIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/messaging/supportedcontent/{supportedContentId}][%d] patchMessagingSupportedcontentSupportedContentIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchMessagingSupportedcontentSupportedContentIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchMessagingSupportedcontentSupportedContentIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMessagingSupportedcontentSupportedContentIDUnsupportedMediaType creates a PatchMessagingSupportedcontentSupportedContentIDUnsupportedMediaType with default headers values
func NewPatchMessagingSupportedcontentSupportedContentIDUnsupportedMediaType() *PatchMessagingSupportedcontentSupportedContentIDUnsupportedMediaType {
	return &PatchMessagingSupportedcontentSupportedContentIDUnsupportedMediaType{}
}

/*PatchMessagingSupportedcontentSupportedContentIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchMessagingSupportedcontentSupportedContentIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchMessagingSupportedcontentSupportedContentIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/messaging/supportedcontent/{supportedContentId}][%d] patchMessagingSupportedcontentSupportedContentIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchMessagingSupportedcontentSupportedContentIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchMessagingSupportedcontentSupportedContentIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMessagingSupportedcontentSupportedContentIDTooManyRequests creates a PatchMessagingSupportedcontentSupportedContentIDTooManyRequests with default headers values
func NewPatchMessagingSupportedcontentSupportedContentIDTooManyRequests() *PatchMessagingSupportedcontentSupportedContentIDTooManyRequests {
	return &PatchMessagingSupportedcontentSupportedContentIDTooManyRequests{}
}

/*PatchMessagingSupportedcontentSupportedContentIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchMessagingSupportedcontentSupportedContentIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchMessagingSupportedcontentSupportedContentIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/messaging/supportedcontent/{supportedContentId}][%d] patchMessagingSupportedcontentSupportedContentIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchMessagingSupportedcontentSupportedContentIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchMessagingSupportedcontentSupportedContentIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMessagingSupportedcontentSupportedContentIDInternalServerError creates a PatchMessagingSupportedcontentSupportedContentIDInternalServerError with default headers values
func NewPatchMessagingSupportedcontentSupportedContentIDInternalServerError() *PatchMessagingSupportedcontentSupportedContentIDInternalServerError {
	return &PatchMessagingSupportedcontentSupportedContentIDInternalServerError{}
}

/*PatchMessagingSupportedcontentSupportedContentIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchMessagingSupportedcontentSupportedContentIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchMessagingSupportedcontentSupportedContentIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/messaging/supportedcontent/{supportedContentId}][%d] patchMessagingSupportedcontentSupportedContentIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchMessagingSupportedcontentSupportedContentIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchMessagingSupportedcontentSupportedContentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMessagingSupportedcontentSupportedContentIDServiceUnavailable creates a PatchMessagingSupportedcontentSupportedContentIDServiceUnavailable with default headers values
func NewPatchMessagingSupportedcontentSupportedContentIDServiceUnavailable() *PatchMessagingSupportedcontentSupportedContentIDServiceUnavailable {
	return &PatchMessagingSupportedcontentSupportedContentIDServiceUnavailable{}
}

/*PatchMessagingSupportedcontentSupportedContentIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchMessagingSupportedcontentSupportedContentIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchMessagingSupportedcontentSupportedContentIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/messaging/supportedcontent/{supportedContentId}][%d] patchMessagingSupportedcontentSupportedContentIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchMessagingSupportedcontentSupportedContentIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchMessagingSupportedcontentSupportedContentIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMessagingSupportedcontentSupportedContentIDGatewayTimeout creates a PatchMessagingSupportedcontentSupportedContentIDGatewayTimeout with default headers values
func NewPatchMessagingSupportedcontentSupportedContentIDGatewayTimeout() *PatchMessagingSupportedcontentSupportedContentIDGatewayTimeout {
	return &PatchMessagingSupportedcontentSupportedContentIDGatewayTimeout{}
}

/*PatchMessagingSupportedcontentSupportedContentIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchMessagingSupportedcontentSupportedContentIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchMessagingSupportedcontentSupportedContentIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/messaging/supportedcontent/{supportedContentId}][%d] patchMessagingSupportedcontentSupportedContentIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchMessagingSupportedcontentSupportedContentIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchMessagingSupportedcontentSupportedContentIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}