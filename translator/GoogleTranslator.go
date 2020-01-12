package translator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Translator interface {
	Translate(text string, src string, dst string)
}

type GoogleTranslator struct {
	url    string
	client *http.Client
}

func New() *GoogleTranslator {
	return &GoogleTranslator{url: "https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s", client: &http.Client{}}
}

/**
 * translating
 */
func (translator *GoogleTranslator) Translate(text string, src string, dst string) (string, error) {
	fmt.Println(fmt.Sprintf("Trying to translate from src = %s to dst = %s given text: %s", src, dst, text))
	req, reqErr := http.NewRequest("GET", fmt.Sprintf(translator.url, "ja", "ko", url.PathEscape(text)), nil)
	if reqErr != nil {
		return "", reqErr
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	resp, respErr := translator.client.Do(req)
	if respErr != nil {
		return "", respErr
	}
	data, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return "", readErr
	}
	var parsedData []interface{}
	parseErr := json.Unmarshal(data, &parsedData)
	if parseErr != nil {
		return "", parseErr
	}
	return parseData(parsedData)
}

func parseData(data []interface{}) (string, error) {
	var translatedData strings.Builder
	for _, value := range data[0].([]interface{}) {
		translatedData.WriteString(value.([]interface{})[0].(string))
	}
	return translatedData.String(), nil
}


// use it as t, _ := translator.Translate("こんにちは", "src", "dst")
