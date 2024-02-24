package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// 定义外层响应的结构体
type APIResponse struct {
	Error           bool       `json:"error"`
	ConsumedFpoint  int        `json:"consumed_fpoint"`
	RequiredFpoints int        `json:"required_fpoints"`
	Size            int64      `json:"size"`
	Page            int        `json:"page"`
	Mode            string     `json:"mode"`
	Query           string     `json:"query"`
	Results         [][]string `json:"results"` // 将Results定义为字符串的二维数组
}

func main() {
	client := openai.NewClient("API_KEY_Openai")
	messages := make([]openai.ChatCompletionMessage, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	B_text := "{\"prompt\": \"如何使用FOFA查询特定的IP地址？\", \"completion\": \"ip=\\\"1.1.1.1\\\"\"}\n{\"prompt\": \"我想要找到所有运行在80端口的资产。\", \"completion\": \"port=\\\"80\\\"\"}\n{\"prompt\": \"如果我想要查找特定操作系统的资产，比如CentOS，我应该怎么做？\", \"completion\": \"os=\\\"centos\\\"\"}\n{\"prompt\": \"我想要找出所有属于微软IIS 10的服务器。\", \"completion\": \"server=\\\"Microsoft-IIS/10\\\"\"}\n{\"prompt\": \"如何查询特定自治系统号（ASN）的资产？\", \"completion\": \"asn=\\\"19551\\\"\"}\n{\"prompt\": \"我怎样才能找到使用nginx的产品？\", \"completion\": \"product=\\\"NGINX\\\"\"}\n{\"prompt\": \"如果我想要查找特定域名的资产，应该怎么操作？\", \"completion\": \"domain=\\\"qq.com\\\"\"}\n{\"prompt\": \"如何筛选出所有含有特定标题的网站？\", \"completion\": \"title=\\\"beijing\\\"\"}\n{\"prompt\": \"我想要找出所有设置了特定响应头的服务器，应该怎么做？\", \"completion\": \"header=\\\"elastic\\\"\"}\n{\"prompt\": \"如何找到所有开放了特定端口，比如6379的资产？\", \"completion\": \"port=\\\"6379\\\"\"}\n{\"prompt\": \"如何查找所有的ipv6资产？\", \"completion\": \"is_ipv6=true\"}\n{\"prompt\": \"我想要找出不属于任何云服务商的资产，应该怎么操作？\", \"completion\": \"is_cloud=false\"}\n{\"prompt\": \"如何利用FOFA查询语法来查找特定云服务商，比如阿里云的资产？\", \"completion\": \"cloud_name=\\\"Aliyundun\\\"\"}\n{\"prompt\": \"我想要通过网站的ICP备案号来查找资产，应该使用什么语法？\", \"completion\": \"icp=\\\"京ICP证030173号\\\"\"}\n{\"prompt\": \"如果我想要找到所有使用特定JavaScript库的网站，比如jQuery，我应该怎么查询？\", \"completion\": \"js_name=\\\"js/jquery.js\\\"\"}\n{\"prompt\": \"如何找到所有域名解析到特定CNAME记录的资产？\", \"completion\": \"cname=\\\"ap21.inst.siteforce.com\\\"\"}\n{\"prompt\": \"我怎样才能查找出所有使用了特定TLS版本，比如TLS 1.3的资产？\", \"completion\": \"tls.version=\\\"TLS 1.3\\\"\"}\n{\"prompt\": \"如何利用FOFA查询特定的操作系统版本，例如Windows 10的资产？\", \"completion\": \"os=\\\"Windows 10\\\"\"}\n{\"prompt\": \"我想要找出所有在特定时间后更新的资产，应该怎么做？\", \"completion\": \"after=\\\"2023-01-01\\\"\"}\n{\"prompt\": \"如何通过FOFA查询语法来查找特定端口范围内的资产，比如开放了1000到2000之间的端口？\", \"completion\": \"port=\\\"1000\\\" || port=\\\"2000\\\"\"}\n{\"prompt\": \"我想找出所有运行Apache服务器的资产。\", \"completion\": \"server=\\\"Apache\\\"\"}\n{\"prompt\": \"如何查询所有位于中国浙江省的资产？\", \"completion\": \"region=\\\"Zhejiang\\\"\"}\n{\"prompt\": \"我需要找到所有使用了特定SSL证书颁发机构颁发的证书的资产，比如Let's Encrypt。\", \"completion\": \"cert.issuer=\\\"Let's Encrypt\\\"\"}\n{\"prompt\": \"如何查找所有部署了蜜罐的资产？\", \"completion\": \"is_honeypot=true\"}\n{\"prompt\": \"我怎样才能找到所有标记为服务的资产，例如HTTP服务？\", \"completion\": \"type=\\\"service\\\"\"}\n{\"prompt\": \"如何通过域名的WHOIS注册信息中的组织名来查找资产？\", \"completion\": \"org=\\\"LLC Baxet\\\"\"}\n{\"prompt\": \"我想要查找所有具有特定favicon图标hash值的网站，应该如何查询？\", \"completion\": \"icon_hash=\\\"-247388890\\\"\"}\n{\"prompt\": \"如何找到所有响应状态码为404的网站？\", \"completion\": \"status_code=\\\"404\\\"\"}\n{\"prompt\": \"我需要查找所有包含特定文本内容的网页，比如包含'欢迎访问'。\", \"completion\": \"body=\\\"欢迎访问\\\"\"}\n{\"prompt\": \"如何找出所有的资产，它们的端口开放数量大于5个？\", \"completion\": \"port_size_gt=\\\"5\\\"\"}\n{\"prompt\": \"我想通过特定的JavaScript文件MD5值来查找使用该文件的网站，怎么做？\", \"completion\": \"js_md5=\\\"82ac3f14327a8b7ba49baa208d4eaa15\\\"\"}\n{\"prompt\": \"如何使用FOFA查找所有部署在特定云服务上的资产，例如AWS？\", \"completion\": \"cloud_name=\\\"AWS\\\"\"}\n{\"prompt\": \"我怎么能找到所有的资产，它们在特定时间之前最后一次更新，例如2020年1月1日之前？\", \"completion\": \"before=\\\"2020-01-01\\\"\"}\n{\"prompt\": \"如何查询所有的资产，它们使用了特定的TLS指纹，例如JARM指纹？\", \"completion\": \"jarm=\\\"2ad2ad0002ad2ad22c2ad2ad2ad2ad2eac92ec34bcc0cf7520e97547f83e81\\\"\"}\n{\"prompt\": \"我需要找出所有使用了特定TLS JA3S指纹的资产，这个指纹是15af977ce25de452b96affa2addb1036，应该怎么查询？\", \"completion\": \"tls.ja3s=\\\"15af977ce25de452b96affa2addb1036\\\"\"}\n{\"prompt\": \"如何查找所有标题包含'登录'的网页？\", \"completion\": \"title=\\\"登录\\\"\"}\n{\"prompt\": \"我想找出所有包含特定HTML标签的网页，例如包含<div>标签。\", \"completion\": \"body=\\\"<div>\\\"\"}\n{\"prompt\": \"如何利用FOFA找到所有运行特定版本的SSH服务的资产，比如SSH 2.0？\", \"completion\": \"banner=\\\"SSH 2.0\\\"\"}\n{\"prompt\": \"我需要找到所有响应头中包含特定X-Powered-By值的资产，比如PHP/5.3.3。\", \"completion\": \"header=\\\"X-Powered-By: PHP/5.3.3\\\"\"}\n{\"prompt\": \"如何查找所有使用了特定操作系统并开放了特定端口的资产，例如使用Windows 10并开放了3389端口？\", \"completion\": \"os=\\\"Windows 10\\\" && port=\\\"3389\\\"\"}\n{\"prompt\": \"我想要查找所有域名为'.edu'结尾的教育机构网站。\", \"completion\": \"domain=\\\".edu\\\"\"}\n{\"prompt\": \"如何找到所有位于特定城市的资产，例如位于上海的资产？\", \"completion\": \"city=\\\"Shanghai\\\"\"}\n{\"prompt\": \"我想通过网站的Meta标签内容来查找资产，比如<meta name='keywords' content='example'>，应该怎么做？\", \"completion\": \"body=\\\"<meta name='keywords' content='example'>\\\"\"}\n{\"prompt\": \"如何利用FOFA查询同时开放了多个特定端口的资产，例如同时开放了22和80端口的资产？\", \"completion\": \"ip_ports=\\\"22,80\\\"\"}\n{\"prompt\": \"我需要找出所有CNAME记录指向特定域名的资产，例如指向example.com的资产。\", \"completion\": \"cname_domain=\\\"example.com\\\"\"}\n{\"prompt\": \"如何查找所有在特定时间范围内更新过的资产，例如在2023年1月到2023年3月之间更新的资产？\", \"completion\": \"after=\\\"2023-01-01\\\" && before=\\\"2023-03-31\\\"\"}\n{\"prompt\": \"我想找到所有的资产，它们使用了特定的Web框架，比如Django。\", \"completion\": \"header=\\\"Django\\\"\"}\n{\"prompt\": \"如何使用FOFA查询语法找到所有部署了特定版本Web应用的资产，例如Wordpress 5.1？\", \"completion\": \"app=\\\"Wordpress\\\" && banner=\\\"5.1\\\"\"}\n{\"prompt\": \"如何找到所有启用了HTTPS协议的资产？\", \"completion\": \"protocol=\\\"https\\\"\"}\n{\"prompt\": \"我想查找所有的资产，它们的标题中包含'欢迎使用'四个字的网站。\", \"completion\": \"title=\\\"欢迎使用\\\"\"}\n{\"prompt\": \"如何利用FOFA找到所有运行在443端口上的NGINX服务器？\", \"completion\": \"port=\\\"443\\\" && server=\\\"nginx\\\"\"}\n{\"prompt\": \"我需要找出所有由特定ISP提供服务的资产，例如由'中国电信'提供服务的资产。\", \"completion\": \"org=\\\"中国电信\\\"\"}\n{\"prompt\": \"如何查找所有包含特定响应头，比如'Access-Control-Allow-Origin: *'的资产？\", \"completion\": \"header=\\\"Access-Control-Allow-Origin: *\\\"\"}\n{\"prompt\": \"我想通过特定的操作系统版本来筛选资产，例如所有运行Windows Server 2012的资产。\", \"completion\": \"os=\\\"Windows Server 2012\\\"\"}\n{\"prompt\": \"如何找到所有具有特定SSL证书特征的资产，比如证书中包含'Let's Encrypt Authority X3'的？\", \"completion\": \"cert.subject=\\\"Let's Encrypt Authority X3\\\"\"}\n{\"prompt\": \"我需要查找所有包含特定JavaScript文件的网站，例如包含'jquery.min.js'的网站。\", \"completion\": \"js_name=\\\"jquery.min.js\\\"\"}\n{\"prompt\": \"如何查询所有部署在亚马逊云服务(AWS)上的资产？\", \"completion\": \"cloud_name=\\\"AWS\\\"\"}\n{\"prompt\": \"我想找出所有响应状态码为200且标题包含'首页'的网站。\", \"completion\": \"status_code=\\\"200\\\" && title=\\\"首页\\\"\"}\n{\"prompt\": \"如何利用FOFA语法来查找所有设置了特定CNAME记录的资产，例如指向'example.cloudfront.net'的资产？\", \"completion\": \"cname=\\\"example.cloudfront.net\\\"\"}\n{\"prompt\": \"我需要查找所有最近在2023年更新过的资产，如何操作？\", \"completion\": \"after=\\\"2023-01-01\\\"\"}\n{\"prompt\": \"如何查找所有使用了特定加密套件的TLS资产，例如使用'ECDHE-RSA-AES128-GCM-SHA256'的资产？\", \"completion\": \"tls.ja3s=\\\"ECDHE-RSA-AES128-GCM-SHA256\\\"\"}\n{\"prompt\": \"如何找到位于中国浙江省，且运行Apache服务器的资产？\", \"completion\": \"region=\\\"Zhejiang\\\" && server=\\\"Apache\\\"\"}\n{\"prompt\": \"我想查找所有运行在80端口，标题包含'欢迎'字样的网站。\", \"completion\": \"port=\\\"80\\\" && title*=\\\"欢迎\\\"\"}\n{\"prompt\": \"如何筛选所有使用了SSL协议，且域名为.com结尾的资产？\", \"completion\": \"protocol=\\\"https\\\" && domain=\\\"*.com\\\"\"}\n{\"prompt\": \"我需要找到所有部署在AWS云服务上，且是IPv6地址的资产。\", \"completion\": \"cloud_name=\\\"AWS\\\" && is_ipv6=true\"}\n{\"prompt\": \"如何查找所有包含特定响应头'X-Powered-By: PHP/5.4.45'，并且是ipv4资产的网站？\", \"completion\": \"header=\\\"X-Powered-By: PHP/5.4.45\\\" && is_ipv6=false\"}"
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: B_text,
	})
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				MaxTokens: 10000,
				Model:     openai.GPT3Dot5Turbo,
				Messages:  messages,
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		content := resp.Choices[0].Message.Content
		fmt.Println(content)
		//增加访问FOFA的逻辑
		Base64content := base64.StdEncoding.EncodeToString([]byte(content))
		apiKey := "API_KEY_FOFA"
		url := fmt.Sprintf("https://fofa.info/api/v1/search/all?qbase64=%s&key=%s", Base64content, apiKey)

		// 创建HTTP客户端
		client := &http.Client{}
		// 创建请求
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		// 发送请求
		respf, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer respf.Body.Close()

		// 读取响应体
		body, err := ioutil.ReadAll(respf.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// 打印响应体
		//fmt.Println(string(body))

		var RespData APIResponse
		err = json.Unmarshal(body, &RespData)
		if err != nil {
			fmt.Println("Error parsing JSON Response:", err)
			return
		}
		//fmt.Println(RespData.Results)
		for _, result := range RespData.Results {
			fmt.Println(result)
		}
		//增加上文逻辑
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
		// excel logic
		f := excelize.NewFile()
		// 创建一个工作表
		index, err := f.NewSheet("Sheet1")
		if err != nil {
			fmt.Println("NewSheet error :", err)
		}
		// 设置工作表的标题行
		f.SetCellValue("Sheet1", "A1", "IP:Port")
		f.SetCellValue("Sheet1", "B1", "IP")
		f.SetCellValue("Sheet1", "C1", "Port")
		f.SetCellValue("Sheet1", "D1", "Extra")

		row := 2 // Excel文件中的数据行开始于第2行
		for _, result := range RespData.Results {
			// 将每个结果添加到工作表
			f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), result[0])
			f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), result[1])
			f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), result[2])
			if len(result) > 3 {
				f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), result[3])
			}
			row++
		}

		// 设置活动的工作表
		f.SetActiveSheet(index)

		// 保存文件
		if err := f.SaveAs("Result.xlsx"); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Results saved to Excel file: Result.xlsx")
	}
}
