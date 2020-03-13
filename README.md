# a test exercise

this application accepts a csv file without header that has number in each row/col, integers or floats but some row/col locations can be nan (not a number) usually indicated by nan

when the app runs it will locate the nan's and then calculate the average of the non diagonal adjacent squares and write that average to the row/col

to run the application you type main and pass it one or two parameters in. the first parameter is the path to the file and the filename.

The second parameter which is optional is the text of the nan. If not provided nan is used but you could for arguments sake us antidisestablishmentarianism if you were so inclined.

# Usage

download the repo and compile. this will give you some sample files to use (see below)

Unit tests can be run go test -v

Assuming compiled to main then

./main -h
will give you the command line parameters the only required one i -fp which is the filepath and file name for the input data

./main -fp=filename -altnan=xxx

# Sample Data

there are several csv files in a subdirectory input_data for you to use

a sample file looks like this

---
|37.454012 |95.071431 |73.199394 |59.865848 |nan |
---
|15.599452 |5.808361  |86.617615 |60.111501 |70.807258
---
|2.058449  |96.990985 |nan       |21.233911 |18.182497
---
|nan       |30.424224 |52.475643 |43.194502 |29.122914
---
|61.185289 |13.949386 |29.214465 |nan       |45.606998
---

