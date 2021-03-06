package problem003

import "testing"

func TestNode_SerializeSimple(t *testing.T) {
	node := Node{
		Val: "root",
		Left: &Node{
			Val:  "left",
			Left: &Node{Val: "left.left"},
		},
		Right: &Node{Val: "right"},
	}

	if serializedNode, err := node.Serialize(); err == nil {
		if deserializedNode, err := serializedNode.Deserialize(); err == nil {
			if deserializedNode.Left.Left.Val != "left.left" {
				t.Log("value mismatched")
				t.Fail()
			}
		} else {
			t.Log(err)
			t.Fail()
		}
	} else {
		t.Log(err)
		t.Fail()
	}
}
