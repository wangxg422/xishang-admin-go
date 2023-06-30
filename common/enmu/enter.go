package enmu

type EnmuGroup struct {
	DelFlagNormal DelFlag
	DelFlagDelete DelFlag

	// UserStatusNormal UserStatus
	// UserStatusDelete UserStatus

	StatusNormal   Status
	StatusDisabled Status

	MenuTypeM MenuType
	MenuTypeC MenuType
	MenuTypeF MenuType
}

var EnmuGroupApp = new(EnmuGroup)


var a int = 5
var DelFlagNormal1 DelFlag = DelFlag(1)
