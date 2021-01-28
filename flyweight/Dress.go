package flyweight

type Dress interface { // dress接口. 会被T和CT两种dress实现
	GetColor() string
}
