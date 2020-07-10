package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fancxxy/algo/list"
)

/*

   f1(x) = 5x^2 + 4x^1 + 2
   f2(x) = 5x^1 + 5
   f1(x) + f2(x) = 5x^2 + 9x^1 +7
   f1(x) * f2(x) = 25x^3 + 45x^2 + 30x^1 + 10

*/

type polynomial struct {
	Coefficient int
	Exponent    int
}

func (p *polynomial) String() string {
	if p.Exponent == 0 {
		return strconv.Itoa(p.Coefficient)
	}
	return strconv.Itoa(p.Coefficient) + "x^" + strconv.Itoa(p.Exponent)
}

func addPolynomial(poly1, poly2 *list.List) *list.List {
	var (
		poly3 = list.New()
		node1 = poly1.Front()
		node2 = poly2.Front()
	)

	for node1 != nil && node2 != nil {
		value1, value2 := node1.Value.(*polynomial), node2.Value.(*polynomial)
		if value1.Exponent > value2.Exponent {
			poly3.PushBack(&polynomial{
				Coefficient: value1.Coefficient,
				Exponent:    value1.Exponent,
			})
			node1 = node1.Next()
		} else if value1.Exponent < value2.Exponent {
			poly3.PushBack(&polynomial{
				Coefficient: value2.Coefficient,
				Exponent:    value2.Exponent,
			})
			node2 = node2.Next()
		} else {
			poly3.PushBack(&polynomial{
				Coefficient: value1.Coefficient + value2.Coefficient,
				Exponent:    value1.Exponent,
			})
			node1 = node1.Next()
			node2 = node2.Next()
		}
	}

	for node1 != nil {
		value1 := node1.Value.(*polynomial)
		poly3.PushBack(&polynomial{
			Coefficient: value1.Coefficient,
			Exponent:    value1.Exponent,
		})
		node1 = node1.Next()
	}
	for node2 != nil {
		value2 := node2.Value.(*polynomial)
		poly3.PushBack(&polynomial{
			Coefficient: value2.Coefficient,
			Exponent:    value2.Exponent,
		})
		node2 = node2.Next()
	}

	return poly3
}

func multiplyPolynomial(poly1, poly2 *list.List) *list.List {
	var (
		poly3 = list.New()
		node1 = poly1.Front()
		node2 = poly2.Front()
	)

	for node1 != nil {
		value1 := node1.Value.(*polynomial)
		node2 = poly2.Front()
		for node2 != nil {
			value2 := node2.Value.(*polynomial)
			poly3.PushBack(&polynomial{
				Coefficient: value1.Coefficient * value2.Coefficient,
				Exponent:    value1.Exponent + value2.Exponent,
			})
			node2 = node2.Next()
		}
		node1 = node1.Next()
	}

	curr := poly3.Front()
	for curr != nil && curr.Next() != nil {
		currValue := curr.Value.(*polynomial)
		// dup是可能需要被合并结点的结点
		dup := curr.Next()
		for dup != nil {
			dupValue := dup.Value.(*polynomial)
			if currValue.Exponent == dupValue.Exponent {
				currValue.Coefficient += dupValue.Coefficient
				poly3.Remove(dup)
			}
			dup = dup.Next()
		}
		curr = curr.Next()
	}

	return poly3
}

func formatPolynomial(poly *list.List) string {
	var ret []string
	for node := poly.Front(); node != nil; node = node.Next() {
		ret = append(ret, node.Value.(*polynomial).String())
	}
	return strings.Join(ret, " + ")
}

func main() {
	poly1 := list.New()
	poly1.PushBack(&polynomial{Coefficient: 5, Exponent: 2})
	poly1.PushBack(&polynomial{Coefficient: 4, Exponent: 1})
	poly1.PushBack(&polynomial{Coefficient: 2, Exponent: 0})

	poly2 := list.New()
	poly2.PushBack(&polynomial{Coefficient: 5, Exponent: 1})
	poly2.PushBack(&polynomial{Coefficient: 5, Exponent: 0})

	fmt.Printf("(%s) + (%s) = %s\n",
		formatPolynomial(poly1), formatPolynomial(poly2), formatPolynomial(addPolynomial(poly1, poly2)))
	fmt.Printf("(%s) * (%s) = %s\n",
		formatPolynomial(poly1), formatPolynomial(poly2), formatPolynomial(multiplyPolynomial(poly1, poly2)))
}
