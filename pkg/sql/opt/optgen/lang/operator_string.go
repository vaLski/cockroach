// Code generated by "stringer -type=Operator operator.og.go"; DO NOT EDIT.

package lang

import "strconv"

const _Operator_name = "UnknownOpRootOpDefineSetOpRuleSetOpDefineOpCommentsOpCommentOpTagsOpTagOpDefineFieldsOpDefineFieldOpRuleOpBindOpRefOpMatchOpOpNamesOpOpNameOpMatchAndOpMatchInvokeOpMatchNotOpMatchAnyOpMatchListAnyOpMatchListFirstOpMatchListLastOpMatchListSingleOpMatchListEmptyOpConstructOpConstructListOpListOpStringOp"

var _Operator_index = [...]uint16{0, 9, 15, 26, 35, 43, 53, 62, 68, 73, 87, 100, 106, 112, 117, 124, 133, 141, 151, 164, 174, 184, 198, 214, 229, 246, 262, 273, 288, 294, 302}

func (i Operator) String() string {
	if i < 0 || i >= Operator(len(_Operator_index)-1) {
		return "Operator(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Operator_name[_Operator_index[i]:_Operator_index[i+1]]
}