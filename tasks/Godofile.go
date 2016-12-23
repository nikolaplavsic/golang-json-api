package main

import (
	"fmt"
	"os"
	"runtime"

	. "gopkg.in/godo.v1"
)

var (
	appName = "jsonapi"
	root    = "./.."
)

func Tasks(p *Project) {
	Env = "GOPATH=$GOPATH"

	p.Task("default", D{"build"})

	p.Task("build", W{root + "/**/*.{go,ini}"}, func() {
		//set Mac and windows commands
		clear_cmd := "clear"
		run_cmd := root + string(os.PathSeparator) + appName
		beep_cmd := `echo -ne \007`

		if runtime.GOOS == "windows" {
			clear_cmd = "cls"
			run_cmd += ".exe"
			beep_cmd = "echo a_ | choice /c_ /n" // http://www.ericphelps.com/batch/samples/beep.txt
		}

		Run(clear_cmd)

		var err error
		err = Run("go generate", In{root})
		if err != nil {
			fmt.Printf("godo: generate error: %+v\n", err)
			Run(beep_cmd)
			return
		}

		fmt.Printf("godo: generate successful.\n")
		fmt.Printf("godo: " + appName + " is building …\n")

		err = Run("go build", In{root})
		if err != nil {
			fmt.Printf("godo: "+appName+" build error: %+v\n", err)
			Run(beep_cmd)
			return
		}

		fmt.Printf("godo: " + appName + " build successful.\n")

		// restart the program
		err = Start(run_cmd, In{root})
		if err != nil {
			fmt.Printf("godo: error starting "+appName+": %+v\n", err)
			Run(beep_cmd)
			return
		}

		fmt.Printf("godo: " + appName + " started.\n")
	})
}

func main() {
	Godo(Tasks)
}
