package admin

import (
	"encoding/json"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/config"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/factory"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/utility"
	"net/http"
)



func GetAllMerchants(datastoreClient factory.MongoClient) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer utility.PanicHandler(writer, request)
		writer.Header().Set(config.CONTENT_TYPE_HEADER, config.APPLICATION_JSON)

		allMerchant, err := datastoreClient.GetAllMerchants()
		if err != nil {
			utility.HandleError(writer, http.StatusBadRequest, err)
		}  else {
			err := json.NewEncoder(writer).Encode(allMerchant)
			if err != nil {
				utility.HandleError(writer, http.StatusInternalServerError, err)
			} else {
				writer.WriteHeader(http.StatusOK)
			}
		}
	}
}
