package translate

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

type LanguageObj struct {
	Data struct {
		Translations []struct {
			TranslatedText         string `json:"translatedText"`
			DetectedSourceLanguage string `json:"detectedSourceLanguage"`
		} `json:"translations"`
	} `json:"data"`
}

type DetectObj struct {
	Data struct {
		Detections [][]struct {
			Language   string      `json:"language"`
			Confidence json.Number `json:"confidence"`
			IsReliable bool        `json:"isReliable"`
		} `json:"detections"`
	} `json:"data"`
}

var (
	baseURL = "https://translation.googleapis.com/language/translate/v2"
	apiKey  = "AIzaSyATmTewyulgKtfNbKJFF5baIaFimyCX5Bs"
)

func TranslateText(ctx context.Context, text string, targetLang string) (languageText string, err error) {

	var (
		client      = g.Client()
		response    *gclient.Response
		languageObj LanguageObj
	)
	//client = client.Proxy("http://127.0.0.1:7890")
	if response, err = client.PostForm(ctx, baseURL, g.MapStrStr{
		"q":      text,
		"target": targetLang,
		"key":    apiKey,
	}); err != nil {
		return
	}
	defer response.Close()
	g.Log().Debug(ctx, response.Raw())

	if err = json.Unmarshal(response.ReadAll(), &languageObj); err != nil {
		return
	}
	languageText = languageObj.Data.Translations[0].TranslatedText
	return
}

func DetectLanguage(ctx context.Context, text string) (Detection *DetectObj, err error) {
	var (
		client   = g.Client()
		response *gclient.Response
	)
	//client = client.Proxy("http://127.0.0.1:7890")
	if response, err = client.Get(ctx, baseURL+"/detect", g.MapStrStr{
		"q":   text,
		"key": apiKey,
	}); err != nil {
		return
	}
	defer response.Close()
	g.Log().Debug(ctx, response.Raw())

	if err = json.Unmarshal(response.ReadAll(), &Detection); err != nil {
		return
	}
	return
}
