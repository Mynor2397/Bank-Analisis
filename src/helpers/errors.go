package helpers

var (
	// IsOk es para decir que el usuario ha sido ingresado correctamente
	IsOk = "Cuenta creada satisfactoriamente"

	//ErrInsert es para informar del conflicto en el insert
	ErrInsert = "Conflicto al crear la cuenta"

	//ErrTransaccion es la advertencia de error en la transacción
	ErrTransaccion = "Error al realizar la transacción"

	//OkTransaccion es para informar el exito de la transacción
	OkTransaccion = "mssql: TRANSFERENCIA EXITOSA"

	//ErrNoData es porqque no hay ningun registro encontrado
	ErrNoData = "Ningun registro encontrado"

	//ErrorSaldoInsuficiente is error of user is incorrect.
	ErrorSaldoInsuficiente = "mssql: CANTIDAD INSUFICIENTE PARA REALIZAR EL DEBITO"

	// ErrorCuentaNotFound es el error de que no hay cuenta en el sistema
	ErrorCuentaNotFound = "mssql: NUMERO DE CUENTA INEXISTENTE"
)
