// Copyright 2025 the libevm authors.
//
// The libevm additions to go-ethereum are free software: you can redistribute
// them and/or modify them under the terms of the GNU Lesser General Public License
// as published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.
//
// The libevm additions are distributed in the hope that they will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser
// General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see
// <http://www.gnu.org/licenses/>.

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

func theirs(cmd *cobra.Command, _ []string) error {
	stdin := bufio.NewScanner(cmd.InOrStdin())

	var files []*fileConflicts
	for stdin.Scan() {
		cs, err := findConflicts(stdin.Text())
		if err != nil {
			return err
		}
		files = append(files, cs)
	}
	if err := stdin.Err(); err != nil {
		return err
	}
	if len(files) == 0 {
		return nil
	}

	ui := createTheirsUI(files)
	if err := ui.app.Run(); err != nil {
		return err
	}
	_, err := cmd.OutOrStdout().Write([]byte(strings.Join(ui.filesToAcceptAllTheirs, "\n")))
	return err
}

type conflict struct {
	Current, Ancestor, Incoming []string
}

type conflictState string

const (
	currentChange  = conflictState("<<<<<<<")
	ancestor       = conflictState("|||||||")
	incomingChange = conflictState("=======")
	endConflict    = conflictState(">>>>>>>")
	noConflict     = endConflict
)

var (
	conflictDelimOrder = []conflictState{noConflict, currentChange, ancestor, incomingChange, endConflict}
	conflictDelimPrev  = make(map[conflictState]conflictState, len(conflictDelimOrder))
)

func init() {
	for i, d := range conflictDelimOrder {
		if i > 0 {
			conflictDelimPrev[d] = conflictDelimOrder[i-1]
		}
	}
}

func nextConflictState(line string) (conflictState, bool) {
	for _, d := range conflictDelimOrder {
		if strings.HasPrefix(line, string(d)) {
			return d, true
		}
	}
	return "", false
}

type fileConflicts struct {
	path      string
	conflicts []*conflict
}

func findConflicts(path string) (*fileConflicts, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := bufio.NewScanner(f)
	var (
		conflicts []*conflict
		curr      *conflict
		chunk     []string
	)
	for state := noConflict; lines.Scan(); {
		line := lines.Text()
		next, ok := nextConflictState(line)
		if !ok {
			if state != noConflict {
				chunk = append(chunk, line)
			}
			continue
		}
		if conflictDelimPrev[next] != state {
			return nil, fmt.Errorf("conflict delimeter %q after %q", next, state)
		}

		switch next {
		case currentChange:
			curr = new(conflict)
			conflicts = append(conflicts, curr)
		case ancestor:
			curr.Current = slices.Clone(chunk)
		case incomingChange:
			curr.Ancestor = slices.Clone(chunk)
		case endConflict:
			curr.Incoming = slices.Clone(chunk)
		}
		state = next
		chunk = chunk[:0]
	}

	if err := lines.Err(); err != nil {
		return nil, err
	}
	return &fileConflicts{path, conflicts}, nil
}

type theirsUI struct {
	files                []*fileConflicts
	fileIdx, conflictIdx int

	app                      *tview.Application
	curr, ancestor, incoming *tview.TextView

	filesToAcceptAllTheirs []string
}

func createTheirsUI(files []*fileConflicts) *theirsUI {

	ui := &theirsUI{
		files: files,
	}

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow) // top to bottom (makes no sense IMO)

	inst := tview.NewTextView().SetText(
		`Should the incoming change be accepted without modification? (y/n)

Typically this will only be the case if *everything* is a Go import. If in doubt, press "n"!

If all conflicts in a file receive a "yes" then that file's path is printed to stdout to be piped into ` + "`git checkout --theirs --`." + `

Press "q" to exit early. Files already receiving all "yes" will still be printed.`,
	)
	inst.SetBorder(true).SetTitle("Instructions").SetTitleAlign(tview.AlignLeft)
	flex.AddItem(inst, 0, 1, false)

	views := []**tview.TextView{&ui.curr, &ui.ancestor, &ui.incoming}
	titles := []string{"Current (ours)", "Ancestor", "Incoming (theirs)"}
	for i, view := range views {
		tv := tview.NewTextView()
		tv.SetBorder(true).SetTitle(titles[i]).SetTitleAlign(tview.AlignLeft)
		*view = tv

		flex.AddItem(tv, 0, 1, true)
	}

	ui.app = tview.NewApplication().SetRoot(flex, true)
	ui.createKeyBindings()
	ui.showConflictOrStopApp()
	return ui
}

func (u *theirsUI) showConflictOrStopApp() {
	if u.fileIdx >= len(u.files) {
		u.app.Stop()
		return
	}

	c := u.files[u.fileIdx].conflicts[u.conflictIdx]
	u.curr.SetText(strings.Join(c.Current, "\n"))
	u.ancestor.SetText(strings.Join(c.Ancestor, "\n"))
	u.incoming.SetText(strings.Join(c.Incoming, "\n"))
}

func (u *theirsUI) nextFile(allImports bool) {
	if allImports {
		u.filesToAcceptAllTheirs = append(u.filesToAcceptAllTheirs, u.files[u.fileIdx].path)
	}
	u.fileIdx++
	u.conflictIdx = 0
}

func (u *theirsUI) createKeyBindings() {
	u.app.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		switch ev.Rune() {
		case 'y':
			u.conflictIdx++
			if u.conflictIdx >= len(u.files[u.fileIdx].conflicts) {
				u.nextFile(true)
			}
		case 'n':
			u.nextFile(false)
		case 'q':
			u.app.Stop()
			return ev
		default:
			return ev
		}

		u.showConflictOrStopApp()
		return ev
	})
}
