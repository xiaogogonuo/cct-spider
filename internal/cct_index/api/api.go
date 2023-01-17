package api

// 东方财富经济数据接口
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

	// 居民消费价格指数
	// CPI 页面展示接口：
	_ = "https://data.eastmoney.com/cjsj/cpi.html"
	// CPI 数据获取接口：
	CPI = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CNATIONAL_SAME%2CNATIONAL_BASE%2CNATIONAL_SEQUENTIAL%2CNATIONAL_ACCUMULATE%2CCITY_SAME%2CCITY_BASE%2CCITY_SEQUENTIAL%2CCITY_ACCUMULATE%2CRURAL_SAME%2CRURAL_BASE%2CRURAL_SEQUENTIAL%2CRURAL_ACCUMULATE&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_CPI"

	// 采购经理人指数
	// PMI 页面展示接口：
	_ = "https://data.eastmoney.com/cjsj/pmi.html"
	// PMI 数据获取接口：
	PMI = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CMAKE_INDEX%2CMAKE_SAME%2CNMAKE_INDEX%2CNMAKE_SAME&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_PMI"

	// 工业品出厂价格指数
	// PPI 页面展示接口：
	_ = "https://data.eastmoney.com/cjsj/ppi.html"
	// PPI 数据获取接口：
	PPI = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CBASE%2CBASE_SAME%2CBASE_ACCUMULATE&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_PPI"

	// 存款准备金率
	// ZBJ 页面展示接口：
	_ = "https://data.eastmoney.com/cjsj/ckzbj.html"
	// ZBJ 数据获取接口：
	ZBJ = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CPUBLISH_DATE%2CTRADE_DATE%2CINTEREST_RATE_BB%2CINTEREST_RATE_BA%2CCHANGE_RATE_B%2CINTEREST_RATE_SB%2CINTEREST_RATE_SA%2CCHANGE_RATE_S%2CNEXT_SH_RATE%2CNEXT_SZ_RATE%2CREMARK&pageNumber=1&pageSize=2000&sortColumns=PUBLISH_DATE%2CTRADE_DATE&sortTypes=-1%2C-1&reportName=RPT_ECONOMY_DEPOSIT_RESERVE"

	// 海关进出口
	// JCK 页面展示接口：
	_ = "https://data.eastmoney.com/cjsj/hgjck.html"
	// JCK 数据获取接口：
	JCK = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CEXIT_BASE%2CIMPORT_BASE%2CEXIT_BASE_SAME%2CIMPORT_BASE_SAME%2CEXIT_BASE_SEQUENTIAL%2CIMPORT_BASE_SEQUENTIAL%2CEXIT_ACCUMULATE%2CIMPORT_ACCUMULATE%2CEXIT_ACCUMULATE_SAME%2CIMPORT_ACCUMULATE_SAME&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_CUSTOMS"

	// 利率调整
	// LL 页面展示接口：
	_ = "https://data.eastmoney.com/cjsj/yhll.html"
	// LL 数据获取接口：
	LL = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CPUBLISH_DATE%2CDEPOSIT_RATE_BB%2CDEPOSIT_RATE_BA%2CDEPOSIT_RATE_B%2CLOAN_RATE_SB%2CLOAN_RATE_SA%2CLOAN_RATE_S%2CNEXT_SH_RATE%2CNEXT_SZ_RATE&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_DEPOSIT_RATE"

	// 外汇和黄金储备
	// WH 页面展示接口：
	_ = "https://data.eastmoney.com/cjsj/hjwh.html"
	// WH 数据获取接口：
	WH = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CGOLD_RESERVES%2CGOLD_RESERVES_SAME%2CGOLD_RESERVES_SEQUENTIAL%2CFOREX%2CFOREX_SAME%2CFOREX_SEQUENTIAL&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_GOLD_CURRENCY"
)

// 东方财富拆借利率接口
const (
	// 银行间拆借利率
	// LendingRate 页面展示接口：
	_ = "https://data.eastmoney.com/shibor/default.html"
	// LendingRate 数据获取接口：
	// 上海银行同业拆借市场-Shibor人民币-隔夜
	// https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22001%22)&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE
	//
	// 上海银行同业拆借市场-Shibor人民币-1周
	// https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22101%22)&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE
	//
	// 上海银行同业拆借市场-Shibor人民币-1月
	// https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22201%22)&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE
	//
	// 上海银行同业拆借市场-Shibor人民币-3月
	// https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22203%22)&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE
	//
	// 上海银行同业拆借市场-Shibor人民币-1年
	// https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22301%22)&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE
	//
	// 伦敦银行同业拆借市场-Libor美元-隔夜(O/N)
	// https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22003%22)(CURRENCY_CODE%3D%22USD%22)(INDICATOR_ID%3D%22001%22)&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE
	//
	// 伦敦银行同业拆借市场-Libor美元-1月
	// https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22003%22)(CURRENCY_CODE%3D%22USD%22)(INDICATOR_ID%3D%22201%22)&pageNumber=1&pageSize=20&sortTypes=-1&sortColumns=REPORT_DATE
	//
	// 伦敦银行同业拆借市场-Libor美元-3月
	// https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22003%22)(CURRENCY_CODE%3D%22USD%22)(INDICATOR_ID%3D%22203%22)&pageNumber=1&pageSize=20&sortTypes=-1&sortColumns=REPORT_DATE
	LendingRate = "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=#&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE"
)

// 东方财富行业指数接口
const (
	// Industry
	// BDI 波罗的海干散货指数
	// https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=2000&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI00107664%22)
	//
	// LPI 物流业景气指数
	// https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=2000&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI00352262%22)

	// CCI 中国大宗商品指数
	// https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=2000&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI00662535%22)

	// CONC 美原油指数
	// https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=2000&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI01508580%22)

	// SOX 费城半导体指数
	// https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=2000&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI00055562%22)

	//Industry 行业数据获取接口：
	Industry = "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=2000&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=#"
)
