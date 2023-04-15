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

// PostExternalcontactsBulkContactsUpdateReader is a Reader for the PostExternalcontactsBulkContactsUpdate structure.
type PostExternalcontactsBulkContactsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkContactsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkContactsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkContactsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkContactsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkContactsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkContactsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkContactsUpdateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkContactsUpdateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkContactsUpdateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsBulkContactsUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkContactsUpdateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkContactsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkContactsUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkContactsUpdateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkContactsUpdateOK creates a PostExternalcontactsBulkContactsUpdateOK with default headers values
func NewPostExternalcontactsBulkContactsUpdateOK() *PostExternalcontactsBulkContactsUpdateOK {
	return &PostExternalcontactsBulkContactsUpdateOK{}
}

/*
PostExternalcontactsBulkContactsUpdateOK describes a response with status code 200, with default header values.

successful operation
*/
type PostExternalcontactsBulkContactsUpdateOK struct {
	Payload *models.BulkContactsResponse
}

// IsSuccess returns true when this post externalcontacts bulk contacts update o k response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post externalcontacts bulk contacts update o k response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update o k response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk contacts update o k response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts update o k response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostExternalcontactsBulkContactsUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateOK) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateOK) GetPayload() *models.BulkContactsResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkContactsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateBadRequest creates a PostExternalcontactsBulkContactsUpdateBadRequest with default headers values
func NewPostExternalcontactsBulkContactsUpdateBadRequest() *PostExternalcontactsBulkContactsUpdateBadRequest {
	return &PostExternalcontactsBulkContactsUpdateBadRequest{}
}

/*
PostExternalcontactsBulkContactsUpdateBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkContactsUpdateBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts update bad request response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts update bad request response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update bad request response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts update bad request response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts update bad request response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostExternalcontactsBulkContactsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateUnauthorized creates a PostExternalcontactsBulkContactsUpdateUnauthorized with default headers values
func NewPostExternalcontactsBulkContactsUpdateUnauthorized() *PostExternalcontactsBulkContactsUpdateUnauthorized {
	return &PostExternalcontactsBulkContactsUpdateUnauthorized{}
}

/*
PostExternalcontactsBulkContactsUpdateUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkContactsUpdateUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts update unauthorized response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts update unauthorized response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update unauthorized response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts update unauthorized response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts update unauthorized response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostExternalcontactsBulkContactsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateForbidden creates a PostExternalcontactsBulkContactsUpdateForbidden with default headers values
func NewPostExternalcontactsBulkContactsUpdateForbidden() *PostExternalcontactsBulkContactsUpdateForbidden {
	return &PostExternalcontactsBulkContactsUpdateForbidden{}
}

/*
PostExternalcontactsBulkContactsUpdateForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkContactsUpdateForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts update forbidden response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts update forbidden response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update forbidden response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts update forbidden response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts update forbidden response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostExternalcontactsBulkContactsUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateNotFound creates a PostExternalcontactsBulkContactsUpdateNotFound with default headers values
func NewPostExternalcontactsBulkContactsUpdateNotFound() *PostExternalcontactsBulkContactsUpdateNotFound {
	return &PostExternalcontactsBulkContactsUpdateNotFound{}
}

/*
PostExternalcontactsBulkContactsUpdateNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkContactsUpdateNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts update not found response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts update not found response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update not found response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts update not found response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts update not found response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostExternalcontactsBulkContactsUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateRequestTimeout creates a PostExternalcontactsBulkContactsUpdateRequestTimeout with default headers values
func NewPostExternalcontactsBulkContactsUpdateRequestTimeout() *PostExternalcontactsBulkContactsUpdateRequestTimeout {
	return &PostExternalcontactsBulkContactsUpdateRequestTimeout{}
}

/*
PostExternalcontactsBulkContactsUpdateRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkContactsUpdateRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts update request timeout response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts update request timeout response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update request timeout response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts update request timeout response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts update request timeout response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostExternalcontactsBulkContactsUpdateRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateRequestEntityTooLarge creates a PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkContactsUpdateRequestEntityTooLarge() *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge {
	return &PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge{}
}

/*
PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts update request entity too large response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts update request entity too large response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update request entity too large response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts update request entity too large response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts update request entity too large response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateUnsupportedMediaType creates a PostExternalcontactsBulkContactsUpdateUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkContactsUpdateUnsupportedMediaType() *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType {
	return &PostExternalcontactsBulkContactsUpdateUnsupportedMediaType{}
}

/*
PostExternalcontactsBulkContactsUpdateUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkContactsUpdateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts update unsupported media type response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts update unsupported media type response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update unsupported media type response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts update unsupported media type response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts update unsupported media type response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateUnprocessableEntity creates a PostExternalcontactsBulkContactsUpdateUnprocessableEntity with default headers values
func NewPostExternalcontactsBulkContactsUpdateUnprocessableEntity() *PostExternalcontactsBulkContactsUpdateUnprocessableEntity {
	return &PostExternalcontactsBulkContactsUpdateUnprocessableEntity{}
}

/*
PostExternalcontactsBulkContactsUpdateUnprocessableEntity describes a response with status code 422, with default header values.

PostExternalcontactsBulkContactsUpdateUnprocessableEntity post externalcontacts bulk contacts update unprocessable entity
*/
type PostExternalcontactsBulkContactsUpdateUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts update unprocessable entity response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts update unprocessable entity response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update unprocessable entity response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts update unprocessable entity response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts update unprocessable entity response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PostExternalcontactsBulkContactsUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateTooManyRequests creates a PostExternalcontactsBulkContactsUpdateTooManyRequests with default headers values
func NewPostExternalcontactsBulkContactsUpdateTooManyRequests() *PostExternalcontactsBulkContactsUpdateTooManyRequests {
	return &PostExternalcontactsBulkContactsUpdateTooManyRequests{}
}

/*
PostExternalcontactsBulkContactsUpdateTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkContactsUpdateTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts update too many requests response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts update too many requests response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update too many requests response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts update too many requests response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts update too many requests response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostExternalcontactsBulkContactsUpdateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateInternalServerError creates a PostExternalcontactsBulkContactsUpdateInternalServerError with default headers values
func NewPostExternalcontactsBulkContactsUpdateInternalServerError() *PostExternalcontactsBulkContactsUpdateInternalServerError {
	return &PostExternalcontactsBulkContactsUpdateInternalServerError{}
}

/*
PostExternalcontactsBulkContactsUpdateInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkContactsUpdateInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts update internal server error response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts update internal server error response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update internal server error response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk contacts update internal server error response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk contacts update internal server error response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostExternalcontactsBulkContactsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateServiceUnavailable creates a PostExternalcontactsBulkContactsUpdateServiceUnavailable with default headers values
func NewPostExternalcontactsBulkContactsUpdateServiceUnavailable() *PostExternalcontactsBulkContactsUpdateServiceUnavailable {
	return &PostExternalcontactsBulkContactsUpdateServiceUnavailable{}
}

/*
PostExternalcontactsBulkContactsUpdateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkContactsUpdateServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts update service unavailable response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts update service unavailable response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update service unavailable response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk contacts update service unavailable response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk contacts update service unavailable response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostExternalcontactsBulkContactsUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateGatewayTimeout creates a PostExternalcontactsBulkContactsUpdateGatewayTimeout with default headers values
func NewPostExternalcontactsBulkContactsUpdateGatewayTimeout() *PostExternalcontactsBulkContactsUpdateGatewayTimeout {
	return &PostExternalcontactsBulkContactsUpdateGatewayTimeout{}
}

/*
PostExternalcontactsBulkContactsUpdateGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostExternalcontactsBulkContactsUpdateGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts update gateway timeout response has a 2xx status code
func (o *PostExternalcontactsBulkContactsUpdateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts update gateway timeout response has a 3xx status code
func (o *PostExternalcontactsBulkContactsUpdateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts update gateway timeout response has a 4xx status code
func (o *PostExternalcontactsBulkContactsUpdateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk contacts update gateway timeout response has a 5xx status code
func (o *PostExternalcontactsBulkContactsUpdateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk contacts update gateway timeout response a status code equal to that given
func (o *PostExternalcontactsBulkContactsUpdateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostExternalcontactsBulkContactsUpdateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
