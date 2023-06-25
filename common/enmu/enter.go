package enmu

type EnmuGroup struct {
	DelFlagNormal DelFlag
	DelFlagDelete DelFlag

	// UserStatusNormal UserStatus
	// UserStatusDelete UserStatus

	StatusNormal   Status
	StatusDisabled Status
}

var EnmuGroupApp = new(EnmuGroup)
