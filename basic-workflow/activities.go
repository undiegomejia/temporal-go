package workflow

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func AddHair(ctx context.Context, face FaceType) (string, error) {
	nose, err := callHairService("hair", face.Eyes, strconv.Itoa(face.Ears), face.Mouth)
	return nose, err
}

func AddVoice(ctx context.Context, face FaceType) (string, error) {
	nose, err := callVoiceService("voice", face.Nose, face.Hair)
	return nose, err
}

func callHairService(stem string, eyes string, ears string, mouth string) (string, error) {
	base := "http://localhost:9999/" + stem + "?eyes=%s&ears=%s&mouth=%s"
	url := fmt.Sprintf(base, url.QueryEscape(eyes), url.QueryEscape(ears), url.QueryEscape(mouth))
	time.Sleep(9 * time.Second)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	response := string(body)
	status := resp.StatusCode
	if status >= 400 {
		return "", fmt.Errorf("HTTP ERROR %d: %s", status, response)
	}
	return response, nil
}

func callVoiceService(stem string, nose string, hair string) (string, error) {
	base := "http://localhost:9999/" + stem + "?nose=%s&hair=%s"
	url := fmt.Sprintf(base, url.QueryEscape(nose), url.QueryEscape(hair))
	time.Sleep(9 * time.Second)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	response := string(body)
	status := resp.StatusCode
	if status >= 400 {
		return "", fmt.Errorf("HTTP ERROR %d: %s", status, response)
	}
	return response, nil
}
