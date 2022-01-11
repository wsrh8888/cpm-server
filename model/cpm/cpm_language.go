package cpm

type CpmLanguage struct {
	LanguageId   string `json:"languageId" gorm:"primary_key;comment:语言ID"`
	LanguageName string `json:"languageName" gorm:"comment:语言名称"`
}
