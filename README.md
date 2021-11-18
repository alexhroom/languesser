# languesser
Languesser is a Go program that uses the distribution of different letters in a text to guess what language it is written in. 

# Process
Languesser 'learns' from text files fed into it. It stores files for each language, which are distributions of each letter's relative prevalence in each language. When it is given more new files, it will average the prevalence of the new files with that of the old file.

It can then 'guess' by comparing the relative prevalence of letters in a new file with the files for each language, and calculating a similarity score. This guessing process accounts for unfamiliar letters; for example, if a text contains the character 'ß', English or French would be discounted as possible languages (as they do not contain those letters).

# Development
The algorithms here are pretty basic. I'd like to see if I can figure out a better way of calculating similarity scores.

I would also like to be able to improve the process so that it can give accurate guesses for smaller files.