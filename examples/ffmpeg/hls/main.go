package main

import (
	"context"

	"github.com/fxkt-tech/liv/ffmpeg"
	"github.com/fxkt-tech/liv/ffmpeg/codec"
	"github.com/fxkt-tech/liv/ffmpeg/filter"
	"github.com/fxkt-tech/liv/ffmpeg/input"
	"github.com/fxkt-tech/liv/ffmpeg/naming"
	"github.com/fxkt-tech/liv/ffmpeg/output"
)

func main() {
	var (
		ctx = context.Background()
		nm  = naming.New()

		input1         = input.WithSimple("in.mp4")
		hlsKeyInfoFile = "hls/file.keyinfo"

		scale1 = filter.Scale(nm.Gen(), -2, -2).Use(filter.SelectStream(0, filter.StreamVideo, true))

		outfolder = "video/"

		mainFile = outfolder + "m.m3u8"
		segFile  = outfolder + "m-%5d.ts"
	)

	ffmpeg.NewFFmpeg(
		ffmpeg.V(ffmpeg.LogLevelError),
		ffmpeg.Debug(true),
		// ffmpeg.Dry(true),
	).AddInput(
		input1,
	).AddFilter(
		scale1,
	).AddOutput(
		output.New(
			output.Map(scale1.Name(0)),
			output.Map("0:a?"),
			output.VideoCodec(codec.X264),
			output.AudioCodec(codec.Copy),
			output.File(mainFile),
			output.MovFlags("faststart"),
			output.HLSSegmentType("mpegts"),
			output.HLSFlags("independent_segments"),
			output.HLSPlaylistType("vod"),
			output.HLSTime(2),
			output.HLSKeyInfoFile(hlsKeyInfoFile), // 加密
			output.HLSSegmentFilename(segFile),
			output.Format(codec.HLS),
		),
	).Run(ctx)
}
