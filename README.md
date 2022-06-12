# go-mp4tag
MP4 tagger written in Go.

## Setup
```
go get github.com/Sorrow446/go-mp4tag
```
```go
import "github.com/Sorrow446/go-mp4tag"
```

## Usage Examples
```go
	tags := &mp4tag.Tags{
		Album: "album",
		AlbumArtist: "album artist",
		Title:       "title",
		TrackNumber: 1,
		TrackTotal:  20,
		Genre: "genre",
		DiskNumber:  3,
		DiskTotal:   10,
		Comment:     "comment",
	}
	err := mp4tag.Write("1.m4a", tags)
	if err != nil {
		panic(err)
	}
```
 Write album, album artist, title, track number, track total, genre, disk number, disk total, and comment tags.
 
 
 ```go
	tags := &mp4tag.Tags{
		Custom: map[string]string{
			"CUSTOMFIELD": "custom field",
			"CUSTOMFIELD2": "custom field 2",
		},
		Delete: []string{"genre", "cover"},
	}
	err := mp4tag.Write("1.m4a", tags)
	if err != nil {
		panic(err)
	}
```
Write two custom fields named `CUSTOMFIELD` and `CUSTOMFIELD2`, delete genre tag, and remove cover.


```go
	coverBytes, err := ioutil.ReadFile("cover.jpg")
	if err != nil {
		panic(err)
	}
	tags := &mp4tag.Tags{
		Cover: coverBytes,
	}
	err = mp4tag.Write("1.m4a", tags)
	if err != nil {
		panic(err)
	}
```
Write cover from `cover.jpg`.

## Thank you
This library relies heavily on abema's go-mp4 library.
