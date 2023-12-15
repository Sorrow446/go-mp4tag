package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/Sorrow446/go-mp4tag"
	"github.com/alexflint/go-arg"
)

func parseArgs() *Args {
	var args Args
	arg.MustParse(&args)
	return &args
}

func parseTags(args *Args) (*mp4tag.Tags, error) {
	_tags := &mp4tag.Tags{
		Album:       args.Album,
		AlbumArtist: args.AlbumArtist,
		Artist:      args.Artist,
		Comment:     args.Comment,
		Composer:    args.Composer,
		Copyright:   args.Copyright,
		Custom:      args.Custom,
		Delete:      args.Delete,
		DiskNumber:  args.DiskNumber,
		DiskTotal:   args.DiskTotal,
		Genre:       args.Genre,
		Label:       args.Label,
		Title:       args.Title,
		TrackNumber: args.TrackNumber,
		TrackTotal:  args.TrackTotal,
		Year:        args.Year,
	}
	if args.Cover != "" {
		coverBytes, err := ioutil.ReadFile(args.Cover)
		if err != nil {
			return nil, err
		}
		_tags.Cover = coverBytes
	}
	return _tags, nil
}

// Icky, but don't think there's any other better way.
func tagsEmpty(_tags *mp4tag.Tags) bool {
	emptyOne := &mp4tag.Tags{DiskNumber: 0, DiskTotal: 0, TrackNumber: 0, TrackTotal: 0}
	emptyTwo := &mp4tag.Tags{DiskNumber: 0, DiskTotal: 0, TrackNumber: 0, TrackTotal: 0, Custom: map[string]string{}}
	return reflect.DeepEqual(_tags, emptyOne) || reflect.DeepEqual(_tags, emptyTwo)
}

func main() {
	args := parseArgs()
	_tags, err := parseTags(args)
	if err != nil {
		panic("Failed to parse tags.\n" + err.Error())
	}
	if tagsEmpty(_tags) {
		fmt.Println("Nothing to write. Exiting...")
		os.Exit(0)
	}
	if err = mp4tag.Write(args.FilePath, _tags); err != nil {
		panic("Failed to write tags.\n" + err.Error())
	}
}
