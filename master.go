package m3u8_parser

import (
  "strings"
  "encoding/csv"
  "strconv"
  "bytes"
"fmt"
)

func (h HLSMaster) Type() string {
  return "master"
}

func (h HLSMaster) insertAt(pos int, playlist HLSPlaylister)  {
  // check if we need throw exception here for master
  return
}

func (h HLSMaster) deleteAt(pos int, length int) {
  return
}

func (h HLSMaster) replaceAt(pos int, playlist HLSPlaylister) {
  return
}

func  (h *HLSMaster)  Parse (m3u8Lines *[]string, segmentCount int) {
  h.header = []string{"#EXTM3U"}
  h.lines = make([]HLSSegmenter, segmentCount)
  curLineNo := 0
  seg := HLSMasterSegment{}
  seg.sequence = 0

  for _, line := range (*m3u8Lines)[1:] {
	
    if strings.HasPrefix(line, "#EXT-X-STREAM-INF") {
      // parsing stream info here.
      start := len("#EXT-X-STREAM-INF:")
      reader := csv.NewReader(bytes.NewBufferString(line[start:]))
      stream_fields, _ := reader.Read()
      for _, field := range stream_fields {
        sep := strings.Index(field, "=")
        if (sep > 0) {
          switch strings.ToUpper(strings.TrimSpace(field[0:sep])) {
            case "BANDWIDTH": seg.bandwidth, _= strconv.Atoi(field[sep+1:])
            case "PROGRAM-ID": seg.programId = strings.TrimSpace(field[sep+1:])
            case "CODECS": seg.codecs = strings.TrimSpace(field[sep+1:])
            case "RESOLUTION": seg.resolution = strings.TrimSpace(field[sep+1:])
            case "AUDIO": seg.audio = strings.TrimSpace(field[sep+1:])
            case "VIDEO": seg.video = strings.TrimSpace(field[sep+1:])
          }
        }
      }
    } else if strings.HasPrefix(line, "#") {
      seg.tags = append(seg.tags, line)
    } else if line == "" {
      // empty line, not doing anything here.
    } else {
      seg.url = line
      h.lines[curLineNo] = seg

      seg = HLSMasterSegment{}
      curLineNo++
      seg.sequence = curLineNo
    }
  }

  return
}