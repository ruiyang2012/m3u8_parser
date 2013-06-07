package m3u8_parser

import "fmt"
// Interface in go following -er ended.

type HLSPlaylister interface {
  fmt.Stringer
  Type() string
  insertAt(pos int, playlist HLSPlaylister)
  replaceAt(pos int, playlist HLSPlaylister)
  deleteAt(pos int, length int)
  getHeader()* []string
  getLines()* []HLSSegmenter
}

type HLSSegmenter interface {

}


type HLSSegment struct {
  sequence int
  url string
  tags [] string
}

type HLSKey struct {
  method string
  uri string
  iv string
  keyFormat string
  keyFormatVersion string
}

type HLSMasterSegment struct {
  HLSSegment
  bandwidth int
  programId string
  codecs string
  resolution string
  audio string
  video string
}

type HLSMediaSegment struct {
  HLSSegment
  key string
  showKey bool
  duration int
  discontinuity bool
  cueOutDuration int
  cueIn bool
  cueOut bool
}

type HLSManifest struct {
  header []string
  lines [] HLSSegmenter
}

type HLSMaster struct {
  HLSManifest
}

type HLSEvent struct {
  HLSManifest
}

type HLSVod struct {
  HLSManifest
}

type HLSSliding struct {
  HLSManifest
}