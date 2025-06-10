package utils

var TransitionOpacityOnGroupHover =  "invisible group-hover:visible opacity-0 group-hover:opacity-100 transition-opacity"
var Clickable = "hover:cursor-pointer"

func HoverableIfNonempty(image string) string {
    if image != "" {
        return TransitionOpacityOnGroupHover
    }
    return ""
}

func ClickableIfNonempty(link string) string {
    if link != "" {
        return Clickable
    }
    return ""
}
