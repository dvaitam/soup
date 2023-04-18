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

func get_tag(s string, i int, j int) Tag {
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
				attrs[att[0]] = att[1]
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
				tags = append(tags, get_tag(s, start_index, i))
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
func GetDivById(s string, id string) string {
	tags := Load(s)
	count := 0
	started := false
	start_index := 0
	for i := 0; i < len(tags); i++ {
		if started {
			if tags[i].closing {
				count--
			} else {
				count++
			}
			if count == 0 {
				return s[tags[start_index].i : tags[i].j+1]
			}
		} else {
			val, ok := tags[i].attrs["id"]
			if ok {
				if val == id {
					started = true
					start_index = i
					count++
				}
			}
		}

	}
	return ""
}
