# Experimental FS Concept

Map FS to immutable tree; Any changes are reflected by adding versioning to
files. All files are merkle tree, checksum, diff hash to calculate percent
difference. 

For each file store:
```
type File struct {
  MerkleHash string
  DiffHash string
  Checksum string

  Filename string
  HeadBytes []bytes // INFO: To support head command
  TailBytes []Bytes // INFO: To support tail command

  Full autocomplete using prefix sorting and fuzzy search

  Live updates so file manager wont have tos mash f5

  No actual structure, structure is overlay relationship networks on files
  existing in key file store with several useful indecies

  Support search via file analysis, i.e: generate text to analyze for audio and
  video files
  ...Add attributes to hold generated tags based on sentimement, word freuqnecy
  and so on and the rest of the file should be encrypted and compressed with
  zstd compression alogrithm...

  **interesting feature ideas**

    * Direct hooking of functions onto any file change. No need to scan, just
      every file has hooks. And those hooks are engaged based on event driven
      emit from any interaction with io stream to file

    * Track harddrives maintained by this system. important files or files
      marked important by user are duplicated across physical disks. and
      break up into blocks, and use this to parallel decrypt/decrompress 
      to access data faster instead of opting for giant blobs

    * Built in backup system, register offsite and using all the built in
      checksums, blocks, compression, push blocks offsite:wq


}
```
