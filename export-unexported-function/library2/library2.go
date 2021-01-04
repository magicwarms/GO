package library2

type Student struct {
	Name string
	// grade int // ini struct yang UNEXPORTED, ubah dulu ke huruf besar
	Grade int // ini baru EXPORTED
}
