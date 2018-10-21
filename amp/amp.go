package amp

import "html/template"

type AMPc struct {
	Img  template.HTML `json:"ampimg"`
	News template.HTML `json:"ampnews"`
}

func AMPC() AMPc {
	amp := AMPc{
		//Img:  template.HTML(`<a class="cglc" href="//{{slug}}.com-http.us"><amp-img width="32" height="32" alt="{{name}}" src="{{img}}"></amp-img><noscript><img src="{{img}}" width="32" height="32" alt="{{name}}"></noscript></a>`),
		Img:  template.HTML(`<a class="cglc" href="//{{slug}}.com-http.us"><amp-img width="32" height="32" alt="{{name}}" src="{{img}}"></amp-img></a>`),
		News: template.HTML(`<a href="{{link}}" target="_blank">{{title}}</a>`),
	}
	return amp
}

type AMPb struct {
	BlockHeight   template.HTML `json:"blockheight"`
	BlockHash     template.HTML `json:"blockhash"`
	Tx            template.HTML `json:"tx"`
	Addr          template.HTML `json:"addr"`
	NextBlockHash template.HTML `json:"nbh"`
	PrevBlockHash template.HTML `json:"pbh"`
}

func AMPB() AMPb {
	amp := AMPb{
		BlockHeight:   template.HTML(`<a href="/block/{{height}}">{{height}}</a>`),
		BlockHash:     template.HTML(`<a href="/blockhash/{{hash}}">{{hash}}</a>`),
		Tx:            template.HTML(`<a href="/tx/{{.}}">{{.}}</a>`),
		Addr:          template.HTML(`<a href="/addr/{{.}}">{{.}}</a>`),
		NextBlockHash: template.HTML(`<a href="/blockhash/{{nextblockhash}}">{{nextblockhash}}</a>`),
		PrevBlockHash: template.HTML(`<a href="/blockhash/{{previousblockhash}}">{{previousblockhash}}</a>`),
	}
	return amp
}
