package models

type Transaction struct {
	Sender   string
	Reciever string
	Sponsor  string
	Amount   string
}

//-------------------------
//--- Adena wallet response

// {
//     code: 0,
//     type: "SIGN_TX",
//     status: "success",
//     message: "Sign Transaction",
//     data: {
//         encodedTransaction: "CnMKDS9iYW5rLk1zZ1NlbmQSYgooZzFqZzhtdHV0dTlraGhmd2M0bnhtdWhjcGZ0ZjBwYWpkaGZ2c3FmNRIoZzFmZnp4aGE1N2RoMHFndjltYTV2MzkzdXIwemV4ZnZwNmxzanBhZRoMNTAwMDAwMHVnbm90EgwIgIl6EgYxdWdub3Qafgo6ChMvdG0uUHViS2V5U2VjcDI1NmsxEiMKIQPhYTbbFx4y30iZNZQfBW4i+Jhj43OdCrfNSexCg5ydshJAK3y9vuIO0BhY+P6f361/RP5QNFPCpHiaNE/cGhRyOV4MSDSl+++WN56NKefZF27MQTHu3lFDTZ7pPTTbkn7DeCIFMTIzMTM="
//     },
// }

type TxHashDecode struct {
	Caller string
	Amount string
}
