package belajargolangcontext

import (
	"context"
	"log"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	log.Println(background)

	todo := context.TODO()
	log.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "bkecik", "Bbesar")
	contextC := context.WithValue(contextA, "ckecik", "Cbesar")

	contextD := context.WithValue(contextB, "dkecik", "Dbesar")
	contextE := context.WithValue(contextB, "ekecik", "Ebesar")

	contextF := context.WithValue(contextC, "fkecik", "Fbesar")
	contextG := context.WithValue(contextF, "gkecik", "Gbesar")

	log.Println(contextA)
	log.Println(contextB)
	log.Println(contextC)
	log.Println(contextD)
	log.Println(contextE)
	log.Println(contextF)
	log.Println(contextG)

	log.Println(contextF.Value("fkecik"))
	log.Println(contextF.Value("ckecik"))
	log.Println(contextF.Value("bkecik"))

	log.Println("PARENT BETUL", contextA.Value("ekecik"))

}

// ini contoh goroutine leak, go routine nya jalan terus tanpa tahu gimana cara berhenti nya
// nah ctx done itu untuk mengirim context cancel -> (Hentikan process nya)
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				// hanya simulasi saja untuk slow process
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

// buat param context dibawah untuk mengirim param ke function diatas
func TestContextWithCancel(t *testing.T) {
	log.Println("STARTED GOROUTINE TOTAL: ", runtime.NumGoroutine())
	parentCtx := context.Background()
	// ini cancel secara manual
	ctx, cancel := context.WithCancel(parentCtx)

	destination := CreateCounter(ctx)
	// harus nya disini masih 3 goroutine karna masih di process di function CreateCounter
	log.Println("STARTED MIDDLE GOROUTINE TOTAL: ", runtime.NumGoroutine())

	for v := range destination {
		log.Println("Counter: ", v)
		if v == 10 {
			break
		}
	}
	cancel() // mengirim sinyal cancel ke context
	// kenapa dikasih sleep, supaya memberi waktu sinyal cancel context nya bener-bener terkirim
	// kalau gak gitu process nya terlalu cepet karna kan async process
	time.Sleep(2 * time.Second)

	log.Println("ENDED GOROUTINE TOTAL: ", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	log.Println("STARTED GOROUTINE TOTAL: ", runtime.NumGoroutine())
	parentCtx := context.Background()
	// ini cancel dengan timeout
	ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
	// kenapa di kasih cancel lagi? kan timeout udah otomatis tercancel?
	// karna ya best practice nya bagus untuk cancel sekali lagi
	// untuk memastikan process nya bener-bener ter-cancel
	defer cancel()

	destination := CreateCounter(ctx)
	// harus nya disini masih 3 goroutine karna masih di process di function CreateCounter
	log.Println("STARTED MIDDLE GOROUTINE TOTAL: ", runtime.NumGoroutine())
	// disini looping tidak pernah berhenti
	for v := range destination {
		log.Println("Counter: ", v)
	}

	log.Println("ENDED GOROUTINE TOTAL: ", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	log.Println("STARTED GOROUTINE TOTAL: ", runtime.NumGoroutine())
	parentCtx := context.Background()
	// ini cancel dengan deadline
	// kalau deadline kita set waktu nya kapan selesai nya, ada set deadline nya misal nya jam 12 siang atau jam 8 malam
	// sedikit berbeda dari timeout,
	ctx, cancel := context.WithDeadline(parentCtx, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	// harus nya disini masih 3 goroutine karna masih di process di function CreateCounter
	log.Println("STARTED MIDDLE GOROUTINE TOTAL: ", runtime.NumGoroutine())
	// disini looping tidak pernah berhenti
	for v := range destination {
		log.Println("Counter: ", v)
	}

	log.Println("ENDED GOROUTINE TOTAL: ", runtime.NumGoroutine())
}
