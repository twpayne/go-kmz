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

func (kmz *KMZ) write(w io.Writer, writeRoots func(io.Writer, []kml.Element) error) error {
	zw := zip.NewWriter(w)
	rootW, err := zw.Create("doc.kml")
	if err != nil {
		return err
	}
	if err := writeRoots(rootW, kmz.roots); err != nil {
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
	return zw.Close()
}

func (kmz *KMZ) Write(w io.Writer) error {
	return kmz.write(w, func(w io.Writer, roots []kml.Element) error {
		return kml.GxKML(kml.Document(kmz.roots...)).Write(w)
	})
}

func (kmz *KMZ) WriteIndent(w io.Writer, prefix, indent string) error {
	return kmz.write(w, func(w io.Writer, roots []kml.Element) error {
		return kml.GxKML(kml.Document(kmz.roots...)).WriteIndent(w, prefix, indent)
	})
}
