package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

const (
	AUDIO_DIR    = "voiceover_intros"
	GREETING_MSG = `Welcome to Voiceovarian! Enter text, and we'll convert it into spoken audio in real-time. Type "exit" to quit.`
)

var (
	otoCtx    *oto.Context
	readyChan chan struct{}
)

func run(ctx context.Context, txt string) error {
	if err := os.MkdirAll(AUDIO_DIR, os.ModePerm); err != nil {
		return err
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	res, err := client.CreateSpeech(ctx, openai.CreateSpeechRequest{
		Model: openai.TTSModel1,
		Input: txt,
		Voice: openai.VoiceOnyx,
	})
	if err != nil {
		return err
	}

	defer res.Close()

	buf, err := io.ReadAll(res)
	if err != nil {
		return err
	}

	p := filepath.Join(AUDIO_DIR, fmt.Sprintf("voiceover_%d.mp3", time.Now().Unix()))
	log.Debug().Str("path", p).Msgf("wrote audio sucessfully!")

	if err := os.WriteFile(p, buf, os.ModePerm); err != nil {
		return err
	}

	return playAudio(ctx, p)
}

func playAudio(ctx context.Context, audioPath string) (err error) {
	log.Debug().Str("path", audioPath).Msg("playing audio..")

	fileBytes, err := os.ReadFile(audioPath)
	if err != nil {
		return err
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		return err
	}

	op := new(oto.NewContextOptions)

	op.SampleRate = 25100
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE

	if otoCtx == nil {
		otoCtx, readyChan, err = oto.NewContext(op)
		if err != nil {
			return err
		}
	}

	<-readyChan

	player := otoCtx.NewPlayer(decodedMp3)

	defer func() {
		_ = player.Close()
	}()

	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	return nil
}

func main() {
	ctx := context.Background()
	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = zerolog.New(multi).With().Caller().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	if err := godotenv.Load(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	fmt.Println(GREETING_MSG)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("\n$ Enter a text to speak it: ")

		scanner.Scan()
		txt := scanner.Text()

		if txt == "exit" {
			break
		}

		if err := run(ctx, txt); err != nil {
			log.Fatal().Msg(err.Error())
		}
	}
}
