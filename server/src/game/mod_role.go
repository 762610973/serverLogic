package game

type ModRole struct {
}

func (m *ModRole) IsHasRole(roleId int) bool {
	return true
}
func (m *ModRole) GetRoleLevel(roleId int) int {
	return 80
}
