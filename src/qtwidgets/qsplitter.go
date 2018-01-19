//  header block begin
// /usr/include/qt/QtWidgets/qsplitter.h
// #include <qsplitter.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QSplitter struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qsplitter.h:59
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSplitter) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:67
// index:0
// void QSplitter(class QWidget *)
func NewQSplitter(parent unsafe.Pointer) *QSplitter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitterC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QSplitter{cthis}
}

// /usr/include/qt/QtWidgets/qsplitter.h:68
// index:1
// void QSplitter(Qt::Orientation, class QWidget *)
func NewQSplitter_1(arg0 int, parent unsafe.Pointer) *QSplitter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitterC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &arg0, parent)
	gopp.ErrPrint(err, rv)
	return &QSplitter{cthis}
}

// /usr/include/qt/QtWidgets/qsplitter.h:69
// index:0
// virtual
// void ~QSplitter()
func DeleteQSplitter(*QSplitter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:71
// index:0
// void addWidget(class QWidget *)
func (this *QSplitter) AddWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter9addWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:72
// index:0
// void insertWidget(int, class QWidget *)
func (this *QSplitter) InsertWidget(index int, widget unsafe.Pointer) {
	// 0: (, int index, QWidget * widget), (&index, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter12insertWidgetEiP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &index, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:73
// index:0
// QWidget * replaceWidget(int, class QWidget *)
func (this *QSplitter) ReplaceWidget(index int, widget unsafe.Pointer) {
	// 0: (, int index, QWidget * widget), (&index, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter13replaceWidgetEiP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &index, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:75
// index:0
// void setOrientation(Qt::Orientation)
func (this *QSplitter) SetOrientation(arg0 int) {
	// 0: (, Qt::Orientation arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:76
// index:0
// Qt::Orientation orientation()
func (this *QSplitter) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter11orientationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:78
// index:0
// void setChildrenCollapsible(_Bool)
func (this *QSplitter) SetChildrenCollapsible(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter22setChildrenCollapsibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:79
// index:0
// bool childrenCollapsible()
func (this *QSplitter) ChildrenCollapsible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter19childrenCollapsibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:81
// index:0
// void setCollapsible(int, _Bool)
func (this *QSplitter) SetCollapsible(index int, arg1 bool) {
	// 0: (, int index, bool arg1), (&index, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter14setCollapsibleEib", ffiqt.FFI_TYPE_VOID, this.cthis, &index, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:82
// index:0
// bool isCollapsible(int)
func (this *QSplitter) IsCollapsible(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter13isCollapsibleEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:83
// index:0
// void setOpaqueResize(_Bool)
func (this *QSplitter) SetOpaqueResize(opaque bool) {
	// 0: (, bool opaque), (&opaque)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter15setOpaqueResizeEb", ffiqt.FFI_TYPE_VOID, this.cthis, &opaque)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:84
// index:0
// bool opaqueResize()
func (this *QSplitter) OpaqueResize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter12opaqueResizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:85
// index:0
// void refresh()
func (this *QSplitter) Refresh() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter7refreshEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:87
// index:0
// virtual
// QSize sizeHint()
func (this *QSplitter) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:88
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QSplitter) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:90
// index:0
// QList<int> sizes()
func (this *QSplitter) Sizes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter5sizesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:93
// index:0
// QByteArray saveState()
func (this *QSplitter) SaveState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter9saveStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:94
// index:0
// bool restoreState(const class QByteArray &)
func (this *QSplitter) RestoreState(state unsafe.Pointer) {
	// 0: (, const QByteArray & state), (state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter12restoreStateERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:96
// index:0
// int handleWidth()
func (this *QSplitter) HandleWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter11handleWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:97
// index:0
// void setHandleWidth(int)
func (this *QSplitter) SetHandleWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter14setHandleWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:99
// index:0
// int indexOf(class QWidget *)
func (this *QSplitter) IndexOf(w unsafe.Pointer) {
	// 0: (, QWidget * w), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter7indexOfEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:100
// index:0
// QWidget * widget(int)
func (this *QSplitter) Widget(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter6widgetEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:101
// index:0
// int count()
func (this *QSplitter) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:103
// index:0
// void getRange(int, int *, int *)
func (this *QSplitter) GetRange(index int, arg1 unsafe.Pointer, arg2 unsafe.Pointer) {
	// 0: (, int index, int * arg1, int * arg2), (&index, arg1, arg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter8getRangeEiPiS0_", ffiqt.FFI_TYPE_VOID, this.cthis, &index, arg1, arg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:104
// index:0
// QSplitterHandle * handle(int)
func (this *QSplitter) Handle(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter6handleEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:106
// index:0
// void setStretchFactor(int, int)
func (this *QSplitter) SetStretchFactor(index int, stretch int) {
	// 0: (, int index, int stretch), (&index, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter16setStretchFactorEii", ffiqt.FFI_TYPE_VOID, this.cthis, &index, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:109
// index:0
// void splitterMoved(int, int)
func (this *QSplitter) SplitterMoved(pos int, index int) {
	// 0: (, int pos, int index), (&pos, &index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter13splitterMovedEii", ffiqt.FFI_TYPE_VOID, this.cthis, &pos, &index)
	gopp.ErrPrint(err, rv)
}

//  body block end