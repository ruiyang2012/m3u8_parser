package m3u8_parser

import (
  "strings"
)

func (h HLSSegment) String() string {
  return strings.Join(h.tags, "\n") + "\n" + h.url
}

func (h HLSMediaSegment) String() string {
  return strings.Join(h.tags, "\n") + "\n" + h.url
}

func (h HLSManifest) String() string {
  return strings.Join(h.header, "\n") + "\n" + strings.Join(mapStr(h.lines), "\n")
}

func (h HLSManifest) getHeader() *[]string {
  return &h.header
}

func (h HLSManifest) getLines() *[]HLSSegmenter {
  return &h.lines
}

func M3U8Parser(m3u8content string) HLSPlaylister {
  m3u8Lines := strings.Split(strings.TrimSpace(m3u8content), "\n")
  segmentCount := 0
  isMaster := false
  isSub := false
  for i, value := range m3u8Lines {
    m3u8Lines[i] = strings.TrimSpace(value)
    if m3u8Lines[i] != "" && !strings.HasPrefix(m3u8Lines[i], "#") { segmentCount++ }
    if strings.HasPrefix(m3u8Lines[i], "#EXT-X-STREAM-INF") { isMaster = true }
    if strings.HasPrefix(m3u8Lines[i], "#EXTINF") { isSub = true }
  }
  if m3u8Lines[0] != "#EXTM3U" { return nil }
  if isMaster { return M3U8MasterParser(&m3u8Lines, segmentCount) }
  if isSub { return M3U8SubParser(&m3u8Lines, segmentCount) }
  return nil
}

func M3U8MasterParser(m3u8Lines *[]string, segmentCount int) *HLSMaster {
  hlsMaster := HLSMaster{}
  hlsMaster.Parse(m3u8Lines, segmentCount)
  return &hlsMaster
}

func M3U8SubParser(m3u8Lines *[]string, segmentCount int) HLSPlaylister {
  // 
  return nil
}