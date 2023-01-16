package api

const (
	// 国内生产总值同比增长
	// GDPTb 页面展示接口：
	_ = "https://data.eastmoney.com/cjsj/gdp.html"
	// GDPTb 数据获取接口：
	GDPTb = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CDOMESTICL_PRODUCT_BASE%2CFIRST_PRODUCT_BASE%2CSECOND_PRODUCT_BASE%2CTHIRD_PRODUCT_BASE%2CSUM_SAME%2CFIRST_SAME%2CSECOND_SAME%2CTHIRD_SAME&pageNumber=1&pageSize=500&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_GDP"

	// 工业增加值
	// IAV 页面展示接口：
	_ = "https://data.eastmoney.com/cjsj/gyzjz.html"
	// IAV 数据获取接口：
	IAV = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CBASE_SAME%2CBASE_ACCUMULATE&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_INDUS_GROW"

	// 社会消费品零售总额
	// XFP 页面展示接口：
	_ = "https://data.eastmoney.com/cjsj/xfp.html"
	// XFP 数据获取接口：
	XFP = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CRETAIL_TOTAL%2CRETAIL_TOTAL_SAME%2CRETAIL_TOTAL_SEQUENTIAL%2CRETAIL_TOTAL_ACCUMULATE%2CRETAIL_ACCUMULATE_SAME&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_TOTAL_RETAIL"

	// 货币和准货币(M2)供应量
	// M2 页面展示接口：
	_ = "https://data.eastmoney.com/cjsj/hbgyl.html"
	// M2 数据获取接口：
	M2 = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CBASIC_CURRENCY%2CBASIC_CURRENCY_SAME%2CBASIC_CURRENCY_SEQUENTIAL%2CCURRENCY%2CCURRENCY_SAME%2CCURRENCY_SEQUENTIAL%2CFREE_CASH%2CFREE_CASH_SAME%2CFREE_CASH_SEQUENTIAL&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_CURRENCY_SUPPLY"
)
