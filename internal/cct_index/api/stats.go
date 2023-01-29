package api

// 国家统计局-统计数据接口
const (
	// 规模以上工业增加值当月同比增速
	// • 页面展示接口：http://www.stats.gov.cn/was5/web/search?page=1&channelid=288041&orderby=-DOCRELTIME&was_custom_expr=DOCTITLE%3D%28like%28规模以上工业增加值%29%2Fsen%29&perpage=100

	// 国有及国有控股企业工业增加值同比增速
	// • 页面展示接口：http://www.stats.gov.cn/was5/web/search?page=1&channelid=288041&orderby=-DOCRELTIME&was_custom_expr=DOCTITLE%3D%28like%28规模以上工业增加值%29%2Fsen%29&perpage=100

	// IndustrialOutput 工业增加值统一获取接口
	// #：第几页
	// $：每一页几条数据
	IndustrialOutput = "http://www.stats.gov.cn/was5/web/search?page=#&channelid=288041&orderby=-DOCRELTIME&was_custom_expr=DOCTITLE%3D%28like%28规模以上工业增加值%29%2Fsen%29&perpage=$"
)
