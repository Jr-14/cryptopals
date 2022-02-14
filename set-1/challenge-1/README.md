# Challenge 1
[Convert Hex to Base64](https://cryptopals.com/sets/1/challenges/1)

The string:
```
49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
```

Should produce:
```
SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
```

So go ahead and make that happen. You'll need to use this code for the rest of the exercises.

Cryptopal Rules:
Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.

# Foundings
Wikipedia has been a vital source of quick information and reading up for [Base64](https://en.wikipedia.org/wiki/Base64)
and likewise with [Hexadecimal](https://en.wikipedia.org/wiki/Hexadecimal).


## My explanation at converting hex to base64
I've encountered Base64 encoding and decoding for cert.pem files, but I've never had a chance to simply explain Base64.
Here is my simplified attempt at explaining how to convert hex to base64.

Note: At any point within my explanation, where there exists numbers such as 16 - they will be read as sixteen in
base10. I'll do my best to prefix hexadecimal with 0x. I've searched for a prefix for Base64 but my quick side searches
were not fruitful. I shall prefix Base64 with 0s (Since sixty-four begins with the character 's')

We can dessiminate similarities between Base64 and Base16 (Hexadecimal), in the sense that Base16 has 16 different
symbols for a single digit, and thus likewise with Base64, from it's name, we're confident that Base64 has 64 different
symbols for a single digit. 

In a single hexadecimal digit, we can store 16 unique characters and thus also count up to 16. Here we notice that 2^4
= 16, and thus a single hexadecimal can also be convereted to 4 bits of data.

Likewise, a single Base64 digit can store 64 unique characters and thus also count up to 64. We make the same
observation that 2^6 = 64, and thus a single Base64 digit can be converted to 6 bits of data.


