package main

import(
	"fmt"
	"os"
	"text/template"
)

type Inventory struct {
    Material string
    Count    uint
}

func main() {
	// sweaters := Inventory{"wool", 17}
	// tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	// tmp2, err1 := template.New("test").Parse("{{.Count}} ite are made of {{.Material}}")
	// tmp3, err2 := template.New("test").Parse("{{.Count}} i are made of {{.Material}}")
	// if err != nil { panic(err) }
	// err = tmpl.Execute(os.Stdout, sweaters)
	// err1 = tmp2.Execute(os.Stdout, sweaters)
	// err2 = tmp3.Execute(os.Stdout, sweaters)
	// if err != nil { panic(err) }
	// if err1 != nil { panic(err1) }
	// if err2 != nil { panic(err2) }
	t := template.Must(template.New("letter").Parse("A{{.}}\n"))
    t1 := template.Must(template.New("letter1").Parse("B{{.}}\n"))
    t.Execute(os.Stdout, "1")
    t1.Execute(os.Stdout, "2")
    fmt.Println(len(t1.Templates()))
    for _, tt := range t1.Templates() {
        fmt.Println(tt.Name())
    }
}