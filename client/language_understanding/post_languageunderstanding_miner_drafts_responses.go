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

// PostLanguageunderstandingMinerDraftsReader is a Reader for the PostLanguageunderstandingMinerDrafts structure.
type PostLanguageunderstandingMinerDraftsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLanguageunderstandingMinerDraftsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLanguageunderstandingMinerDraftsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostLanguageunderstandingMinerDraftsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLanguageunderstandingMinerDraftsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLanguageunderstandingMinerDraftsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLanguageunderstandingMinerDraftsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLanguageunderstandingMinerDraftsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLanguageunderstandingMinerDraftsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLanguageunderstandingMinerDraftsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLanguageunderstandingMinerDraftsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLanguageunderstandingMinerDraftsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLanguageunderstandingMinerDraftsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLanguageunderstandingMinerDraftsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLanguageunderstandingMinerDraftsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLanguageunderstandingMinerDraftsOK creates a PostLanguageunderstandingMinerDraftsOK with default headers values
func NewPostLanguageunderstandingMinerDraftsOK() *PostLanguageunderstandingMinerDraftsOK {
	return &PostLanguageunderstandingMinerDraftsOK{}
}

/*
PostLanguageunderstandingMinerDraftsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostLanguageunderstandingMinerDraftsOK struct {
	Payload *models.Draft
}

// IsSuccess returns true when this post languageunderstanding miner drafts o k response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post languageunderstanding miner drafts o k response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts o k response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding miner drafts o k response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner drafts o k response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostLanguageunderstandingMinerDraftsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsOK  %+v", 200, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsOK  %+v", 200, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsOK) GetPayload() *models.Draft {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Draft)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerDraftsCreated creates a PostLanguageunderstandingMinerDraftsCreated with default headers values
func NewPostLanguageunderstandingMinerDraftsCreated() *PostLanguageunderstandingMinerDraftsCreated {
	return &PostLanguageunderstandingMinerDraftsCreated{}
}

/*
PostLanguageunderstandingMinerDraftsCreated describes a response with status code 201, with default header values.

Draft created successfully
*/
type PostLanguageunderstandingMinerDraftsCreated struct {
	Payload *models.Draft
}

// IsSuccess returns true when this post languageunderstanding miner drafts created response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post languageunderstanding miner drafts created response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts created response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding miner drafts created response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner drafts created response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostLanguageunderstandingMinerDraftsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsCreated  %+v", 201, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsCreated  %+v", 201, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsCreated) GetPayload() *models.Draft {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Draft)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerDraftsBadRequest creates a PostLanguageunderstandingMinerDraftsBadRequest with default headers values
func NewPostLanguageunderstandingMinerDraftsBadRequest() *PostLanguageunderstandingMinerDraftsBadRequest {
	return &PostLanguageunderstandingMinerDraftsBadRequest{}
}

/*
PostLanguageunderstandingMinerDraftsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLanguageunderstandingMinerDraftsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner drafts bad request response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner drafts bad request response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts bad request response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner drafts bad request response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner drafts bad request response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostLanguageunderstandingMinerDraftsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerDraftsUnauthorized creates a PostLanguageunderstandingMinerDraftsUnauthorized with default headers values
func NewPostLanguageunderstandingMinerDraftsUnauthorized() *PostLanguageunderstandingMinerDraftsUnauthorized {
	return &PostLanguageunderstandingMinerDraftsUnauthorized{}
}

/*
PostLanguageunderstandingMinerDraftsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLanguageunderstandingMinerDraftsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner drafts unauthorized response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner drafts unauthorized response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts unauthorized response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner drafts unauthorized response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner drafts unauthorized response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostLanguageunderstandingMinerDraftsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerDraftsForbidden creates a PostLanguageunderstandingMinerDraftsForbidden with default headers values
func NewPostLanguageunderstandingMinerDraftsForbidden() *PostLanguageunderstandingMinerDraftsForbidden {
	return &PostLanguageunderstandingMinerDraftsForbidden{}
}

/*
PostLanguageunderstandingMinerDraftsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostLanguageunderstandingMinerDraftsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner drafts forbidden response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner drafts forbidden response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts forbidden response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner drafts forbidden response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner drafts forbidden response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostLanguageunderstandingMinerDraftsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsForbidden  %+v", 403, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsForbidden  %+v", 403, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerDraftsNotFound creates a PostLanguageunderstandingMinerDraftsNotFound with default headers values
func NewPostLanguageunderstandingMinerDraftsNotFound() *PostLanguageunderstandingMinerDraftsNotFound {
	return &PostLanguageunderstandingMinerDraftsNotFound{}
}

/*
PostLanguageunderstandingMinerDraftsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostLanguageunderstandingMinerDraftsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner drafts not found response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner drafts not found response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts not found response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner drafts not found response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner drafts not found response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostLanguageunderstandingMinerDraftsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsNotFound  %+v", 404, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsNotFound  %+v", 404, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerDraftsRequestTimeout creates a PostLanguageunderstandingMinerDraftsRequestTimeout with default headers values
func NewPostLanguageunderstandingMinerDraftsRequestTimeout() *PostLanguageunderstandingMinerDraftsRequestTimeout {
	return &PostLanguageunderstandingMinerDraftsRequestTimeout{}
}

/*
PostLanguageunderstandingMinerDraftsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLanguageunderstandingMinerDraftsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner drafts request timeout response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner drafts request timeout response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts request timeout response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner drafts request timeout response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner drafts request timeout response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostLanguageunderstandingMinerDraftsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerDraftsRequestEntityTooLarge creates a PostLanguageunderstandingMinerDraftsRequestEntityTooLarge with default headers values
func NewPostLanguageunderstandingMinerDraftsRequestEntityTooLarge() *PostLanguageunderstandingMinerDraftsRequestEntityTooLarge {
	return &PostLanguageunderstandingMinerDraftsRequestEntityTooLarge{}
}

/*
PostLanguageunderstandingMinerDraftsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostLanguageunderstandingMinerDraftsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner drafts request entity too large response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner drafts request entity too large response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts request entity too large response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner drafts request entity too large response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner drafts request entity too large response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostLanguageunderstandingMinerDraftsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerDraftsUnsupportedMediaType creates a PostLanguageunderstandingMinerDraftsUnsupportedMediaType with default headers values
func NewPostLanguageunderstandingMinerDraftsUnsupportedMediaType() *PostLanguageunderstandingMinerDraftsUnsupportedMediaType {
	return &PostLanguageunderstandingMinerDraftsUnsupportedMediaType{}
}

/*
PostLanguageunderstandingMinerDraftsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLanguageunderstandingMinerDraftsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner drafts unsupported media type response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner drafts unsupported media type response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts unsupported media type response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner drafts unsupported media type response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner drafts unsupported media type response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostLanguageunderstandingMinerDraftsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerDraftsTooManyRequests creates a PostLanguageunderstandingMinerDraftsTooManyRequests with default headers values
func NewPostLanguageunderstandingMinerDraftsTooManyRequests() *PostLanguageunderstandingMinerDraftsTooManyRequests {
	return &PostLanguageunderstandingMinerDraftsTooManyRequests{}
}

/*
PostLanguageunderstandingMinerDraftsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLanguageunderstandingMinerDraftsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner drafts too many requests response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner drafts too many requests response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts too many requests response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding miner drafts too many requests response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding miner drafts too many requests response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostLanguageunderstandingMinerDraftsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerDraftsInternalServerError creates a PostLanguageunderstandingMinerDraftsInternalServerError with default headers values
func NewPostLanguageunderstandingMinerDraftsInternalServerError() *PostLanguageunderstandingMinerDraftsInternalServerError {
	return &PostLanguageunderstandingMinerDraftsInternalServerError{}
}

/*
PostLanguageunderstandingMinerDraftsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLanguageunderstandingMinerDraftsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner drafts internal server error response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner drafts internal server error response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts internal server error response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding miner drafts internal server error response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post languageunderstanding miner drafts internal server error response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostLanguageunderstandingMinerDraftsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerDraftsServiceUnavailable creates a PostLanguageunderstandingMinerDraftsServiceUnavailable with default headers values
func NewPostLanguageunderstandingMinerDraftsServiceUnavailable() *PostLanguageunderstandingMinerDraftsServiceUnavailable {
	return &PostLanguageunderstandingMinerDraftsServiceUnavailable{}
}

/*
PostLanguageunderstandingMinerDraftsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLanguageunderstandingMinerDraftsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner drafts service unavailable response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner drafts service unavailable response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts service unavailable response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding miner drafts service unavailable response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post languageunderstanding miner drafts service unavailable response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostLanguageunderstandingMinerDraftsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerDraftsGatewayTimeout creates a PostLanguageunderstandingMinerDraftsGatewayTimeout with default headers values
func NewPostLanguageunderstandingMinerDraftsGatewayTimeout() *PostLanguageunderstandingMinerDraftsGatewayTimeout {
	return &PostLanguageunderstandingMinerDraftsGatewayTimeout{}
}

/*
PostLanguageunderstandingMinerDraftsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostLanguageunderstandingMinerDraftsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding miner drafts gateway timeout response has a 2xx status code
func (o *PostLanguageunderstandingMinerDraftsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding miner drafts gateway timeout response has a 3xx status code
func (o *PostLanguageunderstandingMinerDraftsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding miner drafts gateway timeout response has a 4xx status code
func (o *PostLanguageunderstandingMinerDraftsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding miner drafts gateway timeout response has a 5xx status code
func (o *PostLanguageunderstandingMinerDraftsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post languageunderstanding miner drafts gateway timeout response a status code equal to that given
func (o *PostLanguageunderstandingMinerDraftsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostLanguageunderstandingMinerDraftsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/drafts][%d] postLanguageunderstandingMinerDraftsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLanguageunderstandingMinerDraftsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerDraftsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
