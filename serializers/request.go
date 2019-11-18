package serializers

type AuthorizationHeaderRequest struct {
	Authorization string `header:"Authorization"`
}
