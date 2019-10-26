package problem3

import "encoding/json"

type Node struct {
    Val   string `json:"val"`
    Left  *Node  `json:"left"`
    Right *Node  `json:"right"`
}

type SerializedNode string

func (node Node) Serialize() (*SerializedNode, error) {
    rawResult, err := json.Marshal(node)
    if err != nil {
        return nil, err
    }
    result := SerializedNode(rawResult)
    return &result, nil
}

func (serializedNode SerializedNode) Deserialize() (*Node, error) {
    node := Node{}
    err := json.Unmarshal([]byte(serializedNode), &node)
    if err != nil {
        return nil, err
    }
    return &node, nil
}
