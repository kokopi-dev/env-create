package pages

type ScpCheckCacheMsg struct{}

type ScpCacheCheckedMsg struct {
	CacheExists bool
	Prefill     *ScpPrefill
}

type ScpPrefill struct {
	Username   string
	Host       string
	RemotePath string
}

type ScpSaveConfigMsg struct {
	Username   string
	Host       string
	RemotePath string
}
