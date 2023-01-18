package api

// 新浪财经-汇率接口
const (
	// 人民币汇率
	// • 页面展示接口：https://finance.sina.com.cn/money/forex/hq/CNYUSD.shtml
	// • 数据获取接口：https://hq.sinajs.cn/rn=1674015528333list=fx_scnyusd
	//   - 数据获取接口携带的请求头
	//     + Referer: https://finance.sina.com.cn/money/forex/hq/CNYUSD.shtml

	// 人民币港元
	// • 页面展示接口：https://finance.sina.com.cn/money/forex/hq/CNYHKD.shtml
	// • 数据获取接口：https://hq.sinajs.cn/rn=1674015528333list=fx_scnyhkd
	//   - 数据获取接口携带的请求头
	//     + Referer: https://finance.sina.com.cn/money/forex/hq/CNYHKD.shtml

	// 人民币日元
	// • 页面展示接口：https://finance.sina.com.cn/money/forex/hq/CNYJPY.shtml
	// • 数据获取接口：https://hq.sinajs.cn/rn=1674015528333list=fx_scnyjpy
	//   - 数据获取接口携带的请求头
	//     + Referer: https://finance.sina.com.cn/money/forex/hq/CNYJPY.shtml

	// Forex 汇率统一数据获取接口
	// $：当前时间的unix时间戳: 1674015528333
	// #：A货币兑B货币: cnyusd、cnyhkd、cnyjpy
	Forex        = "https://hq.sinajs.cn/rn=$list=fx_s#"
	ForexReferer = "https://finance.sina.com.cn/money/forex/hq/#.shtml"
)

// 新浪财经-宏观经济接口
const (
	// 地区生产总值
	// • 页面展示接口：http://finance.sina.com.cn/mac/#nation-7-0-31-3
	// • 数据获取接口：https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata?cate=nation&event=7&from=0&num=31&condition=

	// 地区居民消费价格指数
	// • 页面展示接口：http://finance.sina.com.cn/mac/#price-2-0-31-1
	// • 数据获取接口：

	// Marco 宏观经济统一数据获取接口
	// #：指标编码: cate=nation&event=7、cate=price&event=2
	// $：页数
	Marco = "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata?#&from=$&num=31&condition="
)
