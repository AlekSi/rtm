// +build ignore

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/AlekSi/rtm"
)

func main() {
	log.SetFlags(0)

	keyF := flag.String("key", "", "API key")
	secretF := flag.String("secret", "", "API secret")
	flag.Parse()
	if *keyF == "" || *secretF == "" {
		log.Fatal("Both -key and -secret flags should be used.")
	}

	client := &rtm.Client{
		APIKey:    *keyF,
		APISecret: *secretF,
	}
	methods, err := client.Reflection().GetMethods(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	clientT := reflect.TypeOf(client)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	for _, method := range methods {
		docMethod := "TODO"
		parts := strings.SplitN(method, ".", 3)
		serviceName := strings.ToUpper(parts[1][:1]) + parts[1][1:]
		m, ok := clientT.MethodByName(serviceName)
		if ok {
			methodName := strings.ToUpper(parts[2][:1]) + parts[2][1:]
			m, ok = m.Type.Out(0).MethodByName(methodName)
			if ok {
				serviceName += "Service"
				docMethod = fmt.Sprintf("[`%[1]s.%[2]s`](https://godoc.org/github.com/AlekSi/rtm#%[1]s.%[2]s)", serviceName, methodName)
			}
		}

		fmt.Fprintf(w, "| [`%[1]s`](https://www.rememberthemilk.com/services/api/methods/%[1]s.rtm)\t %s\n", method, docMethod)
	}
	w.Flush()
}
