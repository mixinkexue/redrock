package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//url:https://tieba.baidu.com/f?kw=%E6%8A%97%E5%8E%8B%E8%83%8C%E9%94%85&ie=utf-8&pn=150
func main(){
	resp, err := http.Get("https://tieba.baidu.com/f?kw=%E6%8A%97%E5%8E%8B%E8%83%8C%E9%94%85&ie=utf-8&pn=150")
	if err != nil {
		fmt.Println("http get error", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	str :=string(body)
	regexpTable, _ := regexp.Compile(`target="_blank" class="j_th_tit ">([\w\W]*?)</a>`)

	matchTableSliece:= regexpTable.FindAllStringSubmatch(str, -1)
	for _,v:=range(matchTableSliece){
		fmt.Println(v[1])
	}
}
/*func main()  {

	resp, err := http.Get("http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=2018211743")
	if err != nil {
		fmt.Println("http get error", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	str :=string(body)
	regexpTable, _ := regexp.Compile(`<table ><thead ><tr ><td>教学班分类[\w\W]*</tr></tbody></table>					</div>`)

	matchTableSliece:= regexpTable.FindAllString(str, -1)
	matchTable:=matchTableSliece[0]
	matchTable=strings.Replace(matchTable," ","",-1)
	matchTable=strings.Replace(matchTable,string(byte(13)),"",-1)
	matchTable=strings.Replace(matchTable,string(byte(9)),"",-1)
	regexpRow,_:=regexp.Compile(`<[\w\W]*?>`)
	matchRow:=regexpRow.FindAllString(matchTable,-1)
	for _,v:=range(matchRow){
		matchTable=strings.Replace(matchTable,v,"",-1)
	}
	matchTable=strings.Replace(matchTable,"名单","",-1)
	fmt.Println(matchTable)
	/*lables:=[]string{"理论","理论（含实验）","实验实践"}
	var matchTableNew []string
	for _,v :=range (lables){
		matchTableNew=strings.Split(matchTable,v)

	}
	/*newFile, err:= os.Create("class.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.WriteString(matchTable)
	newFile.Close()*/

