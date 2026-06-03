// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ThStoreTerminal is the golang structure for table th_store_terminal.
type ThStoreTerminal struct {
	StoreId    int64 `json:"storeId"    orm:"store_id"    description:"门店ID"`
	TerminalId int64 `json:"terminalId" orm:"terminal_id" description:"终端ID"`
	PrintTimes uint  `json:"printTimes" orm:"print_times" description:"打印联数(次数)"`
}
