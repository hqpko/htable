package htable

type RuleHandler func(group *Group) (pass bool)

// HandlerRuleUnique 字段唯一
func HandlerRuleUnique(group *Group) bool {
	return false
}

// HandlerRuleLink 属性关联
func HandlerRuleLink(group *Group) bool {
	return false
}

// HandlerRuleHas 持有属性
func HandlerRuleHas(group *Group) bool {
	return false
}
