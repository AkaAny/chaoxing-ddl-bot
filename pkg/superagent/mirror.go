package superagent

import (
	"net/http"
	"net/http/cookiejar"
	"reflect"
	"sync"
	"time"
	"unsafe"
	_ "unsafe"
)

type entryMap map[string]entry

type mirrorJar struct {
	psList cookiejar.PublicSuffixList

	// mu locks the remaining fields.
	mu sync.Mutex

	// entries is a set of entries, keyed by their eTLD+1 and subkeyed by
	// their name/domain/path.
	entries map[string]entryMap

	// nextSeqNum is the next sequence number assigned to a new cookie
	// created SetCookies.
	nextSeqNum uint64
}

type entry struct {
	Name       string
	Value      string
	Domain     string
	Path       string
	SameSite   string
	Secure     bool
	HttpOnly   bool
	Persistent bool
	HostOnly   bool
	Expires    time.Time
	Creation   time.Time
	LastAccess time.Time

	// seqNum is a sequence number so that Cookies returns cookies in a
	// deterministic order, even for cookies that have equal Path length and
	// equal Creation time. This simplifies testing.
	seqNum uint64
}

func toMirror(src http.CookieJar) *mirrorJar {
	srcVal := reflect.ValueOf(src)
	var srcPointer = unsafe.Pointer(srcVal.Pointer())
	var srcMirror = (*mirrorJar)(srcPointer)
	return srcMirror
}

func copyJar(src http.CookieJar, dst http.CookieJar) {
	var srcMirror = toMirror(src)
	var dstMirror = toMirror(dst)
	dstMirror.psList = srcMirror.psList
	dstMirror.nextSeqNum = srcMirror.nextSeqNum
	dstMirror.entries = copyEntries(srcMirror.entries)
}

func copyEntries(src map[string]entryMap) map[string]entryMap {
	var dst = make(map[string]entryMap)
	var srcVal = reflect.ValueOf(src)
	for _, key := range srcVal.MapKeys() {
		var em = src[key.String()]
		dst[key.String()] = copyEntryMap(em)
	}
	return dst
}

func copyEntryMap(src entryMap) entryMap {
	var dst = make(entryMap)
	var srcVal = reflect.ValueOf(src)
	for _, key := range srcVal.MapKeys() {
		var e = src[key.String()]
		dst[key.String()] = copyEntryManually(e)
	}
	return dst
}

func copyEntry(src entry) entry {
	var dst entry
	var srcVal = reflect.ValueOf(&src)
	var dstVal = reflect.ValueOf(&dst)
	//for fieldIndex := 0; fieldIndex < srcVal.NumField(); fieldIndex++ {
	//	var fieldValue= srcVal.Field(fieldIndex)
	//	dstVal.Field(fieldIndex).Set(fieldValue)
	//}
	memcpy(reflect.TypeOf(src), unsafe.Pointer(dstVal.Pointer()), unsafe.Pointer(srcVal.Pointer()))
	return dst
}

func copyEntryManually(src entry) entry {
	var dst = entry{
		Name:       src.Name,
		Value:      src.Value,
		Domain:     src.Domain,
		Path:       src.Path,
		SameSite:   src.SameSite,
		Secure:     src.Secure,
		HttpOnly:   src.HttpOnly,
		Persistent: src.Persistent,
		HostOnly:   src.HostOnly,
		Expires:    src.Expires,
		Creation:   src.Creation,
		LastAccess: src.LastAccess,
		seqNum:     src.seqNum,
	}
	return dst
}
