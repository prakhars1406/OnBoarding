package admin

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/config"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/factory"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/model"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/utility"
	"net/http"
	"strings"
)



func AddNewPageType(datastoreClient factory.MongoClient) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer utility.PanicHandler(writer, request)
		writer.Header().Set(config.CONTENT_TYPE_HEADER, config.APPLICATION_JSON)
		merchantId :=""
		for k, v := range mux.Vars(request) {
			if strings.EqualFold(k, config.MERCHANT_ID) {
				merchantId = v
			}
		}
		if len(merchantId)==0{
			err := json.NewEncoder(writer).Encode(struct {
				Result  bool   `json:"result"`
				Error string `json:"err,omitempty"`
			}{Result: false, Error: "invalid merchant id"})
			if err != nil {
				utility.HandleError(writer, http.StatusBadRequest, err)
			} else {
				writer.WriteHeader(http.StatusBadRequest)
			}
			return
		}
		var pageType []model.PageTypeDetails
		err := json.NewDecoder(request.Body).Decode(&pageType)
		if err != nil {
			err := json.NewEncoder(writer).Encode(struct {
				Result  bool   `json:"result"`
				Error string `json:"err,omitempty"`
			}{Result: false, Error: err.Error()})
			if err != nil {
				utility.HandleError(writer, http.StatusBadRequest, err)
			} else {
				writer.WriteHeader(http.StatusBadRequest)
			}
			return
		}
		response, err := datastoreClient.AddNewPageType(merchantId,pageType)
		if err != nil {
			utility.HandleError(writer, http.StatusBadRequest, err)
		}  else {
			err := json.NewEncoder(writer).Encode(response)
			if err != nil {
				utility.HandleError(writer, http.StatusInternalServerError, err)
			} else {
				writer.WriteHeader(http.StatusOK)
			}
		}
	}
}