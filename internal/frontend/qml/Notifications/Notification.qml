// Copyright (c) 2021 Proton Technologies AG
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

import QtQml 2.12
import QtQuick.Controls 2.12

QtObject {
    id: root

    default property var children

    enum NotificationType {
        Info = 0,
        Success = 1,
        Warning = 2,
        Danger = 3
    }

    property string text
    property string description
    property string icon
    property list<Action> action
    property int type
    property int group

    property bool dismissed: false
    property bool active: false
    property bool loading: false
    readonly property var occurred: active ? new Date() : undefined

    property var data

    onActiveChanged: {
        dismissed = false
    }
}