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

function randomString() {
	$result = '';
	$characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz';

	for ($i = 0; $i < 6; $i++) {
		$result .= $characters[rand(0, strlen($characters))];
	}
	return $result;
}

echo randomString();
```

## TODO
- ...

# Credits

Thank you very much to the [Virtual Royaume](https://github.com/Virtual-Royaume) organization who inspired me to take on this challenge with their [Sapin](https://github.com/Virtual-Royaume/Sapin) repository!
