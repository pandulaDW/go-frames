package series

import "github.com/pandulaDW/go-frames/helpers"

func (s *Series) MemSize() int {
	return helpers.GetRealSizeOf(s.Data)
}
