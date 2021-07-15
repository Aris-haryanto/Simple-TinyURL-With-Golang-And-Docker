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

func (ts *TinyService) Store(param api.TinyStore) api.ResStore {
	if param.Url == "" {
		return api.ResStore{
			HttpCode:    400,
			Description: "url is not present",
		}
	}

	if _, _, err := ts.CheckShortcode(param.Shortcode); err == nil {
		return api.ResStore{
			HttpCode:    409,
			Description: "The the desired shortcode is already in use. Shortcodes are case-sensitive.",
		}
	}

	checkShort := regexp.MustCompile("^[0-9a-zA-Z_]{6}$")
	if checkShort.MatchString(param.Shortcode) {
		return api.ResStore{
			HttpCode:    422,
			Description: "The shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{6}$.",
		}
	}

	shortCode := param.Shortcode
	if param.Shortcode == "" {
		shortCode = ts.RandomShort()
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

	return api.ResStore{
		HttpCode:  200,
		Shortcode: shortCode,
	}
}

func (ts *TinyService) Get(shortcode string) api.ResGet {
	if k, v, err := ts.CheckShortcode(shortcode); err == nil {
		// Update lastSeen and Redirect
		ts.tinyList[k].LastSeenDate = time.Now().Format(time.RFC3339)
		ts.tinyList[k].RedirectCount += 1

		// print to cli
		ts.Log()

		return api.ResGet{
			HttpCode: 200,
			Location: v.Url,
		}
	}

	return api.ResGet{
		HttpCode:    404,
		Description: "The shortcode cannot be found in the system",
	}
}

func (ts *TinyService) Stats(shortcode string) api.ResStats {
	if _, v, err := ts.CheckShortcode(shortcode); err == nil {

		var setStats api.ResStats
		setStats.HttpCode = 200
		setStats.StartDate = v.StartDate
		setStats.LastSeenDate = v.LastSeenDate

		if v.RedirectCount > 0 {
			setStats.RedirectCount = v.RedirectCount
		}

		return setStats
	}

	return api.ResStats{
		HttpCode:    404,
		Description: "The shortcode cannot be found in the system",
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
