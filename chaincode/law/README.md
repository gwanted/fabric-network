## chaincode 描述文档

* addLaw 方法  
  方法描述: 添加law方法  
  参数:(string)   
  参数描述： law结构体json字符串
  
    | 字段 | 类型 |含义|  
    |:----:|:----:|:----:|  
    | lawNumber | string | 编号 |   
    | lawTitle | string | 标题 |  
    | lawOrg  | string | 颁布机构 |
    | blockNumber | string | 区块号 |
    | publishTime | int64 | 颁布时间 |
    | lawHash | string | hash |
    | implementationTime | int64 | 施行时间 |
    | file | string | 文件内容（base64） |
    | fileAddress | string | 下载地址 |
  
  例子: {"lawNumber":"1","lawTitle":"中华人民共和国税收法","lawOrg":"司法部","blockNumber":"13","lawHash":"hash3","publishTime":"1402929202020","implementationTime":"1432020392200","file":"8d9wisd99ws0s0s==","fileAddress":"down..."}

* get 方法  
  方法描述: 根据law编号获取信息
  参数:(string)  
  参数描述：编号  
  例子： 1
  返回: lawjson