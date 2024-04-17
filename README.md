# go-torrent

Go-Torrent is a lightweight and efficient BitTorrent client written in Go, designed to provide fast and concurrent downloading capabilities through the use of goroutines. Guidance on coding this project was from this blog post: https://blog.jse.li/posts/torrent/

### Prerequisites

Make sure you have Go installed on your system. If not, you can download and install it from the [official Go website](https://golang.org/doc/install).

## Install

To get started with GoTorrent, follow these steps:

1. Fork the GoTorrent repository to your GitHub account.
2. Clone the forked repository to your local machine:

```bash
git clone https://github.com/timfemey/gotorrent.git
```

3. Navigate to the project directory:

```bash
cd go-torrent
```

4. Build the project using the following command:

```bash
go build
```

## Usage

Try downloading [Debian](https://cdimage.debian.org/debian-cd/current/amd64/bt-cd/#indexlist)!

```sh
go-torrent debian-10.2.0-amd64-netinst.iso.torrent debian.iso
```

## Limitations

- Only supports `.torrent` files (no magnet links)
- Only supports HTTP trackers
- Does not support multi-file torrents
- Strictly leeches (does not support uploading pieces)

## Contributing

Contributions are welcome! If you find any issues or have ideas for improvements, please open an issue or submit a pull request on GitHub.
