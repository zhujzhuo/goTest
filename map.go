package main

import "fmt"

func main() {
    var countryCapitalMap map[string]string /*创建集合 */
    countryCapitalMap = make(map[string]string)
    //myMap = make(map[string] string, 100)

    /* map插入key - value对,各个国家对应的首都 */
    countryCapitalMap [ "France" ] = "Paris"
    countryCapitalMap [ "Italy" ] = "罗马"
    countryCapitalMap [ "Japan" ] = "东京"
    countryCapitalMap [ "India " ] = "新德里"

    /*使用键输出地图值 */ 
    for country,_ := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }

    /*查看元素在集合中是否存在 */
    captial, ok := countryCapitalMap [ "美国" ] /*如果确定是真实的,则存在,否则不存在 */
    /*fmt.Println(captial) */
    /*fmt.Println(ok) */
    if (ok) {
        fmt.Println("美国的首都是", captial)
    } else {
        fmt.Println("美国的首都不存在")
    }
    var countryMap = map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
    fmt.Println(countryMap["India"])
}


