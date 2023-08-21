package vkads

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sintanial/vkads/vkobj"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strconv"
	"strings"
)

type Api struct {
	token Token
	http  *http.Client
	debug bool
}

func New(token Token) *Api {
	return NewWithHttpClient(token, nil)
}

type AuthOptions struct {
	ClientId     string
	ClientSecret string
}

func NewAuth(th TokenHolder, options AuthOptions, http *http.Client) (*Api, error) {
	token, err := th.Retrieve()
	if err == nil {
		return NewWithHttpClient(token, http), nil
	}

	acf := NewAuthCodeFlow(options.ClientId, options.ClientSecret)
	if err := acf.DeleteTokens("", 0); err != nil {
		return nil, err
	}

	token, err = acf.ClientCredentialsGrantToken(true)
	if err != nil {
		return nil, err
	}

	if err := th.Store(token); err != nil {
		return nil, err
	}

	return NewWithHttpClient(token, http), nil
}

func NewWithHttpClient(token Token, client *http.Client) *Api {
	if client == nil {
		client = &http.Client{}
	}

	var rt http.RoundTripper
	if client.Transport != nil {
		rt = client.Transport
	} else {
		rt = http.DefaultTransport
	}

	client.Transport = &authorizedRoundTripper{
		token: token,
		rt:    rt,
	}

	return &Api{
		token: token,
		http:  client,
		debug: true,
	}
}

func (self *Api) Debug(b bool) {
	self.debug = b
}

func (self *Api) GetUser() (response vkobj.User, err error) {
	err = self.getRequestUnmarshal("/api/v3/user.json", &response)
	return
}

func (self *Api) GetAgencyClients() Iterator[[]vkobj.AgencyClient] {
	return createApiIterator[[]vkobj.AgencyClient](self, "/api/v2/agency/clients.json")
}

type UpdateBannerMassRequest []struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func (self *Api) UpdateBannersMassAction(req UpdateBannerMassRequest) error {
	return self.postJsonRequestUnmarshal("/api/v2/banners/mass_action.json", nil, req)
}

type ContentMethod string

const ContentMethodStatic ContentMethod = "static"
const ContentMethodVideo ContentMethod = "video"
const ContentMethodHtml5 ContentMethod = "html5"

type ContentOptions struct {
	MimeType string `json:"mime_type"`
	Ext      string `json:"ext"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

func (self *Api) CreateContent(tp ContentMethod, content io.Reader, opt ContentOptions) (response vkobj.Content, err error) {
	br := bufio.NewReader(content)

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	data, err := json.Marshal(map[string]int{
		"width":  opt.Width,
		"height": opt.Height,
	})
	if err != nil {
		return response, err
	}

	if err = w.WriteField("data", string(data)); err != nil {
		return response, err
	}

	fileHeader := make(textproto.MIMEHeader)
	fileHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="file%s"`, opt.Ext))
	fileHeader.Set("Content-Type", opt.MimeType)
	filePart, err := w.CreatePart(fileHeader)
	if err != nil {
		return response, err
	}

	if _, err = io.Copy(filePart, br); err != nil {
		return
	}

	if err = w.Close(); err != nil {
		return response, err
	}

	err = self.postMultipartRequestUnmarshal("/api/v2/content/"+string(tp)+".json", w.FormDataContentType(), &buf, &response)
	return
}

type GetAdPlansResponse struct {
	Count  int            `json:"count"`
	Items  []vkobj.AdPlan `json:"items"`
	Offset int            `json:"offset"`
}

func (self *Api) GetAdPlans(options ...RequestOptions) Iterator[[]vkobj.AdPlan] {
	return createApiIterator[[]vkobj.AdPlan](self, "/api/v2/ad_plans.json", options...)
}

type CreateAdPlanResponse struct {
	Id int `json:"id"`
}

func (self *Api) CreateAdPlan(adPlan vkobj.AdPlan) (response CreateAdPlanResponse, err error) {
	err = self.postJsonRequestUnmarshal("/api/v2/ad_plans.json", &response, adPlan)
	return
}

var AdPlanAllFieldsOption = []string{
	"id",
	"created",
	"updated",
	"name",
	"status",
	"vkads_status",
	"ad_groups",
	"autobidding_mode",
	"budget_limit",
	"budget_limit_day",
	"date_start",
	"date_end",
	"max_price",
	"objective",
	"priced_goal",
	"pricelist_id",
}

func (self *Api) GetAdPlan(adPlanId int, options ...RequestOptions) (response vkobj.AdPlan, err error) {
	err = self.getRequestUnmarshal("/api/v2/ad_plans/"+strconv.Itoa(adPlanId)+".json", &response, options...)
	return
}

type UpdateAdPlanResponse vkobj.AdPlan

func (self *Api) UpdateAdPlan(adPlanId int, plan vkobj.AdPlan) (response UpdateAdPlanResponse, err error) {
	err = self.postJsonRequestUnmarshal("/api/v2/ad_plan/"+strconv.Itoa(adPlanId)+".json", &response, plan)
	return
}

type GetPackagesPadsResponse struct {
	Items []vkobj.PackagePad `json:"items"`
}

func (self *Api) GetPackagesPads() (response GetPackagesPadsResponse, err error) {
	err = self.getRequestUnmarshal("/api/v2/packages_pads.json", &response)
	return
}

type GetPadsTreeResponse struct {
	Items []struct {
		Tree []struct {
			Name     string `json:"name"`
			Children []struct {
				Name     string `json:"name,omitempty"`
				Children []struct {
					Id int `json:"id"`
				} `json:"children,omitempty"`
				Id int `json:"id,omitempty"`
			} `json:"children"`
		} `json:"tree"`
		Id int `json:"id"`
	} `json:"items"`
}

func (self *Api) GetPadsTree() (response GetPadsTreeResponse, err error) {
	err = self.getRequestUnmarshal("/api/v2/pads_trees.json", &response)
	return
}

type GetPackagesResponse struct {
	Items []vkobj.Package `json:"items"`
}

func (self *Api) GetPackages() (response GetPackagesResponse, err error) {
	err = self.getRequestUnmarshal("/api/v2/packages.json", &response)
	return
}

type UpdateBannerRemoderationResponse struct {
	Banners []struct {
		Id          int         `json:"id"`
		Remoderated interface{} `json:"remoderated"`
	} `json:"banners"`
}

func (self *Api) UpdateBannerRemoderation(bannerIds []int) (response UpdateBannerRemoderationResponse, err error) {
	var params []map[string]int
	for _, id := range bannerIds {
		params = append(params, map[string]int{"id": id})
	}
	err = self.postJsonRequestUnmarshal("/api/v2/banners/remoderate.json", &response, params)
	return
}

type BannerFieldsResponse struct {
	Count int `json:"count"`
	Items []vkobj.BannerField
}

func (self *Api) GetBannerFields() (response BannerFieldsResponse, err error) {
	err = self.getRequestUnmarshal("/api/v2/banner_fields.json", &response)
	return
}

type BannerPatternsResponse struct {
	Count int `json:"count"`
	Items []vkobj.BannerPattern
}

func (self *Api) GetBannerPatterns() (response BannerPatternsResponse, err error) {
	err = self.getRequestUnmarshal("/api/v2/banner_patterns.json", &response)
	return
}

func (self *Api) GetAdGroups(options ...RequestOptions) Iterator[[]vkobj.AdGroup] {
	return createApiIterator[[]vkobj.AdGroup](self, "/api/v2/ad_groups.json", options...)
}

var AdGroupAllFieldsOption = []string{
	"id",
	"created",
	"updated",
	"name",
	"status",
	"ad_plan_id",
	"package_id",
	"age_restrictions",
	"audit_pixels",
	"autobidding_mode",
	"banner_uniq_shows_limit",
	"budget_limit",
	"budget_limit_day",
	"date_end",
	"date_start",
	"delivery",
	"dynamic_banners_use_storelink",
	"dynamic_without_remarketing",
	"enable_offline_goals",
	"enable_utm",
	"issues",
	"language",
	"marketplace_app_client_id",
	"max_price",
	"objective",
	"package_priced_event_type",
	"price",
	"priced_goal",
	"pricelist_id",
	"sk_ad_campaign_id",
	"targetings",
	"uniq_shows_limit",
	"uniq_shows_period",
	"utm",
}

func (self *Api) GetAdGroup(adGroupId int, options ...RequestOptions) (response vkobj.AdGroup, err error) {
	err = self.getRequestUnmarshal("/api/v2/ad_groups/"+strconv.Itoa(adGroupId)+".json", &response, options...)
	return
}

type CreateAdGroupResponse struct {
	Id      int `json:"id"`
	Banners []struct {
		Id int `json:"id"`
	} `json:"banners"`
}

func (self *Api) CreateAdGroup(request vkobj.AdGroup) (response CreateAdGroupResponse, err error) {
	err = self.postJsonRequestUnmarshal("/api/v2/ad_groups.json", &response, request)
	return
}

var BannerAllFieldsOption = []string{
	"id",
	"created",
	"updated",
	"name",
	"status",
	"ad_group_id",
	"content",
	"delivery",
	"issues",
	"moderation_reasons",
	"moderation_status",
	"textblocks",
	"urls",
}

type BannersRequestOptions struct {
	RequestOptions
	AdGroupIdIn     []int
	AdGroupStatusIn []string
}

func (o BannersRequestOptions) SetAdGroupIdIn(ids []int) BannersRequestOptions {
	o.Set("_ad_group_id__in", strings.Join(IntToStringSlice(ids), ","))
	return o
}

func (o BannersRequestOptions) SetAdGroupStatusIn(statuses []string) BannersRequestOptions {
	o.Set("_ad_group_status__in", strings.Join(statuses, ","))
	return o
}

func (self *Api) GetBanners(options ...RequestOptions) Iterator[[]vkobj.Banner] {
	return createApiIterator[[]vkobj.Banner](self, "/api/v2/banners.json", options...)
}

type CreateBannerResponse vkobj.Banner

func (self *Api) CreateBanner(banner vkobj.Banner) (response CreateBannerResponse, err error) {
	err = self.postJsonRequestUnmarshal("/api/v2/banners.json", &response, banner)
	return
}

type GetBannerResponse vkobj.Banner

func (self *Api) GetBanner(bannerId int, options ...RequestOptions) (response GetBannerResponse, err error) {
	err = self.getRequestUnmarshal("/api/v2/banners/"+strconv.Itoa(bannerId)+".json", &response, options...)
	return
}

func (self *Api) UpdateBanner(bannerId int, banner vkobj.Banner) error {
	return self.postJsonRequestUnmarshal("/api/v2/banners/"+strconv.Itoa(bannerId)+".json", nil, banner)
}

func (self *Api) GetGoals() (response vkobj.Goals, err error) {
	err = self.getRequestUnmarshal("/api/v2/goals.json", &response)
	return
}

type GetRegionsResponse struct {
	Count int            `json:"count"`
	Items []vkobj.Region `json:"items"`
}

func (self *Api) GetRegions() (response GetRegionsResponse, err error) {
	err = self.getRequestUnmarshal("/api/v2/regions.json", &response)
	return
}

func (self *Api) GetTargetingsTree() (response vkobj.TargetingsTreeResponse, err error) {
	err = self.getRequestUnmarshal("/api/v2/targetings_tree.json", &response)
	return
}

func (self *Api) GetSegments() Iterator[[]vkobj.Segment] {
	return createApiIterator[[]vkobj.Segment](self, "/api/v2/remarketing/segments.json")
}

func (self *Api) getRequestUnmarshal(uri string, obj interface{}, options ...RequestOptions) error {
	u := host + uri
	if len(options) > 0 {
		var queries []string
		for _, val := range options {
			queries = append(queries, val.Encode())
		}

		u += "?" + strings.Join(queries, "&")
	}

	resp, err := self.http.Get(u)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return self.handleError(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, obj)
}

func (self *Api) postJsonRequestUnmarshal(uri string, obj interface{}, params interface{}) error {
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(params); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, host+uri, &b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := self.http.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return self.handleError(resp)
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	} else if obj == nil {
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, obj)
}

func (self *Api) postMultipartRequestUnmarshal(uri string, contentType string, params *bytes.Buffer, obj interface{}) error {
	req, err := http.NewRequest(http.MethodPost, host+uri, params)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Length", strconv.Itoa(params.Len()))

	res, err := self.http.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return self.handleError(res)
	}

	if res.StatusCode == http.StatusNoContent {
		return nil
	} else if obj == nil {
		return nil
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, obj)
}

func (self *Api) handleError(resp *http.Response) error {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var fr FailedResponse
	if err := json.Unmarshal(data, &fr); err != nil {
		return err
	}

	aerr := &ApiError{
		Code:    "",
		Message: "",
	}

	for key, val := range fr.Error {
		if key == "code" {
			aerr.Code = val.(string)
		} else if key == "message" {
			aerr.Message = val.(string)
		} else {
			if aerr.Extra == nil {
				aerr.Extra = make(map[string]interface{})
			}

			aerr.Extra[key] = val
		}
	}

	return aerr
}

func createApiIterator[T any](api *Api, uri string, options ...RequestOptions) Iterator[T] {
	initialLimit := 50
	if len(options) > 0 && options[0].GetLimit() > 0 {
		initialLimit = options[0].GetLimit()
	}

	return Iterator[T]{
		InitialLimit:  initialLimit,
		InitialOffset: 0,
		next: func(limit int, offset int) (*Iterable[T], error) {
			option := NewRequestOptions()
			if len(options) > 0 {
				option = options[0]
			}
			option.SetOffset(offset)

			var response Iterable[T]
			err := api.getRequestUnmarshal(uri, &response, option)
			return &response, err
		},
	}
}
