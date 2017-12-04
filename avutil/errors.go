// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Corbatto (luca@corbatto.de)

package avutil

//#include <libavutil/error.h>
//#include "errorMacros.h"
import "C"
import "unsafe"

func AVERROR_BSF_NOT_FOUND() int      { return int(C.macro_AVERROR_BSF_NOT_FOUND) }
func AVERROR_BUG() int                { return int(C.macro_AVERROR_BUG) }
func AVERROR_BUFFER_TOO_SMALL() int   { return int(C.macro_AVERROR_BUFFER_TOO_SMALL) }
func AVERROR_DECODER_NOT_FOUND() int  { return int(C.macro_AVERROR_DECODER_NOT_FOUND) }
func AVERROR_DEMUXER_NOT_FOUND() int  { return int(C.macro_AVERROR_DEMUXER_NOT_FOUND) }
func AVERROR_ENCODER_NOT_FOUND() int  { return int(C.macro_AVERROR_ENCODER_NOT_FOUND) }
func AVERROR_EOF() int                { return int(C.macro_AVERROR_EOF) }
func AVERROR_EXIT() int               { return int(C.macro_AVERROR_EXIT) }
func AVERROR_EXTERNAL() int           { return int(C.macro_AVERROR_EXTERNAL) }
func AVERROR_FILTER_NOT_FOUND() int   { return int(C.macro_AVERROR_FILTER_NOT_FOUND) }
func AVERROR_INVALIDDATA() int        { return int(C.macro_AVERROR_INVALIDDATA) }
func AVERROR_MUXER_NOT_FOUND() int    { return int(C.macro_AVERROR_MUXER_NOT_FOUND) }
func AVERROR_OPTION_NOT_FOUND() int   { return int(C.macro_AVERROR_OPTION_NOT_FOUND) }
func AVERROR_PATCHWELCOME() int       { return int(C.macro_AVERROR_PATCHWELCOME) }
func AVERROR_PROTOCOL_NOT_FOUND() int { return int(C.macro_AVERROR_PROTOCOL_NOT_FOUND) }
func AVERROR_STREAM_NOT_FOUND() int   { return int(C.macro_AVERROR_STREAM_NOT_FOUND) }
func AVERROR_UNKNOWN() int            { return int(C.macro_AVERROR_UNKNOWN) }
func AVERROR_EXPERIMENTAL() int       { return int(C.macro_AVERROR_EXPERIMENTAL) }
func AVERROR_INPUT_CHANGED() int      { return int(C.macro_AVERROR_INPUT_CHANGED) }
func AVERROR_OUTPUT_CHANGED() int     { return int(C.macro_AVERROR_OUTPUT_CHANGED) }
func AVERROR_HTTP_BAD_REQUEST() int   { return int(C.macro_AVERROR_HTTP_BAD_REQUEST) }
func AVERROR_HTTP_UNAUTHORIZED() int  { return int(C.macro_AVERROR_HTTP_UNAUTHORIZED) }
func AVERROR_HTTP_FORBIDDEN() int     { return int(C.macro_AVERROR_HTTP_FORBIDDEN) }
func AVERROR_HTTP_NOT_FOUND() int     { return int(C.macro_AVERROR_HTTP_NOT_FOUND) }
func AVERROR_HTTP_OTHER_4XX() int     { return int(C.macro_AVERROR_HTTP_OTHER_4XX) }
func AVERROR_HTTP_SERVER_ERROR() int  { return int(C.macro_AVERROR_HTTP_SERVER_ERROR) }
func AV_ERROR_MAX_STRING_SIZE() int   { return int(C.macro_AV_ERROR_MAX_STRING_SIZE) }

func AVERROR_EAGAIN() int { return AvError(int(C.macro_EAGAIN)) }

// Strerror returns a descriptive string of the given return code.
//
// C-Function: av_strerror
func Strerror(errnum int) string {
	buf := make([]C.char, AV_ERROR_MAX_STRING_SIZE())
	if C.av_strerror(C.int(errnum), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf))) != 0 {
		return "UNKNOWN_ERROR"
	}
	return C.GoString((*C.char)(unsafe.Pointer(&buf[0])))
}

// AvError is a wrapper for the macro AVERROR().
//
// C-Macro: AVERROR
func AvError(code int) int {
	return int(C.macro_AVERROR(C.int(code)))
}
