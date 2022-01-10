package cpm

type CpmProjectLanguage struct {
	LanguageId   string `json:"languageId" gorm:"comment:语言ID"`
	LanguageName string `json:"languageName" gorm:"comment:语言名称"`
}
