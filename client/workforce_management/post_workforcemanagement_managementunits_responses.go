// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostWorkforcemanagementManagementunitsReader is a Reader for the PostWorkforcemanagementManagementunits structure.
type PostWorkforcemanagementManagementunitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementManagementunitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementManagementunitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostWorkforcemanagementManagementunitsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementManagementunitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementManagementunitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementManagementunitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementManagementunitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementManagementunitsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementManagementunitsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementManagementunitsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementManagementunitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementManagementunitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementManagementunitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementManagementunitsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementManagementunitsOK creates a PostWorkforcemanagementManagementunitsOK with default headers values
func NewPostWorkforcemanagementManagementunitsOK() *PostWorkforcemanagementManagementunitsOK {
	return &PostWorkforcemanagementManagementunitsOK{}
}

/*
PostWorkforcemanagementManagementunitsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostWorkforcemanagementManagementunitsOK struct {
	Payload *models.ManagementUnit
}

// IsSuccess returns true when this post workforcemanagement managementunits o k response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post workforcemanagement managementunits o k response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits o k response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement managementunits o k response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement managementunits o k response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostWorkforcemanagementManagementunitsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsOK) GetPayload() *models.ManagementUnit {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsCreated creates a PostWorkforcemanagementManagementunitsCreated with default headers values
func NewPostWorkforcemanagementManagementunitsCreated() *PostWorkforcemanagementManagementunitsCreated {
	return &PostWorkforcemanagementManagementunitsCreated{}
}

/*
PostWorkforcemanagementManagementunitsCreated describes a response with status code 201, with default header values.

The management unit was successfully created
*/
type PostWorkforcemanagementManagementunitsCreated struct {
	Payload *models.ManagementUnit
}

// IsSuccess returns true when this post workforcemanagement managementunits created response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post workforcemanagement managementunits created response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits created response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement managementunits created response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement managementunits created response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostWorkforcemanagementManagementunitsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsCreated  %+v", 201, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsCreated  %+v", 201, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsCreated) GetPayload() *models.ManagementUnit {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsBadRequest creates a PostWorkforcemanagementManagementunitsBadRequest with default headers values
func NewPostWorkforcemanagementManagementunitsBadRequest() *PostWorkforcemanagementManagementunitsBadRequest {
	return &PostWorkforcemanagementManagementunitsBadRequest{}
}

/*
PostWorkforcemanagementManagementunitsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementManagementunitsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement managementunits bad request response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement managementunits bad request response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits bad request response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement managementunits bad request response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement managementunits bad request response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostWorkforcemanagementManagementunitsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsUnauthorized creates a PostWorkforcemanagementManagementunitsUnauthorized with default headers values
func NewPostWorkforcemanagementManagementunitsUnauthorized() *PostWorkforcemanagementManagementunitsUnauthorized {
	return &PostWorkforcemanagementManagementunitsUnauthorized{}
}

/*
PostWorkforcemanagementManagementunitsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementManagementunitsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement managementunits unauthorized response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement managementunits unauthorized response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits unauthorized response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement managementunits unauthorized response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement managementunits unauthorized response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostWorkforcemanagementManagementunitsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsForbidden creates a PostWorkforcemanagementManagementunitsForbidden with default headers values
func NewPostWorkforcemanagementManagementunitsForbidden() *PostWorkforcemanagementManagementunitsForbidden {
	return &PostWorkforcemanagementManagementunitsForbidden{}
}

/*
PostWorkforcemanagementManagementunitsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementManagementunitsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement managementunits forbidden response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement managementunits forbidden response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits forbidden response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement managementunits forbidden response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement managementunits forbidden response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostWorkforcemanagementManagementunitsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsNotFound creates a PostWorkforcemanagementManagementunitsNotFound with default headers values
func NewPostWorkforcemanagementManagementunitsNotFound() *PostWorkforcemanagementManagementunitsNotFound {
	return &PostWorkforcemanagementManagementunitsNotFound{}
}

/*
PostWorkforcemanagementManagementunitsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementManagementunitsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement managementunits not found response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement managementunits not found response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits not found response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement managementunits not found response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement managementunits not found response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostWorkforcemanagementManagementunitsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsRequestTimeout creates a PostWorkforcemanagementManagementunitsRequestTimeout with default headers values
func NewPostWorkforcemanagementManagementunitsRequestTimeout() *PostWorkforcemanagementManagementunitsRequestTimeout {
	return &PostWorkforcemanagementManagementunitsRequestTimeout{}
}

/*
PostWorkforcemanagementManagementunitsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementManagementunitsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement managementunits request timeout response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement managementunits request timeout response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits request timeout response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement managementunits request timeout response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement managementunits request timeout response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostWorkforcemanagementManagementunitsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsRequestEntityTooLarge creates a PostWorkforcemanagementManagementunitsRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementManagementunitsRequestEntityTooLarge() *PostWorkforcemanagementManagementunitsRequestEntityTooLarge {
	return &PostWorkforcemanagementManagementunitsRequestEntityTooLarge{}
}

/*
PostWorkforcemanagementManagementunitsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostWorkforcemanagementManagementunitsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement managementunits request entity too large response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement managementunits request entity too large response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits request entity too large response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement managementunits request entity too large response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement managementunits request entity too large response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostWorkforcemanagementManagementunitsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsUnsupportedMediaType creates a PostWorkforcemanagementManagementunitsUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementManagementunitsUnsupportedMediaType() *PostWorkforcemanagementManagementunitsUnsupportedMediaType {
	return &PostWorkforcemanagementManagementunitsUnsupportedMediaType{}
}

/*
PostWorkforcemanagementManagementunitsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementManagementunitsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement managementunits unsupported media type response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement managementunits unsupported media type response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits unsupported media type response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement managementunits unsupported media type response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement managementunits unsupported media type response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostWorkforcemanagementManagementunitsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsTooManyRequests creates a PostWorkforcemanagementManagementunitsTooManyRequests with default headers values
func NewPostWorkforcemanagementManagementunitsTooManyRequests() *PostWorkforcemanagementManagementunitsTooManyRequests {
	return &PostWorkforcemanagementManagementunitsTooManyRequests{}
}

/*
PostWorkforcemanagementManagementunitsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementManagementunitsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement managementunits too many requests response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement managementunits too many requests response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits too many requests response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement managementunits too many requests response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement managementunits too many requests response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostWorkforcemanagementManagementunitsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsInternalServerError creates a PostWorkforcemanagementManagementunitsInternalServerError with default headers values
func NewPostWorkforcemanagementManagementunitsInternalServerError() *PostWorkforcemanagementManagementunitsInternalServerError {
	return &PostWorkforcemanagementManagementunitsInternalServerError{}
}

/*
PostWorkforcemanagementManagementunitsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementManagementunitsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement managementunits internal server error response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement managementunits internal server error response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits internal server error response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement managementunits internal server error response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement managementunits internal server error response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostWorkforcemanagementManagementunitsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsServiceUnavailable creates a PostWorkforcemanagementManagementunitsServiceUnavailable with default headers values
func NewPostWorkforcemanagementManagementunitsServiceUnavailable() *PostWorkforcemanagementManagementunitsServiceUnavailable {
	return &PostWorkforcemanagementManagementunitsServiceUnavailable{}
}

/*
PostWorkforcemanagementManagementunitsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementManagementunitsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement managementunits service unavailable response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement managementunits service unavailable response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits service unavailable response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement managementunits service unavailable response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement managementunits service unavailable response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostWorkforcemanagementManagementunitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsGatewayTimeout creates a PostWorkforcemanagementManagementunitsGatewayTimeout with default headers values
func NewPostWorkforcemanagementManagementunitsGatewayTimeout() *PostWorkforcemanagementManagementunitsGatewayTimeout {
	return &PostWorkforcemanagementManagementunitsGatewayTimeout{}
}

/*
PostWorkforcemanagementManagementunitsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostWorkforcemanagementManagementunitsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement managementunits gateway timeout response has a 2xx status code
func (o *PostWorkforcemanagementManagementunitsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement managementunits gateway timeout response has a 3xx status code
func (o *PostWorkforcemanagementManagementunitsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement managementunits gateway timeout response has a 4xx status code
func (o *PostWorkforcemanagementManagementunitsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement managementunits gateway timeout response has a 5xx status code
func (o *PostWorkforcemanagementManagementunitsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement managementunits gateway timeout response a status code equal to that given
func (o *PostWorkforcemanagementManagementunitsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostWorkforcemanagementManagementunitsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
