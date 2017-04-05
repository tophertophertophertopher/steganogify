# Steganogify
Security through obscurity via embedding of keys / cyphertext in gif format image data blocks.

![](https://media.giphy.com/media/tivXpQud8dPy0/giphy.gif)

## Encoding

Currently message contents are hex encoded, chunked into RGB values, then


## Encoding Plans
Encoding a given message into a GIF file will eventually follows a few steps in
order to conceal the presence of the payload; obscuring it amongst a list of
normal gif files.

* Incoming message payload btyes are converted into hexadecimal format
* Hex of the payload is chunked into RGB colour values
* colors are added to the carrier gif's colour table. (more on this later)
* the color values mapped to the original message are peppered into the carrier
  gif's image data in a decryptable manner
