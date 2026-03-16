package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func stringToMap(s string) map[string]string {
	result := make(map[string]string)
	pairs := strings.Split(s, ",")
	for _, pair := range pairs {
		parts := strings.Split(pair, ":")
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			result[key] = value
		}
	}
	return result
}

func main() {
	//routerconf := map[string]string{"_[0-9]$|_[1-9][0-9]$|_[1][0-9][0-9]$|_[2][0-4][0-9]$": "GS_DRIVERIDINDEX_YS_4", "_[2][5-9][0-9]$|_[3-4][0-9][0-9]$": "GS_DRIVERIDINDEX_YS_3", "_[5-6][0-9][0-9]$|_[7][0-4][0-9]$": "GS_DRIVERIDINDEX_YS_2", "default": "GS_DRIVERIDINDEX_YS_1"}
	//tbn_list := []string{"driverid_0", "driverid_2", "driverid_281", "driverid_530", "driverid_600", "driverid_800", "driverid_303", "driverid_108", "driverid_108_0", "driverid_480_0", "driverid_0_480"}

	routerconf := map[string]string{"g_comment_complaint_0.*": "GS_API_YS_2", "^g_user_imei$|^g_question$|^g_log$|^g_driver_online_count$|^g_cancel_order$": "GS_API_YS_4", "^g_question_answer$|^g_driver_daily$": "GS_API_YS_4", "^g_driver_online$|^g_transfer$|^g_evaluation_statistics$|^g_activity_driver$|^g_badaccount$": "GS_API_YS_3", "default": "GS_API_YS_1"}
	tbn_list := []string{"g_comment_complaint_010", "g_comment_complaint_0201", "g_user_imei", "g_question", "g_question_answer", "g_driver_online", "g_activity_driver", "g_badaccount", "g_badaccountsadd3"}
	for _, tablename := range tbn_list {
		ismatch := 0
		var tabletname string
		for router, tablet := range routerconf {

			re, err := regexp.Compile(router)
			if err != nil {
				fmt.Println("Error compiling regex:", err)
				return
			}
			if re.MatchString(tablename) {
				ismatch = 1
				tabletname = tablet
			}

		}
		if ismatch == 0 {
			tabletname = routerconf["default"]
		}
		fmt.Println(ismatch, "tablename:", tablename, "tablet:", tabletname)
	}
	//fmt.Println("4 4 3 2 2 1 3 4 4 4 3")
	fmt.Println("2 2 4 4 4 3 3 3 1")
	fmt.Println(time.Now().Format("20060102"))
	fmt.Printf("当天时间：%v", time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location()))

	strTime := "2024-11-07"
	t, err := time.Parse("2006-01-02", strTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)

	getminmaxCom := "get_minmaxpk nil,table has no data"
	getminmax := "FATAL: 2025/01/08 15:50:00 _get_minmaxpk nil,table has no data"
	getminmax2 := "FATAL: 2025/01/08 15:50:00 _get_minmax no data"

	re, err := regexp.Compile(getminmaxCom)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	if re.MatchString(getminmax) {
		fmt.Println("match getminmax")
	}
	if re.MatchString(getminmax2) {
		fmt.Println("match getminmax2")
	} else {
		fmt.Println("not match getminmax2")
	}

}
