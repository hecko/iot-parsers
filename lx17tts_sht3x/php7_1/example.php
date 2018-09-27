<?php

function unpack_sigfox_packet($packet) {

    // PHP does not have specific unpack for int16_t in big endian form, therefore a little hack:
    // if uint16_t in big endian returns >= 2^15, then substract 2^16
    $temp = unpack("n", hex2bin(substr($packet, 0, 4))) [1];
    if ($temp >= pow(2, 15)) {
        $temp = $temp - pow(2, 16);
    }

    $out['temperature'] = $temp / 16.0;
    $out['humidity'] = unpack("C", hex2bin(substr($packet, 6, 2))) [1];
    $out['battery'] = unpack("n", hex2bin(substr($packet, 8, 4))) [1];
    return $out;
}

$packets = array(
    "018100280266",
    "FEFC00440258"
);

foreach($packets as $packet) {
    echo ("Unpacking $packet\n");
    $out = unpack_sigfox_packet($packet);
    print_r($out);
}

?>
