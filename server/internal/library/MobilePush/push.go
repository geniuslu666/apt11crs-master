package MobilePush

import (
	"APT/utility/encrypt"
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

type JSONDataParams struct {
	Source           string      `json:"source"`
	Appkey           string      `json:"appkey"`
	PushTarget       PushTarget  `json:"pushTarget"`
	PushNotify       PushNotify  `json:"pushNotify"`
	PushFactoryExtra g.MapStrAny `json:"pushFactoryExtra"`
}

type HarmonyNotify struct {
	Title string `json:"title"`
}

type PushTarget struct {
	Target int      `json:"target"`
	Alias  []string `json:"alias"`
}
type IosNotify struct {
	Badge     int `json:"badge"`
	BadgeType int `json:"badgeType"`
}
type PushNotify struct {
	Plats         []int               `json:"plats"`
	IosNotify     IosNotify           `json:"iosNotify"`
	IosProduction int                 `json:"iosProduction"`
	Content       string              `json:"content"`
	Type          int                 `json:"type"`
	ExtrasMapList []map[string]string `json:"extrasMapList"`
}

type ResponseJSONData struct {
	Status int `json:"status"`
	Res    struct {
		BatchId    string      `json:"batchId"`
		Fetched    interface{} `json:"fetched"`
		Uninstalls interface{} `json:"uninstalls"`
		Closes     interface{} `json:"closes"`
		NotFounds  interface{} `json:"notFounds"`
	} `json:"res"`
	Error interface{} `json:"error"`
}

var (
	AppKey    = "39b08e4ca9daf"
	AppSecret = "32744a51b42046dbe3fdaebca06f14ce"
	Url       = "http://api.push.mob.com/v3/push/createPush"
)

func SendPush(ctx context.Context, Title string, Content string, PushIds []string, Extras map[string]string) (err error) {
	var (
		client       = g.Client()
		Params       JSONDataParams
		ParamsJson   []byte
		Sign         string
		Response     *gclient.Response
		ResponseBody *ResponseJSONData
	)

	Params = JSONDataParams{
		Source: "webapi",
		Appkey: AppKey,
		PushTarget: PushTarget{
			Target: 2,
			Alias:  PushIds,
		},
		PushNotify: PushNotify{
			Plats: []int{1, 2, 8},
			IosNotify: IosNotify{
				Badge:     1,
				BadgeType: 2,
			},
			IosProduction: 1,
			Content:       Content,
			Type:          1,
		},
		PushFactoryExtra: g.MapStrAny{
			"harmonyExtra": g.MapStrAny{
				"category": "CUSTOMER_SERVICE",
			},
		},
	}
	if ok := Extras["path"]; ok != "" {
		Params.PushNotify.ExtrasMapList = append(Params.PushNotify.ExtrasMapList, g.MapStrStr{
			"key":   "path",
			"value": Extras["path"],
		})
	}
	if ok := Extras["data"]; ok != "" {
		Params.PushNotify.ExtrasMapList = append(Params.PushNotify.ExtrasMapList, g.MapStrStr{
			"key":   "data",
			"value": Extras["data"],
		})
	}
	if ok := Extras["type"]; ok != "" {
		Params.PushNotify.ExtrasMapList = append(Params.PushNotify.ExtrasMapList, g.MapStrStr{
			"key":   "type",
			"value": Extras["type"],
		})
	}
	if ok := Extras["param"]; ok != "" {
		Params.PushNotify.ExtrasMapList = append(Params.PushNotify.ExtrasMapList, g.MapStrStr{
			"key":   "param",
			"value": Extras["param"],
		})
	}
	if ok := Extras["string"]; ok != "" {
		Params.PushNotify.ExtrasMapList = append(Params.PushNotify.ExtrasMapList, g.MapStrStr{
			"key":   "string",
			"value": Extras["string"],
		})
	}

	if ParamsJson, err = json.Marshal(Params); err != nil {
		return
	}
	Sign = encrypt.Md5ToString(string(ParamsJson) + AppSecret)

	client.SetHeader("key", AppKey)
	client.SetHeader("sign", Sign)
	if Response, err = client.ContentJson().Post(ctx, Url, Params); err != nil {
		return
	}
	defer Response.Close()
	g.Log().Debug(ctx, Response.Raw())
	if err = json.Unmarshal(Response.ReadAll(), &ResponseBody); err != nil {
		return
	}
	if ResponseBody.Status != 200 {
		err = gerror.New("推送失败")
		return
	}
	return
}
