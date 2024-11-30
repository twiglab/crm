package msg

func (z *BusinessCircleAuthor) MarshalBinary() ([]byte, error) {
	return z.MarshalMsg(nil)
}

func (z *BusinessCirclePayment) MarshalBinary() ([]byte, error) {
	return z.MarshalMsg(nil)
}

func (z *BusinessCircleRefund) MarshalBinary() ([]byte, error) {
	return z.MarshalMsg(nil)
}
