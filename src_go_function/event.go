package main

type Event struct {
	HTTPMethod                      string              `json:"httpMethod"`
	Headers                         map[string]string   `json:"headers"`
	Path                            string              `json:"path"`
	MultiValueHeaders               map[string][]string `json:"multiValueHeaders"`
	QueryStringParameters           map[string]string   `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string `json:"multiValueQueryStringParameters"`
	RequestContext                  RequestContext      `json:"requestContext"`
	Body                            string              `json:"body"`
	IsBase64Encoded                 bool                `json:"isBase64Encoded"`
}

type RequestContext struct {
	Identity         Identity `json:"identity"`
	HTTPMethod       string   `json:"httpMethod"`
	RequestID        string   `json:"requestId"`
	RequestTime      string   `json:"requestTime"`
	RequestTimeEpoch int64    `json:"requestTimeEpoch"`
}

type Identity struct {
	SourceIP  string `json:"sourceIp"`
	UserAgent string `json:"userAgent"`
}
