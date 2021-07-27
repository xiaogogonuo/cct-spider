package province_code

// 省份代码

const (
	BeiJingCode      = "110000" // 北京市代码
	TianJinCode      = "120000" // 天津市代码
	HeBeiCode        = "130000" // 河北省代码
	ShanXiCode       = "140000" // 山西省代码
	NeiMengGuCode    = "150000" // 内蒙古自治区代码
	LiaoNingCode     = "210000" // 辽宁省代码
	JiLinCode        = "220000" // 吉林省代码
	HeiLongJiangCode = "230000" // 黑龙江省代码
	ShangHaiCode     = "310000" // 上海市代码
	JiangSuCode      = "320000" // 江苏省代码
	ZheJiangCode     = "330000" // 浙江省代码
	AnHuiCode        = "340000" // 安徽省代码
	FuJianCode       = "350000" // 福建省代码
	JiangXiCode      = "360000" // 江西省代码
	ShanDongCode     = "370000" // 山东省代码
	HeNanCode        = "410000" // 河南省代码
	HuBeiCode        = "420000" // 湖北省代码
	HuNanCode        = "430000" // 湖南省代码
	GuangDongCode    = "440000" // 广东省代码
	GuangXiCode      = "450000" // 广西壮族自治区代码
	HaiNanCode       = "460000" // 海南省代码
	ChongQingCode    = "500000" // 重庆市代码
	SiChuanCode      = "510000" // 四川省代码
	GuiZhouCode      = "520000" // 贵州省代码
	YunNanCode       = "530000" // 云南省代码
	XiZangCode       = "540000" // 西藏自治区代码
	ShanXiiCode      = "610000" // 陕西省代码
	GanSuCode        = "620000" // 甘肃省代码
	QingHaiCode      = "630000" // 青海省代码
	NingXiaCode      = "640000" // 宁夏回族自治区代码
	XinJiangCode     = "650000" // 新疆维吾尔族自治区代码
)

var ProvinceCode map[string]string
var CodeProvince map[string]string

func init() {
	ProvinceCode = make(map[string]string)
	ProvinceCode["北京市"] = BeiJingCode
	ProvinceCode["天津市"] = TianJinCode
	ProvinceCode["河北省"] = HeBeiCode
	ProvinceCode["山西省"] = ShanXiCode
	ProvinceCode["内蒙古自治区"] = NeiMengGuCode
	ProvinceCode["辽宁省"] = LiaoNingCode
	ProvinceCode["吉林省"] = JiLinCode
	ProvinceCode["黑龙江省"] = HeiLongJiangCode
	ProvinceCode["上海市"] = ShangHaiCode
	ProvinceCode["江苏省"] = JiangSuCode
	ProvinceCode["浙江省"] = ZheJiangCode
	ProvinceCode["安徽省"] = AnHuiCode
	ProvinceCode["福建省"] = FuJianCode
	ProvinceCode["江西省"] = JiangXiCode
	ProvinceCode["山东省"] = ShanDongCode
	ProvinceCode["河南省"] = HeNanCode
	ProvinceCode["湖北省"] = HuBeiCode
	ProvinceCode["湖南省"] = HuNanCode
	ProvinceCode["广东省"] = GuangDongCode
	ProvinceCode["广西壮族自治区"] = GuangXiCode
	ProvinceCode["海南省"] = HaiNanCode
	ProvinceCode["重庆市"] = ChongQingCode
	ProvinceCode["四川省"] = SiChuanCode
	ProvinceCode["贵州省"] = GuiZhouCode
	ProvinceCode["云南省"] = YunNanCode
	ProvinceCode["西藏自治区"] = XiZangCode
	ProvinceCode["陕西省"] = ShanXiiCode
	ProvinceCode["甘肃省"] = GanSuCode
	ProvinceCode["青海省"] = QingHaiCode
	ProvinceCode["宁夏回族自治区"] = NingXiaCode
	ProvinceCode["新疆维吾尔自治区"] = XinJiangCode

	CodeProvince = make(map[string]string)
	CodeProvince[BeiJingCode] = "北京市"
	CodeProvince[TianJinCode] = "天津市"
	CodeProvince[HeBeiCode] = "河北省"
	CodeProvince[ShanXiCode] = "山西省"
	CodeProvince[NeiMengGuCode] = "内蒙古自治区"
	CodeProvince[LiaoNingCode] = "辽宁省"
	CodeProvince[JiLinCode] = "吉林省"
	CodeProvince[HeiLongJiangCode] = "黑龙江省"
	CodeProvince[ShangHaiCode] = "上海市"
	CodeProvince[JiangSuCode] = "江苏省"
	CodeProvince[ZheJiangCode] = "浙江省"
	CodeProvince[AnHuiCode] = "安徽省"
	CodeProvince[FuJianCode] = "福建省"
	CodeProvince[JiangXiCode] = "江西省"
	CodeProvince[ShanDongCode] = "山东省"
	CodeProvince[HeNanCode] = "河南省"
	CodeProvince[HuBeiCode] = "湖北省"
	CodeProvince[HuNanCode] = "湖南省"
	CodeProvince[GuangDongCode] = "广东省"
	CodeProvince[GuangXiCode] = "广西壮族自治区"
	CodeProvince[HaiNanCode] = "海南省"
	CodeProvince[ChongQingCode] = "重庆市"
	CodeProvince[SiChuanCode] = "四川省"
	CodeProvince[GuiZhouCode] = "贵州省"
	CodeProvince[YunNanCode] = "云南省"
	CodeProvince[XiZangCode] = "西藏自治区"
	CodeProvince[ShanXiiCode] = "陕西省"
	CodeProvince[GanSuCode] = "甘肃省"
	CodeProvince[QingHaiCode] = "青海省"
	CodeProvince[NingXiaCode] = "宁夏回族自治区"
	CodeProvince[XinJiangCode] = "新疆维吾尔自治区"
}