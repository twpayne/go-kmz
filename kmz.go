// Package kmz provides convenience methods for creating and writing KMZ files.
package kmz

import (
	"archive/zip"
	"io"

	"github.com/twpayne/go-kml"
)

type KMZ struct {
	roots []kml.Element
	files map[string][]byte
}

func NewKMZ(roots ...kml.Element) *KMZ {
	return &KMZ{
		roots: roots,
		files: make(map[string][]byte),
	}
}

func (kmz *KMZ) AddRoot(root kml.Element) *KMZ {
	kmz.roots = append(kmz.roots, root)
	return kmz
}

func (kmz *KMZ) AddFile(filename string, content []byte) *KMZ {
	kmz.files[filename] = content
	return kmz
}

func (kmz *KMZ) Write(w io.Writer) error {
	zw := zip.NewWriter(w)
	f, err := zw.Create("doc.kml")
	if err != nil {
		return err
	}
	if err := kml.GxKML(kml.Document(kmz.roots...)).Write(f); err != nil {
		return err
	}
	for filename, content := range kmz.files {
		f, err := zw.Create(filename)
		if err != nil {
			return err
		}
		if _, err := f.Write(content); err != nil {
			return err
		}
	}
	if err := zw.Close(); err != nil {
		return err
	}
	return nil
}
