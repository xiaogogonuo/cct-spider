package index

// ConfigString 指标配置信息
// 一个指标的同比和累计可能来自于一个api，即它们的SourceTargetCode是一致的
// 为了区分同比和累计，增加字段SourceTargetCodeTable，最后将其写入数据库的SourceTargetCode字段
// 对于无区分的指标，SourceTargetCode和SourceTargetCodeTable一致
// 对于有区分的指标，SourceTargetCode和SourceTargetCodeTable不一致
var ConfigString = `[
{
"Name": "国内生产总值同比增长", 
"Case": "eastMoneyHG", 
"TargetNameEN": "GDP", 
"TargetCode": "HG00001", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "20", 
"SourceTargetCodeTable": "mkt200", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "20", 
"PeriodName": "季"
}, 

{
"Name": "工业增加值同比增长", 
"Case": "eastMoneyHG", 
"TargetNameEN": "IavTB", 
"TargetCode": "HG00016", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "0", 
"SourceTargetCodeTable": "mkt00", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "30", 
"PeriodName": "月"
}, 

{
"Name": "工业增加值累计增长", 
"Case": "eastMoneyHG", 
"TargetNameEN": "IavLJ", 
"TargetCode": "HG00017", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "0", 
"SourceTargetCodeTable": "mkt01", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "30", 
"PeriodName": "月"
}, 

{
"Name": "社会消费品零售总额当期值", 
"Case": "eastMoneyHG", 
"TargetNameEN": "ScgDQZ", 
"TargetCode": "HG00027", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "5", 
"SourceTargetCodeTable": "mkt50", 
"IsQuantity": "Y", 
"UnitType": "32", 
"UnitName": "亿元", 
"PeriodType": "30", 
"PeriodName": "月"
}, 

{
"Name": "社会消费品零售总额累计值", 
"Case": "eastMoneyHG", 
"TargetNameEN": "ScgLJZ", 
"TargetCode": "HG00028", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "5", 
"SourceTargetCodeTable": "mkt51", 
"IsQuantity": "Y", 
"UnitType": "32", 
"UnitName": "亿元", 
"PeriodType": "30", 
"PeriodName": "月"
}, 

{
"Name": "社会消费品零售总额同比增长", 
"Case": "eastMoneyHG", 
"TargetNameEN": "ScgTB", 
"TargetCode": "HG00029", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "5", 
"SourceTargetCodeTable": "mkt52", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "30", 
"PeriodName": "月"
}, 

{
"Name": "社会消费品零售总额累计增长", 
"Case": "eastMoneyHG", 
"TargetNameEN": "ScgLJ", 
"TargetCode": "HG00030", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "5", 
"SourceTargetCodeTable": "mkt53", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "30", 
"PeriodName": "月"
}, 

{
"Name": "货币和准货币(M2)供应量期末值", 
"Case": "eastMoneyHG", 
"TargetNameEN": "M2QMZ", 
"TargetCode": "HG00006", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "11", 
"SourceTargetCodeTable": "mkt110", 
"IsQuantity": "Y", 
"UnitType": "32", 
"UnitName": "亿元", 
"PeriodType": "30", 
"PeriodName": "月"
}, 

{
"Name": "货币和准货币(M2)供应量同比增长", 
"Case": "eastMoneyHG", 
"TargetNameEN": "M2TB", 
"TargetCode": "HG00007", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "11", 
"SourceTargetCodeTable": "mkt111", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "30", 
"PeriodName": "月"
}, 

{
"Name": "居民消费价格指数当月", 
"Case": "eastMoneyHG", 
"TargetNameEN": "CPI", 
"TargetCode": "HG00004", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "19", 
"SourceTargetCodeTable": "mkt19", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "30", 
"PeriodName": "月"
}, 

{
"Name": "制造业采购经理指数", 
"Case": "eastMoneyHG", 
"TargetNameEN": "PMI", 
"TargetCode": "HG00020", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "21", 
"SourceTargetCodeTable": "mkt21", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "30", 
"PeriodName": "月"
}, 

{
"Name": "工业品出厂价格指数当月", 
"Case": "eastMoneyHG", 
"TargetNameEN": "PPI", 
"TargetCode": "HG00023", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "22", 
"SourceTargetCodeTable": "mkt220", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "30", 
"PeriodName": "月"
}, 

{
"Name": "贷款基准利率", 
"Case": "eastMoneyHG", 
"TargetNameEN": "LPR", 
"TargetCode": "HY00007", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "13", 
"SourceTargetCodeTable": "mkt130", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "50", 
"PeriodName": "日"
},

{
"Name": "波罗的海干散货指数", 
"Case": "eastMoneyHY", 
"TargetNameEN": "BDI", 
"TargetCode": "HY00003", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "EMI00107664", 
"SourceTargetCodeTable": "EMI00107664", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日"
}, 

{
"Name": "物流业景气指数", 
"Case": "eastMoneyHY", 
"TargetNameEN": "LPI", 
"TargetCode": "HY00001", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "EMI00352262", 
"SourceTargetCodeTable": "EMI00352262", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日"
}, 

{
"Name": "中国大宗商品指数", 
"Case": "eastMoneyHY", 
"TargetNameEN": "CCI", 
"TargetCode": "HY00002", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "EMI00662535", 
"SourceTargetCodeTable": "EMI00662535", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日"
}, 

{
"Name": "上海银行间同业拆放利率", 
"Case": "eastMoneySHIBOR", 
"TargetNameEN": "SHIBOR", 
"TargetCode": "HY00006", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "EMI99221", 
"SourceTargetCodeTable": "EMI99221", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "50", 
"PeriodName": "日"
}, 

{
"Name": "人民币汇率", 
"Case": "sina", 
"TargetNameEN": "RMBExchangeRate", 
"TargetCode": "HY00008", 
"DataSourceCode": "sina", 
"DataSourceName": "新浪财经", 
"SourceTargetCode": "CNYUSD", 
"SourceTargetCodeTable": "CNYUSD", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日"
}, 

{
"Name": "地区生产总值", 
"Case": "sinaRegionGDP", 
"TargetNameEN": "RegionGDP", 
"TargetCode": "HG00002", 
"DataSourceCode": "sina", 
"DataSourceName": "新浪财经", 
"SourceTargetCode": "event7", 
"SourceTargetCodeTable": "event7", 
"IsQuantity": "Y", 
"UnitType": "32", 
"UnitName": "亿元", 
"PeriodType": "10", 
"PeriodName": "年"
}, 

{
"Name": "美元指数", 
"Case": "sina", 
"TargetNameEN": "USDX", 
"TargetCode": "HY00009", 
"DataSourceCode": "sina", 
"DataSourceName": "新浪财经", 
"SourceTargetCode": "DINIW", 
"SourceTargetCodeTable": "DINIW", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日"
}, 

{
"Name": "国债指数", 
"Case": "ifeng", 
"TargetNameEN": "TBI", 
"TargetCode": "HY00005", 
"DataSourceCode": "ifeng", 
"DataSourceName": "凤凰网财经", 
"SourceTargetCode": "sh000012", 
"SourceTargetCodeTable": "sh000012", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日"
}, 

{
"Name": "造纸行业价格指数", 
"Case": "sci", 
"TargetNameEN": "PII", 
"TargetCode": "HY00010", 
"DataSourceCode": "sci", 
"DataSourceName": "卓创资讯", 
"SourceTargetCode": "SCIPII", 
"SourceTargetCodeTable": "SCIPII", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日", 
"HY": "造纸", 
"Level": "0", 
"Path1": "造纸行业价格指数", 
"Path2": "", 
"Path3": "", 
"Path4": "", 
"Type": "2"
}, 

{
"Name": "原油价格指数", 
"Case": "sci", 
"TargetNameEN": "COI", 
"TargetCode": "HY00004", 
"DataSourceCode": "sci", 
"DataSourceName": "卓创资讯", 
"SourceTargetCode": "SCICOI", 
"SourceTargetCodeTable": "SCICOI", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日", 
"HY": "造纸", 
"Level": "0", 
"Path1": "造纸行业价格指数", 
"Path2": "", 
"Path3": "", 
"Path4": "", 
"Type": "2"
}
]`
