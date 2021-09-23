package config

//Format String



//MongoDB
const STRONG_MODE = true
const MONGO_SIMS_DATABASE = "sims"
const MONGO_MERCHANT_COLLECTION = "merchant"
const MONGO_MERCHANT_PLATFORM_COLLECTION = "merchant_platform"
const MONGO_CUSTOMERS_ROLE_COLLECTION = "customers_role"
const MONGO_CUSTOMERS_COLLECTION ="customers"
const MONGO_MERCHANT_CONFIG_COLLECTION = "merchant_config"
const MONGO_PAGE_TYPE_COLLECTION = "pageType"
const MONGO_PAGE_RECOMMENDATIONS_COLLECTION = "page_recommendations"
const MONGO_MERCHANT_PAGES_COLLECTION = "merchant_pages"
const MONGO_ROUTE_INFO_COLLECTION = "route_info"


//Onborading configs
const ROUTE_STATS="/Statistics,/Strategies,/ConversionRatio,/TaggerUsage,/Segmentation"
const ROUTE_CHART ="/Prediction,/Hit,/Accuracy"
const REGISTER_REQUESTED = "REQUESTED"
const REGISTER_IN_PROCESS = "INPROCESS"
const REGISTER_ACTIVE ="ACTIVE"
const REGISTER_INACTIVE ="INACTIVE"
const REGISTER_REJECTED ="REJECTED"
const ARTICLE_SIMS_CONTENT_KEY= "%s:isimsContent:%s"
const ARTICLE_SIMS_RECENCY_KEY= "%s:isimsRecency:%s"
const VIDEO_SIMS_CONTENT_KEY= "%s:videosimsContent:%s"
const VIDEO_SIMS_RECENCY_KEY="%s:videosimsRecency:%s"
const REGISTER_ACTION_ACCEPT ="accept"
const REGISTER_ACTION_REJECT ="reject"

//Logger
const DIR_NAME = "./log"
const MAX_FILES = 20
const FILES_TO_DELETE = 1
const MAX_SIZE = 100
const LOGS_BUCKET = "dss-data-dump"
