package modules

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gookit/color"
)

const (
	__autor__      = "Ch4rum"
	__copyright__  = "Copyright © 2025 Ch4rum -> https://www.instagram.com/Ch4rum/"
	__version__    = "1.1"
	__maintainer__ = "Ch4rum"
)

func Banner() {
	fmt.Printf(color.Green.Render(`                                                    
                  :                                        
           .     t#,                                       
          ;W    ;##W.   j.                                 
         f#E   :#L:WE   EW,       GEEEEEEEL :;;;;;;;;;;;;;.
       .E#f   .KG  ,#D  E##j      ,;;L#K;;.  jWWWWWWWW###L 
      iWW;    EE    ;#f E###D.       t#E             ,W#f  
     L##Lffi f#.     t#iE#jG#W;      t#E            ,##f   
    tLLG##L  :#G     GK E#t t##f     t#E           i##j    
      ,W#i    ;#L   LW. E#t  :K#E:   t#E          i##t     
     j#E.      t#f f#:  E#KDDDD###i  t#E         t##t      
   .D#j         f#D#;   E#f,t#Wi,,,  t#E        t##i       
  ,WK,           G#t    E#t  ;#W:    t#E       j##;        
  EG.             t     DWi   ,KK:    fE      :##,         
  ,                                    :      ,W,          
                                              ::`))
	fmt.Printf(color.Blue.Render("v", __version__))
	fmt.Println("\n\t", color.Red.Render(__copyright__))
	fmt.Println()
}

func ShowVersionInfo() {
	fmt.Printf("Versión: %s\n", __version__)
	fmt.Printf("Autor: %s\n", __autor__)
	fmt.Printf("Copyright: %s\n", __copyright__)
	fmt.Println("Licencia: MIT\n")
}

func UseTypeAlgorit(namealgorithm string, data []int) ([]int, error) {
	alg := &SortAlg{numbers: data, count: len(data)}
	options := map[string]func() []int{
		"bubble":    alg.BubbleSort,
		"selection": alg.SelectionSort,
		"insertion": alg.InsertionSort,
		"mergesort": alg.MergeSort,
		"quicksort": alg.QuickSort,
	}
	if chosen, ok := options[namealgorithm]; ok {
		return chosen(), nil
	} else {
		return nil, fmt.Errorf("Invalid choice, please try again!")
	}
}

func MainManager() {
	arg := GetArgument()
	algorithm, path, err := arg.GetSelectedAlgorithm()
	texterr := color.New(color.FgRed, color.OpBold).Sprint("[x] ERROR:")
	if err != nil {
		fmt.Println(texterr, err)
		return
	}
	if arg.ShowVersion {
		ShowVersionInfo()
		return
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(texterr, err)
		return
	}
	file := Archives{absPath}
	numbers, err := file.OpenArchive()
	if err != nil {
		fmt.Println(texterr, err)
		return
	}
	Banner()
	start := time.Now()
	sortNumbers, err := UseTypeAlgorit(algorithm, numbers)
	if err != nil {
		fmt.Println(texterr, err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("[%v] Time elapsed > %v seconds.\n", color.Green.Render("+"), color.Blue.Render(elapsed.Seconds()))
	if err := file.ExportArchive(sortNumbers); err != nil {
		fmt.Println(texterr, err)
		return
	}
	fmt.Printf("[%v] File successfully exported -> %v\n", color.Green.Render("+"), color.Cyan.Render(absPath))
}
