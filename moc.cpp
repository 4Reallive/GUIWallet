

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>


class QmlBridge40a792: public QObject
{
Q_OBJECT
public:
	QmlBridge40a792(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QmlBridge40a792_QmlBridge40a792_QRegisterMetaType();QmlBridge40a792_QmlBridge40a792_QRegisterMetaTypes();callbackQmlBridge40a792_Constructor(this);};
	void Signal_DisplayTotalBalance(QString balance, QString balanceUSD) { QByteArray t8dfa30 = balance.toUtf8(); Moc_PackedString balancePacked = { const_cast<char*>(t8dfa30.prepend("WHITESPACE").constData()+10), t8dfa30.size()-10 };QByteArray t7a7b4e = balanceUSD.toUtf8(); Moc_PackedString balanceUSDPacked = { const_cast<char*>(t7a7b4e.prepend("WHITESPACE").constData()+10), t7a7b4e.size()-10 };callbackQmlBridge40a792_DisplayTotalBalance(this, balancePacked, balanceUSDPacked); };
	void Signal_DisplayAvailableBalance(QString data) { QByteArray ta17c9a = data.toUtf8(); Moc_PackedString dataPacked = { const_cast<char*>(ta17c9a.prepend("WHITESPACE").constData()+10), ta17c9a.size()-10 };callbackQmlBridge40a792_DisplayAvailableBalance(this, dataPacked); };
	void Signal_DisplayLockedBalance(QString data) { QByteArray ta17c9a = data.toUtf8(); Moc_PackedString dataPacked = { const_cast<char*>(ta17c9a.prepend("WHITESPACE").constData()+10), ta17c9a.size()-10 };callbackQmlBridge40a792_DisplayLockedBalance(this, dataPacked); };
	void Signal_DisplayAddress(QString address, QString wallet, bool displayFiatConversion) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };QByteArray tc307b6 = wallet.toUtf8(); Moc_PackedString walletPacked = { const_cast<char*>(tc307b6.prepend("WHITESPACE").constData()+10), tc307b6.size()-10 };callbackQmlBridge40a792_DisplayAddress(this, addressPacked, walletPacked, displayFiatConversion); };
	void Signal_AddTransactionToList(QString paymentID, QString transactionID, QString amount, QString confirmations, QString time, QString number) { QByteArray tc240c1 = paymentID.toUtf8(); Moc_PackedString paymentIDPacked = { const_cast<char*>(tc240c1.prepend("WHITESPACE").constData()+10), tc240c1.size()-10 };QByteArray tec2ac1 = transactionID.toUtf8(); Moc_PackedString transactionIDPacked = { const_cast<char*>(tec2ac1.prepend("WHITESPACE").constData()+10), tec2ac1.size()-10 };QByteArray t9cb6ff = amount.toUtf8(); Moc_PackedString amountPacked = { const_cast<char*>(t9cb6ff.prepend("WHITESPACE").constData()+10), t9cb6ff.size()-10 };QByteArray t8499a0 = confirmations.toUtf8(); Moc_PackedString confirmationsPacked = { const_cast<char*>(t8499a0.prepend("WHITESPACE").constData()+10), t8499a0.size()-10 };QByteArray t714eea = time.toUtf8(); Moc_PackedString timePacked = { const_cast<char*>(t714eea.prepend("WHITESPACE").constData()+10), t714eea.size()-10 };QByteArray t53b0a1 = number.toUtf8(); Moc_PackedString numberPacked = { const_cast<char*>(t53b0a1.prepend("WHITESPACE").constData()+10), t53b0a1.size()-10 };callbackQmlBridge40a792_AddTransactionToList(this, paymentIDPacked, transactionIDPacked, amountPacked, confirmationsPacked, timePacked, numberPacked); };
	void Signal_AddRemoteNodeToList(QString nodeURL) { QByteArray t04d4d0 = nodeURL.toUtf8(); Moc_PackedString nodeURLPacked = { const_cast<char*>(t04d4d0.prepend("WHITESPACE").constData()+10), t04d4d0.size()-10 };callbackQmlBridge40a792_AddRemoteNodeToList(this, nodeURLPacked); };
	void Signal_SetSelectedRemoteNode(qint32 index) { callbackQmlBridge40a792_SetSelectedRemoteNode(this, index); };
	void Signal_DisplayPopup(QString text, qint32 time) { QByteArray t372ea0 = text.toUtf8(); Moc_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQmlBridge40a792_DisplayPopup(this, textPacked, time); };
	void Signal_DisplaySyncingInfo(QString syncing, QString syncingInfo) { QByteArray t72df87 = syncing.toUtf8(); Moc_PackedString syncingPacked = { const_cast<char*>(t72df87.prepend("WHITESPACE").constData()+10), t72df87.size()-10 };QByteArray t834023 = syncingInfo.toUtf8(); Moc_PackedString syncingInfoPacked = { const_cast<char*>(t834023.prepend("WHITESPACE").constData()+10), t834023.size()-10 };callbackQmlBridge40a792_DisplaySyncingInfo(this, syncingPacked, syncingInfoPacked); };
	void Signal_DisplayErrorDialog(QString errorText, QString errorInformativeText) { QByteArray t5f7b8e = errorText.toUtf8(); Moc_PackedString errorTextPacked = { const_cast<char*>(t5f7b8e.prepend("WHITESPACE").constData()+10), t5f7b8e.size()-10 };QByteArray tae90fd = errorInformativeText.toUtf8(); Moc_PackedString errorInformativeTextPacked = { const_cast<char*>(tae90fd.prepend("WHITESPACE").constData()+10), tae90fd.size()-10 };callbackQmlBridge40a792_DisplayErrorDialog(this, errorTextPacked, errorInformativeTextPacked); };
	void Signal_ClearTransferAmount() { callbackQmlBridge40a792_ClearTransferAmount(this); };
	void Signal_AskForFusion() { callbackQmlBridge40a792_AskForFusion(this); };
	void Signal_ClearListTransactions() { callbackQmlBridge40a792_ClearListTransactions(this); };
	void Signal_DisplayPrivateKeys(QString filename, QString privateViewKey, QString privateSpendKey, QString walletAddress) { QByteArray t08deae = filename.toUtf8(); Moc_PackedString filenamePacked = { const_cast<char*>(t08deae.prepend("WHITESPACE").constData()+10), t08deae.size()-10 };QByteArray tc634db = privateViewKey.toUtf8(); Moc_PackedString privateViewKeyPacked = { const_cast<char*>(tc634db.prepend("WHITESPACE").constData()+10), tc634db.size()-10 };QByteArray t1f6ec2 = privateSpendKey.toUtf8(); Moc_PackedString privateSpendKeyPacked = { const_cast<char*>(t1f6ec2.prepend("WHITESPACE").constData()+10), t1f6ec2.size()-10 };QByteArray t208cac = walletAddress.toUtf8(); Moc_PackedString walletAddressPacked = { const_cast<char*>(t208cac.prepend("WHITESPACE").constData()+10), t208cac.size()-10 };callbackQmlBridge40a792_DisplayPrivateKeys(this, filenamePacked, privateViewKeyPacked, privateSpendKeyPacked, walletAddressPacked); };
	void Signal_DisplaySeed(QString filename, QString mnemonicSeed, QString walletAddress) { QByteArray t08deae = filename.toUtf8(); Moc_PackedString filenamePacked = { const_cast<char*>(t08deae.prepend("WHITESPACE").constData()+10), t08deae.size()-10 };QByteArray tde366e = mnemonicSeed.toUtf8(); Moc_PackedString mnemonicSeedPacked = { const_cast<char*>(tde366e.prepend("WHITESPACE").constData()+10), tde366e.size()-10 };QByteArray t208cac = walletAddress.toUtf8(); Moc_PackedString walletAddressPacked = { const_cast<char*>(t208cac.prepend("WHITESPACE").constData()+10), t208cac.size()-10 };callbackQmlBridge40a792_DisplaySeed(this, filenamePacked, mnemonicSeedPacked, walletAddressPacked); };
	void Signal_DisplayOpenWalletScreen() { callbackQmlBridge40a792_DisplayOpenWalletScreen(this); };
	void Signal_DisplayMainWalletScreen() { callbackQmlBridge40a792_DisplayMainWalletScreen(this); };
	void Signal_FinishedLoadingWalletd() { callbackQmlBridge40a792_FinishedLoadingWalletd(this); };
	void Signal_FinishedCreatingWallet() { callbackQmlBridge40a792_FinishedCreatingWallet(this); };
	void Signal_FinishedSendingTransaction() { callbackQmlBridge40a792_FinishedSendingTransaction(this); };
	void Signal_DisplayPathToPreviousWallet(QString pathToPreviousWallet) { QByteArray t70f286 = pathToPreviousWallet.toUtf8(); Moc_PackedString pathToPreviousWalletPacked = { const_cast<char*>(t70f286.prepend("WHITESPACE").constData()+10), t70f286.size()-10 };callbackQmlBridge40a792_DisplayPathToPreviousWallet(this, pathToPreviousWalletPacked); };
	void Signal_DisplayWalletCreationLocation(QString walletLocation) { QByteArray tabe26e = walletLocation.toUtf8(); Moc_PackedString walletLocationPacked = { const_cast<char*>(tabe26e.prepend("WHITESPACE").constData()+10), tabe26e.size()-10 };callbackQmlBridge40a792_DisplayWalletCreationLocation(this, walletLocationPacked); };
	void Signal_DisplayUseRemoteNode(bool useRemote) { callbackQmlBridge40a792_DisplayUseRemoteNode(this, useRemote); };
	void Signal_HideSettingsScreen() { callbackQmlBridge40a792_HideSettingsScreen(this); };
	void Signal_DisplaySettingsScreen() { callbackQmlBridge40a792_DisplaySettingsScreen(this); };
	void Signal_DisplaySettingsValues(bool displayFiat) { callbackQmlBridge40a792_DisplaySettingsValues(this, displayFiat); };
	void Signal_DisplaySettingsRemoteDaemonInfo(QString remoteNodeAddress, QString remoteNodePort) { QByteArray tc6f2b4 = remoteNodeAddress.toUtf8(); Moc_PackedString remoteNodeAddressPacked = { const_cast<char*>(tc6f2b4.prepend("WHITESPACE").constData()+10), tc6f2b4.size()-10 };QByteArray t1b1377 = remoteNodePort.toUtf8(); Moc_PackedString remoteNodePortPacked = { const_cast<char*>(t1b1377.prepend("WHITESPACE").constData()+10), t1b1377.size()-10 };callbackQmlBridge40a792_DisplaySettingsRemoteDaemonInfo(this, remoteNodeAddressPacked, remoteNodePortPacked); };
	void Signal_DisplayFullBalanceInTransferAmount(QString fullBalance) { QByteArray tccc49e = fullBalance.toUtf8(); Moc_PackedString fullBalancePacked = { const_cast<char*>(tccc49e.prepend("WHITESPACE").constData()+10), tccc49e.size()-10 };callbackQmlBridge40a792_DisplayFullBalanceInTransferAmount(this, fullBalancePacked); };
	void Signal_DisplayDefaultFee(QString fee) { QByteArray t9c15cd = fee.toUtf8(); Moc_PackedString feePacked = { const_cast<char*>(t9c15cd.prepend("WHITESPACE").constData()+10), t9c15cd.size()-10 };callbackQmlBridge40a792_DisplayDefaultFee(this, feePacked); };
	void Signal_DisplayNodeFee(QString nodeFee) { QByteArray t96217f = nodeFee.toUtf8(); Moc_PackedString nodeFeePacked = { const_cast<char*>(t96217f.prepend("WHITESPACE").constData()+10), t96217f.size()-10 };callbackQmlBridge40a792_DisplayNodeFee(this, nodeFeePacked); };
	void Signal_UpdateConfirmationsOfTransaction(qint32 index, QString confirmations) { QByteArray t8499a0 = confirmations.toUtf8(); Moc_PackedString confirmationsPacked = { const_cast<char*>(t8499a0.prepend("WHITESPACE").constData()+10), t8499a0.size()-10 };callbackQmlBridge40a792_UpdateConfirmationsOfTransaction(this, index, confirmationsPacked); };
	void Signal_DisplayInfoDialog() { callbackQmlBridge40a792_DisplayInfoDialog(this); };
	void Signal_AddSavedAddressToList(qint32 dbID, QString name, QString address, QString paymentID) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };QByteArray tc240c1 = paymentID.toUtf8(); Moc_PackedString paymentIDPacked = { const_cast<char*>(tc240c1.prepend("WHITESPACE").constData()+10), tc240c1.size()-10 };callbackQmlBridge40a792_AddSavedAddressToList(this, dbID, namePacked, addressPacked, paymentIDPacked); };
	 ~QmlBridge40a792() { callbackQmlBridge40a792_DestroyQmlBridge(this); };
	bool event(QEvent * e) { return callbackQmlBridge40a792_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQmlBridge40a792_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQmlBridge40a792_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQmlBridge40a792_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQmlBridge40a792_CustomEvent(this, event); };
	void deleteLater() { callbackQmlBridge40a792_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQmlBridge40a792_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQmlBridge40a792_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQmlBridge40a792_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQmlBridge40a792_TimerEvent(this, event); };
signals:
	void displayTotalBalance(QString balance, QString balanceUSD);
	void displayAvailableBalance(QString data);
	void displayLockedBalance(QString data);
	void displayAddress(QString address, QString wallet, bool displayFiatConversion);
	void addTransactionToList(QString paymentID, QString transactionID, QString amount, QString confirmations, QString time, QString number);
	void addRemoteNodeToList(QString nodeURL);
	void setSelectedRemoteNode(qint32 index);
	void displayPopup(QString text, qint32 time);
	void displaySyncingInfo(QString syncing, QString syncingInfo);
	void displayErrorDialog(QString errorText, QString errorInformativeText);
	void clearTransferAmount();
	void askForFusion();
	void clearListTransactions();
	void displayPrivateKeys(QString filename, QString privateViewKey, QString privateSpendKey, QString walletAddress);
	void displaySeed(QString filename, QString mnemonicSeed, QString walletAddress);
	void displayOpenWalletScreen();
	void displayMainWalletScreen();
	void finishedLoadingWalletd();
	void finishedCreatingWallet();
	void finishedSendingTransaction();
	void displayPathToPreviousWallet(QString pathToPreviousWallet);
	void displayWalletCreationLocation(QString walletLocation);
	void displayUseRemoteNode(bool useRemote);
	void hideSettingsScreen();
	void displaySettingsScreen();
	void displaySettingsValues(bool displayFiat);
	void displaySettingsRemoteDaemonInfo(QString remoteNodeAddress, QString remoteNodePort);
	void displayFullBalanceInTransferAmount(QString fullBalance);
	void displayDefaultFee(QString fee);
	void displayNodeFee(QString nodeFee);
	void updateConfirmationsOfTransaction(qint32 index, QString confirmations);
	void displayInfoDialog();
	void addSavedAddressToList(qint32 dbID, QString name, QString address, QString paymentID);
public slots:
	void log(QString msg) { QByteArray t19f34e = msg.toUtf8(); Moc_PackedString msgPacked = { const_cast<char*>(t19f34e.prepend("WHITESPACE").constData()+10), t19f34e.size()-10 };callbackQmlBridge40a792_Log(this, msgPacked); };
	void clickedButtonExplorer(QString transactionID) { QByteArray tec2ac1 = transactionID.toUtf8(); Moc_PackedString transactionIDPacked = { const_cast<char*>(tec2ac1.prepend("WHITESPACE").constData()+10), tec2ac1.size()-10 };callbackQmlBridge40a792_ClickedButtonExplorer(this, transactionIDPacked); };
	void goToWebsite(QString url) { QByteArray t817363 = url.toUtf8(); Moc_PackedString urlPacked = { const_cast<char*>(t817363.prepend("WHITESPACE").constData()+10), t817363.size()-10 };callbackQmlBridge40a792_GoToWebsite(this, urlPacked); };
	void clickedButtonCopyTx(QString transactionID) { QByteArray tec2ac1 = transactionID.toUtf8(); Moc_PackedString transactionIDPacked = { const_cast<char*>(tec2ac1.prepend("WHITESPACE").constData()+10), tec2ac1.size()-10 };callbackQmlBridge40a792_ClickedButtonCopyTx(this, transactionIDPacked); };
	void clickedButtonCopyAddress() { callbackQmlBridge40a792_ClickedButtonCopyAddress(this); };
	void clickedButtonCopyKeys() { callbackQmlBridge40a792_ClickedButtonCopyKeys(this); };
	void clickedButtonCopy(QString stringToCopy) { QByteArray te009d3 = stringToCopy.toUtf8(); Moc_PackedString stringToCopyPacked = { const_cast<char*>(te009d3.prepend("WHITESPACE").constData()+10), te009d3.size()-10 };callbackQmlBridge40a792_ClickedButtonCopy(this, stringToCopyPacked); };
	void clickedButtonSend(QString transferAddress, QString transferAmount, QString transferPaymentID, QString transferFee) { QByteArray tb2289f = transferAddress.toUtf8(); Moc_PackedString transferAddressPacked = { const_cast<char*>(tb2289f.prepend("WHITESPACE").constData()+10), tb2289f.size()-10 };QByteArray t307ef6 = transferAmount.toUtf8(); Moc_PackedString transferAmountPacked = { const_cast<char*>(t307ef6.prepend("WHITESPACE").constData()+10), t307ef6.size()-10 };QByteArray tc16f3f = transferPaymentID.toUtf8(); Moc_PackedString transferPaymentIDPacked = { const_cast<char*>(tc16f3f.prepend("WHITESPACE").constData()+10), tc16f3f.size()-10 };QByteArray t1a861b = transferFee.toUtf8(); Moc_PackedString transferFeePacked = { const_cast<char*>(t1a861b.prepend("WHITESPACE").constData()+10), t1a861b.size()-10 };callbackQmlBridge40a792_ClickedButtonSend(this, transferAddressPacked, transferAmountPacked, transferPaymentIDPacked, transferFeePacked); };
	void clickedButtonBackupWallet() { callbackQmlBridge40a792_ClickedButtonBackupWallet(this); };
	void clickedCloseWallet() { callbackQmlBridge40a792_ClickedCloseWallet(this); };
	void clickedButtonOpen(QString pathToWallet, QString passwordWallet) { QByteArray t6c0b5c = pathToWallet.toUtf8(); Moc_PackedString pathToWalletPacked = { const_cast<char*>(t6c0b5c.prepend("WHITESPACE").constData()+10), t6c0b5c.size()-10 };QByteArray t1b4d89 = passwordWallet.toUtf8(); Moc_PackedString passwordWalletPacked = { const_cast<char*>(t1b4d89.prepend("WHITESPACE").constData()+10), t1b4d89.size()-10 };callbackQmlBridge40a792_ClickedButtonOpen(this, pathToWalletPacked, passwordWalletPacked); };
	void clickedButtonCreate(QString filenameWallet, QString passwordWallet, QString confirmPasswordWallet) { QByteArray t794b1e = filenameWallet.toUtf8(); Moc_PackedString filenameWalletPacked = { const_cast<char*>(t794b1e.prepend("WHITESPACE").constData()+10), t794b1e.size()-10 };QByteArray t1b4d89 = passwordWallet.toUtf8(); Moc_PackedString passwordWalletPacked = { const_cast<char*>(t1b4d89.prepend("WHITESPACE").constData()+10), t1b4d89.size()-10 };QByteArray taf1d27 = confirmPasswordWallet.toUtf8(); Moc_PackedString confirmPasswordWalletPacked = { const_cast<char*>(taf1d27.prepend("WHITESPACE").constData()+10), taf1d27.size()-10 };callbackQmlBridge40a792_ClickedButtonCreate(this, filenameWalletPacked, passwordWalletPacked, confirmPasswordWalletPacked); };
	void clickedButtonImport(QString filenameWallet, QString passwordWallet, QString privateViewKey, QString privateSpendKey, QString mnemonicSeed, QString confirmPasswordWallet, QString scanHeight) { QByteArray t794b1e = filenameWallet.toUtf8(); Moc_PackedString filenameWalletPacked = { const_cast<char*>(t794b1e.prepend("WHITESPACE").constData()+10), t794b1e.size()-10 };QByteArray t1b4d89 = passwordWallet.toUtf8(); Moc_PackedString passwordWalletPacked = { const_cast<char*>(t1b4d89.prepend("WHITESPACE").constData()+10), t1b4d89.size()-10 };QByteArray tc634db = privateViewKey.toUtf8(); Moc_PackedString privateViewKeyPacked = { const_cast<char*>(tc634db.prepend("WHITESPACE").constData()+10), tc634db.size()-10 };QByteArray t1f6ec2 = privateSpendKey.toUtf8(); Moc_PackedString privateSpendKeyPacked = { const_cast<char*>(t1f6ec2.prepend("WHITESPACE").constData()+10), t1f6ec2.size()-10 };QByteArray tde366e = mnemonicSeed.toUtf8(); Moc_PackedString mnemonicSeedPacked = { const_cast<char*>(tde366e.prepend("WHITESPACE").constData()+10), tde366e.size()-10 };QByteArray taf1d27 = confirmPasswordWallet.toUtf8(); Moc_PackedString confirmPasswordWalletPacked = { const_cast<char*>(taf1d27.prepend("WHITESPACE").constData()+10), taf1d27.size()-10 };QByteArray t32972b = scanHeight.toUtf8(); Moc_PackedString scanHeightPacked = { const_cast<char*>(t32972b.prepend("WHITESPACE").constData()+10), t32972b.size()-10 };callbackQmlBridge40a792_ClickedButtonImport(this, filenameWalletPacked, passwordWalletPacked, privateViewKeyPacked, privateSpendKeyPacked, mnemonicSeedPacked, confirmPasswordWalletPacked, scanHeightPacked); };
	void choseRemote(bool remote) { callbackQmlBridge40a792_ChoseRemote(this, remote); };
	void selectedRemoteNode(qint32 index) { callbackQmlBridge40a792_SelectedRemoteNode(this, index); };
	QString getTransferAmountUSD(QString amountTRTL) { QByteArray te38fb0 = amountTRTL.toUtf8(); Moc_PackedString amountTRTLPacked = { const_cast<char*>(te38fb0.prepend("WHITESPACE").constData()+10), te38fb0.size()-10 };return ({ Moc_PackedString tempVal = callbackQmlBridge40a792_GetTransferAmountUSD(this, amountTRTLPacked); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void clickedCloseSettings() { callbackQmlBridge40a792_ClickedCloseSettings(this); };
	void clickedSettingsButton() { callbackQmlBridge40a792_ClickedSettingsButton(this); };
	void choseDisplayFiat(bool displayFiat) { callbackQmlBridge40a792_ChoseDisplayFiat(this, displayFiat); };
	void choseCheckpoints(bool checkpoints) { callbackQmlBridge40a792_ChoseCheckpoints(this, checkpoints); };
	void saveRemoteDaemonInfo(QString daemonAddress, QString daemonPort) { QByteArray td2b37c = daemonAddress.toUtf8(); Moc_PackedString daemonAddressPacked = { const_cast<char*>(td2b37c.prepend("WHITESPACE").constData()+10), td2b37c.size()-10 };QByteArray tcb21fe = daemonPort.toUtf8(); Moc_PackedString daemonPortPacked = { const_cast<char*>(tcb21fe.prepend("WHITESPACE").constData()+10), tcb21fe.size()-10 };callbackQmlBridge40a792_SaveRemoteDaemonInfo(this, daemonAddressPacked, daemonPortPacked); };
	void resetRemoteDaemonInfo() { callbackQmlBridge40a792_ResetRemoteDaemonInfo(this); };
	void getFullBalanceAndDisplayInTransferAmount(QString transferFee) { QByteArray t1a861b = transferFee.toUtf8(); Moc_PackedString transferFeePacked = { const_cast<char*>(t1a861b.prepend("WHITESPACE").constData()+10), t1a861b.size()-10 };callbackQmlBridge40a792_GetFullBalanceAndDisplayInTransferAmount(this, transferFeePacked); };
	void getDefaultFeeAndDisplay() { callbackQmlBridge40a792_GetDefaultFeeAndDisplay(this); };
	void limitDisplayTransactions(bool limit) { callbackQmlBridge40a792_LimitDisplayTransactions(this, limit); };
	QString getVersion() { return ({ Moc_PackedString tempVal = callbackQmlBridge40a792_GetVersion(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString getNewVersion() { return ({ Moc_PackedString tempVal = callbackQmlBridge40a792_GetNewVersion(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString getNewVersionURL() { return ({ Moc_PackedString tempVal = callbackQmlBridge40a792_GetNewVersionURL(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void optimizeWalletWithFusion() { callbackQmlBridge40a792_OptimizeWalletWithFusion(this); };
	void saveAddress(QString name, QString address, QString paymentID) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };QByteArray tc240c1 = paymentID.toUtf8(); Moc_PackedString paymentIDPacked = { const_cast<char*>(tc240c1.prepend("WHITESPACE").constData()+10), tc240c1.size()-10 };callbackQmlBridge40a792_SaveAddress(this, namePacked, addressPacked, paymentIDPacked); };
	void fillListSavedAddresses() { callbackQmlBridge40a792_FillListSavedAddresses(this); };
	void deleteSavedAddress(qint32 dbID) { callbackQmlBridge40a792_DeleteSavedAddress(this, dbID); };
	void registerToGo(QObject* object) { callbackQmlBridge40a792_RegisterToGo(this, object); };
	void deregisterToGo(QString objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQmlBridge40a792_DeregisterToGo(this, objectNamePacked); };
private:
};

Q_DECLARE_METATYPE(QmlBridge40a792*)


void QmlBridge40a792_QmlBridge40a792_QRegisterMetaTypes() {
}

void QmlBridge40a792_ConnectDisplayTotalBalance(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::displayTotalBalance), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::Signal_DisplayTotalBalance));
}

void QmlBridge40a792_DisconnectDisplayTotalBalance(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::displayTotalBalance), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::Signal_DisplayTotalBalance));
}

void QmlBridge40a792_DisplayTotalBalance(void* ptr, struct Moc_PackedString balance, struct Moc_PackedString balanceUSD)
{
	static_cast<QmlBridge40a792*>(ptr)->displayTotalBalance(QString::fromUtf8(balance.data, balance.len), QString::fromUtf8(balanceUSD.data, balanceUSD.len));
}

void QmlBridge40a792_ConnectDisplayAvailableBalance(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayAvailableBalance), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayAvailableBalance));
}

void QmlBridge40a792_DisconnectDisplayAvailableBalance(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayAvailableBalance), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayAvailableBalance));
}

void QmlBridge40a792_DisplayAvailableBalance(void* ptr, struct Moc_PackedString data)
{
	static_cast<QmlBridge40a792*>(ptr)->displayAvailableBalance(QString::fromUtf8(data.data, data.len));
}

void QmlBridge40a792_ConnectDisplayLockedBalance(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayLockedBalance), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayLockedBalance));
}

void QmlBridge40a792_DisconnectDisplayLockedBalance(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayLockedBalance), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayLockedBalance));
}

void QmlBridge40a792_DisplayLockedBalance(void* ptr, struct Moc_PackedString data)
{
	static_cast<QmlBridge40a792*>(ptr)->displayLockedBalance(QString::fromUtf8(data.data, data.len));
}

void QmlBridge40a792_ConnectDisplayAddress(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, bool)>(&QmlBridge40a792::displayAddress), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, bool)>(&QmlBridge40a792::Signal_DisplayAddress));
}

void QmlBridge40a792_DisconnectDisplayAddress(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, bool)>(&QmlBridge40a792::displayAddress), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, bool)>(&QmlBridge40a792::Signal_DisplayAddress));
}

void QmlBridge40a792_DisplayAddress(void* ptr, struct Moc_PackedString address, struct Moc_PackedString wallet, char displayFiatConversion)
{
	static_cast<QmlBridge40a792*>(ptr)->displayAddress(QString::fromUtf8(address.data, address.len), QString::fromUtf8(wallet.data, wallet.len), displayFiatConversion != 0);
}

void QmlBridge40a792_ConnectAddTransactionToList(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, QString, QString, QString, QString)>(&QmlBridge40a792::addTransactionToList), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, QString, QString, QString, QString)>(&QmlBridge40a792::Signal_AddTransactionToList));
}

void QmlBridge40a792_DisconnectAddTransactionToList(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, QString, QString, QString, QString)>(&QmlBridge40a792::addTransactionToList), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, QString, QString, QString, QString)>(&QmlBridge40a792::Signal_AddTransactionToList));
}

void QmlBridge40a792_AddTransactionToList(void* ptr, struct Moc_PackedString paymentID, struct Moc_PackedString transactionID, struct Moc_PackedString amount, struct Moc_PackedString confirmations, struct Moc_PackedString time, struct Moc_PackedString number)
{
	static_cast<QmlBridge40a792*>(ptr)->addTransactionToList(QString::fromUtf8(paymentID.data, paymentID.len), QString::fromUtf8(transactionID.data, transactionID.len), QString::fromUtf8(amount.data, amount.len), QString::fromUtf8(confirmations.data, confirmations.len), QString::fromUtf8(time.data, time.len), QString::fromUtf8(number.data, number.len));
}

void QmlBridge40a792_ConnectAddRemoteNodeToList(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::addRemoteNodeToList), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_AddRemoteNodeToList));
}

void QmlBridge40a792_DisconnectAddRemoteNodeToList(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::addRemoteNodeToList), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_AddRemoteNodeToList));
}

void QmlBridge40a792_AddRemoteNodeToList(void* ptr, struct Moc_PackedString nodeURL)
{
	static_cast<QmlBridge40a792*>(ptr)->addRemoteNodeToList(QString::fromUtf8(nodeURL.data, nodeURL.len));
}

void QmlBridge40a792_ConnectSetSelectedRemoteNode(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(qint32)>(&QmlBridge40a792::setSelectedRemoteNode), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(qint32)>(&QmlBridge40a792::Signal_SetSelectedRemoteNode));
}

void QmlBridge40a792_DisconnectSetSelectedRemoteNode(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(qint32)>(&QmlBridge40a792::setSelectedRemoteNode), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(qint32)>(&QmlBridge40a792::Signal_SetSelectedRemoteNode));
}

void QmlBridge40a792_SetSelectedRemoteNode(void* ptr, int index)
{
	static_cast<QmlBridge40a792*>(ptr)->setSelectedRemoteNode(index);
}

void QmlBridge40a792_ConnectDisplayPopup(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, qint32)>(&QmlBridge40a792::displayPopup), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, qint32)>(&QmlBridge40a792::Signal_DisplayPopup));
}

void QmlBridge40a792_DisconnectDisplayPopup(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, qint32)>(&QmlBridge40a792::displayPopup), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, qint32)>(&QmlBridge40a792::Signal_DisplayPopup));
}

void QmlBridge40a792_DisplayPopup(void* ptr, struct Moc_PackedString text, int time)
{
	static_cast<QmlBridge40a792*>(ptr)->displayPopup(QString::fromUtf8(text.data, text.len), time);
}

void QmlBridge40a792_ConnectDisplaySyncingInfo(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::displaySyncingInfo), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::Signal_DisplaySyncingInfo));
}

void QmlBridge40a792_DisconnectDisplaySyncingInfo(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::displaySyncingInfo), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::Signal_DisplaySyncingInfo));
}

void QmlBridge40a792_DisplaySyncingInfo(void* ptr, struct Moc_PackedString syncing, struct Moc_PackedString syncingInfo)
{
	static_cast<QmlBridge40a792*>(ptr)->displaySyncingInfo(QString::fromUtf8(syncing.data, syncing.len), QString::fromUtf8(syncingInfo.data, syncingInfo.len));
}

void QmlBridge40a792_ConnectDisplayErrorDialog(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::displayErrorDialog), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::Signal_DisplayErrorDialog));
}

void QmlBridge40a792_DisconnectDisplayErrorDialog(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::displayErrorDialog), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::Signal_DisplayErrorDialog));
}

void QmlBridge40a792_DisplayErrorDialog(void* ptr, struct Moc_PackedString errorText, struct Moc_PackedString errorInformativeText)
{
	static_cast<QmlBridge40a792*>(ptr)->displayErrorDialog(QString::fromUtf8(errorText.data, errorText.len), QString::fromUtf8(errorInformativeText.data, errorInformativeText.len));
}

void QmlBridge40a792_ConnectClearTransferAmount(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::clearTransferAmount), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_ClearTransferAmount));
}

void QmlBridge40a792_DisconnectClearTransferAmount(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::clearTransferAmount), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_ClearTransferAmount));
}

void QmlBridge40a792_ClearTransferAmount(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->clearTransferAmount();
}

void QmlBridge40a792_ConnectAskForFusion(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::askForFusion), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_AskForFusion));
}

void QmlBridge40a792_DisconnectAskForFusion(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::askForFusion), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_AskForFusion));
}

void QmlBridge40a792_AskForFusion(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->askForFusion();
}

void QmlBridge40a792_ConnectClearListTransactions(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::clearListTransactions), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_ClearListTransactions));
}

void QmlBridge40a792_DisconnectClearListTransactions(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::clearListTransactions), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_ClearListTransactions));
}

void QmlBridge40a792_ClearListTransactions(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->clearListTransactions();
}

void QmlBridge40a792_ConnectDisplayPrivateKeys(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, QString, QString)>(&QmlBridge40a792::displayPrivateKeys), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, QString, QString)>(&QmlBridge40a792::Signal_DisplayPrivateKeys));
}

void QmlBridge40a792_DisconnectDisplayPrivateKeys(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, QString, QString)>(&QmlBridge40a792::displayPrivateKeys), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, QString, QString)>(&QmlBridge40a792::Signal_DisplayPrivateKeys));
}

void QmlBridge40a792_DisplayPrivateKeys(void* ptr, struct Moc_PackedString filename, struct Moc_PackedString privateViewKey, struct Moc_PackedString privateSpendKey, struct Moc_PackedString walletAddress)
{
	static_cast<QmlBridge40a792*>(ptr)->displayPrivateKeys(QString::fromUtf8(filename.data, filename.len), QString::fromUtf8(privateViewKey.data, privateViewKey.len), QString::fromUtf8(privateSpendKey.data, privateSpendKey.len), QString::fromUtf8(walletAddress.data, walletAddress.len));
}

void QmlBridge40a792_ConnectDisplaySeed(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, QString)>(&QmlBridge40a792::displaySeed), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, QString)>(&QmlBridge40a792::Signal_DisplaySeed));
}

void QmlBridge40a792_DisconnectDisplaySeed(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, QString)>(&QmlBridge40a792::displaySeed), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString, QString)>(&QmlBridge40a792::Signal_DisplaySeed));
}

void QmlBridge40a792_DisplaySeed(void* ptr, struct Moc_PackedString filename, struct Moc_PackedString mnemonicSeed, struct Moc_PackedString walletAddress)
{
	static_cast<QmlBridge40a792*>(ptr)->displaySeed(QString::fromUtf8(filename.data, filename.len), QString::fromUtf8(mnemonicSeed.data, mnemonicSeed.len), QString::fromUtf8(walletAddress.data, walletAddress.len));
}

void QmlBridge40a792_ConnectDisplayOpenWalletScreen(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::displayOpenWalletScreen), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_DisplayOpenWalletScreen));
}

void QmlBridge40a792_DisconnectDisplayOpenWalletScreen(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::displayOpenWalletScreen), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_DisplayOpenWalletScreen));
}

void QmlBridge40a792_DisplayOpenWalletScreen(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->displayOpenWalletScreen();
}

void QmlBridge40a792_ConnectDisplayMainWalletScreen(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::displayMainWalletScreen), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_DisplayMainWalletScreen));
}

void QmlBridge40a792_DisconnectDisplayMainWalletScreen(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::displayMainWalletScreen), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_DisplayMainWalletScreen));
}

void QmlBridge40a792_DisplayMainWalletScreen(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->displayMainWalletScreen();
}

void QmlBridge40a792_ConnectFinishedLoadingWalletd(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::finishedLoadingWalletd), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_FinishedLoadingWalletd));
}

void QmlBridge40a792_DisconnectFinishedLoadingWalletd(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::finishedLoadingWalletd), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_FinishedLoadingWalletd));
}

void QmlBridge40a792_FinishedLoadingWalletd(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->finishedLoadingWalletd();
}

void QmlBridge40a792_ConnectFinishedCreatingWallet(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::finishedCreatingWallet), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_FinishedCreatingWallet));
}

void QmlBridge40a792_DisconnectFinishedCreatingWallet(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::finishedCreatingWallet), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_FinishedCreatingWallet));
}

void QmlBridge40a792_FinishedCreatingWallet(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->finishedCreatingWallet();
}

void QmlBridge40a792_ConnectFinishedSendingTransaction(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::finishedSendingTransaction), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_FinishedSendingTransaction));
}

void QmlBridge40a792_DisconnectFinishedSendingTransaction(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::finishedSendingTransaction), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_FinishedSendingTransaction));
}

void QmlBridge40a792_FinishedSendingTransaction(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->finishedSendingTransaction();
}

void QmlBridge40a792_ConnectDisplayPathToPreviousWallet(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayPathToPreviousWallet), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayPathToPreviousWallet));
}

void QmlBridge40a792_DisconnectDisplayPathToPreviousWallet(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayPathToPreviousWallet), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayPathToPreviousWallet));
}

void QmlBridge40a792_DisplayPathToPreviousWallet(void* ptr, struct Moc_PackedString pathToPreviousWallet)
{
	static_cast<QmlBridge40a792*>(ptr)->displayPathToPreviousWallet(QString::fromUtf8(pathToPreviousWallet.data, pathToPreviousWallet.len));
}

void QmlBridge40a792_ConnectDisplayWalletCreationLocation(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayWalletCreationLocation), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayWalletCreationLocation));
}

void QmlBridge40a792_DisconnectDisplayWalletCreationLocation(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayWalletCreationLocation), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayWalletCreationLocation));
}

void QmlBridge40a792_DisplayWalletCreationLocation(void* ptr, struct Moc_PackedString walletLocation)
{
	static_cast<QmlBridge40a792*>(ptr)->displayWalletCreationLocation(QString::fromUtf8(walletLocation.data, walletLocation.len));
}

void QmlBridge40a792_ConnectDisplayUseRemoteNode(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(bool)>(&QmlBridge40a792::displayUseRemoteNode), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(bool)>(&QmlBridge40a792::Signal_DisplayUseRemoteNode));
}

void QmlBridge40a792_DisconnectDisplayUseRemoteNode(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(bool)>(&QmlBridge40a792::displayUseRemoteNode), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(bool)>(&QmlBridge40a792::Signal_DisplayUseRemoteNode));
}

void QmlBridge40a792_DisplayUseRemoteNode(void* ptr, char useRemote)
{
	static_cast<QmlBridge40a792*>(ptr)->displayUseRemoteNode(useRemote != 0);
}

void QmlBridge40a792_ConnectHideSettingsScreen(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::hideSettingsScreen), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_HideSettingsScreen));
}

void QmlBridge40a792_DisconnectHideSettingsScreen(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::hideSettingsScreen), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_HideSettingsScreen));
}

void QmlBridge40a792_HideSettingsScreen(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->hideSettingsScreen();
}

void QmlBridge40a792_ConnectDisplaySettingsScreen(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::displaySettingsScreen), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_DisplaySettingsScreen));
}

void QmlBridge40a792_DisconnectDisplaySettingsScreen(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::displaySettingsScreen), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_DisplaySettingsScreen));
}

void QmlBridge40a792_DisplaySettingsScreen(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->displaySettingsScreen();
}

void QmlBridge40a792_ConnectDisplaySettingsValues(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(bool)>(&QmlBridge40a792::displaySettingsValues), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(bool)>(&QmlBridge40a792::Signal_DisplaySettingsValues));
}

void QmlBridge40a792_DisconnectDisplaySettingsValues(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(bool)>(&QmlBridge40a792::displaySettingsValues), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(bool)>(&QmlBridge40a792::Signal_DisplaySettingsValues));
}

void QmlBridge40a792_DisplaySettingsValues(void* ptr, char displayFiat)
{
	static_cast<QmlBridge40a792*>(ptr)->displaySettingsValues(displayFiat != 0);
}

void QmlBridge40a792_ConnectDisplaySettingsRemoteDaemonInfo(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::displaySettingsRemoteDaemonInfo), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::Signal_DisplaySettingsRemoteDaemonInfo));
}

void QmlBridge40a792_DisconnectDisplaySettingsRemoteDaemonInfo(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::displaySettingsRemoteDaemonInfo), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString, QString)>(&QmlBridge40a792::Signal_DisplaySettingsRemoteDaemonInfo));
}

void QmlBridge40a792_DisplaySettingsRemoteDaemonInfo(void* ptr, struct Moc_PackedString remoteNodeAddress, struct Moc_PackedString remoteNodePort)
{
	static_cast<QmlBridge40a792*>(ptr)->displaySettingsRemoteDaemonInfo(QString::fromUtf8(remoteNodeAddress.data, remoteNodeAddress.len), QString::fromUtf8(remoteNodePort.data, remoteNodePort.len));
}

void QmlBridge40a792_ConnectDisplayFullBalanceInTransferAmount(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayFullBalanceInTransferAmount), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayFullBalanceInTransferAmount));
}

void QmlBridge40a792_DisconnectDisplayFullBalanceInTransferAmount(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayFullBalanceInTransferAmount), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayFullBalanceInTransferAmount));
}

void QmlBridge40a792_DisplayFullBalanceInTransferAmount(void* ptr, struct Moc_PackedString fullBalance)
{
	static_cast<QmlBridge40a792*>(ptr)->displayFullBalanceInTransferAmount(QString::fromUtf8(fullBalance.data, fullBalance.len));
}

void QmlBridge40a792_ConnectDisplayDefaultFee(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayDefaultFee), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayDefaultFee));
}

void QmlBridge40a792_DisconnectDisplayDefaultFee(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayDefaultFee), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayDefaultFee));
}

void QmlBridge40a792_DisplayDefaultFee(void* ptr, struct Moc_PackedString fee)
{
	static_cast<QmlBridge40a792*>(ptr)->displayDefaultFee(QString::fromUtf8(fee.data, fee.len));
}

void QmlBridge40a792_ConnectDisplayNodeFee(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayNodeFee), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayNodeFee));
}

void QmlBridge40a792_DisconnectDisplayNodeFee(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::displayNodeFee), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(QString)>(&QmlBridge40a792::Signal_DisplayNodeFee));
}

void QmlBridge40a792_DisplayNodeFee(void* ptr, struct Moc_PackedString nodeFee)
{
	static_cast<QmlBridge40a792*>(ptr)->displayNodeFee(QString::fromUtf8(nodeFee.data, nodeFee.len));
}

void QmlBridge40a792_ConnectUpdateConfirmationsOfTransaction(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(qint32, QString)>(&QmlBridge40a792::updateConfirmationsOfTransaction), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(qint32, QString)>(&QmlBridge40a792::Signal_UpdateConfirmationsOfTransaction));
}

void QmlBridge40a792_DisconnectUpdateConfirmationsOfTransaction(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(qint32, QString)>(&QmlBridge40a792::updateConfirmationsOfTransaction), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(qint32, QString)>(&QmlBridge40a792::Signal_UpdateConfirmationsOfTransaction));
}

void QmlBridge40a792_UpdateConfirmationsOfTransaction(void* ptr, int index, struct Moc_PackedString confirmations)
{
	static_cast<QmlBridge40a792*>(ptr)->updateConfirmationsOfTransaction(index, QString::fromUtf8(confirmations.data, confirmations.len));
}

void QmlBridge40a792_ConnectDisplayInfoDialog(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::displayInfoDialog), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_DisplayInfoDialog));
}

void QmlBridge40a792_DisconnectDisplayInfoDialog(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::displayInfoDialog), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)()>(&QmlBridge40a792::Signal_DisplayInfoDialog));
}

void QmlBridge40a792_DisplayInfoDialog(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->displayInfoDialog();
}

void QmlBridge40a792_ConnectAddSavedAddressToList(void* ptr)
{
	QObject::connect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(qint32, QString, QString, QString)>(&QmlBridge40a792::addSavedAddressToList), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(qint32, QString, QString, QString)>(&QmlBridge40a792::Signal_AddSavedAddressToList));
}

void QmlBridge40a792_DisconnectAddSavedAddressToList(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(qint32, QString, QString, QString)>(&QmlBridge40a792::addSavedAddressToList), static_cast<QmlBridge40a792*>(ptr), static_cast<void (QmlBridge40a792::*)(qint32, QString, QString, QString)>(&QmlBridge40a792::Signal_AddSavedAddressToList));
}

void QmlBridge40a792_AddSavedAddressToList(void* ptr, int dbID, struct Moc_PackedString name, struct Moc_PackedString address, struct Moc_PackedString paymentID)
{
	static_cast<QmlBridge40a792*>(ptr)->addSavedAddressToList(dbID, QString::fromUtf8(name.data, name.len), QString::fromUtf8(address.data, address.len), QString::fromUtf8(paymentID.data, paymentID.len));
}

void QmlBridge40a792_Log(void* ptr, struct Moc_PackedString msg)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "log", Q_ARG(QString, QString::fromUtf8(msg.data, msg.len)));
}

void QmlBridge40a792_ClickedButtonExplorer(void* ptr, struct Moc_PackedString transactionID)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedButtonExplorer", Q_ARG(QString, QString::fromUtf8(transactionID.data, transactionID.len)));
}

void QmlBridge40a792_GoToWebsite(void* ptr, struct Moc_PackedString url)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "goToWebsite", Q_ARG(QString, QString::fromUtf8(url.data, url.len)));
}

void QmlBridge40a792_ClickedButtonCopyTx(void* ptr, struct Moc_PackedString transactionID)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedButtonCopyTx", Q_ARG(QString, QString::fromUtf8(transactionID.data, transactionID.len)));
}

void QmlBridge40a792_ClickedButtonCopyAddress(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedButtonCopyAddress");
}

void QmlBridge40a792_ClickedButtonCopyKeys(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedButtonCopyKeys");
}

void QmlBridge40a792_ClickedButtonCopy(void* ptr, struct Moc_PackedString stringToCopy)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedButtonCopy", Q_ARG(QString, QString::fromUtf8(stringToCopy.data, stringToCopy.len)));
}

void QmlBridge40a792_ClickedButtonSend(void* ptr, struct Moc_PackedString transferAddress, struct Moc_PackedString transferAmount, struct Moc_PackedString transferPaymentID, struct Moc_PackedString transferFee)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedButtonSend", Q_ARG(QString, QString::fromUtf8(transferAddress.data, transferAddress.len)), Q_ARG(QString, QString::fromUtf8(transferAmount.data, transferAmount.len)), Q_ARG(QString, QString::fromUtf8(transferPaymentID.data, transferPaymentID.len)), Q_ARG(QString, QString::fromUtf8(transferFee.data, transferFee.len)));
}

void QmlBridge40a792_ClickedButtonBackupWallet(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedButtonBackupWallet");
}

void QmlBridge40a792_ClickedCloseWallet(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedCloseWallet");
}

void QmlBridge40a792_ClickedButtonOpen(void* ptr, struct Moc_PackedString pathToWallet, struct Moc_PackedString passwordWallet)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedButtonOpen", Q_ARG(QString, QString::fromUtf8(pathToWallet.data, pathToWallet.len)), Q_ARG(QString, QString::fromUtf8(passwordWallet.data, passwordWallet.len)));
}

void QmlBridge40a792_ClickedButtonCreate(void* ptr, struct Moc_PackedString filenameWallet, struct Moc_PackedString passwordWallet, struct Moc_PackedString confirmPasswordWallet)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedButtonCreate", Q_ARG(QString, QString::fromUtf8(filenameWallet.data, filenameWallet.len)), Q_ARG(QString, QString::fromUtf8(passwordWallet.data, passwordWallet.len)), Q_ARG(QString, QString::fromUtf8(confirmPasswordWallet.data, confirmPasswordWallet.len)));
}

void QmlBridge40a792_ClickedButtonImport(void* ptr, struct Moc_PackedString filenameWallet, struct Moc_PackedString passwordWallet, struct Moc_PackedString privateViewKey, struct Moc_PackedString privateSpendKey, struct Moc_PackedString mnemonicSeed, struct Moc_PackedString confirmPasswordWallet, struct Moc_PackedString scanHeight)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedButtonImport", Q_ARG(QString, QString::fromUtf8(filenameWallet.data, filenameWallet.len)), Q_ARG(QString, QString::fromUtf8(passwordWallet.data, passwordWallet.len)), Q_ARG(QString, QString::fromUtf8(privateViewKey.data, privateViewKey.len)), Q_ARG(QString, QString::fromUtf8(privateSpendKey.data, privateSpendKey.len)), Q_ARG(QString, QString::fromUtf8(mnemonicSeed.data, mnemonicSeed.len)), Q_ARG(QString, QString::fromUtf8(confirmPasswordWallet.data, confirmPasswordWallet.len)), Q_ARG(QString, QString::fromUtf8(scanHeight.data, scanHeight.len)));
}

void QmlBridge40a792_ChoseRemote(void* ptr, char remote)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "choseRemote", Q_ARG(bool, remote != 0));
}

void QmlBridge40a792_SelectedRemoteNode(void* ptr, int index)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "selectedRemoteNode", Q_ARG(qint32, index));
}

struct Moc_PackedString QmlBridge40a792_GetTransferAmountUSD(void* ptr, struct Moc_PackedString amountTRTL)
{
	QString returnArg;
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "getTransferAmountUSD", Q_RETURN_ARG(QString, returnArg), Q_ARG(QString, QString::fromUtf8(amountTRTL.data, amountTRTL.len)));
	return ({ QByteArray t8e5b69 = returnArg.toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

void QmlBridge40a792_ClickedCloseSettings(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedCloseSettings");
}

void QmlBridge40a792_ClickedSettingsButton(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "clickedSettingsButton");
}

void QmlBridge40a792_ChoseDisplayFiat(void* ptr, char displayFiat)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "choseDisplayFiat", Q_ARG(bool, displayFiat != 0));
}

void QmlBridge40a792_ChoseCheckpoints(void* ptr, char checkpoints)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "choseCheckpoints", Q_ARG(bool, checkpoints != 0));
}

void QmlBridge40a792_SaveRemoteDaemonInfo(void* ptr, struct Moc_PackedString daemonAddress, struct Moc_PackedString daemonPort)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "saveRemoteDaemonInfo", Q_ARG(QString, QString::fromUtf8(daemonAddress.data, daemonAddress.len)), Q_ARG(QString, QString::fromUtf8(daemonPort.data, daemonPort.len)));
}

void QmlBridge40a792_ResetRemoteDaemonInfo(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "resetRemoteDaemonInfo");
}

void QmlBridge40a792_GetFullBalanceAndDisplayInTransferAmount(void* ptr, struct Moc_PackedString transferFee)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "getFullBalanceAndDisplayInTransferAmount", Q_ARG(QString, QString::fromUtf8(transferFee.data, transferFee.len)));
}

void QmlBridge40a792_GetDefaultFeeAndDisplay(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "getDefaultFeeAndDisplay");
}

void QmlBridge40a792_LimitDisplayTransactions(void* ptr, char limit)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "limitDisplayTransactions", Q_ARG(bool, limit != 0));
}

struct Moc_PackedString QmlBridge40a792_GetVersion(void* ptr)
{
	QString returnArg;
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "getVersion", Q_RETURN_ARG(QString, returnArg));
	return ({ QByteArray t8e5b69 = returnArg.toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

struct Moc_PackedString QmlBridge40a792_GetNewVersion(void* ptr)
{
	QString returnArg;
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "getNewVersion", Q_RETURN_ARG(QString, returnArg));
	return ({ QByteArray t8e5b69 = returnArg.toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

struct Moc_PackedString QmlBridge40a792_GetNewVersionURL(void* ptr)
{
	QString returnArg;
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "getNewVersionURL", Q_RETURN_ARG(QString, returnArg));
	return ({ QByteArray t8e5b69 = returnArg.toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

void QmlBridge40a792_OptimizeWalletWithFusion(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "optimizeWalletWithFusion");
}

void QmlBridge40a792_SaveAddress(void* ptr, struct Moc_PackedString name, struct Moc_PackedString address, struct Moc_PackedString paymentID)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "saveAddress", Q_ARG(QString, QString::fromUtf8(name.data, name.len)), Q_ARG(QString, QString::fromUtf8(address.data, address.len)), Q_ARG(QString, QString::fromUtf8(paymentID.data, paymentID.len)));
}

void QmlBridge40a792_FillListSavedAddresses(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "fillListSavedAddresses");
}

void QmlBridge40a792_DeleteSavedAddress(void* ptr, int dbID)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "deleteSavedAddress", Q_ARG(qint32, dbID));
}

void QmlBridge40a792_RegisterToGo(void* ptr, void* object)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "registerToGo", Q_ARG(QObject*, static_cast<QObject*>(object)));
}

void QmlBridge40a792_DeregisterToGo(void* ptr, struct Moc_PackedString objectName)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40a792*>(ptr), "deregisterToGo", Q_ARG(QString, QString::fromUtf8(objectName.data, objectName.len)));
}

int QmlBridge40a792_QmlBridge40a792_QRegisterMetaType()
{
	return qRegisterMetaType<QmlBridge40a792*>();
}

int QmlBridge40a792_QmlBridge40a792_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QmlBridge40a792*>(const_cast<const char*>(typeName));
}

int QmlBridge40a792_QmlBridge40a792_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge40a792>();
#else
	return 0;
#endif
}

int QmlBridge40a792_QmlBridge40a792_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge40a792>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QmlBridge40a792___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QmlBridge40a792___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QmlBridge40a792___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QmlBridge40a792___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge40a792___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge40a792___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridge40a792___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge40a792___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge40a792___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridge40a792___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge40a792___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge40a792___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridge40a792___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge40a792___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge40a792___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QmlBridge40a792_NewQmlBridge(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40a792(static_cast<QWindow*>(parent));
	} else {
		return new QmlBridge40a792(static_cast<QObject*>(parent));
	}
}

void QmlBridge40a792_DestroyQmlBridge(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->~QmlBridge40a792();
}

void QmlBridge40a792_DestroyQmlBridgeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QmlBridge40a792_EventDefault(void* ptr, void* e)
{
	return static_cast<QmlBridge40a792*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QmlBridge40a792_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QmlBridge40a792*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QmlBridge40a792_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge40a792*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlBridge40a792_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge40a792*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge40a792_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge40a792*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlBridge40a792_DeleteLaterDefault(void* ptr)
{
	static_cast<QmlBridge40a792*>(ptr)->QObject::deleteLater();
}

void QmlBridge40a792_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge40a792*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge40a792_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge40a792*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
