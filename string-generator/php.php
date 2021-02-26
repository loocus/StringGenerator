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
