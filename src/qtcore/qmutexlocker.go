package qtcore

// /usr/include/qt/QtCore/qmutex.h
// #include <qmutex.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QMutexLocker struct {
	*qtrt.CObject
}

func (this *QMutexLocker) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMutexLocker) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQMutexLockerFromPointer(cthis unsafe.Pointer) *QMutexLocker {
	return &QMutexLocker{&qtrt.CObject{cthis}}
}
func (*QMutexLocker) NewFromPointer(cthis unsafe.Pointer) *QMutexLocker {
	return NewQMutexLockerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmutex.h:199
// index:0
// Public inline
// void QMutexLocker(class QBasicMutex *)
func NewQMutexLocker(m *QBasicMutex /*777 QBasicMutex **/) *QMutexLocker {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = m.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QMutexLockerC2EP11QBasicMutex", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMutexLockerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmutex.h:213
// index:0
// Public inline
// void ~QMutexLocker()
func DeleteQMutexLocker(*QMutexLocker) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QMutexLockerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:215
// index:0
// Public inline
// void unlock()
func (this *QMutexLocker) Unlock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QMutexLocker6unlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:223
// index:0
// Public inline
// void relock()
func (this *QMutexLocker) Relock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QMutexLocker6relockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:238
// index:0
// Public inline
// QMutex * mutex()
func (this *QMutexLocker) Mutex() *QMutex /*777 QMutex **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QMutexLocker5mutexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMutexFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
