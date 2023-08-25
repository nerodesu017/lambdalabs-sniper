package constants

type LAMBDA_LABS_API_URLS string

const (
	HOST_NAME          LAMBDA_LABS_API_URLS = "cloud.lambdalabs.com"
	BASE_URL           LAMBDA_LABS_API_URLS = "https://" + HOST_NAME
	SPA_INIT_INFO_URL  LAMBDA_LABS_API_URLS = BASE_URL + "/api/cloud/spa-init-info" // for email
	INSTANCE_TYPES_URL LAMBDA_LABS_API_URLS = BASE_URL + "/api/v1/instance-types"   // for getting currently available instances
)