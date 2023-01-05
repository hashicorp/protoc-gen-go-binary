package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/hashicorp/protoc-gen-go-binary/gen"
	"google.golang.org/protobuf/compiler/protogen"
	plugin "google.golang.org/protobuf/types/pluginpb"
)

var (
	importPrefix = flag.String("import_prefix", "", "prefix to be added to go package paths for imported proto files")
	file         = flag.String("file", "-", "where to load data from")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gp *protogen.Plugin) error {

		gp.SupportedFeatures = uint64(plugin.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		for _, name := range gp.Request.FileToGenerate {
			f := gp.FilesByPath[name]

			if len(f.Messages) == 0 {
				glog.V(1).Infof("Skipping %s, no messages", name)
				continue
			}

			glog.V(1).Infof("Processing %s", name)
			glog.V(2).Infof("Generating %s\n", fmt.Sprintf("%s.pb.binary.go", f.GeneratedFilenamePrefix))

			gf := gp.NewGeneratedFile(fmt.Sprintf("%s.pb.binary.go", f.GeneratedFilenamePrefix), f.GoImportPath)

			err := gen.ApplyTemplate(gf, f)
			if err != nil {
				gf.Skip()
				gp.Error(err)
				continue
			}
		}

		return nil
	})
}
