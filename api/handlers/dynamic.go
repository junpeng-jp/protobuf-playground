package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/junpeng.ong/protobuf-playground/api/dto"
	"github.com/junpeng.ong/protobuf-playground/api/serde"
	"google.golang.org/protobuf/encoding/protojson"
)

func HandleMessageDynamically(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(w, "Could not read the request body", http.StatusUnprocessableEntity)
	}

	var container dto.Container
	err = protojson.Unmarshal(body, &container)

	if err != nil {
		http.Error(w, "Could not unmarshal request body as a dto.Container", http.StatusUnprocessableEntity)
	}

	output := make([][]byte, len(container.Objects))

	for _, obj := range container.Objects {
		var text []byte
		var err error

		switch obj.ObjectType {
		case "location":
			text, err = serde.UnmarshalJsonToLocation(body)
		case "error":
			text, err = serde.UnmarshalJsonToError(body)
		default:
			http.Error(w, fmt.Sprintf("Unknown object of type `%s`", obj.ObjectType), http.StatusUnprocessableEntity)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		output = append(output, text)
	}

	w.Write(bytes.Join(output, []byte("\n")))
}
