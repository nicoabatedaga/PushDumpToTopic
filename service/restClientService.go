package services

import (
	"fmt"
	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
	"github.com/mercadolibre/go-meli-toolkit/restful/rest/retry"
	"net/http"
	"time"

	"github.com/mercadolibre/go-meli-toolkit/restful/rest"
)

var postByMP rest.RequestBuilder

func init() {
	h := make(http.Header)
	h.Add("X-Caller-Scopes", "admin")
	h.Add("X-Auth-Token", "")

	postByMP = rest.RequestBuilder{
		Headers:        h,
		Timeout:        3 * time.Second,
		ContentType:    rest.JSON,
		DisableTimeout: false,
		CustomPool: &rest.CustomPool{
			MaxIdleConnsPerHost: 5000,
		},
		RetryStrategy: retry.NewSimpleRetryStrategy(3, 200*time.Millisecond),
	}
}

func NewInternalServerError(msg string) apierrors.ApiError {
	var err = fmt.Errorf(msg)
	return apierrors.NewInternalServerApiError(fmt.Sprintf(msg), err)
}

var RbPostMP = func(url string) *rest.Response {
	return postByMP.Post(url, nil)
}

func PostMsg(id, _type, site_id, user_id string) error {
	url := fmt.Sprintf("https://prod-topic-news-generator_transfer-consumers.furyapps.io/transfer_consumers/ba_producer/id/%v/type/%v/site/%v/user/%v", id, _type, site_id, user_id)
	r := RbPostMP(url)
	if r == nil || r.Response == nil || r.Err != nil || r.StatusCode != http.StatusOK {
		return NewInternalServerError(fmt.Sprintf("Post fail %s. Unknown Error - %v", id, r.Err))
	}
	return nil
}
