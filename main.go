package main

import (
	"os"

	"github.com/martinohmann/collections-gen/internal/generator"
	"k8s.io/gengo/args"

	"k8s.io/klog"
)

func main() {
	klog.InitFlags(nil)
	arguments := args.Default()

	// no boilerplate by default
	arguments.GoHeaderFilePath = ""

	if err := arguments.Execute(
		generator.NameSystems(),
		generator.DefaultNameSystem(),
		generator.Packages,
	); err != nil {
		klog.Errorf("Error: %v", err)
		os.Exit(1)
	}
	klog.V(2).Info("Completed successfully.")
}
