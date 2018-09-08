<?php
$http = new swoole_http_server("127.0.0.1", 8000);
$http->on('request', function ($request, $response) {
    $json = json_encode($request);
    echo $json;
    $response->end("<h1>Hello Swoole. #".rand(1000, 9999)."</h1>");
});
$http->start();

?>