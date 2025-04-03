HashCracker

I'm actively developing this simple rainbow table creation and lookup tool, expanding its functionality as needed for CTF events I participate in. Right now, it's CPU-bound but effective for basic hash-cracking tasks.
Features

•	Generate rainbow tables from common hash functions (unsalted).

•	Search generated rainbow tables for matching hashes.

•	Supports single hash searches or batch hash searches from files.

•	Simple command-line interface (CLI).

Supported hash functions (unsalted):

•	md4

•	md5

•	sha1

•	sha224

•	sha256

•	sha384

•	sha512

•	blake2b_256

•	blake2b_384

•	blake2b_512

•	ripemd160

•	sha3_224

•	sha3_256

•	sha3_384

•	sha3_512

•	sha512_224

•	sha512_256

Usage:

To generate a rainbow table:

HashCracker.exe generate <hash function> <password-dump.txt>

You can also use "g" instead of "generate".

To search a rainbow table for hashes:

HashCracker.exe find <hash> <password-dump-hash.txt>

You can also use "f" instead of "find".

Instead of providing a single hash, you can supply a .txt file containing multiple hashes to search all at once.

Future goals:

•	Salting Support: Currently, there’s no built-in salting. I’ve hard-coded salting for certain CTF events, but plan to add CLI-driven salting support for flexibility.

•	Multi-threaded Brute Force: I currently have a working single-threaded brute force implementation. Future updates will include multi-threaded support to improve performance.

•	Dictionary Attack Mode: I plan to include a direct dictionary attack mode that bypasses rainbow table generation. Although straightforward to implement, this isn't a high-priority feature.

•	GPU Acceleration (long-term goal): Eventually, I'd like to implement GPU acceleration for increased performance. This is a longer-term objective due to the complexity involved and varying hardware and driver compatibility.
