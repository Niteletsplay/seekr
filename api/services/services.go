package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	//"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"

	// "github.com/seekr-osint/seekr/api/history"
	"github.com/seekr-osint/seekr/api/language"
)

func StatusCodeUserExistsFunc(data UserServiceDataToCheck) (bool, error) {
	url, _ := data.GetUserHtmlUrl()
	log.Printf("status code check:%s\n", url)
	return data.StatusCodeUserExistsFunc()
}

func EmptyInfo(data UserServiceDataToCheck) (AccountInfo, error) { // data can sometimes be nil
	return AccountInfo{}, nil
}

var DefaultServices = Services{
	// {
	// 	Name: "Instagram",
	// 	UserExistsFunc: func(data UserServiceDataToCheck) (bool, error) {
	// 		return data.PatternUrlMatchUserExists("user?username={{.Username}}")
	// 	},
	// 	InfoFunc:            InstagramInfo,
	// 	Domain:              "www.instagram.com",
	// 	UserHtmlUrlTemplate: "{{.Domain}}/{{.Username}}/",
	// 	TestData: TestData{
	// 		ExistingUser:    "greg",
	// 		NotExistingUser: "greg2q1412fdwkdfns",
	// 	},
	// },
	{
		Name:                "GitHub",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		InfoFunc:            GitHubInfo,
		Domain:              "github.com",
		UserHtmlUrlTemplate: "{{.Domain}}/{{.Username}}",
		TestData: TestData{
			ExistingUser:    "greg",
			NotExistingUser: "greg2q1412fdwkdfns",
		},
	},

	{
		Name:                "YouTube",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		InfoFunc:            YouTubeInfo,
		Domain:              "youtube.com",
		UserHtmlUrlTemplate: "{{.Domain}}/@{{.Username}}",
		UrlTemplates: map[string]string{
			"bio": "{{.Domain}}/@{{.Username}}/about",
		},
		TestData: TestData{
			ExistingUser:    "mrbeast",
			NotExistingUser: "gregdoesnotexsist",
		},
	},
	{
		Name:                "TikTok",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		InfoFunc:            TikTokInfo,
		Domain:              "tiktok.com",
		UserHtmlUrlTemplate: "{{.Domain}}/@{{.Username}}",
		TestData: TestData{
			ExistingUser:    "greg",
			NotExistingUser: "gregdoesnotexsist",
		},
	},
	//{
	//	Name:           "TryHackMe",
	//	UserExistsFunc: StatusCodeUserExistsFunc,
	//	Domain: "tryhackme.com",
	//	BlocksTor: true,
	//	UserHtmlUrlTemplate: "{{.Domain}}/p/{{.Username}}",
	//	TestData: TestData{
	//		ExistingUser:    "greg",
	//		NotExistingUser: "gregdoesnotexsist",
	//	},
	//},
	{
		Name:                "Npm",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		InfoFunc:            EmptyInfo,
		Domain:              "npmjs.com",
		UserHtmlUrlTemplate: "{{.Domain}}/~{{.Username}}",
		TestData: TestData{
			ExistingUser:    "greg",
			NotExistingUser: "gregdoesnotexsist",
		},
	},
	{
		Name:                "chess.com",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		InfoFunc:            ChessComInfo,
		Domain:              "api.chess.com",
		UserHtmlUrlTemplate: "{{.Domain}}/pub/player/{{.Username}}",
		BlocksTor:           true,
		TestData: TestData{
			ExistingUser:    "danielnaroditsky",
			NotExistingUser: "gregdoesnotexsist",
		},
	},
	{
		Name:                "Asciinema",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		InfoFunc:            AsciinemaInfo,
		Domain:              "asciinema.org",
		UserHtmlUrlTemplate: "{{.Domain}}/~{{.Username}}",
		TestData: TestData{
			ExistingUser:    "greg",
			NotExistingUser: "gregdoesnotexsist",
		},
	},
	// blocks tor
	//{
	//	Name:           "Replit",
	//	UserExistsFunc: StatusCodeUserExistsFunc,
	//	Domain: "replit.com",
	//	UserHtmlUrlTemplate: "{{.Domain}}/{{.Username}}",
	//	TestData: TestData{
	//		ExistingUser:    "greg",
	//		NotExistingUser: "gregdoesnotexsistsfdssfda",
	//	},
	//},

	{
		Name:                "Lichess",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		Domain:              "lichess.org",
		UserHtmlUrlTemplate: "{{.Domain}}/api/user/{{.Username}}",
		BlocksTor:           true, // ???
		TestData: TestData{
			ExistingUser:    "starwars",
			NotExistingUser: "gregdoesnotexsist",
		},
	},

	{
		Name:                "Snapchat",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		InfoFunc:            SnapchatInfo,
		Domain:              "snapchat.com",
		UserHtmlUrlTemplate: "{{.Domain}}/add/{{.Username}}",
		BlocksTor:           true,
		TestData: TestData{
			ExistingUser:    "greg",
			NotExistingUser: "gregdoesnotexsistdsada",
		},
	},
}

func ServicesCheckWorker(s <-chan UserServiceDataToCheck, res chan<- ServiceCheckResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for service := range s {
		status := service.UserExistsFunction()
		status.GetInfo(service)
		res <- status
	}
}

func ScrapeImageTwitterTag(response http.Response) (string, error) {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}

	img := doc.Find(`meta[name="twitter:image"]`).AttrOr("content", "")
	log.Printf("image: %s", img)
	return img, nil
}

func ScrapeBioTwitterTag(response http.Response) (string, error) {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}

	bio := doc.Find(`meta[name="twitter:description"]`).AttrOr("content", "")
	return bio, nil
}

func SnapchatInfo(data UserServiceDataToCheck) (AccountInfo, error) {
	url, err := data.GetUserHtmlUrl()
	if err != nil {
		return AccountInfo{}, err
	}
	response, err := http.Get(url)
	if err != nil {
		return AccountInfo{}, err
	}
	defer response.Body.Close()

	bio, err := ScrapeBioTwitterTag(*response)
	if err != nil {
		return AccountInfo{}, err
	}

	img, err := ScrapeImageTwitterTag(*response)
	if err != nil {
		return AccountInfo{}, err
	}
	_, err = GetImage(RemoveExtension(img, "jpeg"))
	if err != nil { // FIXME no pfp
		return AccountInfo{}, err
	}
	accountInfo := AccountInfo{}
	accountInfo.Bio.AddOrUpdateLatestItem(NewBio(bio))

	// accountInfo.ProfilePicture.AddOrUpdateLatestItem(pfp) // cors
	return accountInfo, nil
}
func YouTubeInfo(data UserServiceDataToCheck) (AccountInfo, error) {
	url, err := data.GetUserHtmlUrl()
	if err != nil {
		return AccountInfo{}, err
	}
	response, err := http.Get(url + "/about")
	if err != nil {
		return AccountInfo{}, err
	}
	defer response.Body.Close()

	bio, err := ScrapeBioTwitterTag(*response)
	if err != nil {
		return AccountInfo{}, err
	}

	img, err := ScrapeImageTwitterTag(*response)
	if err != nil {
		return AccountInfo{}, err
	}
	_, err = GetImage(RemoveExtension(img, "jpg"))
	if err != nil { // FIXME no pfp
		return AccountInfo{}, err
	}
	accountInfo := AccountInfo{}
	accountInfo.Bio.AddOrUpdateLatestItem(NewBio(bio))
	// accountInfo.ProfilePicture.AddOrUpdateLatestItem(pfp) // FIXME cors
	return accountInfo, nil
}

func GitHubInfo(data UserServiceDataToCheck) (AccountInfo, error) {
	url, err := data.GetUserHtmlUrl()
	if err != nil {
		return AccountInfo{}, err
	}
	response, err := http.Get(url)
	if err != nil {
		return AccountInfo{}, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return AccountInfo{}, err
	}
	imgElement := doc.Find("img.avatar.avatar-user")
	imgUrl, _ := imgElement.Attr("src")
	bioText := doc.Find(".p-note[data-bio-text]").Text()
	//ogImageContent, _ = element.Attr("content")
	pfp, err := GetImage(imgUrl)
	if err != nil { // FIXME no pfp
		return AccountInfo{}, err
	}
	accountInfo := AccountInfo{}
	accountInfo.ProfilePicture.AddOrUpdateLatestItem(pfp)
	accountInfo.Bio.AddOrUpdateLatestItem(NewBio(bioText))
	return accountInfo, nil
}
func InstagramInfo(data UserServiceDataToCheck) (AccountInfo, error) {
	url, err := data.GetUserHtmlUrl()
	if err != nil {
		return AccountInfo{}, err
	}
	response, err := http.Get(url)
	if err != nil {
		return AccountInfo{}, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return AccountInfo{}, err
	}
	var imgUrl string
	doc.Find("meta[property='og:image']").Each(func(_ int, item *goquery.Selection) {
		content, exists := item.Attr("content")
		if exists {
			imgUrl = content
		}
	})
	//ogImageContent, _ = element.Attr("content")
	pfp, err := GetImage(imgUrl)
	if err != nil { // FIXME no pfp
		return AccountInfo{}, err
	}

	accountInfo := AccountInfo{}
	accountInfo.ProfilePicture.AddOrUpdateLatestItem(pfp)
	return accountInfo, nil
}

func AsciinemaInfo(data UserServiceDataToCheck) (AccountInfo, error) {
	// NOTE joined date is scrapable
	url, err := data.GetUserHtmlUrl()
	if err != nil {
		return AccountInfo{}, err
	}
	response, err := http.Get(url)
	if err != nil {
		return AccountInfo{}, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return AccountInfo{}, err
	}
	var imgUrl string
	doc.Find("img.avatar").Each(func(i int, img *goquery.Selection) {
		imgUrl, _ = img.Attr("src")
	})

	pfp, err := GetImage("https:" + imgUrl)
	if err != nil { // FIXME no pfp
		return AccountInfo{}, err
	}
	accountInfo := AccountInfo{}
	accountInfo.ProfilePicture.AddOrUpdateLatestItem(pfp)
	return accountInfo, nil
}

type Player struct {
	Avatar     string `json:"avatar"`
	PlayerID   int    `json:"player_id"`
	ID         string `json:"@id"`
	Url        string `json:"url"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Followers  int    `json:"followers"`
	Country    string `json:"country"`
	Location   string `json:"location"`
	LastOnline int64  `json:"last_online"`
	Joined     int64  `json:"joined"`
	Status     string `json:"status"`
	IsStreamer bool   `json:"is_streamer"`
	Verified   bool   `json:"verified"`
}

func ChessComInfo(data UserServiceDataToCheck) (AccountInfo, error) {
	url, err := data.GetUserHtmlUrl()
	if err != nil {
		return AccountInfo{}, err
	}
	response, err := http.Get(url)
	if err != nil {
		return AccountInfo{}, err
	}
	defer response.Body.Close()
	var player Player
	jsonData, err := io.ReadAll(response.Body)
	if err != nil {
		return AccountInfo{}, err
	}
	err = json.Unmarshal([]byte(jsonData), &player)
	if err != nil {
		return AccountInfo{}, err
	}

	img, err := GetImage(player.Avatar)
	if err != nil {
		return AccountInfo{}, err
	}

	accountInfo := AccountInfo{}
	accountInfo.ProfilePicture.AddOrUpdateLatestItem(img)
	accountInfo.Url = player.Url
	return accountInfo, nil
}

func TikTokInfo(data UserServiceDataToCheck) (AccountInfo, error) {
	url, err := data.GetUserHtmlUrl()
	if err != nil {
		return AccountInfo{}, err
	}
	response, err := http.Get(url)
	if err != nil {
		return AccountInfo{}, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return AccountInfo{}, err
	}

	selector := "h2[data-e2e='user-bio']"
	userBioElement := doc.Find(selector)

	userBioText := userBioElement.Text()
	if userBioText == "No bio yet." {
		userBioText = ""
	}
	var imgUrl string
	doc.Find("meta[data-rh=true][property='og:image']").Each(func(index int, element *goquery.Selection) {
		imgUrl, _ = element.Attr("content")
	})
	pfp, err := GetImage(imgUrl)
	if err != nil { // FIXME no pfp
		return AccountInfo{}, err
	}
	accountInfo := AccountInfo{}
	accountInfo.ProfilePicture.AddOrUpdateLatestItem(pfp)
	accountInfo.Bio.AddOrUpdateLatestItem(NewBio(userBioText))
	return accountInfo, nil
}
func NewBio(bio string) Bio { // TODO discord tag regex/username regex (https://www.tiktok.com/@japan)
	if bio != "" {
		return Bio{
			Language: language.DetectLanguage(bio),
			Bio:      bio,
		}
	}
	return Bio{
		Bio: bio,
	}
}
func RemoveExtension(input, extension string) string {
	lastIndex := strings.LastIndex(input, extension)
	if lastIndex == -1 {
		return input
	}
	return input[:lastIndex] + extension
}
