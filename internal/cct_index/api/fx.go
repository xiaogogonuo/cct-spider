package api

// 汇通财经-全球指数接口
const (
	// 美元人民币、港元人民币、欧元人民币
	// • 页面展示接口：https://quote.fx678.com/exchange/WHMP
	// • 数据获取接口：https://quote.fx678.com/exchange/WHMP

	// 欧元美元、美元日元、英镑美元、美元指数
	// • 页面展示接口：https://quote.fx678.com/exchange/WH
	// • 数据获取接口：https://quote.fx678.com/exchange/WH

	// 布伦特原油连续
	// • 页面展示接口：https://quote.fx678.com/exchange/MAINOIL
	// • 数据获取接口：https://quote.fx678.com/exchange/MAINOIL

	// 纽约黄金连续、纽约白银连续
	// • 页面展示接口：https://quote.fx678.com/exchange/MAINGOLD
	// • 数据获取接口：https://quote.fx678.com/exchange/MAINGOLD

	// LME铜、LME镍、LME铝
	// • 页面展示接口：https://quote.fx678.com/exchange/LME
	// • 数据获取接口：https://quote.fx678.com/exchange/LME

	// 美玉米连续、美黄豆连续
	// • 页面展示接口：https://quote.fx678.com/exchange/CBOT
	// • 数据获取接口：https://quote.fx678.com/exchange/CBOT

	// 美债10年收益率、日债10年收益率、德债10年收益率、英债10年收益率
	// • 页面展示接口：https://quote.fx678.com/exchange/GJZQ
	// • 数据获取接口：https://quote.fx678.com/exchange/GJZQ

	// 日经225、英国FTSE100、德国DAX30、法国CAC40、意大利MIB、加拿大SP/TSX、纳斯达克指数、道琼斯工业指数、标普500、恒生指数
	// • 页面展示接口：https://quote.fx678.com/exchange/GJZS
	// • 数据获取接口：https://quote.fx678.com/exchange/GJZS

	// FxExchange 汇通财经全球指数统一数据获取接口
	FxExchange = "https://quote.fx678.com/exchange/#"

	// 斯托克600
	// • 页面展示接口：https://quote.fx678.com/symbol/SXO
	// • 数据获取接口：https://api-q.fx678img.com/getQuote.php?exchName=GJZS&symbol=SXO

	// FxExchangeSpecial 汇通财经一些特殊的全球指数统一数据获取接口
	FxExchangeSpecial      = "https://api-q.fx678img.com/getQuote.php?exchName=#&symbol=$"
	FxExchangeSpecialRefer = "https://quote.fx678.com/"
)
