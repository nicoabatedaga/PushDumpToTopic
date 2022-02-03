package services

/*
import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
	"github.com/mercadolibre/go-meli-toolkit/restful/rest/retry"
	"net/http"
	"time"

	"github.com/mercadolibre/go-meli-toolkit/restful/rest"
)

var getByML rest.RequestBuilder

func init() {
	h := make(http.Header)
	h.Add("X-Caller-Scopes", "admin")

	getByML = rest.RequestBuilder{
		Timeout:        3 * time.Second,
		ContentType:    rest.JSON,
		DisableTimeout: false,
		CustomPool: &rest.CustomPool{
			MaxIdleConnsPerHost: 100,
		},
		RetryStrategy: retry.NewSimpleRetryStrategy(3, 100*time.Millisecond),
		Headers:       copyMapHeaders(h),
	}

}

func copyMapHeaders(originalMap http.Header) http.Header {
	targetMap := make(http.Header)
	for key, value := range originalMap {
		targetMap[key] = value
	}
	return targetMap
}

func NewInternalServerError(msg string) apierrors.ApiError {
	var err = fmt.Errorf(msg)
	return apierrors.NewInternalServerApiError(fmt.Sprintf(msg), err)
}

var RbGetML = func(url string) *rest.Response { return getByML.Get(url) }

type User struct {
	UserID         int64 `json:"id"`
	Identification struct {
		Number string `json:"number"`
	} `json:"identification"`
	Company struct {
		Identification string `json:"identification"`
	} `json:"company"`
}

func (us *User) UnMarshal(rb []byte) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	result := new(User)
	if err := json.Unmarshal(rb, result); err != nil {
		return err
	}
	*us = *result
	return nil
}

func GetUser(id string) (*User, error) {
	url := fmt.Sprintf("https://internal-api.mercadolibre.com/users/%v?&caller.scopes=admin", id)
	//fmt.Println(url)
	r := RbGetML(url)
	if r == nil {
		return nil, apierrors.NewInternalServerApiError(fmt.Sprintf("Get fail %s. Unknown Error", id), nil)
	}
	if r.Response == nil || r.Err != nil {
		return nil, apierrors.NewInternalServerApiError(fmt.Sprintf("Get fail %s. Error:%v", id, r.Err), r.Err)
	}
	if r.StatusCode != http.StatusOK {
		return nil, NewInternalServerError(fmt.Sprintf("StatusCode:%v Get fail: %s. status:%s", r.StatusCode, id, r.Status))
	}
	us := &User{}
	if e := us.UnMarshal(r.Bytes()); e != nil {
		if e.Error() == "not_found" {
			return nil, apierrors.NewNotFoundApiError("Withdrawal not found")
		}
		return nil, apierrors.NewInternalServerApiError(fmt.Sprintf("UnMarshal withdraw error (id: %s)", id), e)
	}
	return us, nil
}
*/
