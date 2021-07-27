package mysql

import (
	"fmt"
	"strings"
)

const (
	targetValue = "t_dmaa_base_target_value"
)

func dataPatch(data [][]string) string {
	var patch []string
	for _, row := range data {
		s := `('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')`
		s = fmt.Sprintf(s,
			row[0],
			row[1],
			row[2],
			row[3],
			row[4],
			row[5],
			row[6],
			row[7],
			row[8],
			row[9],
			row[10],
			row[11],
			row[12],
			row[13],
			row[14],
			row[15],
		)
		patch = append(patch, s)
	}
	ps := strings.Join(patch, ", ")
	return ps
}

func Generator(data [][]string) string {
	s := `
	INSERT INTO %s
		(
		 VALUE_GUID,
		 TARGET_GUID,
		 TARGET_CODE,
		 TARGET_NAME,
         DATA_SOURCE_CODE,
         DATA_SOURCE_NAME,
		 SOURCE_TARGET_CODE,
		 REGION_CODE,
		 REGION_NAME,
		 UNIT_TYPE,
		 UNIT_NAME,
		 ACCT_YEAR,
		 ACCT_QUARTOR,
		 ACCT_MONTH,
		 ACCT_DATE,
		 TARGET_VALUE
		 )
	VALUES
         %s
	ON DUPLICATE KEY UPDATE
		VALUE_GUID = VALUES(VALUE_GUID),
		TARGET_GUID = VALUES(TARGET_GUID),
		TARGET_CODE = VALUES(TARGET_CODE),
		TARGET_NAME = VALUES(TARGET_NAME),
        DATA_SOURCE_CODE = VALUES(DATA_SOURCE_CODE),
        DATA_SOURCE_NAME = VALUES(DATA_SOURCE_NAME),
		SOURCE_TARGET_CODE = VALUES(SOURCE_TARGET_CODE),
		REGION_CODE = VALUES(REGION_CODE),
		REGION_NAME = VALUES(REGION_NAME),
		UNIT_TYPE = VALUES(UNIT_TYPE),
		UNIT_NAME = VALUES(UNIT_NAME),
		ACCT_YEAR = VALUES(ACCT_YEAR),
		ACCT_QUARTOR = VALUES(ACCT_QUARTOR),
		ACCT_MONTH = VALUES(ACCT_MONTH),
		ACCT_DATE = VALUES(ACCT_DATE),
		TARGET_VALUE = VALUES(TARGET_VALUE);
`
    s = fmt.Sprintf(s, targetValue, dataPatch(data))
	return s
}
