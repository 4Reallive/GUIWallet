// Copyright (c) 2018, The TurtleCoin Developers
// Copyright (c) 2018, The Celestial Developers
//
// Please see the included LICENSE file for more information.
//

package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in CELC
	DefaultTransferFee float64 = 0.1

	logWalletdCurrentSessionFilename     = "service-session.log"
	logWalletdAllSessionsFilename        = "service.log"
	logTurtleCoindCurrentSessionFilename = "Celestiald-session.log"
	logTurtleCoindAllSessionsFilename    = "Celestiald.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "service"
	turtlecoindCommandName               = "Celestiald"
)
