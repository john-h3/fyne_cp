package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

func main() {
	a := app.New()

	rds := NewRecords(5)

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("FyneCP")

		InitClipboard()
		var fresh func([]byte)
		fresh = func(content []byte) {
			ok := rds.Add(content)
			if ok {
				items := make([]*fyne.MenuItem, 0)
				for _, Bytes := range rds.Slice() {
					items = append(items, fyne.NewMenuItem(string(Bytes), func() {
						WriteText(Bytes)
						fresh(Bytes)
					}))
				}
				items = append(items, fyne.NewMenuItemSeparator())
				items = append(items, fyne.NewMenuItem("Quit", func() {
					a.Quit()
				}))
				m.Items = items
				m.Refresh()
			}
		}
		go WatchClipboard(fresh)

		desk.SetSystemTrayMenu(m)
	}

	fyne.CurrentApp().Run()
}
