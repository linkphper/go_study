package action

import "fmt"

func JudgeExpressionIf(age int) {
	if age > 0 && age <= 18 {
		fmt.Println("少年")
		return
	}

	if age > 18 && age <= 35 {
		fmt.Println("中年")
		return
	}

	fmt.Println("老年")
}

func JudgeExpressionSwitchExpr(age int) {
	switch {
	case age > 0 && age <= 18:
		fmt.Println("少年")
		break
	case age > 18 && age <= 35:
		fmt.Println("中年")
		break
	default:
		fmt.Println("老年")
		return
	}
}

func JudgeExpressionSwitchEnum(week int) {
	weekday := "未知"

	switch week {
	case 1:
		weekday = "周一"
		break
	case 2:
		weekday = "周二"
		break
	case 3:
		weekday = "周三"
		break
	case 4:
		weekday = "周四"
		break
	case 5:
		weekday = "周五"
		break
	case 6:
		weekday = "周六"
		break
	case 7:
		weekday = "周日"
		break
	default:
		weekday = "错误"
		break
	}

	fmt.Println(weekday)
}
