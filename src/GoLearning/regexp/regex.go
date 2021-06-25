package main

import (
	"fmt"
	re "regexp"
	"strings"
)

type Cfg struct {
	fqdn     string
	url      string
	tenantId string
	extn     string
}

type Req struct {
	provider string
	asset    string
	subtype  string
	brate    string
}

func main() {
	cf := Cfg{
		fqdn:     "abcxyz.com",
		url:      "http://{cfg.fqdn}/vod/{req.provider}/{req.asset}/{cfg.tenant}_{req.provider}_{req.asset}_{cfg.tenant}_hls_{req.bitrate}.{cfg.extn}",
		tenantId: "vs",
		extn:     "m3u8",
	}

	c2re := Req{
		provider: "testproviderblah",
		asset:    "testassetblah",
		subtype:  "index",
		brate:    "3750000",
	}

	c := map[string]interface{}{
		"cfg": cf,
		"req": c2re,
	}

	r := re.MustCompile(`({\w+\.\w+})`)

	s := r.ReplaceAllStringFunc(cf.url, func(s string) string {
		s = strings.ToLower(s)
		ty := strings.Split(s, `.`)

		f := string(ty[1][:len(ty[1])-1])
		switch kField := c[ty[0][1:]].(type) {
		case Cfg:
			switch f {
			case "fqdn":
				s = kField.fqdn
			case "tenant":
				s = kField.tenantId
			case "extn":
				s = kField.extn
			default:
				s = ""
				fmt.Errorf("Invalid Config option for URL pattern")
			}
		case Req:
			switch f {
			case "provider":
				s = kField.provider
			case "asset":
				s = kField.asset
			case "bitrate":
				s = kField.brate
			default:
				s = ""
				fmt.Errorf("Invalid request option for URL pattern")
			}
		}

		return s
	})

	fmt.Println(s)
}
