package helper

import "github.com/jinzhu/copier"

func Copy(src interface{}, dest interface{}) {
	copier.Copy(dest, src)
}
