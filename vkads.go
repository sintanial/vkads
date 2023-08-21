package vkads

import (
	"bufio"
	"bytes"
	"context"
	"github.com/gabriel-vasile/mimetype"
	"gopkg.in/vansante/go-ffprobe.v2"
	"image"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const host = "https://ads.vk.com"

func IntRef(f int) *int             { return &f }
func Float64Ref(f float64) *float64 { return &f }

type authorizedRoundTripper struct {
	token Token
	rt    http.RoundTripper
}

func (t *authorizedRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return t.rt.RoundTrip(t.token.Sign(request))
}

type RequestOptions struct {
	url.Values
}

func NewRequestOptions(values ...url.Values) RequestOptions {
	var r RequestOptions
	if len(values) > 0 {
		r.Values = values[0]
	} else {
		r.Values = url.Values{}
	}
	return r
}

func (o RequestOptions) GetLimit() int {
	limit, _ := strconv.Atoi(o.Get("limit"))
	return limit
}

func (o RequestOptions) SetLimit(limit int) RequestOptions {
	o.Set("limit", strconv.Itoa(limit))
	return o
}

func (o RequestOptions) GetOffset() int {
	limit, _ := strconv.Atoi(o.Get("offset"))
	return limit
}

func (o RequestOptions) SetOffset(limit int) RequestOptions {
	o.Set("offset", strconv.Itoa(limit))
	return o
}

func (o RequestOptions) SetFields(fields []string) RequestOptions {
	o.Set("fields", strings.Join(fields, ","))
	return o
}

func (o RequestOptions) SetIdIn(ids []int) RequestOptions {
	o.Set("_id__in", strings.Join(IntToStringSlice(ids), ","))
	return o
}

func (o RequestOptions) SetStatus(statuses []string) RequestOptions {
	o.Set("_status", strings.Join(statuses, ","))
	return o
}

func (o RequestOptions) SetSorting(sorting []string) RequestOptions {
	o.Set("_sorting", strings.Join(sorting, ","))
	return o
}

type Iterable[T any] struct {
	Items  T   `json:"items"`
	Count  int `json:"count"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Iterator[T any] struct {
	InitialLimit  int
	InitialOffset int

	lastResponse *Iterable[T]

	next func(limit int, offset int) (*Iterable[T], error)
}

func (self *Iterator[T]) HasNext() bool {
	if self.lastResponse == nil {
		return true
	}

	if self.lastResponse.Limit <= 0 {
		return false
	}

	return self.lastResponse.Limit+self.lastResponse.Offset < self.lastResponse.Count
}

func (self *Iterator[T]) Next() (*Iterable[T], error) {
	if !self.HasNext() {
		return nil, nil
	}

	limit := self.InitialLimit
	offset := self.InitialOffset

	offsetlimit := 0
	if self.lastResponse != nil {
		limit = self.lastResponse.Limit
		offset = self.lastResponse.Offset

		offsetlimit = offset + limit
	}

	response, err := self.next(limit, offsetlimit)
	if err != nil {
		return nil, err
	}

	self.lastResponse = response
	return response, nil
}

func (self *Iterator[T]) All() ([]*Iterable[T], error) {
	var result []*Iterable[T]
	for self.HasNext() {
		data, err := self.Next()
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}

func MakeContentOptions(tp ContentMethod, br *bufio.Reader) (opt ContentOptions, err error) {
	peek, err := br.Peek(4096)
	if err != nil {
		return opt, err
	}

	mm := mimetype.Detect(peek)
	opt.MimeType = mm.String()
	opt.Ext = mm.Extension()

	if tp == ContentMethodStatic {
		imageCfg, _, err := image.DecodeConfig(bufio.NewReader(bytes.NewReader(peek)))
		if err != nil {
			return opt, err
		}

		opt.Width = imageCfg.Width
		opt.Height = imageCfg.Height
	} else if tp == ContentMethodVideo {
		data, err := ioutil.ReadAll(br)
		if err != nil {
			return opt, err
		}

		reader := bytes.NewReader(data)

		videoProbe, err := ffprobe.ProbeReader(context.Background(), reader)
		if err != nil {
			return opt, err
		}

		stream := videoProbe.Streams[0]
		opt.Width = stream.Width
		opt.Height = stream.Height

		reader.Seek(0, io.SeekStart)
		br.Reset(reader)
	}

	return opt, nil
}

type CropableImage interface {
	image.Image
	SubImage(r image.Rectangle) image.Image
}

type CropOption struct {
	X      int
	Y      int
	Width  int
	Height int
}

func CropImage(r io.Reader, width int, height int) (io.Reader, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	img.(CropableImage).SubImage(image.Rect(0, 0, width, height))

	return nil, nil
}

func IntToStringSlice(i []int) []string {
	var result []string
	for _, val := range i {
		result = append(result, strconv.Itoa(val))
	}

	return result
}
