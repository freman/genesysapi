// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostExternalcontactsBulkNotesUpdateReader is a Reader for the PostExternalcontactsBulkNotesUpdate structure.
type PostExternalcontactsBulkNotesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkNotesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkNotesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkNotesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkNotesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkNotesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkNotesUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkNotesUpdateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkNotesUpdateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkNotesUpdateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsBulkNotesUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkNotesUpdateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkNotesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkNotesUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkNotesUpdateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkNotesUpdateOK creates a PostExternalcontactsBulkNotesUpdateOK with default headers values
func NewPostExternalcontactsBulkNotesUpdateOK() *PostExternalcontactsBulkNotesUpdateOK {
	return &PostExternalcontactsBulkNotesUpdateOK{}
}

/*
PostExternalcontactsBulkNotesUpdateOK describes a response with status code 200, with default header values.

successful operation
*/
type PostExternalcontactsBulkNotesUpdateOK struct {
	Payload *models.BulkNotesResponse
}

// IsSuccess returns true when this post externalcontacts bulk notes update o k response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post externalcontacts bulk notes update o k response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update o k response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk notes update o k response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk notes update o k response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostExternalcontactsBulkNotesUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateOK) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateOK) GetPayload() *models.BulkNotesResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkNotesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateBadRequest creates a PostExternalcontactsBulkNotesUpdateBadRequest with default headers values
func NewPostExternalcontactsBulkNotesUpdateBadRequest() *PostExternalcontactsBulkNotesUpdateBadRequest {
	return &PostExternalcontactsBulkNotesUpdateBadRequest{}
}

/*
PostExternalcontactsBulkNotesUpdateBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkNotesUpdateBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk notes update bad request response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk notes update bad request response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update bad request response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk notes update bad request response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk notes update bad request response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostExternalcontactsBulkNotesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateUnauthorized creates a PostExternalcontactsBulkNotesUpdateUnauthorized with default headers values
func NewPostExternalcontactsBulkNotesUpdateUnauthorized() *PostExternalcontactsBulkNotesUpdateUnauthorized {
	return &PostExternalcontactsBulkNotesUpdateUnauthorized{}
}

/*
PostExternalcontactsBulkNotesUpdateUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkNotesUpdateUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk notes update unauthorized response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk notes update unauthorized response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update unauthorized response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk notes update unauthorized response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk notes update unauthorized response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostExternalcontactsBulkNotesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateForbidden creates a PostExternalcontactsBulkNotesUpdateForbidden with default headers values
func NewPostExternalcontactsBulkNotesUpdateForbidden() *PostExternalcontactsBulkNotesUpdateForbidden {
	return &PostExternalcontactsBulkNotesUpdateForbidden{}
}

/*
PostExternalcontactsBulkNotesUpdateForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkNotesUpdateForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk notes update forbidden response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk notes update forbidden response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update forbidden response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk notes update forbidden response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk notes update forbidden response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostExternalcontactsBulkNotesUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateNotFound creates a PostExternalcontactsBulkNotesUpdateNotFound with default headers values
func NewPostExternalcontactsBulkNotesUpdateNotFound() *PostExternalcontactsBulkNotesUpdateNotFound {
	return &PostExternalcontactsBulkNotesUpdateNotFound{}
}

/*
PostExternalcontactsBulkNotesUpdateNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkNotesUpdateNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk notes update not found response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk notes update not found response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update not found response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk notes update not found response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk notes update not found response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostExternalcontactsBulkNotesUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateRequestTimeout creates a PostExternalcontactsBulkNotesUpdateRequestTimeout with default headers values
func NewPostExternalcontactsBulkNotesUpdateRequestTimeout() *PostExternalcontactsBulkNotesUpdateRequestTimeout {
	return &PostExternalcontactsBulkNotesUpdateRequestTimeout{}
}

/*
PostExternalcontactsBulkNotesUpdateRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkNotesUpdateRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk notes update request timeout response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk notes update request timeout response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update request timeout response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk notes update request timeout response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk notes update request timeout response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostExternalcontactsBulkNotesUpdateRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateRequestEntityTooLarge creates a PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkNotesUpdateRequestEntityTooLarge() *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge {
	return &PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge{}
}

/*
PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk notes update request entity too large response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk notes update request entity too large response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update request entity too large response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk notes update request entity too large response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk notes update request entity too large response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateUnsupportedMediaType creates a PostExternalcontactsBulkNotesUpdateUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkNotesUpdateUnsupportedMediaType() *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType {
	return &PostExternalcontactsBulkNotesUpdateUnsupportedMediaType{}
}

/*
PostExternalcontactsBulkNotesUpdateUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkNotesUpdateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk notes update unsupported media type response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk notes update unsupported media type response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update unsupported media type response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk notes update unsupported media type response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk notes update unsupported media type response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateUnprocessableEntity creates a PostExternalcontactsBulkNotesUpdateUnprocessableEntity with default headers values
func NewPostExternalcontactsBulkNotesUpdateUnprocessableEntity() *PostExternalcontactsBulkNotesUpdateUnprocessableEntity {
	return &PostExternalcontactsBulkNotesUpdateUnprocessableEntity{}
}

/*
PostExternalcontactsBulkNotesUpdateUnprocessableEntity describes a response with status code 422, with default header values.

PostExternalcontactsBulkNotesUpdateUnprocessableEntity post externalcontacts bulk notes update unprocessable entity
*/
type PostExternalcontactsBulkNotesUpdateUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk notes update unprocessable entity response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk notes update unprocessable entity response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update unprocessable entity response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk notes update unprocessable entity response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk notes update unprocessable entity response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PostExternalcontactsBulkNotesUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateTooManyRequests creates a PostExternalcontactsBulkNotesUpdateTooManyRequests with default headers values
func NewPostExternalcontactsBulkNotesUpdateTooManyRequests() *PostExternalcontactsBulkNotesUpdateTooManyRequests {
	return &PostExternalcontactsBulkNotesUpdateTooManyRequests{}
}

/*
PostExternalcontactsBulkNotesUpdateTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkNotesUpdateTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk notes update too many requests response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk notes update too many requests response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update too many requests response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk notes update too many requests response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk notes update too many requests response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostExternalcontactsBulkNotesUpdateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateInternalServerError creates a PostExternalcontactsBulkNotesUpdateInternalServerError with default headers values
func NewPostExternalcontactsBulkNotesUpdateInternalServerError() *PostExternalcontactsBulkNotesUpdateInternalServerError {
	return &PostExternalcontactsBulkNotesUpdateInternalServerError{}
}

/*
PostExternalcontactsBulkNotesUpdateInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkNotesUpdateInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk notes update internal server error response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk notes update internal server error response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update internal server error response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk notes update internal server error response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk notes update internal server error response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostExternalcontactsBulkNotesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateServiceUnavailable creates a PostExternalcontactsBulkNotesUpdateServiceUnavailable with default headers values
func NewPostExternalcontactsBulkNotesUpdateServiceUnavailable() *PostExternalcontactsBulkNotesUpdateServiceUnavailable {
	return &PostExternalcontactsBulkNotesUpdateServiceUnavailable{}
}

/*
PostExternalcontactsBulkNotesUpdateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkNotesUpdateServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk notes update service unavailable response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk notes update service unavailable response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update service unavailable response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk notes update service unavailable response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk notes update service unavailable response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostExternalcontactsBulkNotesUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateGatewayTimeout creates a PostExternalcontactsBulkNotesUpdateGatewayTimeout with default headers values
func NewPostExternalcontactsBulkNotesUpdateGatewayTimeout() *PostExternalcontactsBulkNotesUpdateGatewayTimeout {
	return &PostExternalcontactsBulkNotesUpdateGatewayTimeout{}
}

/*
PostExternalcontactsBulkNotesUpdateGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostExternalcontactsBulkNotesUpdateGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk notes update gateway timeout response has a 2xx status code
func (o *PostExternalcontactsBulkNotesUpdateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk notes update gateway timeout response has a 3xx status code
func (o *PostExternalcontactsBulkNotesUpdateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk notes update gateway timeout response has a 4xx status code
func (o *PostExternalcontactsBulkNotesUpdateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk notes update gateway timeout response has a 5xx status code
func (o *PostExternalcontactsBulkNotesUpdateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk notes update gateway timeout response a status code equal to that given
func (o *PostExternalcontactsBulkNotesUpdateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostExternalcontactsBulkNotesUpdateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
