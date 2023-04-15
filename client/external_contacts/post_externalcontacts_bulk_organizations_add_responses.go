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

// PostExternalcontactsBulkOrganizationsAddReader is a Reader for the PostExternalcontactsBulkOrganizationsAdd structure.
type PostExternalcontactsBulkOrganizationsAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkOrganizationsAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkOrganizationsAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkOrganizationsAddBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkOrganizationsAddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkOrganizationsAddForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkOrganizationsAddNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkOrganizationsAddRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkOrganizationsAddUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsBulkOrganizationsAddUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkOrganizationsAddTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkOrganizationsAddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkOrganizationsAddServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkOrganizationsAddGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkOrganizationsAddOK creates a PostExternalcontactsBulkOrganizationsAddOK with default headers values
func NewPostExternalcontactsBulkOrganizationsAddOK() *PostExternalcontactsBulkOrganizationsAddOK {
	return &PostExternalcontactsBulkOrganizationsAddOK{}
}

/*
PostExternalcontactsBulkOrganizationsAddOK describes a response with status code 200, with default header values.

successful operation
*/
type PostExternalcontactsBulkOrganizationsAddOK struct {
	Payload *models.BulkOrganizationsResponse
}

// IsSuccess returns true when this post externalcontacts bulk organizations add o k response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post externalcontacts bulk organizations add o k response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add o k response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk organizations add o k response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations add o k response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostExternalcontactsBulkOrganizationsAddOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddOK) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddOK) GetPayload() *models.BulkOrganizationsResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkOrganizationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddBadRequest creates a PostExternalcontactsBulkOrganizationsAddBadRequest with default headers values
func NewPostExternalcontactsBulkOrganizationsAddBadRequest() *PostExternalcontactsBulkOrganizationsAddBadRequest {
	return &PostExternalcontactsBulkOrganizationsAddBadRequest{}
}

/*
PostExternalcontactsBulkOrganizationsAddBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkOrganizationsAddBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations add bad request response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations add bad request response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add bad request response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations add bad request response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations add bad request response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostExternalcontactsBulkOrganizationsAddBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddUnauthorized creates a PostExternalcontactsBulkOrganizationsAddUnauthorized with default headers values
func NewPostExternalcontactsBulkOrganizationsAddUnauthorized() *PostExternalcontactsBulkOrganizationsAddUnauthorized {
	return &PostExternalcontactsBulkOrganizationsAddUnauthorized{}
}

/*
PostExternalcontactsBulkOrganizationsAddUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkOrganizationsAddUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations add unauthorized response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations add unauthorized response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add unauthorized response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations add unauthorized response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations add unauthorized response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostExternalcontactsBulkOrganizationsAddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddForbidden creates a PostExternalcontactsBulkOrganizationsAddForbidden with default headers values
func NewPostExternalcontactsBulkOrganizationsAddForbidden() *PostExternalcontactsBulkOrganizationsAddForbidden {
	return &PostExternalcontactsBulkOrganizationsAddForbidden{}
}

/*
PostExternalcontactsBulkOrganizationsAddForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkOrganizationsAddForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations add forbidden response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations add forbidden response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add forbidden response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations add forbidden response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations add forbidden response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostExternalcontactsBulkOrganizationsAddForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddNotFound creates a PostExternalcontactsBulkOrganizationsAddNotFound with default headers values
func NewPostExternalcontactsBulkOrganizationsAddNotFound() *PostExternalcontactsBulkOrganizationsAddNotFound {
	return &PostExternalcontactsBulkOrganizationsAddNotFound{}
}

/*
PostExternalcontactsBulkOrganizationsAddNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkOrganizationsAddNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations add not found response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations add not found response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add not found response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations add not found response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations add not found response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostExternalcontactsBulkOrganizationsAddNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddRequestTimeout creates a PostExternalcontactsBulkOrganizationsAddRequestTimeout with default headers values
func NewPostExternalcontactsBulkOrganizationsAddRequestTimeout() *PostExternalcontactsBulkOrganizationsAddRequestTimeout {
	return &PostExternalcontactsBulkOrganizationsAddRequestTimeout{}
}

/*
PostExternalcontactsBulkOrganizationsAddRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkOrganizationsAddRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations add request timeout response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations add request timeout response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add request timeout response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations add request timeout response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations add request timeout response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostExternalcontactsBulkOrganizationsAddRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge creates a PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge() *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge {
	return &PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge{}
}

/*
PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations add request entity too large response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations add request entity too large response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add request entity too large response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations add request entity too large response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations add request entity too large response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddUnsupportedMediaType creates a PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkOrganizationsAddUnsupportedMediaType() *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType {
	return &PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType{}
}

/*
PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations add unsupported media type response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations add unsupported media type response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add unsupported media type response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations add unsupported media type response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations add unsupported media type response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddUnprocessableEntity creates a PostExternalcontactsBulkOrganizationsAddUnprocessableEntity with default headers values
func NewPostExternalcontactsBulkOrganizationsAddUnprocessableEntity() *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity {
	return &PostExternalcontactsBulkOrganizationsAddUnprocessableEntity{}
}

/*
PostExternalcontactsBulkOrganizationsAddUnprocessableEntity describes a response with status code 422, with default header values.

PostExternalcontactsBulkOrganizationsAddUnprocessableEntity post externalcontacts bulk organizations add unprocessable entity
*/
type PostExternalcontactsBulkOrganizationsAddUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations add unprocessable entity response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations add unprocessable entity response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add unprocessable entity response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations add unprocessable entity response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations add unprocessable entity response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddTooManyRequests creates a PostExternalcontactsBulkOrganizationsAddTooManyRequests with default headers values
func NewPostExternalcontactsBulkOrganizationsAddTooManyRequests() *PostExternalcontactsBulkOrganizationsAddTooManyRequests {
	return &PostExternalcontactsBulkOrganizationsAddTooManyRequests{}
}

/*
PostExternalcontactsBulkOrganizationsAddTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkOrganizationsAddTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations add too many requests response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations add too many requests response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add too many requests response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk organizations add too many requests response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk organizations add too many requests response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostExternalcontactsBulkOrganizationsAddTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddInternalServerError creates a PostExternalcontactsBulkOrganizationsAddInternalServerError with default headers values
func NewPostExternalcontactsBulkOrganizationsAddInternalServerError() *PostExternalcontactsBulkOrganizationsAddInternalServerError {
	return &PostExternalcontactsBulkOrganizationsAddInternalServerError{}
}

/*
PostExternalcontactsBulkOrganizationsAddInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkOrganizationsAddInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations add internal server error response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations add internal server error response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add internal server error response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk organizations add internal server error response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk organizations add internal server error response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostExternalcontactsBulkOrganizationsAddInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddServiceUnavailable creates a PostExternalcontactsBulkOrganizationsAddServiceUnavailable with default headers values
func NewPostExternalcontactsBulkOrganizationsAddServiceUnavailable() *PostExternalcontactsBulkOrganizationsAddServiceUnavailable {
	return &PostExternalcontactsBulkOrganizationsAddServiceUnavailable{}
}

/*
PostExternalcontactsBulkOrganizationsAddServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkOrganizationsAddServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations add service unavailable response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations add service unavailable response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add service unavailable response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk organizations add service unavailable response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk organizations add service unavailable response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostExternalcontactsBulkOrganizationsAddServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkOrganizationsAddGatewayTimeout creates a PostExternalcontactsBulkOrganizationsAddGatewayTimeout with default headers values
func NewPostExternalcontactsBulkOrganizationsAddGatewayTimeout() *PostExternalcontactsBulkOrganizationsAddGatewayTimeout {
	return &PostExternalcontactsBulkOrganizationsAddGatewayTimeout{}
}

/*
PostExternalcontactsBulkOrganizationsAddGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostExternalcontactsBulkOrganizationsAddGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk organizations add gateway timeout response has a 2xx status code
func (o *PostExternalcontactsBulkOrganizationsAddGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk organizations add gateway timeout response has a 3xx status code
func (o *PostExternalcontactsBulkOrganizationsAddGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk organizations add gateway timeout response has a 4xx status code
func (o *PostExternalcontactsBulkOrganizationsAddGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk organizations add gateway timeout response has a 5xx status code
func (o *PostExternalcontactsBulkOrganizationsAddGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk organizations add gateway timeout response a status code equal to that given
func (o *PostExternalcontactsBulkOrganizationsAddGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostExternalcontactsBulkOrganizationsAddGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/organizations/add][%d] postExternalcontactsBulkOrganizationsAddGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkOrganizationsAddGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkOrganizationsAddGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
