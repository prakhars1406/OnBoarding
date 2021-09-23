package admin

import (
	"encoding/json"
	"github.com/prakhars1406/OnBoarding/config"
	"github.com/prakhars1406/OnBoarding/factory"
	"github.com/prakhars1406/OnBoarding/model"
	"github.com/prakhars1406/OnBoarding/utility"
	"net/http"
)



func RegisterHandler(datastoreClient factory.MongoClient) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer utility.PanicHandler(writer, request)
		writer.Header().Set(config.CONTENT_TYPE_HEADER, config.APPLICATION_JSON)
		var registrationDetails model.RegisterationDetailsRequest
		err := json.NewDecoder(request.Body).Decode(&registrationDetails)
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
		registrationResponse, err := datastoreClient.AddNewMerchant(registrationDetails)
		if err != nil {
			utility.HandleError(writer, http.StatusBadRequest, err)
		}  else {
			err := json.NewEncoder(writer).Encode(struct {
				CustomerRole model.CustomerRoleResponse `json:"result,omitempty"`
			}{CustomerRole: model.CustomerRoleResponse{Username:registrationResponse.Username,Role:registrationResponse.Role,Password: registrationResponse.Password,RouteStats: registrationResponse.RouteStats,
				RouteRecommendations: registrationResponse.RouteRecommendations,RouteCharts: registrationResponse.RouteCharts,ClientsAccess: registrationResponse.ClientsAccess,
			Platform: registrationResponse.Platform,Email: registrationResponse.Email,Address: registrationResponse.Address,AmpSelection: registrationResponse.AmpSelection,
			PhoneNumber: registrationResponse.PhoneNumber,OnboardingStatus: registrationResponse.OnBoard,MerchantId: registrationResponse.MerchantId}})
			if err != nil {
				utility.HandleError(writer, http.StatusInternalServerError, err)
			} else {
				writer.WriteHeader(http.StatusOK)
			}
		}
	}
}
