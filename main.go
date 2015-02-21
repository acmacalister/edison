package main

import (
	"fmt"
	"log"
	"time"

	"github.com/acmacalister/gatt"
)

func main() {

	srv := gatt.NewServer(gatt.Name("gophergatt"))
	svc := srv.AddService(gatt.MustParseUUID("09fc95c0-c111-11e3-9904-0002a5d5c51b"))

	// Add a read characteristic that prints how many times it has been read
	n := 0
	rchar := svc.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51b"))
	rchar.HandleRead(
		gatt.ReadHandlerFunc(
			func(resp gatt.ReadResponseWriter, req *gatt.ReadRequest) {
				fmt.Fprintf(resp, "count: %d", n)
				n++
			}),
	)

	// Add a write characteristic that logs when written to
	wchar := svc.AddCharacteristic(gatt.MustParseUUID("16fe0d80-c111-11e3-b8c8-0002a5d5c51b"))
	wchar.HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			log.Println("Wrote:", string(data))
			return gatt.StatusSuccess
		})

	// Add a notify characteristic that updates once a second
	nchar := svc.AddCharacteristic(gatt.MustParseUUID("1c927b50-c116-11e3-8a33-0800200c9a66"))
	nchar.HandleNotifyFunc(
		func(r gatt.Request, n gatt.Notifier) {
			go func() {
				count := 0
				for !n.Done() {
					fmt.Fprintf(n, "Count: %d", count)
					count++
					time.Sleep(time.Second)
				}
			}()
		})

	fmt.Println("advertising...")
	// Start the server
	log.Fatal(srv.AdvertiseAndServe())
}

// c := &serial.Config{Name: "/dev/ttyMFD1", Baud: 9600}
// s, err := serial.OpenPort(c)
// if err != nil {
// 	log.Fatal(err)
// }

// go func() {
// 	for {
// 		reader := bufio.NewReader(os.Stdin)
// 		fmt.Print("Enter text: ")
// 		text, _ := reader.ReadString('\n')
// 		_, err := s.Write([]byte(text))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }()

// for {
// 	buf := make([]byte, 128)
// 	n, err := s.Read(buf)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("stuff:", string(buf[:n]))
// }
