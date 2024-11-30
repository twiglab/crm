package msg

func (z *BusinessCircleAuthor) MarshalBinary() ([]byte, error) {
	return z.MarshalMsg(nil)
}
