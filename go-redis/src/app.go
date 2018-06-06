//   See the License for the specific language governing permissions and
//   limitations under the License.
//
package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

var client *redis.Client

func main(){
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	testSetGet()
	testScan()
	tesMGet()
	testHashMap()
}

func testSetGet(){
	fmt.Println("\n- testSetGet() -");

	fmt.Printf("try set someValue := 100500\n")

	client.Set("someValue", "100500", 0);

	val, err := client.Get("someValue").Result();
	if err != nil {
		panic(err)
	} else{
			fmt.Printf("got value: %s", val)
	}

	fmt.Println("\n- end testSetGet() -");
}

func tesMGet(){
	values, err := client.MGet("key10", "key20", "key30").Result()
	if err != nil {
		panic(err)
	} else{
		for key, val := range values {
			fmt.Printf("%d) Val: %s\n", key, val)
		}

	}
}

func testScan(){
	fmt.Println("\n- testScan() -");

	client.FlushDB()
	for i := 0; i < 33; i++ {
		err := client.Set(fmt.Sprintf("key%d", i), "value", 0).Err()
		if err != nil {
			panic(err)
		}
	}

	var cursor uint64
	var n int
	c :=0;
	for {
		var keys []string
		var err error
		keys, cursor, err = client.Scan(cursor, "key2*", 10).Result()
		c++
		if err != nil {
			panic(err)
		}
		n += len(keys)
		if cursor == 0 {
			break
		}
	}

	fmt.Printf("found %d keys. scanned %d times\n", n, c)
}

func testHashMap(){
	fmt.Println("\n- testHashMap() -");

	m := map[string]interface{}{"three": 3, "four": 4, "two": 2}
	client.HMSet("myHashMap", m);

	values, err := client.HMGet("myHashMap", "four", "two", "non-exist").Result();

	if err != nil {
		panic(err)
	}

	for _, val := range values{
		fmt.Printf("value: %v;\n",val)
	}

	fmt.Println("- end testHashMap() -");
}