package vkobj

type AdditionalUserInfo struct {
	Address    string `json:"address"`
	ClientInfo string `json:"client_info"`
	ClientName string `json:"client_name"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
}

type User struct {
	Id            int      `json:"id"`
	Username      string   `json:"username"`
	Language      string   `json:"language"`
	Firstname     string   `json:"firstname"`
	Lastname      string   `json:"lastname"`
	Email         string   `json:"email"`
	Types         []string `json:"types"`
	Status        string   `json:"status"`
	InfoCurrency  string   `json:"info_currency"`
	Currency      string   `json:"currency"`
	Timezone      int      `json:"timezone"`
	Country       int      `json:"country"`
	EmailSettings []struct {
		Type                    string `json:"type"`
		Email                   string `json:"email"`
		EmailVerificationStatus string `json:"email_verification_status"`
	} `json:"email_settings"`
	Mailings struct {
		News struct {
			Email []interface{} `json:"email"`
		} `json:"news"`
		AdvCampaigns struct {
			Email []string `json:"email"`
		} `json:"adv_campaigns"`
		LeadAds struct {
			Email []interface{} `json:"email"`
		} `json:"lead_ads"`
		Finance struct {
			Email []string `json:"email"`
		} `json:"finance"`
		ManagementRule struct {
			Email []string `json:"email"`
		} `json:"management_rule"`
		Event struct {
			Email []string `json:"email"`
		} `json:"event"`
		Moderation struct {
			Email []interface{} `json:"email"`
		} `json:"moderation"`
		Other struct {
			Email []interface{} `json:"email"`
		} `json:"other"`
		ApiChanges struct {
			Email []interface{} `json:"email"`
		} `json:"api_changes"`
	} `json:"mailings"`
	Regions struct {
		Allowed       []int `json:"allowed"`
		Required      []int `json:"required"`
		RequiredOneOf []int `json:"required_one_of"`
	} `json:"regions"`
}

type AgencyClient struct {
	AccessType string `json:"access_type"`
	Status     string `json:"status"`
	User       User   `json:"user"`
}

type AdPlanVkadsStatus struct {
	Codes       []string `json:"codes"`
	MajorStatus string   `json:"major_status"`
	Status      string   `json:"status"`
}

type AdPlanOrStatus struct {
	Codes       []string `json:"codes"`
	MajorStatus string   `json:"major_status"`
	Status      string   `json:"status"`
}

type AdPlanIssue struct {
	Arguments struct{} `json:"arguments"`
	Code      string   `json:"code"`
	Message   string   `json:"message"`
}

type Objective string

const ObjectiveSiteConversions Objective = "site_conversions"

type UrlTypes = struct {
	Primary                 [][]string `json:"primary,omitempty"`
	PrimaryDynamic          [][]string `json:"primary_dynamic,omitempty"`
	ShopitemUrl             [][]string `json:"shopitem_url,omitempty"`
	SlideClick              [][]string `json:"slide_click,omitempty"`
	UrlStaticSlide          [][]string `json:"url_static_slide,omitempty"`
	DeeplinkStaticSlide     [][]string `json:"deeplink_static_slide,omitempty"`
	DeeplinkStaticSlideRmkt [][]string `json:"deeplink_static_slide_rmkt,omitempty"`
	IosStoreUrl             [][]string `json:"ios_store_url,omitempty"`
	IosTrackingUrl          [][]string `json:"ios_tracking_url,omitempty"`
	IosUrl                  [][]string `json:"ios_url,omitempty"`
	AndroidStoreUrl         [][]string `json:"android_store_url,omitempty"`
	AndroidTrackingUrl      [][]string `json:"android_tracking_url,omitempty"`
	AndroidUrl              [][]string `json:"android_url,omitempty"`
	VkPost                  [][]string `json:"vk_post,omitempty"`
	DeeplinkUrl             [][]string `json:"deeplink_url,omitempty"`
	HeaderClick             [][]string `json:"header_click,omitempty"`
	Slide1Click             [][]string `json:"slide_1_click,omitempty"`
	Slide2Click             [][]string `json:"slide_2_click,omitempty"`
	Slide3Click             [][]string `json:"slide_3_click,omitempty"`
	UrlSlide1               [][]string `json:"url_slide_1,omitempty"`
	UrlSlide2               [][]string `json:"url_slide_2,omitempty"`
	UrlSlide3               [][]string `json:"url_slide_3,omitempty"`
	Ok                      [][]string `json:"ok,omitempty"`
	Vk                      [][]string `json:"vk1,omitempty"`
}

type PricedGoal struct {
	Name     string `json:"name"`
	SourceID int    `json:"source_id"`
}

type PricedGoalName string

const PricedGoalNameShow PricedGoalName = "shows"
const PricedGoalNameClick PricedGoalName = "clicks"
const PricedGoalNameAgSuccess PricedGoalName = "ag:success"

type AdPlanCampaign struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	PackageId int    `json:"package_id,omitempty"`
}

type AutobiddingMode string

const AutobiddingModeMaxGoals AutobiddingMode = "max_goals"

type AdPlanStatus string

const AdPlanStatusActive AdPlanStatus = "active"
const AdPlanStatusDeleted AdPlanStatus = "deleted"
const AdPlanStatusBlocked AdPlanStatus = "blocked"

type AdPlan struct {
	Id              int             `json:"id,omitempty"`
	Name            string          `json:"name"`
	Status          AdPlanStatus    `json:"status,omitempty"`
	VkadsStatus     interface{}     `json:"vkads_status,omitempty"`
	Objective       Objective       `json:"objective,omitempty"`
	AutobiddingMode AutobiddingMode `json:"autobidding_mode,omitempty"`
	BudgetLimit     *float64        `json:"budget_limit,omitempty"`
	BudgetLimitDay  *float64        `json:"budget_limit_day,omitempty"`
	MaxPrice        *float64        `json:"max_price,string,omitempty"`
	DateStart       Date            `json:"date_start,omitempty"`
	DateEnd         *Date           `json:"date_end,omitempty"`
	PricedGoal      PricedGoal      `json:"priced_goal,omitempty"`
	AdGroups        []AdGroup       `json:"ad_groups,omitempty"`
	Created         string          `json:"created,omitempty"`
	Updated         string          `json:"updated,omitempty"`
}

type AgeTargeting struct {
	AgeList []int `json:"age_list"`
	Expand  bool  `json:"expand"`
}

func AgeList(from int, to int) []int {
	var s []int
	for i := from; i <= to; i++ {
		s = append(s, i)
	}

	return s
}

var DefaultAgeList = AgeList(12, 75)

type FulltimeTargetingFlag = string

const FulltimeTargetingFlagUseHolidaysMoving = "use_holidays_moving"
const FulltimeTargetingFlagCrossTimezone = "cross_timezone"

type FulltimeTargeting struct {
	Flags []FulltimeTargetingFlag `json:"flags"`
	Fri   []int                   `json:"fri"`
	Mon   []int                   `json:"mon"`
	Sat   []int                   `json:"sat"`
	Sun   []int                   `json:"sun"`
	Thu   []int                   `json:"thu"`
	Tue   []int                   `json:"tue"`
	Wed   []int                   `json:"wed"`
}

var DefaultFulltimeTargetings = FulltimeTargeting{
	Flags: []FulltimeTargetingFlag{FulltimeTargetingFlagUseHolidaysMoving, FulltimeTargetingFlagCrossTimezone},
	Fri:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
	Mon:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
	Sat:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
	Sun:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
	Thu:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
	Tue:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
	Wed:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
}

type GeoTargeting struct {
	Regions []int `json:"regions"`
}

type Sex string

const SexMale Sex = "male"
const SexFemale Sex = "female"

type Targetings struct {
	Age       AgeTargeting      `json:"age,omitempty"`
	Fulltime  FulltimeTargeting `json:"fulltime,omitempty"`
	Geo       GeoTargeting      `json:"geo,omitempty"`
	Interests []int             `json:"interests,omitempty"`
	Pads      []int             `json:"pads,omitempty"`
	Segments  []int             `json:"segments,omitempty"`
	Sex       []Sex             `json:"sex,omitempty"`
}

type AgeRestriction string

const AgeRestriction0 AgeRestriction = "0+"
const AgeRestrictio6 AgeRestriction = "6+"
const AgeRestriction12 AgeRestriction = "12+"
const AgeRestriction16 AgeRestriction = "16+"
const AgeRestriction18 AgeRestriction = "18+"

type AdGroup struct {
	Id              int             `json:"id,omitempty"`
	Name            string          `json:"name"`
	Status          string          `json:"status,omitempty"`
	AdPlanId        Int             `json:"ad_plan_id,omitempty"`
	PackageId       int             `json:"package_id,omitempty"`
	AgeRestrictions AgeRestriction  `json:"age_restrictions,omitempty"`
	AutobiddingMode AutobiddingMode `json:"autobidding_mode,omitempty"`
	BudgetLimit     *Float64        `json:"budget_limit,omitempty"`
	BudgetLimitDay  *Float64        `json:"budget_limit_day,omitempty"`
	MaxPrice        *Float64        `json:"max_price,omitempty"`
	DateStart       Date            `json:"date_start,omitempty"`
	DateEnd         *Date           `json:"date_end,omitempty"`
	Objective       Objective       `json:"objective,omitempty"`
	EnableUtm       bool            `json:"enable_utm,omitempty"`
	Utm             *string         `json:"utm,omitempty"`
	Social          bool            `json:"social,omitempty"`
	Targetings      Targetings      `json:"targetings,omitempty"`
	Banners         []Banner        `json:"banners"`
	Created         string          `json:"created,omitempty"`
	Updated         string          `json:"updated,omitempty"`
}

type ContentType = string

const ContentTypeStatic ContentType = "static"
const ContentTypeAnimated ContentType = "animated"
const ContentTypeRollovered ContentType = "rollovered"
const ContentTypeVideo ContentType = "video"
const ContentTypeAudio ContentType = "audio"
const ContentTypeHtml5 ContentType = "html5"

type ContentVariant struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Size   int    `json:"size"`
	Url    string `json:"url"`
}

type Content struct {
	Id       int                       `json:"id"`
	Variants map[string]ContentVariant `json:"variants,omitempty"`
}

type BannerContent struct {
	Content
	Type ContentType `json:"type,omitempty"`
}

type Textblock struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

type Urls struct {
	Id            int      `json:"id"`
	Url           string   `json:"url,omitempty"`
	UrlObjectId   string   `json:"url_object_id,omitempty"`
	UrlObjectType string   `json:"url_object_type,omitempty"`
	UrlTypes      []string `json:"url_types,omitempty"`
}

type Banner struct {
	Id               int                      `json:"id,omitempty"`
	Name             string                   `json:"name"`
	Status           string                   `json:"status,omitempty"`
	AdGroupId        int                      `json:"ad_group_id,omitempty"`
	Content          map[string]BannerContent `json:"content,omitempty"`
	Delivery         string                   `json:"delivery,omitempty"`
	Issues           interface{}              `json:"issues,omitempty"`
	ModerationStatus string                   `json:"moderation_status,omitempty"`
	Textblocks       map[string]Textblock     `json:"textblocks,omitempty"`
	Urls             map[string]Urls          `json:"urls,omitempty"`
}

type Region struct {
	Id        int      `json:"id"`
	ParentId  int      `json:"parent_id"`
	Name      string   `json:"name"`
	IsoAlpha3 string   `json:"iso_alpha_3"`
	Type      string   `json:"type"`
	Flags     []string `json:"flags"`
}

type TargetingsTreeElement struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Synonyms []string `json:"synonyms"`
	Children []struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"children"`
}

type TargetingsTree struct {
	Interests       []TargetingsTreeElement `json:"interests"`
	InterestsShort  []TargetingsTreeElement `json:"interests_short"`
	InterestsSocDem []TargetingsTreeElement `json:"interests_soc_dem"`
	InterestsStable []TargetingsTreeElement `json:"interests_stable"`
}

type TargetingsTreeResponse []TargetingsTree

type SegmentRelation struct {
	Id         int         `json:"id"`
	ObjectId   int         `json:"object_id"`
	ObjectType string      `json:"object_type"`
	Params     interface{} `json:"params"`
}

type SegmentUser struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	Username string `json:"username"`
}

type Segment struct {
	Id             int               `json:"id"`
	Name           string            `json:"name"`
	CampaignIds    []int             `json:"campaign_ids"`
	Flags          []string          `json:"flags"`
	PassCondition  int               `json:"pass_condition"`
	Relations      []SegmentRelation `json:"relations"`
	RelationsCount int               `json:"relations_count"`
	Users          []SegmentUser     `json:"users"`
	Created        DateTime          `json:"created"`
	Updated        DateTime          `json:"updated"`
}

type PackagePad struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Patterns    string `json:"patterns"`
	EyeUrl      struct {
		Id          int    `json:"id,omitempty"`
		Description string `json:"description,omitempty"`
		Url         string `json:"url,omitempty"`
	} `json:"eye_url"`
}

type Package struct {
	Id                      int         `json:"id"`
	Name                    string      `json:"name"`
	Description             string      `json:"description"`
	Price                   string      `json:"price"`
	PricedEventType         int         `json:"priced_event_type"`
	PaidEventType           int         `json:"paid_event_type"`
	MaxPricePerUnit         string      `json:"max_price_per_unit"`
	MaxUniqShowsLimit       int         `json:"max_uniq_shows_limit"`
	MaxBannersInOneCampaign interface{} `json:"max_banners_in_one_campaign"`
	RelatedPackageIds       []int       `json:"related_package_ids"`
	Flags                   []string    `json:"flags"`
	UrlType                 interface{} `json:"url_type"`
	BannerFormatId          int         `json:"banner_format_id"`
	BannerUrlGetParams      string      `json:"banner_url_get_params"`
	PadsTreeId              int         `json:"pads_tree_id"`
	UrlTypes                UrlTypes    `json:"url_types"`
	Status                  string      `json:"status"`
	Objective               []string    `json:"objective"`
	Created                 string      `json:"created"`
	Updated                 string      `json:"updated"`
}

type TopmainlruGoal struct {
	Goal        string `json:"goal"`
	Description string `json:"description"`
	Id          int    `json:"id"`
	CounterId   int    `json:"counter_id"`
	CounterName string `json:"counter_name"`
}

type Goals struct {
	Topmailru     []TopmainlruGoal `json:"topmailru"`
	MobileInstall []interface{}    `json:"mobile_install"`
	OkGroup       []interface{}    `json:"ok_group"`
	OkGame        []interface{}    `json:"ok_game"`
}

type BannerPattern struct {
	Description string `json:"description"`
	Format      []struct {
		Field    string `json:"field"`
		Required bool   `json:"required"`
		Role     string `json:"role"`
	} `json:"format"`
	Id        int `json:"id"`
	Interface struct {
		ProjectionFactor float64 `json:"projectionFactor,omitempty"`
	} `json:"interface"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type BannerField struct {
	Description string `json:"description"`
	Field       string `json:"field"`
	Format      struct {
		MaxLength int `json:"max_length,omitempty"`
		Text      struct {
			MaxLength int      `json:"max_length,omitempty"`
			Choices   []string `json:"choices,omitempty"`
		} `json:"text,omitempty"`
		MaxHeight int         `json:"max_height,omitempty"`
		MaxWidth  int         `json:"max_width,omitempty"`
		MinHeight int         `json:"min_height,omitempty"`
		MinWidth  int         `json:"min_width,omitempty"`
		Type      interface{} `json:"type,omitempty"`
		Variants  map[string]struct {
			Height int `json:"height"`
			Width  int `json:"width"`
		} `json:"variants,omitempty"`
		Height        int    `json:"height,omitempty"`
		Size          int    `json:"size,omitempty"`
		Width         int    `json:"width,omitempty"`
		AspectRatio   string `json:"aspect_ratio,omitempty"`
		MinLength     int    `json:"min_length,omitempty"`
		Choices       []int  `json:"choices,omitempty"`
		FromPricelist bool   `json:"from_pricelist,omitempty"`
	} `json:"format"`
	Id        int `json:"id"`
	Interface struct {
		Placeholder  string `json:"placeholder,omitempty"`
		Type         string `json:"type"`
		Help         string `json:"help,omitempty"`
		AutoComplete bool   `json:"autoComplete,omitempty"`
		Affect       []struct {
			Field string `json:"field"`
			Path  string `json:"path"`
			Role  string `json:"role"`
		} `json:"affect,omitempty"`
	} `json:"interface"`
	Role   string `json:"role"`
	Status string `json:"status"`
}
