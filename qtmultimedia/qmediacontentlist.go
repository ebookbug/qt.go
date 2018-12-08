package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediacontent.h
// #include <qmediacontent.h>
// #include <QtMultimedia>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end

//  body block begin
type QMediaContentList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QMediaContentList) Operator_equal0() *QMediaContentList {
	// QMediaContentList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QMediaContentList) Operator_equal1() *QMediaContentList {
	// QMediaContentList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QMediaContentList) Swap0() {
	// QMediaContentList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QMediaContentList) Operator_equal_equal0() bool {
	// QMediaContentList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QMediaContentList) Operator_not_equal0() bool {
	// QMediaContentList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QMediaContentList) Size0() int {
	// QMediaContentList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QMediaContentList) Detach0() {
	// QMediaContentList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QMediaContentList) DetachShared0() {
	// QMediaContentList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QMediaContentList) IsDetached0() bool {
	// QMediaContentList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QMediaContentList) SetSharable0() {
	// QMediaContentList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QMediaContentList) IsSharedWith0() bool {
	// QMediaContentList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QMediaContentList) IsEmpty0() bool {
	// QMediaContentList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QMediaContentList) Clear0() {
	// QMediaContentList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QMediaContentList) At0() *QMediaContent {
	// QMediaContentList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// const T & operator[](int)
func (this *QMediaContentList) Operator_get_index0() *QMediaContent {
	// QMediaContentList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// T & operator[](int)
func (this *QMediaContentList) Operator_get_index1() *QMediaContent {
	// QMediaContentList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// void reserve(int)
func (this *QMediaContentList) Reserve0() {
	// QMediaContentList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QMediaContentList) Append0() {
	// QMediaContentList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QMediaContentList) Append1() {
	// QMediaContentList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QMediaContentList) Prepend0() {
	// QMediaContentList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QMediaContentList) Insert0() {
	// QMediaContentList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QMediaContentList) Replace0() {
	// QMediaContentList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QMediaContentList) RemoveAt0() {
	// QMediaContentList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QMediaContentList) RemoveAll0() int {
	// QMediaContentList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QMediaContentList) RemoveOne0() bool {
	// QMediaContentList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QMediaContentList) TakeAt0() *QMediaContent {
	// QMediaContentList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// T takeFirst()
func (this *QMediaContentList) TakeFirst0() *QMediaContent {
	// QMediaContentList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// T takeLast()
func (this *QMediaContentList) TakeLast0() *QMediaContent {
	// QMediaContentList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// void move(int, int)
func (this *QMediaContentList) Move0() {
	// QMediaContentList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QMediaContentList) Swap1() {
	// QMediaContentList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QMediaContentList) IndexOf0() int {
	// QMediaContentList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QMediaContentList) LastIndexOf0() int {
	// QMediaContentList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QMediaContentList) Contains0() bool {
	// QMediaContentList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QMediaContentList) Count0() int {
	// QMediaContentList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QMediaContentList) Begin0() {
	// QMediaContentList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QMediaContentList) Begin1() {
	// QMediaContentList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QMediaContentList) Cbegin0() {
	// QMediaContentList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QMediaContentList) ConstBegin0() {
	// QMediaContentList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QMediaContentList) End0() {
	// QMediaContentList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QMediaContentList) End1() {
	// QMediaContentList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QMediaContentList) Cend0() {
	// QMediaContentList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QMediaContentList) ConstEnd0() {
	// QMediaContentList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QMediaContentList) Rbegin0() {
	// QMediaContentList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QMediaContentList) Rend0() {
	// QMediaContentList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QMediaContentList) Rbegin1() {
	// QMediaContentList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QMediaContentList) Rend1() {
	// QMediaContentList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QMediaContentList) Crbegin0() {
	// QMediaContentList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QMediaContentList) Crend0() {
	// QMediaContentList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QMediaContentList) Insert1() {
	// QMediaContentList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QMediaContentList) Erase0() {
	// QMediaContentList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QMediaContentList) Erase1() {
	// QMediaContentList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QMediaContentList) Count1() int {
	// QMediaContentList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QMediaContentList) Length0() int {
	// QMediaContentList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QMediaContentList) First0() *QMediaContent {
	// QMediaContentList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// const T & constFirst()
func (this *QMediaContentList) ConstFirst0() *QMediaContent {
	// QMediaContentList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// const T & first()
func (this *QMediaContentList) First1() *QMediaContent {
	// QMediaContentList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// T & last()
func (this *QMediaContentList) Last0() *QMediaContent {
	// QMediaContentList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// const T & last()
func (this *QMediaContentList) Last1() *QMediaContent {
	// QMediaContentList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// const T & constLast()
func (this *QMediaContentList) ConstLast0() *QMediaContent {
	// QMediaContentList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// void removeFirst()
func (this *QMediaContentList) RemoveFirst0() {
	// QMediaContentList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QMediaContentList) RemoveLast0() {
	// QMediaContentList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QMediaContentList) StartsWith0() bool {
	// QMediaContentList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QMediaContentList) EndsWith0() bool {
	// QMediaContentList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QMediaContentList) Mid0() *QMediaContentList {
	// QMediaContentList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QMediaContentList) Value0() *QMediaContent {
	// QMediaContentList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// T value(int, const T &)
func (this *QMediaContentList) Value1() *QMediaContent {
	// QMediaContentList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// void push_back(const T &)
func (this *QMediaContentList) Push_back0() {
	// QMediaContentList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QMediaContentList) Push_front0() {
	// QMediaContentList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QMediaContentList) Front0() *QMediaContent {
	// QMediaContentList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// const T & front()
func (this *QMediaContentList) Front1() *QMediaContent {
	// QMediaContentList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// T & back()
func (this *QMediaContentList) Back0() *QMediaContent {
	// QMediaContentList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// const T & back()
func (this *QMediaContentList) Back1() *QMediaContent {
	// QMediaContentList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaContent{}
}

// void pop_front()
func (this *QMediaContentList) Pop_front0() {
	// QMediaContentList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QMediaContentList) Pop_back0() {
	// QMediaContentList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QMediaContentList) Empty0() bool {
	// QMediaContentList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QMediaContentList) Operator_add_equal0() *QMediaContentList {
	// QMediaContentList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QMediaContentList) Operator_add0() *QMediaContentList {
	// QMediaContentList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QMediaContentList) Operator_add_equal1() *QMediaContentList {
	// QMediaContentList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QMediaContentList) Operator_left_shift0() *QMediaContentList {
	// QMediaContentList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QMediaContentList) Operator_left_shift1() *QMediaContentList {
	// QMediaContentList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QMediaContentList) ToVector0() {
	// QMediaContentList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QMediaContentList) ToSet0() {
	// QMediaContentList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QMediaContentList) FromVector0() *QMediaContentList {
	// QMediaContentList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QMediaContentList) FromSet0() *QMediaContentList {
	// QMediaContentList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QMediaContentList) FromStdList0() *QMediaContentList {
	// QMediaContentList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QMediaContentList) ToStdList0() {
	// QMediaContentList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QMediaContentList) Detach_helper_grow0() {
	// QMediaContentList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QMediaContentList) Detach_helper0() {
	// QMediaContentList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QMediaContentList) Detach_helper1() {
	// QMediaContentList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QMediaContentList) Dealloc0() {
	// QMediaContentList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QMediaContentList) Node_construct0() {
	// QMediaContentList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QMediaContentList) Node_destruct0() {
	// QMediaContentList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QMediaContentList) Node_copy0() {
	// QMediaContentList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QMediaContentList) Node_destruct1() {
	// QMediaContentList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QMediaContentList) IsValidIterator0() bool {
	// QMediaContentList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QMediaContentList) Op_eq_impl0() bool {
	// QMediaContentList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QMediaContentList) Op_eq_impl1() bool {
	// QMediaContentList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QMediaContentList) Contains_impl0() bool {
	// QMediaContentList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QMediaContentList) Contains_impl1() bool {
	// QMediaContentList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QMediaContentList) Count_impl0() int {
	// QMediaContentList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QMediaContentList) Count_impl1() int {
	// QMediaContentList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaContentList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
