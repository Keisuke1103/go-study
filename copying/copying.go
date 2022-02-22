package copying

import "fmt"

func Copy() {
	inta := 5
	intb := inta
	inta = 7
	fmt.Printf("inta は %d\n", inta)
	fmt.Printf("intb は %d\n", intb)

	intb = 9
	fmt.Printf("intb は %d\n", intb)
	fmt.Printf("inta は %d\n", inta)

	fmt.Printf("intb のアドレスは %p\n", &intb)
	fmt.Printf("inta のアドレスは %p\n", &inta)
}

func Pointout() {
	inta := 5
	adinta := &inta

	// & はポインタ書式（アドレス）
	// * はポインタ書式（実際の値）
	fmt.Printf("adinta の値 %p\n", adinta)
	fmt.Printf("adinta の格納されたアドレスの値 %d\n", *adinta)
	fmt.Printf("inta のアドレス %p\n", &inta)
	fmt.Printf("inta の値 %d\n", inta)
	*adinta = 7
	fmt.Println("アドレス指定で adinta の値変更")
	fmt.Printf("inta の値 %d\n", inta)

}
