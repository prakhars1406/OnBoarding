package factory

import (
	"errors"
	"fmt"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/config"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/logger"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"sort"
	"strconv"
	"strings"
)

type MongoClientImpl struct {
	mongoServer string
	session     *mgo.Session
}

func (client *MongoClientImpl) GetMaxMerchantId() (int, error) {
	collection := fmt.Sprintf(config.MONGO_MERCHANT_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	result := []bson.M{}
	err := c.Find(bson.M{}).Select(bson.M{"merchant_id": 1, config.MONGO_ID: 0}).All(&result)
	if err != nil {
		return -1, err
	}
	merchant_id:=[]int{}
	for _, v := range result {
		var resultStruct model.MongoMerchant
		bsonBytes, _ := bson.Marshal(v)
		bson.Unmarshal(bsonBytes, &resultStruct)
		resultStructInt, err := strconv.Atoi(resultStruct.Merchant_id)
		if err != nil {
				continue
		}
		merchant_id=append(merchant_id,resultStructInt)
	}
	sort.Ints(merchant_id)
	return merchant_id[len(merchant_id)-1], nil
}
func (client *MongoClientImpl) CheckUsernameExists(userName string) (bool, error) {
	collection := fmt.Sprintf(config.MONGO_CUSTOMERS_ROLE_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	count, err := c.Find(bson.M{"username": strings.ToLower(userName)}).Count()
	if err != nil {
		return true, err
	}
	if count!=0{
		return true,nil
	}else{
		return false,nil
	}
}
func (client *MongoClientImpl) CheckMerchantExists(merchantId string,collection string,registerationStep string) (bool, error) {
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	result := []bson.M{}
	err := c.Find(bson.M{"merchant_id": merchantId}).Select(bson.M{}).All(&result)
	if err != nil {
		return false, errors.New("<mongo> Unable to query collection")
	}
	if len(result) == 0 {
		return false, errors.New("merchant_id doesn't match")
	}
	var resultStruct model.CustomerRole
	bsonBytes, _ := bson.Marshal(result[0])
	bson.Unmarshal(bsonBytes, &resultStruct)
	if resultStruct.OnBoard==registerationStep{
		return true,nil
	}else{
		return false,nil
	}
}
func (client *MongoClientImpl) CheckMerchantConfigNotExist(merchantId string) (bool, error) {
	collection := fmt.Sprintf(config.MONGO_MERCHANT_CONFIG_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	count, err := c.Find(bson.M{"merchant_id": merchantId}).Count()
	if err != nil {
		return true, err
	}
	if count!=0{
		return true,nil
	}else{
		return false,nil
	}
}
func (client *MongoClientImpl) CreateRegisterationTransaction(userName,pass,merchantId,platform,email,address,ampSelection,phoneNumber,defaultImage,defaultImageSource,language string) (bool, error) {
	collection := fmt.Sprintf("txns")
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	db := session.DB(config.MONGO_SIMS_DATABASE)
	c := db.C(collection)
	runner := txn.NewRunner(c)

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return false, err
	}

	type M map[string]interface{}
	ops := []txn.Op{{
		C:      config.MONGO_CUSTOMERS_COLLECTION,
		Id:bson.NewObjectId(),
		Insert: M{"username": strings.ToLower(userName),
			"password":string(hash),
			"role":strings.ToLower(userName),
			"merchant_id": merchantId},
	}, {
		C:      config.MONGO_CUSTOMERS_ROLE_COLLECTION,
		Id:     bson.NewObjectId(),
		Insert: M{"username": strings.ToLower(userName),
			"password":string(hash),
			"role":strings.ToUpper(userName),
			"route_stats":config.ROUTE_STATS,
			"route_recommendations": "/"+config.RECOMMENDATIONS,
			"route_charts":config.ROUTE_CHART ,
			"platform":platform,
			"email": email,
			"address": address,
			"ampSelection": strings.ToLower(ampSelection),
			"phoneNumber": phoneNumber,
			"clients_access": merchantId,
			"onboard":config.REGISTER_REQUESTED,
			"language":language,
			"merchant_id": merchantId},
	}, {
		C:      config.MONGO_MERCHANT_COLLECTION,
		Id:     bson.NewObjectId(),
		Insert: M{
			"merchant_id": merchantId,
			"merchant_name":strings.ToUpper(userName),
			"article_sims_content_key": config.ARTICLE_SIMS_CONTENT_KEY,
			"article_sims_recency_key": config.ARTICLE_SIMS_RECENCY_KEY,
			"video_sims_content_key": config.VIDEO_SIMS_CONTENT_KEY,
			"video_sims_recency_key": config.VIDEO_SIMS_RECENCY_KEY,
			"video_accountId": 0,
			"video_clientId": "",
			"video_secretKey": ""},
	}, {
		C:      config.MONGO_MERCHANT_PLATFORM_COLLECTION,
		Id:     bson.NewObjectId(),
		Insert: M{
			"merchant_id": merchantId,
			"platform":platform},
	},{
		C:      config.MONGO_ROUTE_INFO_COLLECTION,
		Id:     bson.NewObjectId(),
		Insert: M{
			"merchant_id": merchantId,
			"route":"/"+merchantId,
			"name": userName,
			"component": config.RECOMMENDATIONS,
			"description": config.DESCRIPTION+userName,
			"images": config.IMAGES,
			"videoSrc": "",
			"defaultImage": defaultImage,
			"defaultSlug": platform,
			"defaultImageSource":defaultImageSource},
	}}

	err = runner.Run(ops,"", nil)
	if err != nil {
		return false, err
	}
	return true,nil
}
func (client *MongoClientImpl) AddNewMerchant(registerData model.RegisterationDetailsRequest) (model.CustomerRole, error) {
	exists,err:=client.CheckUsernameExists(registerData.Username)
	if err != nil {
		return model.CustomerRole{}, err
	}
	if exists{
		return model.CustomerRole{}, errors.New(fmt.Sprintf("<mongo> Merchant with username:%s already present", registerData.Username))
	}
	max_merchant_id,err:=client.GetMaxMerchantId()
	if err != nil {
		return model.CustomerRole{}, err
	}
	success,err:=client.CreateRegisterationTransaction(registerData.Username,registerData.Password,strconv.Itoa(max_merchant_id+1),registerData.Platform,
		registerData.Email,registerData.Address,registerData.AmpSelection,registerData.PhoneNumber,registerData.DefaultImage,registerData.DefaultImageSource,registerData.Language)
	if err != nil {
		return model.CustomerRole{}, err
	}
	if success {
		customerRole, err := client.GetCustomerRole(registerData.Username,registerData.Password)
		if err != nil {
			return model.CustomerRole{}, err
		}else{
			return customerRole, nil
		}
	}
	return model.CustomerRole{}, err
}
func (client *MongoClientImpl) GetCustomerRole(userName, password string) (model.CustomerRole, error) {
	collection := fmt.Sprintf(config.MONGO_CUSTOMERS_ROLE_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	result := []bson.M{}
	err := c.Find(bson.M{"username": strings.ToLower(userName)}).Select(bson.M{}).All(&result)
	if err != nil {
		return model.CustomerRole{}, errors.New("<mongo> Unable to query collection")
	}
	if len(result) == 0 {
		return model.CustomerRole{}, errors.New("username doesn't match")
	}
	var resultStruct model.CustomerRole
	bsonBytes, _ := bson.Marshal(result[0])
	bson.Unmarshal(bsonBytes, &resultStruct)
	err = bcrypt.CompareHashAndPassword([]byte(resultStruct.Password), []byte(password))
	if err == nil {
		return resultStruct, nil
	} else {
		return model.CustomerRole{}, errors.New("password doesn't match")
	}
}
func (client *MongoClientImpl) GetAllMerchants() ([]model.AllMerchantResponse, error) {
	collection := fmt.Sprintf(config.MONGO_CUSTOMERS_ROLE_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	result := []bson.M{}
	err := c.Find(bson.M{}).Select(bson.M{"username": 1, "onboard": 1, config.MONGO_ID: 0}).All(&result)
	if err != nil {
		return []model.AllMerchantResponse{}, errors.New("<mongo> Unable to query collection")
	}
	if len(result) == 0 {
		return []model.AllMerchantResponse{}, errors.New("no merchants found")
	}
	var allMercahants[] model.AllMerchantResponse
	for _, v := range result {
		var resultStruct model.CustomerRole
		bsonBytes, _ := bson.Marshal(v)
		err:=bson.Unmarshal(bsonBytes, &resultStruct)
		if err == nil {
			if len(resultStruct.OnBoard)==0{
				resultStruct.OnBoard=config.REGISTER_ACTIVE
			}
			allMercahants=append(allMercahants,model.AllMerchantResponse{Username: resultStruct.Username,OnboardingStatus: resultStruct.OnBoard})
		} else {
			continue
		}
	}
	return  allMercahants,nil
}
func (client *MongoClientImpl) AddNewMerchantConfig(merchantConfig model.MerchantConfigDetails) (bool, error) {
	status,err:=client.CheckMerchantExists(merchantConfig.MerchantId,config.MONGO_CUSTOMERS_ROLE_COLLECTION,config.REGISTER_IN_PROCESS)
	if err != nil {
		return false, err
	}
	exists:=true
	if status{
		exists,err=client.CheckMerchantConfigNotExist(merchantConfig.MerchantId)
	}else{
		return false, errors.New(fmt.Sprintf("<mongo> Merchant with ID:%s not in process", merchantConfig.MerchantId))
	}
	if !exists{
		collection := config.MONGO_MERCHANT_CONFIG_COLLECTION
		session := client.session.Copy()
		defer session.Close()
		newMerchantConfig:=model.MongoMerchantConfigDetails{MerchantId: merchantConfig.MerchantId,NoOfSimilarContents: merchantConfig.NoOfSimilarContents,NoOfCollabContents: merchantConfig.NoOfCollabContents,
			NoOfCollabCategories: merchantConfig.NoOfCollabCategories,NoOfTopInterests: merchantConfig.NoOfTopInterests,ContentsForEveryUserInterest: merchantConfig.ContentsForEveryUserInterest,
		UseUniqueUserInterests: merchantConfig.UseUniqueUserInterests,UseInterestBasedTags: merchantConfig.UseInterestBasedTags,ArticleRecencyInDays: merchantConfig.ArticleRecencyInDays,
		FilterHistoryInDays: merchantConfig.FilterHistoryInDays,RecentHistoryCountInDays: merchantConfig.RecentHistoryCountInDays}
		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, config.STRONG_MODE)
		c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
		err = c.Insert(newMerchantConfig)
		if err != nil {
			msg := fmt.Sprintf("<mongo> Adding new record to merchant_config collection failed for merchant Id: %s, error : %v", merchantConfig.MerchantId, err)
			logger.Error(msg)
			return false,err
		}
		return true,nil
	}else{
		return false, errors.New(fmt.Sprintf("<mongo> Merchant config with ID:%s already exists", merchantConfig.MerchantId))
	}
}
func (client *MongoClientImpl) UpdateRegisterStatusInCustomerRole(merchantId,registerStatus string) error {
	collection := config.MONGO_CUSTOMERS_ROLE_COLLECTION
	session := client.session.Copy()
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	filter := bson.M{"merchant_id": bson.M{"$eq": merchantId}}
	update := bson.M{"$set": bson.M{"onboard": registerStatus}}
	err := c.Update(filter, update)
	if err != nil {
		msg := fmt.Sprintf("<mongo> Updating register field in customer role collection failed for merchant Id: %s, error : %v", merchantId, err)
		logger.Error(msg)
		return err
	}
	return nil
}
func (client *MongoClientImpl) AcceptRegistrationRequest(merchantId string) (bool, error) {
	status,err:=client.CheckMerchantExists(merchantId,config.MONGO_CUSTOMERS_ROLE_COLLECTION,config.REGISTER_REQUESTED)
	if err != nil {
		return false, err
	}
	if status{
		err=client.UpdateRegisterStatusInCustomerRole(merchantId,config.REGISTER_IN_PROCESS)
		if err != nil {
			return false, err
		}
		return true,nil
	}else{
		return false, errors.New(fmt.Sprintf("<mongo> Merchant with ID:%s not in requested state", merchantId))
	}
}
func (client *MongoClientImpl) RejectRegistrationRequest(merchantId string) (bool, error) {
	status,err:=client.CheckMerchantExists(merchantId,config.MONGO_CUSTOMERS_ROLE_COLLECTION,config.REGISTER_REQUESTED)
	if err != nil {
		return false, err
	}
	if status{
		err=client.UpdateRegisterStatusInCustomerRole(merchantId,config.REGISTER_REJECTED)
		if err != nil {
			return false, err
		}
		return true,nil
	}else{
		return false, errors.New(fmt.Sprintf("<mongo> Merchant with ID:%s not in requested state", merchantId))
	}
}
func (client *MongoClientImpl) CheckStrategyAndPageTypeExists(strategy,pageType string) (bool, error) {
	collection := fmt.Sprintf(config.MONGO_PAGE_TYPE_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	count, err := c.Find(bson.M{"strategy": strategy,"pageType": pageType}).Count()
	if err != nil {
		return true, err
	}
	if count!=0{
		return true,nil
	}else{
		return false,nil
	}
}
func (client *MongoClientImpl) CheckPageTypeExists(pageType string) (bool, error) {
	collection := fmt.Sprintf(config.MONGO_PAGE_TYPE_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	count, err := c.Find(bson.M{"pageType": pageType}).Count()
	if err != nil {
		return false, err
	}
	if count!=0{
		return true,nil
	}else{
		return false,nil
	}
}
func (client *MongoClientImpl) CheckPageTypeAndMerhantIdUnique(merhcantid,pageType string) (bool, error) {
	collection := fmt.Sprintf(config.MONGO_MERCHANT_PAGES_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	count, err := c.Find(bson.M{"merchantId": merhcantid,"PAGETYPE": pageType}).Count()
	if err != nil {
		return true, err
	}
	if count!=0{
		return true,nil
	}else{
		return false,nil
	}
}
func (client *MongoClientImpl) CheckPageTypeIdExists(pageTypeId int) (bool, error) {
	collection := fmt.Sprintf(config.MONGO_PAGE_RECOMMENDATIONS_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	count, err := c.Find(bson.M{"pageTypeId": pageTypeId}).Count()
	if err != nil {
		return false, err
	}
	if count!=0{
		return true,nil
	}else{
		return false,nil
	}
}
func (client *MongoClientImpl) CheckRecommendationTypeExists(recommendationType string) (bool, error) {
	collection := fmt.Sprintf(config.MONGO_PAGE_RECOMMENDATIONS_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	count, err := c.Find(bson.M{"recommendationType": recommendationType}).Count()
	if err != nil {
		return true, err
	}
	if count!=0{
		return true,nil
	}else{
		return false,nil
	}
}
func (client *MongoClientImpl) CheckPageTypeCount() (int, error) {
	collection := fmt.Sprintf(config.MONGO_PAGE_TYPE_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	count, err := c.Find(bson.M{}).Count()
	if err != nil {
		return count, err
	}
	return count,nil
}
func (client *MongoClientImpl) AddNewPageType(merchantId string,pageTypes []model.PageTypeDetails) ([]model.PageTypeDetailsResponse, error) {
	status,err:=client.CheckMerchantExists(merchantId,config.MONGO_CUSTOMERS_ROLE_COLLECTION,config.REGISTER_IN_PROCESS)
	if err != nil {
		return []model.PageTypeDetailsResponse{},err
	}
	if !status{
		return []model.PageTypeDetailsResponse{}, errors.New(fmt.Sprintf("<mongo> Merchant with ID:%s not in process", merchantId))
	}
	count,err:=client.CheckPageTypeCount()
	if err != nil {
		return []model.PageTypeDetailsResponse{},err
	}
	if count==0{
		return []model.PageTypeDetailsResponse{},errors.New(fmt.Sprintf("<mongo> Unable to query pageType Colletion"))
	}
	pageTypeDetailsResponse:=make([]model.PageTypeDetailsResponse,0)
	count++
	for _,pageType:=range pageTypes{
		exists,err:=client.CheckStrategyAndPageTypeExists(strings.ToLower(pageType.Strategy),strings.ToUpper(pageType.PageType))
		if err != nil {
			pageTypeDetailsResponse=append(pageTypeDetailsResponse,model.PageTypeDetailsResponse{
				Strategy: strings.ToLower(pageType.Strategy),PageType: strings.ToUpper(pageType.PageType),Message: err.Error(),
			})
			continue
		}
		if exists {
			pageTypeDetailsResponse=append(pageTypeDetailsResponse,model.PageTypeDetailsResponse{
				Strategy: strings.ToLower(pageType.Strategy),PageType: strings.ToUpper(pageType.PageType),Message: "Already exists",
			})
			continue
		}
		collection := config.MONGO_PAGE_TYPE_COLLECTION
		session := client.session.Copy()
		defer session.Close()
		pageTypeDetails:=model.MongoPageTypeDetails{PageTypeId: count,Strategy: strings.ToLower(pageType.Strategy),PageType: strings.ToUpper(pageType.PageType)}
		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, config.STRONG_MODE)
		c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
		err = c.Insert(pageTypeDetails)
		if err != nil {
			msg := fmt.Sprintf("<mongo> Adding new record to pageType collection failed,error : %v", err)
			logger.Error(msg)
			pageTypeDetailsResponse=append(pageTypeDetailsResponse,model.PageTypeDetailsResponse{
				Strategy: strings.ToLower(pageType.Strategy),PageType: strings.ToUpper(pageType.PageType),Message: msg,
			})
			continue
		}
		pageTypeDetailsResponse=append(pageTypeDetailsResponse,model.PageTypeDetailsResponse{
			Strategy: strings.ToLower(pageType.Strategy),PageType: strings.ToUpper(pageType.PageType),Message: "Added",
		})
		count++
	}
	return pageTypeDetailsResponse,nil
}

func (client *MongoClientImpl) GetAllPageRecommendations() ([]model.AllPageRecommendations, error) {
	collection := fmt.Sprintf(config.MONGO_PAGE_RECOMMENDATIONS_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	result := []bson.M{}
	err := c.Find(bson.M{}).Select(bson.M{}).All(&result)
	if err != nil {
		return []model.AllPageRecommendations{}, errors.New("<mongo> Unable to query collection")
	}
	if len(result) == 0 {
		return []model.AllPageRecommendations{}, errors.New("no pageRecommendations found")
	}
	var allMercahants[] model.AllPageRecommendations
	for _, v := range result {
		var resultStruct model.MongoPageRecommendations
		bsonBytes, _ := bson.Marshal(v)
		err:=bson.Unmarshal(bsonBytes, &resultStruct)
		if err == nil {
			allMercahants=append(allMercahants,model.AllPageRecommendations{PageTypeId: resultStruct.PageTypeId,
				RecommendationType: resultStruct.RecommendationType,Percentage: resultStruct.Percentage,Preference: resultStruct.Preference})
		} else {
			continue
		}
	}
	return  allMercahants,nil
}
func (client *MongoClientImpl) AddNewPageRecommendations(pageRecommendations []model.AllPageRecommendations) ([]model.AllPageRecommendationsResponse, error) {
	pageRecommendationResponse:=make([]model.AllPageRecommendationsResponse,0)
	for _,pageRecommendation:=range pageRecommendations{
		exists,err:=client.CheckRecommendationTypeExists(strings.ToLower(pageRecommendation.RecommendationType))
		if err != nil {
			pageRecommendationResponse=append(pageRecommendationResponse,model.AllPageRecommendationsResponse{PageTypeId: pageRecommendation.PageTypeId,
				RecommendationType: pageRecommendation.RecommendationType,Percentage: pageRecommendation.Percentage,Preference: pageRecommendation.Preference,
				Message: err.Error(),
			})
			continue
		}
		if exists {
			pageRecommendationResponse=append(pageRecommendationResponse,model.AllPageRecommendationsResponse{PageTypeId: pageRecommendation.PageTypeId,
				RecommendationType: pageRecommendation.RecommendationType,Percentage: pageRecommendation.Percentage,Preference: pageRecommendation.Preference,
				Message: "Already exists",
			})
			continue
		}
		collection := config.MONGO_PAGE_RECOMMENDATIONS_COLLECTION
		session := client.session.Copy()
		defer session.Close()
		pageRecommendationDetails:=model.MongoPageRecommendations{PageTypeId: pageRecommendation.PageTypeId,
			RecommendationType: strings.ToLower(pageRecommendation.RecommendationType),Percentage: pageRecommendation.Percentage,Preference: pageRecommendation.Preference,}
		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, config.STRONG_MODE)
		c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
		err = c.Insert(pageRecommendationDetails)
		if err != nil {
			msg := fmt.Sprintf("<mongo> Adding new record to page_recommendations collection failed,error : %v", err)
			logger.Error(msg)
			pageRecommendationResponse=append(pageRecommendationResponse,model.AllPageRecommendationsResponse{PageTypeId: pageRecommendation.PageTypeId,
				RecommendationType: pageRecommendation.RecommendationType,Percentage: pageRecommendation.Percentage,Preference: pageRecommendation.Preference,
				Message: msg,
			})
			continue
		}
		pageRecommendationResponse=append(pageRecommendationResponse,model.AllPageRecommendationsResponse{PageTypeId: pageRecommendation.PageTypeId,
			RecommendationType: pageRecommendation.RecommendationType,Percentage: pageRecommendation.Percentage,Preference: pageRecommendation.Preference,
			Message: "Added",
		})
	}
	return pageRecommendationResponse,nil
}
func (client *MongoClientImpl) AddNewMerchantPages(merchantId string,merchantpages []model.MerchantPagesDetails) ([]model.MerchantPagesResponse, error){
	exists,err:=client.CheckMerchantConfigNotExist(merchantId)
	if err != nil {
		return []model.MerchantPagesResponse{},err
	}
	if !exists{
		return []model.MerchantPagesResponse{}, errors.New(fmt.Sprintf("<mongo> Merchant with ID:%s not in merchant_config", merchantId))
	}
	merchantPagesResponse:=make([]model.MerchantPagesResponse,0)
	for _,merchantPage:=range merchantpages{
		if merchantId!=merchantPage.MerchantId{
			merchantPagesResponse=append(merchantPagesResponse,model.MerchantPagesResponse{PageTypeId: merchantPage.PageTypeId,MerchantId: merchantPage.MerchantId,
				TotalRecommendations: merchantPage.TotalRecommendations,Sorting: strings.ToUpper(merchantPage.Sorting),PAGETYPE: strings.ToUpper(merchantPage.PAGETYPE),
				FilterUserHistory:strings.ToUpper(merchantPage.FilterUserHistory),Tracker: merchantPage.Tracker,Message: "invalid merchant Id",
			})
			continue
		}
		exists,err:=client.CheckPageTypeIdExists(merchantPage.PageTypeId)
		if err != nil {
			merchantPagesResponse=append(merchantPagesResponse,model.MerchantPagesResponse{PageTypeId: merchantPage.PageTypeId,MerchantId: merchantPage.MerchantId,
				TotalRecommendations: merchantPage.TotalRecommendations,Sorting: strings.ToUpper(merchantPage.Sorting),PAGETYPE: strings.ToUpper(merchantPage.PAGETYPE),
				FilterUserHistory:strings.ToUpper(merchantPage.FilterUserHistory),Tracker: merchantPage.Tracker,Message: err.Error(),
			})
			continue
		}
		if !exists {
			merchantPagesResponse=append(merchantPagesResponse,model.MerchantPagesResponse{PageTypeId: merchantPage.PageTypeId,MerchantId: merchantPage.MerchantId,
				TotalRecommendations: merchantPage.TotalRecommendations,Sorting: strings.ToUpper(merchantPage.Sorting),PAGETYPE: strings.ToUpper(merchantPage.PAGETYPE),
				FilterUserHistory:strings.ToUpper(merchantPage.FilterUserHistory),Tracker: merchantPage.Tracker,Message: "Page type Id doesn't exists",
			})
			continue
		}
		exists,err=client.CheckPageTypeExists(merchantPage.PAGETYPE)
		if err != nil {
			merchantPagesResponse=append(merchantPagesResponse,model.MerchantPagesResponse{PageTypeId: merchantPage.PageTypeId,MerchantId: merchantPage.MerchantId,
				TotalRecommendations: merchantPage.TotalRecommendations,Sorting: strings.ToUpper(merchantPage.Sorting),PAGETYPE: strings.ToUpper(merchantPage.PAGETYPE),
				FilterUserHistory:strings.ToUpper(merchantPage.FilterUserHistory),Tracker: merchantPage.Tracker,Message: err.Error(),
			})
			continue
		}
		if !exists {
			merchantPagesResponse=append(merchantPagesResponse,model.MerchantPagesResponse{PageTypeId: merchantPage.PageTypeId,MerchantId: merchantPage.MerchantId,
				TotalRecommendations: merchantPage.TotalRecommendations,Sorting: strings.ToUpper(merchantPage.Sorting),PAGETYPE: strings.ToUpper(merchantPage.PAGETYPE),
				FilterUserHistory:strings.ToUpper(merchantPage.FilterUserHistory),Tracker: merchantPage.Tracker,Message: "Page type doesn't exists",
			})
			continue
		}
		exists,err=client.CheckPageTypeAndMerhantIdUnique(merchantPage.MerchantId,merchantPage.PAGETYPE)
		if err != nil {
			merchantPagesResponse=append(merchantPagesResponse,model.MerchantPagesResponse{PageTypeId: merchantPage.PageTypeId,MerchantId: merchantPage.MerchantId,
				TotalRecommendations: merchantPage.TotalRecommendations,Sorting: strings.ToUpper(merchantPage.Sorting),PAGETYPE: strings.ToUpper(merchantPage.PAGETYPE),
				FilterUserHistory:strings.ToUpper(merchantPage.FilterUserHistory),Tracker: merchantPage.Tracker,Message: err.Error(),
			})
			continue
		}
		if exists {
			merchantPagesResponse=append(merchantPagesResponse,model.MerchantPagesResponse{PageTypeId: merchantPage.PageTypeId,MerchantId: merchantPage.MerchantId,
				TotalRecommendations: merchantPage.TotalRecommendations,Sorting: strings.ToUpper(merchantPage.Sorting),PAGETYPE: strings.ToUpper(merchantPage.PAGETYPE),
				FilterUserHistory:strings.ToUpper(merchantPage.FilterUserHistory),Tracker: merchantPage.Tracker,Message: "PAGETYPE with same merhcant id already exists",
			})
			continue
		}

		collection := config.MONGO_MERCHANT_PAGES_COLLECTION
		session := client.session.Copy()
		defer session.Close()
		merchantPages :=model.MongoMerchantPages{PageTypeId: merchantPage.PageTypeId,MerchantId: merchantPage.MerchantId,
			TotalRecommendations: merchantPage.TotalRecommendations,Sorting: strings.ToUpper(merchantPage.Sorting),PAGETYPE: strings.ToUpper(merchantPage.PAGETYPE),
			FilterUserHistory:strings.ToUpper(merchantPage.FilterUserHistory),Tracker: merchantPage.Tracker}
		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, config.STRONG_MODE)
		c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
		err = c.Insert(merchantPages)
		if err != nil {
			msg := fmt.Sprintf("<mongo> Adding new record to merchant_pages collection failed,error : %v", err)
			logger.Error(msg)
			merchantPagesResponse=append(merchantPagesResponse,model.MerchantPagesResponse{PageTypeId: merchantPage.PageTypeId,MerchantId: merchantPage.MerchantId,
				TotalRecommendations: merchantPage.TotalRecommendations,Sorting: strings.ToUpper(merchantPage.Sorting),PAGETYPE: strings.ToUpper(merchantPage.PAGETYPE),
				FilterUserHistory:strings.ToUpper(merchantPage.FilterUserHistory),Tracker: merchantPage.Tracker,Message: msg,
			})
			continue
		}
		merchantPagesResponse=append(merchantPagesResponse,model.MerchantPagesResponse{PageTypeId: merchantPage.PageTypeId,MerchantId: merchantPage.MerchantId,
			TotalRecommendations: merchantPage.TotalRecommendations,Sorting: strings.ToUpper(merchantPage.Sorting),PAGETYPE: strings.ToUpper(merchantPage.PAGETYPE),
			FilterUserHistory:strings.ToUpper(merchantPage.FilterUserHistory),Tracker: merchantPage.Tracker,Message: "Added",
		})
	}
	return merchantPagesResponse,nil
}
func (client *MongoClientImpl) GetAllRouteInfo(merchantId string) (model.RouteInfoResponse, error) {
	collection := fmt.Sprintf(config.MONGO_ROUTE_INFO_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	result := bson.M{}
	err := c.Find(bson.M{"merchant_id": merchantId}).Select(bson.M{}).One(&result)
	if err != nil {
		return model.RouteInfoResponse{}, errors.New("<mongo> Unable to query collection")
	}
	if len(result) == 0 {
		return model.RouteInfoResponse{}, errors.New("no route info found")
	}
	var resultStruct model.RouteInfo
	bsonBytes, _ := bson.Marshal(result)
	err=bson.Unmarshal(bsonBytes, &resultStruct)
	if err == nil {
		return  model.RouteInfoResponse{Route:resultStruct.Route,Name: resultStruct.Name,Component: resultStruct.Component,Description: resultStruct.Description,
			Images: resultStruct.Images,VideoSrc: resultStruct.VideoSrc,DefaultImage: resultStruct.DefaultImage,DefaultSlug: resultStruct.DefaultSlug,
			DefaultImageSource: resultStruct.DefaultImageSource},nil
	} else {
		return model.RouteInfoResponse{}, errors.New("unable to unmarshall result")
	}
}
/*func (client *MongoClientImpl) CheckIsMerchantOnBoarded(merchantId string) (bool, error) {
	collection := fmt.Sprintf(config.MONGO_MERCHANT_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	result := bson.M{}
	err := c.Find(bson.M{"merchant_id": merchantId}).One(&result)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	var resultStruct model.MongoMerchant
	bsonBytes, err := bson.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("<mongo> Unable to marshall bson. Error: %v",err.Error()))
		return false, err
	}
	bson.Unmarshal(bsonBytes, &resultStruct)
	if resultStruct.Register==config.REGISTER_COMPLETED{
		return true,nil
	}else{
		return false,nil
	}
}
//check if page type sent in merchant_pages present in pagetType collection
//if not present then it should be present in request body in pagetType section
//update merchant_config
//

func (client *MongoClientImpl) IsPageTypePresent(onBoardingDetails model.OnBoardDetailsRequest) (bool, error) {
	pageTypes:=make([]string,0)
	for _,merchant_page:= range onBoardingDetails.MerchantPagesDetails{
		pageTypes=append(pageTypes,merchant_page.PAGETYPE)
	}
	if len(onBoardingDetails.PageTypeDetails)==0{
		//check merchant_pages pagetype should be present in pageType collection

	}else{
		pageTypeInPageTypeDetailsMap:=make(map[string]string)
		for _,pageTypeDetail:= range onBoardingDetails.PageTypeDetails{
			pageTypeInPageTypeDetailsMap[pageTypeDetail.PageType]=pageTypeDetail.PageType
		}
		for _,pageType:=range pageTypes{
			if _,ok:=pageTypeInPageTypeDetailsMap[pageType];ok{
				continue
			}else{
				return false,nil
			}
		}
	}
	collection := fmt.Sprintf(config.MONGO_MERCHANT_COLLECTION)
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.MONGO_SIMS_DATABASE).C(collection)
	result := bson.M{}
	err := c.Find(bson.M{"merchant_id": merchantId}).One(&result)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	var resultStruct model.MongoMerchant
	bsonBytes, err := bson.Marshal(result)
	if err != nil {
		logger.Error(fmt.Sprintf("<mongo> Unable to marshall bson. Error: %v",err.Error()))
		return false, err
	}
	bson.Unmarshal(bsonBytes, &resultStruct)
	if resultStruct.Register==config.REGISTER_COMPLETED{
		return true,nil
	}else{
		return false,nil
	}
}
func (client *MongoClientImpl) OnBoardNewMerchant(onBoardingDetails model.OnBoardDetailsRequest) (model.RegisterationDetailsResponse, error) {
	onBoarded,err:=client.CheckIsMerchantOnBoarded(onBoardingDetails.MerchantConfigDetails.MerchantId)
	if err != nil {
		logger.Error(err.Error())
		return model.RegisterationDetailsResponse{}, err
	}
	if onBoarded{
		logger.Info("<mongo> Merchant with merchantId:%s already onBoarded", onBoardingDetails.MerchantConfigDetails.MerchantId)
		return model.RegisterationDetailsResponse{}, errors.New(fmt.Sprintf("<mongo> Merchant with merchantId:%s already onBoarded", onBoardingDetails.MerchantConfigDetails.MerchantId))
	}
	max_merchant_id,err:=client.GetMaxMerchantId()
	if err != nil {
		logger.Error(err.Error())
		return model.RegisterationDetailsResponse{}, err
	}
	success,err:=client.CreateRegisterationTransaction(userName,password,strconv.Itoa(max_merchant_id+1))
	if err != nil {
		logger.Error(err.Error())
		return model.RegisterationDetailsResponse{}, err
	}
	if success {
		return model.RegisterationDetailsResponse{Username: userName,Merchant_id: strconv.Itoa(max_merchant_id+1)}, nil
	}
	return model.RegisterationDetailsResponse{}, err
}*/