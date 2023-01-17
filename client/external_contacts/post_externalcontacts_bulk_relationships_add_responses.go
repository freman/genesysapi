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

// PostExternalcontactsBulkRelationshipsAddReader is a Reader for the PostExternalcontactsBulkRelationshipsAdd structure.
type PostExternalcontactsBulkRelationshipsAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkRelationshipsAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkRelationshipsAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkRelationshipsAddBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkRelationshipsAddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkRelationshipsAddForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkRelationshipsAddNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkRelationshipsAddRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkRelationshipsAddUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsBulkRelationshipsAddUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkRelationshipsAddTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkRelationshipsAddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkRelationshipsAddServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkRelationshipsAddGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkRelationshipsAddOK creates a PostExternalcontactsBulkRelationshipsAddOK with default headers values
func NewPostExternalcontactsBulkRelationshipsAddOK() *PostExternalcontactsBulkRelationshipsAddOK {
	return &PostExternalcontactsBulkRelationshipsAddOK{}
}

/*
PostExternalcontactsBulkRelationshipsAddOK describes a response with status code 200, with default header values.

successful operation
*/
type PostExternalcontactsBulkRelationshipsAddOK struct {
	Payload *models.BulkRelationshipsResponse
}

// IsSuccess returns true when this post externalcontacts bulk relationships add o k response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post externalcontacts bulk relationships add o k response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add o k response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk relationships add o k response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk relationships add o k response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostExternalcontactsBulkRelationshipsAddOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddOK) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddOK) GetPayload() *models.BulkRelationshipsResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkRelationshipsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddBadRequest creates a PostExternalcontactsBulkRelationshipsAddBadRequest with default headers values
func NewPostExternalcontactsBulkRelationshipsAddBadRequest() *PostExternalcontactsBulkRelationshipsAddBadRequest {
	return &PostExternalcontactsBulkRelationshipsAddBadRequest{}
}

/*
PostExternalcontactsBulkRelationshipsAddBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkRelationshipsAddBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk relationships add bad request response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk relationships add bad request response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add bad request response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk relationships add bad request response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk relationships add bad request response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostExternalcontactsBulkRelationshipsAddBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddUnauthorized creates a PostExternalcontactsBulkRelationshipsAddUnauthorized with default headers values
func NewPostExternalcontactsBulkRelationshipsAddUnauthorized() *PostExternalcontactsBulkRelationshipsAddUnauthorized {
	return &PostExternalcontactsBulkRelationshipsAddUnauthorized{}
}

/*
PostExternalcontactsBulkRelationshipsAddUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkRelationshipsAddUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk relationships add unauthorized response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk relationships add unauthorized response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add unauthorized response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk relationships add unauthorized response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk relationships add unauthorized response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostExternalcontactsBulkRelationshipsAddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddForbidden creates a PostExternalcontactsBulkRelationshipsAddForbidden with default headers values
func NewPostExternalcontactsBulkRelationshipsAddForbidden() *PostExternalcontactsBulkRelationshipsAddForbidden {
	return &PostExternalcontactsBulkRelationshipsAddForbidden{}
}

/*
PostExternalcontactsBulkRelationshipsAddForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkRelationshipsAddForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk relationships add forbidden response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk relationships add forbidden response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add forbidden response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk relationships add forbidden response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk relationships add forbidden response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostExternalcontactsBulkRelationshipsAddForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddNotFound creates a PostExternalcontactsBulkRelationshipsAddNotFound with default headers values
func NewPostExternalcontactsBulkRelationshipsAddNotFound() *PostExternalcontactsBulkRelationshipsAddNotFound {
	return &PostExternalcontactsBulkRelationshipsAddNotFound{}
}

/*
PostExternalcontactsBulkRelationshipsAddNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkRelationshipsAddNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk relationships add not found response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk relationships add not found response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add not found response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk relationships add not found response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk relationships add not found response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostExternalcontactsBulkRelationshipsAddNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddRequestTimeout creates a PostExternalcontactsBulkRelationshipsAddRequestTimeout with default headers values
func NewPostExternalcontactsBulkRelationshipsAddRequestTimeout() *PostExternalcontactsBulkRelationshipsAddRequestTimeout {
	return &PostExternalcontactsBulkRelationshipsAddRequestTimeout{}
}

/*
PostExternalcontactsBulkRelationshipsAddRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkRelationshipsAddRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk relationships add request timeout response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk relationships add request timeout response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add request timeout response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk relationships add request timeout response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk relationships add request timeout response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge creates a PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge() *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge {
	return &PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge{}
}

/*
PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk relationships add request entity too large response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk relationships add request entity too large response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add request entity too large response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk relationships add request entity too large response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk relationships add request entity too large response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddUnsupportedMediaType creates a PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkRelationshipsAddUnsupportedMediaType() *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType {
	return &PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType{}
}

/*
PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk relationships add unsupported media type response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk relationships add unsupported media type response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add unsupported media type response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk relationships add unsupported media type response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk relationships add unsupported media type response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddUnprocessableEntity creates a PostExternalcontactsBulkRelationshipsAddUnprocessableEntity with default headers values
func NewPostExternalcontactsBulkRelationshipsAddUnprocessableEntity() *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity {
	return &PostExternalcontactsBulkRelationshipsAddUnprocessableEntity{}
}

/*
PostExternalcontactsBulkRelationshipsAddUnprocessableEntity describes a response with status code 422, with default header values.

PostExternalcontactsBulkRelationshipsAddUnprocessableEntity post externalcontacts bulk relationships add unprocessable entity
*/
type PostExternalcontactsBulkRelationshipsAddUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk relationships add unprocessable entity response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk relationships add unprocessable entity response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add unprocessable entity response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk relationships add unprocessable entity response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk relationships add unprocessable entity response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddTooManyRequests creates a PostExternalcontactsBulkRelationshipsAddTooManyRequests with default headers values
func NewPostExternalcontactsBulkRelationshipsAddTooManyRequests() *PostExternalcontactsBulkRelationshipsAddTooManyRequests {
	return &PostExternalcontactsBulkRelationshipsAddTooManyRequests{}
}

/*
PostExternalcontactsBulkRelationshipsAddTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkRelationshipsAddTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk relationships add too many requests response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk relationships add too many requests response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add too many requests response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk relationships add too many requests response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk relationships add too many requests response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostExternalcontactsBulkRelationshipsAddTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddInternalServerError creates a PostExternalcontactsBulkRelationshipsAddInternalServerError with default headers values
func NewPostExternalcontactsBulkRelationshipsAddInternalServerError() *PostExternalcontactsBulkRelationshipsAddInternalServerError {
	return &PostExternalcontactsBulkRelationshipsAddInternalServerError{}
}

/*
PostExternalcontactsBulkRelationshipsAddInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkRelationshipsAddInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk relationships add internal server error response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk relationships add internal server error response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add internal server error response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk relationships add internal server error response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk relationships add internal server error response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostExternalcontactsBulkRelationshipsAddInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddServiceUnavailable creates a PostExternalcontactsBulkRelationshipsAddServiceUnavailable with default headers values
func NewPostExternalcontactsBulkRelationshipsAddServiceUnavailable() *PostExternalcontactsBulkRelationshipsAddServiceUnavailable {
	return &PostExternalcontactsBulkRelationshipsAddServiceUnavailable{}
}

/*
PostExternalcontactsBulkRelationshipsAddServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkRelationshipsAddServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk relationships add service unavailable response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk relationships add service unavailable response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add service unavailable response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk relationships add service unavailable response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk relationships add service unavailable response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostExternalcontactsBulkRelationshipsAddServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddGatewayTimeout creates a PostExternalcontactsBulkRelationshipsAddGatewayTimeout with default headers values
func NewPostExternalcontactsBulkRelationshipsAddGatewayTimeout() *PostExternalcontactsBulkRelationshipsAddGatewayTimeout {
	return &PostExternalcontactsBulkRelationshipsAddGatewayTimeout{}
}

/*
PostExternalcontactsBulkRelationshipsAddGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostExternalcontactsBulkRelationshipsAddGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk relationships add gateway timeout response has a 2xx status code
func (o *PostExternalcontactsBulkRelationshipsAddGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk relationships add gateway timeout response has a 3xx status code
func (o *PostExternalcontactsBulkRelationshipsAddGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk relationships add gateway timeout response has a 4xx status code
func (o *PostExternalcontactsBulkRelationshipsAddGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk relationships add gateway timeout response has a 5xx status code
func (o *PostExternalcontactsBulkRelationshipsAddGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk relationships add gateway timeout response a status code equal to that given
func (o *PostExternalcontactsBulkRelationshipsAddGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostExternalcontactsBulkRelationshipsAddGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
