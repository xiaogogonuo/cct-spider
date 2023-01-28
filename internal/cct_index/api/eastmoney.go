package api

// 东方财富-经济数据接口
const (
	// GDP 国内生产总值同比增长 、国内生产总值环比增长
	// • 页面展示接口：https://data.eastmoney.com/cjsj/gdp.html
	// • 数据获取接口：
	GDP = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CDOMESTICL_PRODUCT_BASE%2CFIRST_PRODUCT_BASE%2CSECOND_PRODUCT_BASE%2CTHIRD_PRODUCT_BASE%2CSUM_SAME%2CFIRST_SAME%2CSECOND_SAME%2CTHIRD_SAME&pageNumber=1&pageSize=500&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_GDP"

	// IAV 工业增加值
	// • 页面展示接口：https://data.eastmoney.com/cjsj/gyzjz.html
	// • 数据获取接口：
	IAV = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CBASE_SAME%2CBASE_ACCUMULATE&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_INDUS_GROW"

	// XFP 社会消费品零售总额
	// • 页面展示接口：https://data.eastmoney.com/cjsj/xfp.html
	// • 数据获取接口：
	XFP = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CRETAIL_TOTAL%2CRETAIL_TOTAL_SAME%2CRETAIL_TOTAL_SEQUENTIAL%2CRETAIL_TOTAL_ACCUMULATE%2CRETAIL_ACCUMULATE_SAME&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_TOTAL_RETAIL"

	// M2 货币和准货币(M2)供应量
	// • 页面展示接口：https://data.eastmoney.com/cjsj/hbgyl.html
	// • 数据获取接口：
	M2 = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CBASIC_CURRENCY%2CBASIC_CURRENCY_SAME%2CBASIC_CURRENCY_SEQUENTIAL%2CCURRENCY%2CCURRENCY_SAME%2CCURRENCY_SEQUENTIAL%2CFREE_CASH%2CFREE_CASH_SAME%2CFREE_CASH_SEQUENTIAL&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_CURRENCY_SUPPLY"

	// CPI 居民消费价格指数
	// • 页面展示接口：https://data.eastmoney.com/cjsj/cpi.html
	// • 数据获取接口：
	CPI = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CNATIONAL_SAME%2CNATIONAL_BASE%2CNATIONAL_SEQUENTIAL%2CNATIONAL_ACCUMULATE%2CCITY_SAME%2CCITY_BASE%2CCITY_SEQUENTIAL%2CCITY_ACCUMULATE%2CRURAL_SAME%2CRURAL_BASE%2CRURAL_SEQUENTIAL%2CRURAL_ACCUMULATE&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_CPI"

	// PMI 采购经理人指数
	// • 页面展示接口：https://data.eastmoney.com/cjsj/pmi.html
	// • 数据获取接口：
	PMI = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CMAKE_INDEX%2CMAKE_SAME%2CNMAKE_INDEX%2CNMAKE_SAME&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_PMI"

	// PPI 工业品出厂价格指数
	// • 页面展示接口：https://data.eastmoney.com/cjsj/ppi.html
	// • 数据获取接口：
	PPI = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CBASE%2CBASE_SAME%2CBASE_ACCUMULATE&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_PPI"

	// ZBJ 存款准备金率
	// • 页面展示接口：https://data.eastmoney.com/cjsj/ckzbj.html
	// • 数据获取接口：
	ZBJ = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CPUBLISH_DATE%2CTRADE_DATE%2CINTEREST_RATE_BB%2CINTEREST_RATE_BA%2CCHANGE_RATE_B%2CINTEREST_RATE_SB%2CINTEREST_RATE_SA%2CCHANGE_RATE_S%2CNEXT_SH_RATE%2CNEXT_SZ_RATE%2CREMARK&pageNumber=1&pageSize=2000&sortColumns=PUBLISH_DATE%2CTRADE_DATE&sortTypes=-1%2C-1&reportName=RPT_ECONOMY_DEPOSIT_RESERVE"

	// JCK 海关进出口
	// • 页面展示接口：https://data.eastmoney.com/cjsj/hgjck.html
	// • 数据获取接口：
	JCK = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CEXIT_BASE%2CIMPORT_BASE%2CEXIT_BASE_SAME%2CIMPORT_BASE_SAME%2CEXIT_BASE_SEQUENTIAL%2CIMPORT_BASE_SEQUENTIAL%2CEXIT_ACCUMULATE%2CIMPORT_ACCUMULATE%2CEXIT_ACCUMULATE_SAME%2CIMPORT_ACCUMULATE_SAME&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_CUSTOMS"

	// LL 利率调整
	// • 页面展示接口：https://data.eastmoney.com/cjsj/yhll.html
	// • 数据获取接口：
	LL = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CPUBLISH_DATE%2CDEPOSIT_RATE_BB%2CDEPOSIT_RATE_BA%2CDEPOSIT_RATE_B%2CLOAN_RATE_SB%2CLOAN_RATE_SA%2CLOAN_RATE_S%2CNEXT_SH_RATE%2CNEXT_SZ_RATE&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_DEPOSIT_RATE"

	// WH 外汇和黄金储备
	// • 页面展示接口：https://data.eastmoney.com/cjsj/hjwh.html
	// • 数据获取接口：
	WH = "https://datacenter-web.eastmoney.com/api/data/v1/get?columns=REPORT_DATE%2CTIME%2CGOLD_RESERVES%2CGOLD_RESERVES_SAME%2CGOLD_RESERVES_SEQUENTIAL%2CFOREX%2CFOREX_SAME%2CFOREX_SEQUENTIAL&pageNumber=1&pageSize=2000&sortColumns=REPORT_DATE&sortTypes=-1&reportName=RPT_ECONOMY_GOLD_CURRENCY"
)

// 东方财富-拆借利率接口
const (
	// 上海银行同业拆借市场-Shibor人民币-隔夜
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22001%22)&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE

	// 上海银行同业拆借市场-Shibor人民币-1周
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22101%22)&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE

	// 上海银行同业拆借市场-Shibor人民币-1月
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22201%22)&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE

	// 上海银行同业拆借市场-Shibor人民币-3月
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22203%22)&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE

	// 上海银行同业拆借市场-Shibor人民币-1年
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22001%22)(CURRENCY_CODE%3D%22CNY%22)(INDICATOR_ID%3D%22301%22)&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE

	// 伦敦银行同业拆借市场-Libor美元-隔夜(O/N)
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22003%22)(CURRENCY_CODE%3D%22USD%22)(INDICATOR_ID%3D%22001%22)&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE

	// 伦敦银行同业拆借市场-Libor美元-1月
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22003%22)(CURRENCY_CODE%3D%22USD%22)(INDICATOR_ID%3D%22201%22)&pageNumber=1&pageSize=20&sortTypes=-1&sortColumns=REPORT_DATE

	// 伦敦银行同业拆借市场-Libor美元-3月
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=(MARKET_CODE%3D%22003%22)(CURRENCY_CODE%3D%22USD%22)(INDICATOR_ID%3D%22203%22)&pageNumber=1&pageSize=20&sortTypes=-1&sortColumns=REPORT_DATE

	// LendingRate 拆借利率统一数据获取接口
	LendingRate = "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_IMP_INTRESTRATEN&columns=REPORT_DATE%2CIR_RATE%2CCHANGE_RATE%2C&filter=#&pageNumber=1&pageSize=365&sortTypes=-1&sortColumns=REPORT_DATE"
)

// 东方财富-行业指数接口
const (
	// 波罗的海干散货指数
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=2000&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI00107664%22)

	// 物流业景气指数
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=2000&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI00352262%22)

	// 中国大宗商品指数
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=2000&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI00662535%22)

	// 美原油指数
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=2000&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI01508580%22)

	// 费城半导体指数
	// • 页面展示接口：
	// • 数据获取接口：https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=2000&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=(INDICATOR_ID%3D%22EMI00055562%22)

	// Industry 行业指数统一数据获取接口
	Industry = "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=REPORT_DATE&sortTypes=-1&pageSize=2000&pageNumber=1&reportName=RPT_INDUSTRY_INDEX&columns=REPORT_DATE%2CINDICATOR_VALUE%2CCHANGE_RATE%2CCHANGERATE_3M%2CCHANGERATE_6M%2CCHANGERATE_1Y%2CCHANGERATE_2Y%2CCHANGERATE_3Y&filter=#"
)

// 东方财富-全球指数接口
const (
	// 国债指数
	// • 页面展示接口：http://quote.eastmoney.com/zs000012.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=1.000012

	// 台湾加权
	// • 页面展示接口：http://quote.eastmoney.com/gb/zsTWII.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=100.TWII

	// 韩国KOSPI
	// • 页面展示接口：http://quote.eastmoney.com/gb/zsKS11.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=100.KS11

	// 俄罗斯RTS
	// • 页面展示接口：http://quote.eastmoney.com/gb/zsRTS.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=100.RTS

	// 澳大利亚标普200
	// • 页面展示接口：http://quote.eastmoney.com/gb/zsAS51.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=100.AS51

	// 路透CRB商品指数
	// • 页面展示接口：http://quote.eastmoney.com/gb/zsCRB.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=100.CRB

	// 中证国有企业综合指数
	// • 页面展示接口：http://quote.eastmoney.com/zs000955.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=1.000955

	// 中证央企结构调整指数
	// • 页面展示接口：http://quote.eastmoney.com/zs000860.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=1.000860

	// 中证国企一带一路指数
	// • 页面展示接口：http://quote.eastmoney.com/zz/2.000859.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=2.000859

	// 中证中央企业100指数
	// • 页面展示接口：http://quote.eastmoney.com/zs000927.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=1.000927

	// 中证央企创新驱动指数
	// • 页面展示接口：http://quote.eastmoney.com/zz/2.000861.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=2.000861

	// 中证国有企业改革指数
	// • 页面展示接口：http://quote.eastmoney.com/zs399974.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=0.399974

	// 中证中央企业综合指数
	// • 页面展示接口：http://quote.eastmoney.com/zs000926.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=1.000926

	// 银华央企结构调整ETF
	// • 页面展示接口：http://quote.eastmoney.com/sz159959.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=0.159959

	// 博时央企结构调整ETF
	// • 页面展示接口：http://quote.eastmoney.com/sh512960.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=1.512960

	// 华夏央企结构调整ETF
	// • 页面展示接口：http://quote.eastmoney.com/sh512950.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=1.512950

	// 原油指数
	// • 页面展示接口：http://quote.eastmoney.com/q/159.scfi.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=159.scfi

	// AMAC造纸
	// • 页面展示接口：http://quote.eastmoney.com/zz/2.H30049.html
	// • 数据获取接口：http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&secid=2.H30049

	// Global 全球指数统一数据获取接口
	Global = "http://push2.eastmoney.com/api/qt/stock/get?invt=2&fltt=1&fields=f43%2Cf44%2Cf45%2Cf46%2Cf60%2Cf86%2Cf169%2Cf170&#"
)

// 东方财富-中美国债接口
const (
	// NationalDebt 中美国债收益率
	// • 页面展示接口：https://data.eastmoney.com/cjsj/zmgzsyl.html
	// • 数据获取接口：
	NationalDebt = "https://datacenter-web.eastmoney.com/api/data/get?type=RPTA_WEB_TREASURYYIELD&sty=ALL&st=SOLAR_DATE&sr=-1&p=1&ps=100"
)
