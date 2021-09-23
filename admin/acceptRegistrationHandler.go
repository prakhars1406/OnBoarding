package admin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/config"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/factory"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/model"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/utility"
	"net/http"
	"strings"
)



func ActionHandle(datastoreClient factory.MongoClient) http.HandlerFunc {
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
		var actionRequest model.ActionRequest
		err := json.NewDecoder(request.Body).Decode(&actionRequest)
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
		if merchantId!=actionRequest.MerchantId{
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
		if strings.ToLower(actionRequest.Action)==config.REGISTER_ACTION_ACCEPT{
			actionRequest.Action="accepted"
			_, err = datastoreClient.AcceptRegistrationRequest(merchantId)
		}else if strings.ToLower(actionRequest.Action)==config.REGISTER_ACTION_REJECT{
			actionRequest.Action="rejected"
			_, err = datastoreClient.RejectRegistrationRequest(merchantId)
		}else{
			err=errors.New(fmt.Sprintf("invalid action"))
		}

		if err != nil {
			utility.HandleError(writer, http.StatusBadRequest, err)
		} else {
			err := json.NewEncoder(writer).Encode(struct {
				Message string `json:"result,omitempty"`
			}{Message:fmt.Sprintf("registration request %s",strings.ToLower(actionRequest.Action))})
			if err != nil {
				utility.HandleError(writer, http.StatusInternalServerError, err)
			}else {
				writer.WriteHeader(http.StatusOK)
			}
		}
	}
}
