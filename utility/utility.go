package utility

import (
	"errors"
	. "gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/config"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/model"
	"net/http"
)

var MailIdMatcher = "[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+.[A-Za-z]{1,}."
var HtmlTagMatcher = "(<.*?>|&quot)"
var Nbsp = "\u00A0"


var merchantDetails *model.MerchantDetails = nil

func HandleError(writer http.ResponseWriter, code int, err error) {

	writer.WriteHeader(code)
	writer.Write([]byte(err.Error()))
}

type PairList []Pair

type Pair struct {
	Key   string
	Value int
}

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func PanicHandler(w http.ResponseWriter, r *http.Request) {
	if r := recover(); r != nil {
		var err error
		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = errors.New("Unknown panic")
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func GetMongoServer() string {
	return GetAppConfiguration().MONGO_SERVER
}

func GetDataStore() string {
	return GetAppConfiguration().DATASTORE
}

