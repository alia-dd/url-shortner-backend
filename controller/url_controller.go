package controller

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"url-shortener/schema"

	"gorm.io/gorm"
)

var SchemaM = schema.Url_Data{}

func PostUrl(longUrl string) (string, error) {
	exist, existing, err := checkifexist("Long_url", longUrl)
	if err != nil {
		fmt.Errorf("somthing went wrong: %s", err)
	}
	if exist {
		fmt.Println("The Url is already saved.")
		return existing.Short_url, nil
	}
	data := schema.Url_Data{
		Long_url: longUrl,
	}
	if res := schema.DBConnect.DB.Create(&data); res.Error != nil {
		return "", res.Error
	}
	short, err := updateUrl(&data)
	if err != nil {
		return short, err
	}
	return short, nil
}

func updateUrl(data *schema.Url_Data) (string, error) {
	short := encodeID(data.ID)
	return short, schema.DBConnect.DB.Model(data).Update("Short_url", short).Error
}

func GetUrls() ([]schema.Url_Data, error) {
	var urls []schema.Url_Data
	err := schema.DBConnect.DB.Find(&urls).Error
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func ListUrls() {

	urls, err := GetUrls()
	if err != nil {
		fmt.Errorf("Could not locate the data:", err)
		return
	}
	fmt.Println("Short url	| Long url	")
	fmt.Println("----------------|-------------------")
	for _, u := range urls {
		fmt.Printf("%s	| %s\n", u.Short_url, u.Long_url)
	}
}

func GetUrlLongUrl(shortUrl string) (string, error) {
	var longurl schema.Url_Data
	err := schema.DBConnect.DB.Where(
		"Short_url = ? ", shortUrl).First(&longurl).Error
	if err != nil {
		return "", err
	}
	return longurl.Long_url, nil
}

func checkifexist(fieldName string, url string) (bool, schema.Url_Data, error) {
	var data schema.Url_Data
	var err error
	switch fieldName {
	case "Long_url":
		err = schema.DBConnect.DB.Where("Long_url = ? ", url).First(&data).Error
	case "Short_url":
		err = schema.DBConnect.DB.Where("Short_url = ? ", url).First(&data).Error
	default:
		return false, data, fmt.Errorf("invalid field name: %s", fieldName)
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, data, nil
		}
		return false, data, err
	}
	return true, data, nil
}
func encodeID(id uint) string {
	s := strconv.Itoa(int(id))
	return base64.RawURLEncoding.EncodeToString([]byte(s))
}
