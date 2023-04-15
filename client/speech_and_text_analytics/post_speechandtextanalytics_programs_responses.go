// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostSpeechandtextanalyticsProgramsReader is a Reader for the PostSpeechandtextanalyticsPrograms structure.
type PostSpeechandtextanalyticsProgramsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSpeechandtextanalyticsProgramsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSpeechandtextanalyticsProgramsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSpeechandtextanalyticsProgramsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSpeechandtextanalyticsProgramsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSpeechandtextanalyticsProgramsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostSpeechandtextanalyticsProgramsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostSpeechandtextanalyticsProgramsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostSpeechandtextanalyticsProgramsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostSpeechandtextanalyticsProgramsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostSpeechandtextanalyticsProgramsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostSpeechandtextanalyticsProgramsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSpeechandtextanalyticsProgramsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostSpeechandtextanalyticsProgramsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostSpeechandtextanalyticsProgramsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostSpeechandtextanalyticsProgramsOK creates a PostSpeechandtextanalyticsProgramsOK with default headers values
func NewPostSpeechandtextanalyticsProgramsOK() *PostSpeechandtextanalyticsProgramsOK {
	return &PostSpeechandtextanalyticsProgramsOK{}
}

/*
PostSpeechandtextanalyticsProgramsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostSpeechandtextanalyticsProgramsOK struct {
	Payload *models.Program
}

// IsSuccess returns true when this post speechandtextanalytics programs o k response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post speechandtextanalytics programs o k response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs o k response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post speechandtextanalytics programs o k response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post speechandtextanalytics programs o k response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostSpeechandtextanalyticsProgramsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsOK  %+v", 200, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsOK  %+v", 200, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsOK) GetPayload() *models.Program {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Program)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsProgramsBadRequest creates a PostSpeechandtextanalyticsProgramsBadRequest with default headers values
func NewPostSpeechandtextanalyticsProgramsBadRequest() *PostSpeechandtextanalyticsProgramsBadRequest {
	return &PostSpeechandtextanalyticsProgramsBadRequest{}
}

/*
PostSpeechandtextanalyticsProgramsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostSpeechandtextanalyticsProgramsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post speechandtextanalytics programs bad request response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post speechandtextanalytics programs bad request response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs bad request response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post speechandtextanalytics programs bad request response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post speechandtextanalytics programs bad request response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostSpeechandtextanalyticsProgramsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsBadRequest  %+v", 400, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsBadRequest  %+v", 400, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsProgramsUnauthorized creates a PostSpeechandtextanalyticsProgramsUnauthorized with default headers values
func NewPostSpeechandtextanalyticsProgramsUnauthorized() *PostSpeechandtextanalyticsProgramsUnauthorized {
	return &PostSpeechandtextanalyticsProgramsUnauthorized{}
}

/*
PostSpeechandtextanalyticsProgramsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostSpeechandtextanalyticsProgramsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post speechandtextanalytics programs unauthorized response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post speechandtextanalytics programs unauthorized response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs unauthorized response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post speechandtextanalytics programs unauthorized response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post speechandtextanalytics programs unauthorized response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostSpeechandtextanalyticsProgramsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsProgramsForbidden creates a PostSpeechandtextanalyticsProgramsForbidden with default headers values
func NewPostSpeechandtextanalyticsProgramsForbidden() *PostSpeechandtextanalyticsProgramsForbidden {
	return &PostSpeechandtextanalyticsProgramsForbidden{}
}

/*
PostSpeechandtextanalyticsProgramsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostSpeechandtextanalyticsProgramsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post speechandtextanalytics programs forbidden response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post speechandtextanalytics programs forbidden response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs forbidden response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post speechandtextanalytics programs forbidden response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post speechandtextanalytics programs forbidden response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostSpeechandtextanalyticsProgramsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsForbidden  %+v", 403, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsForbidden  %+v", 403, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsProgramsNotFound creates a PostSpeechandtextanalyticsProgramsNotFound with default headers values
func NewPostSpeechandtextanalyticsProgramsNotFound() *PostSpeechandtextanalyticsProgramsNotFound {
	return &PostSpeechandtextanalyticsProgramsNotFound{}
}

/*
PostSpeechandtextanalyticsProgramsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostSpeechandtextanalyticsProgramsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post speechandtextanalytics programs not found response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post speechandtextanalytics programs not found response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs not found response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post speechandtextanalytics programs not found response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post speechandtextanalytics programs not found response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostSpeechandtextanalyticsProgramsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsNotFound  %+v", 404, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsNotFound  %+v", 404, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsProgramsRequestTimeout creates a PostSpeechandtextanalyticsProgramsRequestTimeout with default headers values
func NewPostSpeechandtextanalyticsProgramsRequestTimeout() *PostSpeechandtextanalyticsProgramsRequestTimeout {
	return &PostSpeechandtextanalyticsProgramsRequestTimeout{}
}

/*
PostSpeechandtextanalyticsProgramsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostSpeechandtextanalyticsProgramsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post speechandtextanalytics programs request timeout response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post speechandtextanalytics programs request timeout response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs request timeout response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post speechandtextanalytics programs request timeout response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post speechandtextanalytics programs request timeout response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostSpeechandtextanalyticsProgramsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsProgramsConflict creates a PostSpeechandtextanalyticsProgramsConflict with default headers values
func NewPostSpeechandtextanalyticsProgramsConflict() *PostSpeechandtextanalyticsProgramsConflict {
	return &PostSpeechandtextanalyticsProgramsConflict{}
}

/*
PostSpeechandtextanalyticsProgramsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostSpeechandtextanalyticsProgramsConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post speechandtextanalytics programs conflict response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post speechandtextanalytics programs conflict response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs conflict response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post speechandtextanalytics programs conflict response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post speechandtextanalytics programs conflict response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostSpeechandtextanalyticsProgramsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsConflict  %+v", 409, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsConflict  %+v", 409, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsProgramsRequestEntityTooLarge creates a PostSpeechandtextanalyticsProgramsRequestEntityTooLarge with default headers values
func NewPostSpeechandtextanalyticsProgramsRequestEntityTooLarge() *PostSpeechandtextanalyticsProgramsRequestEntityTooLarge {
	return &PostSpeechandtextanalyticsProgramsRequestEntityTooLarge{}
}

/*
PostSpeechandtextanalyticsProgramsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostSpeechandtextanalyticsProgramsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post speechandtextanalytics programs request entity too large response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post speechandtextanalytics programs request entity too large response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs request entity too large response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post speechandtextanalytics programs request entity too large response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post speechandtextanalytics programs request entity too large response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostSpeechandtextanalyticsProgramsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsProgramsUnsupportedMediaType creates a PostSpeechandtextanalyticsProgramsUnsupportedMediaType with default headers values
func NewPostSpeechandtextanalyticsProgramsUnsupportedMediaType() *PostSpeechandtextanalyticsProgramsUnsupportedMediaType {
	return &PostSpeechandtextanalyticsProgramsUnsupportedMediaType{}
}

/*
PostSpeechandtextanalyticsProgramsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostSpeechandtextanalyticsProgramsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post speechandtextanalytics programs unsupported media type response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post speechandtextanalytics programs unsupported media type response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs unsupported media type response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post speechandtextanalytics programs unsupported media type response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post speechandtextanalytics programs unsupported media type response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostSpeechandtextanalyticsProgramsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsProgramsTooManyRequests creates a PostSpeechandtextanalyticsProgramsTooManyRequests with default headers values
func NewPostSpeechandtextanalyticsProgramsTooManyRequests() *PostSpeechandtextanalyticsProgramsTooManyRequests {
	return &PostSpeechandtextanalyticsProgramsTooManyRequests{}
}

/*
PostSpeechandtextanalyticsProgramsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostSpeechandtextanalyticsProgramsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post speechandtextanalytics programs too many requests response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post speechandtextanalytics programs too many requests response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs too many requests response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post speechandtextanalytics programs too many requests response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post speechandtextanalytics programs too many requests response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostSpeechandtextanalyticsProgramsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsProgramsInternalServerError creates a PostSpeechandtextanalyticsProgramsInternalServerError with default headers values
func NewPostSpeechandtextanalyticsProgramsInternalServerError() *PostSpeechandtextanalyticsProgramsInternalServerError {
	return &PostSpeechandtextanalyticsProgramsInternalServerError{}
}

/*
PostSpeechandtextanalyticsProgramsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostSpeechandtextanalyticsProgramsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post speechandtextanalytics programs internal server error response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post speechandtextanalytics programs internal server error response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs internal server error response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post speechandtextanalytics programs internal server error response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post speechandtextanalytics programs internal server error response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostSpeechandtextanalyticsProgramsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsProgramsServiceUnavailable creates a PostSpeechandtextanalyticsProgramsServiceUnavailable with default headers values
func NewPostSpeechandtextanalyticsProgramsServiceUnavailable() *PostSpeechandtextanalyticsProgramsServiceUnavailable {
	return &PostSpeechandtextanalyticsProgramsServiceUnavailable{}
}

/*
PostSpeechandtextanalyticsProgramsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostSpeechandtextanalyticsProgramsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post speechandtextanalytics programs service unavailable response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post speechandtextanalytics programs service unavailable response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs service unavailable response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post speechandtextanalytics programs service unavailable response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post speechandtextanalytics programs service unavailable response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostSpeechandtextanalyticsProgramsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsProgramsGatewayTimeout creates a PostSpeechandtextanalyticsProgramsGatewayTimeout with default headers values
func NewPostSpeechandtextanalyticsProgramsGatewayTimeout() *PostSpeechandtextanalyticsProgramsGatewayTimeout {
	return &PostSpeechandtextanalyticsProgramsGatewayTimeout{}
}

/*
PostSpeechandtextanalyticsProgramsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostSpeechandtextanalyticsProgramsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post speechandtextanalytics programs gateway timeout response has a 2xx status code
func (o *PostSpeechandtextanalyticsProgramsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post speechandtextanalytics programs gateway timeout response has a 3xx status code
func (o *PostSpeechandtextanalyticsProgramsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post speechandtextanalytics programs gateway timeout response has a 4xx status code
func (o *PostSpeechandtextanalyticsProgramsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post speechandtextanalytics programs gateway timeout response has a 5xx status code
func (o *PostSpeechandtextanalyticsProgramsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post speechandtextanalytics programs gateway timeout response a status code equal to that given
func (o *PostSpeechandtextanalyticsProgramsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostSpeechandtextanalyticsProgramsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/programs][%d] postSpeechandtextanalyticsProgramsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostSpeechandtextanalyticsProgramsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsProgramsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
