# go-kmz

Package kmz provides convenience methods for creating and writing KMZ files.

See https://godoc.org/github.com/twpayne/go-kmz.

Example:

```go
func ExampleNewKMZ() {
	kmz := NewKMZ(
		kml.Placemark(
			kml.Name("Simple placemark"),
			kml.Description("Attached to the ground. Intelligently places itself at the height of the underlying terrain."),
			kml.Point(
				kml.Coordinates(kml.Coordinate{Lon: -122.0822035425683, Lat: 37.42228990140251}),
			),
		),
	)
	w := &bytes.Buffer{}
	if err := kmz.WriteIndent(w, "", "\t"); err != nil {
		log.Fatal(err)
	}
}
```

[Licence](LICENCE)
