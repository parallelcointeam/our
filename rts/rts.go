package rts

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parallelcointeam/our/conf"
	"github.com/parallelcointeam/our/jdb"
	"github.com/parallelcointeam/our/mod"
	"github.com/parallelcointeam/our/tools"
)

var cf = conf.CsYsConf()
var ComServer = cf.ComServer
var templates = make(map[string]*template.Template)

var last string

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["404"] = template.Must(template.ParseFiles("tpl/hlp/plgs.gohtml", "tpl/css/boot.gohtml", "tpl/css/grid.gohtml", "tpl/css/typo.gohtml", "tpl/css/btn.gohtml", "tpl/hlp/base.gohtml", "tpl/hlp/body.gohtml", "tpl/hlp/head.gohtml", "tpl/hlp/style.gohtml", "tpl/hlp/spectre.gohtml", "tpl/404.gohtml", "tpl/hlp/search.gohtml"))

	// templates["cmc"] = template.Must(template.ParseFiles("tpl/frame/cmc.gohtml"))

	templates["ccw"] = template.Must(template.ParseFiles("tpl/frames/prices/ccw.gohtml"))
	templates["trades"] = template.Must(template.ParseFiles("tpl/frames/prices/trades.gohtml"))
	templates["nodes"] = template.Must(template.ParseFiles("tpl/frames/maps/nodes.gohtml"))
	templates["magnet"] = template.Must(template.ParseFiles("tpl/frames/games/magnet.gohtml"))
	templates["spiralogo"] = template.Must(template.ParseFiles("tpl/frames/anim/spiralogo.gohtml"))

	templates["proxy"] = template.Must(template.ParseFiles("tpl/proxy/proxy.gohtml"))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	data := tools.GetData("index", "http://127.0.0.1:3553/")
	renderTemplate(w, "proxy", "proxy", data)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := "COINS"
	renderTemplate(w, "home", "base", data)
}

func FrameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	frame := vars["frame"]
	gdb, err := jdb.OpenDB()
	if err != nil {
	}
	vCoin := mod.VCoin{}
	if err := gdb.Read("coins", coin, &vCoin); err != nil {
		fmt.Println("Error", err)
	}
	nodesUrl := ComServer + "a/n/" + coin
	gamp, err := http.Get(nodesUrl)
	if err != nil {
		fmt.Println("AMP gampgampgampgamp", gamp)
	}

	fmt.Println("Read error", err)
	defer gamp.Body.Close()
	mapNodes, err := ioutil.ReadAll(gamp.Body)
	gnodes := make(map[string]interface{})
	json.Unmarshal(mapNodes, &gnodes)
	if err != nil {
		fmt.Println("Read error", err)
	}

	vCoin.Nodes = gnodes["nodes"]
	data := vCoin
	renderTemplate(w, frame, frame, data)
}
