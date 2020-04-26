<?php

$addr = 'unix:///tmp/cached_service.sock';

$handle = stream_socket_client($addr, $errorCode, $errorMessage, 1);

if ($handle === false) {
    var_dump($errorMessage, $errorCode);
    exit;
}



$result = [
    'method' => 'Cached.GetUserList',
    'params' => [
         //[
         //    'Id' => 2,
         //]
    ],
];

$body= json_encode($result);



$writeResult = fwrite($handle, $body . "\n");
if ($writeResult === false) {
    echo 'Unable to write to socket' . PHP_EOL;
    exit;
}

$readResult = fgets($handle);
if ($readResult === false) {
    echo 'Unable to read from socket' . PHP_EOL;
    exit;
}


echo $readResult . PHP_EOL;
