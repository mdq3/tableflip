package main

import (
	"fmt"
	"math/rand"
)

var tableFlippers = []string{
	"(╯°□°）╯︵ ┻━┻",
	"(ノಠ益ಠ)ノ彡┻━┻",
	"┻━┻︵ \\(°□°)/ ︵ ┻━┻",
	"ʕノ•ᴥ•ʔノ ︵ ┻━┻",
	"(┛❍ᴥ❍)┛彡┻━┻",
	"(ノಥ,_｣ಥ)ノ彡┻━┻",
	"(/¯◡ ‿ ◡)/¯ ~ ┻━┻",
	"(┛ಠ_ಠ)┛彡┻━┻",
}

func main() {
	fmt.Println(tableFlippers[rand.Intn(len(tableFlippers))])
}
