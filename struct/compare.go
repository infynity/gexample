package main

type A2 struct {
	name string
	bs struct{
		sd *int
	}
}

type A struct {
	Iptables struct {
		Content []string `json:"content"`  //切片不能比较
	} `json:"iptables"`
	Screensaver struct {
		Enable bool `json:"enable"`
		Delay  int  `json:"delay"`
	} `json:"screensaver"`
	Power struct {
		DisplayStandby int `json:"display_standby"`
	} `json:"power"`
	PamCracklib struct {
		Minlen         int  `json:"minlen"`
		RejectUsername bool `json:"reject_username"`
		Minclass       int  `json:"minclass"`
	} `json:"pam_cracklib"`
	Wallpaper struct {
		Image string `json:"image"`
	} `json:"wallpaper"`
	Script struct {
		Name     string       `json:"name"`
		Content  string       `json:"content"`
		//UploadAt utils.ExTime `json:"upload_at" sql:"-"`
	} `json:"script"`
	CanLogin bool `json:"can_login"`
}

func main(){

	//a := getStrct()
	a := getstruct()

	var b A
	if a==b{

	}

}

func getStrct() *A{
	return &A{}
}

func getstruct() A{
	return A{}
}
