# StringGenerator

## Fun project
StringGenerator in all programming languages. The goal of this repo is to generate 4 caracthers and 2 numbers

## Example Output
```
bmCv15
```

## Pull requests 
If you want to take part in this fun little project, don't hesitate to make pull requests !

**Code example in PHP :**
```php
<?php

$letter = substr(str_shuffle(str_repeat($x = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",  ceil(4 / strlen($x)))), 1, 4);
$number = substr(str_shuffle(str_repeat($x = "0123456789",  ceil(2 / strlen($x)))), 1, 2);

echo $letter + $number;
```

## TODO
- ...

# Credits

Thank you very much to the [Virtual Royaume](https://github.com/Virtual-Royaume) organization who inspired me to take on this challenge with their [Sapin](https://github.com/Virtual-Royaume/Sapin) repository!
