package models

type LinkShortner struct{
	UrlKey string `json:"url_key" gorm:"column:url_key"`
	OriginalLink string `json:"original_link"`
}


type DatabaseCred struct{
   Name string 
   User string
   Password string
}