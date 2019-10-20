package main

import (
	"os"

	"github.com/martinohmann/collections-gen/generators"
	"k8s.io/gengo/args"

	"k8s.io/klog"
)

func main() {
	klog.InitFlags(nil)
	arguments := args.Default()

	// no boilerplate by default
	arguments.GoHeaderFilePath = ""

	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		klog.Errorf("Error: %v", err)
		os.Exit(1)
	}
	klog.V(2).Info("Completed successfully.")
}
