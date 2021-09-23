package model

import (
	"encoding/json"
	"net/http"
	"time"
)

type RecommendationRequest struct {
	Cid                    string
	Uid                    string
	MerchantId             string
	SimsThreshold          float64
	ST                     string
	Sid                    string
	Source                 bool
	Timestamp              string
	CategoryId             string
	ChronologySort         bool
	IP                     string
	LogPrefix              string
	StrategyType           int64
	RecommendationStrategy []PageRecommendations
	ConfigKeys             map[string]string
	MerchantWiseConfig     MerchantWiseConfig
	Category               string
	Offset                 int64
	Count                  int64
	Tag                    string
	Request                *http.Request
}

type UserSaveRequest struct {
	Uid            string            `json:"uid,omitempty"`
	Cid            string            `json:"cid,omitempty"`
	Event          string            `json:"eid,omitempty"`
	MerchantId     string            `json:"merchantId,omitempty"`
	Sid            string            `json:"sessionId,omitempty"`
	Ref            string            `json:"ref,omitempty"`
	Timestamp      string            `json:"timestamp,omitempty"`
	Recommendation string            `json:"recommendation,omitempty"`
	EventType      string            `json:"eventType,omitempty"`
	Ip             string            `json:"ip,omitempty"`
	UserAgent      string            `json:"userAgent,omitempty"`
	ConfigKeys     map[string]string `json:"-"`
	Payload        json.RawMessage   `json:"payload,omitempty"`
	UserAccountId  string            `json:"useraccountid,omitempty"`
	StockId        string            `json:"stockid,omitempty"`
	Platform       string            `json:"platform,omitempty"`
}

type DssHandlerEventListener struct {
	SaveHits          bool
	Session           bool
	PopulateIP        bool
	PopulateUserEvent bool
	SaveKafka         bool
}

type Article struct {
	Id                   string `json:"id,omitempty"`
	MerchantId           string `json:"mid,omitempty"`
	Headline             string `json:"headline,omitempty"`
	AuthorName           string `json:"author-name,omitempty"`
	BreakingNewsLinkedId string `json:"breaking-news-linked-story-id,omitempty"`
	Text                 string `json:"text,omitempty"`
	PublishedAt          int64  `json:"published-at,omitempty"`
	FirstPublishedAt     int64  `json:"first_published_at,omitempty"`
	PublisherId          string `json:"publisher-id,omitempty"`
	Section_id           string `json:"section_id,omitempty"`
	Section_name         string `json:"section_name,omitempty"`
	Tag_id               string `json:"tag_id,omitempty"`
	Tag_name             string `json:"tag_name,omitempty"`
	Time_To_Read         int64  `json:"time_to_read,omitempty"`
	Slug                 string `json:"slug,omitempty"`
	ContentType          string `json:"content_type,omitempty"`
	VideoId              string `json:"video_id,omitempty"`
	ImageUrl             string `json:"ImageUrl,omitempty"`
	StoryTemplate        string `json:"StoryTemplate,omitempty"`
	ImageMetaData        string `json:"ImageMetaData,omitempty"`
	StockIds             string `json:"StockIds,omitempty"`
	AuthorNames          string `json:"author-names,omitempty"`
	MetaDescription      string `json:"meta-description,omitempty"`
	SubHeadline          string `json:"subheadline,omitempty"`
	Access               string `json:"access,omitempty"`
	AccessLevelValue     int64  `json:"access-level-value,omitempty"`
}

type ArticleMongo struct {
	Id                   string    `json:"Id" bson:"Id"`
	MerchantId           string    `json:"MerchantId" bson:"MerchantId"`
	Headline             string    `json:"Headline" bson:"Headline,omitempty"`
	AuthorName           string    `json:"AuthorName" bson:"AuthorName,omitempty"`
	BreakingNewsLinkedId string    `json:"BreakingNewsLinkedId" bson:"BreakingNewsLinkedId,omitempty"`
	Text                 string    `json:"Text" bson:"Text,omitempty"`
	PublishedAt          int64     `json:"PublishedAt" bson:"PublishedAt,omitempty"`
	FirstPublishedAt     int64     `json:"FirstPublishedAt" bson:"FirstPublishedAt,omitempty"`
	PublisherId          string    `json:"PublisherId" bson:"PublisherId,omitempty"`
	Section_id           string    `json:"Section_id" bson:"Section_id,omitempty"`
	Section_name         string    `json:"Section_name" bson:"Section_name,omitempty"`
	Tag_id               string    `json:"Tag_id" bson:"Tag_id,omitempty"`
	Tag_name             string    `json:"Tag_name" bson:"Tag_name,omitempty"`
	Time_To_Read         int64     `json:"Time_To_Read" bson:"Time_To_Read,omitempty"`
	Slug                 string    `json:"Slug" bson:"Slug,omitempty"`
	ContentType          string    `json:"ContentType" bson:"ContentType,omitempty"`
	VideoId              string    `json:"VideoId" bson:"VideoId,omitempty"`
	ImageUrl             string    `json:"ImageUrl" bson:"ImageUrl,omitempty"`
	PublishedAtDateTime  time.Time `json:"PublishedAtDateTime" bson:"PublishedAtDateTime"`
	StoryTemplate        string    `json:"StoryTemplate" bson:"StoryTemplate,omitempty"`
	ImageMetaData        string    `json:"ImageMetaData" bson:"ImageMetaData,omitempty"`
}

type RecommendationEvents struct {
	Timestamp          string `json:"Timestamp,omitempty"`
	MerchantId         string `json:"MerchantId,omitempty"`
	Cid                string `json:"Cid,omitempty"`
	Uid                string `json:"Uid,omitempty"`
	CategoryId         string `json:"CategoryId,omitempty"`
	Ip                 string `json:"Ip,omitempty"`
	RequestUrl         string `json:"RequestUrl,omitempty"`
	ResponseStatus     string `json:"ResponseStatus,omitempty"`
	SimsType           string `json:"SimsType,omitempty"`
	RecommendationType string `json:"RecommendationType,omitempty"`
	RecommendedCids    string `json:"RecommendedCids,omitempty"`
	RecommendedModels  string `json:"RecommendedModels,omitempty"`
}

type RecommendedModel struct {
	Model string `json:"Model,omitempty"`
	Cids  string `json:"Cids,omitempty"`
}

type Stock struct {
	Name    string `json:"name,omitempty"`
	IsIn    string `json:"isin,omitempty"`
	NseCode string `json:"nse-code,omitempty"`
	BseCode string `json:"bse-code,omitempty"`
	StockId string `json:"stock-id,omitempty"`
}

type Stocks_Watchlist struct {
	Stock_ids []string `json:"stock_ids,omitempty"`
}

type StockId_WithTimestamp struct {
	StockId   string `json:"stock_ids,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

type Sections struct {
	Preferences   []string `json:"preferences,omitempty"`
	UserAccountId string   `json:"user_account_id,omitempty"`
}
type UserProfile struct {
	UserDetails     Sections         `json:"user_details,omitempty"`
	StockViewed     Stock            `json:"stock,omitempty"`
	StocksWatchlist Stocks_Watchlist `json:"stocks_watchlist,omitempty"`
}
type UserEvent struct {
	Timestamp               string            `json:"timestamp,omitempty"`
	MerchantId              string            `json:"merchantId,omitempty"`
	Cid                     string            `json:"cid,omitempty"`
	Uid                     string            `json:"uid,omitempty"`
	Sid                     string            `json:"sessionId,omitempty"`
	EventId                 string            `json:"event,omitempty"`
	EventType               string            `json:"eventType,omitempty"`
	Headline                string            `json:"headline,omitempty"`
	Section                 string            `json:"section,omitempty"`
	Tag                     string            `json:"tag,omitempty"`
	Ip                      string            `json:"ip,omitempty"`
	RefUrl                  string            `json:"refUrl,omitempty"`
	Referrer                string            `json:"referrer,omitempty"`
	LocationDetails         LocationDetails   `json:"locationDetails,omitempty"`
	PlatformDetails         PlatformDetails   `json:"platformDetails,omitempty"`
	ConfigKeys              map[string]string `json:"-"`
	TrendingService         interface{}       `json:"-"`
	SaveTrendingKeysService interface{}       `json:"-"`
	WruSid                  string            `json:"wruSid,omitempty"`
	Payload                 json.RawMessage   `json:"payload,omitempty"`
	Platform                string            `json:"platform,omitempty"`
	Analytics               Analytics         `json:"analytics,omitempty"`
}

type Analytics struct {
	ExperimentRecommendationTypeId int64 `json:"experimentRecommendationTypeId,omitempty"`
}

type RrsResponse struct {
	Category    string        `json:"category,omitempty"`
	Author      string        `json:"author,omitempty"`
	CategoryId  string        `json:"categoryId,omitempty"`
	StockId     string        `json:"StockId,omitempty"`
	StockName   string        `json:"StockName,omitempty"`
	StockSlug   string        `json:"StockSlug,omitempty"`
	Vsims       []ContentData `json:"vsims,omitempty"`
	Isims       []ContentData `json:"isims,omitempty"`
	Csims       []ContentData `json:"csims,omitempty"`
	Usims       []ContentData `json:"usims,omitempty"`
	Tsims       []ContentData `json:"tsims,omitempty"`
	Ctsims      []ContentData `json:"ctsims,omitempty"`
	Default     []ContentData `json:"recommendations,omitempty"`
	CurrentData ContentData   `json:"currentData,omitempty"`
}

type StockData struct {
	StockId     string `json:"StockId,omitempty"`
	StockName   string `json:"StockName,omitempty"`
	StockSlug   string `json:"StockSlug,omitempty"`
	StockSector string `json:"StockSector,omitempty"`
	IsIn        string `json:"isin,omitempty"`
	StockAlias  string `json:"aliases,omitempty"`
}

type ContentData struct {
	Cid              string                 `json:"Cid,omitempty"`
	SimsScore        float64                `json:"SimsScore,omitempty"`
	Description      string                 `json:"Description,omitempty"`
	ImageLink        string                 `json:"ImageLink,omitempty"`
	ProductLink      string                 `json:"ProductLink,omitempty"`
	Category         string                 `json:"Category,omitempty"`
	Tags             string                 `json:"Tags,omitempty"`
	PublishedAt      string                 `json:"PublishedAt,omitempty"`
	FirstPublishedAt string                 `json:"first_published_at,omitempty"`
	Timestamp        time.Time              `json:"Timestamp,omitempty"`
	Slug             string                 `json:"Slug,omitempty"`
	Model            string                 `json:"Model,omitempty"`
	Reason           string                 `json:"Reason,omitempty"`
	Analysis         string                 `json:"Analysis,omitempty"`
	Thumbnail        string                 `json:"Thumbnail,omitempty"`
	Duration         float64                `json:"Duration,omitempty"`
	ImageUrl         string                 `json:"ImageUrl"`
	ImageMetaData    string                 `json:"ImageMetaData"`
	AuthorName       string                 `json:"AuthorName"`
	Additional       map[string]interface{} `json:"Additional,omitempty"`
	StockId          string                 `json:"StockId,omitempty"`
	StockName        string                 `json:"StockName,omitempty"`
	StockSlug        string                 `json:"StockSlug,omitempty"`
	CategoryList     []string               `json:"CategoryList,omitempty"`
	MetaDescription  string                 `json:"MetaDescription,omitempty"`
	SubHeadline      string                 `json:"SubHeadline,omitempty"`
	Access           string                 `json:"Access,omitempty"`
	AccessLevelValue int64                  `json:"AccessLevelValue,omitempty"`
}
type SortedGroupedContentData struct{
	Count int								`json:"Count,omitempty"`
	Category string							`json:"Category,omitempty"`
	CategoryData []GroupedContentData 		`json:"CategoryData,omitempty"`
}
type CategoryWithSize struct{
	Count int								`json:"Count,omitempty"`
	Category string							`json:"Category,omitempty"`
}
type GroupedContentData struct {
	Cid              string                 `json:"Cid,omitempty"`
	Headline        string                `json:"Headline,omitempty"`
	SubHeadline        string                `json:"SubHeadline,omitempty"`
	Category         string                 `json:"Category,omitempty"`
	Timestamp        time.Time              `json:"Timestamp,omitempty"`
	ImageUrl         string                 `json:"ImageUrl,omitempty"`
}

type ContentDetail struct {
	SimsScore float64
	Thumbnail string
	Duration  float64
}

type UserCredentials struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}
type RegisterationDetailsRequest struct {
	Username     string `json:"user"`
	Password     string `json:"pass"`
	Platform     string `json:"platform"`
	Email        string `json:"email"`
	Language        string `json:"language"`
	Address      string `json:"address"`
	AmpSelection string `json:"ampSelection"`
	PhoneNumber  string `json:"phoneNumber"`
	Images             string      `json:"images"`
	DefaultImage       string      `json:"defaultImage"`
	DefaultImageSource string      `json:"defaultImageSource"`
}
type RegisterationDetailsResponse struct {
	Username string `json:"user"`
	Merchant_id string `json:"merchant_id"`
}
type OnBoardDetailsRequest struct {
	MerchantConfigDetails MerchantConfigDetails `json:"merchantConfigDetails"`
	PageTypeDetails []PageTypeDetails `json:"pageTypeDetails"`
	MerchantPagesDetails []MerchantPagesDetails `json:"merchantPagesDetails"`
	PageRecommendationDetails []PageRecommendationDetails `json:"pageRecommendationDetails"`
}
type MerchantConfigDetails struct {
	MerchantId                   string `json:"merchant_id"`
	NoOfSimilarContents          int    `json:"no_of_similar_contents"`
	NoOfCollabCategories         int    `json:"no_of_collab_categories"`
	NoOfCollabContents           int    `json:"no_of_collab_contents"`
	NoOfTopInterests             int    `json:"no_of_top_interests"`
	ContentsForEveryUserInterest int    `json:"contents_for_every_user_interest"`
	UseUniqueUserInterests       int    `json:"use_unique_user_interests"`
	UseInterestBasedTags         int    `json:"use_interest_based_tags"`
	ArticleRecencyInDays         int    `json:"article_recency_in_days"`
	FilterHistoryInDays          int    `json:"filter_history_in_days"`
	RecentHistoryCountInDays     int    `json:"recent_history_count_in_days"`
}
type MongoMerchantConfigDetails struct {
	MerchantId                   string `bson:"merchant_id"`
	NoOfSimilarContents          int    `bson:"no_of_similar_contents"`
	NoOfCollabCategories         int    `bson:"no_of_collab_categories"`
	NoOfCollabContents           int    `bson:"no_of_collab_contents"`
	NoOfTopInterests             int    `bson:"no_of_top_interests"`
	ContentsForEveryUserInterest int    `bson:"contents_for_every_user_interest"`
	UseUniqueUserInterests       int    `bson:"use_unique_user_interests"`
	UseInterestBasedTags         int    `bson:"use_interest_based_tags"`
	ArticleRecencyInDays         int    `bson:"article_recency_in_days"`
	FilterHistoryInDays          int    `bson:"filter_history_in_days"`
	RecentHistoryCountInDays     int    `bson:"recent_history_count_in_days"`
}
type PageTypeDetails struct {
	PageTypeId int    `json:"pageTypeId,omitempty"`
	Strategy   string `json:"strategy,omitempty"`
	PageType   string `json:"pageType,omitempty"`
}
type MongoPageTypeDetails struct {
	PageTypeId int    `bson:"pageTypeId"`
	Strategy   string `bson:"strategy"`
	PageType   string `bson:"pageType"`
}
type PageTypeDetailsResponse struct {
	Strategy   string `json:"strategy,omitempty"`
	PageType   string `json:"pageType,omitempty"`
	Message   string `json:"message,omitempty"`
}

type PageRecommendationDetails struct {
	PageTypeId         int     `json:"pageTypeId"`
	RecommendationType string  `json:"recommendationType"`
	Percentage         float64 `json:"percentage"`
	Preference         int     `json:"preference"`
}
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type Response struct {
	Data string `json:"data"`
}

type AccessToken struct {
	Token string `json:"token"`
}

type MerchantExtractRef struct {
	Version     int
	Merchant_id string
	Type        string
	Input_path  string
	Output_path string
	Columns     []ColumnMappings
}
type MongoMerchantExtractRef struct {
	ID          int    `bson:"id"`
	Version     int    `bson:"version"`
	Merchant_id string `bson:"merchant_id"`
	Type        string `bson:"type"`
	Input_path  string `bson:"input_path"`
	Output_path string `bson:"output_path"`
}

type ColumnMappings struct {
	Id               string `bson:"extract_id"`
	Column_name      string `bson:"column_name"`
	Column_index     string `bson:"column_index"`
	Column_weight    string `bson:"column_weight"`
	Column_algorithm string `bson:"column_algorithm"`
	DataType         string `bson:"data_type"`
}
type CustomerRole struct {
	Username             string `bson:"username"`
	Role             string `bson:"role"`
	Password             string `bson:"password"`
	RouteStats           string `bson:"route_stats"`
	RouteRecommendations string `bson:"route_recommendations"`
	RouteCharts          string `bson:"route_charts"`
	ClientsAccess        string `bson:"clients_access"`
	Platform string `bson:"platform"`
	Email string `bson:"email"`
	Address string `bson:"address"`
	AmpSelection string `bson:"ampSelection"`
	PhoneNumber string `bson:"phoneNumber"`
	OnBoard string `bson:"onboard"`
	MerchantId string `bson:"merchant_id"`
}
type CustomerRoleResponse struct {
	Username string `json:"user"`
	Role             string `json:"role"`
	Password             string `json:"password"`
	RouteStats           string `json:"route_stats"`
	RouteRecommendations string `json:"route_recommendations"`
	RouteCharts          string `json:"route_charts"`
	ClientsAccess        string `json:"clients_access"`
	Platform string `json:"platform"`
	Email string `json:"email"`
	Address string `json:"address"`
	AmpSelection string `json:"ampSelection"`
	PhoneNumber string `json:"phoneNumber"`
	OnboardingStatus string `json:"onboarding_status"`
	MerchantId string `json:"merchant_id"`
}
type AllMerchantResponse struct {
	Username string `json:"user"`
	OnboardingStatus string `json:"onboarding_status"`
}
type MerchantPagesDetails struct {
	PageTypeId           int    `json:"pageTypeId"`
	MerchantId           string `json:"merchantId"`
	TotalRecommendations int    `json:"total_recommendations"`
	PageDescription      string `json:"pageDescription"`
	Sorting              string `json:"sorting"`
	PAGETYPE             string `json:"PAGETYPE"`
	FilterUserHistory    string `json:"filterUserHistory"`
	Tracker              string `json:"tracker"`
}
type MongoMerchantPages struct {
	PageTypeId           int    `bson:"pageTypeId"`
	MerchantId           string `bson:"merchantId"`
	TotalRecommendations int    `bson:"total_recommendations"`
	PageDescription      string `bson:"pageDescription"`
	Sorting              string `bson:"sorting"`
	PAGETYPE             string `bson:"PAGETYPE"`
	FilterUserHistory    string `bson:"filterUserHistory"`
	Tracker              string `bson:"tracker"`
}
type MerchantPagesResponse struct {
	PageTypeId           int    `json:"pageTypeId"`
	MerchantId           string `json:"merchantId"`
	TotalRecommendations int    `json:"total_recommendations"`
	PageDescription      string `json:"pageDescription"`
	Sorting              string `json:"sorting"`
	PAGETYPE             string `json:"PAGETYPE"`
	FilterUserHistory    string `json:"filterUserHistory"`
	Tracker              string `json:"tracker"`
	Message   string `json:"message,omitempty"`
}
type AllPageRecommendations struct {
	PageTypeId         int     `json:"pageTypeId"`
	RecommendationType string  `json:"recommendationType"`
	Percentage         float64 `json:"percentage"`
	Preference         int     `json:"preference"`
}
type AllPageRecommendationsResponse struct {
	PageTypeId         int     `json:"pageTypeId"`
	RecommendationType string  `json:"recommendationType"`
	Percentage         float64 `json:"percentage"`
	Preference         int     `json:"preference"`
	Message   string `json:"message,omitempty"`
}
type MongoPageRecommendations struct {
	PageTypeId         int     `bson:"pageTypeId"`
	RecommendationType string  `bson:"recommendationType"`
	Percentage         float64 `bson:"percentage"`
	Preference         int     `bson:"preference"`
}
type ActionRequest struct {
	MerchantId         string `json:"merchant_id"`
	Action string  `json:"Action"`
}
type MongoMerchantPlatform struct {
	Merchant_id string `bson:"merchant_id"`
	Platform    string `bson:"platform"`
}
type MongoMerchant struct {
	Merchant_id string `bson:"merchant_id"`
	Merchant_name    string `bson:"merchant_name"`
	Article_sims_content_key    string `bson:"article_sims_content_key"`
	Article_sims_recency_key    string `bson:"article_sims_recency_key"`
	Video_sims_content_key    string `bson:"video_sims_content_key"`
	Video_sims_recency_key    string `bson:"video_sims_recency_key"`
	Video_accountId    string `bson:"video_accountId"`
	Video_clientId    string `bson:"video_clientId"`
	Video_secretKey    string `bson:"video_secretKey"`
	Register    string `bson:"register"`
}
type MongoMerhcantConfig struct {
	Merchant_id                  string `bson:"merchant_id"`
	NoOfSimilarContents          int    `bson:"no_of_similar_contents"`
	NoOfCollabCategories         int    `bson:"no_of_collab_categories"`
	NoOfCollabContents           int    `bson:"no_of_collab_contents"`
	NoOfTopInterests             int    `bson:"no_of_top_interests"`
	ContentsForEveryUserInterest int    `bson:"contents_for_every_user_interest"`
	UseUniqueUserInterests       int    `bson:"use_unique_user_interests"`
	UseInterestBasedTags         int    `bson:"use_interest_based_tags"`
	ArticleRecencyInDays         int    `bson:"article_recency_in_days"`
	FilterHistoryInDays          int    `bson:"filter_history_in_days"`
	RecentHistoryCountInDays     int    `bson:"recent_history_count_in_days"`
}
type RouteInfo struct {
	Route              string `bson:"route"`
	Name               string `bson:"name"`
	Component          string `bson:"component"`
	Description        string `bson:"description"`
	Images             string `bson:"images"`
	VideoSrc           string `bson:"videoSrc"`
	DefaultImage       string `bson:"defaultImage"`
	DefaultSlug        string `bson:"defaultSlug"`
	DefaultImageSource string `bson:"defaultImageSource"`
}
type RouteInfoResponse struct {
	Route              string `json:"route"`
	Name               string `json:"name"`
	Component          string `json:"component"`
	Description        string `json:"description"`
	Images             string `json:"images"`
	VideoSrc           string `json:"videoSrc"`
	DefaultImage       string `json:"defaultImage"`
	DefaultSlug        string `json:"defaultSlug"`
	DefaultImageSource string `json:"defaultImageSource"`
}
type Merchant struct {
	Id                   string
	Name                 string
	ArticleKey           string
	ArticleRecencyKey    string
	VideoKey             string
	VideoRecencyKey      string
	VideoAccountId       string
	VideoClientId        string
	VideoSecretKey       string
	TotalRecommendations map[int64]int
	Recommendations      map[int64][]PageRecommendations
	Experiments          map[string][]Experiment
}
type Experiment struct {
	RecommendationTypeId int64
	Description          string
	CreatedAt            string
	ExperimentType       string
}

type MerchantDetails struct {
	Merchants      map[string]Merchant
	Trending       map[string]string
	TrendingVideos map[string]string
}

type LocationDetails struct {
	City      string  `json:"city,omitempty"`
	State     string  `json:"state,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Country   string  `json:"country,omitempty"`
	TimeZone  string  `json:"timezone,omitempty"`
}

type PlatformDetails struct {
	Browser                string `json:"browser,omitempty"`
	BrowserVersion         string `json:"browserVersion,omitempty"`
	BrowserMajorVer        string `json:"browserMajorVer,omitempty"`
	BrowserMinorVer        string `json:"browserMinorVer,omitempty"`
	BrowserType            string `json:"browserType,omitempty"`
	DeviceName             string `json:"deviceName,omitempty"`
	DeviceType             string `json:"deviceType,omitempty"`
	DeviceBrand            string `json:"deviceBrand,omitempty"`
	Platform               string `json:"platform,omitempty"`
	PlatformVersion        string `json:"platformVersion,omitempty"`
	PlatformShort          string `json:"platformShort,omitempty"`
	DeviceIsMobile         bool   `json:"deviceIsMobile,omitempty"`
	DeviceIsAndroid        bool   `json:"deviceIsAndroid,omitempty"`
	DeviceIsConsole        bool   `json:"deviceIsConsole,omitempty"`
	DeviceIsCrawler        bool   `json:"deviceIsCrawler,omitempty"`
	DeviceIsDesktop        bool   `json:"deviceIsDesktop,omitempty"`
	DeviceIsIPad           bool   `json:"deviceIsIPad,omitempty"`
	DeviceIsIPhone         bool   `json:"deviceIsIPhone,omitempty"`
	DeviceIsTablet         bool   `json:"deviceIsTablet,omitempty"`
	DeviceIsTv             bool   `json:"deviceIsTv,omitempty"`
	DeviceIsWinPhone       bool   `json:"deviceIsWinPhone,omitempty"`
	RenderingEngineName    string `json:"renderingEngineName,omitempty"`
	RenderingEngineVersion string `json:"renderingEngineVersion,omitempty"`
	JavaScript             string `json:"javaScript,omitempty"`
	Cookies                string `json:"cookies,omitempty"`
	Crawler                string `json:"crawler,omitempty"`
}

type BlockedUsersList struct {
	Country   string
	City      string
	Timestamp string
}

type LocationBasedFetch struct {
	MerchantId string
	Geo        string
	FetchCount int64
}

type SplitRecommendationReq struct {
	RedisKey   string
	FetchCount int64
}

type PageRecommendations struct {
	Type              string
	PageType          string
	Percentage        int64
	Preference        int64
	Sorting           string
	FilterUserHistory string
	Tracker           string
}

type PredictEvent struct {
	Id           string  `json:"id"`
	PredictScore float64 `json:"predictionViews"`
	CurrentViews float64 `json:"currentViews"`
	Slug         string  `json:"Slug,omitempty"`
}

type AnalysisLocationEvent struct {
	Location      string  `json:"location"`
	TotalUsers    int64   `json:"totalUsers"`
	DistinctUsers int64   `json:"distinctUsers"`
	Average       float64 `json:"average"`
}

type AnalysisStrategyEvent struct {
	Type          string  `json:"type"`
	TotalUsers    int64   `json:"totalUsers"`
	DistinctUsers int64   `json:"distinctUsers"`
	Average       float64 `json:"average"`
}

type HitsCount struct {
	Type      string  `json:"requestType"`
	HitsCount float64 `json:"hitsCount"`
}

type ArticleVector struct {
	Cid    string
	Vector []float64
}

type ArticleCategoryVector struct {
	Cid      string
	Vector   []float64
	Category string
}

type ArticleClusterVector struct {
	Cid     string
	Vector  []float64
	Cluster string
}

type ClusterPageReference struct {
}

type CategoryScores struct {
	Score      float64
	Categories []string
}

type RecommReason struct {
	Cid          string
	Model        string
	TechReason   string
	ClientReason string
	Entity       string
}

type UserHistoryData struct {
	Cid       string  `json:"Cid,omitempty"`
	Timestamp float64 `json:"Timestamp,omitempty"`
	Sid       string  `json:"Sid,omitempty"`
}

type UserStockHistoryData struct {
	StockId   string  `json:"StockId,omitempty"`
	Timestamp float64 `json:"Timestamp,omitempty"`
}

type MongoUserEvent struct {
	Key          string            `bson:"_id"`
	History      []UserHistoryData `json:"history,omitempty"`
	VideoHistory []UserHistoryData `json:"videohistory,omitempty"`
	UpdatedAt    time.Time         `json:"updatedAt,omitempty" bson:"updatedAt"`
}

type MongoUserEventForCategoryNew struct {
	Key                   string            `bson:"_id"`
	Interested_Categories interface{}       `json:"interested_categories" bson:"interested_categories"`
	History               []UserHistoryData `json:"history,omitempty"`
	VideoHistory          []UserHistoryData `json:"videohistory,omitempty"`
	Preferences           []string          `json:"preferences,omitempty"`
}

type MongoUserEventForCategory struct {
	Key                   string            `bson:"_id"`
	Interested_Categories json.RawMessage   `json:"interested_categories"`
	History               []UserHistoryData `json:"history,omitempty"`
	VideoHistory          []UserHistoryData `json:"videohistory,omitempty"`
	Preferences           []string          `json:"preferences,omitempty"`
}

type VideoData struct {
	AccountId    string `json:"account_id,omitempty"`
	VideoId      string `json:"video_id,omitempty"`
	Description  string `json:"description,omitempty"`
	Duration     int64  `json:"duration,omitempty"`
	ThumbnailUrl string `json:"thumbnail_url,omitempty"`
	Headline     string `json:"name,omitempty"`
	PublishedAt  int64  `json:"published_at,omitempty"`
	State        string `json:"state,omitempty"`
	Section_name string `json:"section_name,omitempty"`
	Tag_name     string `json:"tags,omitempty"`
	UpdatedAt    int64  `json:"updated_at,omitempty"`
}

type IndividualScore struct {
	RecencyScore  string `json:"recency_score,omitempty"`
	VectorScore   string `json:"vector_score,omitempty"`
	CategoryScore string `json:"category_score,omitempty"`
	ReadScore     string `json:"read_score,omitempty"`
	TagsScore     string `json:"tags_score,omitempty"`
	TitleScore    string `json:"title_score,omitempty"`
	AuthorScore   string `json:"author_score,omitempty"`
}

type JobResult struct {
	ID         string
	Mid        string
	Suffix     string
	Collection string
	Result     interface{}
}

type TaggerRecommendation struct {
	Cid             string `json:"cid"`
	MerchantId      string `json:"merchant_id"`
	Timestamp       string `json:"tagger_timestamp,omitempty"`
	PublisherTags   string `json:"publisher_tags,omitempty"`
	RecommendedTags string `json:"recommended_tags,omitempty"`
	Headline        string `json:"headline,omitempty"`
	Slug            string `json:"slug,omitempty"`
	Tags            string `json:"tags,omitempty"`
}

type UserSession struct {
	Timestamp float64
	SessionId string
}

type ClusterArticles struct {
	Key      string               `json:"_id,omitempty"`
	Articles []ArticleClusterData `json:"articles,omitempty"`
}

type ArticleClusterData struct {
	ArticleId   string `json:"cid,omitempty"`
	Relevance   string `json:"relevance,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	PublishedAt string `json:"publishedAt,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
	Tags        string `json:"tags,omitempty"`
}

type MerchantWiseConfig struct {
	NoOfSimilarContents          int64
	NoOfCollabCategories         int64
	NoOfCollabContents           int64
	NoOfTopInterests             int64
	ContentsForEveryUserInterest int64
	UseUniqueUserInterests       bool
	UseInterestBasedTags         bool
	ArticleRecencyInDays         int
	FilterHistoryInDays          int
	RecentHistoryCountInDays     int
}

type SimsJobRequest struct {
	MerchantID     string   `json:"merchantId"`
	SimilarityType string   `json:"similarityType"`
	JobType        int      `json:"jobType"`
	Filter         int      `json:"filter"`
	CodePath       string   `json:"codePath"`
	TopicName      string   `json:"topicName"`
	ParamMap       ParamMap `json:"paramMap"`
}

type SimsUsingVectorizerRequest struct {
	Vector       string `json:"vector"`
	Merchant     string `json:"merchant"`
	NoOfArticles int    `json:"no_of_articles"`
}

type SimsUsingVectorizerResponse struct {
	Cids string `json:"sim_cids"`
}

type ParamMap struct {
	JobId  string `json:"jobId"`
	Vector string `json:"vector"`
	Limit  string `json:"limit"`
}

type SuggestedImagesJobResponse struct {
	Status  string                     `json:"status"`
	Results []SuggestedImagesJobResult `json:"result,omitempty"`
}

type SuggestedImagesResponse struct {
	Results []SuggestedImagesJobResult `json:"result,omitempty"`
}

type SuggestedImagesJobResult struct {
	Slug     string `json:"Slug,omitempty"`
	ImageURL string `json:"ImageUrl,omitempty"`
}

type Rules struct {
	Id         float64 `json:"id,omitempty"`
	ColumnName string  `json:"columnName,omitempty"`
	DataType   string  `json:"dataType,omitempty"`
	Weightage  float64 `json:"weightage,omitempty"`
	Algo       float64 `json:"algo,omitempty"`
}

type MongoRules struct {
	Key   string  `bson:"_id"`
	Rules []Rules `json:"rules,omitempty"`
}

type MongoBookmarks struct {
	Key       string     `bson:"_id"`
	Bookmarks []Bookmark `json:"bookmarks"`
}

type Bookmark struct {
	Cid       string  `json:"Cid"`
	Timestamp float64 `json:"Timestamp"`
}

type BookmarksResponse struct {
	Bookmarks []BookmarkResponse `json:"bookmarks"`
}

type BookmarkResponse struct {
	Cid              string `json:"Cid,omitempty"`
	Description      string `json:"Description,omitempty"`
	PublishedAt      string `json:"PublishedAt,omitempty"`
	Slug             string `json:"Slug,omitempty"`
	ImageUrl         string `json:"ImageUrl,omitempty"`
	FirstPublishedAt string `json:"first_published_at,omitempty"`
	ImageMetaData    string `json:"ImageMetaData,omitempty"`
}

type RulesResponse struct {
	MerchantId string `json:"merchantId,omitempty"`
}

type SlackMessage struct {
	Subject string         `json:"text"`
	Error   []ErrorMessage `json:"attachments,omitempty"`
}

type ErrorMessage struct {
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}

type DSSEventLog struct {
	MerchantId string
	DSS        string
	Status     *string
	Event      string
}


type ResultPageRecommendations struct {
	RecommendationType string  `bson:"recommendationType"`
	PRPercentage       float32 `bson:"percentage"`
	Preference         int64   `bson:"preference"`
}
type ResultMerchantConfig struct {
	MerchantId                   string `bson:"merchant_id"`
	NoOfSimilarContents          int    `bson:"no_of_similar_contents"`
	NoOfCollabCategories         int    `bson:"no_of_collab_categories"`
	NoOfCollabContents           int    `bson:"no_of_collab_contents"`
	NoOfTopInterests             int    `bson:"no_of_top_interests"`
	ContentsForEveryUserInterest int    `bson:"contents_for_every_user_interest"`
	UseUniqueUserInterests       int    `bson:"use_unique_user_interests"`
	UseInterestBasedTags         int    `bson:"use_interest_based_tags"`
	ArticleRecencyInDays         int    `bson:"article_recency_in_days"`
	FilterHistoryInDays          int    `bson:"filter_history_in_days"`
	RecentHistoryCountInDays     int    `bson:"recent_history_count_in_days"`
}
type ResultExperiments struct {
	PageTypeId  int    `bson:"page_type_id"`
	MerchantId  int    `bson:"merchant_id"`
	Strategy    string `bson:"strategy"`
	Description string `bson:"description"`
	CreatedAt   string `bson:"created_at"`
	Type        string `bson:"type"`
}

type ResultPageType struct {
	Strategy                string                        `bson:"strategy"`
	FieldIndicesForSimsType ResultFieldIndicesForSimsType `bson:"fieldIndicesForSimsType"`
}
type MongoPageType struct {
	Strategy string `bson:"strategy"`
	PageType string `bson:"pageType"`
}
type ResultFieldIndicesForSimsType struct {
	Field   string `bson:"field"`
	Indices int    `bson:"indices"`
}

type PollMerchant struct {
	Mid        string `bson:"mid"`
	Slug       string `bson:"slug"`
	Rediskey   string `bson:"rediskey"`
	Pick_limit int    `bson:"pick_limit"`
}

type ResultGetMerchants struct {
	MerchantId   string `bson:"merchant_id"`
	MerchantName string `bson:"merchant_name"`
}

type ResultWRUTaggerstruct struct {
	ContentId       string `bson:"content_id"`
	PublisherTags   string `bson:"publisher_tags"`
	Recommendedtags string `bson:"recommended_tags"`
	CreatedAt       string `bson:"created_at"`
}
