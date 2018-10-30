

#pragma once

#ifndef GO_MOC_H
#define GO_MOC_H

#include <stdint.h>

#ifdef __cplusplus
class QmlBridge40a792;
void QmlBridge40a792_QmlBridge40a792_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void QmlBridge40a792_ConnectDisplayTotalBalance(void* ptr);
void QmlBridge40a792_DisconnectDisplayTotalBalance(void* ptr);
void QmlBridge40a792_DisplayTotalBalance(void* ptr, struct Moc_PackedString balance, struct Moc_PackedString balanceUSD);
void QmlBridge40a792_ConnectDisplayAvailableBalance(void* ptr);
void QmlBridge40a792_DisconnectDisplayAvailableBalance(void* ptr);
void QmlBridge40a792_DisplayAvailableBalance(void* ptr, struct Moc_PackedString data);
void QmlBridge40a792_ConnectDisplayLockedBalance(void* ptr);
void QmlBridge40a792_DisconnectDisplayLockedBalance(void* ptr);
void QmlBridge40a792_DisplayLockedBalance(void* ptr, struct Moc_PackedString data);
void QmlBridge40a792_ConnectDisplayAddress(void* ptr);
void QmlBridge40a792_DisconnectDisplayAddress(void* ptr);
void QmlBridge40a792_DisplayAddress(void* ptr, struct Moc_PackedString address, struct Moc_PackedString wallet, char displayFiatConversion);
void QmlBridge40a792_ConnectAddTransactionToList(void* ptr);
void QmlBridge40a792_DisconnectAddTransactionToList(void* ptr);
void QmlBridge40a792_AddTransactionToList(void* ptr, struct Moc_PackedString paymentID, struct Moc_PackedString transactionID, struct Moc_PackedString amount, struct Moc_PackedString confirmations, struct Moc_PackedString time, struct Moc_PackedString number);
void QmlBridge40a792_ConnectAddRemoteNodeToList(void* ptr);
void QmlBridge40a792_DisconnectAddRemoteNodeToList(void* ptr);
void QmlBridge40a792_AddRemoteNodeToList(void* ptr, struct Moc_PackedString nodeURL);
void QmlBridge40a792_ConnectSetSelectedRemoteNode(void* ptr);
void QmlBridge40a792_DisconnectSetSelectedRemoteNode(void* ptr);
void QmlBridge40a792_SetSelectedRemoteNode(void* ptr, int index);
void QmlBridge40a792_ConnectDisplayPopup(void* ptr);
void QmlBridge40a792_DisconnectDisplayPopup(void* ptr);
void QmlBridge40a792_DisplayPopup(void* ptr, struct Moc_PackedString text, int time);
void QmlBridge40a792_ConnectDisplaySyncingInfo(void* ptr);
void QmlBridge40a792_DisconnectDisplaySyncingInfo(void* ptr);
void QmlBridge40a792_DisplaySyncingInfo(void* ptr, struct Moc_PackedString syncing, struct Moc_PackedString syncingInfo);
void QmlBridge40a792_ConnectDisplayErrorDialog(void* ptr);
void QmlBridge40a792_DisconnectDisplayErrorDialog(void* ptr);
void QmlBridge40a792_DisplayErrorDialog(void* ptr, struct Moc_PackedString errorText, struct Moc_PackedString errorInformativeText);
void QmlBridge40a792_ConnectClearTransferAmount(void* ptr);
void QmlBridge40a792_DisconnectClearTransferAmount(void* ptr);
void QmlBridge40a792_ClearTransferAmount(void* ptr);
void QmlBridge40a792_ConnectAskForFusion(void* ptr);
void QmlBridge40a792_DisconnectAskForFusion(void* ptr);
void QmlBridge40a792_AskForFusion(void* ptr);
void QmlBridge40a792_ConnectClearListTransactions(void* ptr);
void QmlBridge40a792_DisconnectClearListTransactions(void* ptr);
void QmlBridge40a792_ClearListTransactions(void* ptr);
void QmlBridge40a792_ConnectDisplayPrivateKeys(void* ptr);
void QmlBridge40a792_DisconnectDisplayPrivateKeys(void* ptr);
void QmlBridge40a792_DisplayPrivateKeys(void* ptr, struct Moc_PackedString filename, struct Moc_PackedString privateViewKey, struct Moc_PackedString privateSpendKey, struct Moc_PackedString walletAddress);
void QmlBridge40a792_ConnectDisplaySeed(void* ptr);
void QmlBridge40a792_DisconnectDisplaySeed(void* ptr);
void QmlBridge40a792_DisplaySeed(void* ptr, struct Moc_PackedString filename, struct Moc_PackedString mnemonicSeed, struct Moc_PackedString walletAddress);
void QmlBridge40a792_ConnectDisplayOpenWalletScreen(void* ptr);
void QmlBridge40a792_DisconnectDisplayOpenWalletScreen(void* ptr);
void QmlBridge40a792_DisplayOpenWalletScreen(void* ptr);
void QmlBridge40a792_ConnectDisplayMainWalletScreen(void* ptr);
void QmlBridge40a792_DisconnectDisplayMainWalletScreen(void* ptr);
void QmlBridge40a792_DisplayMainWalletScreen(void* ptr);
void QmlBridge40a792_ConnectFinishedLoadingWalletd(void* ptr);
void QmlBridge40a792_DisconnectFinishedLoadingWalletd(void* ptr);
void QmlBridge40a792_FinishedLoadingWalletd(void* ptr);
void QmlBridge40a792_ConnectFinishedCreatingWallet(void* ptr);
void QmlBridge40a792_DisconnectFinishedCreatingWallet(void* ptr);
void QmlBridge40a792_FinishedCreatingWallet(void* ptr);
void QmlBridge40a792_ConnectFinishedSendingTransaction(void* ptr);
void QmlBridge40a792_DisconnectFinishedSendingTransaction(void* ptr);
void QmlBridge40a792_FinishedSendingTransaction(void* ptr);
void QmlBridge40a792_ConnectDisplayPathToPreviousWallet(void* ptr);
void QmlBridge40a792_DisconnectDisplayPathToPreviousWallet(void* ptr);
void QmlBridge40a792_DisplayPathToPreviousWallet(void* ptr, struct Moc_PackedString pathToPreviousWallet);
void QmlBridge40a792_ConnectDisplayWalletCreationLocation(void* ptr);
void QmlBridge40a792_DisconnectDisplayWalletCreationLocation(void* ptr);
void QmlBridge40a792_DisplayWalletCreationLocation(void* ptr, struct Moc_PackedString walletLocation);
void QmlBridge40a792_ConnectDisplayUseRemoteNode(void* ptr);
void QmlBridge40a792_DisconnectDisplayUseRemoteNode(void* ptr);
void QmlBridge40a792_DisplayUseRemoteNode(void* ptr, char useRemote);
void QmlBridge40a792_ConnectHideSettingsScreen(void* ptr);
void QmlBridge40a792_DisconnectHideSettingsScreen(void* ptr);
void QmlBridge40a792_HideSettingsScreen(void* ptr);
void QmlBridge40a792_ConnectDisplaySettingsScreen(void* ptr);
void QmlBridge40a792_DisconnectDisplaySettingsScreen(void* ptr);
void QmlBridge40a792_DisplaySettingsScreen(void* ptr);
void QmlBridge40a792_ConnectDisplaySettingsValues(void* ptr);
void QmlBridge40a792_DisconnectDisplaySettingsValues(void* ptr);
void QmlBridge40a792_DisplaySettingsValues(void* ptr, char displayFiat);
void QmlBridge40a792_ConnectDisplaySettingsRemoteDaemonInfo(void* ptr);
void QmlBridge40a792_DisconnectDisplaySettingsRemoteDaemonInfo(void* ptr);
void QmlBridge40a792_DisplaySettingsRemoteDaemonInfo(void* ptr, struct Moc_PackedString remoteNodeAddress, struct Moc_PackedString remoteNodePort);
void QmlBridge40a792_ConnectDisplayFullBalanceInTransferAmount(void* ptr);
void QmlBridge40a792_DisconnectDisplayFullBalanceInTransferAmount(void* ptr);
void QmlBridge40a792_DisplayFullBalanceInTransferAmount(void* ptr, struct Moc_PackedString fullBalance);
void QmlBridge40a792_ConnectDisplayDefaultFee(void* ptr);
void QmlBridge40a792_DisconnectDisplayDefaultFee(void* ptr);
void QmlBridge40a792_DisplayDefaultFee(void* ptr, struct Moc_PackedString fee);
void QmlBridge40a792_ConnectDisplayNodeFee(void* ptr);
void QmlBridge40a792_DisconnectDisplayNodeFee(void* ptr);
void QmlBridge40a792_DisplayNodeFee(void* ptr, struct Moc_PackedString nodeFee);
void QmlBridge40a792_ConnectUpdateConfirmationsOfTransaction(void* ptr);
void QmlBridge40a792_DisconnectUpdateConfirmationsOfTransaction(void* ptr);
void QmlBridge40a792_UpdateConfirmationsOfTransaction(void* ptr, int index, struct Moc_PackedString confirmations);
void QmlBridge40a792_ConnectDisplayInfoDialog(void* ptr);
void QmlBridge40a792_DisconnectDisplayInfoDialog(void* ptr);
void QmlBridge40a792_DisplayInfoDialog(void* ptr);
void QmlBridge40a792_ConnectAddSavedAddressToList(void* ptr);
void QmlBridge40a792_DisconnectAddSavedAddressToList(void* ptr);
void QmlBridge40a792_AddSavedAddressToList(void* ptr, int dbID, struct Moc_PackedString name, struct Moc_PackedString address, struct Moc_PackedString paymentID);
void QmlBridge40a792_Log(void* ptr, struct Moc_PackedString msg);
void QmlBridge40a792_ClickedButtonExplorer(void* ptr, struct Moc_PackedString transactionID);
void QmlBridge40a792_GoToWebsite(void* ptr, struct Moc_PackedString url);
void QmlBridge40a792_ClickedButtonCopyTx(void* ptr, struct Moc_PackedString transactionID);
void QmlBridge40a792_ClickedButtonCopyAddress(void* ptr);
void QmlBridge40a792_ClickedButtonCopyKeys(void* ptr);
void QmlBridge40a792_ClickedButtonCopy(void* ptr, struct Moc_PackedString stringToCopy);
void QmlBridge40a792_ClickedButtonSend(void* ptr, struct Moc_PackedString transferAddress, struct Moc_PackedString transferAmount, struct Moc_PackedString transferPaymentID, struct Moc_PackedString transferFee);
void QmlBridge40a792_ClickedButtonBackupWallet(void* ptr);
void QmlBridge40a792_ClickedCloseWallet(void* ptr);
void QmlBridge40a792_ClickedButtonOpen(void* ptr, struct Moc_PackedString pathToWallet, struct Moc_PackedString passwordWallet);
void QmlBridge40a792_ClickedButtonCreate(void* ptr, struct Moc_PackedString filenameWallet, struct Moc_PackedString passwordWallet, struct Moc_PackedString confirmPasswordWallet);
void QmlBridge40a792_ClickedButtonImport(void* ptr, struct Moc_PackedString filenameWallet, struct Moc_PackedString passwordWallet, struct Moc_PackedString privateViewKey, struct Moc_PackedString privateSpendKey, struct Moc_PackedString mnemonicSeed, struct Moc_PackedString confirmPasswordWallet, struct Moc_PackedString scanHeight);
void QmlBridge40a792_ChoseRemote(void* ptr, char remote);
void QmlBridge40a792_SelectedRemoteNode(void* ptr, int index);
struct Moc_PackedString QmlBridge40a792_GetTransferAmountUSD(void* ptr, struct Moc_PackedString amountTRTL);
void QmlBridge40a792_ClickedCloseSettings(void* ptr);
void QmlBridge40a792_ClickedSettingsButton(void* ptr);
void QmlBridge40a792_ChoseDisplayFiat(void* ptr, char displayFiat);
void QmlBridge40a792_ChoseCheckpoints(void* ptr, char checkpoints);
void QmlBridge40a792_SaveRemoteDaemonInfo(void* ptr, struct Moc_PackedString daemonAddress, struct Moc_PackedString daemonPort);
void QmlBridge40a792_ResetRemoteDaemonInfo(void* ptr);
void QmlBridge40a792_GetFullBalanceAndDisplayInTransferAmount(void* ptr, struct Moc_PackedString transferFee);
void QmlBridge40a792_GetDefaultFeeAndDisplay(void* ptr);
void QmlBridge40a792_LimitDisplayTransactions(void* ptr, char limit);
struct Moc_PackedString QmlBridge40a792_GetVersion(void* ptr);
struct Moc_PackedString QmlBridge40a792_GetNewVersion(void* ptr);
struct Moc_PackedString QmlBridge40a792_GetNewVersionURL(void* ptr);
void QmlBridge40a792_OptimizeWalletWithFusion(void* ptr);
void QmlBridge40a792_SaveAddress(void* ptr, struct Moc_PackedString name, struct Moc_PackedString address, struct Moc_PackedString paymentID);
void QmlBridge40a792_FillListSavedAddresses(void* ptr);
void QmlBridge40a792_DeleteSavedAddress(void* ptr, int dbID);
void QmlBridge40a792_RegisterToGo(void* ptr, void* object);
void QmlBridge40a792_DeregisterToGo(void* ptr, struct Moc_PackedString objectName);
int QmlBridge40a792_QmlBridge40a792_QRegisterMetaType();
int QmlBridge40a792_QmlBridge40a792_QRegisterMetaType2(char* typeName);
int QmlBridge40a792_QmlBridge40a792_QmlRegisterType();
int QmlBridge40a792_QmlBridge40a792_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QmlBridge40a792___dynamicPropertyNames_atList(void* ptr, int i);
void QmlBridge40a792___dynamicPropertyNames_setList(void* ptr, void* i);
void* QmlBridge40a792___dynamicPropertyNames_newList(void* ptr);
void* QmlBridge40a792___findChildren_atList2(void* ptr, int i);
void QmlBridge40a792___findChildren_setList2(void* ptr, void* i);
void* QmlBridge40a792___findChildren_newList2(void* ptr);
void* QmlBridge40a792___findChildren_atList3(void* ptr, int i);
void QmlBridge40a792___findChildren_setList3(void* ptr, void* i);
void* QmlBridge40a792___findChildren_newList3(void* ptr);
void* QmlBridge40a792___findChildren_atList(void* ptr, int i);
void QmlBridge40a792___findChildren_setList(void* ptr, void* i);
void* QmlBridge40a792___findChildren_newList(void* ptr);
void* QmlBridge40a792___children_atList(void* ptr, int i);
void QmlBridge40a792___children_setList(void* ptr, void* i);
void* QmlBridge40a792___children_newList(void* ptr);
void* QmlBridge40a792_NewQmlBridge(void* parent);
void QmlBridge40a792_DestroyQmlBridge(void* ptr);
void QmlBridge40a792_DestroyQmlBridgeDefault(void* ptr);
char QmlBridge40a792_EventDefault(void* ptr, void* e);
char QmlBridge40a792_EventFilterDefault(void* ptr, void* watched, void* event);
void QmlBridge40a792_ChildEventDefault(void* ptr, void* event);
void QmlBridge40a792_ConnectNotifyDefault(void* ptr, void* sign);
void QmlBridge40a792_CustomEventDefault(void* ptr, void* event);
void QmlBridge40a792_DeleteLaterDefault(void* ptr);
void QmlBridge40a792_DisconnectNotifyDefault(void* ptr, void* sign);
void QmlBridge40a792_TimerEventDefault(void* ptr, void* event);
;

#ifdef __cplusplus
}
#endif

#endif