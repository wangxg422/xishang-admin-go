package permission

import (
	"backend/global"
	sysModel "backend/model/system"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"gorm.io/plugin/dbresolver"
)

type CachedAdapter struct {
}

// LoadPolicy loads all policy rules from the storage.
func (m *CachedAdapter) LoadPolicy(model model.Model) error {
	var lines []sysModel.CasbinRule
	if err := global.DB.Order("ID").Find(&lines).Error; err != nil {
		return err
	}
	err := m.Preview(&lines, model)
	if err != nil {
		return err
	}
	for _, line := range lines {
		err := loadPolicyLine(line, model)
		if err != nil {
			return err
		}
	}

	return nil
}

// SavePolicy saves all policy rules to the storage.
func (m *CachedAdapter) SavePolicy(model model.Model) error {
	var err error
	tx := global.DB.Clauses(dbresolver.Write).Begin()

	if err != nil {
		tx.Rollback()
		return err
	}

	var lines []sysModel.CasbinRule
	flushEvery := 1000
	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			lines = append(lines, m.savePolicyLine(ptype, rule))
			if len(lines) > flushEvery {
				if err := tx.Create(&lines).Error; err != nil {
					tx.Rollback()
					return err
				}
				lines = nil
			}
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			lines = append(lines, m.savePolicyLine(ptype, rule))
			if len(lines) > flushEvery {
				if err := tx.Create(&lines).Error; err != nil {
					tx.Rollback()
					return err
				}
				lines = nil
			}
		}
	}
	if len(lines) > 0 {
		if err := tx.Create(&lines).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit().Error
	return err
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (m *CachedAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (m *CachedAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (m *CachedAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}

func loadPolicyLine(line sysModel.CasbinRule, model model.Model) error {
	var p = []string{line.Ptype,
		line.V0, line.V1, line.V2,
		line.V3, line.V4, line.V5}

	index := len(p) - 1
	for p[index] == "" {
		index--
	}
	index += 1
	p = p[:index]
	err := persist.LoadPolicyArray(p, model)
	if err != nil {
		return err
	}
	return nil
}

func (m *CachedAdapter) Preview(rules *[]sysModel.CasbinRule, model model.Model) error {
	j := 0
	for i, rule := range *rules {
		r := []string{rule.Ptype,
			rule.V0, rule.V1, rule.V2,
			rule.V3, rule.V4, rule.V5}
		index := len(r) - 1
		for r[index] == "" {
			index--
		}
		index += 1
		p := r[:index]
		key := p[0]
		sec := key[:1]
		ok, err := model.HasPolicyEx(sec, key, p[1:])
		if err != nil {
			return err
		}
		if ok {
			(*rules)[j], (*rules)[i] = rule, (*rules)[j]
			j++
		}
	}
	(*rules) = (*rules)[j:]
	return nil
}

func (m *CachedAdapter) savePolicyLine(ptype string, rule []string) sysModel.CasbinRule {
	line := &sysModel.CasbinRule{}

	line.Ptype = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}

	return *line
}
