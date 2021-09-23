package admin

import (
	"encoding/json"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/config"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/factory"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/model"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/utility"
	"net/http"
)

func OnBoardHandler(datastoreClient factory.MongoClient) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer utility.PanicHandler(writer, request)
		writer.Header().Set(config.CONTENT_TYPE_HEADER, config.APPLICATION_JSON)
		var onBoardingDetails model.OnBoardDetailsRequest
		err := json.NewDecoder(request.Body).Decode(&onBoardingDetails)
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
		customerRole, err := datastoreClient.GetCustomerRole("","")
		if err != nil {
			utility.HandleError(writer, http.StatusBadRequest, err)
		} else {
			err := json.NewEncoder(writer).Encode(struct {
				CustomerRole model.CustomerRoleResponse `json:"result,omitempty"`
			}{CustomerRole: model.CustomerRoleResponse{Username: customerRole.Username,Password: customerRole.Password,RouteStats: customerRole.RouteStats,
				RouteRecommendations: customerRole.RouteRecommendations,RouteCharts: customerRole.RouteCharts,
				ClientsAccess: customerRole.ClientsAccess,Platform: customerRole.Platform,Email: customerRole.Email,Address: customerRole.Address,
				AmpSelection: customerRole.AmpSelection,PhoneNumber: customerRole.PhoneNumber,OnboardingStatus: customerRole.OnBoard}})
			if err != nil {
				utility.HandleError(writer, http.StatusInternalServerError, err)
			} else {
				writer.WriteHeader(http.StatusOK)
			}
		}
	}
}