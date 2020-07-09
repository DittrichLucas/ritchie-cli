package tree

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/ZupIT/ritchie-cli/pkg/formula"
	"github.com/ZupIT/ritchie-cli/pkg/formula/repo"
	"github.com/ZupIT/ritchie-cli/pkg/stream"
)

func TestGenerate(t *testing.T) {
	fileManager := stream.NewFileManager()
	dirManager := stream.NewDirManager(fileManager)
	generator := NewGenerator(dirManager, fileManager)


	adder := repo.NewAdder(os.TempDir(), http.DefaultClient, generator, dirManager, fileManager)
	adder.Add(formula.Repo{
		Name:     "commons",
		ZipUrl:   "http://localhost:8882/repos/ZupIT/ritchie-formulas/zipball/v2.0.0",
		Version:  "v2.0.0",
		Priority: 0,
	})
	
	tree, err := generator.Generate(os.TempDir() + "/commons")
	if err != nil {
		t.Error(err)
	}

	bytes, _ := json.MarshalIndent(tree, "", "\t")

	fmt.Println(string(bytes))
}