package main

import (
	"saltfish/internal/job"
)

func m() {
	manager := job.NewManager()
	//str := binding.NewString()
	//str.Set(fmt.Sprintf("Rewards : %d Fishing : %d", manager.GetCountReward(), manager.GetCountFish()))
	//workLabel := widget.NewLabelWithData(str)
	//
	//widget.NewButton("hide", func() {
	//	manager.HideMode()
	//}),
	//	widget.NewButton("show", func() {
	//		manager.ShowMode()
	//	}),
	//	widget.NewButton("reward", func() {
	//		go manager.SetCallBack(func() {
	//			str.Set(fmt.Sprintf("Rewards : %d Fishing : %d", manager.GetCountReward(), manager.GetCountFish()))
	//		}).GetReward()
	//	}),
	//	widget.NewButton("fishing", func() {
	//		go manager.SetCallBack(func() {
	//			str.Set(fmt.Sprintf("Rewards : %d Fishing : %d", manager.GetCountReward(), manager.GetCountFish()))
	//		}).GetFish()
	//	}),
	//
	//	menu := fyne.NewMenu("MyApp",
	//	fyne.NewMenuItem("open", func() {
	//		log.Println("Tapped show")
	//		w.Show()
	//	}),
	//	fyne.NewMenuItem("show", func() {
	//		manager.ShowMode()
	//	}),
	//	fyne.NewMenuItem("hide", func() {
	//		manager.HideMode()
	//	}),
	//	fyne.NewMenuItem("reward", func() {
	//		go manager.GetReward()
	//	}),
	//)

}
