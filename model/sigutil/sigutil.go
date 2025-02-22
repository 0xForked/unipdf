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

package sigutil ;import (_a "bytes";_ca "crypto";_f "crypto/x509";_gf "encoding/asn1";_gd "encoding/pem";_cd "errors";_ef "fmt";_ab "github.com/unidoc/timestamp";_gdg "github.com/unidoc/unipdf/v3/common";_gfb "golang.org/x/crypto/ocsp";_ae "io";_e "io/ioutil";
_g "net/http";_b "time";);

// NewCRLClient returns a new CRL client.
func NewCRLClient ()*CRLClient {return &CRLClient {HTTPClient :_ddb ()}};

// MakeRequest makes a OCSP request to the specified server and returns
// the parsed and raw responses. If a server URL is not provided, it is
// extracted from the certificate.
func (_dc *OCSPClient )MakeRequest (serverURL string ,cert ,issuer *_f .Certificate )(*_gfb .Response ,[]byte ,error ){if _dc .HTTPClient ==nil {_dc .HTTPClient =_ddb ();};if serverURL ==""{if len (cert .OCSPServer )==0{return nil ,nil ,_cd .New ("\u0063e\u0072\u0074i\u0066\u0069\u0063a\u0074\u0065\u0020\u0064\u006f\u0065\u0073 \u006e\u006f\u0074\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0079\u0020\u004f\u0043S\u0050\u0020\u0073\u0065\u0072\u0076\u0065\u0072\u0073");
};serverURL =cert .OCSPServer [0];};_fe ,_cc :=_gfb .CreateRequest (cert ,issuer ,&_gfb .RequestOptions {Hash :_dc .Hash });if _cc !=nil {return nil ,nil ,_cc ;};_edf ,_cc :=_dc .HTTPClient .Post (serverURL ,"\u0061p\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u006fc\u0073\u0070\u002d\u0072\u0065\u0071\u0075\u0065\u0073\u0074",_a .NewReader (_fe ));
if _cc !=nil {return nil ,nil ,_cc ;};defer _edf .Body .Close ();_ad ,_cc :=_e .ReadAll (_edf .Body );if _cc !=nil {return nil ,nil ,_cc ;};if _cf ,_ :=_gd .Decode (_ad );_cf !=nil {_ad =_cf .Bytes ;};_dac ,_cc :=_gfb .ParseResponseForCert (_ad ,cert ,issuer );
if _cc !=nil {return nil ,nil ,_cc ;};return _dac ,_ad ,nil ;};

// GetIssuer retrieves the issuer of the provided certificate.
func (_bg *CertClient )GetIssuer (cert *_f .Certificate )(*_f .Certificate ,error ){for _ ,_ff :=range cert .IssuingCertificateURL {_ed ,_be :=_bg .Get (_ff );if _be !=nil {_gdg .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0064\u006f\u0077\u006e\u006c\u006f\u0061\u0064\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0066\u006f\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0025\u0076\u003a\u0020\u0025\u0076",cert .Subject .CommonName ,_be );
continue ;};return _ed ,nil ;};return nil ,_ef .Errorf ("\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063a\u0074e\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064");};

// MakeRequest makes a CRL request to the specified server and returns the
// response. If a server URL is not provided, it is extracted from the certificate.
func (_ecb *CRLClient )MakeRequest (serverURL string ,cert *_f .Certificate )([]byte ,error ){if _ecb .HTTPClient ==nil {_ecb .HTTPClient =_ddb ();};if serverURL ==""{if len (cert .CRLDistributionPoints )==0{return nil ,_cd .New ("\u0063e\u0072\u0074i\u0066\u0069\u0063\u0061t\u0065\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070ec\u0069\u0066\u0079 \u0061\u006ey\u0020\u0043\u0052\u004c\u0020\u0073e\u0072\u0076e\u0072\u0073");
};serverURL =cert .CRLDistributionPoints [0];};_ea ,_cda :=_ecb .HTTPClient .Get (serverURL );if _cda !=nil {return nil ,_cda ;};defer _ea .Body .Close ();_fdb ,_cda :=_e .ReadAll (_ea .Body );if _cda !=nil {return nil ,_cda ;};if _eae ,_ :=_gd .Decode (_fdb );
_eae !=nil {_fdb =_eae .Bytes ;};return _fdb ,nil ;};

// IsCA returns true if the provided certificate appears to be a CA certificate.
func (_fd *CertClient )IsCA (cert *_f .Certificate )bool {return cert .IsCA &&_a .Equal (cert .RawIssuer ,cert .RawSubject );};

// TimestampClient represents a RFC 3161 timestamp client.
// It is used to obtain signed tokens from timestamp authority servers.
type TimestampClient struct{

// HTTPClient is the HTTP client used to make timestamp requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_g .Client ;

// Callbacks.
BeforeHTTPRequest func (_ee *_g .Request )error ;};

// Get retrieves the certificate at the specified URL.
func (_bc *CertClient )Get (url string )(*_f .Certificate ,error ){if _bc .HTTPClient ==nil {_bc .HTTPClient =_ddb ();};_ge ,_d :=_bc .HTTPClient .Get (url );if _d !=nil {return nil ,_d ;};defer _ge .Body .Close ();_fg ,_d :=_e .ReadAll (_ge .Body );if _d !=nil {return nil ,_d ;
};if _efd ,_ :=_gd .Decode (_fg );_efd !=nil {_fg =_efd .Bytes ;};_ec ,_d :=_f .ParseCertificate (_fg );if _d !=nil {return nil ,_d ;};return _ec ,nil ;};

// OCSPClient represents a OCSP (Online Certificate Status Protocol) client.
// It is used to request revocation data from OCSP servers.
type OCSPClient struct{

// HTTPClient is the HTTP client used to make OCSP requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_g .Client ;

// Hash is the hash function  used when constructing the OCSP
// requests. If zero, SHA-1 will be used.
Hash _ca .Hash ;};

// NewTimestampRequest returns a new timestamp request based
// on the specified options.
func NewTimestampRequest (body _ae .Reader ,opts *_ab .RequestOptions )(*_ab .Request ,error ){if opts ==nil {opts =&_ab .RequestOptions {};};if opts .Hash ==0{opts .Hash =_ca .SHA256 ;};if !opts .Hash .Available (){return nil ,_f .ErrUnsupportedAlgorithm ;
};_abf :=opts .Hash .New ();if _ ,_bgc :=_ae .Copy (_abf ,body );_bgc !=nil {return nil ,_bgc ;};return &_ab .Request {HashAlgorithm :opts .Hash ,HashedMessage :_abf .Sum (nil ),Certificates :opts .Certificates ,TSAPolicyOID :opts .TSAPolicyOID ,Nonce :opts .Nonce },nil ;
};

// GetEncodedToken executes the timestamp request and returns the DER encoded
// timestamp token bytes.
func (_eg *TimestampClient )GetEncodedToken (serverURL string ,req *_ab .Request )([]byte ,error ){if serverURL ==""{return nil ,_ef .Errorf ("\u006d\u0075\u0073\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061m\u0070\u0020\u0073\u0065\u0072\u0076\u0065r\u0020\u0055\u0052\u004c");
};if req ==nil {return nil ,_ef .Errorf ("\u0074\u0069\u006de\u0073\u0074\u0061\u006dp\u0020\u0072\u0065\u0071\u0075\u0065\u0073t\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c");};_gef ,_ddf :=req .Marshal ();if _ddf !=nil {return nil ,_ddf ;
};_cdd ,_ddf :=_g .NewRequest ("\u0050\u004f\u0053\u0054",serverURL ,_a .NewBuffer (_gef ));if _ddf !=nil {return nil ,_ddf ;};_cdd .Header .Set ("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065","a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079");
if _eg .BeforeHTTPRequest !=nil {if _add :=_eg .BeforeHTTPRequest (_cdd );_add !=nil {return nil ,_add ;};};_bf :=_eg .HTTPClient ;if _bf ==nil {_bf =_ddb ();};_db ,_ddf :=_bf .Do (_cdd );if _ddf !=nil {return nil ,_ddf ;};defer _db .Body .Close ();_ag ,_ddf :=_e .ReadAll (_db .Body );
if _ddf !=nil {return nil ,_ddf ;};if _db .StatusCode !=_g .StatusOK {return nil ,_ef .Errorf ("\u0075\u006e\u0065x\u0070\u0065\u0063\u0074e\u0064\u0020\u0048\u0054\u0054\u0050\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0064",_db .StatusCode );
};var _cb struct{Version _gf .RawValue ;Content _gf .RawValue ;};if _ ,_ddf =_gf .Unmarshal (_ag ,&_cb );_ddf !=nil {return nil ,_ddf ;};return _cb .Content .FullBytes ,nil ;};

// NewTimestampClient returns a new timestamp client.
func NewTimestampClient ()*TimestampClient {return &TimestampClient {HTTPClient :_ddb ()}};

// CRLClient represents a CRL (Certificate revocation list) client.
// It is used to request revocation data from CRL servers.
type CRLClient struct{

// HTTPClient is the HTTP client used to make CRL requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_g .Client ;};

// NewOCSPClient returns a new OCSP client.
func NewOCSPClient ()*OCSPClient {return &OCSPClient {HTTPClient :_ddb (),Hash :_ca .SHA1 }};

// NewCertClient returns a new certificate client.
func NewCertClient ()*CertClient {return &CertClient {HTTPClient :_ddb ()}};func _ddb ()*_g .Client {return &_g .Client {Timeout :5*_b .Second }};

// CertClient represents a X.509 certificate client. Its primary purpose
// is to download certificates.
type CertClient struct{

// HTTPClient is the HTTP client used to make certificate requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_g .Client ;};