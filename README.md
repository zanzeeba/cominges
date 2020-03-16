# a test exercise

this application accepts as input a csv file without a header line that has number in each row/col, integers or floats but some row/col locations can be nan (not a number) usually indicated by nan

when the app runs it will locate the nan's and then calculate the average of the non diagonal adjacent squares and write that average to the row/col

to run the application you type main and pass it one or two parameters in. the first parameter is the path to the file and the filename.

The second parameter which is optional is the text of the nan. If not provided nan is used but you could for arguments sake us antidisestablishmentarianism if you were so inclined.

written in go 1.13

# Usage

download the repo and either use the main executable there or recompile. this will give you some sample files to use (see below)

Unit tests can be run go test -v

Assuming compiled to main then

./main -h

will give you the command line parameters the only required one i -fp which is the filepath and file name for the input data

./main -fp=./input_data/input_test_data.csv

will run to the output file

## testing
to run tests:- go test -v

# Todo

* started to add functionality to pass alternative nan 
* improve testing. current testing only checks that the correct average is calculated needs to be extended cover incorrect input. (sorry ran out of time)

# Sample Data

there are several csv files in a subdirectory input_data for you to use


input_test_data.csv is the primary file for testing

# Versioning
Tried out adding semantic versioning
