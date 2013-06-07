package m3u8_parser

import (
  "log"
  "testing"
  "strconv"
  "reflect"
  
)

const SIMPLE_MASTER string = `
    #EXTM3U
    #TEST
    #EXT-X-STREAM-INF:BANDWIDTH=500000,PROGRAM-ID=1
    http://www.example.com/
    #EXT-X-STREAM-INF:BANDWIDTH=800000,PROGRAM-ID=1
    http://www.someotherexample.com/
  `

func TestM3U8Parser_header(t *testing.T) {
  hls := M3U8Parser(SIMPLE_MASTER)
  header := *hls.getHeader()
  lines := *hls.getLines()
  verify(t, "header size", "M3U8Parser", "SIMPLE_MASTER", "1", strconv.Itoa(len(header)))
  verify(t, "header match", "M3U8Parser", "SIMPLE_MASTER", "#EXTM3U", header[0])
  verify(t, "header size", "M3U8Parser", "SIMPLE_MASTER", "2", strconv.Itoa(len(lines)))
  l := reflect.ValueOf(lines[0]).Interface().(HLSMasterSegment)

  verify(t, "first sequence", "M3U8Parser", "SIMPLE_MASTER", "0", strconv.Itoa(l.sequence))
  verify(t, "first sequence", "M3U8Parser", "SIMPLE_MASTER", "http://www.example.com/", l.url)
  verify(t, "first sequence", "M3U8Parser", "SIMPLE_MASTER", "#TEST", l.tags[0])
  verify(t, "bandwidth", "M3U8Parser", "SIMPLE_MASTER", "500000", strconv.Itoa(l.bandwidth))

  line1 := reflect.ValueOf(lines[1]).Interface().(HLSMasterSegment)

  verify(t, "first sequence", "M3U8Parser", "SIMPLE_MASTER", "1", strconv.Itoa(line1.sequence))
  verify(t, "prog id", "M3U8Parser", "SIMPLE_MASTER", "1", line1.programId)
  verify(t, "bandwidth", "M3U8Parser", "SIMPLE_MASTER", "800000", strconv.Itoa(line1.bandwidth))
}


func BenchmarkM3U8Parser(b *testing.B) {
    for i := 0; i < b.N; i++ {
      hls := M3U8Parser(SIMPLE_MASTER)
      log.Println(hls.String())
    }
}