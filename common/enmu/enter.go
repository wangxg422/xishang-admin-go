package enmu

const (
	DelFlagNormal  DelFlag = DelFlag(DelFlag_Normal)
	DelFlagDeleted DelFlag = DelFlag(DelFlag_Deleted)

	StatusNormal   Status = Status(Status_Normal)
	StatusDisabled Status = Status(Status_Disabled)

	MenuTypeM MenuType = MenuType(MenuType_M)
	MenuTypeC MenuType = MenuType(MenuType_C)
	MenuTypeF MenuType = MenuType(MenuType_F)

	MenuIsFrame    MenuFrame = MenuFrame(Menu_Frame)
	MenuIsNotFrame MenuFrame = MenuFrame(Menu_Not_Frame)

	MenuIsCache    MenuCache = MenuCache(Menu_Cache)
	MenuIsNotCache MenuCache = MenuCache(Menu_Not_Cache)

	MenuIsVisible  MenuVisible = MenuVisible(Menu_Visible)
	MenuIsUnisible MenuVisible = MenuVisible(Menu_Unvisible)
)
