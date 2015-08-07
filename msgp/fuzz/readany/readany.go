package readany

import "github.com/tinylib/msgp/msgp"

func Fuzz(data []byte) int {
	var err error
	knowntype := true
	tp := msgp.NextType(data)
	switch tp {
	case msgp.InvalidType:
		return 0
	case msgp.NilType:
		_, err = msgp.ReadNilBytes(data)
	case msgp.BoolType:
		_, _, err = msgp.ReadBoolBytes(data)
	case msgp.BinType:
		_, _, err = msgp.ReadBytesZC(data)
	case msgp.StrType:
		_, _, err = msgp.ReadStringZC(data)
	case msgp.MapType:
		_, _, err = msgp.ReadMapHeaderBytes(data)
	case msgp.ArrayType:
		_, _, err = msgp.ReadArrayHeaderBytes(data)
	case msgp.UintType:
		_, _, err = msgp.ReadUint64Bytes(data)
	case msgp.IntType:
		_, _, err = msgp.ReadInt64Bytes(data)
	case msgp.Float32Type:
		_, _, err = msgp.ReadFloat32Bytes(data)
	case msgp.Float64Type:
		_, _, err = msgp.ReadFloat64Bytes(data)
	case msgp.ExtensionType:
		var m msgp.RawExtension
		_, err = msgp.ReadExtensionBytes(data, &m)
	case msgp.TimeType:
		knowntype = false
		_, _, err = msgp.ReadTimeBytes(data)
	case msgp.Complex64Type:
		knowntype = false
		_, _, err = msgp.ReadComplex64Bytes(data)
	case msgp.Complex128Type:
		knowntype = false
		_, _, err = msgp.ReadComplex128Bytes(data)
	default:
		panic("impossible type")
	}
	if err != nil {
		// for all but the extension types (time.Time, complex64/128),
		// we shouldn't get a type error.
		if te, ok := err.(msgp.TypeError); ok && knowntype {
			panic(tp.String() + ": " + te.Error())
		}
		return 0
	}
	return 1
}
