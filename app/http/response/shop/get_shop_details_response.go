package shop_response

type GetShopDetailsResponse struct {
	ShopID      int        `json:"shopId"`
	Name        string     `json:"name"`
	Location    string     `json:"location"`
	ImageUrl    string     `json:"imageUrl"`
	Worktimes   []Worktime `json:"worktimes"`
	Coffees     []Coffee   `json:"coffees"`
	Phones      []Phone    `json:"phones"`
	Socials     []Social   `json:"socials"`
}

type Worktime struct {
	WorktimeID  int    `json:"worktimeId"`
	DayRange    string `json:"dayRange"`
	OpeningTime string `json:"openingTime"`
	ClosingTime string `json:"closingTime"`
}

type Phone struct {
	PhoneID     int    `json:"phoneId"`
	PhoneNumber string `json:"phoneNumber"`
}

type Social struct {
	SocialID int    `json:"socialId"`
	Platform string `json:"platform"`
	Url      string `json:"url"`
}