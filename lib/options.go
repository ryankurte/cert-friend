package cafriend

type Options struct {
	NewCA NewCA `command:"new-ca" description:"Create a new CA"`

	Version Version `command:"version" description:"Show version and exit"`
	Verbose bool    `short:"v" long:"verbose" description:"Enable verbose logging"`
}

type NewCA struct {
}

type Version struct{}

var version string = "NOT SET"

type General struct {
	KeyType string `short:"t" long:"key-type" description:"Public key type" option:"rsa" option:"ecdsa"`
	KeySize uint   `short:"s" long:"key-size" description:"Private key length`
}
