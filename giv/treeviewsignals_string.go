// Code generated by "stringer -type=TreeViewSignals"; DO NOT EDIT.

package giv

import (
	"fmt"
	"strconv"
)

const _TreeViewSignals_name = "TreeViewSelectedTreeViewUnselectedTreeViewAllSelectedTreeViewAllUnselectedTreeViewOpenedTreeViewClosedTreeViewSignalsN"

var _TreeViewSignals_index = [...]uint8{0, 16, 34, 53, 74, 88, 102, 118}

func (i TreeViewSignals) String() string {
	if i < 0 || i >= TreeViewSignals(len(_TreeViewSignals_index)-1) {
		return "TreeViewSignals(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TreeViewSignals_name[_TreeViewSignals_index[i]:_TreeViewSignals_index[i+1]]
}

func (i *TreeViewSignals) FromString(s string) error {
	for j := 0; j < len(_TreeViewSignals_index)-1; j++ {
		if s == _TreeViewSignals_name[_TreeViewSignals_index[j]:_TreeViewSignals_index[j+1]] {
			*i = TreeViewSignals(j)
			return nil
		}
	}
	return fmt.Errorf("String %v is not a valid option for type TreeViewSignals", s)
}