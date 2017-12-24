// 但是如何去发送一个真正的http request而不去真正的启动一个http server
// （亦或者请求任意的server）？答案是使用Go 的httptest包，
// 这个包可以非常简单的创建一个测试的http server


package person

import (
"net/http"
"fmt"
"io/ioutil"
"encoding/json"

"github.com/astaxie/beego/logs"
)

const (
	ADDRESS = "shanghai"
)

type Person struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Age int `json:"age"`
}

func GetInfo(api string) ([]Person, error) {
	url := fmt.Sprintf("%s/person?addr=%s", api, ADDRESS)
	resp, err := http.Get(url)

	defer resp.Body.Close()

	if err != nil {
		return []Person{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []Person{}, fmt.Errorf("get info didn’t respond 200 OK: %s", resp.Status)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	personList := make([]Person,0)
	err = json.Unmarshal(bodyBytes, &personList)
	if err != nil {
		logs.Error("decode data fail")
		return []Person{}, fmt.Errorf("decode data fail")
	}

	return personList, nil
}


