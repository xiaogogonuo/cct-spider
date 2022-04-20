# (可选操作)如果远程服务器在插入数据的过程中，屡次出现主键冲突的问题，则需要对本地爬虫服务器和远程数据库服务器做同步操作
## 第一步、从远程数据库服务器查询所有历史记录，保存到自己电脑，IDE默认另存为t_dmaa_base_target_value.csv
```sql
SELECT TARGET_CODE, REGION_CODE, ACCT_YEAR, ACCT_QUARTOR, ACCT_MONTH, ACCT_DATE FROM t_dmaa_base_target_value;
```
## 第二步、计算每一条记录的md5值，代码参考target/pkg/synchronous/synchronous.go
