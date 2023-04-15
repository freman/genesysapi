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

// GetLanguageunderstandingMinerTopicPhraseReader is a Reader for the GetLanguageunderstandingMinerTopicPhrase structure.
type GetLanguageunderstandingMinerTopicPhraseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguageunderstandingMinerTopicPhraseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguageunderstandingMinerTopicPhraseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguageunderstandingMinerTopicPhraseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguageunderstandingMinerTopicPhraseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguageunderstandingMinerTopicPhraseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguageunderstandingMinerTopicPhraseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLanguageunderstandingMinerTopicPhraseRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguageunderstandingMinerTopicPhraseTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguageunderstandingMinerTopicPhraseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguageunderstandingMinerTopicPhraseServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguageunderstandingMinerTopicPhraseGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguageunderstandingMinerTopicPhraseOK creates a GetLanguageunderstandingMinerTopicPhraseOK with default headers values
func NewGetLanguageunderstandingMinerTopicPhraseOK() *GetLanguageunderstandingMinerTopicPhraseOK {
	return &GetLanguageunderstandingMinerTopicPhraseOK{}
}

/*
GetLanguageunderstandingMinerTopicPhraseOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLanguageunderstandingMinerTopicPhraseOK struct {
	Payload *models.MinerTopicPhrase
}

// IsSuccess returns true when this get languageunderstanding miner topic phrase o k response has a 2xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get languageunderstanding miner topic phrase o k response has a 3xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner topic phrase o k response has a 4xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languageunderstanding miner topic phrase o k response has a 5xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner topic phrase o k response a status code equal to that given
func (o *GetLanguageunderstandingMinerTopicPhraseOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLanguageunderstandingMinerTopicPhraseOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseOK) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseOK) GetPayload() *models.MinerTopicPhrase {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerTopicPhraseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MinerTopicPhrase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerTopicPhraseBadRequest creates a GetLanguageunderstandingMinerTopicPhraseBadRequest with default headers values
func NewGetLanguageunderstandingMinerTopicPhraseBadRequest() *GetLanguageunderstandingMinerTopicPhraseBadRequest {
	return &GetLanguageunderstandingMinerTopicPhraseBadRequest{}
}

/*
GetLanguageunderstandingMinerTopicPhraseBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguageunderstandingMinerTopicPhraseBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner topic phrase bad request response has a 2xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner topic phrase bad request response has a 3xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner topic phrase bad request response has a 4xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner topic phrase bad request response has a 5xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner topic phrase bad request response a status code equal to that given
func (o *GetLanguageunderstandingMinerTopicPhraseBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLanguageunderstandingMinerTopicPhraseBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerTopicPhraseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerTopicPhraseUnauthorized creates a GetLanguageunderstandingMinerTopicPhraseUnauthorized with default headers values
func NewGetLanguageunderstandingMinerTopicPhraseUnauthorized() *GetLanguageunderstandingMinerTopicPhraseUnauthorized {
	return &GetLanguageunderstandingMinerTopicPhraseUnauthorized{}
}

/*
GetLanguageunderstandingMinerTopicPhraseUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguageunderstandingMinerTopicPhraseUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner topic phrase unauthorized response has a 2xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner topic phrase unauthorized response has a 3xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner topic phrase unauthorized response has a 4xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner topic phrase unauthorized response has a 5xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner topic phrase unauthorized response a status code equal to that given
func (o *GetLanguageunderstandingMinerTopicPhraseUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLanguageunderstandingMinerTopicPhraseUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerTopicPhraseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerTopicPhraseForbidden creates a GetLanguageunderstandingMinerTopicPhraseForbidden with default headers values
func NewGetLanguageunderstandingMinerTopicPhraseForbidden() *GetLanguageunderstandingMinerTopicPhraseForbidden {
	return &GetLanguageunderstandingMinerTopicPhraseForbidden{}
}

/*
GetLanguageunderstandingMinerTopicPhraseForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguageunderstandingMinerTopicPhraseForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner topic phrase forbidden response has a 2xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner topic phrase forbidden response has a 3xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner topic phrase forbidden response has a 4xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner topic phrase forbidden response has a 5xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner topic phrase forbidden response a status code equal to that given
func (o *GetLanguageunderstandingMinerTopicPhraseForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLanguageunderstandingMinerTopicPhraseForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerTopicPhraseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerTopicPhraseNotFound creates a GetLanguageunderstandingMinerTopicPhraseNotFound with default headers values
func NewGetLanguageunderstandingMinerTopicPhraseNotFound() *GetLanguageunderstandingMinerTopicPhraseNotFound {
	return &GetLanguageunderstandingMinerTopicPhraseNotFound{}
}

/*
GetLanguageunderstandingMinerTopicPhraseNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetLanguageunderstandingMinerTopicPhraseNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner topic phrase not found response has a 2xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner topic phrase not found response has a 3xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner topic phrase not found response has a 4xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner topic phrase not found response has a 5xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner topic phrase not found response a status code equal to that given
func (o *GetLanguageunderstandingMinerTopicPhraseNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLanguageunderstandingMinerTopicPhraseNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerTopicPhraseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerTopicPhraseRequestTimeout creates a GetLanguageunderstandingMinerTopicPhraseRequestTimeout with default headers values
func NewGetLanguageunderstandingMinerTopicPhraseRequestTimeout() *GetLanguageunderstandingMinerTopicPhraseRequestTimeout {
	return &GetLanguageunderstandingMinerTopicPhraseRequestTimeout{}
}

/*
GetLanguageunderstandingMinerTopicPhraseRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLanguageunderstandingMinerTopicPhraseRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner topic phrase request timeout response has a 2xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner topic phrase request timeout response has a 3xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner topic phrase request timeout response has a 4xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner topic phrase request timeout response has a 5xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner topic phrase request timeout response a status code equal to that given
func (o *GetLanguageunderstandingMinerTopicPhraseRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetLanguageunderstandingMinerTopicPhraseRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerTopicPhraseRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge creates a GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge with default headers values
func NewGetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge() *GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge {
	return &GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge{}
}

/*
GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner topic phrase request entity too large response has a 2xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner topic phrase request entity too large response has a 3xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner topic phrase request entity too large response has a 4xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner topic phrase request entity too large response has a 5xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner topic phrase request entity too large response a status code equal to that given
func (o *GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerTopicPhraseRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType creates a GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType with default headers values
func NewGetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType() *GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType {
	return &GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType{}
}

/*
GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner topic phrase unsupported media type response has a 2xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner topic phrase unsupported media type response has a 3xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner topic phrase unsupported media type response has a 4xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner topic phrase unsupported media type response has a 5xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner topic phrase unsupported media type response a status code equal to that given
func (o *GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerTopicPhraseUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerTopicPhraseTooManyRequests creates a GetLanguageunderstandingMinerTopicPhraseTooManyRequests with default headers values
func NewGetLanguageunderstandingMinerTopicPhraseTooManyRequests() *GetLanguageunderstandingMinerTopicPhraseTooManyRequests {
	return &GetLanguageunderstandingMinerTopicPhraseTooManyRequests{}
}

/*
GetLanguageunderstandingMinerTopicPhraseTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLanguageunderstandingMinerTopicPhraseTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner topic phrase too many requests response has a 2xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner topic phrase too many requests response has a 3xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner topic phrase too many requests response has a 4xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding miner topic phrase too many requests response has a 5xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding miner topic phrase too many requests response a status code equal to that given
func (o *GetLanguageunderstandingMinerTopicPhraseTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetLanguageunderstandingMinerTopicPhraseTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerTopicPhraseTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerTopicPhraseInternalServerError creates a GetLanguageunderstandingMinerTopicPhraseInternalServerError with default headers values
func NewGetLanguageunderstandingMinerTopicPhraseInternalServerError() *GetLanguageunderstandingMinerTopicPhraseInternalServerError {
	return &GetLanguageunderstandingMinerTopicPhraseInternalServerError{}
}

/*
GetLanguageunderstandingMinerTopicPhraseInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguageunderstandingMinerTopicPhraseInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner topic phrase internal server error response has a 2xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner topic phrase internal server error response has a 3xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner topic phrase internal server error response has a 4xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languageunderstanding miner topic phrase internal server error response has a 5xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get languageunderstanding miner topic phrase internal server error response a status code equal to that given
func (o *GetLanguageunderstandingMinerTopicPhraseInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLanguageunderstandingMinerTopicPhraseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerTopicPhraseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerTopicPhraseServiceUnavailable creates a GetLanguageunderstandingMinerTopicPhraseServiceUnavailable with default headers values
func NewGetLanguageunderstandingMinerTopicPhraseServiceUnavailable() *GetLanguageunderstandingMinerTopicPhraseServiceUnavailable {
	return &GetLanguageunderstandingMinerTopicPhraseServiceUnavailable{}
}

/*
GetLanguageunderstandingMinerTopicPhraseServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguageunderstandingMinerTopicPhraseServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner topic phrase service unavailable response has a 2xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner topic phrase service unavailable response has a 3xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner topic phrase service unavailable response has a 4xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languageunderstanding miner topic phrase service unavailable response has a 5xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get languageunderstanding miner topic phrase service unavailable response a status code equal to that given
func (o *GetLanguageunderstandingMinerTopicPhraseServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetLanguageunderstandingMinerTopicPhraseServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerTopicPhraseServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingMinerTopicPhraseGatewayTimeout creates a GetLanguageunderstandingMinerTopicPhraseGatewayTimeout with default headers values
func NewGetLanguageunderstandingMinerTopicPhraseGatewayTimeout() *GetLanguageunderstandingMinerTopicPhraseGatewayTimeout {
	return &GetLanguageunderstandingMinerTopicPhraseGatewayTimeout{}
}

/*
GetLanguageunderstandingMinerTopicPhraseGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetLanguageunderstandingMinerTopicPhraseGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding miner topic phrase gateway timeout response has a 2xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding miner topic phrase gateway timeout response has a 3xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding miner topic phrase gateway timeout response has a 4xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languageunderstanding miner topic phrase gateway timeout response has a 5xx status code
func (o *GetLanguageunderstandingMinerTopicPhraseGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get languageunderstanding miner topic phrase gateway timeout response a status code equal to that given
func (o *GetLanguageunderstandingMinerTopicPhraseGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetLanguageunderstandingMinerTopicPhraseGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/miners/{minerId}/topics/{topicId}/phrases/{phraseId}][%d] getLanguageunderstandingMinerTopicPhraseGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingMinerTopicPhraseGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingMinerTopicPhraseGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
