# Voiceovarian

Voiceovarian is a command-line tool written in Go that utilizes OpenAI's Text-to-Speech (TTS) model to convert written text into spoken audio. This tool allows users to input text interactively and generates corresponding voiceovers in MP3 format. The generated audio files are then played back in real-time.

<center>
<img alt="voiceovarian-like creature in rick and morty but with golang mascot " src="./img/voiceovarian_golang" style="max-width: 100%; height: 500px"/>
</center>

## Installation

Before you begin, make sure you have the following prerequisites installed:

- Go: [https://golang.org/doc/install](https://golang.org/doc/install)
- OpenAI API Key: Get your API key from OpenAI and set it in the environment (`.env`) variable `OPENAI_API_KEY`.

```sh
# Clone the repository
git clone https://github.com/navicstein/voiceovarian.git

# Navigate to the project directory
cd voiceovarian

# Build the executable
go build -o voiceovarian .

# Set the OpenAI API key
echo "OPENAI_API_KEY=your_api_key" >> .env

# Run the tool
./voiceovarian
```

## Key Features:

- Utilizes OpenAI's TTS Model for high-quality speech synthesis.
- Interactive command-line interface for user-friendly input.
- Real-time playback of generated voiceovers.
- Audio files are saved in the "voiceover_intros" directory for future use.

## Usage

Once the tool is running, enter the text you want to convert into spoken audio when prompted. The generated audio file will be saved in the `voiceover_intros` directory and played back in real-time.

```sh
$ Enter a text to speak it: Hello, this is Voiceovarian!
```

## Dependencies

- [github.com/joho/godotenv](github.com/joho/godotenv): For loading environment variables from a file.
- [github.com/rs/zerolog](github.com/rs/zerolog): For logging.
- [github.com/sashabaranov/go-openai](github.com/sashabaranov/go-openai): Go client for OpenAI API.
- [github.com/ebitengine/oto/v3](github.com/ebitengine/oto/v3): For audio playback.
- [github.com/hajimehoshi/go-mp3](github.com/hajimehoshi/go-mp3): MP3 decoding library.
