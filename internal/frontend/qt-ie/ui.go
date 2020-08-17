// Copyright (c) 2020 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

// +build !nogui

package qtie

import (
	"runtime"

	"github.com/therecipe/qt/core"
)

// GoQMLInterface between go and qml
//
// Here we implements all the signals / methods.
type GoQMLInterface struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"currentAddress"`
	_ string `property:"goos"`
	_ string `property:"credits"`
	_ bool   `property:"isFirstStart"`
	_ bool   `property:"isRestarting"`
	_ bool   `property:"isConnectionOK"`

	_ string  `property:lastError`
	_ float32 `property:progress`
	_ string  `property:progressDescription`
	_ int     `property:progressFails`
	_ int     `property:total`
	_ string  `property:importLogFileName`

	_ string `property:"programTitle"`
	_ string `property:"newversion"`
	_ string `property:"downloadLink"`
	_ string `property:"landingPage"`
	_ string `property:"changelog"`
	_ string `property:"bugfixes"`

	// translations
	_ string `property:"wrongCredentials"`
	_ string `property:"wrongMailboxPassword"`
	_ string `property:"canNotReachAPI"`
	_ string `property:"credentialsNotRemoved"`
	_ string `property:"versionCheckFailed"`
	//
	_ func(isAvailable bool)   `signal:"setConnectionStatus"`
	_ func(updateState string) `signal:"setUpdateState"`
	_ func()                   `slot:"checkInternet"`

	_ func()                 `signal:"processFinished"`
	_ func(okay bool)        `signal:"exportStructureLoadFinished"`
	_ func(okay bool)        `signal:"importStructuresLoadFinished"`
	_ func()                 `signal:"openManual"`
	_ func(showMessage bool) `signal:"runCheckVersion"`
	_ func()                 `slot:"getLocalVersionInfo"`
	_ func()                 `slot:"loadImportReports"`

	_ func() `slot:"quit"`
	_ func() `slot:"loadAccounts"`
	_ func() `slot:"openLogs"`
	_ func() `slot:"openDownloadLink"`
	_ func() `slot:"openReport"`
	_ func() `slot:"clearCache"`
	_ func() `slot:"clearKeychain"`
	_ func() `signal:"highlightSystray"`
	_ func() `signal:"normalSystray"`

	_ func(showMessage bool) `slot:"isNewVersionAvailable"`
	_ func() string          `slot:"getBackendVersion"`

	_ func(description, client, address string) bool                                       `slot:"sendBug"`
	_ func(address string) bool                                                            `slot:"sendImportReport"`
	_ func(address string)                                                                 `slot:"loadStructureForExport"`
	_ func() string                                                                        `slot:"leastUsedColor"`
	_ func(username string, name string, color string, isLabel bool, sourceID string) bool `slot:"createLabelOrFolder"`
	_ func(fpath, address, fileType string, attachEncryptedBody bool)                      `slot:"startExport"`
	_ func(email string)                                                                   `slot:"startImport"`
	_ func()                                                                               `slot:"resetSource"`

	_ func(isFromIMAP bool, sourcePath, sourceEmail, sourcePassword, sourceServe, sourcePort, targetAddress string) `slot:"setupAndLoadForImport"`

	_ string `property:"progressInit"`

	_ func(path string) int `slot:"checkPathStatus"`

	_ func(evType string, msg string)    `signal:"emitEvent"`
	_ func(tabIndex int, message string) `signal:"notifyBubble"`

	_ func() `signal:"bubbleClosed"`
	_ func() `signal:"simpleErrorHappen"`
	_ func() `signal:"askErrorHappen"`
	_ func() `signal:"retryErrorHappen"`
	_ func() `signal:"pauseProcess"`
	_ func() `signal:"resumeProcess"`
	_ func() `signal:"cancelProcess"`

	_ func(iAccount int, prefRem bool)      `slot:"deleteAccount"`
	_ func(iAccount int)                    `slot:"logoutAccount"`
	_ func(login, password string) int      `slot:"login"`
	_ func(twoFacAuth string) int           `slot:"auth2FA"`
	_ func(mailboxPassword string) int      `slot:"addAccount"`
	_ func(message string, changeIndex int) `signal:"setAddAccountWarning"`

	_ func()               `signal:"notifyVersionIsTheLatest"`
	_ func()               `signal:"notifyKeychainRebuild"`
	_ func()               `signal:"notifyHasNoKeychain"`
	_ func()               `signal:"notifyUpdate"`
	_ func(accname string) `signal:"notifyLogout"`
	_ func(accname string) `signal:"notifyAddressChanged"`
	_ func(accname string) `signal:"notifyAddressChangedLogout"`

	_ func()              `slot:"startUpdate"`
	_ func(hasError bool) `signal:"updateFinished"`

	// errors
	_ func()            `signal:"answerRetry"`
	_ func(all bool)    `signal:"answerSkip"`
	_ func(errCode int) `signal:"notifyError"`
	_ string            `property:"errorDescription"`
}

// Constructor
func (s *GoQMLInterface) init() {}

// SetFrontend connects all slots and signals from Go to QML
func (s *GoQMLInterface) SetFrontend(f *FrontendQt) {
	s.ConnectQuit(f.App.Quit)

	s.ConnectLoadAccounts(f.Accounts.LoadAccounts)
	s.ConnectOpenLogs(f.openLogs)
	s.ConnectOpenDownloadLink(f.openDownloadLink)
	s.ConnectOpenReport(f.openReport)
	s.ConnectClearCache(f.Accounts.ClearCache)
	s.ConnectClearKeychain(f.Accounts.ClearKeychain)

	s.ConnectSendBug(f.sendBug)
	s.ConnectSendImportReport(f.sendImportReport)

	s.ConnectDeleteAccount(f.Accounts.DeleteAccount)
	s.ConnectLogoutAccount(f.Accounts.LogoutAccount)
	s.ConnectLogin(f.Accounts.Login)
	s.ConnectAuth2FA(f.Accounts.Auth2FA)
	s.ConnectAddAccount(f.Accounts.AddAccount)

	s.SetGoos(runtime.GOOS)
	s.SetIsRestarting(false)
	s.SetProgramTitle(f.programName)

	s.ConnectGetLocalVersionInfo(f.getLocalVersionInfo)
	s.ConnectIsNewVersionAvailable(f.isNewVersionAvailable)
	s.ConnectGetBackendVersion(func() string {
		return f.programVersion
	})

	s.ConnectCheckInternet(f.checkInternet)

	s.ConnectLoadStructureForExport(f.LoadStructureForExport)
	s.ConnectSetupAndLoadForImport(f.setupAndLoadForImport)
	s.ConnectResetSource(f.resetSource)
	s.ConnectLeastUsedColor(f.leastUsedColor)
	s.ConnectCreateLabelOrFolder(f.createLabelOrFolder)

	s.ConnectStartExport(f.StartExport)
	s.ConnectStartImport(f.StartImport)

	s.ConnectCheckPathStatus(CheckPathStatus)

	s.ConnectStartUpdate(f.StartUpdate)

	s.ConnectEmitEvent(f.emitEvent)
}
