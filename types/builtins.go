package types

// TODO:
// 'builtin' does not mean 'external' always. For example, we might introduce polymorphic function
// `print: 'a -> ()` in the future as built-in function. We need to separate external symbols and
// built-in symbols for implementing special functions or intrinsics.

func builtinPopulatedTable() map[string]*External {
	return map[string]*External{
		"argv":                       &External{&Array{StringType}, "argv"},
		"print_int":                  &External{&Fun{UnitType, []Type{IntType}}, "print_int"},
		"print_bool":                 &External{&Fun{UnitType, []Type{BoolType}}, "print_bool"},
		"print_float":                &External{&Fun{UnitType, []Type{FloatType}}, "print_float"},
		"print_str":                  &External{&Fun{UnitType, []Type{StringType}}, "print_str"},
		"println_int":                &External{&Fun{UnitType, []Type{IntType}}, "println_int"},
		"println_bool":               &External{&Fun{UnitType, []Type{BoolType}}, "println_bool"},
		"println_float":              &External{&Fun{UnitType, []Type{FloatType}}, "println_float"},
		"println_str":                &External{&Fun{UnitType, []Type{StringType}}, "println_str"},
		"float_to_int":               &External{&Fun{IntType, []Type{FloatType}}, "float_to_int"},
		"int_to_float":               &External{&Fun{FloatType, []Type{IntType}}, "int_to_float"},
		"str_length":                 &External{&Fun{IntType, []Type{StringType}}, "str_length"},
		"__str_equal$builtin":        &External{&Fun{BoolType, []Type{StringType, StringType}}, "__str_equal"},
		"str_concat":                 &External{&Fun{StringType, []Type{StringType, StringType}}, "str_concat"},
		"str_sub":                    &External{&Fun{StringType, []Type{StringType, IntType, IntType}}, "str_sub"},
		"int_to_str":                 &External{&Fun{StringType, []Type{IntType}}, "int_to_str"},
		"float_to_str":               &External{&Fun{StringType, []Type{FloatType}}, "float_to_str"},
		"str_to_int":                 &External{&Fun{IntType, []Type{StringType}}, "str_to_int"},
		"str_to_float":               &External{&Fun{FloatType, []Type{StringType}}, "str_to_float"},
		"get_line":                   &External{&Fun{StringType, []Type{UnitType}}, "get_line"},
		"get_char":                   &External{&Fun{StringType, []Type{UnitType}}, "get_char"},
		"to_char_code":               &External{&Fun{IntType, []Type{StringType}}, "to_char_code"},
		"from_char_code":             &External{&Fun{StringType, []Type{IntType}}, "from_char_code"},
		"bit_and":                    &External{&Fun{IntType, []Type{IntType, IntType}}, "bit_and"},
		"bit_or":                     &External{&Fun{IntType, []Type{IntType, IntType}}, "bit_or"},
		"bit_xor":                    &External{&Fun{IntType, []Type{IntType, IntType}}, "bit_xor"},
		"bit_rsft":                   &External{&Fun{IntType, []Type{IntType, IntType}}, "bit_rsft"},
		"bit_lsft":                   &External{&Fun{IntType, []Type{IntType, IntType}}, "bit_lsft"},
		"bit_inv":                    &External{&Fun{IntType, []Type{IntType}}, "bit_inv"},
		"time_now":                   &External{&Fun{IntType, []Type{UnitType}}, "time_now"},
		"read_file":                  &External{&Fun{&Option{StringType}, []Type{StringType}}, "read_file"},
		"write_file":                 &External{&Fun{BoolType, []Type{StringType, StringType}}, "write_file"},
		"do_garbage_collection":      &External{&Fun{UnitType, []Type{UnitType}}, "do_garbage_collection"},
		"enable_garbage_collection":  &External{&Fun{UnitType, []Type{UnitType}}, "enable_garbage_collection"},
		"disable_garbage_collection": &External{&Fun{UnitType, []Type{UnitType}}, "disable_garbage_collection"},
	}
}
