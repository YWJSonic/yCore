package rss

import "strings"

const (
	Space_Code string = "&nbsp;"
	Etc_Code   string = "&hellip;"
	cr_Code    string = "&crarr;"
	ua_Code    string = "&uarr;"
	ra_Code    string = "&rarr;"
	da_Code    string = "&darr;"
	ha_Code    string = "&harr;"
	lA_Code    string = "&lArr;"
	uA_Code    string = "&uArr;"
	rA_Code    string = "&rArr;"
	dA_Code    string = "&dArr;"
	hA_Code    string = "&hArr;"
)
const (
	Spcae_Str string = " "
	Etc_Str   string = "…"
	cr_Str    string = "↵"
	lu_Str    string = "↑"
	ra_Str    string = "→"
	da_Str    string = "↓"
	ha_Str    string = "↔"
	lA_Str    string = "⇐"
	uA_Str    string = "⇑"
	rA_Str    string = "⇒"
	dA_Str    string = "⇓"
	hA_Str    string = "⇔"
)

const (
	NodeType_Element NodeType = iota
	NodeType_ProcInst
	NodeType_Other
)

const (
	Space_Num   uint8 = 9
	EndLine_Num uint8 = 10
)

var replacer *strings.Replacer = strings.NewReplacer(
	Space_Code, Spcae_Str,
	Etc_Code, Etc_Str,
	cr_Code, cr_Str,
	ua_Code, lu_Str,
	ra_Code, ra_Str,
	da_Code, da_Str,
	ha_Code, ha_Str,
	lA_Code, lA_Str,
	uA_Code, uA_Str,
	rA_Code, rA_Str,
	dA_Code, dA_Str,
	hA_Code, hA_Str)
