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

// Code generated by ./release-notes.sh at 'Mon Sep 21 01:29:10 PM CEST 2020'. DO NOT EDIT.

package bridge

const ReleaseNotes = `• Bulletproofing against any potential data loss and/or duplication
• Performance improvements for handling attachments and non-standard formatting
• Better stability of the message parser
• Additional foreign encoding support for outgoing messages
• Complete refactor of the way messages are parsed to simplify code maintenance
• Improved User-Agent detection
• Added MacOS Big Sur compatibility
• Added persistent anonymous API cookies
`

const ReleaseFixedBugs = `• Fixed rare mail loss when moving from Spam folder
• Limited log size
• Fixed Linux font issues (mouse hover).
`
