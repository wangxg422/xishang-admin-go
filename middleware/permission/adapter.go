package permission

import (
	"github.com/casbin/casbin/v2/model"
)

type CachedAdapter struct {
}

// LoadPolicy loads all policy rules from the storage.
func (m *CachedAdapter) LoadPolicy(model model.Model) error {
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (m *CachedAdapter) SavePolicy(model model.Model) error {
	return nil
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
