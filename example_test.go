package unionfs_test

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"strconv"

	"arhat.dev/unionfs"
)

var (
	//go:embed trie
	sampleEmbedFS embed.FS

	sampleOSFS = os.DirFS("/some/rootfs")
)

func ExampleUnionFS() {
	ufs := unionfs.New()
	ufs.Map("/go/arhat.dev/unionfs/trie", "trie", sampleEmbedFS)
	ufs.Map("/", "/some/rootfs", sampleOSFS)

	f, err := ufs.Open("/go/arhat.dev/unionfs/trie/trie.go")
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	info, err := f.Stat()
	if err != nil {
		return
	}
	println("FILE trie.go NAME:", info.Name())

	_, _ = io.ReadAll(f)

	entries, err := fs.ReadDir(ufs, "/go/arhat.dev/unionfs/trie")
	if err != nil {
		panic(err)
	}

	for i, ent := range entries {
		println("DIR trie/#"+strconv.Itoa(i)+" NAME:", ent.Name())
	}
}
