package ui

import "github.com/lxn/walk"

var cachedSystemIconsForWidthAndDllIdx = make(map[widthDllIdx]*walk.Icon)

func loadSysIcon(dll string, index int32, size int) (icon *walk.Icon) {
	icon = cachedSystemIconsForWidthAndDllIdx[widthDllIdx{size, index, dll}]
	if icon != nil {
		return
	}
	var err error
	icon, err = walk.NewIconFromSysDLLWithSize(dll, int(index), size)
	if err == nil {
		cachedSystemIconsForWidthAndDllIdx[widthDllIdx{size, index, dll}] = icon
	}
	return
}

type widthDllIdx struct {
	width int
	idx   int32
	dll   string
}

type widthAndState struct {
	width int
	state ServiceState
}

var cachedIconsForWidthAndState = make(map[widthAndState]*walk.Icon)

func iconForState(state ServiceState, size int) (icon *walk.Icon) {
	icon = cachedIconsForWidthAndState[widthAndState{size, state}]
	if icon != nil {
		return
	}
	switch state {
	case StateStarted:
		icon = loadSysIcon("imageres", 101, size)
	case StateStopped:
		icon = loadSysIcon("imageres", 100, size)
	default:
		icon = loadSysIcon("shell32", 238, size)
	}
	cachedIconsForWidthAndState[widthAndState{size, state}] = icon
	return
}

var cachedLogoIconsForWidth = make(map[int]*walk.Icon)

func loadLogoIcon(size int) (icon *walk.Icon, err error) {
	icon = cachedLogoIconsForWidth[size]
	if icon != nil {
		return
	}
	icon, err = walk.NewIconFromResourceIdWithSize(11, walk.Size{size, size})
	if err == nil {
		cachedLogoIconsForWidth[size] = icon
	}
	return
}
