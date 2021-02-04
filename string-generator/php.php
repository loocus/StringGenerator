<?php

$letter = substr(str_shuffle(str_repeat($x = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",  ceil(4 / strlen($x)))), 1, 4);
$number = substr(str_shuffle(str_repeat($x = "0123456789",  ceil(2 / strlen($x)))), 1, 2);

echo $letter + $number;