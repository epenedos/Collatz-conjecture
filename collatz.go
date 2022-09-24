import (
	//    "os"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"

	//    "math/rand"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	//    "github.com/go-echarts/go-echarts/v2/types"
)

func main() {

	i := 7

	for {
		r = coll(i)
	}

}

func coll(r int) int {
	if r%2 == 0 {
		r = r % 2
	} else {
		r = r*3 + 1
	}

}
