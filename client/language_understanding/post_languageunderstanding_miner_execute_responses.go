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

// PostLanguageunderstandingMinerExecuteReader is a Reader for the PostLanguageunderstandingMinerExecute structure.
type PostLanguageunderstandingMinerExecuteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLanguageunderstandingMinerExecuteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLanguageunderstandingMinerExecuteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostLanguageunderstandingMinerExecuteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLanguageunderstandingMinerExecuteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLanguageunderstandingMinerExecuteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLanguageunderstandingMinerExecuteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLanguageunderstandingMinerExecuteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLanguageunderstandingMinerExecuteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLanguageunderstandingMinerExecuteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLanguageunderstandingMinerExecuteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLanguageunderstandingMinerExecuteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLanguageunderstandingMinerExecuteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLanguageunderstandingMinerExecuteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLanguageunderstandingMinerExecuteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLanguageunderstandingMinerExecuteOK creates a PostLanguageunderstandingMinerExecuteOK with default headers values
func NewPostLanguageunderstandingMinerExecuteOK() *PostLanguageunderstandingMinerExecuteOK {
	return &PostLanguageunderstandingMinerExecuteOK{}
}

/*
PostLanguageunderstandingMinerExecuteOK describes a response with status code 200, with default header values.

successful operation
*/
type PostLanguageunderstandingMinerExecuteOK struct {
	Payload *models.Miner
}

// IsSuccess returns true when this post languageunderstanding miner execute o k response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post languageunderstanding miner execute o k response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute o k response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding miner execute o k response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner execute o k response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostLanguageunderstandingMinerExecuteOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteOK  %+v", 200, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteOK) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteOK  %+v", 200, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteOK) GetPayload() *models.Miner {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Miner)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerExecuteAccepted creates a PostLanguageunderstandingMinerExecuteAccepted with default headers values
func NewPostLanguageunderstandingMinerExecuteAccepted() *PostLanguageunderstandingMinerExecuteAccepted {
	return &PostLanguageunderstandingMinerExecuteAccepted{}
}

/*
PostLanguageunderstandingMinerExecuteAccepted describes a response with status code 202, with default header values.

Processing request
*/
type PostLanguageunderstandingMinerExecuteAccepted struct {
	Payload *models.Miner
}

// IsSuccess returns true when this post languageunderstanding miner execute accepted response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post languageunderstanding miner execute accepted response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute accepted response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding miner execute accepted response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner execute accepted response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostLanguageunderstandingMinerExecuteAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteAccepted  %+v", 202, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteAccepted  %+v", 202, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteAccepted) GetPayload() *models.Miner {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Miner)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerExecuteBadRequest creates a PostLanguageunderstandingMinerExecuteBadRequest with default headers values
func NewPostLanguageunderstandingMinerExecuteBadRequest() *PostLanguageunderstandingMinerExecuteBadRequest {
	return &PostLanguageunderstandingMinerExecuteBadRequest{}
}

/*
PostLanguageunderstandingMinerExecuteBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLanguageunderstandingMinerExecuteBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner execute bad request response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner execute bad request response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute bad request response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner execute bad request response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner execute bad request response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostLanguageunderstandingMinerExecuteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteBadRequest  %+v", 400, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteBadRequest  %+v", 400, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerExecuteUnauthorized creates a PostLanguageunderstandingMinerExecuteUnauthorized with default headers values
func NewPostLanguageunderstandingMinerExecuteUnauthorized() *PostLanguageunderstandingMinerExecuteUnauthorized {
	return &PostLanguageunderstandingMinerExecuteUnauthorized{}
}

/*
PostLanguageunderstandingMinerExecuteUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLanguageunderstandingMinerExecuteUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner execute unauthorized response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner execute unauthorized response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute unauthorized response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner execute unauthorized response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner execute unauthorized response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostLanguageunderstandingMinerExecuteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerExecuteForbidden creates a PostLanguageunderstandingMinerExecuteForbidden with default headers values
func NewPostLanguageunderstandingMinerExecuteForbidden() *PostLanguageunderstandingMinerExecuteForbidden {
	return &PostLanguageunderstandingMinerExecuteForbidden{}
}

/*
PostLanguageunderstandingMinerExecuteForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostLanguageunderstandingMinerExecuteForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner execute forbidden response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner execute forbidden response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute forbidden response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner execute forbidden response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner execute forbidden response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostLanguageunderstandingMinerExecuteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteForbidden  %+v", 403, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteForbidden  %+v", 403, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerExecuteNotFound creates a PostLanguageunderstandingMinerExecuteNotFound with default headers values
func NewPostLanguageunderstandingMinerExecuteNotFound() *PostLanguageunderstandingMinerExecuteNotFound {
	return &PostLanguageunderstandingMinerExecuteNotFound{}
}

/*
PostLanguageunderstandingMinerExecuteNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostLanguageunderstandingMinerExecuteNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner execute not found response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner execute not found response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute not found response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner execute not found response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner execute not found response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostLanguageunderstandingMinerExecuteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteNotFound  %+v", 404, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteNotFound  %+v", 404, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerExecuteRequestTimeout creates a PostLanguageunderstandingMinerExecuteRequestTimeout with default headers values
func NewPostLanguageunderstandingMinerExecuteRequestTimeout() *PostLanguageunderstandingMinerExecuteRequestTimeout {
	return &PostLanguageunderstandingMinerExecuteRequestTimeout{}
}

/*
PostLanguageunderstandingMinerExecuteRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLanguageunderstandingMinerExecuteRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner execute request timeout response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner execute request timeout response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute request timeout response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner execute request timeout response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner execute request timeout response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostLanguageunderstandingMinerExecuteRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerExecuteRequestEntityTooLarge creates a PostLanguageunderstandingMinerExecuteRequestEntityTooLarge with default headers values
func NewPostLanguageunderstandingMinerExecuteRequestEntityTooLarge() *PostLanguageunderstandingMinerExecuteRequestEntityTooLarge {
	return &PostLanguageunderstandingMinerExecuteRequestEntityTooLarge{}
}

/*
PostLanguageunderstandingMinerExecuteRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostLanguageunderstandingMinerExecuteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner execute request entity too large response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner execute request entity too large response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute request entity too large response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner execute request entity too large response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner execute request entity too large response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostLanguageunderstandingMinerExecuteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerExecuteUnsupportedMediaType creates a PostLanguageunderstandingMinerExecuteUnsupportedMediaType with default headers values
func NewPostLanguageunderstandingMinerExecuteUnsupportedMediaType() *PostLanguageunderstandingMinerExecuteUnsupportedMediaType {
	return &PostLanguageunderstandingMinerExecuteUnsupportedMediaType{}
}

/*
PostLanguageunderstandingMinerExecuteUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLanguageunderstandingMinerExecuteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner execute unsupported media type response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner execute unsupported media type response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute unsupported media type response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner execute unsupported media type response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner execute unsupported media type response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostLanguageunderstandingMinerExecuteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerExecuteTooManyRequests creates a PostLanguageunderstandingMinerExecuteTooManyRequests with default headers values
func NewPostLanguageunderstandingMinerExecuteTooManyRequests() *PostLanguageunderstandingMinerExecuteTooManyRequests {
	return &PostLanguageunderstandingMinerExecuteTooManyRequests{}
}

/*
PostLanguageunderstandingMinerExecuteTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLanguageunderstandingMinerExecuteTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner execute too many requests response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner execute too many requests response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute too many requests response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner execute too many requests response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner execute too many requests response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostLanguageunderstandingMinerExecuteTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerExecuteInternalServerError creates a PostLanguageunderstandingMinerExecuteInternalServerError with default headers values
func NewPostLanguageunderstandingMinerExecuteInternalServerError() *PostLanguageunderstandingMinerExecuteInternalServerError {
	return &PostLanguageunderstandingMinerExecuteInternalServerError{}
}

/*
PostLanguageunderstandingMinerExecuteInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLanguageunderstandingMinerExecuteInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner execute internal server error response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner execute internal server error response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute internal server error response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding miner execute internal server error response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post languageunderstanding miner execute internal server error response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostLanguageunderstandingMinerExecuteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerExecuteServiceUnavailable creates a PostLanguageunderstandingMinerExecuteServiceUnavailable with default headers values
func NewPostLanguageunderstandingMinerExecuteServiceUnavailable() *PostLanguageunderstandingMinerExecuteServiceUnavailable {
	return &PostLanguageunderstandingMinerExecuteServiceUnavailable{}
}

/*
PostLanguageunderstandingMinerExecuteServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLanguageunderstandingMinerExecuteServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner execute service unavailable response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner execute service unavailable response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute service unavailable response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding miner execute service unavailable response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post languageunderstanding miner execute service unavailable response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostLanguageunderstandingMinerExecuteServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerExecuteGatewayTimeout creates a PostLanguageunderstandingMinerExecuteGatewayTimeout with default headers values
func NewPostLanguageunderstandingMinerExecuteGatewayTimeout() *PostLanguageunderstandingMinerExecuteGatewayTimeout {
	return &PostLanguageunderstandingMinerExecuteGatewayTimeout{}
}

/*
PostLanguageunderstandingMinerExecuteGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostLanguageunderstandingMinerExecuteGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner execute gateway timeout response has a 2xx status code
func (o *PostLanguageunderstandingMinerExecuteGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner execute gateway timeout response has a 3xx status code
func (o *PostLanguageunderstandingMinerExecuteGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner execute gateway timeout response has a 4xx status code
func (o *PostLanguageunderstandingMinerExecuteGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding miner execute gateway timeout response has a 5xx status code
func (o *PostLanguageunderstandingMinerExecuteGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post languageunderstanding miner execute gateway timeout response a status code equal to that given
func (o *PostLanguageunderstandingMinerExecuteGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostLanguageunderstandingMinerExecuteGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/execute][%d] postLanguageunderstandingMinerExecuteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLanguageunderstandingMinerExecuteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerExecuteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
