package indicator

var (
	table = "T_DMAA_BASE_TARGET_VALUE"
)

var (
	IndustrySQL = "SELECT ACCT_DATE FROM %s WHERE TARGET_CODE = '%s'"
)