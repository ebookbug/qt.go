package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qstringlist.h
// dst-file: /src/core/qstringlist.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QStringList)=1
type QStringList struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QStringList) lastIndexOf(args ...interface{}) () {
  // lastIndexOf(const class QRegularExpression &, int)
  // lastIndexOf(const class QRegExp &, int)
  // lastIndexOf(class QRegExp &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStringList11lastIndexOfERK18QRegularExpressioni
  case 1:
    // invoke: _ZNK11QStringList11lastIndexOfERK7QRegExpi
  case 2:
    // invoke: _ZNK11QStringList11lastIndexOfER7QRegExpi
  default:
    qtrt.ErrorResolve("QStringList", "lastIndexOf", args)
  }

}


func NewQStringList(args ...interface{}) QStringList {
  return QStringList{}
}


func (this *QStringList) indexOf(args ...interface{}) () {
  // indexOf(const class QRegExp &, int)
  // indexOf(class QRegExp &, int)
  // indexOf(const class QRegularExpression &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStringList7indexOfERK7QRegExpi
  case 1:
    // invoke: _ZNK11QStringList7indexOfER7QRegExpi
  case 2:
    // invoke: _ZNK11QStringList7indexOfERK18QRegularExpressioni
  default:
    qtrt.ErrorResolve("QStringList", "indexOf", args)
  }

}

// <= body block end
