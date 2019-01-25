package main

import (
	"fmt"
	/*导入 chaincode shim 包和 peer protobuf 包*/
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"

	"encoding/json"
	"time"
	// "strconv"
	"bytes"
	"strings"
)

const prefix = "urun"

type LawChaincode struct {
}

// 云证
type Law struct {
	LawNumber          string `json:"lawNumber"`//编号
	LawTitle   		   string `json:"lawTitle"`//标题
	BlockNumber        string `json:"blockNumber"`//区块号
	LawOrg             string `json:"lawOrg"`//颁布机构
	PublishTime        int64  `json:"publishTime"`//颁布时间
	ImplementationTime int64  `json:"implementationTime"`//施行时间
	Time               int64  `json:"time"`//创建时间
	LawHash            string `json:"lawHash"`//hash
	File               string `json:"file"`//文件
	FileAddress        string `json:"fileAddress"`//下载地址
}

//初始化方法
func (s *LawChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

//调用Chaincode
func (s *LawChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	//获取要调用的方法名和方法参数
	fn, args := stub.GetFunctionAndParameters()

	fmt.Printf("方法: %s  参数 ： %s \n", fn, args)

	if fn == "addLaw" {
		return s.addLaw(stub, args)
	} else if fn == "get" {
		return s.get(stub, args)
	}

	return shim.Error("方法不存在")
}

func (s *LawChaincode) addLaw(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("参数出错")
	}

	lawStr := args[0]

	var law Law
	//这里就是实际的解码和相关的错误检查
	if err := json.Unmarshal([]byte(lawStr), &law); err != nil {
		return shim.Error("json反序列化失败")
	}

	t := time.Now()
	id := law.LawNumber
	law.Time = t.Unix()

	bys, err := json.Marshal(law)
	fmt.Println("json:" + string(bys))

	if err != nil {
		return shim.Error("json序列化失败")
	}

	err = stub.PutState(id, bys)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(id))
}

func (s *LawChaincode) get(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("要输入一个键")
	}
	//读出
	value, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("json:" + string(value))

	var c Law
	err = json.Unmarshal(value, &c)
	if err != nil {
		return shim.Error("json unmarshal failed "+err.Error())
	}
	//截取一半hash
	c.LawHash = c.LawHash[:len(c.LawHash)/2]

	law,err := json.Marshal(c)
	if err != nil{
		return shim.Error("json marshal failed "+err.Error())
	}

	return shim.Success(law)
}

func selectionCriteria(pmap map[string]string) string {

	var buffer bytes.Buffer
	buffer.WriteString(`{"selector":{`)
	buffer.WriteString(`"lawNumber":{"$regex": "^` + prefix + `.*"},`)

	if pmap["startTime"] != "" && pmap["endTime"] != "" {
		pmap["startTime-endTime"] = pmap["startTime"] + "-" + pmap["endTime"]
		delete(pmap, "startTime")
		delete(pmap, "endTime")
	}

	for k, v := range pmap {

		switch k {

		case "lawHash":
			if v != "" {
				buffer.WriteString(`"lawHash":{"$regex": "^` + v + `.*"},`)
			}
		case "startTime":
			if v != "" {
				buffer.WriteString(`"time":{"$gte": ` + v + `},`)
			}
		case "endTime":
			if v != "" {
				buffer.WriteString(`"time":{"$lte": ` + v + `},`)
			}
		case "startTime-endTime":
			if v != "" {
				args := strings.Split(v, "-")
				buffer.WriteString(`"time":{"$gte": ` + args[0] + `,"$lte": ` + args[1] + `},`)
			}

		case "fileType":
			if v != "" && v != "," {
				types := `"fileType":{"$or":[`

				args := strings.Split(v, ",")
				for i, tyv := range args {
					if i != 0 {
						types += `,`
					}
					types += `{"$eq":"` + tyv + `"}`
				}
				types += `]},`

				buffer.WriteString(types)
			}

		case "fileLabel":
			if v != "" && v != "," {
				buffer.WriteString(`"fileLabel":{"$regex": ".*` + v + `.*"},`)
			}

		default:
			if k != "" && v != "" {
				buffer.WriteString(`"` + k + `":{"$eq": "` + v + `"},`)
			}

		}
	}
	buffer.Truncate(buffer.Len() - 1)

	buffer.WriteString(`}}`)

	return buffer.String()
}

func main() {

	if err := shim.Start(new(LawChaincode)); err != nil {
		fmt.Println("LawChaincode start error")
	}
}
