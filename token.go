package vkads

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Token struct {
	AccessToken  string   `json:"access_token"`
	TokenType    string   `json:"token_type"`
	Scope        []string `json:"scope"`
	ExpiresIn    *int     `json:"expires_in"`
	RefreshToken string   `json:"refresh_token"`
	TokensLeft   int      `json:"tokens_left"`
}

func (self Token) Sign(req *http.Request) *http.Request {
	req.Header.Set("Authorization", "Bearer "+self.AccessToken)
	return req
}

type TokenHolder interface {
	Store(Token) error
	Retrieve() (Token, error)
}

type JsonFileTokenHolder struct {
	File string
}

func (j JsonFileTokenHolder) Store(token Token) error {
	f, err := os.OpenFile(j.File, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	return json.NewEncoder(f).Encode(token)
}

func (j JsonFileTokenHolder) Retrieve() (t Token, err error) {
	var f *os.File
	f, err = os.Open(j.File)
	if err != nil {
		return
	}

	err = json.NewDecoder(f).Decode(&t)
	return
}

func NewJsonFileTokenHolder(file string) *JsonFileTokenHolder {
	return &JsonFileTokenHolder{File: file}
}

type AuthCodeFlow struct {
	clientId     string
	clientSecret string
}

func NewAuthCodeFlow(clientId string, clientSecret string) *AuthCodeFlow {
	return &AuthCodeFlow{clientId: clientId, clientSecret: clientSecret}
}

func (self *AuthCodeFlow) ClientCredentialsGrantToken(permanent bool) (Token, error) {
	body := url.Values{}
	body.Set("grant_type", "client_credentials")
	body.Set("client_id", self.clientId)
	body.Set("client_secret", self.clientSecret)
	if permanent {
		body.Set("permanent", "true")
	}

	return self.getToken("/api/v2/oauth2/token.json", body)
}

func (self *AuthCodeFlow) AgencyClientCredentialsGrantToken(permanent bool, agencyClientName string, agencyClientId int) (Token, error) {
	body := url.Values{}
	body.Set("grant_type", "agency_client_credentials")
	body.Set("client_id", self.clientId)
	body.Set("client_secret", self.clientSecret)
	if permanent {
		body.Set("permanent", "true")
	}

	if agencyClientName != "" {
		body.Set("agency_client_name", agencyClientName)
	} else if agencyClientId != 0 {
		body.Set("agency_client_id", strconv.Itoa(agencyClientId))
	} else {
		return Token{}, errors.New("agency_client_name or agency_client_id must be set")
	}

	return self.getToken("/api/v2/oauth2/token.json", body)
}

func (self *AuthCodeFlow) DeleteTokens(username string, userId int) error {
	body := url.Values{}
	body.Set("client_id", self.clientId)
	body.Set("client_secret", self.clientSecret)
	if username != "" {
		body.Set("username", username)
	} else if userId != 0 {
		body.Set("user_id", strconv.Itoa(userId))
	}

	return self.doRequest("/api/v2/oauth2/token/delete.json", body, nil)
}

func (self *AuthCodeFlow) getToken(uri string, body url.Values) (Token, error) {
	var token Token
	if err := self.doRequest(uri, body, &token); err != nil {
		return Token{}, err
	}

	return token, nil
}

func (self *AuthCodeFlow) doRequest(uri string, params url.Values, result interface{}) error {
	resp, err := http.DefaultClient.Post(
		host+uri,
		"application/x-www-form-urlencoded",
		strings.NewReader(params.Encode()),
	)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		body, _ := httputil.DumpResponse(resp, true)
		return errors.New("invalid status code: " + resp.Status + ", params:" + string(body))
	}

	if result == nil {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, result)
}
