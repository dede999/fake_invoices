# Invoice Printing simulation

This is an experiment I made in GO to mock the work of querying a hypothetical
database and print the results in a file (currently, it's being written in the
standard output). Other topic I'm learning is how to deal with parallelism and
this is being quite interesting.

## How was the idea born?

In my path to learn GO, I'm seeking ways to master important concepts of the
language, and this is why I'm looking for small or medium-sized projects to
learn  and display this knowledge.

On a previous job, I created a script to generate my punch clock sheet, as I
found very boring to create a new LaTeX file every month from a fixed template.
On the last company I worked with, I had a request to work on a feature about
invoices being created with database resources. So I decided to do a similar
thing.

## How to make it work?

```shell
$ ./invoices_print -h
# Usage of ./invoices_print:
#  -items int
#        The number of invoices to generate (default 20)
#  -parallel
#        Will it be run in parallel mode? (default: FALSE)
#  -projects int
#        The number of projects to generate (default 1)
```

Arguments:
* `items`
* `parallel`
* `projects`

## Roadmap

- [x] Create classes and methods to create invoice items
- [x] Create methods to display the generated data
- [x] Make it work
- [x] Fix problem related to errors being treated on the base files
  - Now, it just panics, and breaks gracefully 
- [ ] Fix over usage of pointers on runners file as suggested by [@Jonss](https://github.com/Jonss)
- [ ] Create tests
- [ ] Implement writing data to files
- [ ] Allow users to pass a project name as parameter 
  - if the user passes 5 to project flag, and passes 2 names, than 7 projects are created 
  - in other words, this option would not overwrite the `projects` flag
- [ ] Allow users to choose if they want to print to Std out our to a file