package htable

import "fmt"

type Group struct {
	Tables map[string]*Table
	Rules  map[string]RuleHandler
}

func NewGroup() *Group {
	return &Group{
		Tables: map[string]*Table{},
		Rules:  map[string]RuleHandler{},
	}
}

func (g *Group) AddTable(table *Table) {
	if g.Tables[table.Name] != nil {
		panic(fmt.Sprintf("table %s already existing", table.Name))
	}
	g.Tables[table.Name] = table
}

func (g *Group) AddRule(name string, handler RuleHandler) {
	if g.Rules[name] != nil {
		panic(fmt.Sprintf("role %s already existing", name))
	}
	g.Rules[name] = handler
}

// RuleCheck 规则检查
// name 规则名称
// src 数据来源，例如 [a,b] 表示 a 表的 b 字段
// des 匹配目标，例如 [c,d] 表示匹配到 c 表的 d 字段，可以为空
func (g *Group) RuleCheck(name string, src []string, des ...string) {

}
