package ext

type Mod struct {
	Host struct {
		Hid        int    `yaml:"Hid"`
		Hostname   string `yaml:"Hostname"`
		Name       string `yaml:"Name"`
		Desc       string `yaml:"Desc"`
		Createdate string `yaml:"Createdate"`
		Usetoken   string `yaml:"Usetoken"`
		Token      string `yaml:"Token"`
	} `yaml:"Host"`
	User struct {
		UID        int    `yaml:"Uid"`
		Username   string `yaml:"Username"`
		Userid     int    `yaml:"Userid"`
		Createdate string `yaml:"Createdate"`
		Sshpk      string `yaml:"sshpk"`
		Desc       string `yaml:"Desc"`
	} `yaml:"User"`
}
