package log

//Level describe el nivel del error
type Level int

//Kind describe el tipo del error
type Kind int

//Code describe el código de error
type Code int

const (
	KindDatabase Kind = 1
	KindLogic    Kind = 2
	KindView     Kind = 3
)

const (
	LevelFatal   Level = 0
	LevelWarning Level = 1
	LevelInfo    Level = 2
)

const (
	CodeServer          Code = 500
	CodeNotFound        Code = 404
	CodePaymentDeclined Code = 402
	CodeMailNotSend     Code = 202
	CodeNoError         Code = 0
)

//Throwable describe cualquier error
type Throwable interface {
	Error() string
	Level() Level
	Kind() Kind
	Code() Code
}

//String devuelve una cadena con el nivel de error
func (l Level) String() string {
	switch l {
	case LevelFatal:
		return "Fatal"
	case LevelWarning:
		return "Warning"
	case LevelInfo:
		return "Info"
	default:
		return "Unknown"
	}
}

//Error implementa Throwable
type Error struct {
	Mensaje string
	Nivel   Level
	Tipo    Kind
	Codigo  Code
}

//NewError crea una instancia de Error
func NewError(err string, level Level, kind Kind, code Code) Error {
	e := Error{
		Mensaje: err,
		Nivel:   level,
		Tipo:    kind,
		Codigo:  code,
	}
	RegistrarError(e)
	return e
}

func (c Code) Int() int { return int(c) }

func (e Error) Error() string { return e.Mensaje }
func (e Error) Level() Level  { return e.Nivel }
func (e Error) Kind() Kind    { return e.Tipo }
func (e Error) Code() Code    { return e.Codigo }
