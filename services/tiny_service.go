package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"tinyurl/api"
)

type TinyService struct {
	tinyList []api.TinyStore
}

func (ts *TinyService) Store(param api.TinyStore) api.Response {
	if param.Url == "" {
		return api.Response{
			HttpCode:    400,
			Description: "url is not present",
			Data:        "",
		}
	}

	if _, _, err := ts.CheckShortcode(param.Shortcode); err == nil {
		return api.Response{
			HttpCode:    409,
			Description: "The the desired shortcode is already in use. Shortcodes are case-sensitive.",
			Data:        "",
		}
	}

	shortCode := param.Shortcode
	if param.Shortcode == "" {
		shortCode = ts.RandomShort()
	}

	checkShort := regexp.MustCompile("^[0-9a-zA-Z_]{6}$")
	if checkShort.MatchString(shortCode) != true {
		return api.Response{
			HttpCode:    422,
			Description: "The shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{6}$.",
			Data:        "",
		}
	}

	ts.tinyList = append(ts.tinyList, api.TinyStore{
		Url:           param.Url,
		Shortcode:     shortCode,
		StartDate:     time.Now().Format(time.RFC3339),
		LastSeenDate:  "",
		RedirectCount: 0,
	})

	// print to cli
	ts.Log()

	setStore := make(map[string]interface{}) //set custom response with interface
	setStore["shortcode"] = shortCode

	return api.Response{
		HttpCode:    200,
		Description: "Success",
		Data:        setStore,
	}
}

func (ts *TinyService) Get(shortcode string) api.Response {
	if k, v, err := ts.CheckShortcode(shortcode); err == nil {
		// Update lastSeen and Redirect
		ts.tinyList[k].LastSeenDate = time.Now().Format(time.RFC3339)
		ts.tinyList[k].RedirectCount += 1

		// print to cli
		ts.Log()

		setUrl := make(map[string]interface{}) //set custom response with interface
		setUrl["location"] = v.Url

		return api.Response{
			HttpCode:    302,
			Description: "Success",
			Data:        setUrl,
		}
	}

	return api.Response{
		HttpCode:    404,
		Description: "The shortcode cannot be found in the system",
		Data:        "",
	}
}

func (ts *TinyService) Stats(shortcode string) api.Response {
	if _, v, err := ts.CheckShortcode(shortcode); err == nil {

		setStats := make(map[string]interface{}) //set custom response with interface

		setStats["redirectCount"] = v.RedirectCount

		if v.RedirectCount > 0 {
			setStats["LastSeenDate"] = v.LastSeenDate
		}

		return api.Response{
			HttpCode:    200,
			Description: "Success",
			Data:        setStats,
		}
	}

	return api.Response{
		HttpCode:    404,
		Description: "The shortcode cannot be found in the system",
		Data:        "",
	}
}

func (ts *TinyService) CheckShortcode(shortcode string) (int, api.TinyStore, error) {
	for k, v := range ts.tinyList {
		if v.Shortcode == shortcode {
			return k, v, nil
		}
	}
	return 0, api.TinyStore{}, errors.New("Shortcode Not Found")
}

func (ts *TinyService) Log() {
	// showing every tinyurl list
	jsonEncode, _ := json.Marshal(ts.tinyList)
	fmt.Println(string(jsonEncode))
}

func (ts *TinyService) RandomShort() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 6
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
