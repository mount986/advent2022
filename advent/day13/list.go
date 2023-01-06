package day13

import (
	"bufio"
	"fmt"
	"reflect"
	"strconv"
)

type List struct {
	Values []interface{}
}

func BuildList(scanner *bufio.Scanner) (List, error) {
	var list List

loop:
	for scanner.Scan() {
		c := scanner.Text()

		switch c {
		case "[":
			l, err := BuildList(scanner)
			if err != nil {
				return list, err
			}
			list.Values = append(list.Values, l)
		case ",":
			continue
		case "]":
			break loop
		
		default:
			fAppend := func(c string) error {
				v, err := strconv.Atoi(c)
				if err != nil {
					return err
				}
				list.Values = append(list.Values, v)

				return nil
			}

			for scanner.Scan() {
				b := scanner.Text()

				switch b {
				case ",":
					fAppend(c)
					continue loop
				case "]":
					fAppend(c)
					break loop
				default:
					c = c + b
				}
			}
		}
	}

	return list, nil
}

func NewList(n int) List {
	var list List

	list.Values = append(list.Values, n)

	return list
}

func (l List) Compare(r List) int {
	for index := range l.Values {
		if len(r.Values) == index {
			return 1
		}

		lItem := l.Values[index]
		rItem := r.Values[index]

		if reflect.TypeOf(lItem).Kind() == reflect.Int {
			if reflect.TypeOf(rItem).Kind() == reflect.Int {
				li := reflect.ValueOf(lItem).Interface().(int)
				ri := reflect.ValueOf(rItem).Interface().(int)
				if li > ri {
					return 1
				} else if li < ri {
					return -1
				} else {
					continue
				}
			} else {
				li := reflect.ValueOf(lItem).Interface().(int)
				ri := reflect.ValueOf(rItem).Interface().(List)
				comp := NewList(li).Compare(ri)
				if comp == 0 {
					continue
				} else {
					return comp
				}
			}
		} else {
			if reflect.TypeOf(rItem).Kind() == reflect.Int {
				li := reflect.ValueOf(lItem).Interface().(List)
				ri := reflect.ValueOf(rItem).Interface().(int)
				comp := li.Compare(NewList(ri))
				if comp == 0 {
					continue
				} else {
					return comp
				}
			} else {
				li := reflect.ValueOf(lItem).Interface().(List)
				ri := reflect.ValueOf(rItem).Interface().(List)
				comp := li.Compare(ri)
				if comp == 0 {
					continue
				} else {
					return comp
				}
			}
		}
	}

	if len(r.Values) > len(l.Values) {
		return -1
	}

	return 0
}

func (l List) ToString() string {
	s := "["

	for _, item := range l.Values {		
		if reflect.TypeOf(item).Kind() == reflect.Int {
			s += fmt.Sprint(reflect.ValueOf(item).Interface().(int))
			s += ","
		} else {
			s += reflect.ValueOf(item).Interface().(List).ToString()
			s += ","
		}
	}	

	if len(s) == 1 {
		s += ","
	}
	s = s[:len(s) - 1] + "]"

	return s
}