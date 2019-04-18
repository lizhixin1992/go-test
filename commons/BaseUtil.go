package commons

import "testproject/models/conditions"

func SetLimitSize(condition *conditions.UserCondition) *conditions.UserCondition {
	if condition.Page > 0 {
		condition.Page = condition.Page - 1
	} else {
		condition.Page = 0
	}
	condition.StartRow = condition.Page * condition.Size
	condition.EndRow = condition.Size
	return condition
}