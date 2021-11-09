//go:build windows
// +build windows

package solver

func dialStream(id string) (net.Conn, error) {
	if !strings.HasPrefix(id, unixPrefix) {
		return nil, fmt.Errorf("invalid socket forward key %s", id)
	}

	id = strings.TrimPrefix(id, npipePrefix)
	dur := time.Second
	return winio.DialPipe(id, &dur)
}