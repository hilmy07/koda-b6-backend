package lib

func AcceptedMime(mimetype string, group string) bool {
	switch group {
	case "picture":
		acceptedMimeType := []string{"image/jpeg", "image/jpg", "image/png"}
		for x := range acceptedMimeType {
			if mimetype == acceptedMimeType[x] {
				return true
			}
		}
	case "document":
		acceptedMimeType := []string{"application/pdf", "text/plain"}
		for x := range acceptedMimeType {
			if mimetype == acceptedMimeType[x] {
				return true
			}
		}
	}

	return false
}