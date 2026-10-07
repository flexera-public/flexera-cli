package catalog

// Offline, reviewed vocabulary; no model or external index is involved.
var synonyms = map[string]string{
	"remove": "delete", "destroy": "delete", "show": "get", "describe": "get", "fetch": "get", "read": "get",
	"add": "create", "new": "create", "make": "create", "edit": "update", "modify": "update", "patch": "update", "change": "update",
	"spend": "cost", "spending": "cost", "bills": "bill", "execute": "evaluate", "run": "evaluate", "who": "user", "org": "organization",
}
