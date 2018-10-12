/*************************************
    __    __    __    __    __    __
   /  \  /  \  /  \  /  \  /  \  /  \
  /  __\/  __\/  __\/  __\/  __\/  __\
 /  /__/  /__/  /__/  /__/  /__/  /__/
  \   / \   / \   / \   / \   / \   /
   \_/   \_/   \_/   \ /   \_/   \_/

App : our
Name: Ouroboros

*************************************/
package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/our/rts"
	"github.com/robfig/cron"
)

const (
	ComServer = "http://127.0.0.1:9998"
)

func main() {
	cr := cron.New()
	cr.AddFunc("@every 15s", func() {
		fmt.Println("Radi kron")
		//ser.GetData(ComServer)
	})
	cr.Start()
	jsonHandler := http.FileServer(http.Dir("./JDB/"))
	jsonHandler = http.StripPrefix("/json/", jsonHandler)
	r := mux.NewRouter()
	r.PathPrefix("/json/").Handler(jsonHandler)
	r.Host("com-http.us").Path("/").HandlerFunc(rts.IndexHandler).Name("index")
	r.Host("com-http.us").Path("/coins/").HandlerFunc(rts.CoinsHandler).Name("coins")
	r.Host("com-http.us").Path("/home").HandlerFunc(rts.HomeHandler).Name("home")
	r.Host("{subdomain}.com-http.us").Path("/").HandlerFunc(rts.CoinHandler).Name("coin")
	r.Host("{subdomain}.com-http.us").Path("/explorer").HandlerFunc(rts.ExplorerHandler).Name("explorer")
	r.Host("{subdomain}.com-http.us").Path("/block/{id}").HandlerFunc(rts.ViewBlock).Name("block")
	r.Host("{subdomain}.com-http.us").Path("/tx/{id}").HandlerFunc(rts.ViewTx).Name("tx")
	r.Host("{subdomain}.com-http.us").Path("/addr/{id}").HandlerFunc(rts.ViewAddr).Name("addr")
	//api
	r.Host("{subdomain}.com-http.us").Path("/a/news").HandlerFunc(rts.CoinNewsHandler).Name("news")
	r.Host("{subdomain}.com-http.us").Path("/a/last").HandlerFunc(rts.ApiLast).Name("last")
	r.Host("{subdomain}.com-http.us").Path("/a/info").HandlerFunc(rts.ApiInfo).Name("info")
	r.Host("{subdomain}.com-http.us").Path("/a/mining").HandlerFunc(rts.ApiMiningInfo).Name("mining")
	r.Host("{subdomain}.com-http.us").Path("/a/rawpool").HandlerFunc(rts.ApiRawPool).Name("rawpool")
	r.Host("{subdomain}.com-http.us").Path("/a/search").HandlerFunc(rts.DoSearch).Name("search")
	r.Host("{subdomain}.com-http.us").Path("/a/{type}/{id}").HandlerFunc(rts.ApiData).Name("coin")

	log.Fatal(http.ListenAndServe(":8985", handlers.CORS()(r)))
}
