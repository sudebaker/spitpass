# spitpass
Returns human readable passwords with some posible modifications


### NAME:
   - SpitPass - SpitPass

### USAGE:
   - SpitPass [global options] command [command options] [arguments...]

### AUTHOR:
   - Jesus Cifuentes <jcifuentesfonseca@protonmail.com>

### COMMANDS:
   - help, h  Shows a list of commands or help for one command

### GLOBAL OPTIONS:
  - --suffix, -u              Add some random special character at the end of password (default: false)
  - --preffix, -p             Add some random special character at the beginning of password (default: false)
  - --eleet, -e               Convert password to a semi eleet language (default: false)
  - --capfirst, -c            Capitalize the first letter of password (default: false)
  - --caplast, -l             Capitalize the last letter of password (default: false)
  - --text FILE, -t FILE      Read words from FILE
  - --lenght value, -s value  Lenght of password (default: 8)
  - --help, -h                show help
  - --version, -v             print the version


  ### Examples

  ~~~
  spitpass -s 12 -e
  ~~~

  ~~~
  spitpass -u -c
  ~~~