package models

type Call struct {
	Number               string `json:"number"`
	Type                 string `json:"type"`
	Language             string `json:"language"`
	NotificationCallback string `json:"notification_callback"`
	Platform             string `json:"platform"`
	AndroidAppHash       string `json:"android_app_hash"`
}

type CallResponse struct {
	ID             string         `json:"id"`
	Type           string         `json:"type"`
	PinHash        string         `json:"pin_hash"`
	CliPrefix      string         `json:"cli_prefix"`
	ValidationInfo ValidationInfo `json:"validation_info"`
}

type ValidationInfo struct {
	CountryCode    int    `json:"country_code"`
	CountryIsoCode string `json:"country_iso_code"`
	Carrier        string `json:"carrier"`
	IsMobile       bool   `json:"is_mobile"`
	El64Format     string `json:"e164_format"`
	Formatting     string `json:"formatting"`
}

type Verify struct {
	ID  string `json:"id"`
	Pin string `json:"pin"`
}

type VerifyResponse struct {
	Number         string  `json:"number"`
	Validated      bool    `json:"validated"`
	ValidationDate any     `json:"validation_date"`
	ChargedAmount  float64 `json:"charged_amount"`
}

type ErrorVerifyLimit struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
