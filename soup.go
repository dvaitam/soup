package soup

import (
	"strings"
)

type Tag struct {
	i, j    int
	elem    string
	attrs   map[string]string
	closing bool
}

var single_elements = map[string]bool{"br": true, "link": true}

func GetTag(s string, i int, j int) Tag {
	if s[i+1] == '/' {
		return Tag{i: i, j: j, elem: s[i+2 : j], closing: true}
	} else {
		parts := strings.Fields(s[i+1 : j])
		tag := Tag{i: i, j: j}
		tag.elem = parts[0]
		attrs := map[string]string{}

		for i := 1; i < len(parts); i++ {
			att := strings.Split(parts[i], "=")
			if len(att) > 1 {
				if len(att[1]) == 0 {
					continue
				}
				val := att[1]
				ll := len(val)
				if val[0:1] == "\"" && val[ll-1:ll] == "\"" {
					val = val[1 : ll-1]
				}
				attrs[att[0]] = val
			}
		}
		tag.attrs = attrs
		return tag
	}
}
func Load(s string) []Tag {
	tags := make([]Tag, 0)
	started := false
	start_index := 0
	for i := 0; i < len(s); i++ {
		if started {
			if s[i] == '>' {
				tags = append(tags, GetTag(s, start_index, i))
				started = false
			}
		} else {
			if s[i] == '<' {
				started = true
				start_index = i
			}
		}
	}
	return tags
}
func GetTagsByClassOrId(s string, IdOrClass string, Id bool) [][]Tag {
	tags := Load(s)
	count := 0
	started := false
	start_index := 0
	all_tags := make([][]Tag, 0)
	for i := 0; i < len(tags); i++ {
		if started {
			if tags[i].closing {
				count--
			} else if !single_elements[tags[i].elem] {
				count++
			}
			if count == 0 {
				all_tags = append(all_tags, tags[start_index:i+1])
				started = false
			}
		} else {
			val, ok := tags[i].attrs[IdOrClass]
			if ok {
				if Id {
					if val == IdOrClass {
						started = true
						start_index = i
						count++
					}
				} else {
					classes := strings.Fields(val)
					found := false
					for _, class := range classes {
						if class == IdOrClass {
							found = true
							break
						}
					}
					if found {
						started = true
						start_index = i
						count++
					}
				}
			}
		}
	}
	return all_tags
}
func GetTagsById(s string, Id string) []Tag {
	all_tags := GetTagsByClassOrId(s, Id, true)
	if len(all_tags) > 0 {
		return all_tags[0]
	}
	return nil
}
func GetDivById(s string, Id string) string {
	tags := GetTagsById(s, Id)
	if len(tags) == 0 {
		return ""
	}
	n := len(tags)
	return s[tags[0].i : tags[n-1].j+1]
}
func GetHtmlById(s string, Id string) string {
	tags := GetTagsById(s, Id)
	if len(tags) == 0 {
		return ""
	}
	n := len(tags)
	return s[tags[0].i : tags[n-1].j+1]
}
func GetTextFromTags(s string, tags []Tag) string {
	if len(tags) <= 1 {
		return ""
	}
	var sb strings.Builder
	for i := 1; i <= len(tags); i++ {
		curr, prev := tags[i], tags[i-1]
		sb.WriteString(s[prev.j+1 : curr.i])
	}
	return sb.String()
}
func GetTextById(s string, Id string) string {
	tags := GetTagsById(s, Id)
	return GetTextFromTags(s, tags)
}
func GetTextsFromClass(s string, Class string) []string {
	all_tags := GetTagsByClassOrId(s, Class, false)
	texts := make([]string, 0)
	for i := 0; i < len(all_tags); i++ {
		texts = append(texts, GetTextFromTags(s, all_tags[i]))
	}
	return texts
}
