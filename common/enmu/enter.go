package enmu

const (
	DelFlagNormal  DelFlag = DelFlag(DelFlag_Normal)
	DelFlagDeleted DelFlag = DelFlag(DelFlag_Deleted)

	StatusNormal   Status = Status(Status_Normal)
	StatusDisabled Status = Status(Status_Disabled)

	MenuTypeDir  MenuType = MenuType(MenuType_DIR)
	MenuTypeMenu MenuType = MenuType(MenuType_MENU)
	MenuTypeBTN  MenuType = MenuType(MenuType_BTN)

	MenuIsFrame    MenuFrame = MenuFrame(Menu_Frame)
	MenuIsNotFrame MenuFrame = MenuFrame(Menu_Not_Frame)

	MenuIsCache    MenuCache = MenuCache(Menu_Cache)
	MenuIsNotCache MenuCache = MenuCache(Menu_Not_Cache)

	SexM = Sex(Sex_M)
	SexF = Sex(Sex_F)
	SexU = Sex(Sex_U)
)
