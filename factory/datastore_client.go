package factory

import (
	"github.com/prakhars1406/OnBoarding/model"
)

type MongoClient interface {
	GetCustomerRole(userName, password string) (model.CustomerRole, error)
	AddNewMerchant(registerData model.RegisterationDetailsRequest) (model.CustomerRole, error)
	GetMaxMerchantId() (int, error)
	CheckUsernameExists(userName string) (bool, error)
	CreateRegisterationTransaction(userName,pass,merchantId,platform,email,address,ampSelection,phoneNumber,defaultImage,defaultImageSource,language string) (bool, error)
	GetAllMerchants() ([]model.AllMerchantResponse, error)
	AddNewMerchantConfig(merchantConfig model.MerchantConfigDetails) (bool, error)
	AcceptRegistrationRequest(merchantId string) (bool, error)
	RejectRegistrationRequest(merchantId string) (bool, error)
	UpdateRegisterStatusInCustomerRole(merchantId,registerStatus string) error
	AddNewPageType(merchantId string,pageTypes []model.PageTypeDetails) ([]model.PageTypeDetailsResponse, error)
	CheckPageTypeCount() (int, error)
	CheckStrategyAndPageTypeExists(strategy,pageType string) (bool, error)
	CheckPageTypeExists(pageType string) (bool, error)
	GetAllPageRecommendations() ([]model.AllPageRecommendations, error)
	AddNewPageRecommendations(pageRecommendations []model.AllPageRecommendations) ([]model.AllPageRecommendationsResponse, error)
	AddNewMerchantPages(merchantId string,merchantpages []model.MerchantPagesDetails) ([]model.MerchantPagesResponse, error)
	GetAllRouteInfo(merchantId string) (model.RouteInfoResponse, error)
	//OnBoardNewMerchant(onBoardingDetails model.OnBoardDetailsRequest) (model.RegisterationDetailsResponse, error)
}
