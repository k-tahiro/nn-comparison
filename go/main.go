package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/k-tahiro/tahiremogon/middleware"
	"github.com/k-tahiro/tahiremogon/util"
)

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func main() {
	predictionModel, err := middleware.LoadPredictionModel(os.Getenv("ONNX_MODEL_FILE"))
	if err != nil {
		panic(err)
	}

	input_dir := os.Getenv("INPUT_DIR")
	paths := dirwalk(input_dir)
	for _, path := range paths {
		input, err := util.ReadImage(path)
		if err != nil {
			panic(err)
		}

		label, err := predictionModel.Predict(input)
		if err != nil {
			panic(err)
		}

		fmt.Println(label)
	}

}
