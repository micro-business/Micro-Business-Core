package system

// UUIDGeneratorServiceImpl is the default implementation of UUID generator service
type UUIDGeneratorServiceImpl struct {
}

// GenerateRandomUUID generates random UUID value.
// Returns either the new random generated UUID or an error if something goes wrong.
func (UUIDGeneratorServiceImpl) GenerateRandomUUID() (UUID, error) {
	return RandomUUID()
}
