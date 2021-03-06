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
"TargetNameEN": "GDP同比", 
"TargetCode": "HG00001", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "20", 
"SourceTargetCodeTable": "mkt200", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "20", 
"PeriodName": "季",
"RunTime": "20:00~22:00"
}, 

{
"Name": "工业增加值同比增长", 
"Case": "eastMoneyHG", 
"TargetNameEN": "工业增加值同比", 
"TargetCode": "HG00016", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "0", 
"SourceTargetCodeTable": "mkt00", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "30", 
"PeriodName": "月",
"RunTime": "20:00~22:00"
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
"PeriodName": "月",
"RunTime": "20:00~22:00"
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
"PeriodName": "月",
"RunTime": "20:00~22:00"
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
"PeriodName": "月",
"RunTime": "20:00~22:00"
}, 

{
"Name": "社会消费品零售总额同比增长", 
"Case": "eastMoneyHG", 
"TargetNameEN": "社会消费品零售总额同比", 
"TargetCode": "HG00029", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "5", 
"SourceTargetCodeTable": "mkt52", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "30", 
"PeriodName": "月",
"RunTime": "20:00~22:00"
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
"PeriodName": "月",
"RunTime": "20:00~22:00"
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
"PeriodName": "月",
"RunTime": "20:00~22:00"
}, 

{
"Name": "货币和准货币(M2)供应量同比增长", 
"Case": "eastMoneyHG", 
"TargetNameEN": "M2同比", 
"TargetCode": "HG00007", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "11", 
"SourceTargetCodeTable": "mkt111", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "30", 
"PeriodName": "月",
"RunTime": "20:00~22:00"
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
"PeriodName": "月",
"RunTime": "20:00~22:00"
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
"PeriodName": "月",
"RunTime": "20:00~22:00"
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
"PeriodName": "月",
"RunTime": "20:00~22:00"
}, 

{
"Name": "存款准备金率", 
"Case": "eastMoneyHG", 
"TargetNameEN": "CKZBJ", 
"TargetCode": "HG00066", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "23", 
"SourceTargetCodeTable": "mkt230", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "30", 
"PeriodName": "月",
"RunTime": "20:00~22:00"
}, 

{
"Name": "出口当月同比增速", 
"Case": "eastMoneyHG", 
"TargetNameEN": "HGCK", 
"TargetCode": "HG00065", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "1", 
"SourceTargetCodeTable": "mkt10", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "30", 
"PeriodName": "月",
"RunTime": "20:00~22:00"
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
"PeriodName": "日",
"RunTime": "20:00~22:00"
},

{
"Name": "存款基准利率", 
"Case": "eastMoneyHG", 
"TargetNameEN": "DPR", 
"TargetCode": "HY00011", 
"DataSourceCode": "eastmoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "13", 
"SourceTargetCodeTable": "mkt131", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "20:00~22:00"
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
"PeriodName": "日",
"RunTime": "20:00~22:00"
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
"PeriodName": "日",
"RunTime": "20:00~22:00"
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
"PeriodName": "日",
"RunTime": "20:00~22:00"
}, 

{
"Name": "铁矿石主力合约", 
"Case": "eastMoneyIM", 
"TargetNameEN": "IM", 
"TargetCode": "HG00061", 
"DataSourceCode": "eastMoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "IM", 
"SourceTargetCodeTable": "IM", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "20:00~22:00"
},

{
"Name": "沪银主力", 
"Case": "eastMoneyAGM", 
"TargetNameEN": "AGM", 
"TargetCode": "TODO", 
"DataSourceCode": "eastMoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "AGM", 
"SourceTargetCodeTable": "AGM", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "20:00~22:00"
},

{
"Name": "沪金主力", 
"Case": "eastMoneyAUM", 
"TargetNameEN": "AUM", 
"TargetCode": "TODO", 
"DataSourceCode": "eastMoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "AUM", 
"SourceTargetCodeTable": "AUM", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "20:00~22:00"
},

{
"Name": "沪铜主力", 
"Case": "eastMoneyCUM", 
"TargetNameEN": "CUM", 
"TargetCode": "TODO", 
"DataSourceCode": "eastMoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "CUM", 
"SourceTargetCodeTable": "CUM", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "20:00~22:00"
},

{
"Name": "沪镍主力", 
"Case": "eastMoneyNIM", 
"TargetNameEN": "NIM", 
"TargetCode": "TODO", 
"DataSourceCode": "eastMoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "NIM", 
"SourceTargetCodeTable": "NIM", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "20:00~22:00"
},

{
"Name": "豆油主力", 
"Case": "eastMoneyYM", 
"TargetNameEN": "YM", 
"TargetCode": "TODO", 
"DataSourceCode": "eastMoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "YM", 
"SourceTargetCodeTable": "YM", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "20:00~22:00"
},

{
"Name": "玉米主力", 
"Case": "eastMoneyCM", 
"TargetNameEN": "CM", 
"TargetCode": "TODO", 
"DataSourceCode": "eastMoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "CM", 
"SourceTargetCodeTable": "CM", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "20:00~22:00"
},

{
"Name": "原油主力", 
"Case": "eastMoneySCM", 
"TargetNameEN": "SCM", 
"TargetCode": "TODO", 
"DataSourceCode": "eastMoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "SCM", 
"SourceTargetCodeTable": "SCM", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "20:00~22:00"
},

{
"Name": "中债10年期国债到期收益率", 
"Case": "eastMoneyGCHN10", 
"TargetNameEN": "GCHN10", 
"TargetCode": "HG00062", 
"DataSourceCode": "eastMoney", 
"DataSourceName": "东方财富", 
"SourceTargetCode": "GCHN10", 
"SourceTargetCodeTable": "GCHN10", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "20:00~22:00"
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
"PeriodName": "日",
"RunTime": "20:00~22:00"
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
"PeriodName": "年",
"RunTime": "20:00~22:00"
}, 

{
"Name": "地区居民消费价格指数", 
"Case": "sinaRegionCPI", 
"TargetNameEN": "RegionCPI", 
"TargetCode": "HG00040", 
"DataSourceCode": "sina", 
"DataSourceName": "新浪财经", 
"SourceTargetCode": "event2", 
"SourceTargetCodeTable": "event2", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "30", 
"PeriodName": "月",
"RunTime": "20:00~22:00"
}, 

{
"Name": "居民消费价格指数", 
"Case": "sinaCPI", 
"TargetNameEN": "CPI", 
"TargetCode": "HG00003", 
"DataSourceCode": "sina", 
"DataSourceName": "新浪财经", 
"SourceTargetCode": "event0", 
"SourceTargetCodeTable": "event0", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "30", 
"PeriodName": "月",
"RunTime": "20:00~22:00"
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
"PeriodName": "日",
"RunTime": "20:00~22:00"
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
"Type": "2",
"RunTime": "20:00~22:00"
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
"Type": "2",
"RunTime": "20:00~22:00"
},

{
"Name": "美元人民币", 
"Case": "fxWHMP", 
"TargetNameEN": "MUSD", 
"TargetCode": "HG00059", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "MUSD", 
"SourceTargetCodeTable": "MUSD", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日", 
"RunTime": "00:00~00:00"
},

{
"Name": "港元人民币", 
"Case": "fxWHMP", 
"TargetNameEN": "MHKD", 
"TargetCode": "HG00060", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "MHKD", 
"SourceTargetCodeTable": "MHKD", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日", 
"RunTime": "00:00~00:00"
}, 

{
"Name": "欧元美元", 
"Case": "fxWH", 
"TargetNameEN": "EURUSD", 
"TargetCode": "HG00045", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "EURUSD", 
"SourceTargetCodeTable": "EURUSD", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "欧元人民币", 
"Case": "fxWH", 
"TargetNameEN": "EURCNY", 
"TargetCode": "TODO", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "EURCNY", 
"SourceTargetCodeTable": "EURCNY", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "美元日元", 
"Case": "fxWH", 
"TargetNameEN": "USDJPY", 
"TargetCode": "HG00046", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "USDJPY", 
"SourceTargetCodeTable": "USDJPY", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "英镑美元", 
"Case": "fxWH", 
"TargetNameEN": "GBPUSD", 
"TargetCode": "HG00047", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "GBPUSD", 
"SourceTargetCodeTable": "GBPUSD", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "美元指数", 
"Case": "fxWH", 
"TargetNameEN": "USDX", 
"TargetCode": "HY00009", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "USDX", 
"SourceTargetCodeTable": "USDX", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "布伦特原油连续", 
"Case": "fxIPE", 
"TargetNameEN": "OILC", 
"TargetCode": "HG00048", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "OILC", 
"SourceTargetCodeTable": "OILC", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "纽约黄金连续", 
"Case": "fxCOMEX", 
"TargetNameEN": "GLNC", 
"TargetCode": "HG00049", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "GLNC", 
"SourceTargetCodeTable": "GLNC", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "纽约白银连续", 
"Case": "fxCOMEX", 
"TargetNameEN": "SLNC", 
"TargetCode": "HG00050", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "SLNC", 
"SourceTargetCodeTable": "SLNC", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "LME铜", 
"Case": "fxLME", 
"TargetNameEN": "LMCI", 
"TargetCode": "HG00051", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "LMCI", 
"SourceTargetCodeTable": "LMCI", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "LME镍", 
"Case": "fxLME", 
"TargetNameEN": "LMNI", 
"TargetCode": "TODO", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "LMNI", 
"SourceTargetCodeTable": "LMNI", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "LME铝", 
"Case": "fxLME", 
"TargetNameEN": "LMAI", 
"TargetCode": "TODO", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "LMAI", 
"SourceTargetCodeTable": "LMAI", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "美玉米连续", 
"Case": "fxCBOT", 
"TargetNameEN": "CRCC", 
"TargetCode": "TODO", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "CRCC", 
"SourceTargetCodeTable": "CRCC", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "美黄豆连续", 
"Case": "fxCBOT", 
"TargetNameEN": "ZSCC", 
"TargetCode": "TODO", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "ZSCC", 
"SourceTargetCodeTable": "ZSCC", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "美债10年收益率", 
"Case": "fxGJZQ", 
"TargetNameEN": "USG10Y", 
"TargetCode": "HG00052", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "USG10Y", 
"SourceTargetCodeTable": "USG10Y", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
},

{
"Name": "日债10年收益率", 
"Case": "fxGJZQ", 
"TargetNameEN": "GJGB10", 
"TargetCode": "HG00055", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "GJGB10", 
"SourceTargetCodeTable": "GJGB10", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "16:00~17:00"
},

{
"Name": "日经225", 
"Case": "fxGJZS", 
"TargetNameEN": "NIKKI", 
"TargetCode": "HG00041", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "NIKKI", 
"SourceTargetCodeTable": "NIKKI", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "16:00~17:00"
}, 

{
"Name": "英国FTSE100", 
"Case": "fxGJZS", 
"TargetNameEN": "FTSE", 
"TargetCode": "HG00056", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "FTSE", 
"SourceTargetCodeTable": "FTSE", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
},

{
"Name": "德国DAX30", 
"Case": "fxGJZS", 
"TargetNameEN": "DAX", 
"TargetCode": "TODO", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "DAX", 
"SourceTargetCodeTable": "DAX", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
},

{
"Name": "法国CAC40", 
"Case": "fxGJZS", 
"TargetNameEN": "CAC", 
"TargetCode": "TODO", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "CAC", 
"SourceTargetCodeTable": "CAC", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
},

{
"Name": "意大利MIB", 
"Case": "fxGJZS", 
"TargetNameEN": "MIB", 
"TargetCode": "TODO", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "MIB", 
"SourceTargetCodeTable": "MIB", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
},

{
"Name": "(加拿大)S&P/TSX 60", 
"Case": "fxGJZS", 
"TargetNameEN": "SPSX", 
"TargetCode": "TODO", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "SPSX", 
"SourceTargetCodeTable": "SPSX", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
},

{
"Name": "英债10年收益率", 
"Case": "fxGJZQ", 
"TargetNameEN": "GUKG10", 
"TargetCode": "HG00054", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "GUKG10", 
"SourceTargetCodeTable": "GUKG10", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "斯托克600", 
"Case": "fxGJZS", 
"TargetNameEN": "SXO", 
"TargetCode": "HG00042", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "SXO", 
"SourceTargetCodeTable": "SXO", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "德债10年收益率", 
"Case": "fxGJZQ", 
"TargetNameEN": "GDBR10", 
"TargetCode": "HG00053", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "GDBR10", 
"SourceTargetCodeTable": "GDBR10", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "纳斯达克指数", 
"Case": "fxGJZS", 
"TargetNameEN": "NASDAQ", 
"TargetCode": "HG00057", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "NASDAQ", 
"SourceTargetCodeTable": "NASDAQ", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
},

{
"Name": "道琼斯工业指数", 
"Case": "fxGJZS", 
"TargetNameEN": "DJIA", 
"TargetCode": "HG00043", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "DJIA", 
"SourceTargetCodeTable": "DJIA", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "标普500", 
"Case": "fxGJZS", 
"TargetNameEN": "SP500", 
"TargetCode": "HG00044", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "SP500", 
"SourceTargetCodeTable": "SP500", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "00:00~00:00"
}, 

{
"Name": "恒生指数", 
"Case": "fxGJZS", 
"TargetNameEN": "HSZS", 
"TargetCode": "HG00058", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "hkHSI", 
"SourceTargetCodeTable": "hkHSI", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日", 
"RunTime": "16:00~17:00"
}, 

{
"Name": "美元LIBOR隔夜", 
"Case": "fxLibor", 
"TargetNameEN": "USLibor", 
"TargetCode": "HG00063", 
"DataSourceCode": "fx678", 
"DataSourceName": "汇通财经", 
"SourceTargetCode": "USLibor", 
"SourceTargetCodeTable": "USLibor", 
"IsQuantity": "Y", 
"UnitType": "50", 
"UnitName": "%", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "20:00~22:00"
}, 

{
"Name": "上海银行间同业拆放利率隔夜", 
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
"PeriodName": "日",
"RunTime": "20:00~22:00"
}, 

{
"Name": "人民币指数", 
"Case": "cni", 
"TargetNameEN": "CNYX", 
"TargetCode": "HG00064", 
"DataSourceCode": "cni", 
"DataSourceName": "国证指数", 
"SourceTargetCode": "CNYX", 
"SourceTargetCodeTable": "CNYX", 
"IsQuantity": "Y", 
"UnitType": "", 
"UnitName": "", 
"PeriodType": "50", 
"PeriodName": "日",
"RunTime": "17:00~19:00"
}

]`
