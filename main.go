package main

func main() {}

// 判断接口是否正常运行
/*
gopher@192 ~ % curl http://106.37.165.121:18088/inf/chengtong/py/sy/baseTargetValue/saveRequest
{"timestamp":"2023-01-17 00:25:13","status":400,"error":"Bad Request","exception":"org.springframework.http.converter.HttpMessageNotReadableException","message":"Required request body is missing: public com.baosight.audit.api.dm.aa.response.operation.DmaaBaseTargetValueSaveResp com.baosight.audit.site.controller.py.sy.DmaaBaseTargetValueController.save(com.baosight.audit.api.dm.aa.command.operation.DmaaBaseTargetValueSaveCmd)","path":"/chengtong/py/sy/baseTargetValue/saveRequest"}

gopher@192 ~ % curl http://106.37.165.121:18088/inf/chengtong/py/sy/baseTargetValue/saveRequestx
{"timestamp":"2023-01-17 00:25:20","status":404,"error":"Not Found","message":"No message available","path":"/chengtong/py/sy/baseTargetValue/saveRequestx"}%                                                                                         gopher@192 ~ % curl http://106.37.165.121/inf/chengtong/py/sy/baseTargetValue/saveRequest
curl: (7) Failed to connect to 106.37.165.121 port 80 after 40 ms: Connection refused
*/
