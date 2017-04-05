# Steganogify
Security through obscurity via embedding of keys / cyphertext in gif format image data blocks.

![](https://media.giphy.com/media/tivXpQud8dPy0/giphy.gif)

## Encoding

Currently message contents are hex encoded, chunked into RGB values, then a new
gif image is generated with frames of each color; this first pass is to prove
out the feasibility, and determine a strategy for packing the colors into a
carrier gif image.

``


## Encoding Plans
Encoding a given message into a GIF file will eventually follows a few steps in
order to conceal the presence of the payload; obscuring it amongst a list of
normal gif files.

* Incoming message payload btyes are converted into hexadecimal format
* Hex of the payload is chunked into RGB colour values
* colors are added to the carrier gif's colour table. (more on this later)
* the color values mapped to the original message are peppered into the carrier
  gif's image data in a decryptable manner

Currently the gif file size scales drastically with the message size, but
ideally there should be as little impact on the end gif size.. and the message
contents should be packed as efficiently as possible into the color table of the
image.
