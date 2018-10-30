package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QmlBridge_ITF interface {
	std_core.QObject_ITF
	QmlBridge_PTR() *QmlBridge
}

func (ptr *QmlBridge) QmlBridge_PTR() *QmlBridge {
	return ptr
}

func (ptr *QmlBridge) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlBridge) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlBridge(ptr QmlBridge_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlBridge_PTR().Pointer()
	}
	return nil
}

func NewQmlBridgeFromPointer(ptr unsafe.Pointer) (n *QmlBridge) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QmlBridge)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QmlBridge:
			n = deduced

		case *std_core.QObject:
			n = &QmlBridge{QObject: *deduced}

		default:
			n = new(QmlBridge)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackQmlBridge40a792_Constructor
func callbackQmlBridge40a792_Constructor(ptr unsafe.Pointer) {
	this := NewQmlBridgeFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackQmlBridge40a792_DisplayTotalBalance
func callbackQmlBridge40a792_DisplayTotalBalance(ptr unsafe.Pointer, balance C.struct_Moc_PackedString, balanceUSD C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayTotalBalance"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(balance), cGoUnpackString(balanceUSD))
	}

}

func (ptr *QmlBridge) ConnectDisplayTotalBalance(f func(balance string, balanceUSD string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayTotalBalance") {
			C.QmlBridge40a792_ConnectDisplayTotalBalance(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayTotalBalance"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayTotalBalance", func(balance string, balanceUSD string) {
				signal.(func(string, string))(balance, balanceUSD)
				f(balance, balanceUSD)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayTotalBalance", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayTotalBalance() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayTotalBalance(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayTotalBalance")
	}
}

func (ptr *QmlBridge) DisplayTotalBalance(balance string, balanceUSD string) {
	if ptr.Pointer() != nil {
		var balanceC *C.char
		if balance != "" {
			balanceC = C.CString(balance)
			defer C.free(unsafe.Pointer(balanceC))
		}
		var balanceUSDC *C.char
		if balanceUSD != "" {
			balanceUSDC = C.CString(balanceUSD)
			defer C.free(unsafe.Pointer(balanceUSDC))
		}
		C.QmlBridge40a792_DisplayTotalBalance(ptr.Pointer(), C.struct_Moc_PackedString{data: balanceC, len: C.longlong(len(balance))}, C.struct_Moc_PackedString{data: balanceUSDC, len: C.longlong(len(balanceUSD))})
	}
}

//export callbackQmlBridge40a792_DisplayAvailableBalance
func callbackQmlBridge40a792_DisplayAvailableBalance(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayAvailableBalance"); signal != nil {
		signal.(func(string))(cGoUnpackString(data))
	}

}

func (ptr *QmlBridge) ConnectDisplayAvailableBalance(f func(data string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayAvailableBalance") {
			C.QmlBridge40a792_ConnectDisplayAvailableBalance(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayAvailableBalance"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayAvailableBalance", func(data string) {
				signal.(func(string))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayAvailableBalance", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayAvailableBalance() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayAvailableBalance(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayAvailableBalance")
	}
}

func (ptr *QmlBridge) DisplayAvailableBalance(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.QmlBridge40a792_DisplayAvailableBalance(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackQmlBridge40a792_DisplayLockedBalance
func callbackQmlBridge40a792_DisplayLockedBalance(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayLockedBalance"); signal != nil {
		signal.(func(string))(cGoUnpackString(data))
	}

}

func (ptr *QmlBridge) ConnectDisplayLockedBalance(f func(data string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayLockedBalance") {
			C.QmlBridge40a792_ConnectDisplayLockedBalance(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayLockedBalance"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayLockedBalance", func(data string) {
				signal.(func(string))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayLockedBalance", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayLockedBalance() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayLockedBalance(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayLockedBalance")
	}
}

func (ptr *QmlBridge) DisplayLockedBalance(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.QmlBridge40a792_DisplayLockedBalance(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackQmlBridge40a792_DisplayAddress
func callbackQmlBridge40a792_DisplayAddress(ptr unsafe.Pointer, address C.struct_Moc_PackedString, wallet C.struct_Moc_PackedString, displayFiatConversion C.char) {
	if signal := qt.GetSignal(ptr, "displayAddress"); signal != nil {
		signal.(func(string, string, bool))(cGoUnpackString(address), cGoUnpackString(wallet), int8(displayFiatConversion) != 0)
	}

}

func (ptr *QmlBridge) ConnectDisplayAddress(f func(address string, wallet string, displayFiatConversion bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayAddress") {
			C.QmlBridge40a792_ConnectDisplayAddress(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayAddress"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayAddress", func(address string, wallet string, displayFiatConversion bool) {
				signal.(func(string, string, bool))(address, wallet, displayFiatConversion)
				f(address, wallet, displayFiatConversion)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayAddress", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayAddress() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayAddress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayAddress")
	}
}

func (ptr *QmlBridge) DisplayAddress(address string, wallet string, displayFiatConversion bool) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		var walletC *C.char
		if wallet != "" {
			walletC = C.CString(wallet)
			defer C.free(unsafe.Pointer(walletC))
		}
		C.QmlBridge40a792_DisplayAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))}, C.struct_Moc_PackedString{data: walletC, len: C.longlong(len(wallet))}, C.char(int8(qt.GoBoolToInt(displayFiatConversion))))
	}
}

//export callbackQmlBridge40a792_AddTransactionToList
func callbackQmlBridge40a792_AddTransactionToList(ptr unsafe.Pointer, paymentID C.struct_Moc_PackedString, transactionID C.struct_Moc_PackedString, amount C.struct_Moc_PackedString, confirmations C.struct_Moc_PackedString, time C.struct_Moc_PackedString, number C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addTransactionToList"); signal != nil {
		signal.(func(string, string, string, string, string, string))(cGoUnpackString(paymentID), cGoUnpackString(transactionID), cGoUnpackString(amount), cGoUnpackString(confirmations), cGoUnpackString(time), cGoUnpackString(number))
	}

}

func (ptr *QmlBridge) ConnectAddTransactionToList(f func(paymentID string, transactionID string, amount string, confirmations string, time string, number string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addTransactionToList") {
			C.QmlBridge40a792_ConnectAddTransactionToList(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addTransactionToList"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addTransactionToList", func(paymentID string, transactionID string, amount string, confirmations string, time string, number string) {
				signal.(func(string, string, string, string, string, string))(paymentID, transactionID, amount, confirmations, time, number)
				f(paymentID, transactionID, amount, confirmations, time, number)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addTransactionToList", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectAddTransactionToList() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectAddTransactionToList(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addTransactionToList")
	}
}

func (ptr *QmlBridge) AddTransactionToList(paymentID string, transactionID string, amount string, confirmations string, time string, number string) {
	if ptr.Pointer() != nil {
		var paymentIDC *C.char
		if paymentID != "" {
			paymentIDC = C.CString(paymentID)
			defer C.free(unsafe.Pointer(paymentIDC))
		}
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		var amountC *C.char
		if amount != "" {
			amountC = C.CString(amount)
			defer C.free(unsafe.Pointer(amountC))
		}
		var confirmationsC *C.char
		if confirmations != "" {
			confirmationsC = C.CString(confirmations)
			defer C.free(unsafe.Pointer(confirmationsC))
		}
		var timeC *C.char
		if time != "" {
			timeC = C.CString(time)
			defer C.free(unsafe.Pointer(timeC))
		}
		var numberC *C.char
		if number != "" {
			numberC = C.CString(number)
			defer C.free(unsafe.Pointer(numberC))
		}
		C.QmlBridge40a792_AddTransactionToList(ptr.Pointer(), C.struct_Moc_PackedString{data: paymentIDC, len: C.longlong(len(paymentID))}, C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))}, C.struct_Moc_PackedString{data: amountC, len: C.longlong(len(amount))}, C.struct_Moc_PackedString{data: confirmationsC, len: C.longlong(len(confirmations))}, C.struct_Moc_PackedString{data: timeC, len: C.longlong(len(time))}, C.struct_Moc_PackedString{data: numberC, len: C.longlong(len(number))})
	}
}

//export callbackQmlBridge40a792_AddRemoteNodeToList
func callbackQmlBridge40a792_AddRemoteNodeToList(ptr unsafe.Pointer, nodeURL C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addRemoteNodeToList"); signal != nil {
		signal.(func(string))(cGoUnpackString(nodeURL))
	}

}

func (ptr *QmlBridge) ConnectAddRemoteNodeToList(f func(nodeURL string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addRemoteNodeToList") {
			C.QmlBridge40a792_ConnectAddRemoteNodeToList(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addRemoteNodeToList"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addRemoteNodeToList", func(nodeURL string) {
				signal.(func(string))(nodeURL)
				f(nodeURL)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addRemoteNodeToList", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectAddRemoteNodeToList() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectAddRemoteNodeToList(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addRemoteNodeToList")
	}
}

func (ptr *QmlBridge) AddRemoteNodeToList(nodeURL string) {
	if ptr.Pointer() != nil {
		var nodeURLC *C.char
		if nodeURL != "" {
			nodeURLC = C.CString(nodeURL)
			defer C.free(unsafe.Pointer(nodeURLC))
		}
		C.QmlBridge40a792_AddRemoteNodeToList(ptr.Pointer(), C.struct_Moc_PackedString{data: nodeURLC, len: C.longlong(len(nodeURL))})
	}
}

//export callbackQmlBridge40a792_SetSelectedRemoteNode
func callbackQmlBridge40a792_SetSelectedRemoteNode(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "setSelectedRemoteNode"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QmlBridge) ConnectSetSelectedRemoteNode(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "setSelectedRemoteNode") {
			C.QmlBridge40a792_ConnectSetSelectedRemoteNode(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "setSelectedRemoteNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setSelectedRemoteNode", func(index int) {
				signal.(func(int))(index)
				f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSelectedRemoteNode", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectSetSelectedRemoteNode() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectSetSelectedRemoteNode(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "setSelectedRemoteNode")
	}
}

func (ptr *QmlBridge) SetSelectedRemoteNode(index int) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_SetSelectedRemoteNode(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQmlBridge40a792_DisplayPopup
func callbackQmlBridge40a792_DisplayPopup(ptr unsafe.Pointer, text C.struct_Moc_PackedString, time C.int) {
	if signal := qt.GetSignal(ptr, "displayPopup"); signal != nil {
		signal.(func(string, int))(cGoUnpackString(text), int(int32(time)))
	}

}

func (ptr *QmlBridge) ConnectDisplayPopup(f func(text string, time int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayPopup") {
			C.QmlBridge40a792_ConnectDisplayPopup(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayPopup"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayPopup", func(text string, time int) {
				signal.(func(string, int))(text, time)
				f(text, time)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayPopup", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayPopup() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayPopup(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayPopup")
	}
}

func (ptr *QmlBridge) DisplayPopup(text string, time int) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QmlBridge40a792_DisplayPopup(ptr.Pointer(), C.struct_Moc_PackedString{data: textC, len: C.longlong(len(text))}, C.int(int32(time)))
	}
}

//export callbackQmlBridge40a792_DisplaySyncingInfo
func callbackQmlBridge40a792_DisplaySyncingInfo(ptr unsafe.Pointer, syncing C.struct_Moc_PackedString, syncingInfo C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displaySyncingInfo"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(syncing), cGoUnpackString(syncingInfo))
	}

}

func (ptr *QmlBridge) ConnectDisplaySyncingInfo(f func(syncing string, syncingInfo string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displaySyncingInfo") {
			C.QmlBridge40a792_ConnectDisplaySyncingInfo(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displaySyncingInfo"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displaySyncingInfo", func(syncing string, syncingInfo string) {
				signal.(func(string, string))(syncing, syncingInfo)
				f(syncing, syncingInfo)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displaySyncingInfo", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplaySyncingInfo() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplaySyncingInfo(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displaySyncingInfo")
	}
}

func (ptr *QmlBridge) DisplaySyncingInfo(syncing string, syncingInfo string) {
	if ptr.Pointer() != nil {
		var syncingC *C.char
		if syncing != "" {
			syncingC = C.CString(syncing)
			defer C.free(unsafe.Pointer(syncingC))
		}
		var syncingInfoC *C.char
		if syncingInfo != "" {
			syncingInfoC = C.CString(syncingInfo)
			defer C.free(unsafe.Pointer(syncingInfoC))
		}
		C.QmlBridge40a792_DisplaySyncingInfo(ptr.Pointer(), C.struct_Moc_PackedString{data: syncingC, len: C.longlong(len(syncing))}, C.struct_Moc_PackedString{data: syncingInfoC, len: C.longlong(len(syncingInfo))})
	}
}

//export callbackQmlBridge40a792_DisplayErrorDialog
func callbackQmlBridge40a792_DisplayErrorDialog(ptr unsafe.Pointer, errorText C.struct_Moc_PackedString, errorInformativeText C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayErrorDialog"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(errorText), cGoUnpackString(errorInformativeText))
	}

}

func (ptr *QmlBridge) ConnectDisplayErrorDialog(f func(errorText string, errorInformativeText string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayErrorDialog") {
			C.QmlBridge40a792_ConnectDisplayErrorDialog(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayErrorDialog"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayErrorDialog", func(errorText string, errorInformativeText string) {
				signal.(func(string, string))(errorText, errorInformativeText)
				f(errorText, errorInformativeText)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayErrorDialog", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayErrorDialog() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayErrorDialog(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayErrorDialog")
	}
}

func (ptr *QmlBridge) DisplayErrorDialog(errorText string, errorInformativeText string) {
	if ptr.Pointer() != nil {
		var errorTextC *C.char
		if errorText != "" {
			errorTextC = C.CString(errorText)
			defer C.free(unsafe.Pointer(errorTextC))
		}
		var errorInformativeTextC *C.char
		if errorInformativeText != "" {
			errorInformativeTextC = C.CString(errorInformativeText)
			defer C.free(unsafe.Pointer(errorInformativeTextC))
		}
		C.QmlBridge40a792_DisplayErrorDialog(ptr.Pointer(), C.struct_Moc_PackedString{data: errorTextC, len: C.longlong(len(errorText))}, C.struct_Moc_PackedString{data: errorInformativeTextC, len: C.longlong(len(errorInformativeText))})
	}
}

//export callbackQmlBridge40a792_ClearTransferAmount
func callbackQmlBridge40a792_ClearTransferAmount(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clearTransferAmount"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectClearTransferAmount(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clearTransferAmount") {
			C.QmlBridge40a792_ConnectClearTransferAmount(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clearTransferAmount"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clearTransferAmount", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clearTransferAmount", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClearTransferAmount() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectClearTransferAmount(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clearTransferAmount")
	}
}

func (ptr *QmlBridge) ClearTransferAmount() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ClearTransferAmount(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_AskForFusion
func callbackQmlBridge40a792_AskForFusion(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "askForFusion"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectAskForFusion(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "askForFusion") {
			C.QmlBridge40a792_ConnectAskForFusion(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "askForFusion"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "askForFusion", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "askForFusion", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectAskForFusion() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectAskForFusion(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "askForFusion")
	}
}

func (ptr *QmlBridge) AskForFusion() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_AskForFusion(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_ClearListTransactions
func callbackQmlBridge40a792_ClearListTransactions(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clearListTransactions"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectClearListTransactions(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clearListTransactions") {
			C.QmlBridge40a792_ConnectClearListTransactions(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clearListTransactions"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clearListTransactions", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clearListTransactions", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClearListTransactions() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectClearListTransactions(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clearListTransactions")
	}
}

func (ptr *QmlBridge) ClearListTransactions() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ClearListTransactions(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_DisplayPrivateKeys
func callbackQmlBridge40a792_DisplayPrivateKeys(ptr unsafe.Pointer, filename C.struct_Moc_PackedString, privateViewKey C.struct_Moc_PackedString, privateSpendKey C.struct_Moc_PackedString, walletAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayPrivateKeys"); signal != nil {
		signal.(func(string, string, string, string))(cGoUnpackString(filename), cGoUnpackString(privateViewKey), cGoUnpackString(privateSpendKey), cGoUnpackString(walletAddress))
	}

}

func (ptr *QmlBridge) ConnectDisplayPrivateKeys(f func(filename string, privateViewKey string, privateSpendKey string, walletAddress string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayPrivateKeys") {
			C.QmlBridge40a792_ConnectDisplayPrivateKeys(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayPrivateKeys"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayPrivateKeys", func(filename string, privateViewKey string, privateSpendKey string, walletAddress string) {
				signal.(func(string, string, string, string))(filename, privateViewKey, privateSpendKey, walletAddress)
				f(filename, privateViewKey, privateSpendKey, walletAddress)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayPrivateKeys", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayPrivateKeys() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayPrivateKeys(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayPrivateKeys")
	}
}

func (ptr *QmlBridge) DisplayPrivateKeys(filename string, privateViewKey string, privateSpendKey string, walletAddress string) {
	if ptr.Pointer() != nil {
		var filenameC *C.char
		if filename != "" {
			filenameC = C.CString(filename)
			defer C.free(unsafe.Pointer(filenameC))
		}
		var privateViewKeyC *C.char
		if privateViewKey != "" {
			privateViewKeyC = C.CString(privateViewKey)
			defer C.free(unsafe.Pointer(privateViewKeyC))
		}
		var privateSpendKeyC *C.char
		if privateSpendKey != "" {
			privateSpendKeyC = C.CString(privateSpendKey)
			defer C.free(unsafe.Pointer(privateSpendKeyC))
		}
		var walletAddressC *C.char
		if walletAddress != "" {
			walletAddressC = C.CString(walletAddress)
			defer C.free(unsafe.Pointer(walletAddressC))
		}
		C.QmlBridge40a792_DisplayPrivateKeys(ptr.Pointer(), C.struct_Moc_PackedString{data: filenameC, len: C.longlong(len(filename))}, C.struct_Moc_PackedString{data: privateViewKeyC, len: C.longlong(len(privateViewKey))}, C.struct_Moc_PackedString{data: privateSpendKeyC, len: C.longlong(len(privateSpendKey))}, C.struct_Moc_PackedString{data: walletAddressC, len: C.longlong(len(walletAddress))})
	}
}

//export callbackQmlBridge40a792_DisplaySeed
func callbackQmlBridge40a792_DisplaySeed(ptr unsafe.Pointer, filename C.struct_Moc_PackedString, mnemonicSeed C.struct_Moc_PackedString, walletAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displaySeed"); signal != nil {
		signal.(func(string, string, string))(cGoUnpackString(filename), cGoUnpackString(mnemonicSeed), cGoUnpackString(walletAddress))
	}

}

func (ptr *QmlBridge) ConnectDisplaySeed(f func(filename string, mnemonicSeed string, walletAddress string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displaySeed") {
			C.QmlBridge40a792_ConnectDisplaySeed(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displaySeed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displaySeed", func(filename string, mnemonicSeed string, walletAddress string) {
				signal.(func(string, string, string))(filename, mnemonicSeed, walletAddress)
				f(filename, mnemonicSeed, walletAddress)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displaySeed", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplaySeed() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplaySeed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displaySeed")
	}
}

func (ptr *QmlBridge) DisplaySeed(filename string, mnemonicSeed string, walletAddress string) {
	if ptr.Pointer() != nil {
		var filenameC *C.char
		if filename != "" {
			filenameC = C.CString(filename)
			defer C.free(unsafe.Pointer(filenameC))
		}
		var mnemonicSeedC *C.char
		if mnemonicSeed != "" {
			mnemonicSeedC = C.CString(mnemonicSeed)
			defer C.free(unsafe.Pointer(mnemonicSeedC))
		}
		var walletAddressC *C.char
		if walletAddress != "" {
			walletAddressC = C.CString(walletAddress)
			defer C.free(unsafe.Pointer(walletAddressC))
		}
		C.QmlBridge40a792_DisplaySeed(ptr.Pointer(), C.struct_Moc_PackedString{data: filenameC, len: C.longlong(len(filename))}, C.struct_Moc_PackedString{data: mnemonicSeedC, len: C.longlong(len(mnemonicSeed))}, C.struct_Moc_PackedString{data: walletAddressC, len: C.longlong(len(walletAddress))})
	}
}

//export callbackQmlBridge40a792_DisplayOpenWalletScreen
func callbackQmlBridge40a792_DisplayOpenWalletScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "displayOpenWalletScreen"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectDisplayOpenWalletScreen(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayOpenWalletScreen") {
			C.QmlBridge40a792_ConnectDisplayOpenWalletScreen(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayOpenWalletScreen"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayOpenWalletScreen", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayOpenWalletScreen", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayOpenWalletScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayOpenWalletScreen(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayOpenWalletScreen")
	}
}

func (ptr *QmlBridge) DisplayOpenWalletScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisplayOpenWalletScreen(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_DisplayMainWalletScreen
func callbackQmlBridge40a792_DisplayMainWalletScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "displayMainWalletScreen"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectDisplayMainWalletScreen(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayMainWalletScreen") {
			C.QmlBridge40a792_ConnectDisplayMainWalletScreen(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayMainWalletScreen"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayMainWalletScreen", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayMainWalletScreen", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayMainWalletScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayMainWalletScreen(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayMainWalletScreen")
	}
}

func (ptr *QmlBridge) DisplayMainWalletScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisplayMainWalletScreen(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_FinishedLoadingWalletd
func callbackQmlBridge40a792_FinishedLoadingWalletd(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finishedLoadingWalletd"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectFinishedLoadingWalletd(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finishedLoadingWalletd") {
			C.QmlBridge40a792_ConnectFinishedLoadingWalletd(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finishedLoadingWalletd"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "finishedLoadingWalletd", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finishedLoadingWalletd", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectFinishedLoadingWalletd() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectFinishedLoadingWalletd(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finishedLoadingWalletd")
	}
}

func (ptr *QmlBridge) FinishedLoadingWalletd() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_FinishedLoadingWalletd(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_FinishedCreatingWallet
func callbackQmlBridge40a792_FinishedCreatingWallet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finishedCreatingWallet"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectFinishedCreatingWallet(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finishedCreatingWallet") {
			C.QmlBridge40a792_ConnectFinishedCreatingWallet(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finishedCreatingWallet"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "finishedCreatingWallet", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finishedCreatingWallet", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectFinishedCreatingWallet() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectFinishedCreatingWallet(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finishedCreatingWallet")
	}
}

func (ptr *QmlBridge) FinishedCreatingWallet() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_FinishedCreatingWallet(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_FinishedSendingTransaction
func callbackQmlBridge40a792_FinishedSendingTransaction(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finishedSendingTransaction"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectFinishedSendingTransaction(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finishedSendingTransaction") {
			C.QmlBridge40a792_ConnectFinishedSendingTransaction(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finishedSendingTransaction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "finishedSendingTransaction", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finishedSendingTransaction", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectFinishedSendingTransaction() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectFinishedSendingTransaction(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finishedSendingTransaction")
	}
}

func (ptr *QmlBridge) FinishedSendingTransaction() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_FinishedSendingTransaction(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_DisplayPathToPreviousWallet
func callbackQmlBridge40a792_DisplayPathToPreviousWallet(ptr unsafe.Pointer, pathToPreviousWallet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayPathToPreviousWallet"); signal != nil {
		signal.(func(string))(cGoUnpackString(pathToPreviousWallet))
	}

}

func (ptr *QmlBridge) ConnectDisplayPathToPreviousWallet(f func(pathToPreviousWallet string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayPathToPreviousWallet") {
			C.QmlBridge40a792_ConnectDisplayPathToPreviousWallet(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayPathToPreviousWallet"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayPathToPreviousWallet", func(pathToPreviousWallet string) {
				signal.(func(string))(pathToPreviousWallet)
				f(pathToPreviousWallet)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayPathToPreviousWallet", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayPathToPreviousWallet() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayPathToPreviousWallet(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayPathToPreviousWallet")
	}
}

func (ptr *QmlBridge) DisplayPathToPreviousWallet(pathToPreviousWallet string) {
	if ptr.Pointer() != nil {
		var pathToPreviousWalletC *C.char
		if pathToPreviousWallet != "" {
			pathToPreviousWalletC = C.CString(pathToPreviousWallet)
			defer C.free(unsafe.Pointer(pathToPreviousWalletC))
		}
		C.QmlBridge40a792_DisplayPathToPreviousWallet(ptr.Pointer(), C.struct_Moc_PackedString{data: pathToPreviousWalletC, len: C.longlong(len(pathToPreviousWallet))})
	}
}

//export callbackQmlBridge40a792_DisplayWalletCreationLocation
func callbackQmlBridge40a792_DisplayWalletCreationLocation(ptr unsafe.Pointer, walletLocation C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayWalletCreationLocation"); signal != nil {
		signal.(func(string))(cGoUnpackString(walletLocation))
	}

}

func (ptr *QmlBridge) ConnectDisplayWalletCreationLocation(f func(walletLocation string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayWalletCreationLocation") {
			C.QmlBridge40a792_ConnectDisplayWalletCreationLocation(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayWalletCreationLocation"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayWalletCreationLocation", func(walletLocation string) {
				signal.(func(string))(walletLocation)
				f(walletLocation)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayWalletCreationLocation", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayWalletCreationLocation() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayWalletCreationLocation(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayWalletCreationLocation")
	}
}

func (ptr *QmlBridge) DisplayWalletCreationLocation(walletLocation string) {
	if ptr.Pointer() != nil {
		var walletLocationC *C.char
		if walletLocation != "" {
			walletLocationC = C.CString(walletLocation)
			defer C.free(unsafe.Pointer(walletLocationC))
		}
		C.QmlBridge40a792_DisplayWalletCreationLocation(ptr.Pointer(), C.struct_Moc_PackedString{data: walletLocationC, len: C.longlong(len(walletLocation))})
	}
}

//export callbackQmlBridge40a792_DisplayUseRemoteNode
func callbackQmlBridge40a792_DisplayUseRemoteNode(ptr unsafe.Pointer, useRemote C.char) {
	if signal := qt.GetSignal(ptr, "displayUseRemoteNode"); signal != nil {
		signal.(func(bool))(int8(useRemote) != 0)
	}

}

func (ptr *QmlBridge) ConnectDisplayUseRemoteNode(f func(useRemote bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayUseRemoteNode") {
			C.QmlBridge40a792_ConnectDisplayUseRemoteNode(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayUseRemoteNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayUseRemoteNode", func(useRemote bool) {
				signal.(func(bool))(useRemote)
				f(useRemote)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayUseRemoteNode", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayUseRemoteNode() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayUseRemoteNode(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayUseRemoteNode")
	}
}

func (ptr *QmlBridge) DisplayUseRemoteNode(useRemote bool) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisplayUseRemoteNode(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(useRemote))))
	}
}

//export callbackQmlBridge40a792_HideSettingsScreen
func callbackQmlBridge40a792_HideSettingsScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideSettingsScreen"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectHideSettingsScreen(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hideSettingsScreen") {
			C.QmlBridge40a792_ConnectHideSettingsScreen(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hideSettingsScreen"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hideSettingsScreen", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hideSettingsScreen", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectHideSettingsScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectHideSettingsScreen(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hideSettingsScreen")
	}
}

func (ptr *QmlBridge) HideSettingsScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_HideSettingsScreen(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_DisplaySettingsScreen
func callbackQmlBridge40a792_DisplaySettingsScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "displaySettingsScreen"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectDisplaySettingsScreen(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displaySettingsScreen") {
			C.QmlBridge40a792_ConnectDisplaySettingsScreen(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displaySettingsScreen"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displaySettingsScreen", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displaySettingsScreen", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplaySettingsScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplaySettingsScreen(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displaySettingsScreen")
	}
}

func (ptr *QmlBridge) DisplaySettingsScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisplaySettingsScreen(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_DisplaySettingsValues
func callbackQmlBridge40a792_DisplaySettingsValues(ptr unsafe.Pointer, displayFiat C.char) {
	if signal := qt.GetSignal(ptr, "displaySettingsValues"); signal != nil {
		signal.(func(bool))(int8(displayFiat) != 0)
	}

}

func (ptr *QmlBridge) ConnectDisplaySettingsValues(f func(displayFiat bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displaySettingsValues") {
			C.QmlBridge40a792_ConnectDisplaySettingsValues(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displaySettingsValues"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displaySettingsValues", func(displayFiat bool) {
				signal.(func(bool))(displayFiat)
				f(displayFiat)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displaySettingsValues", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplaySettingsValues() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplaySettingsValues(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displaySettingsValues")
	}
}

func (ptr *QmlBridge) DisplaySettingsValues(displayFiat bool) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisplaySettingsValues(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(displayFiat))))
	}
}

//export callbackQmlBridge40a792_DisplaySettingsRemoteDaemonInfo
func callbackQmlBridge40a792_DisplaySettingsRemoteDaemonInfo(ptr unsafe.Pointer, remoteNodeAddress C.struct_Moc_PackedString, remoteNodePort C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displaySettingsRemoteDaemonInfo"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(remoteNodeAddress), cGoUnpackString(remoteNodePort))
	}

}

func (ptr *QmlBridge) ConnectDisplaySettingsRemoteDaemonInfo(f func(remoteNodeAddress string, remoteNodePort string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displaySettingsRemoteDaemonInfo") {
			C.QmlBridge40a792_ConnectDisplaySettingsRemoteDaemonInfo(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displaySettingsRemoteDaemonInfo"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displaySettingsRemoteDaemonInfo", func(remoteNodeAddress string, remoteNodePort string) {
				signal.(func(string, string))(remoteNodeAddress, remoteNodePort)
				f(remoteNodeAddress, remoteNodePort)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displaySettingsRemoteDaemonInfo", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplaySettingsRemoteDaemonInfo() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplaySettingsRemoteDaemonInfo(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displaySettingsRemoteDaemonInfo")
	}
}

func (ptr *QmlBridge) DisplaySettingsRemoteDaemonInfo(remoteNodeAddress string, remoteNodePort string) {
	if ptr.Pointer() != nil {
		var remoteNodeAddressC *C.char
		if remoteNodeAddress != "" {
			remoteNodeAddressC = C.CString(remoteNodeAddress)
			defer C.free(unsafe.Pointer(remoteNodeAddressC))
		}
		var remoteNodePortC *C.char
		if remoteNodePort != "" {
			remoteNodePortC = C.CString(remoteNodePort)
			defer C.free(unsafe.Pointer(remoteNodePortC))
		}
		C.QmlBridge40a792_DisplaySettingsRemoteDaemonInfo(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteNodeAddressC, len: C.longlong(len(remoteNodeAddress))}, C.struct_Moc_PackedString{data: remoteNodePortC, len: C.longlong(len(remoteNodePort))})
	}
}

//export callbackQmlBridge40a792_DisplayFullBalanceInTransferAmount
func callbackQmlBridge40a792_DisplayFullBalanceInTransferAmount(ptr unsafe.Pointer, fullBalance C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayFullBalanceInTransferAmount"); signal != nil {
		signal.(func(string))(cGoUnpackString(fullBalance))
	}

}

func (ptr *QmlBridge) ConnectDisplayFullBalanceInTransferAmount(f func(fullBalance string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayFullBalanceInTransferAmount") {
			C.QmlBridge40a792_ConnectDisplayFullBalanceInTransferAmount(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayFullBalanceInTransferAmount"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayFullBalanceInTransferAmount", func(fullBalance string) {
				signal.(func(string))(fullBalance)
				f(fullBalance)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayFullBalanceInTransferAmount", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayFullBalanceInTransferAmount() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayFullBalanceInTransferAmount(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayFullBalanceInTransferAmount")
	}
}

func (ptr *QmlBridge) DisplayFullBalanceInTransferAmount(fullBalance string) {
	if ptr.Pointer() != nil {
		var fullBalanceC *C.char
		if fullBalance != "" {
			fullBalanceC = C.CString(fullBalance)
			defer C.free(unsafe.Pointer(fullBalanceC))
		}
		C.QmlBridge40a792_DisplayFullBalanceInTransferAmount(ptr.Pointer(), C.struct_Moc_PackedString{data: fullBalanceC, len: C.longlong(len(fullBalance))})
	}
}

//export callbackQmlBridge40a792_DisplayDefaultFee
func callbackQmlBridge40a792_DisplayDefaultFee(ptr unsafe.Pointer, fee C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayDefaultFee"); signal != nil {
		signal.(func(string))(cGoUnpackString(fee))
	}

}

func (ptr *QmlBridge) ConnectDisplayDefaultFee(f func(fee string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayDefaultFee") {
			C.QmlBridge40a792_ConnectDisplayDefaultFee(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayDefaultFee"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayDefaultFee", func(fee string) {
				signal.(func(string))(fee)
				f(fee)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayDefaultFee", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayDefaultFee() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayDefaultFee(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayDefaultFee")
	}
}

func (ptr *QmlBridge) DisplayDefaultFee(fee string) {
	if ptr.Pointer() != nil {
		var feeC *C.char
		if fee != "" {
			feeC = C.CString(fee)
			defer C.free(unsafe.Pointer(feeC))
		}
		C.QmlBridge40a792_DisplayDefaultFee(ptr.Pointer(), C.struct_Moc_PackedString{data: feeC, len: C.longlong(len(fee))})
	}
}

//export callbackQmlBridge40a792_DisplayNodeFee
func callbackQmlBridge40a792_DisplayNodeFee(ptr unsafe.Pointer, nodeFee C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayNodeFee"); signal != nil {
		signal.(func(string))(cGoUnpackString(nodeFee))
	}

}

func (ptr *QmlBridge) ConnectDisplayNodeFee(f func(nodeFee string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayNodeFee") {
			C.QmlBridge40a792_ConnectDisplayNodeFee(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayNodeFee"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayNodeFee", func(nodeFee string) {
				signal.(func(string))(nodeFee)
				f(nodeFee)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayNodeFee", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayNodeFee() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayNodeFee(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayNodeFee")
	}
}

func (ptr *QmlBridge) DisplayNodeFee(nodeFee string) {
	if ptr.Pointer() != nil {
		var nodeFeeC *C.char
		if nodeFee != "" {
			nodeFeeC = C.CString(nodeFee)
			defer C.free(unsafe.Pointer(nodeFeeC))
		}
		C.QmlBridge40a792_DisplayNodeFee(ptr.Pointer(), C.struct_Moc_PackedString{data: nodeFeeC, len: C.longlong(len(nodeFee))})
	}
}

//export callbackQmlBridge40a792_UpdateConfirmationsOfTransaction
func callbackQmlBridge40a792_UpdateConfirmationsOfTransaction(ptr unsafe.Pointer, index C.int, confirmations C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "updateConfirmationsOfTransaction"); signal != nil {
		signal.(func(int, string))(int(int32(index)), cGoUnpackString(confirmations))
	}

}

func (ptr *QmlBridge) ConnectUpdateConfirmationsOfTransaction(f func(index int, confirmations string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateConfirmationsOfTransaction") {
			C.QmlBridge40a792_ConnectUpdateConfirmationsOfTransaction(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateConfirmationsOfTransaction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateConfirmationsOfTransaction", func(index int, confirmations string) {
				signal.(func(int, string))(index, confirmations)
				f(index, confirmations)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateConfirmationsOfTransaction", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectUpdateConfirmationsOfTransaction() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectUpdateConfirmationsOfTransaction(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateConfirmationsOfTransaction")
	}
}

func (ptr *QmlBridge) UpdateConfirmationsOfTransaction(index int, confirmations string) {
	if ptr.Pointer() != nil {
		var confirmationsC *C.char
		if confirmations != "" {
			confirmationsC = C.CString(confirmations)
			defer C.free(unsafe.Pointer(confirmationsC))
		}
		C.QmlBridge40a792_UpdateConfirmationsOfTransaction(ptr.Pointer(), C.int(int32(index)), C.struct_Moc_PackedString{data: confirmationsC, len: C.longlong(len(confirmations))})
	}
}

//export callbackQmlBridge40a792_DisplayInfoDialog
func callbackQmlBridge40a792_DisplayInfoDialog(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "displayInfoDialog"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectDisplayInfoDialog(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayInfoDialog") {
			C.QmlBridge40a792_ConnectDisplayInfoDialog(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayInfoDialog"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "displayInfoDialog", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayInfoDialog", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayInfoDialog() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectDisplayInfoDialog(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayInfoDialog")
	}
}

func (ptr *QmlBridge) DisplayInfoDialog() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisplayInfoDialog(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_AddSavedAddressToList
func callbackQmlBridge40a792_AddSavedAddressToList(ptr unsafe.Pointer, dbID C.int, name C.struct_Moc_PackedString, address C.struct_Moc_PackedString, paymentID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addSavedAddressToList"); signal != nil {
		signal.(func(int, string, string, string))(int(int32(dbID)), cGoUnpackString(name), cGoUnpackString(address), cGoUnpackString(paymentID))
	}

}

func (ptr *QmlBridge) ConnectAddSavedAddressToList(f func(dbID int, name string, address string, paymentID string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addSavedAddressToList") {
			C.QmlBridge40a792_ConnectAddSavedAddressToList(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addSavedAddressToList"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addSavedAddressToList", func(dbID int, name string, address string, paymentID string) {
				signal.(func(int, string, string, string))(dbID, name, address, paymentID)
				f(dbID, name, address, paymentID)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addSavedAddressToList", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectAddSavedAddressToList() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectAddSavedAddressToList(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addSavedAddressToList")
	}
}

func (ptr *QmlBridge) AddSavedAddressToList(dbID int, name string, address string, paymentID string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		var paymentIDC *C.char
		if paymentID != "" {
			paymentIDC = C.CString(paymentID)
			defer C.free(unsafe.Pointer(paymentIDC))
		}
		C.QmlBridge40a792_AddSavedAddressToList(ptr.Pointer(), C.int(int32(dbID)), C.struct_Moc_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))}, C.struct_Moc_PackedString{data: paymentIDC, len: C.longlong(len(paymentID))})
	}
}

//export callbackQmlBridge40a792_Log
func callbackQmlBridge40a792_Log(ptr unsafe.Pointer, msg C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "log"); signal != nil {
		signal.(func(string))(cGoUnpackString(msg))
	}

}

func (ptr *QmlBridge) ConnectLog(f func(msg string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "log"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "log", func(msg string) {
				signal.(func(string))(msg)
				f(msg)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "log", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectLog() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "log")
	}
}

func (ptr *QmlBridge) Log(msg string) {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.QmlBridge40a792_Log(ptr.Pointer(), C.struct_Moc_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

//export callbackQmlBridge40a792_ClickedButtonExplorer
func callbackQmlBridge40a792_ClickedButtonExplorer(ptr unsafe.Pointer, transactionID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonExplorer"); signal != nil {
		signal.(func(string))(cGoUnpackString(transactionID))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonExplorer(f func(transactionID string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonExplorer"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonExplorer", func(transactionID string) {
				signal.(func(string))(transactionID)
				f(transactionID)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonExplorer", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonExplorer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonExplorer")
	}
}

func (ptr *QmlBridge) ClickedButtonExplorer(transactionID string) {
	if ptr.Pointer() != nil {
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		C.QmlBridge40a792_ClickedButtonExplorer(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

//export callbackQmlBridge40a792_GoToWebsite
func callbackQmlBridge40a792_GoToWebsite(ptr unsafe.Pointer, url C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "goToWebsite"); signal != nil {
		signal.(func(string))(cGoUnpackString(url))
	}

}

func (ptr *QmlBridge) ConnectGoToWebsite(f func(url string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "goToWebsite"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "goToWebsite", func(url string) {
				signal.(func(string))(url)
				f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "goToWebsite", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectGoToWebsite() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "goToWebsite")
	}
}

func (ptr *QmlBridge) GoToWebsite(url string) {
	if ptr.Pointer() != nil {
		var urlC *C.char
		if url != "" {
			urlC = C.CString(url)
			defer C.free(unsafe.Pointer(urlC))
		}
		C.QmlBridge40a792_GoToWebsite(ptr.Pointer(), C.struct_Moc_PackedString{data: urlC, len: C.longlong(len(url))})
	}
}

//export callbackQmlBridge40a792_ClickedButtonCopyTx
func callbackQmlBridge40a792_ClickedButtonCopyTx(ptr unsafe.Pointer, transactionID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonCopyTx"); signal != nil {
		signal.(func(string))(cGoUnpackString(transactionID))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonCopyTx(f func(transactionID string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonCopyTx"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopyTx", func(transactionID string) {
				signal.(func(string))(transactionID)
				f(transactionID)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopyTx", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonCopyTx() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonCopyTx")
	}
}

func (ptr *QmlBridge) ClickedButtonCopyTx(transactionID string) {
	if ptr.Pointer() != nil {
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		C.QmlBridge40a792_ClickedButtonCopyTx(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

//export callbackQmlBridge40a792_ClickedButtonCopyAddress
func callbackQmlBridge40a792_ClickedButtonCopyAddress(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clickedButtonCopyAddress"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectClickedButtonCopyAddress(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonCopyAddress"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopyAddress", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopyAddress", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonCopyAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonCopyAddress")
	}
}

func (ptr *QmlBridge) ClickedButtonCopyAddress() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ClickedButtonCopyAddress(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_ClickedButtonCopyKeys
func callbackQmlBridge40a792_ClickedButtonCopyKeys(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clickedButtonCopyKeys"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectClickedButtonCopyKeys(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonCopyKeys"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopyKeys", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopyKeys", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonCopyKeys() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonCopyKeys")
	}
}

func (ptr *QmlBridge) ClickedButtonCopyKeys() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ClickedButtonCopyKeys(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_ClickedButtonCopy
func callbackQmlBridge40a792_ClickedButtonCopy(ptr unsafe.Pointer, stringToCopy C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonCopy"); signal != nil {
		signal.(func(string))(cGoUnpackString(stringToCopy))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonCopy(f func(stringToCopy string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonCopy"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopy", func(stringToCopy string) {
				signal.(func(string))(stringToCopy)
				f(stringToCopy)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopy", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonCopy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonCopy")
	}
}

func (ptr *QmlBridge) ClickedButtonCopy(stringToCopy string) {
	if ptr.Pointer() != nil {
		var stringToCopyC *C.char
		if stringToCopy != "" {
			stringToCopyC = C.CString(stringToCopy)
			defer C.free(unsafe.Pointer(stringToCopyC))
		}
		C.QmlBridge40a792_ClickedButtonCopy(ptr.Pointer(), C.struct_Moc_PackedString{data: stringToCopyC, len: C.longlong(len(stringToCopy))})
	}
}

//export callbackQmlBridge40a792_ClickedButtonSend
func callbackQmlBridge40a792_ClickedButtonSend(ptr unsafe.Pointer, transferAddress C.struct_Moc_PackedString, transferAmount C.struct_Moc_PackedString, transferPaymentID C.struct_Moc_PackedString, transferFee C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonSend"); signal != nil {
		signal.(func(string, string, string, string))(cGoUnpackString(transferAddress), cGoUnpackString(transferAmount), cGoUnpackString(transferPaymentID), cGoUnpackString(transferFee))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonSend(f func(transferAddress string, transferAmount string, transferPaymentID string, transferFee string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonSend"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonSend", func(transferAddress string, transferAmount string, transferPaymentID string, transferFee string) {
				signal.(func(string, string, string, string))(transferAddress, transferAmount, transferPaymentID, transferFee)
				f(transferAddress, transferAmount, transferPaymentID, transferFee)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonSend", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonSend() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonSend")
	}
}

func (ptr *QmlBridge) ClickedButtonSend(transferAddress string, transferAmount string, transferPaymentID string, transferFee string) {
	if ptr.Pointer() != nil {
		var transferAddressC *C.char
		if transferAddress != "" {
			transferAddressC = C.CString(transferAddress)
			defer C.free(unsafe.Pointer(transferAddressC))
		}
		var transferAmountC *C.char
		if transferAmount != "" {
			transferAmountC = C.CString(transferAmount)
			defer C.free(unsafe.Pointer(transferAmountC))
		}
		var transferPaymentIDC *C.char
		if transferPaymentID != "" {
			transferPaymentIDC = C.CString(transferPaymentID)
			defer C.free(unsafe.Pointer(transferPaymentIDC))
		}
		var transferFeeC *C.char
		if transferFee != "" {
			transferFeeC = C.CString(transferFee)
			defer C.free(unsafe.Pointer(transferFeeC))
		}
		C.QmlBridge40a792_ClickedButtonSend(ptr.Pointer(), C.struct_Moc_PackedString{data: transferAddressC, len: C.longlong(len(transferAddress))}, C.struct_Moc_PackedString{data: transferAmountC, len: C.longlong(len(transferAmount))}, C.struct_Moc_PackedString{data: transferPaymentIDC, len: C.longlong(len(transferPaymentID))}, C.struct_Moc_PackedString{data: transferFeeC, len: C.longlong(len(transferFee))})
	}
}

//export callbackQmlBridge40a792_ClickedButtonBackupWallet
func callbackQmlBridge40a792_ClickedButtonBackupWallet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clickedButtonBackupWallet"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectClickedButtonBackupWallet(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonBackupWallet"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonBackupWallet", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonBackupWallet", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonBackupWallet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonBackupWallet")
	}
}

func (ptr *QmlBridge) ClickedButtonBackupWallet() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ClickedButtonBackupWallet(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_ClickedCloseWallet
func callbackQmlBridge40a792_ClickedCloseWallet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clickedCloseWallet"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectClickedCloseWallet(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedCloseWallet"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedCloseWallet", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedCloseWallet", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedCloseWallet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedCloseWallet")
	}
}

func (ptr *QmlBridge) ClickedCloseWallet() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ClickedCloseWallet(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_ClickedButtonOpen
func callbackQmlBridge40a792_ClickedButtonOpen(ptr unsafe.Pointer, pathToWallet C.struct_Moc_PackedString, passwordWallet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonOpen"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(pathToWallet), cGoUnpackString(passwordWallet))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonOpen(f func(pathToWallet string, passwordWallet string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonOpen"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonOpen", func(pathToWallet string, passwordWallet string) {
				signal.(func(string, string))(pathToWallet, passwordWallet)
				f(pathToWallet, passwordWallet)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonOpen", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonOpen")
	}
}

func (ptr *QmlBridge) ClickedButtonOpen(pathToWallet string, passwordWallet string) {
	if ptr.Pointer() != nil {
		var pathToWalletC *C.char
		if pathToWallet != "" {
			pathToWalletC = C.CString(pathToWallet)
			defer C.free(unsafe.Pointer(pathToWalletC))
		}
		var passwordWalletC *C.char
		if passwordWallet != "" {
			passwordWalletC = C.CString(passwordWallet)
			defer C.free(unsafe.Pointer(passwordWalletC))
		}
		C.QmlBridge40a792_ClickedButtonOpen(ptr.Pointer(), C.struct_Moc_PackedString{data: pathToWalletC, len: C.longlong(len(pathToWallet))}, C.struct_Moc_PackedString{data: passwordWalletC, len: C.longlong(len(passwordWallet))})
	}
}

//export callbackQmlBridge40a792_ClickedButtonCreate
func callbackQmlBridge40a792_ClickedButtonCreate(ptr unsafe.Pointer, filenameWallet C.struct_Moc_PackedString, passwordWallet C.struct_Moc_PackedString, confirmPasswordWallet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonCreate"); signal != nil {
		signal.(func(string, string, string))(cGoUnpackString(filenameWallet), cGoUnpackString(passwordWallet), cGoUnpackString(confirmPasswordWallet))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonCreate(f func(filenameWallet string, passwordWallet string, confirmPasswordWallet string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonCreate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCreate", func(filenameWallet string, passwordWallet string, confirmPasswordWallet string) {
				signal.(func(string, string, string))(filenameWallet, passwordWallet, confirmPasswordWallet)
				f(filenameWallet, passwordWallet, confirmPasswordWallet)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCreate", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonCreate")
	}
}

func (ptr *QmlBridge) ClickedButtonCreate(filenameWallet string, passwordWallet string, confirmPasswordWallet string) {
	if ptr.Pointer() != nil {
		var filenameWalletC *C.char
		if filenameWallet != "" {
			filenameWalletC = C.CString(filenameWallet)
			defer C.free(unsafe.Pointer(filenameWalletC))
		}
		var passwordWalletC *C.char
		if passwordWallet != "" {
			passwordWalletC = C.CString(passwordWallet)
			defer C.free(unsafe.Pointer(passwordWalletC))
		}
		var confirmPasswordWalletC *C.char
		if confirmPasswordWallet != "" {
			confirmPasswordWalletC = C.CString(confirmPasswordWallet)
			defer C.free(unsafe.Pointer(confirmPasswordWalletC))
		}
		C.QmlBridge40a792_ClickedButtonCreate(ptr.Pointer(), C.struct_Moc_PackedString{data: filenameWalletC, len: C.longlong(len(filenameWallet))}, C.struct_Moc_PackedString{data: passwordWalletC, len: C.longlong(len(passwordWallet))}, C.struct_Moc_PackedString{data: confirmPasswordWalletC, len: C.longlong(len(confirmPasswordWallet))})
	}
}

//export callbackQmlBridge40a792_ClickedButtonImport
func callbackQmlBridge40a792_ClickedButtonImport(ptr unsafe.Pointer, filenameWallet C.struct_Moc_PackedString, passwordWallet C.struct_Moc_PackedString, privateViewKey C.struct_Moc_PackedString, privateSpendKey C.struct_Moc_PackedString, mnemonicSeed C.struct_Moc_PackedString, confirmPasswordWallet C.struct_Moc_PackedString, scanHeight C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonImport"); signal != nil {
		signal.(func(string, string, string, string, string, string, string))(cGoUnpackString(filenameWallet), cGoUnpackString(passwordWallet), cGoUnpackString(privateViewKey), cGoUnpackString(privateSpendKey), cGoUnpackString(mnemonicSeed), cGoUnpackString(confirmPasswordWallet), cGoUnpackString(scanHeight))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonImport(f func(filenameWallet string, passwordWallet string, privateViewKey string, privateSpendKey string, mnemonicSeed string, confirmPasswordWallet string, scanHeight string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonImport"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonImport", func(filenameWallet string, passwordWallet string, privateViewKey string, privateSpendKey string, mnemonicSeed string, confirmPasswordWallet string, scanHeight string) {
				signal.(func(string, string, string, string, string, string, string))(filenameWallet, passwordWallet, privateViewKey, privateSpendKey, mnemonicSeed, confirmPasswordWallet, scanHeight)
				f(filenameWallet, passwordWallet, privateViewKey, privateSpendKey, mnemonicSeed, confirmPasswordWallet, scanHeight)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonImport", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonImport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonImport")
	}
}

func (ptr *QmlBridge) ClickedButtonImport(filenameWallet string, passwordWallet string, privateViewKey string, privateSpendKey string, mnemonicSeed string, confirmPasswordWallet string, scanHeight string) {
	if ptr.Pointer() != nil {
		var filenameWalletC *C.char
		if filenameWallet != "" {
			filenameWalletC = C.CString(filenameWallet)
			defer C.free(unsafe.Pointer(filenameWalletC))
		}
		var passwordWalletC *C.char
		if passwordWallet != "" {
			passwordWalletC = C.CString(passwordWallet)
			defer C.free(unsafe.Pointer(passwordWalletC))
		}
		var privateViewKeyC *C.char
		if privateViewKey != "" {
			privateViewKeyC = C.CString(privateViewKey)
			defer C.free(unsafe.Pointer(privateViewKeyC))
		}
		var privateSpendKeyC *C.char
		if privateSpendKey != "" {
			privateSpendKeyC = C.CString(privateSpendKey)
			defer C.free(unsafe.Pointer(privateSpendKeyC))
		}
		var mnemonicSeedC *C.char
		if mnemonicSeed != "" {
			mnemonicSeedC = C.CString(mnemonicSeed)
			defer C.free(unsafe.Pointer(mnemonicSeedC))
		}
		var confirmPasswordWalletC *C.char
		if confirmPasswordWallet != "" {
			confirmPasswordWalletC = C.CString(confirmPasswordWallet)
			defer C.free(unsafe.Pointer(confirmPasswordWalletC))
		}
		var scanHeightC *C.char
		if scanHeight != "" {
			scanHeightC = C.CString(scanHeight)
			defer C.free(unsafe.Pointer(scanHeightC))
		}
		C.QmlBridge40a792_ClickedButtonImport(ptr.Pointer(), C.struct_Moc_PackedString{data: filenameWalletC, len: C.longlong(len(filenameWallet))}, C.struct_Moc_PackedString{data: passwordWalletC, len: C.longlong(len(passwordWallet))}, C.struct_Moc_PackedString{data: privateViewKeyC, len: C.longlong(len(privateViewKey))}, C.struct_Moc_PackedString{data: privateSpendKeyC, len: C.longlong(len(privateSpendKey))}, C.struct_Moc_PackedString{data: mnemonicSeedC, len: C.longlong(len(mnemonicSeed))}, C.struct_Moc_PackedString{data: confirmPasswordWalletC, len: C.longlong(len(confirmPasswordWallet))}, C.struct_Moc_PackedString{data: scanHeightC, len: C.longlong(len(scanHeight))})
	}
}

//export callbackQmlBridge40a792_ChoseRemote
func callbackQmlBridge40a792_ChoseRemote(ptr unsafe.Pointer, remote C.char) {
	if signal := qt.GetSignal(ptr, "choseRemote"); signal != nil {
		signal.(func(bool))(int8(remote) != 0)
	}

}

func (ptr *QmlBridge) ConnectChoseRemote(f func(remote bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "choseRemote"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "choseRemote", func(remote bool) {
				signal.(func(bool))(remote)
				f(remote)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "choseRemote", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectChoseRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "choseRemote")
	}
}

func (ptr *QmlBridge) ChoseRemote(remote bool) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ChoseRemote(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(remote))))
	}
}

//export callbackQmlBridge40a792_SelectedRemoteNode
func callbackQmlBridge40a792_SelectedRemoteNode(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "selectedRemoteNode"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QmlBridge) ConnectSelectedRemoteNode(f func(index int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectedRemoteNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "selectedRemoteNode", func(index int) {
				signal.(func(int))(index)
				f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectedRemoteNode", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectSelectedRemoteNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectedRemoteNode")
	}
}

func (ptr *QmlBridge) SelectedRemoteNode(index int) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_SelectedRemoteNode(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQmlBridge40a792_GetTransferAmountUSD
func callbackQmlBridge40a792_GetTransferAmountUSD(ptr unsafe.Pointer, amountTRTL C.struct_Moc_PackedString) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "getTransferAmountUSD"); signal != nil {
		tempVal := signal.(func(string) string)(cGoUnpackString(amountTRTL))
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QmlBridge) ConnectGetTransferAmountUSD(f func(amountTRTL string) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getTransferAmountUSD"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "getTransferAmountUSD", func(amountTRTL string) string {
				signal.(func(string) string)(amountTRTL)
				return f(amountTRTL)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getTransferAmountUSD", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectGetTransferAmountUSD() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getTransferAmountUSD")
	}
}

func (ptr *QmlBridge) GetTransferAmountUSD(amountTRTL string) string {
	if ptr.Pointer() != nil {
		var amountTRTLC *C.char
		if amountTRTL != "" {
			amountTRTLC = C.CString(amountTRTL)
			defer C.free(unsafe.Pointer(amountTRTLC))
		}
		return cGoUnpackString(C.QmlBridge40a792_GetTransferAmountUSD(ptr.Pointer(), C.struct_Moc_PackedString{data: amountTRTLC, len: C.longlong(len(amountTRTL))}))
	}
	return ""
}

//export callbackQmlBridge40a792_ClickedCloseSettings
func callbackQmlBridge40a792_ClickedCloseSettings(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clickedCloseSettings"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectClickedCloseSettings(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedCloseSettings"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedCloseSettings", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedCloseSettings", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedCloseSettings() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedCloseSettings")
	}
}

func (ptr *QmlBridge) ClickedCloseSettings() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ClickedCloseSettings(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_ClickedSettingsButton
func callbackQmlBridge40a792_ClickedSettingsButton(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clickedSettingsButton"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectClickedSettingsButton(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedSettingsButton"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clickedSettingsButton", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedSettingsButton", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedSettingsButton() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedSettingsButton")
	}
}

func (ptr *QmlBridge) ClickedSettingsButton() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ClickedSettingsButton(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_ChoseDisplayFiat
func callbackQmlBridge40a792_ChoseDisplayFiat(ptr unsafe.Pointer, displayFiat C.char) {
	if signal := qt.GetSignal(ptr, "choseDisplayFiat"); signal != nil {
		signal.(func(bool))(int8(displayFiat) != 0)
	}

}

func (ptr *QmlBridge) ConnectChoseDisplayFiat(f func(displayFiat bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "choseDisplayFiat"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "choseDisplayFiat", func(displayFiat bool) {
				signal.(func(bool))(displayFiat)
				f(displayFiat)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "choseDisplayFiat", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectChoseDisplayFiat() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "choseDisplayFiat")
	}
}

func (ptr *QmlBridge) ChoseDisplayFiat(displayFiat bool) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ChoseDisplayFiat(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(displayFiat))))
	}
}

//export callbackQmlBridge40a792_ChoseCheckpoints
func callbackQmlBridge40a792_ChoseCheckpoints(ptr unsafe.Pointer, checkpoints C.char) {
	if signal := qt.GetSignal(ptr, "choseCheckpoints"); signal != nil {
		signal.(func(bool))(int8(checkpoints) != 0)
	}

}

func (ptr *QmlBridge) ConnectChoseCheckpoints(f func(checkpoints bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "choseCheckpoints"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "choseCheckpoints", func(checkpoints bool) {
				signal.(func(bool))(checkpoints)
				f(checkpoints)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "choseCheckpoints", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectChoseCheckpoints() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "choseCheckpoints")
	}
}

func (ptr *QmlBridge) ChoseCheckpoints(checkpoints bool) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ChoseCheckpoints(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(checkpoints))))
	}
}

//export callbackQmlBridge40a792_SaveRemoteDaemonInfo
func callbackQmlBridge40a792_SaveRemoteDaemonInfo(ptr unsafe.Pointer, daemonAddress C.struct_Moc_PackedString, daemonPort C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "saveRemoteDaemonInfo"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(daemonAddress), cGoUnpackString(daemonPort))
	}

}

func (ptr *QmlBridge) ConnectSaveRemoteDaemonInfo(f func(daemonAddress string, daemonPort string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "saveRemoteDaemonInfo"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "saveRemoteDaemonInfo", func(daemonAddress string, daemonPort string) {
				signal.(func(string, string))(daemonAddress, daemonPort)
				f(daemonAddress, daemonPort)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "saveRemoteDaemonInfo", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectSaveRemoteDaemonInfo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "saveRemoteDaemonInfo")
	}
}

func (ptr *QmlBridge) SaveRemoteDaemonInfo(daemonAddress string, daemonPort string) {
	if ptr.Pointer() != nil {
		var daemonAddressC *C.char
		if daemonAddress != "" {
			daemonAddressC = C.CString(daemonAddress)
			defer C.free(unsafe.Pointer(daemonAddressC))
		}
		var daemonPortC *C.char
		if daemonPort != "" {
			daemonPortC = C.CString(daemonPort)
			defer C.free(unsafe.Pointer(daemonPortC))
		}
		C.QmlBridge40a792_SaveRemoteDaemonInfo(ptr.Pointer(), C.struct_Moc_PackedString{data: daemonAddressC, len: C.longlong(len(daemonAddress))}, C.struct_Moc_PackedString{data: daemonPortC, len: C.longlong(len(daemonPort))})
	}
}

//export callbackQmlBridge40a792_ResetRemoteDaemonInfo
func callbackQmlBridge40a792_ResetRemoteDaemonInfo(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetRemoteDaemonInfo"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectResetRemoteDaemonInfo(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resetRemoteDaemonInfo"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "resetRemoteDaemonInfo", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resetRemoteDaemonInfo", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectResetRemoteDaemonInfo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resetRemoteDaemonInfo")
	}
}

func (ptr *QmlBridge) ResetRemoteDaemonInfo() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ResetRemoteDaemonInfo(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_GetFullBalanceAndDisplayInTransferAmount
func callbackQmlBridge40a792_GetFullBalanceAndDisplayInTransferAmount(ptr unsafe.Pointer, transferFee C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "getFullBalanceAndDisplayInTransferAmount"); signal != nil {
		signal.(func(string))(cGoUnpackString(transferFee))
	}

}

func (ptr *QmlBridge) ConnectGetFullBalanceAndDisplayInTransferAmount(f func(transferFee string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getFullBalanceAndDisplayInTransferAmount"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "getFullBalanceAndDisplayInTransferAmount", func(transferFee string) {
				signal.(func(string))(transferFee)
				f(transferFee)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getFullBalanceAndDisplayInTransferAmount", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectGetFullBalanceAndDisplayInTransferAmount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getFullBalanceAndDisplayInTransferAmount")
	}
}

func (ptr *QmlBridge) GetFullBalanceAndDisplayInTransferAmount(transferFee string) {
	if ptr.Pointer() != nil {
		var transferFeeC *C.char
		if transferFee != "" {
			transferFeeC = C.CString(transferFee)
			defer C.free(unsafe.Pointer(transferFeeC))
		}
		C.QmlBridge40a792_GetFullBalanceAndDisplayInTransferAmount(ptr.Pointer(), C.struct_Moc_PackedString{data: transferFeeC, len: C.longlong(len(transferFee))})
	}
}

//export callbackQmlBridge40a792_GetDefaultFeeAndDisplay
func callbackQmlBridge40a792_GetDefaultFeeAndDisplay(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "getDefaultFeeAndDisplay"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectGetDefaultFeeAndDisplay(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getDefaultFeeAndDisplay"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "getDefaultFeeAndDisplay", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getDefaultFeeAndDisplay", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectGetDefaultFeeAndDisplay() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getDefaultFeeAndDisplay")
	}
}

func (ptr *QmlBridge) GetDefaultFeeAndDisplay() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_GetDefaultFeeAndDisplay(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_LimitDisplayTransactions
func callbackQmlBridge40a792_LimitDisplayTransactions(ptr unsafe.Pointer, limit C.char) {
	if signal := qt.GetSignal(ptr, "limitDisplayTransactions"); signal != nil {
		signal.(func(bool))(int8(limit) != 0)
	}

}

func (ptr *QmlBridge) ConnectLimitDisplayTransactions(f func(limit bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "limitDisplayTransactions"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "limitDisplayTransactions", func(limit bool) {
				signal.(func(bool))(limit)
				f(limit)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "limitDisplayTransactions", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectLimitDisplayTransactions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "limitDisplayTransactions")
	}
}

func (ptr *QmlBridge) LimitDisplayTransactions(limit bool) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_LimitDisplayTransactions(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(limit))))
	}
}

//export callbackQmlBridge40a792_GetVersion
func callbackQmlBridge40a792_GetVersion(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "getVersion"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QmlBridge) ConnectGetVersion(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getVersion"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "getVersion", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getVersion", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectGetVersion() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getVersion")
	}
}

func (ptr *QmlBridge) GetVersion() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QmlBridge40a792_GetVersion(ptr.Pointer()))
	}
	return ""
}

//export callbackQmlBridge40a792_GetNewVersion
func callbackQmlBridge40a792_GetNewVersion(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "getNewVersion"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QmlBridge) ConnectGetNewVersion(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getNewVersion"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "getNewVersion", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getNewVersion", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectGetNewVersion() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getNewVersion")
	}
}

func (ptr *QmlBridge) GetNewVersion() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QmlBridge40a792_GetNewVersion(ptr.Pointer()))
	}
	return ""
}

//export callbackQmlBridge40a792_GetNewVersionURL
func callbackQmlBridge40a792_GetNewVersionURL(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "getNewVersionURL"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QmlBridge) ConnectGetNewVersionURL(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getNewVersionURL"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "getNewVersionURL", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getNewVersionURL", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectGetNewVersionURL() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getNewVersionURL")
	}
}

func (ptr *QmlBridge) GetNewVersionURL() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QmlBridge40a792_GetNewVersionURL(ptr.Pointer()))
	}
	return ""
}

//export callbackQmlBridge40a792_OptimizeWalletWithFusion
func callbackQmlBridge40a792_OptimizeWalletWithFusion(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "optimizeWalletWithFusion"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectOptimizeWalletWithFusion(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "optimizeWalletWithFusion"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "optimizeWalletWithFusion", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "optimizeWalletWithFusion", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectOptimizeWalletWithFusion() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "optimizeWalletWithFusion")
	}
}

func (ptr *QmlBridge) OptimizeWalletWithFusion() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_OptimizeWalletWithFusion(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_SaveAddress
func callbackQmlBridge40a792_SaveAddress(ptr unsafe.Pointer, name C.struct_Moc_PackedString, address C.struct_Moc_PackedString, paymentID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "saveAddress"); signal != nil {
		signal.(func(string, string, string))(cGoUnpackString(name), cGoUnpackString(address), cGoUnpackString(paymentID))
	}

}

func (ptr *QmlBridge) ConnectSaveAddress(f func(name string, address string, paymentID string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "saveAddress"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "saveAddress", func(name string, address string, paymentID string) {
				signal.(func(string, string, string))(name, address, paymentID)
				f(name, address, paymentID)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "saveAddress", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectSaveAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "saveAddress")
	}
}

func (ptr *QmlBridge) SaveAddress(name string, address string, paymentID string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		var paymentIDC *C.char
		if paymentID != "" {
			paymentIDC = C.CString(paymentID)
			defer C.free(unsafe.Pointer(paymentIDC))
		}
		C.QmlBridge40a792_SaveAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))}, C.struct_Moc_PackedString{data: paymentIDC, len: C.longlong(len(paymentID))})
	}
}

//export callbackQmlBridge40a792_FillListSavedAddresses
func callbackQmlBridge40a792_FillListSavedAddresses(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fillListSavedAddresses"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QmlBridge) ConnectFillListSavedAddresses(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fillListSavedAddresses"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fillListSavedAddresses", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fillListSavedAddresses", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectFillListSavedAddresses() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fillListSavedAddresses")
	}
}

func (ptr *QmlBridge) FillListSavedAddresses() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_FillListSavedAddresses(ptr.Pointer())
	}
}

//export callbackQmlBridge40a792_DeleteSavedAddress
func callbackQmlBridge40a792_DeleteSavedAddress(ptr unsafe.Pointer, dbID C.int) {
	if signal := qt.GetSignal(ptr, "deleteSavedAddress"); signal != nil {
		signal.(func(int))(int(int32(dbID)))
	}

}

func (ptr *QmlBridge) ConnectDeleteSavedAddress(f func(dbID int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "deleteSavedAddress"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "deleteSavedAddress", func(dbID int) {
				signal.(func(int))(dbID)
				f(dbID)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deleteSavedAddress", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDeleteSavedAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "deleteSavedAddress")
	}
}

func (ptr *QmlBridge) DeleteSavedAddress(dbID int) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DeleteSavedAddress(ptr.Pointer(), C.int(int32(dbID)))
	}
}

//export callbackQmlBridge40a792_RegisterToGo
func callbackQmlBridge40a792_RegisterToGo(ptr unsafe.Pointer, object unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "registerToGo"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(object))
	}

}

func (ptr *QmlBridge) ConnectRegisterToGo(f func(object *std_core.QObject)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "registerToGo"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "registerToGo", func(object *std_core.QObject) {
				signal.(func(*std_core.QObject))(object)
				f(object)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "registerToGo", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectRegisterToGo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "registerToGo")
	}
}

func (ptr *QmlBridge) RegisterToGo(object std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_RegisterToGo(ptr.Pointer(), std_core.PointerFromQObject(object))
	}
}

//export callbackQmlBridge40a792_DeregisterToGo
func callbackQmlBridge40a792_DeregisterToGo(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "deregisterToGo"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

func (ptr *QmlBridge) ConnectDeregisterToGo(f func(objectName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "deregisterToGo"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "deregisterToGo", func(objectName string) {
				signal.(func(string))(objectName)
				f(objectName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deregisterToGo", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDeregisterToGo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "deregisterToGo")
	}
}

func (ptr *QmlBridge) DeregisterToGo(objectName string) {
	if ptr.Pointer() != nil {
		var objectNameC *C.char
		if objectName != "" {
			objectNameC = C.CString(objectName)
			defer C.free(unsafe.Pointer(objectNameC))
		}
		C.QmlBridge40a792_DeregisterToGo(ptr.Pointer(), C.struct_Moc_PackedString{data: objectNameC, len: C.longlong(len(objectName))})
	}
}

func QmlBridge_QRegisterMetaType() int {
	return int(int32(C.QmlBridge40a792_QmlBridge40a792_QRegisterMetaType()))
}

func (ptr *QmlBridge) QRegisterMetaType() int {
	return int(int32(C.QmlBridge40a792_QmlBridge40a792_QRegisterMetaType()))
}

func QmlBridge_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridge40a792_QmlBridge40a792_QRegisterMetaType2(typeNameC)))
}

func (ptr *QmlBridge) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridge40a792_QmlBridge40a792_QRegisterMetaType2(typeNameC)))
}

func QmlBridge_QmlRegisterType() int {
	return int(int32(C.QmlBridge40a792_QmlBridge40a792_QmlRegisterType()))
}

func (ptr *QmlBridge) QmlRegisterType() int {
	return int(int32(C.QmlBridge40a792_QmlBridge40a792_QmlRegisterType()))
}

func QmlBridge_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlBridge40a792_QmlBridge40a792_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlBridge40a792_QmlBridge40a792_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.QmlBridge40a792___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QmlBridge) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QmlBridge40a792___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge40a792___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList2() unsafe.Pointer {
	return C.QmlBridge40a792___findChildren_newList2(ptr.Pointer())
}

func (ptr *QmlBridge) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge40a792___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList3() unsafe.Pointer {
	return C.QmlBridge40a792___findChildren_newList3(ptr.Pointer())
}

func (ptr *QmlBridge) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge40a792___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList() unsafe.Pointer {
	return C.QmlBridge40a792___findChildren_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge40a792___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __children_newList() unsafe.Pointer {
	return C.QmlBridge40a792___children_newList(ptr.Pointer())
}

func NewQmlBridge(parent std_core.QObject_ITF) *QmlBridge {
	tmpValue := NewQmlBridgeFromPointer(C.QmlBridge40a792_NewQmlBridge(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQmlBridge40a792_DestroyQmlBridge
func callbackQmlBridge40a792_DestroyQmlBridge(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QmlBridge"); signal != nil {
		signal.(func())()
	} else {
		NewQmlBridgeFromPointer(ptr).DestroyQmlBridgeDefault()
	}
}

func (ptr *QmlBridge) ConnectDestroyQmlBridge(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QmlBridge"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDestroyQmlBridge() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QmlBridge")
	}
}

func (ptr *QmlBridge) DestroyQmlBridge() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DestroyQmlBridge(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QmlBridge) DestroyQmlBridgeDefault() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DestroyQmlBridgeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlBridge40a792_Event
func callbackQmlBridge40a792_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QmlBridge) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlBridge40a792_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQmlBridge40a792_EventFilter
func callbackQmlBridge40a792_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QmlBridge) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlBridge40a792_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQmlBridge40a792_ChildEvent
func callbackQmlBridge40a792_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlBridge) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQmlBridge40a792_ConnectNotify
func callbackQmlBridge40a792_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge40a792_CustomEvent
func callbackQmlBridge40a792_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlBridge) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQmlBridge40a792_DeleteLater
func callbackQmlBridge40a792_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQmlBridgeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QmlBridge) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlBridge40a792_Destroyed
func callbackQmlBridge40a792_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackQmlBridge40a792_DisconnectNotify
func callbackQmlBridge40a792_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge40a792_ObjectNameChanged
func callbackQmlBridge40a792_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQmlBridge40a792_TimerEvent
func callbackQmlBridge40a792_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlBridge) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge40a792_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
