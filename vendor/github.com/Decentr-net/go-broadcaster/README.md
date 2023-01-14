# broadcaster
![img](https://img.shields.io/github/go-mod/go-version/TessorNetwork/go-broadcaster) ![img](https://img.shields.io/github/v/tag/TessorNetwork/go-broadcaster?label=version)

Package which simplifies broadcasting messages to furya blockchain node

## Example

```
import (
    . "github.com/TessorNetwork/furya/testutil"
    communitytypes "github.com/TessorNetwork/furya/x/community/types"
    "github.com/TessorNetwork/go-broadcaster"
)

func main() {
	b, err := broadcaster.New(Config{
		KeyringRootDir:     "~/.furya",
		KeyringBackend:     "test",
		KeyringPromptInput: "",
		NodeURI:            "http://localhost:26657",
		BroadcastMode:      "sync",
		From:               "jack",
		ChainID:            "local",
	})
	if err != nil {
		panic(err)
	}

	if _, err := b.BroadcastMsg(&communitytypes.MsgFollow{
		Owner: b.From(),
		Whom:  NewAccAddress(),
	}, "follow me back"); err != nil {
		panic(err)
	}
}
```
