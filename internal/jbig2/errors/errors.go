//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package errors ;import (_ga "fmt";_gf "golang.org/x/xerrors";);type processError struct{_b string ;_c string ;_gc string ;_gcg error ;};func Wrap (err error ,processName ,message string )error {if _f ,_fa :=err .(*processError );_fa {_f ._b ="";};_a :=_cfc (message ,processName );
_a ._gcg =err ;return _a ;};func _cfc (_e ,_ce string )*processError {return &processError {_b :"\u005b\u0055\u006e\u0069\u0050\u0044\u0046\u005d",_gc :_e ,_c :_ce };};func Error (processName ,message string )error {return _cfc (message ,processName )};
func Errorf (processName ,message string ,arguments ...interface{})error {return _cfc (_ga .Sprintf (message ,arguments ...),processName );};func (_cf *processError )Unwrap ()error {return _cf ._gcg };func (_bc *processError )Error ()string {var _ge string ;
if _bc ._b !=""{_ge =_bc ._b ;};_ge +="\u0050r\u006f\u0063\u0065\u0073\u0073\u003a "+_bc ._c ;if _bc ._gc !=""{_ge +="\u0020\u004d\u0065\u0073\u0073\u0061\u0067\u0065\u003a\u0020"+_bc ._gc ;};if _bc ._gcg !=nil {_ge +="\u002e\u0020"+_bc ._gcg .Error ();
};return _ge ;};func Wrapf (err error ,processName ,message string ,arguments ...interface{})error {if _de ,_ag :=err .(*processError );_ag {_de ._b ="";};_gd :=_cfc (_ga .Sprintf (message ,arguments ...),processName );_gd ._gcg =err ;return _gd ;};var _ _gf .Wrapper =(*processError )(nil );
