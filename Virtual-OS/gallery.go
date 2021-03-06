package main

import(
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fmt"
    "io/ioutil"
    "log"
	"strings"  
	"fyne.io/fyne/v2/canvas"
)

func showGallery(w fyne.Window) {
	// a := app.New()
	// w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(800,600));
	root_src:="C:\\Users\\Unknown\\OneDrive\\Desktop";

	files, err := ioutil.ReadDir(root_src);

    if err != nil {
        log.Fatal(err)
    }

		
	tabs := container.NewAppTabs()

    for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir(){
			extension:=strings.Split(file.Name(),".")[1];
			if extension =="png" || extension == "jpeg"{
				image:=canvas.NewImageFromFile(root_src+"\\"+file.Name());
				image.FillMode = canvas.ImageFillOriginal;
				tabs.Append(container.NewTabItem(file.Name(),image));
			}
		}

    }
	



	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,tabs),);
	w.ShowAndRun()
}



