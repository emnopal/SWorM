package agent

import (
	"github.com/SWorM/v2/templater"
	"github.com/go-resty/resty/v2"
)

type Request struct {
	*resty.Request
}

func parseMapWithEnv(values map[string]string, envs map[string]string) map[string]string {
	for k, v := range values {
		parsedValue, _ := templater.ParseRegex(v, envs)
		values[k] = parsedValue
	}
	return values
}

// param
func (request *Request) SetupParam(pathParam map[string]string, queryParam map[string]string, envs map[string]string) *Request {
	pathParam = parseMapWithEnv(pathParam, envs)
	queryParam = parseMapWithEnv(queryParam, envs)
	request = request.inurlParam(pathParam)
	request = request.posturlParam(queryParam)
	return request
}
func (request *Request) inurlParam(pathParam map[string]string) *Request {
	request.SetPathParams(pathParam)
	return request
}
func (request *Request) posturlParam(queryParam map[string]string) *Request {
	request.SetQueryParams(queryParam)
	return request
}

// header
func (request *Request) SetupHeader(header map[string]string, envs map[string]string) *Request {
	header = parseMapWithEnv(header, envs)
	request.SetHeaders(header)
	return request
}

// body
func (request *Request) SetupBody(body string, envs map[string]string) *Request {
	parsedBody, _ := templater.ParseTemplate(body, envs)
	request.SetBody(parsedBody)
	return request
}
