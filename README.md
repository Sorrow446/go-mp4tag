# go-mp4tag
MP4 tag library written in Go.

### Setup
```
go get github.com/Sorrow446/go-mp4tag
```

### Usage Examples
```go
import "github.com/Sorrow446/go-mp4tag"
```
Opening is omitted from the examples.
```go
mp4, err := mp4tag.Open("1.m4a")
if err != nil {
	panic(err)
}
defer mp4.Close()
```

Read album title:
```go
tags, err := mp4.Read()
if err != nil {
	panic(err)
}
fmt.Println(tags.Album)
```

Extract all covers:
```go
tags, err := mp4.Read()
if err != nil {
	panic(err)
}

for idx, pic := range tags.Pictures {
	fname := fmt.Sprintf("out_%03d.jpg", idx+1)
	err = os.WriteFile(fname, pic.Data, 0666)
	if err != nil {
		panic(err)
	}
}
```

Write two covers:
```go
	picOneData, err := os.ReadFile("1.jpg")
	if err != nil {
		panic(err)
	}

	picTwoData, err := os.ReadFile("2.jpg")
	if err != nil {
		panic(err)
	}

	picOne := &mp4tag.MP4Picture{Data: picOneData}
	picTwo := &mp4tag.MP4Picture{Data: picTwoData}

	tags := &mp4tag.MP4Tags{
		Pictures: []*mp4tag.MP4Picture{picOne, picTwo},
	}

	err = mp4.Write(tags, []string{})
	if err != nil {
		panic(err)
	}
```


Write track number and total:
```go
tags := &mp4tag.MP4Tags{
	TrackNumber: 1,
	TrackTotal: 10,
}

err = mp4.Write(tags, []string{})
if err != nil {
	panic(err)
}
```

Delete comment:
```go
err = mp4.Write(&mp4tag.MP4Tags{}, []string{"comment"})
if err != nil {
	panic(err)
}
```
