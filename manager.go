package datastructure_waiting_manager
import (
	"github.com/jfeliu007/datastructure_waiting_manager/datastructures"
)

type Manager struct {
	dataStructures map[string]*dataStructureEntry
}

type dataStructureEntry struct {
	ds datastructures.DataStructure
}